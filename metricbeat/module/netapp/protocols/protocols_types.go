package protocols

// endpoint: /api/protocols/san/iscsi/services
type ISCSIServicesResponse struct {
	Records    []ISCSIService `json:"records"`
	NumRecords int            `json:"num_records"`
}

type ISCSIService struct {
	SVM     SVMInfo    `json:"svm"`
	Enabled bool       `json:"enabled"`
	Target  TargetInfo `json:"target"`
}

type SVMInfo struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type TargetInfo struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

// endpoint: /api/protocols/san/iscsi/sessions

type ISCSISessionsResponse struct {
	Links      Links          `json:"_links"`
	NumRecords int            `json:"num_records"`
	Records    []ISCSISession `json:"records"`
}

type ISCSISession struct {
	Links                Links             `json:"_links"`
	Connections          []ISCSIConnection `json:"connections"`
	Igroups              []ISCSIGroup      `json:"igroups"`
	Initiator            ISCSIInitiator    `json:"initiator"`
	ISID                 string            `json:"isid"`
	SVM                  ISCSISessionSVM   `json:"svm"`
	TargetPortalGroup    string            `json:"target_portal_group"`
	TargetPortalGroupTag int               `json:"target_portal_group_tag"`
	TSIH                 int               `json:"tsih"`
}

type ISCSIConnection struct {
	Links              Links          `json:"_links"`
	AuthenticationType string         `json:"authentication_type"`
	CID                int            `json:"cid"`
	InitiatorAddress   ISCSIAddress   `json:"initiator_address"`
	Interface          ISCSIInterface `json:"interface"`
}

type ISCSIAddress struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type ISCSIInterface struct {
	Links Links        `json:"_links"`
	IP    ISCSIAddress `json:"ip"`
	Name  string       `json:"name"`
	UUID  string       `json:"uuid"`
}

type ISCSIGroup struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type ISCSIInitiator struct {
	Alias   string `json:"alias"`
	Comment string `json:"comment"`
	Name    string `json:"name"`
}

type ISCSISessionSVM struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type Links struct {
	Next *Link `json:"next,omitempty"`
	Self *Link `json:"self,omitempty"`
}

type Link struct {
	Href string `json:"href"`
}

// endpoint: /api/protocols/cifs/services

type CIFSServicesResponse struct {
	Links      Links         `json:"_links"`
	NumRecords int           `json:"num_records"`
	Records    []CIFSService `json:"records"`
}

type CIFSService struct {
	Links                    Links          `json:"_links"`
	AdDomain                 CIFSAdDomain   `json:"ad_domain"`
	AuthStyle                string         `json:"auth-style"`
	AuthUserType             string         `json:"auth_user_type"`
	AuthenticationMethod     string         `json:"authentication_method"`
	ClientID                 string         `json:"client_id"`
	Comment                  string         `json:"comment"`
	DefaultUnixUser          string         `json:"default_unix_user"`
	Enabled                  bool           `json:"enabled"`
	GroupPolicyObjectEnabled bool           `json:"group_policy_object_enabled"`
	KeyVaultURI              string         `json:"key_vault_uri"`
	Metric                   CIFSMetric     `json:"metric"`
	Name                     string         `json:"name"`
	Netbios                  CIFSNetbios    `json:"netbios"`
	OAuthHost                string         `json:"oauth_host"`
	Options                  CIFSOptions    `json:"options"`
	ProxyHost                string         `json:"proxy_host"`
	ProxyPort                int            `json:"proxy_port"`
	ProxyType                string         `json:"proxy_type"`
	ProxyUsername            string         `json:"proxy_username"`
	Security                 CIFSSecurity   `json:"security"`
	Statistics               CIFSStatistics `json:"statistics"`
	SVM                      CIFSSVM        `json:"svm"`
	TenantID                 string         `json:"tenant_id"`
	Timeout                  int            `json:"timeout"`
	VerifyHost               bool           `json:"verify_host"`
	Workgroup                string         `json:"workgroup"`
}

type CIFSAdDomain struct {
	DefaultSite        string `json:"default_site"`
	FQDN               string `json:"fqdn"`
	OrganizationalUnit string `json:"organizational_unit"`
}

type CIFSMetric struct {
	Links      Links          `json:"_links"`
	Duration   string         `json:"duration"`
	IOPS       CIFSIOPS       `json:"iops"`
	Latency    CIFSIOPS       `json:"latency"`
	Status     string         `json:"status"`
	Throughput CIFSThroughput `json:"throughput"`
	Timestamp  string         `json:"timestamp"`
}

