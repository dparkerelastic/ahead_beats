package storage

import (
	"time"

	"github.com/elastic/beats/v7/metricbeat/module/netapp"
)

type StorageMetrics struct {
	Timestamp  time.Time        `json:"timestamp"`
	Duration   string           `json:"duration"`
	Status     string           `json:"status"`
	Throughput netapp.IOLatency `json:"throughput"`
	IOPS       netapp.IOLatency `json:"iops"`
	Latency    netapp.IOLatency `json:"latency"`
}

type StorageStatistics struct {
	Timestamp     time.Time        `json:"timestamp"`
	Status        string           `json:"status"`
	ThroughputRaw netapp.IOLatency `json:"throughput_raw"`
	IOPSRaw       netapp.IOLatency `json:"iops_raw"`
	LatencyRaw    netapp.IOLatency `json:"latency_raw"`
}

// Specific storage objects

type Disk struct {
	Name                 string               `json:"name"`
	UID                  string               `json:"uid"`
	SerialNumber         string               `json:"serial_number"`
	Model                string               `json:"model"`
	Vendor               string               `json:"vendor"`
	FirmwareVersion      string               `json:"firmware_version"`
	UsableSize           int64                `json:"usable_size"`
	RatedLifeUsedPercent int                  `json:"rated_life_used_percent"`
	Type                 string               `json:"type"`
	EffectiveType        string               `json:"effective_type"`
	Class                string               `json:"class"`
	ContainerType        string               `json:"container_type"`
	Pool                 string               `json:"pool"`
	State                string               `json:"state"`
	Node                 netapp.NamedObject   `json:"node"`
	HomeNode             netapp.NamedObject   `json:"home_node"`
	Aggregates           []netapp.NamedObject `json:"aggregates"`
	Shelf                Shelf                `json:"shelf"`
	Local                bool                 `json:"local"`
	Paths                []DiskPath           `json:"paths"`
	Bay                  int                  `json:"bay"`
	SelfEncrypting       bool                 `json:"self_encrypting"`
	FipsCertified        bool                 `json:"fips_certified"`
	BytesPerSector       int64                `json:"bytes_per_sector"`
	SectorCount          int64                `json:"sector_count"`
	RightSizeSectorCount int64                `json:"right_size_sector_count"`
	PhysicalSize         int64                `json:"physical_size"`
	Stats                DiskStats            `json:"stats"`
}

type DiskPath struct {
	DiskPathName string             `json:"disk_path_name"`
	Initiator    string             `json:"initiator"`
	PortName     string             `json:"port_name"`
	PortType     string             `json:"port_type"`
	WWNN         string             `json:"wwnn"`
	WWPN         string             `json:"wwpn"`
	Node         netapp.NamedObject `json:"node"`
}

type DiskStats struct {
	AverageLatency int `json:"average_latency"`
	Throughput     int `json:"throughput"`
	IOPSTotal      int `json:"iops_total"`
	PathErrorCount int `json:"path_error_count"`
	PowerOnHours   int `json:"power_on_hours"`
}

type SVM struct {
	UUID                                string             `json:"uuid"`
	Name                                string             `json:"name"`
	Subtype                             string             `json:"subtype"`
	Language                            string             `json:"language"`
	Aggregates                          []Aggregate        `json:"aggregates"`
	State                               string             `json:"state"`
	Comment                             string             `json:"comment"`
	IPSpace                             netapp.NamedObject `json:"ipspace"`
	IPInterfaces                        []IPInterface      `json:"ip_interfaces"`
	SnapshotPolicy                      netapp.NamedObject `json:"snapshot_policy"`
	NSSwitch                            NSSwitch           `json:"nsswitch"`
	NIS                                 NIS                `json:"nis"`
	LDAP                                LDAP               `json:"ldap"`
	NFS                                 ProtocolStatus     `json:"nfs"`
	CIFS                                ProtocolStatus     `json:"cifs"`
	ISCSI                               ProtocolStatus     `json:"iscsi"`
	FCP                                 ProtocolStatus     `json:"fcp"`
	NVMe                                ProtocolStatus     `json:"nvme"`
	NDMP                                NDMP               `json:"ndmp"`
	S3                                  ProtocolStatus     `json:"s3"`
	Certificate                         Certificate        `json:"certificate"`
	AggregatesDelegated                 bool               `json:"aggregates_delegated"`
	RetentionPeriod                     int                `json:"retention_period"`
	MaxVolumes                          string             `json:"max_volumes"`
	AntiRansomwareDefaultVolumeState    string             `json:"anti_ransomware_default_volume_state"`
	IsSpaceReportingLogical             bool               `json:"is_space_reporting_logical"`
	IsSpaceEnforcementLogical           bool               `json:"is_space_enforcement_logical"`
	AutoEnableAnalytics                 bool               `json:"auto_enable_analytics"`
	AutoEnableActivityTracking          bool               `json:"auto_enable_activity_tracking"`
	AntiRansomwareAutoSwitchEnabled     bool               `json:"anti_ransomware_auto_switch_from_learning_to_enabled"`
	AntiRansomwareAutoSwitchDataPercent int                `json:"anti_ransomware_auto_switch_minimum_incoming_data_percent"`
	AntiRansomwareAutoSwitchNoExtDays   int                `json:"anti_ransomware_auto_switch_duration_without_new_file_extension"`
	AntiRansomwareAutoSwitchMinPeriod   int                `json:"anti_ransomware_auto_switch_minimum_learning_period"`
	AntiRansomwareAutoSwitchMinFiles    int                `json:"anti_ransomware_auto_switch_minimum_file_count"`
	AntiRansomwareAutoSwitchMinExts     int                `json:"anti_ransomware_auto_switch_minimum_file_extension"`
}

