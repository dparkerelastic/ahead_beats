package cluster

import (
	"time"

	"github.com/elastic/beats/v7/metricbeat/module/netapp"
)

type Cluster struct {
	Name                    string                `json:"name"`
	UUID                    string                `json:"uuid"`
	Location                string                `json:"location"`
	Contact                 string                `json:"contact"`
	Version                 NodeVersion           `json:"version"`
	DNSDomains              []string              `json:"dns_domains"`
	NameServers             []string              `json:"name_servers"`
	NTPServers              []string              `json:"ntp_servers"`
	ManagementInterfaces    []netapp.IPInterface  `json:"management_interfaces"`
	Metric                  netapp.Metrics        `json:"metric"`
	Statistics              netapp.Statistics     `json:"statistics"`
	Timezone                *ClusterTimezone      `json:"timezone"`
	Certificate             *ClusterCertificate   `json:"certificate"`
	SANOptimized            bool                  `json:"san_optimized"`
	Disaggregated           bool                  `json:"disaggregated"`
	PeeringPolicy           *ClusterPeeringPolicy `json:"peering_policy"`
	AutoEnableAnalytics     bool                  `json:"auto_enable_analytics"`
	AutoEnableActivityTrack bool                  `json:"auto_enable_activity_tracking"`
}

type ClusterTimezone struct {
	Name string `json:"name"`
}

type ClusterCertificate struct {
	UUID string `json:"uuid"`
}

type ClusterPeeringPolicy struct {
	MinimumPassphraseLength int  `json:"minimum_passphrase_length"`
	AuthenticationRequired  bool `json:"authentication_required"`
	EncryptionRequired      bool `json:"encryption_required"`
}

type Node struct {
	UUID                  string               `json:"uuid"`
	Name                  string               `json:"name"`
	SerialNumber          string               `json:"serial_number"`
	Location              string               `json:"location"`
	Owner                 string               `json:"owner"`
	Model                 string               `json:"model"`
	SystemID              string               `json:"system_id"`
	Version               NodeVersion          `json:"version"`
	Date                  time.Time            `json:"date"`
	Uptime                int64                `json:"uptime"`
	State                 string               `json:"state"`
	Membership            string               `json:"membership"`
	ManagementInterfaces  []netapp.IPInterface `json:"management_interfaces"`
	ClusterInterfaces     []netapp.IPInterface `json:"cluster_interfaces"`
	StorageConfiguration  string               `json:"storage_configuration"`
	SystemAggregate       NodeAggregate        `json:"system_aggregate"`
	Controller            NodeController       `json:"controller"`
	HA                    NodeHA               `json:"ha"`
	ServiceProcessor      NodeServiceProcessor `json:"service_processor"`
	NVRAM                 NodeNVRAM            `json:"nvram"`
	ExternalCache         NodeExternalCache    `json:"external_cache"`
	HWAssist              NodeHWAssist         `json:"hw_assist"`
	AntiRansomwareVersion string               `json:"anti_ransomware_version"`
	Metric                netapp.Metrics       `json:"metric"`
	Statistics            netapp.Statistics    `json:"statistics"`
}

type NodeVersion struct {
	Full       string `json:"full"`
	Generation int    `json:"generation"`
	Major      int    `json:"major"`
	Minor      int    `json:"minor"`
}

type NodeAggregate struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type NodeController struct {
	Board             string                  `json:"board"`
	MemorySize        int64                   `json:"memory_size"`
	OverTemperature   string                  `json:"over_temperature"`
	FailedFan         NodeControllerComponent `json:"failed_fan"`
	FailedPowerSupply NodeControllerComponent `json:"failed_power_supply"`
	CPU               NodeCPU                 `json:"cpu"`
}

type NodeControllerComponent struct {
	Count   int                   `json:"count"`
	Message NodeControllerMessage `json:"message"`
}

type NodeControllerMessage struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type NodeCPU struct {
	FirmwareRelease string `json:"firmware_release"`
	Processor       string `json:"processor"`
	Count           int    `json:"count"`
}

type NodeHA struct {
	Enabled      bool               `json:"enabled"`
	AutoGiveback bool               `json:"auto_giveback"`
	Partners     []NodeHAPartner    `json:"partners"`
	Giveback     NodeHAGiveback     `json:"giveback"`
	Takeover     NodeHATakeover     `json:"takeover"`
	Interconnect NodeHAInterconnect `json:"interconnect"`
	Ports        []NodeHAPort       `json:"ports"`
}

type NodeHAPartner struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type NodeHAGiveback struct {
	State  string           `json:"state"`
	Status []NodeHAGBStatus `json:"status"`
}

type NodeHAGBStatus struct {
	State     string           `json:"state"`
	Aggregate NodeAggregateRef `json:"aggregate"`
}

type NodeAggregateRef struct {
	Name string `json:"name"`
}

type NodeHATakeover struct {
	State string `json:"state"`
}

type NodeHAInterconnect struct {
	Adapter string `json:"adapter"`
	State   string `json:"state"`
}

type NodeHAPort struct {
	Number int    `json:"number"`
	State  string `json:"state"`
}