type CIFSIOPS struct {
	Other int `json:"other"`
	Read  int `json:"read"`
	Total int `json:"total"`
	Write int `json:"write"`
}

type CIFSThroughput struct {
	Read  int `json:"read"`
	Total int `json:"total"`
	Write int `json:"write"`
}

type CIFSNetbios struct {
	Aliases     []string `json:"aliases"`
	Enabled     bool     `json:"enabled"`
	WinsServers []string `json:"wins_servers"`
}

type CIFSOptions struct {
	AdminToRootMapping               bool     `json:"admin_to_root_mapping"`
	AdvancedSparseFile               bool     `json:"advanced_sparse_file"`
	BackupSymlinkEnabled             bool     `json:"backup_symlink_enabled"`
	ClientDupDetectionEnabled        bool     `json:"client_dup_detection_enabled"`
	ClientVersionReportingEnabled    bool     `json:"client_version_reporting_enabled"`
	CopyOffload                      bool     `json:"copy_offload"`
	DacEnabled                       bool     `json:"dac_enabled"`
	ExportPolicyEnabled              bool     `json:"export_policy_enabled"`
	FakeOpen                         bool     `json:"fake_open"`
	FsctlTrim                        bool     `json:"fsctl_trim"`
	JunctionReparse                  bool     `json:"junction_reparse"`
	LargeMTU                         bool     `json:"large_mtu"`
	MaxConnectionsPerSession         int      `json:"max_connections_per_session"`
	MaxLifsPerSession                int      `json:"max_lifs_per_session"`
	MaxOpensSameFilePerTree          int      `json:"max_opens_same_file_per_tree"`
	MaxSameTreeConnectPerSession     int      `json:"max_same_tree_connect_per_session"`
	MaxSameUserSessionsPerConnection int      `json:"max_same_user_sessions_per_connection"`
	MaxWatchesSetPerTree             int      `json:"max_watches_set_per_tree"`
	Multichannel                     bool     `json:"multichannel"`
	NullUserWindowsName              string   `json:"null_user_windows_name"`
	PathComponentCache               bool     `json:"path_component_cache"`
	Referral                         bool     `json:"referral"`
	Shadowcopy                       bool     `json:"shadowcopy"`
	ShadowcopyDirDepth               int      `json:"shadowcopy_dir_depth"`
	SmbCredits                       int      `json:"smb_credits"`
	TrustedDomainEnumSearchEnabled   bool     `json:"trusted_domain_enum_search_enabled"`
	WidelinkReparseVersions          []string `json:"widelink_reparse_versions"`
}

type CIFSSecurity struct {
	AdvertisedKDCEncryptions []string `json:"advertised_kdc_encryptions"`
	AESNetlogonEnabled       bool     `json:"aes_netlogon_enabled"`
	EncryptDCConnection      bool     `json:"encrypt_dc_connection"`
	KDCEncryption            bool     `json:"kdc_encryption"`
	LDAPReferralEnabled      bool     `json:"ldap_referral_enabled"`
	LMCompatibilityLevel     string   `json:"lm_compatibility_level"`
	RestrictAnonymous        string   `json:"restrict_anonymous"`
	SessionSecurity          string   `json:"session_security"`
	SMBEncryption            bool     `json:"smb_encryption"`
	SMBSigning               bool     `json:"smb_signing"`
	TryLDAPChannelBinding    bool     `json:"try_ldap_channel_binding"`
	UseLDAPS                 bool     `json:"use_ldaps"`
	UseStartTLS              bool     `json:"use_start_tls"`
}

type CIFSStatistics struct {
	IOPSRaw       CIFSIOPS       `json:"iops_raw"`
	LatencyRaw    CIFSIOPS       `json:"latency_raw"`
	Status        string         `json:"status"`
	ThroughputRaw CIFSThroughput `json:"throughput_raw"`
	Timestamp     string         `json:"timestamp"`
}

type CIFSSVM struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

// endpoint: /api/protocols/cifs/shares
type CIFSSharesResponse struct {
	Links      Links       `json:"_links"`
	NumRecords int         `json:"num_records"`
	Records    []CIFSShare `json:"records"`
}