type IPInterface struct {
	UUID     string   `json:"uuid"`
	Name     string   `json:"name"`
	IP       IP       `json:"ip"`
	Services []string `json:"services"`
}

type IP struct {
	Address string `json:"address"`
}

type NSSwitch struct {
	Hosts    []string `json:"hosts"`
	Group    []string `json:"group"`
	Passwd   []string `json:"passwd"`
	Netgroup []string `json:"netgroup"`
	NameMap  []string `json:"namemap"`
}

type NIS struct {
	Enabled bool `json:"enabled"`
}

type LDAP struct {
	Enabled bool `json:"enabled"`
}

type ProtocolStatus struct {
	Allowed bool `json:"allowed"`
	Enabled bool `json:"enabled"`
}

type NDMP struct {
	Allowed bool `json:"allowed"`
}

type Certificate struct {
	UUID string `json:"uuid"`
}

type Volume struct {
	UUID                          string               `json:"uuid"`
	Comment                       string               `json:"comment"`
	CreateTime                    time.Time            `json:"create_time"`
	Language                      string               `json:"language"`
	Name                          string               `json:"name"`
	Size                          int64                `json:"size"`
	State                         string               `json:"state"`
	Style                         string               `json:"style"`
	Tiering                       Tiering              `json:"tiering"`
	CloudRetrievalPolicy          string               `json:"cloud_retrieval_policy"`
	Type                          string               `json:"type"`
	Aggregates                    []netapp.NamedObject `json:"aggregates"`
	SnapshotCount                 int                  `json:"snapshot_count"`
	MSID                          int64                `json:"msid"`
	ScheduledSnapshotNamingScheme string               `json:"scheduled_snapshot_naming_scheme"`
	Clone                         CloneInfo            `json:"clone"`
	NAS                           NASInfo              `json:"nas"`
	SnapshotLockingEnabled        bool                 `json:"snapshot_locking_enabled"`
	NamedObject                   netapp.NamedObject   `json:"snapshot_policy"`
	SVM                           netapp.NamedObject   `json:"svm"`
	Space                         VolumeSpace          `json:"space"`
	Metrics                       StorageMetrics       `json:"metric"`
	Snapmirror                    SnapmirrorInfo       `json:"snapmirror"`
	ActivityTracking              ActivityTracking     `json:"activity_tracking"`
	GranularData                  bool                 `json:"granular_data"`
	GranularDataMode              string               `json:"granular_data_mode"`
	// Analytics                     AnalyticsState   `json:"analytics"`
}

type Tiering struct {
	Policy string `json:"policy"`
}

type CloneInfo struct {
	IsFlexclone  bool `json:"is_flexclone"`
	HasFlexclone bool `json:"has_flexclone"`
}

type NASInfo struct {
	GID             int            `json:"gid"`
	SecurityStyle   string         `json:"security_style"`
	UID             int            `json:"uid"`
	UnixPermissions int            `json:"unix_permissions"`
	ExportPolicy    ExportPolicyID `json:"export_policy"`
}

type ExportPolicy struct {
	Name string `json:"name"`
}

