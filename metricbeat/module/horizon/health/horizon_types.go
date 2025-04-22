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
