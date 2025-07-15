package protocols

import "github.com/elastic/beats/v7/metricbeat/module/netapp"

// endpoint: /api/protocols/san/iscsi/services
type ISCSIService struct {
	SVM        netapp.NamedObject `json:"svm"`
	Enabled    bool               `json:"enabled"`
	Target     TargetInfo         `json:"target"`
	Metric     netapp.Metrics     `json:"metric"`
	Statistics netapp.Statistics  `json:"statistics"`
}

type TargetInfo struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

// endpoint: /api/protocols/san/iscsi/sessions

type ISCSISession struct {
	Connections          []ISCSIConnection    `json:"connections"`
	Igroups              []netapp.NamedObject `json:"igroups"`
	Initiator            ISCSIInitiator       `json:"initiator"`
	ISID                 string               `json:"isid"`
	SVM                  netapp.NamedObject   `json:"svm"`
	TargetPortalGroup    string               `json:"target_portal_group"`
	TargetPortalGroupTag int                  `json:"target_portal_group_tag"`
	TSIH                 int                  `json:"tsih"`
}

type ISCSIConnection struct {
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
	IP   ISCSIAddress `json:"ip"`
	Name string       `json:"name"`
	UUID string       `json:"uuid"`
}

type ISCSIInitiator struct {
	Alias   string `json:"alias"`
	Comment string `json:"comment"`
	Name    string `json:"name"`
}

// endpoint: /api/protocols/cifs/services

type CIFSService struct {
	AdDomain                 CIFSAdDomain       `json:"ad_domain"`
	AuthStyle                string             `json:"auth-style"`
	AuthUserType             string             `json:"auth_user_type"`
	AuthenticationMethod     string             `json:"authentication_method"`
	ClientID                 string             `json:"client_id"`
	Comment                  string             `json:"comment"`
	DefaultUnixUser          string             `json:"default_unix_user"`
	Enabled                  bool               `json:"enabled"`
	GroupPolicyObjectEnabled bool               `json:"group_policy_object_enabled"`
	KeyVaultURI              string             `json:"key_vault_uri"`
	Name                     string             `json:"name"`
	Netbios                  CIFSNetbios        `json:"netbios"`
	OAuthHost                string             `json:"oauth_host"`
	Options                  CIFSOptions        `json:"options"`
	ProxyHost                string             `json:"proxy_host"`
	ProxyPort                int                `json:"proxy_port"`
	ProxyType                string             `json:"proxy_type"`
	ProxyUsername            string             `json:"proxy_username"`
	Security                 CIFSSecurity       `json:"security"`
	SVM                      netapp.NamedObject `json:"svm"`
	TenantID                 string             `json:"tenant_id"`
	Timeout                  int                `json:"timeout"`
	VerifyHost               bool               `json:"verify_host"`
	Workgroup                string             `json:"workgroup"`
	Metric                   netapp.Metrics     `json:"metric"`
	Statistics               netapp.Statistics  `json:"statistics"`
}

