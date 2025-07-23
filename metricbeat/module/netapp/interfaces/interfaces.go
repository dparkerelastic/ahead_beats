package interfaces

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/gosnmp/gosnmp"
)

var (
	interfaceDescriptions     map[int]string
	interfaceDescriptionsErr  error
	initInterfaceDescriptions sync.Once
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("netapp", "interfaces", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	config     *netapp.Config
	logger     *logp.Logger
	snmpClient *NetAppSnmpClient
}

var fullyQualifiedName string
var logger *logp.Logger
var OidToInterfaceName map[int]string

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The netapp interfaces metricset is beta.")
	logger = logp.NewLogger(fullyQualifiedName)

	config, err := netapp.NewConfig(base, logger)
	if err != nil {
		return nil, err
	}

	snmpClient, err := GetSnmpClient(config, base)

	if err != nil {
		logger.Errorf("Failed to get SNMP client: %v", err)
		return nil, err
	}

	OidToInterfaceName, err = loadInterfaceDescriptions(snmpClient)
	if err != nil {
		logger.Errorf("Failed to get interface descriptions: %v", err)
		return nil, err
	}
	return &MetricSet{
		BaseMetricSet: base,
		config:        config,
		logger:        logger,
		snmpClient:    snmpClient,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {

	var results []SNMPResult
	for oid := range OidToName {
		m.logger.Debugf("Processing OID key: %s", oid)
		snmpResults, err := m.snmpClient.GetByOID(oid)
		if err != nil {
			errstr := fmt.Sprintf("failed to get SNMP data for OID %s: %v", oid, err)
			m.logger.Errorf(errstr)
			return fmt.Errorf("%s", errstr)
		}
		results = append(results, snmpResults...)
	}

	events := m.createEvents(results)

	for _, event := range events {
		report.Event(event)
	}

	return nil
}

func (m *MetricSet) createEvents(fields []SNMPResult) []mb.Event {

	interfaceValues := loadResults(fields)

	var events []mb.Event
	for ifName, oids := range interfaceValues {
		for metric, value := range oids {
			event := mb.Event{
				MetricSetFields: mapstr.M{
					"if_name":     ifName,
					"metric_name": metric,
					"value":       value,
				},
				RootFields: netapp.CreateECSFields(),
			}
			events = append(events, event)
		}
	}
	return events
}

type InterfaceOIDValues map[string]map[string]int64

func loadResults(results []SNMPResult) InterfaceOIDValues {
	// Map: interface name -> map[oid prefix]value
	interfaceData := make(InterfaceOIDValues)

	for _, result := range results {
		// Get base oid and last element of OID (interface index)
		oidPrefix, ifIndex, err := basePlusIndex(result.OID)
		if err != nil {
			logger.Warnf("Skipping invalid OID %s: %v", result.OID, err)
			continue // skip invalid OIDs
		}

		ifName, ok := OidToInterfaceName[ifIndex]
		if !ok {
			logger.Warnf("Skipping unknown interface index %d for OID %s", ifIndex, result.OID)
			continue // skip unknown interfaces
		}

		oidField := OidToName[oidPrefix]
		// Create map entry if it doesn't exist
		if _, exists := interfaceData[ifName]; !exists {
			interfaceData[ifName] = make(map[string]int64)
		}

		// Store value by OID prefix
		val, err := strconv.ParseInt(result.Value, 10, 64)
		if err != nil {
			logger.Warnf("Could not convert SNMP value '%s' for OID '%s' to int64: %v", result.Value, result.OID, err)
			continue // skip values that can't be converted
		}
		interfaceData[ifName][oidField.OIDFieldName] = val
	}

	return interfaceData
}

func basePlusIndex(oid string) (string, int, error) {
	parts := strings.Split(oid, ".")
	if len(parts) < 2 {
		return "", 0, fmt.Errorf("invalid OID: %s", oid)
	}
	base := strings.Join(parts[:len(parts)-1], ".")
	indexStr := parts[len(parts)-1]
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid index in OID %s: %w", oid, err)
	}
	return base, index, nil
}

func OIDLastElement(oid string) (int, error) {
	parts := strings.Split(oid, ".")
	ifIndexStr := parts[len(parts)-1]
	ifIndex, err := strconv.Atoi(ifIndexStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ifIndex in OID %s: %w", oid, err)
	}
	return ifIndex, nil
}

// We want to load the interface names during startup to avoid doing for every fetch,
// under the assumption that these values do not change frequently.
// This function is called once to populate the OidToInterfaceName map.
// It uses a sync.Once to ensure it is only called once, though New() gets called more than once.
func loadInterfaceDescriptions(c *NetAppSnmpClient) (map[int]string, error) {

	initInterfaceDescriptions.Do(func() {
		results := make(map[int]string)
		oid := "1.3.6.1.4.1.789.1.22.1.2.1.2"

		if err := c.client.Connect(); err != nil {
			interfaceDescriptionsErr = fmt.Errorf("error connecting to SNMP target %s: %v", c.target, err)
			return
		}
		defer c.client.Conn.Close()

		err := c.client.Walk(oid, func(variable gosnmp.SnmpPDU) error {
			ifIndex, err := OIDLastElement(variable.Name)
			if err != nil {
				return err
			}

			val := valToString(variable.Value)
			results[ifIndex] = val
			return nil
		})
		if err != nil {
			interfaceDescriptionsErr = fmt.Errorf("error performing SNMP Walk: %v", err)
			return
		}

		interfaceDescriptions = results
	})

	return interfaceDescriptions, interfaceDescriptionsErr
}

func valToString(val interface{}) string {
	var value string
	switch v := val.(type) {
	case []byte:
		value = string(v)
	case int:
		value = strconv.Itoa(v)
	case int64:
		value = strconv.FormatInt(v, 10)
	case uint:
		value = strconv.FormatUint(uint64(v), 10)
	case uint64:
		value = strconv.FormatUint(v, 10)
	case string:
		value = v
	default:
		value = fmt.Sprintf("%v", v) // fallback for unexpected types
	}

	return value
}