type VolumeSpace struct {
	Size                            int64          `json:"size"`
	Available                       int64          `json:"available"`
	Used                            int64          `json:"used"`
	IsUsedStale                     bool           `json:"is_used_stale"`
	BlockStorageInactiveUserData    int64          `json:"block_storage_inactive_user_data"`
	LocalTierFootprint              int64          `json:"local_tier_footprint"`
	Footprint                       int64          `json:"footprint"`
	OverProvisioned                 int64          `json:"over_provisioned"`
	Metadata                        int64          `json:"metadata"`
	TotalFootprint                  int64          `json:"total_footprint"`
	DelayedFreeFootprint            int64          `json:"delayed_free_footprint"`
	FileOperationMetadata           int64          `json:"file_operation_metadata"`
	VolumeGuaranteeFootprint        int64          `json:"volume_guarantee_footprint"`
	EffectiveTotalFootprint         int64          `json:"effective_total_footprint"`
	UserData                        int64          `json:"user_data"`
	UsedByAFS                       int64          `json:"used_by_afs"`
	AvailablePercent                int            `json:"available_percent"`
	AFSTotal                        int64          `json:"afs_total"`
	FullThresholdPercent            int            `json:"full_threshold_percent"`
	NearlyFullThresholdPercent      int            `json:"nearly_full_threshold_percent"`
	OverwriteReserve                int64          `json:"overwrite_reserve"`
	OverwriteReserveUsed            int64          `json:"overwrite_reserve_used"`
	SizeAvailableForSnapshots       int64          `json:"size_available_for_snapshots"`
	PercentUsed                     int            `json:"percent_used"`
	FractionalReserve               int64          `json:"fractional_reserve"`
	BlockStorageInactiveUserDataPct int            `json:"block_storage_inactive_user_data_percent"`
	PhysicalUsedPercent             int            `json:"physical_used_percent"`
	PhysicalUsed                    int64          `json:"physical_used"`
	ExpectedAvailable               int64          `json:"expected_available"`
	FilesystemSize                  int64          `json:"filesystem_size"`
	FilesystemSizeFixed             bool           `json:"filesystem_size_fixed"`
	LargeSizeEnabled                bool           `json:"large_size_enabled"`
	TotalMetadata                   int64          `json:"total_metadata"`
	TotalMetadataFootprint          int64          `json:"total_metadata_footprint"`
	LogicalSpace                    LogicalSpace   `json:"logical_space"`
	Snapshot                        VolumeSnapshot `json:"snapshot"`
}

type VolumeSnapshot struct {
	Used              int64                `json:"used"`
	ReservePercent    int                  `json:"reserve_percent"`
	AutodeleteEnabled bool                 `json:"autodelete_enabled"`
	ReserveSize       int64                `json:"reserve_size"`
	SpaceUsedPercent  int                  `json:"space_used_percent"`
	ReserveAvailable  int64                `json:"reserve_available"`
	AutodeleteTrigger string               `json:"autodelete_trigger"`
	Autodelete        VolumeAutodeleteInfo `json:"autodelete"`
}

type VolumeAutodeleteInfo struct {
	Enabled         bool   `json:"enabled"`
	Trigger         string `json:"trigger"`
	DeleteOrder     string `json:"delete_order"`
	DeferDelete     string `json:"defer_delete"`
	Commitment      string `json:"commitment"`
	TargetFreeSpace int    `json:"target_free_space"`
	Prefix          string `json:"prefix"`
}
type LogicalSpace struct {
	Reporting       bool  `json:"reporting"`
	Enforcement     bool  `json:"enforcement"`
	UsedByAFS       int64 `json:"used_by_afs"`
	UsedPercent     int   `json:"used_percent"`
	Used            int64 `json:"used"`
	UsedBySnapshots int64 `json:"used_by_snapshots"`
}

type SnapmirrorInfo struct {
	IsProtected  bool             `json:"is_protected"`
	Destinations SnapDestinations `json:"destinations"`
}

type SnapDestinations struct {
	IsONTAP bool `json:"is_ontap"`
	IsCloud bool `json:"is_cloud"`
}

type AnalyticsState struct {
	State string `json:"state"`
}

type ActivityTracking struct {
	Supported bool   `json:"supported"`
	State     string `json:"state"`
}