type CIFSShare struct {
	Links                  Links          `json:"_links"`
	AccessBasedEnumeration bool           `json:"access_based_enumeration"`
	Acls                   []CIFSShareACL `json:"acls"`
	AllowUnencryptedAccess bool           `json:"allow_unencrypted_access"`
	AttributeCache         bool           `json:"attribute_cache"`
	Browsable              bool           `json:"browsable"`
	ChangeNotify           bool           `json:"change_notify"`
	Comment                string         `json:"comment"`
	ContinuouslyAvailable  bool           `json:"continuously_available"`
	DirUmask               string         `json:"dir_umask"`
	Encryption             bool           `json:"encryption"`
	FileUmask              string         `json:"file_umask"`
	ForceGroupForCreate    string         `json:"force_group_for_create"`
	HomeDirectory          bool           `json:"home_directory"`
	MaxConnectionsPerShare int            `json:"max_connections_per_share"`
	Name                   string         `json:"name"`
	NamespaceCaching       bool           `json:"namespace_caching"`
	NoStrictSecurity       bool           `json:"no_strict_security"`
	OfflineFiles           string         `json:"offline_files"`
	Oplocks                bool           `json:"oplocks"`
	Path                   string         `json:"path"`
	ShowPreviousVersions   bool           `json:"show_previous_versions"`
	ShowSnapshot           bool           `json:"show_snapshot"`
	SVM                    CIFSSVM        `json:"svm"`
	UnixSymlink            string         `json:"unix_symlink"`
	Volume                 CIFSVolume     `json:"volume"`
	VscanProfile           string         `json:"vscan_profile"`
}

type CIFSShareACL struct {
	Links        Links  `json:"_links"`
	Permission   string `json:"permission"`
	Type         string `json:"type"`
	UserOrGroup  string `json:"user_or_group"`
	WinSidUnixID string `json:"win_sid_unix_id"`
}

type CIFSVolume struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

// endpoint: /api/protocols/san/igroups
type IGrouplistResponse struct {
	Links      Links    `json:"_links"`
	NumRecords int      `json:"num_records"`
	Records    []IGroup `json:"records"`
}

type IGroup struct {
	Links                Links               `json:"_links"`
	Comment              string              `json:"comment"`
	ConnectivityTracking *IGroupConnectivity `json:"connectivity_tracking,omitempty"`
	DeleteOnUnmap        bool                `json:"delete_on_unmap"`
	Igroups              []IGroupNested      `json:"igroups"`
	Initiators           []IGroupInitiator   `json:"initiators"`
	LunMaps              []IGroupLunMap      `json:"lun_maps"`
	Name                 string              `json:"name"`
	OsType               string              `json:"os_type"`
	ParentIgroups        []IGroupParent      `json:"parent_igroups"`
	Portset              *IGroupPortset      `json:"portset,omitempty"`
	Protocol             string              `json:"protocol"`
	Replication          *IGroupReplication  `json:"replication,omitempty"`
	SupportsIgroups      bool                `json:"supports_igroups"`
	SVM                  IGroupSVM           `json:"svm"`
	Target               *IGroupTarget       `json:"target,omitempty"`
	UUID                 string              `json:"uuid"`
}

type IGroupConnectivity struct {
	Alerts          []IGroupAlert `json:"alerts"`
	ConnectionState string        `json:"connection_state"`
	RequiredNodes   []IGroupNode  `json:"required_nodes"`
}

type IGroupAlert struct {
	Summary IGroupAlertSummary `json:"summary"`
}

type IGroupAlertSummary struct {
	Arguments []IGroupAlertArgument `json:"arguments"`
	Code      string                `json:"code"`
	Message   string                `json:"message"`
}

type IGroupAlertArgument struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type IGroupNode struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type IGroupNested struct {
	Links   Links           `json:"_links"`
	Comment string          `json:"comment"`
	Igroups []*IGroupNested `json:"igroups"`
	Name    string          `json:"name"`
	UUID    string          `json:"uuid"`
}

type IGroupInitiator struct {
	Links                IGroupInitiatorLinks `json:"_links"`
	Comment              string               `json:"comment"`
	ConnectivityTracking *IGroupInitiatorConn `json:"connectivity_tracking,omitempty"`
	Igroup               *IGroupNested        `json:"igroup,omitempty"`
	Name                 string               `json:"name"`
	Proximity            *IGroupInitiatorProx `json:"proximity,omitempty"`
}

type IGroupInitiatorLinks struct {
	ConnectivityTracking *Link `json:"connectivity_tracking,omitempty"`
	Self                 *Link `json:"self,omitempty"`
}

type IGroupInitiatorConn struct {
	ConnectionState string `json:"connection_state"`
}

