package health

import (
	"errors"
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/horizon"
	"github.com/elastic/elastic-agent-libs/logp"
)

type Endpoint struct {
	Name     string
	Endpoint string
	Fn       func(*MetricSet) ([]mb.Event, error)
}

var (
	endpoints map[string]Endpoint
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {

	endpoints = map[string]Endpoint{
		"ConnectionServers": {Name: "ConnectionServers", Endpoint: "rest/monitor/connection-servers", Fn: getConnectionServers},
		"DesktopPools":      {Name: "DesktopPools", Endpoint: "rest/inventory/v1/desktop-pools", Fn: getDesktopPools},
		"Sessions":          {Name: "Sessions", Endpoint: "rest/inventory/v1/sessions", Fn: getSessions},
		"Gateways":          {Name: "Gateways", Endpoint: "rest/monitor/gateways", Fn: getGateways},
		"VirtualCenters":    {Name: "VirtualCenters", Endpoint: "rest/config/v1/virtual-centers", Fn: getVirtualCenters},
		"Machines":          {Name: "Machines", Endpoint: "rest/inventory/v1/machines", Fn: getMachines},
		"RDSServers":        {Name: "RDSServers", Endpoint: "rest/monitor/rds-servers", Fn: getRDSServers},
		"Farms":             {Name: "Farms", Endpoint: "rest/inventory/v1/farms", Fn: getFarms},
		"CertificateData":   {Name: "CertificateData", Endpoint: "rest/config/v1/connection-servers/certificates", Fn: getCertificateData},
		"LicenseData":       {Name: "LicenseData", Endpoint: "rest/config/v1/licenses", Fn: getLicenseData},
	}

	mb.Registry.MustAddMetricSet("horizon", "health", New)
}

func getEndpoint(name string) (Endpoint, error) {
	endpoint, ok := endpoints[name]
	if !ok {
		return Endpoint{}, fmt.Errorf("%s not found in the map", name)
	}
	return endpoint, nil
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	config        *horizon.Config
	logger        *logp.Logger
	horizonClient *HorizonRestClient
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The horizon health metricset is beta.")

	logger := logp.NewLogger(base.FullyQualifiedName())

	config, err := horizon.NewConfig(base, logger)
	if err != nil {
		logger.Errorf("failed to get load config: %v", err)
		return nil, err
	}

	// Get the session cookie
	ecsClient, err := GetClient(config, base)
	if err != nil {
		logger.Errorf("failed to get a session client: %v", err)
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		config:        config,
		logger:        logger,
		horizonClient: ecsClient,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	// accumulate errs and report them all at the end so that we don't
	// stop processing events if one of the fetches fails
	var errs []error

	err := m.horizonClient.login()
	if err != nil {
		m.logger.Errorf("failed to login: %v", err)
		return err
	}

	for _, endpoint := range endpoints {
		m.logger.Debugf("Calling endpoint %s ....", endpoint.Name)
		events, err := endpoint.Fn(m)
		m.logger.Debugf("Fetched %d %s events", len(events), endpoint.Name)

		if err != nil {
			m.logger.Errorf("Error getting %s events: %s", endpoint.Name, err)
			errs = append(errs, err)
		} else {
			for _, event := range events {
				report.Event(event)
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("error while fetching system metrics: %w", errors.Join(errs...))
	}

	return nil
}