type Snapshot struct {
	Volume           netapp.NamedObject `json:"volume"`
	UUID             string             `json:"uuid"`
	SVM              netapp.NamedObject `json:"svm"`
	Name             string             `json:"name"`
	CreateTime       time.Time          `json:"create_time"`
	SnapmirrorLabel  string             `json:"snapmirror_label"`
	Size             int64              `json:"size"`
	VersionUUID      string             `json:"version_uuid"`
	ProvenanceVolume netapp.NamedObject `json:"provenance_volume"`
	LogicalSize      int64              `json:"logical_size"`
	CompressSavings  int64              `json:"compress_savings"`
	DedupSavings     int64              `json:"dedup_savings"`
	VBN0Savings      int64              `json:"vbn0_savings"`
}

type Qtree struct {
	Volume          netapp.NamedObject `json:"volume"`
	ID              int                `json:"id"`
	SVM             netapp.NamedObject `json:"svm"`
	Name            string             `json:"name"`
	SecurityStyle   string             `json:"security_style"`
	UnixPermissions int                `json:"unix_permissions"`
	ExportPolicy    ExportPolicyID     `json:"export_policy"`
	Path            string             `json:"path"`
	NAS             NASPath            `json:"nas"`
	User            UnixID             `json:"user"`
	Group           UnixID             `json:"group"`
	Metrics         QtreeMetrics       `json:"metricss"`
	Statistics      QtreeStatistics    `json:"statistics"`
}

type ExportPolicyID struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

type NASPath struct {
	Path string `json:"path"`
}

type UnixID struct {
	ID string `json:"id"`
}

type QtreeMetrics struct {
	Duration        string             `json:"duration"`
	IOPS            netapp.IOLatency   `json:"iops"`
	Latency         netapp.IOLatency   `json:"latency"`
	Throughput      netapp.IOLatency   `json:"throughput"`
	Qtree           QtreeRef           `json:"qtree"`
	Status          string             `json:"status"`
	SVM             netapp.NamedObject `json:"svm"`
	MetricTimestamp string             `json:"timestamp"`
	Volume          netapp.NamedObject `json:"volume"`
}

type QtreeStatistics struct {
	IOPSRaw             netapp.IOLatency `json:"iops_raw"`
	LatencyRaw          netapp.IOLatency `json:"latency_raw"`
	Status              string           `json:"status"`
	ThroughputRaw       netapp.IOLatency `json:"throughput_raw"`
	StatisticsTimestamp string           `json:"timestamp"`
}

type QtreeRef struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type QuotaReport struct {
	Files  QuotaUsage           `json:"files"`
	Group  netapp.NamedObject   `json:"group"`
	Index  int                  `json:"index"`
	Qtree  QtreeRef             `json:"qtree"`
	Space  QuotaUsage           `json:"space"`
	SVM    netapp.NamedObject   `json:"svm"`
	Type   string               `json:"type"`
	Users  []netapp.NamedObject `json:"users"`
	Volume netapp.NamedObject   `json:"volume"`
}

type QuotaUsage struct {
	HardLimit int64         `json:"hard_limit"`
	SoftLimit int64         `json:"soft_limit"`
	Used      QuotaUsedInfo `json:"used"`
}

type QuotaUsedInfo struct {
	HardLimitPercent int   `json:"hard_limit_percent"`
	SoftLimitPercent int   `json:"soft_limit_percent"`
	Total            int64 `json:"total"`
}

type QuotaRule struct {
	Files       QuotaLimit     `json:"files"`
	Qtree       QtreeRef       `json:"qtree"`
	Space       QuotaLimit     `json:"space"`
	SVM         SVMNameOnly    `json:"svm"` // FIXME: needed?
	Type        string         `json:"type"`
	UserMapping bool           `json:"user_mapping"`
	Users       []UserOnly     `json:"users"` // FIXME: needed?
	UUID        string         `json:"uuid"`
	Volume      VolumeNameOnly `json:"volume"` // FIXME: needed?
}

type QuotaLimit struct {
	HardLimit int64 `json:"hard_limit"`
	SoftLimit int64 `json:"soft_limit"`
}

type SVMNameOnly struct {
	Name string `json:"name"`
}

type UserOnly struct {
	Name string `json:"name"`
}

type VolumeNameOnly struct {
	Name string `json:"name"`
}

type FailoverInfo struct {
	Error  netapp.Status    `json:"error"`
	State  string           `json:"state"`
	Status SnapmirrorStatus `json:"status"`
	Type   string           `json:"type"`
}

type SnapmirrorStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SnapMirrorEndpoint struct {
	Cluster                 netapp.NamedObject `json:"cluster"`
	ConsistencyGroupVolumes []VolumeNameOnly   `json:"consistency_group_volumes"`
	LUNs                    netapp.NamedObject `json:"luns"`
	Path                    string             `json:"path"`
	SVM                     netapp.NamedObject `json:"svm"`
}

type SnapmirrorPolicy struct {
	Name string `json:"name"`
	Type string `json:"type"`
	UUID string `json:"uuid"`
}

type SnapmirrorTransfer struct {
	BytesTransferred int64  `json:"bytes_transferred"`
	EndTime          string `json:"end_time"`
	LastUpdatedTime  string `json:"last_updated_time"`
	State            string `json:"state"`
	TotalDuration    string `json:"total_duration"`
	Type             string `json:"type"`
	UUID             string `json:"uuid"`
}

type Shelf struct {
	UID            string              `json:"uid"`
	Name           string              `json:"name"`
	ID             string              `json:"id"`
	SerialNumber   string              `json:"serial_number"`
	Model          string              `json:"model"`
	ModuleType     string              `json:"module_type"`
	Internal       bool                `json:"internal"`
	Local          bool                `json:"local"`
	Manufacturer   Manufacturer        `json:"manufacturer"`
	State          string              `json:"state"`
	ConnectionType string              `json:"connection_type"`
	DiskCount      int                 `json:"disk_count"`
	LocationLED    string              `json:"location_led"`
	Paths          []ShelfPath         `json:"paths"`
	Bays           []Bay               `json:"bays"`
	FRUs           []FRU               `json:"frus"`
	Ports          []ShelfPort         `json:"ports"`
	Fans           []Fan               `json:"fans"`
	TempSensors    []TemperatureSensor `json:"temperature_sensors"`
	VoltageSensors []VoltageSensor     `json:"voltage_sensors"`
	CurrentSensors []CurrentSensor     `json:"current_sensors"`
	ACPs           []ACP               `json:"acps"`
}

type Manufacturer struct {
	Name string `json:"name"`
}

type ShelfPath struct {
	Name string             `json:"name"`
	Node netapp.NamedObject `json:"node"`
}

type Bay struct {
	ID      int    `json:"id"`
	HasDisk bool   `json:"has_disk"`
	Type    string `json:"type"`
	State   string `json:"state"`
}

type FRU struct {
	Type            string `json:"type"`
	ID              int    `json:"id"`
	State           string `json:"state"`
	PartNumber      string `json:"part_number"`
	SerialNumber    string `json:"serial_number"`
	FirmwareVersion string `json:"firmware_version"`
	Installed       bool   `json:"installed"`
	PSU             *PSU   `json:"psu,omitempty"`
}

type PSU struct {
	Model       string `json:"model"`
	PowerDrawn  int    `json:"power_drawn"`
	PowerRating int    `json:"power_rating"`
	CrestFactor int    `json:"crest_factor"`
}

type ShelfPort struct {
	ID         int     `json:"id"`
	ModuleID   string  `json:"module_id"`
	Designator string  `json:"designator"`
	State      string  `json:"state"`
	Internal   bool    `json:"internal"`
	WWN        string  `json:"wwn,omitempty"`
	Cable      *Cable  `json:"cable,omitempty"`
	Remote     *Remote `json:"remote,omitempty"`
}

type Cable struct {
	Identifier string `json:"identifier"`
}

type Remote struct {
	WWN string `json:"wwn"`
}

type Fan struct {
	ID        int    `json:"id"`
	Location  string `json:"location"`
	RPM       int    `json:"rpm"`
	State     string `json:"state"`
	Installed bool   `json:"installed"`
}

type TemperatureSensor struct {
	ID          int        `json:"id"`
	Location    string     `json:"location"`
	Temperature int        `json:"temperature"`
	Ambient     bool       `json:"ambient"`
	State       string     `json:"state"`
	Installed   bool       `json:"installed"`
	Threshold   Thresholds `json:"threshold"`
}

type Thresholds struct {
	High SensorLimit `json:"high"`
	Low  SensorLimit `json:"low"`
}

type SensorLimit struct {
	Critical int `json:"critical"`
	Warning  int `json:"warning"`
}

type VoltageSensor struct {
	ID        int     `json:"id"`
	Location  string  `json:"location"`
	Voltage   float64 `json:"voltage"`
	State     string  `json:"state"`
	Installed bool    `json:"installed"`
}