type IGroupInitiatorProx struct {
	LocalSVM bool            `json:"local_svm"`
	PeerSVMs []IGroupPeerSVM `json:"peer_svms"`
}

type IGroupPeerSVM struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type IGroupLunMap struct {
	Links             Links     `json:"_links"`
	LogicalUnitNumber int       `json:"logical_unit_number"`
	Lun               IGroupLun `json:"lun"`
}

type IGroupLun struct {
	Links Links       `json:"_links"`
	Name  string      `json:"name"`
	Node  *IGroupNode `json:"node,omitempty"`
	UUID  string      `json:"uuid"`
}

type IGroupParent struct {
	Links         Links           `json:"_links"`
	Comment       string          `json:"comment"`
	Name          string          `json:"name"`
	ParentIgroups []*IGroupParent `json:"parent_igroups"`
	UUID          string          `json:"uuid"`
}

type IGroupPortset struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type IGroupReplication struct {
	Error   *IGroupReplicationError `json:"error,omitempty"`
	PeerSVM *IGroupPeerSVM          `json:"peer_svm,omitempty"`
	State   string                  `json:"state"`
}

type IGroupReplicationError struct {
	Igroup  IGroupReplicationErrorIgroup `json:"igroup"`
	Summary IGroupAlertSummary           `json:"summary"`
}

type IGroupReplicationErrorIgroup struct {
	LocalSVM bool   `json:"local_svm"`
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
}

type IGroupSVM struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type IGroupTarget struct {
	FirmwareRevision string `json:"firmware_revision"`
	ProductID        string `json:"product_id"`
	VendorID         string `json:"vendor_id"`
}

// endpoint: /api/network/fc/interfaces

// endpoint: /api/network/fc/interfaces

type FCInterfacesResponse struct {
	Links      Links         `json:"_links"`
	NumRecords int           `json:"num_records"`
	Recommend  *FCRecommend  `json:"recommend,omitempty"`
	Records    []FCInterface `json:"records"`
}

type FCRecommend struct {
	Messages []FCRecommendMessage `json:"messages"`
}

type FCRecommendMessage struct {
	Arguments []FCRecommendArgument `json:"arguments"`
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Severity  string                `json:"severity"`
}

type FCRecommendArgument struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type FCInterface struct {
	Links        Links         `json:"_links"`
	Comment      string        `json:"comment"`
	DataProtocol string        `json:"data_protocol"`
	Enabled      bool          `json:"enabled"`
	Location     FCLocation    `json:"location"`
	Metric       *FCMetric     `json:"metric,omitempty"`
	Name         string        `json:"name"`
	PortAddress  string        `json:"port_address"`
	State        string        `json:"state"`
	Statistics   *FCStatistics `json:"statistics,omitempty"`
	SVM          CIFSSVM       `json:"svm"`
	UUID         string        `json:"uuid"`
	WWNN         string        `json:"wwnn"`
	WWPN         string        `json:"wwpn"`
}

type FCLocation struct {
	HomeNode FCNode `json:"home_node"`
	HomePort FCPort `json:"home_port"`
	IsHome   bool   `json:"is_home"`
	Node     FCNode `json:"node"`
	Port     FCPort `json:"port"`
}

type FCNode struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type FCMetric struct {
	Links      Links          `json:"_links"`
	Duration   string         `json:"duration"`
	IOPS       CIFSIOPS       `json:"iops"`
	Latency    CIFSIOPS       `json:"latency"`
	Status     string         `json:"status"`
	Throughput CIFSThroughput `json:"throughput"`
	Timestamp  string         `json:"timestamp"`
}

type FCStatistics struct {
	IOPSRaw       CIFSIOPS       `json:"iops_raw"`
	LatencyRaw    CIFSIOPS       `json:"latency_raw"`
	Status        string         `json:"status"`
	ThroughputRaw CIFSThroughput `json:"throughput_raw"`
	Timestamp     string         `json:"timestamp"`
}

// endpoint: /api/network/fc/port

type FCPortsResponse struct {
	NumRecords int      `json:"num_records"`
	Records    []FCPort `json:"records"`
}

type FCPort struct {
	Links              Links              `json:"_links"`
	Node               FCNode             `json:"node"`
	Name               string             `json:"name"`
	UUID               string             `json:"uuid"`
	Description        string             `json:"description"`
	Enabled            bool               `json:"enabled"`
	Fabric             FCPortFabric       `json:"fabric"`
	PhysicalProtocol   string             `json:"physical_protocol"`
	Speed              FCPortSpeed        `json:"speed"`
	State              string             `json:"state"`
	SupportedProtocols []string           `json:"supported_protocols"`
	Transceiver        *FCPortTransceiver `json:"transceiver,omitempty"`
	WWNN               string             `json:"wwnn"`
	WWPN               string             `json:"wwpn"`
}

