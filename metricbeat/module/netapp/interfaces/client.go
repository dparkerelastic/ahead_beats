package interfaces

import (
	"fmt"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/gosnmp/gosnmp"
)

type NetAppSnmpClient struct {
	config    *netapp.Config
	target    string
	client    *gosnmp.GoSNMP
	community string
	baseOID   string
}

type SNMPResult struct {
	OIDName string `json:"OIDName"`
	OID     string `json:"OID"`
	Value   string `json:"Value"`
}

type OidField struct {
	OID          string
	OIDName      string
	OIDFieldName string
}

var OidToName = map[string]OidField{
	".1.3.6.1.4.1.789.1.22.1.2.1.25": {"1.3.6.1.4.1.789.1.22.1.2.1.25", "InOctets", "in_octets"},
	".1.3.6.1.4.1.789.1.22.1.2.1.26": {"1.3.6.1.4.1.789.1.22.1.2.1.26", "InUcastPkts", "in_ucast_pkts"},
	".1.3.6.1.4.1.789.1.22.1.2.1.27": {"1.3.6.1.4.1.789.1.22.1.2.1.27", "InNUcastPkts", "in_nucast_pkts"},
	".1.3.6.1.4.1.789.1.22.1.2.1.28": {"1.3.6.1.4.1.789.1.22.1.2.1.28", "InDiscards", "in_discards"},
	".1.3.6.1.4.1.789.1.22.1.2.1.29": {"1.3.6.1.4.1.789.1.22.1.2.1.29", "InErrors", "in_errors"},
	".1.3.6.1.4.1.789.1.22.1.2.1.31": {"1.3.6.1.4.1.789.1.22.1.2.1.31", "OutOctets", "out_octets"},
	".1.3.6.1.4.1.789.1.22.1.2.1.32": {"1.3.6.1.4.1.789.1.22.1.2.1.32", "OutUcastPkts", "out_ucast_pkts"},
	".1.3.6.1.4.1.789.1.22.1.2.1.33": {"1.3.6.1.4.1.789.1.22.1.2.1.33", "OutNUcastPkts", "out_nucast_pkts"},
	".1.3.6.1.4.1.789.1.22.1.2.1.34": {"1.3.6.1.4.1.789.1.22.1.2.1.34", "OutDiscards", "out_discards"},
	".1.3.6.1.4.1.789.1.22.1.2.1.35": {"1.3.6.1.4.1.789.1.22.1.2.1.35", "OutErrors", "out_errors"},
}

func GetSnmpClient(config *netapp.Config, base mb.BaseMetricSet) (*NetAppSnmpClient, error) {
	snmp := &gosnmp.GoSNMP{
		Target:    config.Host,
		Port:      config.SnmpPort,
		Community: config.SnmpCommunity,
		Version:   gosnmp.Version2c,
		Timeout:   gosnmp.Default.Timeout,
	}

	client := NetAppSnmpClient{
		config:    config,
		target:    config.Host,
		client:    snmp,
		community: config.SnmpCommunity,
		baseOID:   config.SnmpBaseOID,
	}

	return &client, nil
}

func (c *NetAppSnmpClient) Get() ([]SNMPResult, error) {

	if err := c.client.Connect(); err != nil {

		return nil, fmt.Errorf("error connecting to SNMP target %s: %v", c.target, err)
	}
	defer c.client.Conn.Close()

	var results []SNMPResult
	err := c.client.Walk(c.baseOID, func(variable gosnmp.SnmpPDU) error {
		name := OidToName[variable.Name].OIDName // Lookup the human-readable name
		if name == "" {
			name = "Unknown OID"
		}

		value := valToString(variable.Value)
		results = append(results, SNMPResult{
			OID:   variable.Name,
			Value: value,
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error performing SNMP Walk: %v", err)
	}

	return results, nil
}

func (c *NetAppSnmpClient) GetByOID(oid string) ([]SNMPResult, error) {
	if err := c.client.Connect(); err != nil {
		return nil, fmt.Errorf("error connecting to SNMP target %s: %v", c.target, err)
	}
	defer c.client.Conn.Close()

	var results []SNMPResult
	err := c.client.Walk(oid, func(variable gosnmp.SnmpPDU) error {

		value := valToString(variable.Value)
		results = append(results, SNMPResult{
			OID:   variable.Name,
			Value: value,
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error performing SNMP Walk for OID %s: %v", oid, err)
	}

	return results, nil
}