type CurrentSensor struct {
	ID        int    `json:"id"`
	Location  string `json:"location"`
	Current   int    `json:"current"`
	State     string `json:"state"`
	Installed bool   `json:"installed"`
}

type ACP struct {
	Enabled         bool               `json:"enabled"`
	Channel         string             `json:"channel"`
	ConnectionState string             `json:"connection_state"`
	Node            netapp.NamedObject `json:"node"`
}

type SnapMirrorRelationship struct {
	BackoffLevel             string                   `json:"backoff_level"`
	ConsistencyGroupFailover ConsistencyGroupFailover `json:"consistency_group_failover"`
	Destination              SnapMirrorEndpoint       `json:"destination"`
	ExportedSnapshot         string                   `json:"exported_snapshot"`
	GroupType                string                   `json:"group_type"`
	Healthy                  bool                     `json:"healthy"`
	IdentityPreservation     string                   `json:"identity_preservation"`
	IOServingCopy            string                   `json:"io_serving_copy"`
	LagTime                  string                   `json:"lag_time"` //ISO 8601 duration format
	LastTransferNetworkRatio string                   `json:"last_transfer_network_compression_ratio"`
	LastTransferType         string                   `json:"last_transfer_type"`
	MasterBiasActivatedSite  string                   `json:"master_bias_activated_site"`
	Policy                   Policy                   `json:"policy"`
	PreferredSite            string                   `json:"preferred_site"`
	Restore                  bool                     `json:"restore"`
	Source                   SnapMirrorEndpoint       `json:"source"`
	State                    string                   `json:"state"`
	SvmdrVolumes             []NamedVolume            `json:"svmdr_volumes"`
	Throttle                 int                      `json:"throttle"`
	TotalTransferBytes       int64                    `json:"total_transfer_bytes"`
	TotalTransferDuration    string                   `json:"total_transfer_duration"`
	Transfer                 Transfer                 `json:"transfer"`
	TransferSchedule         netapp.NamedObject       `json:"transfer_schedule"`
	UnhealthyReason          []netapp.Status          `json:"unhealthy_reason"`
	UUID                     string                   `json:"uuid"`
}

type ConsistencyGroupFailover struct {
	Error  netapp.Status `json:"error"`
	State  string        `json:"state"`
	Status netapp.Status `json:"status"`
	Type   string        `json:"type"`
}

type NamedVolume struct {
	Name string `json:"name"`
}

type Policy struct {
	Name string `json:"name"`
	Type string `json:"type"`
	UUID string `json:"uuid"`
}

type Transfer struct {
	BytesTransferred int64  `json:"bytes_transferred"`
	EndTime          string `json:"end_time"`
	LastUpdatedTime  string `json:"last_updated_time"`
	State            string `json:"state"`
	TotalDuration    string `json:"total_duration"`
	Type             string `json:"type"`
	UUID             string `json:"uuid"`
}

type Aggregate struct {
	UUID               string                `json:"uuid"`
	Name               string                `json:"name"`
	Node               netapp.NamedObject    `json:"node"`
	HomeNode           netapp.NamedObject    `json:"home_node"`
	Snapshot           AggregateSnapshot     `json:"snapshot"`
	Space              AggregateSpace        `json:"space"`
	State              string                `json:"state"`
	SnaplockType       string                `json:"snaplock_type"`
	CreateTime         time.Time             `json:"create_time"`
	DataEncryption     AggregateEncryption   `json:"data_encryption"`
	BlockStorage       AggregateBlockStorage `json:"block_storage"`
	CloudStorage       AggregateCloudStorage `json:"cloud_storage"`
	InactiveDataReport InactiveDataReport    `json:"inactive_data_reporting"`
	InodeAttributes    InodeAttributes       `json:"inode_attributes"`
	VolumeCount        int                   `json:"volume_count"`
	Metrics            StorageMetrics        `json:"metrics"`
	Statistics         StorageStatistics     `json:"statistics"`
}

type AggregateSnapshot struct {
	FilesTotal        int `json:"files_total"`
	FilesUsed         int `json:"files_used"`
	MaxFilesAvailable int `json:"max_files_available"`
	MaxFilesUsed      int `json:"max_files_used"`
}

