package health

import (
	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

const (
	metricsetName = "health"
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
		"Cluster Nodes":           {Name: "Cluster Nodes", Endpoint: "/api/v1/cluster/nodes", Fn: getClusterNodesEvents},
		"Cluster Status":          {Name: "Cluster Status", Endpoint: "/api/v1/cluster/status", Fn: getClusterStatusEvents},
		"Edge Clusters":           {Name: "Edge Clusters", Endpoint: "/api/v1/edge-clusters", Fn: getEdgeClustersEvents},
		"Firewall Sections":       {Name: "Firewall Sections", Endpoint: "/api/v1/firewall/sections", Fn: getFirewallSectionsEvents},
		"Logical Router Ports":    {Name: "Logical Router Ports", Endpoint: "/api/v1/logical-router-ports", Fn: getLogicalRouterPortsEvents},
		"Node Network Interfaces": {Name: "Node Network Interfaces", Endpoint: "/api/v1/node/network/interfaces", Fn: getNodeNetworkInterfacesEvents},
		"IP Pools":                {Name: "IP Pools", Endpoint: "/api/v1/pools/ip-pools", Fn: getIPPoolsEvents},
		"Transport Nodes":         {Name: "Transport Nodes", Endpoint: "/api/v1/transport-nodes", Fn: getTransportNodesEvents},
		"Transport Zones":         {Name: "Transport Zones", Endpoint: "/api/v1/transport-zones", Fn: getTransportZonesEvents},
		"Cluster Backup History":  {Name: "Cluster Backup History", Endpoint: "/policy/api/v1/cluster/backups/history", Fn: getClusterBackupHistoryEvents},
		"Infrastructure Tier-0s":  {Name: "Infrastructure Tier-0s", Endpoint: "/policy/api/v1/infra/tier-0s", Fn: getInfraTier0sEvents},
		"Infrastructure Tier-1s":  {Name: "Infrastructure Tier-1s", Endpoint: "/policy/api/v1/infra/tier-1s", Fn: getInfraTier1sEvents},
	}
	mb.Registry.MustAddMetricSet("nsxt", "health", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	counter int
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The nsxt health metricset is beta.")

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		counter:       1,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	report.Event(mb.Event{
		MetricSetFields: mapstr.M{
			"counter": m.counter,
		},
	})
	m.counter++

	return nil
}
