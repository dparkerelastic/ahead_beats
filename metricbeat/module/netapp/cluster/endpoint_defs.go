package cluster

import (
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
)

var basicEndpoints map[string]netapp.Endpoint
var customEndpoints map[string]netapp.Endpoint

// endpoint: /api/cluster
// endpoint: /api/cluster/nodes
// endpoint: /api/cluster/peers
// endpoint: /api/cluster/sensors?type=battery-life,fan,thermal,voltage
// endpoint: /api/cluster/counter/tables/host_adapter/rows
// endpoint: /api/cluster/counter/tables/system:constituent/rows

const (
	ClusterName        = "Cluster"
	ClusterNodeName    = "ClusterNode"
	SensorsName        = "Sensors"
	AdapterCounterName = "AdapterCounter"
	SystemCounterName  = "SystemCounter"
	PeersName          = "Peers"
)

func init() {

	basicEndpoints = map[string]netapp.Endpoint{
		ClusterName: {Name: ClusterName,
			Endpoint:    "/api/cluster",
			QueryFields: ClusterFields,
		},
		PeersName: {Name: PeersName,
			Endpoint:    "/api/cluster/peers",
			QueryFields: PeerFields,
		},
	}

	customEndpoints = map[string]netapp.Endpoint{
		ClusterNodeName: {Name: ClusterNodeName,
			Endpoint:    "/api/cluster/nodes",
			GetFunc:     getClusterNodes,
			QueryFields: ClusterNodeFields,
		},
		SensorsName: {Name: SensorsName,
			Endpoint:    "/api/cluster/sensors",
			GetFunc:     getSensors,
			QueryFields: SensorsFields,
		},
		AdapterCounterName: {Name: AdapterCounterName,
			Endpoint:    "/api/cluster/counter/tables/host_adapter/rows",
			GetFunc:     getSystemCounters,
			QueryFields: CounterFields,
		},
		SystemCounterName: {Name: SystemCounterName,
			Endpoint:    "/api/cluster/counter/tables/system:constituent/rows",
			GetFunc:     getSystemCounters,
			QueryFields: CounterFields,
		},
	}
}

// // For processing basic endpoints we need a type-specific function to create the fields.
// // ProcessEndpoint is a generic function that can be used for all basic endpoints,
// // but it needs to know how to create the fields for the specific type.

var endpointDispatchers = map[string]netapp.DispatchFunc{
	ClusterName: netapp.MakeDispatchFunc(createClusterFields),
	PeersName:   netapp.MakeDispatchFunc(createPeerFields),
}

// Fields to be used in endpoint calls - ONTAP API only returns name and uuid unless fields are specified.
// Further, the NetApp doc discourages using fields=* for various reasons, but also metrics and statistics don't
// get retuned unless explicitly requested.

// endpoint: /api/protocols/san/iscsi/services
var ClusterFields = []string{
	"uuid",
	"name",
	"serial_number",
	"location",
	"owner",
	"model",
	"system_id",
	"version",
	"date",
	"uptime",
	"state",
	"membership",
	"management_interfaces",
	"cluster_interfaces",
	"storage_configuration",
	"system_aggregate",
	"controller",
	"ha",
	"service_processor",
	"nvram",
	"external_cache",
	"hw_assist",
	"anti_ransomware_version",
	"metric",
	"statistics",
}

var ClusterNodeFields = []string{
	"uuid",
	"name",
	"serial_number",
	"location",
	"owner",
	"model",
	"system_id",
	"version",
	"date",
	"uptime",
	"state",
	"membership",
	"management_interfaces",
	"cluster_interfaces",
	"storage_configuration",
	"system_aggregate",
	"controller",
	"ha",
	"service_processor",
	"nvram",
	"external_cache",
	"hw_assist",
	"anti_ransomware_version",
	"metric",
	"statistics",
}

var SensorsFields = []string{
	"node",
	"index",
	"name",
	"type",
	"value",
	"value_units",
	"threshold_state",
	"critical_low_threshold",
	"warning_low_threshold",
	"warning_high_threshold",
	"critical_high_threshold",
}

var CounterFields = []string{
	"counter_table",
	"id",
	"properties",
	"counters",
}

var PeerFields = []string{
	"authentication",
	"encryption",
	"initial_allowed_svms",
	"ip_address",
	"ipspace",
	"name",
	"peer_applications",
	"remote",
	"status",
	"uuid",
	"version",
}
