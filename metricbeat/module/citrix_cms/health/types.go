package health

import "time"

type Serial string

type ErrorResponse struct {
	Error struct {
		ErrorCode      int `json:"errorCode"`
		HTTPStatusCode int `json:"httpStatusCode"`
		Messages       []struct {
			EnUS string `json:"en-US"`
		} `json:"messages"`
		Created time.Time `json:"created"`
	} `json:"error"`
}

type CMSData struct {
	machineCurrentLoadIndex MachinesResponse_JSON
}

type MachinesResponse_JSON struct {
	OdataContext string    `json:"@odata.context"`
	Value        []Machine `json:"value"`
	Message      string
}

type Machine struct {
	ID                  string  `json:"Id"`
	Sid                 string  `json:"Sid"`
	Name                string  `json:"Name"`
	DnsName             string  `json:"DnsName"`
	LifecycleState      int     `json:"LifecycleState"`
	IPAddress           *string `json:"IPAddress"`
	HostedMachineId     string  `json:"HostedMachineId"`
	HostingServerName   string  `json:"HostingServerName"`
	HostedMachineName   string  `json:"HostedMachineName"`
	IsAssigned          bool    `json:"IsAssigned"`
	IsInMaintenanceMode bool    `json:"IsInMaintenanceMode"`
	IsPendingUpdate     bool    `json:"IsPendingUpdate"`
	IsPreparing         bool    `json:"IsPreparing"`
	CurrentSessionCount int     `json:"CurrentSessionCount"`
	CurrentPowerState   int     `json:"CurrentPowerState"`
	FaultState          int     `json:"FaultState"`
	// AgentVersion                 string           `json:"AgentVersion"`
	// AssociatedUserFullNames      string           `json:"AssociatedUserFullNames"`
	// AssociatedUserNames          string           `json:"AssociatedUserNames"`
	// AssociatedUserUPNs           string           `json:"AssociatedUserUPNs"`
	// CurrentRegistrationState     int              `json:"CurrentRegistrationState"`
	// RegistrationStateChangeDate  time.Time        `json:"RegistrationStateChangeDate"`
	// LastDeregisteredCode         int              `json:"LastDeregisteredCode"`
	// LastDeregisteredDate         time.Time        `json:"LastDeregisteredDate"`
	// ControllerDnsName            string           `json:"ControllerDnsName"`
	// PoweredOnDate                time.Time        `json:"PoweredOnDate"`
	// FailureDate                  *time.Time       `json:"FailureDate"`
	// PowerStateChangeDate         time.Time        `json:"PowerStateChangeDate"`
	// FunctionalLevel              int              `json:"FunctionalLevel"`
	// WindowsConnectionSetting     int              `json:"WindowsConnectionSetting"`
	// OSType                       string           `json:"OSType"`
	// CurrentLoadIndexId           int              `json:"CurrentLoadIndexId"`
	// CatalogId                    string           `json:"CatalogId"`
	// DesktopGroupId               string           `json:"DesktopGroupId"`
	// HypervisorId                 string           `json:"HypervisorId"`
	// LastPowerActionType          int              `json:"LastPowerActionType"`
	// LastPowerActionReason        int              `json:"LastPowerActionReason"`
	// LastPowerActionFailureReason int              `json:"LastPowerActionFailureReason"`
	// LastPowerActionCompletedDate time.Time        `json:"LastPowerActionCompletedDate"`
	// LastUpgradeState             *string          `json:"LastUpgradeState"`
	// LastUpgradeStateChangeDate   *time.Time       `json:"LastUpgradeStateChangeDate"`
	// Tags                         []string         `json:"Tags"`
	// Hash                         string           `json:"Hash"`
	MachineRole      int              `json:"MachineRole"`
	CreatedDate      time.Time        `json:"CreatedDate"`
	ModifiedDate     time.Time        `json:"ModifiedDate"`
	CurrentLoadIndex MachineLoadIndex `json:"CurrentLoadIndex"`
}
type MachineLoadIndex struct {
	ID                 int       `json:"Id"`
	EffectiveLoadIndex int       `json:"EffectiveLoadIndex"`
	Cpu                int       `json:"Cpu"`
	Memory             int       `json:"Memory"`
	Disk               int       `json:"Disk"`
	Network            int       `json:"Network"`
	SessionCount       int       `json:"SessionCount"`
	MachineId          string    `json:"MachineId"`
	CreatedDate        time.Time `json:"CreatedDate"`
	// ModifiedDate       time.Time `json:"ModifiedDate"`
}
