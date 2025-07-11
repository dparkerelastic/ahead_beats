package storage

import (
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
)

var basicEndpoints map[string]netapp.Endpoint
var customEndpoints map[string]netapp.Endpoint

const (
	SnapmirrorRelationshipsName = "SnapmirrorRelationships"
	AggregatesName              = "Aggregates"
	DisksName                   = "Disks"
	LUNsName                    = "LUNs"
	QosPoliciesName             = "QosPolicies"
	QtreesName                  = "Qtrees"
	QuotaReportsName            = "QuotaReports"
	QuotaRulesName              = "QuotaRules"
	ShelvesName                 = "Shelves"
	VolumesName                 = "Volumes"
	SvmPeersName                = "SvmPeers"
	SvmsName                    = "Svms"
)

func init() {

	basicEndpoints = map[string]netapp.Endpoint{
		SnapmirrorRelationshipsName: {Name: SnapmirrorRelationshipsName,
			Endpoint:    "/api/snapmirror/relationships",
			QueryFields: SnapMirrorFields,
		},
		AggregatesName: {Name: AggregatesName,
			Endpoint:    "/api/storage/aggregates",
			QueryFields: AggregateFields,
		},
		LUNsName: {Name: LUNsName,
			Endpoint:    "/api/storage/luns",
			QueryFields: LunFields,
		},
		QosPoliciesName: {Name: QosPoliciesName,
			Endpoint:    "/api/storage/qos/policies",
			QueryFields: QosPolicyFields,
		},
		QtreesName: {Name: QtreesName,
			Endpoint:    "/api/storage/qtrees",
			QueryFields: QTreeFields,
		},
		QuotaReportsName: {Name: QuotaReportsName,
			Endpoint:    "/api/storage/quota/reports",
			QueryFields: QuotaReportFields,
		},
		QuotaRulesName: {Name: QuotaRulesName,
			Endpoint:    "/api/storage/quota/rules",
			QueryFields: QuotaRulesFields,
		},
		VolumesName: {Name: VolumesName,
			Endpoint:    "/api/storage/volumes",
			QueryFields: VolumeFields,
		},
		SvmPeersName: {Name: SvmPeersName,
			Endpoint:    "/api/svm/peers",
			QueryFields: SvmPeerFields,
		},
		SvmsName: {Name: SvmsName,
			Endpoint:    "/api/svm/svms",
			QueryFields: SvmFields,
		},
	}

	customEndpoints = map[string]netapp.Endpoint{
		DisksName: {Name: DisksName,
			Endpoint:    "/api/storage/disks",
			GetFunc:     getDisks,
			QueryFields: DiskFields},
		ShelvesName: {Name: ShelvesName,
			Endpoint:    "/api/storage/shelves",
			GetFunc:     getShelves,
			QueryFields: ShelfFields},
	}

}

// For processing basic endpoints we need a type-specific function to create the fields.
// ProcessEndpoint is a generic function that can be used for all basic endpoints,
// but it needs to know how to create the fields for the specific type.
var endpointDispatchers = map[string]netapp.DispatchFunc{
	SnapmirrorRelationshipsName: netapp.MakeDispatchFunc(createSnapMirrorRelationshipFields),
	AggregatesName:              netapp.MakeDispatchFunc(createAggregateFields),
	LUNsName:                    netapp.MakeDispatchFunc(createLUNFields),
	QosPoliciesName:             netapp.MakeDispatchFunc(createQosPolicyFields),
	QtreesName:                  netapp.MakeDispatchFunc(createQTreeFields),
	QuotaReportsName:            netapp.MakeDispatchFunc(createQuotaReportFields),
	QuotaRulesName:              netapp.MakeDispatchFunc(createQuotaRuleFields),
	VolumesName:                 netapp.MakeDispatchFunc(createVolumeFields),
	SvmPeersName:                netapp.MakeDispatchFunc(createSVMPeerFields),
	SvmsName:                    netapp.MakeDispatchFunc(createSVMFields),
}

// Fields to be used in endpoint calls - ONTAP API only returns name and uuid unless fields are specified.
// Further, the NetApp doc discourages using fields=* for various reasons, but also metrics and statistics don't
// get retuned unless explicitly requested.
var QTreeFields = []string{
	"volume",
	"id",
	"svm",
	"name",
	"security_style",
	"unix_permissions",
	"export_policy",
	"path",
	"nas",
	"user",
	"group",
	"metric",
	"statistics",
}