type AggregateSpace struct {
	BlockStorage                         BlockStorageSpace `json:"block_storage"`
	Snapshot                             SnapshotSpace     `json:"snapshot"`
	CloudStorage                         CloudStorageSpace `json:"cloud_storage"`
	Efficiency                           Efficiency        `json:"efficiency"`
	EfficiencyWithoutSnapshots           EfficiencySimple  `json:"efficiency_without_snapshots"`
	EfficiencyWithoutSnapshotsFlexclones EfficiencySimple  `json:"efficiency_without_snapshots_flexclones"`
}

type BlockStorageSpace struct {
	Size                                 int64 `json:"size"`
	Available                            int64 `json:"available"`
	Used                                 int64 `json:"used"`
	UsedPercent                          int   `json:"used_percent"`
	FullThresholdPercent                 int   `json:"full_threshold_percent"`
	PhysicalUsed                         int64 `json:"physical_used"`
	PhysicalUsedPercent                  int   `json:"physical_used_percent"`
	DataCompactedCount                   int   `json:"data_compacted_count"`
	DataCompactionSpaceSaved             int64 `json:"data_compaction_space_saved"`
	DataCompactionSpaceSavedPercent      int   `json:"data_compaction_space_saved_percent"`
	VolumeDeduplicationSharedCount       int   `json:"volume_deduplication_shared_count"`
	VolumeDeduplicationSpaceSaved        int64 `json:"volume_deduplication_space_saved"`
	VolumeDeduplicationSpaceSavedPercent int   `json:"volume_deduplication_space_saved_percent"`
}

type SnapshotSpace struct {
	UsedPercent    int   `json:"used_percent"`
	Available      int64 `json:"available"`
	Total          int64 `json:"total"`
	Used           int64 `json:"used"`
	ReservePercent int   `json:"reserve_percent"`
}

type CloudStorageSpace struct {
	Used int64 `json:"used"`
}

type Efficiency struct {
	Savings                        int64   `json:"savings"`
	Ratio                          float64 `json:"ratio"`
	LogicalUsed                    int64   `json:"logical_used"`
	CrossVolumeBackgroundDedupe    bool    `json:"cross_volume_background_dedupe"`
	CrossVolumeInlineDedupe        bool    `json:"cross_volume_inline_dedupe"`
	CrossVolumeDedupeSavings       bool    `json:"cross_volume_dedupe_savings"`
	AutoAdaptiveCompressionSavings bool    `json:"auto_adaptive_compression_savings"`
	EnableWorkloadInformedTSSE     bool    `json:"enable_workload_informed_tsse"`
	WiseTSSEMinUsedCapacityPct     int     `json:"wise_tsse_min_used_capacity_pct"`
}

type EfficiencySimple struct {
	Savings     int64   `json:"savings"`
	Ratio       float64 `json:"ratio"`
	LogicalUsed int64   `json:"logical_used"`
}

type AggregateEncryption struct {
	SoftwareEncryptionEnabled bool `json:"software_encryption_enabled"`
	DriveProtectionEnabled    bool `json:"drive_protection_enabled"`
}

type AggregateBlockStorage struct {
	UsesPartitions bool                 `json:"uses_partitions"`
	StorageType    string               `json:"storage_type"`
	Primary        AggregatePrimary     `json:"primary"`
	HybridCache    AggregateHybridCache `json:"hybrid_cache"`
	Mirror         AggregateMirror      `json:"mirror"`
	Plexes         []AggregatePlex      `json:"plexes"`
}

type AggregatePrimary struct {
	DiskCount     int    `json:"disk_count"`
	DiskClass     string `json:"disk_class"`
	RaidType      string `json:"raid_type"`
	RaidSize      int    `json:"raid_size"`
	ChecksumStyle string `json:"checksum_style"`
	DiskType      string `json:"disk_type"`
}

type AggregateHybridCache struct {
	Enabled bool `json:"enabled"`
}

type AggregateMirror struct {
	Enabled bool   `json:"enabled"`
	State   string `json:"state"`
}

type AggregatePlex struct {
	Name string `json:"name"`
}

type AggregateCloudStorage struct {
	AttachEligible bool                `json:"attach_eligible"`
	Stores         []AggregateStoreRef `json:"stores"`
}

type AggregateStoreRef struct {
	CloudStore netapp.NamedObject `json:"cloud_store"`
	Used       int64              `json:"used"`
}

type InactiveDataReport struct {
	Enabled   bool      `json:"enabled"`
	StartTime time.Time `json:"start_time"`
}