type FCPortFabric struct {
	Connected      bool   `json:"connected"`
	ConnectedSpeed int    `json:"connected_speed"`
	PortAddress    string `json:"port_address"`
	SwitchPort     string `json:"switch_port"`
}

type FCPortSpeed struct {
	Maximum    string `json:"maximum"`
	Configured string `json:"configured"`
}

type FCPortTransceiver struct {
	FormFactor   string `json:"form_factor"`
	Manufacturer string `json:"manufacturer"`
	Capabilities []int  `json:"capabilities"`
	PartNumber   string `json:"part_number"`
}

// endpoint: /api/protocols/san/fcp/services

type FCPServicesResponse struct {
	Links      Links        `json:"_links"`
	NumRecords int          `json:"num_records"`
	Records    []FCPService `json:"records"`
}

type FCPService struct {
	Links      Links            `json:"_links"`
	Enabled    bool             `json:"enabled"`
	Metric     *FCMetric        `json:"metric,omitempty"`
	Statistics *FCStatistics    `json:"statistics,omitempty"`
	SVM        CIFSSVM          `json:"svm"`
	Target     FCPServiceTarget `json:"target"`
}

type FCPServiceTarget struct {
	Name string `json:"name"`
}

// endpoint: /api/protocols/nfs/services

type NFSServicesResponse struct {
	Records    []NFSService `json:"records"`
	NumRecords int          `json:"num_records"`
}

type NFSService struct {
	SVM                           SVMInfo              `json:"svm"`
	Enabled                       bool                 `json:"enabled"`
	State                         string               `json:"state"`
	Transport                     NFSTransport         `json:"transport"`
	Protocol                      NFSProtocol          `json:"protocol"`
	VstorageEnabled               bool                 `json:"vstorage_enabled"`
	RquotaEnabled                 bool                 `json:"rquota_enabled"`
	ShowmountEnabled              bool                 `json:"showmount_enabled"`
	AuthSysExtendedGroupsEnabled  bool                 `json:"auth_sys_extended_groups_enabled"`
	ExtendedGroupsLimit           int                  `json:"extended_groups_limit"`
	CredentialCache               NFSCredentialCache   `json:"credential_cache"`
	Qtree                         NFSQtree             `json:"qtree"`
	AccessCacheConfig             NFSAccessCacheConfig `json:"access_cache_config"`
	FileSessionIOGroupingCount    int                  `json:"file_session_io_grouping_count"`
	FileSessionIOGroupingDuration int                  `json:"file_session_io_grouping_duration"`
	Exports                       NFSExports           `json:"exports"`
	Security                      NFSSecurity          `json:"security"`
	Windows                       NFSWindows           `json:"windows"`
}

type NFSTransport struct {
	UDPEnabled  bool `json:"udp_enabled"`
	TCPEnabled  bool `json:"tcp_enabled"`
	RDMAEnabled bool `json:"rdma_enabled"`
}

type NFSProtocol struct {
	V3Enabled                 bool           `json:"v3_enabled"`
	V364bitIdentifiersEnabled bool           `json:"v3_64bit_identifiers_enabled"`
	V4IDDomain                string         `json:"v4_id_domain"`
	V464bitIdentifiersEnabled bool           `json:"v4_64bit_identifiers_enabled"`
	V40Enabled                bool           `json:"v40_enabled"`
	V41Enabled                bool           `json:"v41_enabled"`
	V4GraceSeconds            int            `json:"v4_grace_seconds"`
	V40Features               NFSV40Features `json:"v40_features"`
	V41Features               NFSV41Features `json:"v41_features"`
	V3Features                NFSV3Features  `json:"v3_features"`
}

type NFSV40Features struct {
	ACLEnabled             bool `json:"acl_enabled"`
	ReadDelegationEnabled  bool `json:"read_delegation_enabled"`
	WriteDelegationEnabled bool `json:"write_delegation_enabled"`
	ACLPreserve            bool `json:"acl_preserve"`
}

type NFSV41Features struct {
	ACLEnabled             bool `json:"acl_enabled"`
	ReadDelegationEnabled  bool `json:"read_delegation_enabled"`
	WriteDelegationEnabled bool `json:"write_delegation_enabled"`
	PnfsEnabled            bool `json:"pnfs_enabled"`
}

