package health

import (
	"encoding/json"
	"fmt"
)

type BackupHistory struct {
	ClusterBackupStatuses   []BackupStatus `json:"cluster_backup_statuses"`
	NodeBackupStatuses      []BackupStatus `json:"node_backup_statuses"`
	InventoryBackupStatuses []BackupStatus `json:"inventory_backup_statuses"`
}

// BackupStatus represents an individual backup status object
type BackupStatus struct {
	BackupID     string `json:"backup_id"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Success      bool   `json:"success"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type TransportNodesResponse struct {
	ResultCount   int             `json:"result_count"`
	Results       []TransportNode `json:"results"`
	SortAscending bool            `json:"sort_ascending"`
	SortBy        string          `json:"sort_by"`
}

type TransportNode struct {
	CreateTime             int64              `json:"_create_time"`
	CreateUser             string             `json:"_create_user"`
	LastModifiedTime       int64              `json:"_last_modified_time"`
	LastModifiedUser       string             `json:"_last_modified_user"`
	Protection             string             `json:"_protection"`
	Revision               int                `json:"_revision"`
	SystemOwned            bool               `json:"_system_owned"`
	ID                     string             `json:"id"`
	NodeID                 string             `json:"node_id"`
	DisplayName            string             `json:"display_name"`
	Description            string             `json:"description,omitempty"`
	FailureDomainID        string             `json:"failure_domain_id,omitempty"`
	IsOverridden           bool               `json:"is_overridden"`
	MaintenanceMode        string             `json:"maintenance_mode"`
	ResourceType           string             `json:"resource_type"`
	Tags                   []Tag              `json:"tags,omitempty"`
	HostSwitchSpec         HostSwitchSpec     `json:"host_switch_spec"`
	TransportZoneEndpoints []TransportZoneRef `json:"transport_zone_endpoints"`
	NodeDeploymentInfo     NodeDeploymentInfo `json:"node_deployment_info"`
}

type Tag struct {
	Scope string `json:"scope"`
	Tag   string `json:"tag"`
}

type HostSwitchSpec struct {
	HostSwitches []HostSwitch `json:"host_switches"`
	ResourceType string       `json:"resource_type"`
}

type HostSwitch struct {
	CPUConfig               []interface{}      `json:"cpu_config"`
	HostSwitchID            string             `json:"host_switch_id"`
	HostSwitchMode          string             `json:"host_switch_mode"`
	HostSwitchName          string             `json:"host_switch_name"`
	HostSwitchProfileIDs    []KeyValue         `json:"host_switch_profile_ids"`
	HostSwitchType          string             `json:"host_switch_type"`
	IPAssignmentSpec        IPAssignmentSpec   `json:"ip_assignment_spec"`
	IsMigratePnics          bool               `json:"is_migrate_pnics"`
	NotReady                bool               `json:"not_ready"`
	Pnics                   []Pnic             `json:"pnics"`
	PnicsUninstallMigration []interface{}      `json:"pnics_uninstall_migration"`
	VmkInstallMigration     []interface{}      `json:"vmk_install_migration"`
	VmkUninstallMigration   []interface{}      `json:"vmk_uninstall_migration"`
	TransportZoneEndpoints  []TransportZoneRef `json:"transport_zone_endpoints"`
	Uplinks                 []Uplink           `json:"uplinks,omitempty"` // Only present in some entries
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type IPAssignmentSpec struct {
	ResourceType   string   `json:"resource_type"`
	IPPoolID       string   `json:"ip_pool_id,omitempty"`
	DefaultGateway string   `json:"default_gateway,omitempty"`
	IPList         []string `json:"ip_list,omitempty"`
	SubnetMask     string   `json:"subnet_mask,omitempty"`
}

type Pnic struct {
	DeviceName string `json:"device_name"`
	UplinkName string `json:"uplink_name"`
}

type Uplink struct {
	UplinkName    string `json:"uplink_name"`
	VdsUplinkName string `json:"vds_uplink_name"`
}

type TransportZoneRef struct {
	TransportZoneID         string                 `json:"transport_zone_id"`
	TransportZoneProfileIDs []TransportZoneProfile `json:"transport_zone_profile_ids"`
}

type TransportZoneProfile struct {
	ProfileID    string `json:"profile_id"`
	ResourceType string `json:"resource_type"`
}

type NodeDeploymentInfo struct {
	CreateTime       int64            `json:"_create_time"`
	CreateUser       string           `json:"_create_user"`
	LastModifiedTime int64            `json:"_last_modified_time"`
	LastModifiedUser string           `json:"_last_modified_user"`
	Protection       string           `json:"_protection"`
	Revision         int              `json:"_revision"`
	SystemOwned      bool             `json:"_system_owned"`
	ResourceType     string           `json:"resource_type"`
	DeploymentType   string           `json:"deployment_type,omitempty"`
	DeploymentConfig DeploymentConfig `json:"deployment_config,omitempty"`
	DisplayName      string           `json:"display_name"`
	Description      string           `json:"description,omitempty"`
	ExternalID       string           `json:"external_id"`
	ID               string           `json:"id"`
	IPAddresses      []string         `json:"ip_addresses"`
	NodeSettings     NodeSettings     `json:"node_settings,omitempty"`
	DiscoveredNodeID string           `json:"discovered_node_id,omitempty"`
	FQDN             string           `json:"fqdn,omitempty"`
	ManagedByServer  string           `json:"managed_by_server,omitempty"`
	DiscoveredIPs    []string         `json:"discovered_ip_addresses,omitempty"`
	OSType           string           `json:"os_type,omitempty"`
	OSVersion        string           `json:"os_version,omitempty"`
}

type DeploymentConfig struct {
	FormFactor       string             `json:"form_factor"`
	NodeUserSettings NodeUserSettings   `json:"node_user_settings"`
	VMDeployment     VMDeploymentConfig `json:"vm_deployment_config"`
}

type NodeUserSettings struct {
	AuditUsername string `json:"audit_username"`
	CLIUsername   string `json:"cli_username"`
}

type VMDeploymentConfig struct {
	AllowSSHRootLogin       bool               `json:"allow_ssh_root_login"`
	ComputeFolderID         string             `json:"compute_folder_id"`
	ComputeID               string             `json:"compute_id"`
	DataNetworkIDs          []string           `json:"data_network_ids"`
	DefaultGatewayAddresses []string           `json:"default_gateway_addresses"`
	DNSServers              []string           `json:"dns_servers"`
	EnableSSH               bool               `json:"enable_ssh"`
	HostID                  string             `json:"host_id"`
	Hostname                string             `json:"hostname"`
	ManagementNetworkID     string             `json:"management_network_id"`
	ManagementPortSubnets   []Subnet           `json:"management_port_subnets"`
	NTPServers              []string           `json:"ntp_servers"`
	PlacementType           string             `json:"placement_type"`
	ReservationInfo         ReservationInfo    `json:"reservation_info"`
	ResourceAllocation      ResourceAllocation `json:"resource_allocation"`
	SearchDomains           []string           `json:"search_domains"`
	StorageID               string             `json:"storage_id"`
	VCID                    string             `json:"vc_id"`
}

type Subnet struct {
	IPAddresses  []string `json:"ip_addresses"`
	PrefixLength int      `json:"prefix_length"`
}

type ReservationInfo struct {
	CPUReservation    CPUReservation    `json:"cpu_reservation"`
	MemoryReservation MemoryReservation `json:"memory_reservation"`
}

type CPUReservation struct {
	ReservationMHz    int    `json:"reservation_in_mhz"`
	ReservationShares string `json:"reservation_in_shares"`
}

type MemoryReservation struct {
	ReservationPercent int `json:"reservation_percentage"`
}

type ResourceAllocation struct {
	CPUCount           int `json:"cpu_count"`
	MemoryAllocationMB int `json:"memory_allocation_in_mb"`
}

type NodeSettings struct {
	AllowSSHRootLogin bool     `json:"allow_ssh_root_login"`
	DNSServers        []string `json:"dns_servers"`
	EnableSSH         bool     `json:"enable_ssh"`
	Hostname          string   `json:"hostname"`
	NTPServers        []string `json:"ntp_servers"`
	SearchDomains     []string `json:"search_domains"`
}

type ClusterStatus struct {
	ClusterID            string                `json:"cluster_id"`
	ControlClusterStatus ControlClusterStatus  `json:"control_cluster_status"`
	DetailedStatus       DetailedClusterStatus `json:"detailed_cluster_status"`
	MgmtClusterStatus    MgmtClusterStatus     `json:"mgmt_cluster_status"`
}

type ControlClusterStatus struct {
	Status string `json:"status"`
}

type DetailedClusterStatus struct {
	ClusterID     string         `json:"cluster_id"`
	Groups        []ClusterGroup `json:"groups"`
	OverallStatus string         `json:"overall_status"`
}

type ClusterGroup struct {
	GroupID     string        `json:"group_id"`
	GroupStatus string        `json:"group_status"`
	GroupType   string        `json:"group_type"`
	Leaders     []Leader      `json:"leaders"`
	Members     []GroupMember `json:"members"`
}

type Leader struct {
	LeaderUUID   string `json:"leader_uuid"`
	LeaseVersion int64  `json:"lease_version"`
	ServiceName  string `json:"service_name"`
}

type GroupMember struct {
	FQDN   string `json:"member_fqdn"`
	IP     string `json:"member_ip"`
	Status string `json:"member_status"`
	UUID   string `json:"member_uuid"`
}

type MgmtClusterStatus struct {
	OnlineNodes []MgmtNode `json:"online_nodes"`
	Status      string     `json:"status"`
}

type MgmtNode struct {
	IP   string `json:"mgmt_cluster_listen_ip_address"`
	UUID string `json:"uuid"`
}

type Tier0ListResponse struct {
	ResultCount   int     `json:"result_count"`
	Results       []Tier0 `json:"results"`
	SortAscending bool    `json:"sort_ascending"`
	SortBy        string  `json:"sort_by"`
}

type Tier0 struct {
	CreateTime             int64          `json:"_create_time"`
	CreateUser             string         `json:"_create_user"`
	LastModifiedTime       int64          `json:"_last_modified_time"`
	LastModifiedUser       string         `json:"_last_modified_user"`
	Protection             string         `json:"_protection"`
	Revision               int            `json:"_revision"`
	SystemOwned            bool           `json:"_system_owned"`
	ID                     string         `json:"id"`
	DisplayName            string         `json:"display_name"`
	Description            string         `json:"description,omitempty"`
	ResourceType           string         `json:"resource_type"`
	Path                   string         `json:"path"`
	ParentPath             string         `json:"parent_path"`
	RelativePath           string         `json:"relative_path"`
	MarkedForDelete        bool           `json:"marked_for_delete"`
	Overridden             bool           `json:"overridden"`
	DefaultRuleLogging     bool           `json:"default_rule_logging"`
	DisableFirewall        bool           `json:"disable_firewall"`
	ForceWhitelisting      bool           `json:"force_whitelisting"`
	FailoverMode           string         `json:"failover_mode"`
	HAMode                 string         `json:"ha_mode"`
	UniqueID               string         `json:"unique_id"`
	AdvancedConfig         AdvancedConfig `json:"advanced_config"`
	InternalTransitSubnets []string       `json:"internal_transit_subnets"`
	TransitSubnets         []string       `json:"transit_subnets"`
	IPv6ProfilePaths       []string       `json:"ipv6_profile_paths"`
	Tags                   []Tag          `json:"tags"`
}

type AdvancedConfig struct {
	Connectivity      string `json:"connectivity"`
	ForwardingUpTimer int    `json:"forwarding_up_timer"`
}

type Tier1ListResponse struct {
	ResultCount   int     `json:"result_count"`
	Results       []Tier1 `json:"results"`
	SortAscending bool    `json:"sort_ascending"`
	SortBy        string  `json:"sort_by"`
}

type Tier1 struct {
	CreateTime              int64    `json:"_create_time"`
	CreateUser              string   `json:"_create_user"`
	LastModifiedTime        int64    `json:"_last_modified_time"`
	LastModifiedUser        string   `json:"_last_modified_user"`
	Protection              string   `json:"_protection"`
	Revision                int      `json:"_revision"`
	SystemOwned             bool     `json:"_system_owned"`
	ID                      string   `json:"id"`
	DisplayName             string   `json:"display_name"`
	Description             string   `json:"description,omitempty"`
	ResourceType            string   `json:"resource_type"`
	Path                    string   `json:"path"`
	ParentPath              string   `json:"parent_path"`
	RelativePath            string   `json:"relative_path"`
	MarkedForDelete         bool     `json:"marked_for_delete"`
	Overridden              bool     `json:"overridden"`
	DefaultRuleLogging      bool     `json:"default_rule_logging"`
	DisableFirewall         bool     `json:"disable_firewall"`
	ForceWhitelisting       bool     `json:"force_whitelisting"`
	FailoverMode            string   `json:"failover_mode"`
	EnableStandbyRelocation bool     `json:"enable_standby_relocation"`
	PoolAllocation          string   `json:"pool_allocation"`
	IPv6ProfilePaths        []string `json:"ipv6_profile_paths"`
	RouteAdvertisementTypes []string `json:"route_advertisement_types"`
	Tier0Path               string   `json:"tier0_path"`
	UniqueID                string   `json:"unique_id"`
	Tags                    []Tag    `json:"tags"`
}

type EdgeClustersResponse struct {
	ResultCount int           `json:"result_count"`
	Results     []EdgeCluster `json:"results"`
}

type EdgeCluster struct {
	CreateTime                int64                   `json:"_create_time"`
	CreateUser                string                  `json:"_create_user"`
	LastModifiedTime          int64                   `json:"_last_modified_time"`
	LastModifiedUser          string                  `json:"_last_modified_user"`
	Protection                string                  `json:"_protection"`
	Revision                  int                     `json:"_revision"`
	SystemOwned               bool                    `json:"_system_owned"`
	ID                        string                  `json:"id"`
	DisplayName               string                  `json:"display_name"`
	Description               string                  `json:"description,omitempty"`
	DeploymentType            string                  `json:"deployment_type"`
	EnableInterSiteForwarding bool                    `json:"enable_inter_site_forwarding"`
	MemberNodeType            string                  `json:"member_node_type"`
	Members                   []EdgeClusterMember     `json:"members"`
	ClusterProfileBindings    []ClusterProfileBinding `json:"cluster_profile_bindings"`
	AllocationRules           []interface{}           `json:"allocation_rules"` // Currently empty; adjust if needed
	ResourceType              string                  `json:"resource_type"`
	Tags                      []Tag                   `json:"tags"`
}

type EdgeClusterMember struct {
	MemberIndex     int    `json:"member_index"`
	TransportNodeID string `json:"transport_node_id"`
}

type ClusterProfileBinding struct {
	ProfileID    string `json:"profile_id"`
	ResourceType string `json:"resource_type"`
}

type TransportZonesResponse struct {
	ResultCount   int             `json:"result_count"`
	Results       []TransportZone `json:"results"`
	SortAscending bool            `json:"sort_ascending"`
	SortBy        string          `json:"sort_by"`
}

type TransportZone struct {
	CreateTime               int64                  `json:"_create_time"`
	CreateUser               string                 `json:"_create_user"`
	LastModifiedTime         int64                  `json:"_last_modified_time"`
	LastModifiedUser         string                 `json:"_last_modified_user"`
	Protection               string                 `json:"_protection"`
	Revision                 int                    `json:"_revision"`
	Schema                   string                 `json:"_schema"`
	SystemOwned              bool                   `json:"_system_owned"`
	ID                       string                 `json:"id"`
	DisplayName              string                 `json:"display_name"`
	HostSwitchID             string                 `json:"host_switch_id"`
	HostSwitchName           string                 `json:"host_switch_name"`
	HostSwitchMode           string                 `json:"host_switch_mode"`
	IsDefault                bool                   `json:"is_default"`
	NestedNSX                bool                   `json:"nested_nsx"`
	ResourceType             string                 `json:"resource_type"`
	TransportType            string                 `json:"transport_type"`
	Tags                     []Tag                  `json:"tags,omitempty"`
	TransportZoneProfileIDs  []TransportZoneProfile `json:"transport_zone_profile_ids"`
	UplinkTeamingPolicyNames []string               `json:"uplink_teaming_policy_names,omitempty"`
}

type ClusterNodesResponse struct {
	ResultCount int           `json:"result_count"`
	Results     []ClusterNode `json:"results"`
}

type ClusterNode struct {
	ID                      string          `json:"id"`
	DisplayName             string          `json:"display_name"`
	ExternalID              string          `json:"external_id,omitempty"`
	ApplianceMgmtListenAddr string          `json:"appliance_mgmt_listen_addr,omitempty"`
	ResourceType            string          `json:"resource_type"`
	ManagerRole             *ManagerRole    `json:"manager_role,omitempty"`
	ControllerRole          *ControllerRole `json:"controller_role,omitempty"`
	CreateTime              int64           `json:"_create_time"`
	CreateUser              string          `json:"_create_user"`
	LastModifiedTime        int64           `json:"_last_modified_time"`
	LastModifiedUser        string          `json:"_last_modified_user"`
	Protection              string          `json:"_protection"`
	Revision                int             `json:"_revision"`
	SystemOwned             bool            `json:"_system_owned"`
}

type ManagerRole struct {
	Type                  string              `json:"type"`
	APIListenAddr         ListenAddr          `json:"api_listen_addr"`
	ApplianceConnection   ApplianceConnection `json:"appliance_connection_info"`
	MgmtClusterListenAddr ListenAddr          `json:"mgmt_cluster_listen_addr"`
	MgmtPlaneListenAddr   ListenAddr          `json:"mgmt_plane_listen_addr"`
	MPAMsgClientInfo      MsgClientInfo       `json:"mpa_msg_client_info"`
}

type ControllerRole struct {
	Type                     string        `json:"type"`
	ControlClusterListenAddr ListenAddr    `json:"control_cluster_listen_addr"`
	ControlPlaneListenAddr   ListenAddr    `json:"control_plane_listen_addr"`
	HostMsgClientInfo        MsgClientInfo `json:"host_msg_client_info"`
	MPAMsgClientInfo         MsgClientInfo `json:"mpa_msg_client_info"`
}

type ListenAddr struct {
	IPAddress                   string `json:"ip_address"`
	Port                        int    `json:"port"`
	Certificate                 string `json:"certificate,omitempty"`
	CertificateSHA256Thumbprint string `json:"certificate_sha256_thumbprint,omitempty"`
}

type ApplianceConnection struct {
	IPAddress                   string         `json:"ip_address"`
	Port                        int            `json:"port"`
	Certificate                 string         `json:"certificate"`
	CertificateSHA256Thumbprint string         `json:"certificate_sha256_thumbprint"`
	ServiceEndpointUUID         string         `json:"service_endpoint_uuid"`
	EntitiesHosted              []EntityHosted `json:"entities_hosted"`
}

type EntityHosted struct {
	EntityType string `json:"entity_type"`
	EntityUUID string `json:"entity_uuid"`
}

type MsgClientInfo struct {
	AccountName string `json:"account_name"`
}

type IPPoolsResponse struct {
	ResultCount int      `json:"result_count"`
	Results     []IPPool `json:"results"`
}

type IPPool struct {
	ID               string         `json:"id"`
	DisplayName      string         `json:"display_name"`
	Description      string         `json:"description,omitempty"`
	PoolUsage        PoolUsage      `json:"pool_usage"`
	ResourceType     string         `json:"resource_type"`
	Subnets          []IPPoolSubnet `json:"subnets"`
	Tags             []Tag          `json:"tags,omitempty"`
	CreateTime       int64          `json:"_create_time"`
	CreateUser       string         `json:"_create_user"`
	LastModifiedTime int64          `json:"_last_modified_time"`
	LastModifiedUser string         `json:"_last_modified_user"`
	Protection       string         `json:"_protection"`
	Revision         int            `json:"_revision"`
	SystemOwned      bool           `json:"_system_owned"`
}

type PoolUsage struct {
	AllocatedIDs int `json:"allocated_ids"`
	FreeIDs      int `json:"free_ids"`
	TotalIDs     int `json:"total_ids"`
}

type IPPoolSubnet struct {
	CIDR             string            `json:"cidr"`
	AllocationRanges []AllocationRange `json:"allocation_ranges"`
	DNSNameservers   []string          `json:"dns_nameservers,omitempty"`
	DNSSuffix        string            `json:"dns_suffix,omitempty"`
	GatewayIP        string            `json:"gateway_ip,omitempty"`
}

type AllocationRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type LogicalRouterPortList struct {
	ResultCount int                 `json:"result_count"`
	Results     []LogicalRouterPort `json:"results"`
}

type LogicalRouterPort struct {
	ID              string   `json:"id"`
	DisplayName     string   `json:"display_name"`
	Description     string   `json:"description"`
	ResourceType    string   `json:"resource_type"`
	LogicalRouterID string   `json:"logical_router_id"`
	MacAddress      string   `json:"mac_address"`
	Subnets         []Subnet `json:"subnets"`
	// linked_logical_router_port_id can be either a string or a LinkPortID object
	LinkedLogicalRouterPortID LinkedLogicalRouterPortID `json:"linked_logical_router_port_id,omitempty"`
	LinkedLogicalSwitchPortID LinkedLogicalRouterPortID `json:"linked_logical_switch_port_id,omitempty"`
	EdgeClusterMemberIndex    []int                     `json:"edge_cluster_member_index,omitempty"`
	EnableMulticast           *bool                     `json:"enable_multicast,omitempty"`
	UrpFMode                  string                    `json:"urpf_mode,omitempty"`
	Mode                      string                    `json:"mode,omitempty"`
	MTU                       int                       `json:"mtu,omitempty"`
	Tags                      []Tag                     `json:"tags,omitempty"`
	ServiceBindings           []ServiceBinding          `json:"service_bindings,omitempty"`
}

type LinkedLogicalRouterPortID struct {
	ID     string        // if it's a string
	Object *LinkedPortID // if it's an object
}

type LinkedPortID struct {
	IsValid           bool   `json:"is_valid"`
	TargetID          string `json:"target_id"`
	TargetType        string `json:"target_type"`
	TargetDisplayName string `json:"target_display_name"`
}

func (l LinkedLogicalRouterPortID) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as string
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		l.ID = id
		return nil
	}

	// Try unmarshaling as LinkedPortID
	var obj LinkedPortID
	if err := json.Unmarshal(data, &obj); err == nil {
		l.Object = &obj
		return nil
	}

	return fmt.Errorf("linked_logical_router_port_id: unsupported format: %s", string(data))
}