type NodeServiceProcessor struct {
	DHCPEnabled       bool                 `json:"dhcp_enabled"`
	State             string               `json:"state"`
	MACAddress        string               `json:"mac_address"`
	FirmwareVersion   string               `json:"firmware_version"`
	LinkStatus        string               `json:"link_status"`
	Type              string               `json:"type"`
	IsIPConfigured    bool                 `json:"is_ip_configured"`
	AutoupdateEnabled bool                 `json:"autoupdate_enabled"`
	LastUpdateState   string               `json:"last_update_state"`
	IPv4Interface     *NodeSPIPv4Interface `json:"ipv4_interface"`
	IPv6Interface     *NodeSPIPv6Interface `json:"ipv6_interface"`
	SSHInfo           *NodeSPSSHInfo       `json:"ssh_info"`
	Primary           *NodeSPFirmware      `json:"primary"`
	Backup            *NodeSPFirmware      `json:"backup"`
	APIService        *NodeSPAPIService    `json:"api_service"`
	WebService        *NodeSPWebService    `json:"web_service"`
}

type NodeSPIPv4Interface struct {
	Address    string `json:"address"`
	Netmask    string `json:"netmask"`
	Gateway    string `json:"gateway"`
	Enabled    bool   `json:"enabled"`
	SetupState string `json:"setup_state"`
}

type NodeSPIPv6Interface struct {
	Enabled bool `json:"enabled"`
}

type NodeSPSSHInfo struct {
	AllowedAddresses []string `json:"allowed_addresses"`
}

type NodeSPFirmware struct {
	IsCurrent bool   `json:"is_current"`
	State     string `json:"state"`
	Version   string `json:"version"`
}

type NodeSPAPIService struct {
	Enabled     bool `json:"enabled"`
	LimitAccess bool `json:"limit_access"`
	Port        int  `json:"port"`
}

type NodeSPWebService struct {
	Enabled     bool `json:"enabled"`
	LimitAccess bool `json:"limit_access"`
}

type NodeNVRAM struct {
	ID           int64  `json:"id"`
	BatteryState string `json:"battery_state"`
}

type NodeExternalCache struct {
	IsEnabled       bool `json:"is_enabled"`
	IsHYAEnabled    bool `json:"is_hya_enabled"`
	IsRewarmEnabled bool `json:"is_rewarm_enabled"`
	PCSSize         int  `json:"pcs_size"`
}

type NodeHWAssist struct {
	Status NodeHWAssistStatus `json:"status"`
}

type NodeHWAssistStatus struct {
	Enabled bool                `json:"enabled"`
	Local   NodeHWAssistPartner `json:"local"`
	Partner NodeHWAssistPartner `json:"partner"`
}

type NodeHWAssistPartner struct {
	State string `json:"state"`
	IP    string `json:"ip"`
	Port  int    `json:"port"`
}

// /api/cluster/sensors?type=voltage,thermal,battery-life,fan

type Sensor struct {
	Node                  netapp.NamedObject `json:"node"`
	Index                 int                `json:"index"`
	Name                  string             `json:"name"`
	Type                  string             `json:"type"`
	Value                 int                `json:"value"`
	ValueUnits            string             `json:"value_units"`
	ThresholdState        string             `json:"threshold_state"`
	CriticalLowThreshold  *int               `json:"critical_low_threshold,omitempty"`
	WarningLowThreshold   *int               `json:"warning_low_threshold,omitempty"`
	WarningHighThreshold  *int               `json:"warning_high_threshold,omitempty"`
	CriticalHighThreshold *int               `json:"critical_high_threshold,omitempty"`
}

// endpoint": "/api/cluster/counter/tables/host_adapter/rows

type CounterTableRow struct {
	CounterTable CounterTableRef   `json:"counter_table"`
	ID           string            `json:"id"`
	Properties   []CounterProperty `json:"properties"`
	Counters     []CounterValue    `json:"counters"`
}

type CounterTableRef struct {
	Name string `json:"name"`
}

type CounterProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CounterValue struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type CounterType string

// CounterSchema represents a single counter's schema.
type CounterSchema struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        CounterType `json:"type"`
	Unit        string      `json:"unit"`
	// denominator is populated for types like "percent" or "average"
	Denominator *struct {
		Name string `json:"name"`
	} `json:"denominator,omitempty"`
}

// CounterTable represents a table of counters.
type CounterTable struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	CounterSchemas []CounterSchema `json:"counter_schemas"`
}

// Peers
type Authentication struct {
	ExpiryTime string `json:"expiry_time"`
	InUse      string `json:"in_use"`
	Passphrase string `json:"passphrase"`
	State      string `json:"state"`
}

type Encryption struct {
	Proposed string `json:"proposed"`
	State    string `json:"state"`
}

type Remote struct {
	IPAddresses  []string `json:"ip_addresses"`
	Name         string   `json:"name"`
	SerialNumber string   `json:"serial_number"`
}

type Status struct {
	State      string `json:"state"`
	UpdateTime string `json:"update_time"`
}

type Version struct {
	Full       string `json:"full"`
	Generation int    `json:"generation"`
	Major      int    `json:"major"`
	Minor      int    `json:"minor"`
}

type ClusterPeer struct {
	Authentication     Authentication       `json:"authentication"`
	Encryption         Encryption           `json:"encryption"`
	InitialAllowedSVMs []netapp.NamedObject `json:"initial_allowed_svms"`
	IPAddress          string               `json:"ip_address"`
	Ipspace            netapp.NamedObject   `json:"ipspace"`
	Name               string               `json:"name"`
	PeerApplications   []string             `json:"peer_applications"`
	Remote             Remote               `json:"remote"`
	Status             Status               `json:"status"`
	UUID               string               `json:"uuid"`
	Version            Version              `json:"version"`
}
