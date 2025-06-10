package health

// Endpoint: /rest/monitor/connection-servers
type ConnectionServer struct {
	ID                    string            `json:"id"`
	Name                  string            `json:"name"`
	Status                string            `json:"status"`
	ConnectionCount       int               `json:"connection_count"`
	TunnelConnectionCount int               `json:"tunnel_connection_count"`
	DefaultCertificate    bool              `json:"default_certificate"`
	Certificate           Certificate       `json:"certificate"`
	Services              []Service         `json:"services"`
	CSReplications        []CSReplication   `json:"cs_replications"`
	Details               Details           `json:"details"`
	SessionProtocols      []SessionProtocol `json:"session_protocol_data"`
}

// Certificate represents the certificate details.
type Certificate struct {
	Valid     bool  `json:"valid"`
	ValidFrom int64 `json:"valid_from"`
	ValidTo   int64 `json:"valid_to"`
}

// Service represents a service running on the connection server.
type Service struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
}

// CSReplication represents replication details of the connection server.
type CSReplication struct {
	ServerName string `json:"server_name"`
	Status     string `json:"status"`
}

// Details represents additional details about the connection server.
type Details struct {
	Version string `json:"version"`
	Build   string `json:"build"`
}

// SessionProtocol represents session protocol data.
type SessionProtocol struct {
	Protocol     string `json:"session_protocol"`
	SessionCount int    `json:"session_count"`
}

// Endpoint: /rest/inventory/v1/desktop-pools
type DesktopPool struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	DisplayName string       `json:"display_name"`
	Description string       `json:"description,omitempty"`
	Type        string       `json:"type"`
	Source      string       `json:"source"`
	Enabled     bool         `json:"enabled"`
	Settings    PoolSettings `json:"settings"`
}

type PoolSettings struct {
	DeleteInProgress             bool                    `json:"delete_in_progress"`
	EnableClientRestrictions     bool                    `json:"enable_client_restrictions"`
	AllowMultipleSessionsPerUser bool                    `json:"allow_mutilple_sessions_per_user"`
	SessionType                  string                  `json:"session_type"`
	CloudManaged                 bool                    `json:"cloud_managed"`
	CloudAssigned                bool                    `json:"cloud_assigned"`
	SessionSettings              SessionSettings         `json:"session_settings"`
	DisplayProtocolSettings      DisplayProtocolSettings `json:"display_protocol_settings"`
}

type SessionSettings struct {
	PowerPolicy                       string `json:"power_policy"`
	DisconnectedSessionTimeoutPolicy  string `json:"disconnected_session_timeout_policy"`
	DisconnectedSessionTimeoutMinutes int    `json:"disconnected_session_timeout_minutes,omitempty"`
	AllowUsersToResetMachines         bool   `json:"allow_users_to_reset_machines"`
	AllowMultipleSessionsPerUser      bool   `json:"allow_multiple_sessions_per_user"`
	DeleteOrRefreshMachineAfterLogoff string `json:"delete_or_refresh_machine_after_logoff"`
	RefreshOSDiskAfterLogoff          string `json:"refresh_os_disk_after_logoff"`
}

type DisplayProtocolSettings struct {
	DisplayProtocols             []string `json:"display_protocols"`
	DefaultDisplayProtocol       string   `json:"default_display_protocol"`
	AllowUsersToChooseProtocol   bool     `json:"allow_users_to_choose_protocol"`
	HTMLAccessEnabled            bool     `json:"html_access_enabled"`
	SessionCollaborationEnabled  bool     `json:"session_collaboration_enabled"`
	Renderer3D                   string   `json:"renderer3d"`
	GridVGPUsEnabled             bool     `json:"grid_vgpus_enabled"`
	MaxNumberOfMonitors          int      `json:"max_number_of_monitors,omitempty"`
	MaxResolutionOfAnyOneMonitor string   `json:"max_resolution_of_any_one_monitor,omitempty"`
}

// Endpoint: /inventory/v1/desktop-pools/{id}/installed-applications
type InstalledApplication struct {
	Name           string          `json:"name"`
	Version        string          `json:"version,omitempty"`
	Publisher      string          `json:"publisher,omitempty"`
	ExecutablePath string          `json:"executable_path"`
	FileTypes      []FileType      `json:"file_types,omitempty"`
	OtherFileTypes []OtherFileType `json:"other_file_types,omitempty"`
}

// FileType represents a file type associated with an application.
type FileType struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