type NFSV3Features struct {
	MountRootOnly bool `json:"mount_root_only"`
}

type NFSCredentialCache struct {
	PositiveTTL int `json:"positive_ttl"`
}

type NFSQtree struct {
	ExportEnabled  bool `json:"export_enabled"`
	ValidateExport bool `json:"validate_export"`
}

type NFSAccessCacheConfig struct {
	TTLPositive     int  `json:"ttl_positive"`
	TTLNegative     int  `json:"ttl_negative"`
	HarvestTimeout  int  `json:"harvest_timeout"`
	IsDnsTTLEnabled bool `json:"isDnsTTLEnabled"`
}

type NFSExports struct {
	NameServiceLookupProtocol string `json:"name_service_lookup_protocol"`
}

type NFSSecurity struct {
	PermittedEncryptionTypes []string `json:"permitted_encryption_types"`
}

type NFSWindows struct {
	V3MsDosClientEnabled bool `json:"v3_ms_dos_client_enabled"`
}

// endpoint: /api/protocols/nfs/services/{svm.uuid}/metrics

type NFSServicesMetricsResponse struct {
	Records    []NFSServiceMetrics `json:"records"`
	NumRecords int                 `json:"num_records"`
}

type NFSServiceMetrics struct {
	UUID      string             `json:"uuid"`
	Timestamp string             `json:"timestamp"`
	SVM       SVMInfo            `json:"svm"`
	V3        NFSServiceMetricsV `json:"v3"`
	V4        NFSServiceMetricsV `json:"v4"`
	V41       NFSServiceMetricsV `json:"v41"`
}

type NFSServiceMetricsV struct {
	Status     string         `json:"status"`
	Duration   string         `json:"duration"`
	Throughput CIFSThroughput `json:"throughput"`
	IOPS       CIFSIOPS       `json:"iops"`
	Latency    CIFSIOPS       `json:"latency"`
}

// endpoint: /api/protocols/nfs/export-policies

type NFSExportPoliciesResponse struct {
	Records    []NFSExportPolicy `json:"records"`
	NumRecords int               `json:"num_records"`
}

type NFSExportPolicy struct {
	SVM  SVMInfo `json:"svm"`
	ID   int64   `json:"id"`
	Name string  `json:"name"`
}

// endpoint: /api/network/ip/interfaces

type IPInterfacesResponse struct {
	Links      Links         `json:"_links"`
	NumRecords int           `json:"num_records"`
	Recommend  *FCRecommend  `json:"recommend,omitempty"`
	Records    []IPInterface `json:"records"`
}

type IPInterface struct {
	Links         Links           `json:"_links"`
	DDNSEnabled   bool            `json:"ddns_enabled"`
	DNSZone       string          `json:"dns_zone"`
	Enabled       bool            `json:"enabled"`
	IP            IPAddress       `json:"ip"`
	IPSpace       IPSpace         `json:"ipspace"`
	Location      IPLocation      `json:"location"`
	Metric        *FCMetric       `json:"metric,omitempty"`
	Name          string          `json:"name"`
	ProbePort     int             `json:"probe_port"`
	RDMAProtocols []string        `json:"rdma_protocols"`
	Scope         string          `json:"scope"`
	ServicePolicy IPServicePolicy `json:"service_policy"`
	Services      []string        `json:"services"`
	State         string          `json:"state"`
	Statistics    *FCStatistics   `json:"statistics,omitempty"`
	Subnet        IPSubnet        `json:"subnet"`
	SVM           CIFSSVM         `json:"svm"`
	UUID          string          `json:"uuid"`
	VIP           bool            `json:"vip"`
}

type IPAddress struct {
	Address string `json:"address"`
	Family  string `json:"family"`
	Netmask string `json:"netmask"`
}

type IPSpace struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type IPLocation struct {
	AutoRevert bool           `json:"auto_revert"`
	Failover   string         `json:"failover"`
	HomeNode   FCNode         `json:"home_node"`
	HomePort   IPLocationPort `json:"home_port"`
	IsHome     bool           `json:"is_home"`
	Node       FCNode         `json:"node"`
	Port       IPLocationPort `json:"port"`
}

type IPLocationPort struct {
	Links Links   `json:"_links"`
	Name  string  `json:"name"`
	Node  *FCNode `json:"node,omitempty"`
	UUID  string  `json:"uuid"`
}

type IPServicePolicy struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}

type IPSubnet struct {
	Links Links  `json:"_links"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}