type InodeAttributes struct {
	FilesTotal        int `json:"files_total"`
	FilesUsed         int `json:"files_used"`
	MaxFilesAvailable int `json:"max_files_available"`
	MaxFilesPossible  int `json:"max_files_possible"`
	MaxFilesUsed      int `json:"max_files_used"`
	UsedPercent       int `json:"used_percent"`
}

type QosPolicy struct {
	/*
		From NetApp documentation:
		"A QoS policy defines measurable service level objectives (SLOs) that apply to the storage objects with which the policy is associated. There are two types of policies that can be configured: fixed, which defines a fixed SLO, or adaptive which defines a variable SLO for a storage object. Adaptive policies vary the SLO depending on the space usage of the storage object.
		 A policy can be either a fixed policy or an adaptive one, not both."
		Hence pointers for Adaptive and Fixed - only one will be set at a time.
	*/
	Adaptive    *QosAdaptive       `json:"adaptive,omitempty"`
	Fixed       *QosFixed          `json:"fixed,omitempty"`
	Name        string             `json:"name"`
	ObjectCount int                `json:"object_count"`
	Pgid        int                `json:"pgid"`
	PolicyClass string             `json:"policy_class"`
	Scope       string             `json:"scope"`
	SVM         netapp.NamedObject `json:"svm"`
	UUID        string             `json:"uuid"`
}

type QosAdaptive struct {
	AbsoluteMinIops        int    `json:"absolute_min_iops"`
	BlockSize              string `json:"block_size"`
	ExpectedIops           int    `json:"expected_iops"`
	ExpectedIopsAllocation string `json:"expected_iops_allocation"`
	PeakIops               int    `json:"peak_iops"`
	PeakIopsAllocation     string `json:"peak_iops_allocation"`
}

type QosFixed struct {
	CapacityShared    bool `json:"capacity_shared"`
	MaxThroughputIops int  `json:"max_throughput_iops"`
	MaxThroughputMbps int  `json:"max_throughput_mbps"`
	MinThroughputIops int  `json:"min_throughput_iops"`
	MinThroughputMbps int  `json:"min_throughput_mbps"`
}

type LUN struct {
	UUID         string             `json:"uuid"`
	SVM          netapp.NamedObject `json:"svm"`
	Name         string             `json:"name"`
	Location     LunLocation        `json:"location"`
	Class        string             `json:"class"`
	CreateTime   time.Time          `json:"create_time"`
	Enabled      bool               `json:"enabled"`
	OsType       string             `json:"os_type"`
	SerialNumber string             `json:"serial_number"`
	Space        LunSpace           `json:"space"`
	Status       LunStatus          `json:"status"`
	VVol         LunVVol            `json:"vvol"`
	Metrics      StorageMetrics     `json:"metric"`
	Statistics   StorageStatistics  `json:"statistics"`
}

type LunLocation struct {
	LogicalUnit string             `json:"logical_unit"`
	Node        netapp.NamedObject `json:"node"`
	Volume      netapp.NamedObject `json:"volume"`
}

type LunSpace struct {
	SCSIThinProvisioningSupportEnabled bool         `json:"scsi_thin_provisioning_support_enabled"`
	Size                               int64        `json:"size"`
	Used                               int64        `json:"used"`
	Guarantee                          LunGuarantee `json:"guarantee"`
}

type LunGuarantee struct {
	Requested bool `json:"requested"`
	Reserved  bool `json:"reserved"`
}

type LunStatus struct {
	ContainerState string `json:"container_state"`
	Mapped         bool   `json:"mapped"`
	ReadOnly       bool   `json:"read_only"`
	State          string `json:"state"`
}

type LunVVol struct {
	IsBound  bool         `json:"is_bound"`
	Bindings []LunBinding `json:"bindings"`
}

type LunBinding struct {
	ID          int        `json:"id"`
	Partner     LunPartner `json:"partner"`
	SecondaryID string     `json:"secondary_id"`
}

type LunPartner struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type PeerInfo struct {
	Cluster netapp.NamedObject `json:"cluster"`
	SVM     netapp.NamedObject `json:"svm"`
}

type SVMPeer struct {
	Applications []string           `json:"applications"`
	Name         string             `json:"name"`
	Peer         PeerInfo           `json:"peer"`
	State        string             `json:"state"`
	SVM          netapp.NamedObject `json:"svm"`
	UUID         string             `json:"uuid"`
}