// OtherFileType represents other file types or protocols associated with an application.
type OtherFileType struct {
	Type        string `json:"type"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Endpoint: /inventory/v1/sessions

type Session struct {
	ID                    string              `json:"id"`
	UserID                string              `json:"user_id"`
	BrokerUserID          string              `json:"broker_user_id"`
	AccessGroupID         string              `json:"access_group_id"`
	MachineID             string              `json:"machine_id"`
	DesktopPoolID         string              `json:"desktop_pool_id"`
	AgentVersion          string              `json:"agent_version"`
	ClientData            ClientData          `json:"client_data"`
	SecurityGatewayData   SecurityGatewayData `json:"security_gateway_data"`
	SessionType           string              `json:"session_type"`
	SessionProtocol       string              `json:"session_protocol"`
	SessionState          string              `json:"session_state"`
	StartTime             int64               `json:"start_time"`
	DisconnectedTime      int64               `json:"disconnected_time,omitempty"`
	LastSessionDurationMs int64               `json:"last_session_duration_ms"`
	ResourcedRemotely     bool                `json:"resourced_remotely"`
	Unauthenticated       bool                `json:"unauthenticated"`
	IdleDuration          int64               `json:"idle_duration,omitempty"`
}

type ClientData struct {
	LocationID string `json:"location_id,omitempty"`
	Type       string `json:"type"`
	Address    string `json:"address,omitempty"`
	Name       string `json:"name,omitempty"`
	Version    string `json:"version,omitempty"`
}

type SecurityGatewayData struct {
	DomainName string `json:"domain_name,omitempty"`
	Address    string `json:"address,omitempty"`
	Location   string `json:"location,omitempty"`
}

// Endpoint: /rest/monitor/gateways
// Gateway represents a gateway in the system.
type Gateway struct {
	ID                    string         `json:"id"`
	Name                  string         `json:"name"`
	Status                string         `json:"status"`
	ActiveConnectionCount int            `json:"active_connection_count"`
	PCoIPConnectionCount  int            `json:"pcoip_connection_count"`
	BlastConnectionCount  int            `json:"blast_connection_count"`
	Details               GatewayDetails `json:"details"`
}

// GatewayDetails represents additional details about a gateway.
type GatewayDetails struct {
	Type     string `json:"type"`
	Address  string `json:"address"`
	Internal bool   `json:"internal"`
	Version  string `json:"version"`
}

// Endpoint: /rest/config/v1/virtual-centers
// VirtualCenter represents a virtual center in the system.
type VirtualCenter struct {
	ID                         string                 `json:"id"`
	Version                    string                 `json:"version"`
	Description                string                 `json:"description"`
	InstanceUUID               string                 `json:"instance_uuid"`
	ServerName                 string                 `json:"server_name"`
	Port                       int                    `json:"port"`
	UseSSL                     bool                   `json:"use_ssl"`
	UserName                   string                 `json:"user_name"`
	SeSparseReclamationEnabled bool                   `json:"se_sparse_reclamation_enabled"`
	Enabled                    bool                   `json:"enabled"`
	VMCDeployment              bool                   `json:"vmc_deployment"`
	Limits                     Limits                 `json:"limits"`
	StorageAcceleratorData     StorageAcceleratorData `json:"storage_accelerator_data"`
	CertificateOverride        CertificateOverride    `json:"certificate_override"`
}

// Limits represents the limits for a virtual center.
type Limits struct {
	ProvisioningLimit                   int `json:"provisioning_limit"`
	PowerOperationsLimit                int `json:"power_operations_limit"`
	InstantCloneEngineProvisioningLimit int `json:"instant_clone_engine_provisioning_limit"`
}

// StorageAcceleratorData represents storage accelerator data for a virtual center.
type StorageAcceleratorData struct {
	Enabled            bool `json:"enabled"`
	DefaultCacheSizeMB int  `json:"default_cache_size_mb"`
}

// CertificateOverride represents certificate override details for a virtual center.
type CertificateOverride struct {
	Certificate string `json:"certificate"`
	Type        string `json:"type"`
}

// Endpoint: /rest/inventory/v1/machines
// Machine represents a machine in the system.
type Machine struct {
	ID                                   string             `json:"id"`
	Name                                 string             `json:"name"`
	DNSName                              string             `json:"dns_name"`
	DesktopPoolID                        string             `json:"desktop_pool_id"`
	State                                string             `json:"state"`
	Type                                 string             `json:"type"`
	OperatingSystem                      string             `json:"operating_system"`
	OperatingSystemArchitecture          string             `json:"operating_system_architecture"`
	AgentVersion                         string             `json:"agent_version,omitempty"`
	AgentBuildNumber                     string             `json:"agent_build_number,omitempty"`
	RemoteExperienceAgentBuildNumber     string             `json:"remote_experience_agent_build_number,omitempty"`
	MessageSecurityMode                  string             `json:"message_security_mode"`
	MessageSecurityEnhancedModeSupported bool               `json:"message_security_enhanced_mode_supported"`
	PairingState                         string             `json:"pairing_state"`
	ConfiguredByConnectionServer         []string           `json:"configured_by_connection_server"`
	UserIDs                              []string           `json:"user_ids,omitempty"`
	ManagedMachineData                   ManagedMachineData `json:"managed_machine_data"`
}

// ManagedMachineData represents additional data for a managed machine.
type ManagedMachineData struct {
	VirtualCenterID          string        `json:"virtual_center_id"`
	HostName                 string        `json:"host_name"`
	Path                     string        `json:"path"`
	VirtualMachinePowerState string        `json:"virtual_machine_power_state"`
	StorageAcceleratorState  string        `json:"storage_accelerator_state,omitempty"`
	MemoryMB                 int           `json:"memory_mb"`
	VirtualDisks             []VirtualDisk `json:"virtual_disks"`
	MissingInVCenter         bool          `json:"missing_in_vcenter"`
	InHoldCustomization      bool          `json:"in_hold_customization"`
	CreateTime               int64         `json:"create_time"`
	InMaintenanceMode        bool          `json:"in_maintenance_mode"`
}

// VirtualDisk represents a virtual disk associated with a machine.
type VirtualDisk struct {
	Path          string `json:"path"`
	DatastorePath string `json:"datastore_path"`
	CapacityMB    int    `json:"capacity_mb"`
}

// Token represents the structure of an access and refresh token.
type LoginToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RDSFarmServer represents a server in an RDS farm.
type RDSServer struct {
	AccessGroupID                        string  `json:"access_group_id"`
	AgentBuildNumber                     int     `json:"agent_build_number"`
	AgentVersion                         float64 `json:"agent_version"`
	BaseVMID                             string  `json:"base_vm_id"`
	BaseVMSnapshotID                     string  `json:"base_vm_snapshot_id"`
	Description                          string  `json:"description"`
	DNSName                              string  `json:"dns_name"`
	Enabled                              bool    `json:"enabled"`
	FarmID                               string  `json:"farm_id"`
	ID                                   string  `json:"id"`
	ImageManagementStreamID              string  `json:"image_management_stream_id"`
	ImageManagementTagID                 string  `json:"image_management_tag_id"`
	LoadIndex                            int     `json:"load_index"`
	LoadPreference                       string  `json:"load_preference"`
	LogoffPolicy                         string  `json:"logoff_policy"`
	MaxSessionsCount                     int     `json:"max_sessions_count"`
	MaxSessionsCountConfigured           int     `json:"max_sessions_count_configured"`
	MaxSessionsType                      string  `json:"max_sessions_type"`
	MaxSessionsTypeConfigured            string  `json:"max_sessions_type_configured"`
	MessageSecurityEnhancedModeSupported bool    `json:"message_security_enhanced_mode_supported"`
	MessageSecurityMode                  string  `json:"message_security_mode"`
	Name                                 string  `json:"name"`
	OperatingSystem                      string  `json:"operating_system"`
	Operation                            string  `json:"operation"`
	OperationState                       string  `json:"operation_state"`
	PendingBaseVMID                      string  `json:"pending_base_vm_id"`
	PendingBaseVMSnapshotID              string  `json:"pending_base_vm_snapshot_id"`
	PendingImageManagementStreamID       string  `json:"pending_image_management_stream_id"`
	PendingImageManagementTagID          string  `json:"pending_image_management_tag_id"`
	RemoteExperienceAgentBuildNumber     int     `json:"remote_experience_agent_build_number"`
	RemoteExperienceAgentVersion         float64 `json:"remote_experience_agent_version"`
	SessionCount                         int     `json:"session_count"`
	State                                string  `json:"state"`
}

// Farm represents a farm in the system.
type Farm struct {
	Description string       `json:"description"`
	DisplayName string       `json:"display_name"`
	Enabled     bool         `json:"enabled"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Settings    FarmSettings `json:"settings"`
	Source      string       `json:"source"`
	Type        string       `json:"type"`
}

