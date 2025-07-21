package cluster

import (
	"errors"
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/logp"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("netapp", "cluster", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	config       *netapp.Config
	logger       *logp.Logger
	netappClient *netapp.NetAppRestClient
}

var fullyQualifiedName string
var logger *logp.Logger

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The netapp protocols metricset is beta.")
	fullyQualifiedName = base.FullyQualifiedName()
	logger = logp.NewLogger(fullyQualifiedName)

	config, err := netapp.NewConfig(base, logger)
	if err != nil {
		logger.Errorf("failed to get load config: %v", err)
		return nil, err
	}

	// Get the session cookie
	netappClient, err := netapp.GetClient(config, base)
	if err != nil {
		logger.Errorf("failed to get a session client: %v", err)
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		config:        config,
		netappClient:  netappClient,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	// accumulate errs and report them all at the end so that we don't
	// stop processing events if one of the fetches fails
	var errs []error

	// basic endpoints can all be processed with the ProcessEndpoint function
	for _, endpoint := range basicEndpoints {

		dispatch, ok := endpointDispatchers[endpoint.Name]
		if !ok {
			logger.Errorf("No dispatcher registered for endpoint %s", endpoint.Name)
			continue
		}

		logger.Infof("Calling basic endpoint %s ....", endpoint.Name)
		events, err := dispatch(m.netappClient, endpoint)
		logger.Infof("Fetched %d %s events", len(events), endpoint.Name)

		if err != nil {
			logger.Errorf("Error getting %s events: %s", endpoint.Name, err)
			errs = append(errs, err)
		} else {
			for _, event := range events {
				report.Event(event)
			}
		}
	}
	// custom endpoints are processed with their own GetFunc - mainly to handle unrolling
	// arrays into individual events
	for _, endpoint := range customEndpoints {
		logger.Infof("Calling custom endpoint %s ....", endpoint.Name)
		events, err := endpoint.GetFunc(m.netappClient, endpoint)
		logger.Infof("Fetched %d %s events", len(events), endpoint.Name)

		if err != nil {
			logger.Errorf("Error getting %s events: %s", endpoint.Name, err)
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
