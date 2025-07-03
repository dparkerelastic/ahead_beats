package storage

import (
	"fmt"

	"github.com/elastic/beats/v7/metricbeat/mb"
)

type Endpoint struct {
	Name        string
	Endpoint    string
	Fn          func(*MetricSet) ([]mb.Event, error)
	QueryFields []string
}

var endpoints map[string]Endpoint

func init() {
	endpoints = map[string]Endpoint{
		"SnapmirrorRelationships": {Name: "SnapmirrorRelationships",
			Endpoint:    "/api/snapmirror/relationships",
			Fn:          getSnapmirrorRelationships,
			QueryFields: SnapMirrorFields},
		"Aggregates": {Name: "Aggregates",
			Endpoint:    "/api/storage/aggregates",
			Fn:          getAggregates,
			QueryFields: AggregateFields},
		"Disks": {Name: "Disks",
			Endpoint:    "/api/storage/disks",
			Fn:          getDisks,
			QueryFields: DiskFields},
		"LUNs": {Name: "LUNs",
			Endpoint:    "/api/storage/luns",
			Fn:          getLUNs,
			QueryFields: LunFields},
		"QosPolicies": {Name: "QosPolicies",
			Endpoint:    "/api/storage/qos/policies",
			Fn:          getQosPolicies,
			QueryFields: QosPolicyFields},
		"Qtrees": {Name: "Qtrees",
			Endpoint:    "/api/storage/qtrees",
			Fn:          getQtrees,
			QueryFields: QTreeFields},
		"QuotaReports": {Name: "QuotaReports",
			Endpoint:    "/api/storage/quota/reports",
			Fn:          getQuotaReports,
			QueryFields: QuotaReportFields},
		"QuotaRules": {Name: "QuotaRules",
			Endpoint:    "/api/storage/quota/rules",
			Fn:          getQuotaRules,
			QueryFields: QuotaRulesFields},
		"Shelves": {Name: "Shelves",
			Endpoint:    "/api/storage/shelves",
			Fn:          getShelves,
			QueryFields: ShelfFields},
		"Volumes": {Name: "Volumes",
			Endpoint:    "/api/storage/volumes",
			Fn:          getVolumes,
			QueryFields: VolumeFields},
		"SvmPeers": {Name: "SvmPeers",
			Endpoint:    "/api/svm/peers",
			Fn:          getSvmPeers,
			QueryFields: SvmPeerFields},
		"Svms": {Name: "Svms",
			Endpoint:    "/api/svm/svms",
			Fn:          getSvms,
			QueryFields: SvmFields},
	}
}

func getEndpoint(name string) (Endpoint, error) {
	endpoint, ok := endpoints[name]
	if !ok {
		return Endpoint{}, fmt.Errorf("%s not found in the map", name)
	}
	return endpoint, nil
}

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
	"metric.*",
	"statistics.*",
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
	"stats.*",
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