// FarmSettings represents the settings for a farm.
type FarmSettings struct {
	DeleteInProgress        bool                        `json:"delete_in_progess"`
	DesktopID               string                      `json:"desktop_id"`
	DisplayProtocolSettings FarmDisplayProtocolSettings `json:"display_protocol_settings"`
	LoadBalancerSettings    LoadBalancerSettings        `json:"load_balancer_settings"`
	ServerErrorThreshold    int                         `json:"server_error_threshold"`
	SessionSettings         FarmSessionSettings         `json:"session_settings"`
}

// FarmDisplayProtocolSettings represents display protocol settings for a farm.
type FarmDisplayProtocolSettings struct {
	AllowDisplayProtocolOverride bool   `json:"allow_display_protocol_override"`
	DefaultDisplayProtocol       string `json:"default_display_protocol"`
	GridVGPUsEnabled             bool   `json:"grid_vgpus_enabled"`
	HTMLAccessEnabled            bool   `json:"html_access_enabled"`
	SessionCollaborationEnabled  bool   `json:"session_collaboration_enabled"`
	VGPUGridProfile              string `json:"vgpu_grid_profile"`
}

// LoadBalancerSettings represents load balancer settings for a farm.
type LoadBalancerSettings struct {
	CustomScriptInUse bool             `json:"custom_script_in_use"`
	LBMetricSettings  LBMetricSettings `json:"lb_metric_settings"`
}

