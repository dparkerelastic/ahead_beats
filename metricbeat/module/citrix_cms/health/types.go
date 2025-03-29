package health

import "time"

type Serial string

// type ErrorResponse struct {
// 	Error struct {
// 		ErrorCode      int `json:"errorCode"`
// 		HTTPStatusCode int `json:"httpStatusCode"`
// 		Messages       []struct {
// 			EnUS string `json:"en-US"`
// 		} `json:"messages"`
// 		Created time.Time `json:"created"`
// 	} `json:"error"`
// }

type CMSData struct {
	machineCurrentLoadIndex  MachineCurrentLoadIndex_JSON
	serverOSDesktopSummaries ServerOSDesktopSummaries_JSON
	loadIndexSummaries       LoadIndexSummaries_JSON
	machineMetric            MachineMetric_JSON
	machineDetails           MachineDetails_JSON
}

type MachineCurrentLoadIndex_JSON struct {
	OdataContext string `json:"@odata.context"`
	OdataCount   int    `json:"@odata.count"`
	Value        []struct {
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
		CurrentRegistrationState int `json:"CurrentRegistrationState"`
		// RegistrationStateChangeDate  time.Time        `json:"RegistrationStateChangeDate"`
		LastDeregisteredCode int `json:"LastDeregisteredCode"`
		// LastDeregisteredDate         time.Time        `json:"LastDeregisteredDate"`
		// ControllerDnsName            string           `json:"ControllerDnsName"`
		// PoweredOnDate                time.Time        `json:"PoweredOnDate"`
		// FailureDate                  *time.Time       `json:"FailureDate"`
		// PowerStateChangeDate         time.Time        `json:"PowerStateChangeDate"`
		FunctionalLevel          int `json:"FunctionalLevel"`
		WindowsConnectionSetting int `json:"WindowsConnectionSetting"`
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
		MachineRole      int       `json:"MachineRole"`
		CreatedDate      time.Time `json:"CreatedDate"`
		ModifiedDate     time.Time `json:"ModifiedDate"`
		CurrentLoadIndex struct {
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
		} `json:"CurrentLoadIndex"`
	} `json:"value"`
	Message string
}

type ServerOSDesktopSummaries_JSON struct {
	OdataContext string `json:"@odata.context"`
	OdataCount   int    `json:"@odata.count"`
	Value        []struct {
		ID                          int       `json:"Id"`
		SummaryDate                 time.Time `json:"SummaryDate"`
		DesktopGroupId              string    `json:"DesktopGroupId"`
		PeakConcurrentInstanceCount int       `json:"PeakConcurrentInstanceCount"`
		TotalUsageDuration          int       `json:"TotalUsageDuration"`
		TotalLaunchesCount          int       `json:"TotalLaunchesCount"`
		StartingInstanceCount       int       `json:"StartingInstanceCount"`
		Granularity                 int       `json:"Granularity"`
		CreatedDate                 time.Time `json:"CreatedDate"`
		ModifiedDate                time.Time `json:"ModifiedDate"`
	} `json:"value"`
	Message string
}

type LoadIndexSummaries_JSON struct {
	OdataContext string `json:"@odata.context"`
	OdataCount   int    `json:"@odata.count"`
	Value        []struct {
		ID              int       `json:"Id"`
		SummaryDate     time.Time `json:"SummaryDate"`
		MachineId       string    `json:"MachineId"`
		SumCount        int       `json:"SumCount"`
		SumLoadIndex    int       `json:"SumLoadIndex"`
		SumCpu          int       `json:"SumCpu"`
		SumNetwork      int       `json:"SumNetwork"`
		SumDisk         int       `json:"SumDisk"`
		SumMemory       int       `json:"SumMemory"`
		SumSessionCount int       `json:"SumSessionCount"`
		Granularity     int       `json:"Granularity"`
		CreatedDate     time.Time `json:"CreatedDate"`
		ModifiedDate    time.Time `json:"ModifiedDate"`
	} `json:"value"`
	Message string
}

type MachineMetric_JSON struct {
	OdataContext string `json:"@odata.context"`
	OdataCount   int    `json:"@odata.count"`
	Value        []struct {
		MachineID     string    `json:"MachineId"`
		CollectedDate time.Time `json:"CollectedDate"`
		Iops          float64   `json:"Iops"`    // Changed to float64 to allow for decimal values
		Latency       float64   `json:"Latency"` // Changed to float64 to allow for decimal values
	} `json:"value"`
	Message string
}