var DiskFields = []string{
	"name",
	"uid",
	"serial_number",
	"model",
	"vendor",
	"firmware_version",
	"usable_size",
	"rated_life_used_percent",
	"type",
	"effective_type",
	"class",
	"container_type",
	"pool",
	"state",
	"node",
	"home_node",
	"aggregates",
	"shelf",
	"local",
	"paths",
	"bay",
	"self_encrypting",
	"fips_certified",
	"bytes_per_sector",
	"sector_count",
	"right_size_sector_count",
	"physical_size",
	"stats",
}

var SvmFields = []string{
	"uuid",
	"name",
	"subtype",
	"language",
	"aggregates",
	"state",
	"comment",
	"ipspace",
	"ip_interfaces",
	"snapshot_policy",
	"nsswitch",
	"nis",
	"ldap",
	"nfs",
	"cifs",
	"iscsi",
	"fcp",
	"nvme",
	"ndmp",
	"s3",
	"certificate",
	"aggregates_delegated",
	"retention_period",
	"max_volumes",
	"anti_ransomware_default_volume_state",
	"is_space_reporting_logical",
	"is_space_enforcement_logical",
	"auto_enable_analytics",
	"auto_enable_activity_tracking",
	"anti_ransomware_auto_switch_from_learning_to_enabled",
	"anti_ransomware_auto_switch_minimum_incoming_data_percent",
	"anti_ransomware_auto_switch_duration_without_new_file_extension",
	"anti_ransomware_auto_switch_minimum_learning_period",
	"anti_ransomware_auto_switch_minimum_file_count",
	"anti_ransomware_auto_switch_minimum_file_extension",
}

var VolumeFields = []string{
	"uuid",
	"comment",
	"create_time",
	"language",
	"name",
	"size",
	"state",
	"style",
	"tiering",
	"cloud_retrieval_policy",
	"type",
	"aggregates",
	"snapshot_count",
	"msid",
	"scheduled_snapshot_naming_scheme",
	"clone",
	"nas",
	"snapshot_locking_enabled",
	"snapshot_policy",
	"svm",
	"space",
	"metric",
	"snapmirror",
	// Analytics not always supported
	// "analytics",
	"activity_tracking",
	"granular_data",
	"granular_data_mode",
}

var QuotaReportFields = []string{
	"files",
	"group",
	"index",
	"qtree",
	"space",
	"svm",
	"type",
	"users",
	"volume",
}

var QuotaRulesFields = []string{
	"files",
	"qtree",
	"space",
	"svm",
	"type",
	"user_mapping",
	"users",
	"uuid",
	"volume",
}

var ShelfFields = []string{
	"uid",
	"name",
	"id",
	"serial_number",
	"model",
	"module_type",
	"internal",
	"local",
	"manufacturer",
	"state",
	"connection_type",
	"disk_count",
	"location_led",
	"paths",
	"bays",
	"frus",
	"ports",
	"fans",
	"temperature_sensors",
	"voltage_sensors",
	"current_sensors",
	"acps",
}

var SnapMirrorFields = []string{
	"backoff_level",
	"consistency_group_failover",
	"destination",
	"exported_snapshot",
	"group_type",
	"healthy",
	"identity_preservation",
	"io_serving_copy",
	"lag_time",
	"last_transfer_network_compression_ratio",
	"last_transfer_type",
	"master_bias_activated_site",
	"policy",
	"preferred_site",
	"restore",
	"source",
	"state",
	"svmdr_volumes",
	"throttle",
	"total_transfer_bytes",
	"total_transfer_duration",
	"transfer",
	"transfer_schedule",
	"unhealthy_reason",
	"uuid",
}

var AggregateFields = []string{
	"uuid",
	"name",
	"node",
	"home_node",
	"snapshot",
	"space",
	"state",
	"snaplock_type",
	"create_time",
	"data_encryption",
	"block_storage",
	"cloud_storage",
	"inactive_data_reporting",
	"inode_attributes",
	"volume_count",
	"metric",
	"statistics",
}

var LunFields = []string{
	"uuid",
	"svm",
	"name",
	"location",
	"class",
	"create_time",
	"enabled",
	"os_type",
	"serial_number",
	"space",
	"status",
	"vvol",
}

var SvmPeerFields = []string{
	"applications",
	"name",
	"peer",
	"state",
	"svm",
	"uuid",
}

var QosPolicyFields = []string{
	"adaptive",
	"fixed",
	"name",
	"object_count",
	"pgid",
	"policy_class",
	"scope",
	"svm",
	"uuid",
}