// LBMetricSettings represents load balancer metric settings.
type LBMetricSettings struct {
	CPUThreshold              int  `json:"cpu_threshold"`
	DiskQueueLengthThreshold  int  `json:"disk_queue_length_threshold"`
	DiskReadLatencyThreshold  int  `json:"disk_read_latency_threshold"`
	DiskWriteLatencyThreshold int  `json:"disk_write_latency_threshold"`
	IncludeSessionCount       bool `json:"include_session_count"`
	MemoryThreshold           int  `json:"memory_threshold"`
}

// FarmSessionSettings represents session settings for a farm.
type FarmSessionSettings struct {
	DisconnectedSessionTimeoutMinutes int    `json:"disconnected_session_timeout_minutes"`
	DisconnectedSessionTimeoutPolicy  string `json:"disconnected_session_timeout_policy"`
	EmptySessionTimeoutMinutes        int    `json:"empty_session_timeout_minutes"`
	EmptySessionTimeoutPolicy         string `json:"empty_session_timeout_policy"`
	LogoffAfterTimeout                bool   `json:"logoff_after_timeout"`
	PreLaunchSessionTimeoutMinutes    int    `json:"pre_launch_session_timeout_minutes"`
	PreLaunchSessionTimeoutPolicy     string `json:"pre_launch_session_timeout_policy"`
}

// CertificateDetails represents the details of a certificate.
type CertificateData struct {
	CertificateUsage           string   `json:"certificate_usage"`
	DNSSubjectAlternativeNames []string `json:"dnssubject_alternative_names"`
	InUse                      bool     `json:"in_use"`
	InvalidReasons             []string `json:"invalid_reasons"`
	IsValid                    bool     `json:"is_valid"`
	IssuerName                 string   `json:"issuer_name"`
	SerialNumber               string   `json:"serial_number"`
	SHA1Thumbprint             string   `json:"sha1_thumbprint"`
	SignatureAlgorithm         string   `json:"signature_algorithm"`
	SubjectName                string   `json:"subject_name"`
	ValidFrom                  string   `json:"valid_from"`
	ValidUntil                 string   `json:"valid_until"`
}

// LicenseDetails represents the details of a license.
type LicenseData struct {
	ApplicationPoolLaunchEnabled bool   `json:"application_pool_launch_enabled"`
	DesktopPoolLaunchEnabled     bool   `json:"desktop_pool_launch_enabled"`
	ExpirationTime               int64  `json:"expiration_time"`
	GracePeriodDays              int    `json:"grace_period_days"`
	HelpDeskEnabled              bool   `json:"help_desk_enabled"`
	InstantCloneEnabled          bool   `json:"instant_clone_enabled"`
	LicenseEdition               string `json:"license_edition"`
	LicenseHealth                string `json:"license_health"`
	LicenseKey                   string `json:"license_key"`
	LicenseMode                  string `json:"license_mode"`
	Licensed                     bool   `json:"licensed"`
	SessionCollaborationEnabled  bool   `json:"session_collaboration_enabled"`
	SubscriptionSliceExpiry      int64  `json:"subscription_slice_expiry"`
	UsageModel                   string `json:"usage_model"`
}