type MachineDetails_JSON struct {
	OdataContext string `json:"@odata.context"`
	OdataCount   int    `json:"@odata.count"`
	Message      string
	Value        []struct {
		ID                           string    `json:"Id"`
		Sid                          string    `json:"Sid"`
		Name                         string    `json:"Name"`
		DnsName                      string    `json:"DnsName"`
		LifecycleState               int       `json:"LifecycleState"`
		IPAddress                    string    `json:"IPAddress"`
		HostedMachineId              string    `json:"HostedMachineId"`
		HostingServerName            string    `json:"HostingServerName"`
		HostedMachineName            string    `json:"HostedMachineName"`
		IsAssigned                   bool      `json:"IsAssigned"`
		IsInMaintenanceMode          bool      `json:"IsInMaintenanceMode"`
		IsPendingUpdate              bool      `json:"IsPendingUpdate"`
		AgentVersion                 string    `json:"AgentVersion"`
		AssociatedUserFullNames      string    `json:"AssociatedUserFullNames"`
		AssociatedUserNames          string    `json:"AssociatedUserNames"`
		AssociatedUserUPNs           string    `json:"AssociatedUserUPNs"`
		CurrentRegistrationState     int       `json:"CurrentRegistrationState"`
		RegistrationStateChangeDate  time.Time `json:"RegistrationStateChangeDate"`
		LastDeregisteredCode         int       `json:"LastDeregisteredCode"`
		LastDeregisteredDate         time.Time `json:"LastDeregisteredDate"`
		CurrentPowerState            int       `json:"CurrentPowerState"`
		CurrentSessionCount          int       `json:"CurrentSessionCount"`
		ControllerDnsName            string    `json:"ControllerDnsName"`
		PoweredOnDate                time.Time `json:"PoweredOnDate"`
		PowerStateChangeDate         time.Time `json:"PowerStateChangeDate"`
		FunctionalLevel              int       `json:"FunctionalLevel"`
		FailureDate                  time.Time `json:"FailureDate"`
		WindowsConnectionSetting     int       `json:"WindowsConnectionSetting"`
		IsPreparing                  bool      `json:"IsPreparing"`
		FaultState                   int       `json:"FaultState"`
		OSType                       string    `json:"OSType"`
		CurrentLoadIndexId           int       `json:"CurrentLoadIndexId"`
		CatalogId                    string    `json:"CatalogId"`
		DesktopGroupId               string    `json:"DesktopGroupId"`
		HypervisorId                 string    `json:"HypervisorId"`
		LastPowerActionCompletedDate time.Time `json:"LastPowerActionCompletedDate"`
		LastUpgradeState             int       `json:"LastUpgradeState"`
		LastUpgradeStateChangeDate   time.Time `json:"LastUpgradeStateChangeDate"`
		Hash                         string    `json:"Hash"`
		MachineRole                  int       `json:"MachineRole"`
		CreatedDate                  time.Time `json:"CreatedDate"`
		ModifiedDate                 time.Time `json:"ModifiedDate"`
		// CurrentLoadIndex             struct {
		// 	ID                 int       `json:"Id"`
		// 	EffectiveLoadIndex int       `json:"EffectiveLoadIndex"`
		// 	Cpu                int       `json:"Cpu"`
		// 	Memory             int       `json:"Memory"`
		// 	Disk               int       `json:"Disk"`
		// 	Network            int       `json:"Network"`
		// 	SessionCount       int       `json:"SessionCount"`
		// 	MachineId          string    `json:"MachineId"`
		// 	CreatedDate        time.Time `json:"CreatedDate"`
		// 	ModifiedDate       time.Time `json:"ModifiedDate"`
		// } `json:"CurrentLoadIndex"`
		// Catalog struct {
		// 	ID                    string        `json:"Id"`
		// 	Name                  string        `json:"Name"`
		// 	LifecycleState        int           `json:"LifecycleState"`
		// 	ProvisioningType      int           `json:"ProvisioningType"`
		// 	PersistentUserChanges int           `json:"PersistentUserChanges"`
		// 	IsMachinePhysical     bool          `json:"IsMachinePhysical"`
		// 	AllocationType        int           `json:"AllocationType"`
		// 	SessionSupport        int           `json:"SessionSupport"`
		// 	ProvisioningSchemeId  string        `json:"ProvisioningSchemeId"`
		// 	ZoneUid               string        `json:"ZoneUid"`
		// 	ZoneName              string        `json:"ZoneName"`
		// 	CreatedDate           time.Time     `json:"CreatedDate"`
		// 	ModifiedDate          time.Time     `json:"ModifiedDate"`
		// 	Machines              []interface{} `json:"Machines"`
		// } `json:"Catalog"`
		// Hypervisor struct {
		// 	ID             string        `json:"Id"`
		// 	Name           string        `json:"Name"`
		// 	Type           string        `json:"Type"`
		// 	LifecycleState int           `json:"LifecycleState"`
		// 	CreatedDate    time.Time     `json:"CreatedDate"`
		// 	ModifiedDate   time.Time     `json:"ModifiedDate"`
		// 	Machines       []interface{} `json:"Machines"`
		// } `json:"Hypervisor"`
		// MachineCost struct {
		// 	MachineId                  string    `json:"MachineId"`
		// 	SpecId                     int       `json:"SpecId"`
		// 	CostPerHour                int       `json:"CostPerHour"`
		// 	PowerOnComputeCostPerHour  int       `json:"PowerOnComputeCostPerHour"`
		// 	PowerOnStorageCostPerHour  int       `json:"PowerOnStorageCostPerHour"`
		// 	PowerOffStorageCostPerHour int       `json:"PowerOffStorageCostPerHour"`
		// 	CreatedDate                time.Time `json:"CreatedDate"`
		// 	ModifiedDate               time.Time `json:"ModifiedDate"`
		// } `json:"MachineCost"`
		// LoadIndex []struct {
		// 	ID                 int       `json:"Id"`
		// 	EffectiveLoadIndex int       `json:"EffectiveLoadIndex"`
		// 	Cpu                int       `json:"Cpu"`
		// 	Memory             int       `json:"Memory"`
		// 	Disk               int       `json:"Disk"`
		// 	Network            int       `json:"Network"`
		// 	SessionCount       int       `json:"SessionCount"`
		// 	MachineId          string    `json:"MachineId"`
		// 	CreatedDate        time.Time `json:"CreatedDate"`
		// 	ModifiedDate       time.Time `json:"ModifiedDate"`
		// } `json:"LoadIndex"`
		// ResourceUtilization []struct {
		// 	MachineId      string    `json:"MachineId"`
		// 	CollectedDate  time.Time `json:"CollectedDate"`
		// 	PercentCpu     int       `json:"PercentCpu"`
		// 	UsedMemory     int       `json:"UsedMemory"`
		// 	TotalMemory    int       `json:"TotalMemory"`
		// 	CreatedDate    time.Time `json:"CreatedDate"`
		// 	ModifiedDate   time.Time `json:"ModifiedDate"`
		// 	DesktopGroupId string    `json:"DesktopGroupId"`
		// } `json:"ResourceUtilization"`
		// MachineFailures []struct {
		// 	ID                   int       `json:"Id"`
		// 	MachineId            string    `json:"MachineId"`
		// 	FailureStartDate     time.Time `json:"FailureStartDate"`
		// 	FailureEndDate       time.Time `json:"FailureEndDate"`
		// 	FaultState           int       `json:"FaultState"`
		// 	LastDeregisteredCode int       `json:"LastDeregisteredCode"`
		// 	CreatedDate          time.Time `json:"CreatedDate"`
		// 	ModifiedDate         time.Time `json:"ModifiedDate"`
		// } `json:"MachineFailures"`
		// MachineMetric []struct {
		// 	MachineId     string    `json:"MachineId"`
		// 	CollectedDate time.Time `json:"CollectedDate"`
		// 	Iops          int       `json:"Iops"`
		// 	Latency       int       `json:"Latency"`
		// } `json:"MachineMetric"`
		// MachineMetricSummary []struct {
		// 	MachineId   string    `json:"MachineId"`
		// 	SummaryDate time.Time `json:"SummaryDate"`
		// 	AvgIops     int       `json:"AvgIops"`
		// 	PeakIops    int       `json:"PeakIops"`
		// 	AvgLatency  int       `json:"AvgLatency"`
		// } `json:"MachineMetricSummary"`
		// MachineHotfixLogs []struct {
		// 	ID           int       `json:"Id"`
		// 	MachineId    string    `json:"MachineId"`
		// 	HotfixId     string    `json:"HotfixId"`
		// 	ChangeType   int       `json:"ChangeType"`
		// 	CurrentState bool      `json:"CurrentState"`
		// 	CreatedDate  time.Time `json:"CreatedDate"`
		// 	ModifiedDate time.Time `json:"ModifiedDate"`
		// 	Hotfix       struct {
		// 		ID                string        `json:"Id"`
		// 		Name              string        `json:"Name"`
		// 		Article           string        `json:"Article"`
		// 		ArticleName       string        `json:"ArticleName"`
		// 		FileName          string        `json:"FileName"`
		// 		FileFormat        string        `json:"FileFormat"`
		// 		Version           string        `json:"Version"`
		// 		ComponentName     string        `json:"ComponentName"`
		// 		ComponentVersion  string        `json:"ComponentVersion"`
		// 		CreatedDate       time.Time     `json:"CreatedDate"`
		// 		ModifiedDate      time.Time     `json:"ModifiedDate"`
		// 		MachineHotfixLogs []interface{} `json:"MachineHotfixLogs"`
		// 	} `json:"Hotfix"`
		// } `json:"MachineHotfixLogs"`
		// ApplicationFaults []struct {
		// 	ID                      int       `json:"Id"`
		// 	FaultingApplicationPath string    `json:"FaultingApplicationPath"`
		// 	ProcessName             string    `json:"ProcessName"`
		// 	SessionKey              string    `json:"SessionKey"`
		// 	Version                 string    `json:"Version"`
		// 	Description             string    `json:"Description"`
		// 	FaultReportedDate       time.Time `json:"FaultReportedDate"`
		// 	BrowserNames            string    `json:"BrowserNames"`
		// 	MachineId               string    `json:"MachineId"`
		// } `json:"ApplicationFaults"`
		// ApplicationErrors []struct {
		// 	ID                      int       `json:"Id"`
		// 	FaultingApplicationPath string    `json:"FaultingApplicationPath"`
		// 	ProcessName             string    `json:"ProcessName"`
		// 	SessionKey              string    `json:"SessionKey"`
		// 	Version                 string    `json:"Version"`
		// 	Description             string    `json:"Description"`
		// 	ErrorReportedDate       time.Time `json:"ErrorReportedDate"`
		// 	BrowserNames            string    `json:"BrowserNames"`
		// 	MachineId               string    `json:"MachineId"`
		// } `json:"ApplicationErrors"`
		// HistoricalApplicationLaunches []struct {
		// 	ID                int       `json:"Id"`
		// 	AppId             string    `json:"AppId"`
		// 	ProcessName       string    `json:"ProcessName"`
		// 	StartDate         time.Time `json:"StartDate"`
		// 	EndDate           time.Time `json:"EndDate"`
		// 	AppVersion        string    `json:"AppVersion"`
		// 	ExecutablePath    string    `json:"ExecutablePath"`
		// 	UserName          string    `json:"UserName"`
		// 	MachineId         string    `json:"MachineId"`
		// 	DesktopGroupId    string    `json:"DesktopGroupId"`
		// 	SessionGUID       string    `json:"SessionGUID"`
		// 	ProcessGUID       string    `json:"ProcessGUID"`
		// 	StartupTimeMs     int       `json:"StartupTimeMs"`
		// 	StartupIOPS       int       `json:"StartupIOPS"`
		// 	ProcessParentGUID string    `json:"ProcessParentGUID"`
		// 	ProcessParentName string    `json:"ProcessParentName"`
		// 	ModifiedDate      time.Time `json:"ModifiedDate"`
		// 	IsPublishedApp    bool      `json:"IsPublishedApp"`
		// } `json:"HistoricalApplicationLaunches"`
		// ActiveApplicationLaunches []struct {
		// 	ID                int       `json:"Id"`
		// 	AppId             string    `json:"AppId"`
		// 	ProcessName       string    `json:"ProcessName"`
		// 	StartDate         time.Time `json:"StartDate"`
		// 	AppVersion        string    `json:"AppVersion"`
		// 	ExecutablePath    string    `json:"ExecutablePath"`
		// 	UserName          string    `json:"UserName"`
		// 	MachineId         string    `json:"MachineId"`
		// 	DesktopGroupId    string    `json:"DesktopGroupId"`
		// 	SessionGUID       string    `json:"SessionGUID"`
		// 	ProcessGUID       string    `json:"ProcessGUID"`
		// 	StartupTimeMs     int       `json:"StartupTimeMs"`
		// 	StartupIOPS       int       `json:"StartupIOPS"`
		// 	ProcessParentGUID string    `json:"ProcessParentGUID"`
		// 	ProcessParentName string    `json:"ProcessParentName"`
		// 	ModifiedDate      time.Time `json:"ModifiedDate"`
		// 	IsPublishedApp    bool      `json:"IsPublishedApp"`
		// 	IsAppEnded        bool      `json:"IsAppEnded"`
		// } `json:"ActiveApplicationLaunches"`
	} `json:"value"`
}