type ServiceBinding struct {
	ServiceID LinkedPortID `json:"service_id"`
}

type FirewallSectionList struct {
	ResultCount int               `json:"result_count"`
	Results     []FirewallSection `json:"results"`
	SortBy      string            `json:"sort_by"`
}

type FirewallSection struct {
	ID               string   `json:"id"`
	DisplayName      string   `json:"display_name"`
	Description      string   `json:"description"`
	Comments         string   `json:"comments"`
	ResourceType     string   `json:"resource_type"`
	Category         string   `json:"category"`
	SectionType      string   `json:"section_type"`
	EnforcedOn       string   `json:"enforced_on"`
	IsDefault        bool     `json:"is_default"`
	Locked           bool     `json:"locked"`
	LockModifiedBy   string   `json:"lock_modified_by"`
	LockModifiedTime int64    `json:"lock_modified_time"`
	Stateful         bool     `json:"stateful"`
	TcpStrict        bool     `json:"tcp_strict"`
	RuleCount        int      `json:"rule_count"`
	Priority         int64    `json:"priority"`
	AutoPlumbed      bool     `json:"autoplumbed"`
	AppliedTos       []Target `json:"applied_tos,omitempty"`
	Tags             []Tag    `json:"tags,omitempty"`
	CreateTime       int64    `json:"_create_time"`
	CreateUser       string   `json:"_create_user"`
	LastModifiedTime int64    `json:"_last_modified_time"`
	LastModifiedUser string   `json:"_last_modified_user"`
	Protection       string   `json:"_protection"`
	Revision         int      `json:"_revision"`
	SystemOwned      bool     `json:"_system_owned"`
}

type Target struct {
	IsValid           bool   `json:"is_valid"`
	TargetDisplayName string `json:"target_display_name"`
	TargetID          string `json:"target_id"`
	TargetType        string `json:"target_type"`
}

type NetworkInterfaceList struct {
	Schema      string             `json:"_schema"`
	Self        Link               `json:"_self"`
	ResultCount int                `json:"result_count"`
	Results     []NetworkInterface `json:"results"`
}

type NetworkInterface struct {
	Schema           string      `json:"_schema"`
	Self             Link        `json:"_self"`
	AdminStatus      string      `json:"admin_status"`
	BroadcastAddress string      `json:"broadcast_address"`
	DefaultGateway   string      `json:"default_gateway,omitempty"`
	InterfaceID      string      `json:"interface_id"`
	IPAddresses      []IPAddress `json:"ip_addresses"`
	IPConfiguration  string      `json:"ip_configuration"`
	LinkStatus       string      `json:"link_status"`
	MTU              int         `json:"mtu"`
	PhysicalAddress  string      `json:"physical_address"`
}

type IPAddress struct {
	IPAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}