type CIFSAdDomain struct {
	DefaultSite        string `json:"default_site"`
	FQDN               string `json:"fqdn"`
	OrganizationalUnit string `json:"organizational_unit"`
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

// endpoint: /api/protocols/cifs/shares

type CIFSShare struct {
	AccessBasedEnumeration bool               `json:"access_based_enumeration"`
	Acls                   []CIFSShareACL     `json:"acls"`
	AllowUnencryptedAccess bool               `json:"allow_unencrypted_access"`
	AttributeCache         bool               `json:"attribute_cache"`
	Browsable              bool               `json:"browsable"`
	ChangeNotify           bool               `json:"change_notify"`
	Comment                string             `json:"comment"`
	ContinuouslyAvailable  bool               `json:"continuously_available"`
	DirUmask               string             `json:"dir_umask"`
	Encryption             bool               `json:"encryption"`
	FileUmask              string             `json:"file_umask"`
	ForceGroupForCreate    string             `json:"force_group_for_create"`
	HomeDirectory          bool               `json:"home_directory"`
	MaxConnectionsPerShare int                `json:"max_connections_per_share"`
	Name                   string             `json:"name"`
	NamespaceCaching       bool               `json:"namespace_caching"`
	NoStrictSecurity       bool               `json:"no_strict_security"`
	OfflineFiles           string             `json:"offline_files"`
	Oplocks                bool               `json:"oplocks"`
	Path                   string             `json:"path"`
	ShowPreviousVersions   bool               `json:"show_previous_versions"`
	ShowSnapshot           bool               `json:"show_snapshot"`
	SVM                    netapp.NamedObject `json:"svm"`
	UnixSymlink            string             `json:"unix_symlink"`
	Volume                 netapp.NamedObject `json:"volume"`
	VscanProfile           string             `json:"vscan_profile"`
}

type CIFSShareACL struct {
	Permission   string `json:"permission"`
	Type         string `json:"type"`
	UserOrGroup  string `json:"user_or_group"`
	WinSidUnixID string `json:"win_sid_unix_id"`
}

// endpoint: /api/protocols/san/igroups

type IGroup struct {
	Comment              string              `json:"comment"`
	ConnectivityTracking IGroupConnectivity  `json:"connectivity_tracking,omitempty"`
	DeleteOnUnmap        bool                `json:"delete_on_unmap"`
	Igroups              []IGroupNested      `json:"igroups"`
	Initiators           []IGroupInitiator   `json:"initiators"`
	LunMaps              []IGroupLunMap      `json:"lun_maps"`
	Name                 string              `json:"name"`
	OsType               string              `json:"os_type"`
	ParentIgroups        []IGroupParent      `json:"parent_igroups"`
	Portset              *netapp.NamedObject `json:"portset,omitempty"`
	Protocol             string              `json:"protocol"`
	Replication          *IGroupReplication  `json:"replication,omitempty"`
	SupportsIgroups      bool                `json:"supports_igroups"`
	SVM                  netapp.NamedObject  `json:"svm"`
	Target               *IGroupTarget       `json:"target,omitempty"`
	UUID                 string              `json:"uuid"`
}

type IGroupConnectivity struct {
	Alerts          []IGroupAlert        `json:"alerts"`
	ConnectionState string               `json:"connection_state"`
	RequiredNodes   []netapp.NamedObject `json:"required_nodes"`
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

type IGroupNested struct {
	Comment string          `json:"comment"`
	Igroups []*IGroupNested `json:"igroups"`
	Name    string          `json:"name"`
	UUID    string          `json:"uuid"`
}

type IGroupInitiator struct {
	Comment              string                   `json:"comment"`
	ConnectivityTracking IGroupInitiatorConn      `json:"connectivity_tracking,omitempty"`
	Igroup               IGroupNested             `json:"igroup,omitempty"`
	Name                 string                   `json:"name"`
	Proximity            IGroupInitiatorProximity `json:"proximity,omitempty"`
}

type IGroupInitiatorLinks struct {
	ConnectivityTracking *netapp.Link `json:"connectivity_tracking,omitempty"`
	Self                 *netapp.Link `json:"self,omitempty"`
}

type IGroupInitiatorConn struct {
	ConnectionState string `json:"connection_state"`
}

type IGroupInitiatorProximity struct {
	LocalSVM bool                 `json:"local_svm"`
	PeerSVMs []netapp.NamedObject `json:"peer_svms"`
}

type IGroupLunMap struct {
	LogicalUnitNumber int       `json:"logical_unit_number"`
	Lun               IGroupLun `json:"lun"`
}

type IGroupLun struct {
	Name string              `json:"name"`
	Node *netapp.NamedObject `json:"node,omitempty"`
	UUID string              `json:"uuid"`
}

type IGroupParent struct {
	Comment       string          `json:"comment"`
	Name          string          `json:"name"`
	ParentIgroups []*IGroupParent `json:"parent_igroups"`
	UUID          string          `json:"uuid"`
}

type IGroupReplication struct {
	Error   *IGroupReplicationError `json:"error,omitempty"`
	PeerSVM *netapp.NamedObject     `json:"peer_svm,omitempty"`
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

type IGroupTarget struct {
	FirmwareRevision string `json:"firmware_revision"`
	ProductID        string `json:"product_id"`
	VendorID         string `json:"vendor_id"`
}

// endpoint: /api/network/fc/interfaces

type FCInterface struct {
	Comment      string             `json:"comment"`
	DataProtocol string             `json:"data_protocol"`
	Enabled      bool               `json:"enabled"`
	Location     FCLocation         `json:"location"`
	Name         string             `json:"name"`
	PortAddress  string             `json:"port_address"`
	State        string             `json:"state"`
	SVM          netapp.NamedObject `json:"svm"`
	UUID         string             `json:"uuid"`
	WWNN         string             `json:"wwnn"`
	WWPN         string             `json:"wwpn"`
	Metric       netapp.Metrics     `json:"metric,omitempty"`
	Statistics   netapp.Statistics  `json:"statistics,omitempty"`
}

type FCLocation struct {
	HomeNode netapp.NamedObject `json:"home_node"`
	HomePort HomePort           `json:"home_port"`
	IsHome   bool               `json:"is_home"`
	Node     netapp.NamedObject `json:"node"`
	Port     HomePort           `json:"port"`
}

type HomePort struct {
	Name string   `json:"name"`
	Node PortNode `json:"node"`
	UUID string   `json:"uuid"`
}

type PortNode struct {
	Name string `json:"name"`
}

// endpoint: /api/network/fc/ports

type FCPort struct {
	Node               netapp.NamedObject `json:"node"`
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
	Metric             netapp.Metrics     `json:"metric"`
	Statistics         netapp.Statistics  `json:"statistics"`
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

type FCPService struct {
	Enabled    bool               `json:"enabled"`
	Metric     netapp.Metrics     `json:"metric,omitempty"`
	Statistics netapp.Statistics  `json:"statistics,omitempty"`
	SVM        netapp.NamedObject `json:"svm"`
	Target     FCPServiceTarget   `json:"target"`
}

type FCPServiceTarget struct {
	Name string `json:"name"`
}

// endpoint: /api/protocols/nfs/services

type NFSService struct {
	SVM                           netapp.NamedObject   `json:"svm"`
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
	Metric                        NFSMetric            `json:"metric"`
	Statistics                    NFSStatistics        `json:"statistics"`
}

type NFSMetric struct {
	V3  netapp.Metrics `json:"v3"`
	V4  netapp.Metrics `json:"v4"`
	V41 netapp.Metrics `json:"v41"`
}

type NFSStatistics struct {
	V3  netapp.Statistics `json:"v3"`
	V4  netapp.Statistics `json:"v4"`
	V41 netapp.Statistics `json:"v41"`
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

// endpoint: /api/protocols/nfs/export-policies

type NFSExportPolicy struct {
	SVM  netapp.NamedObject `json:"svm"`
	ID   int64              `json:"id"`
	Name string             `json:"name"`
}

// endpoint: /api/network/ip/interfaces

type IPInterface struct {
	DDNSEnabled   bool               `json:"ddns_enabled"`
	DNSZone       string             `json:"dns_zone"`
	Enabled       bool               `json:"enabled"`
	IP            IPAddress          `json:"ip"`
	IPSpace       netapp.NamedObject `json:"ipspace"`
	Location      IPLocation         `json:"location"`
	Metric        netapp.Metrics     `json:"metric,omitempty"`
	Name          string             `json:"name"`
	ProbePort     int                `json:"probe_port"`
	RDMAProtocols []string           `json:"rdma_protocols"`
	Scope         string             `json:"scope"`
	ServicePolicy netapp.NamedObject `json:"service_policy"`
	Services      []string           `json:"services"`
	State         string             `json:"state"`
	Statistics    netapp.Statistics  `json:"statistics,omitempty"`
	Subnet        netapp.NamedObject `json:"subnet"`
	SVM           netapp.NamedObject `json:"svm"`
	UUID          string             `json:"uuid"`
	VIP           bool               `json:"vip"`
}

type IPAddress struct {
	Address string `json:"address"`
	Family  string `json:"family"`
	Netmask string `json:"netmask"`
}

type IPLocation struct {
	AutoRevert bool               `json:"auto_revert"`
	Failover   string             `json:"failover"`
	HomeNode   netapp.NamedObject `json:"home_node"`
	HomePort   IPLocationPort     `json:"home_port"`
	IsHome     bool               `json:"is_home"`
	Node       netapp.NamedObject `json:"node"`
	Port       IPLocationPort     `json:"port"`
}

type IPLocationPort struct {
	Name string              `json:"name"`
	Node *netapp.NamedObject `json:"node,omitempty"`
	UUID string              `json:"uuid"`
}
