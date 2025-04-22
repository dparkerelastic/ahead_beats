package health

import (
	"time"
)

type Serial string

type CMSData struct {
	machineCurrentLoadIndex       MachineCurrentLoadIndex_JSON
	serverOSDesktopSummaries      ServerOSDesktopSummaries_JSON
	loadIndexSummaries            LoadIndexSummaries_JSON
	machineMetric                 MachineMetric_JSON
	sessionActivitySummaries_Agg1 SessionActivitySummaries_Agg1_JSON
	machineMetricDetails          MachineMetricDetails_JSON
	//logonMetricDetails            LogonMetricsDetails_JSON
	sessionMetricsDetails SessionMetricsDetails_JSON
	sessionDetails        SessionsDetails_JSON
	sessionFailureDetails SessionsDetails_JSON
}

type MachineCurrentLoadIndex_JSON struct {
	OdataContext string `json:"@odata.context,omitempty"`
	OdataCount   int    `json:"@odata.count,omitempty"`
	Message      string
	Value        []struct {
		ID                  string  `json:"Id,omitempty"`
		Sid                 string  `json:"Sid,omitempty"`
		Name                string  `json:"Name,omitempty"`
		DnsName             string  `json:"DnsName,omitempty"`
		LifecycleState      int     `json:"LifecycleState,omitempty"`
		IPAddress           *string `json:"IPAddress,omitempty"`
		HostedMachineId     string  `json:"HostedMachineId,omitempty"`
		HostingServerName   string  `json:"HostingServerName,omitempty"`
		HostedMachineName   string  `json:"HostedMachineName,omitempty"`
		IsAssigned          bool    `json:"IsAssigned,omitempty"`
		IsInMaintenanceMode bool    `json:"IsInMaintenanceMode,omitempty"`
		IsPendingUpdate     bool    `json:"IsPendingUpdate,omitempty"`
		IsPreparing         bool    `json:"IsPreparing,omitempty"`
		CurrentSessionCount int     `json:"CurrentSessionCount,omitempty"`
		CurrentPowerState   int     `json:"CurrentPowerState,omitempty"`
		FaultState          int     `json:"FaultState,omitempty"`
		// AgentVersion                 string           `json:"AgentVersion,omitempty"`
		// AssociatedUserFullNames      string           `json:"AssociatedUserFullNames,omitempty"`
		// AssociatedUserNames          string           `json:"AssociatedUserNames,omitempty"`
		// AssociatedUserUPNs           string           `json:"AssociatedUserUPNs,omitempty"`
		CurrentRegistrationState int `json:"CurrentRegistrationState,omitempty"`
		// RegistrationStateChangeDate *time.Time        `json:"RegistrationStateChangeDate,omitempty"`
		LastDeregisteredCode int `json:"LastDeregisteredCode,omitempty"`
		// LastDeregisteredDate        *time.Time        `json:"LastDeregisteredDate,omitempty"`
		// ControllerDnsName            string           `json:"ControllerDnsName,omitempty"`
		// PoweredOnDate               *time.Time        `json:"PoweredOnDate,omitempty"`
		// FailureDate                 *time.Time       `json:"FailureDate,omitempty"`
		// PowerStateChangeDate        *time.Time        `json:"PowerStateChangeDate,omitempty"`
		FunctionalLevel          int `json:"FunctionalLevel,omitempty"`
		WindowsConnectionSetting int `json:"WindowsConnectionSetting,omitempty"`
		// OSType                       string           `json:"OSType,omitempty"`
		// CurrentLoadIndexId           int              `json:"CurrentLoadIndexId,omitempty"`
		// CatalogId                    string           `json:"CatalogId,omitempty"`
		// DesktopGroupId               string           `json:"DesktopGroupId,omitempty"`
		// HypervisorId                 string           `json:"HypervisorId,omitempty"`
		// LastPowerActionType          int              `json:"LastPowerActionType,omitempty"`
		// LastPowerActionReason        int              `json:"LastPowerActionReason,omitempty"`
		// LastPowerActionFailureReason int              `json:"LastPowerActionFailureReason,omitempty"`
		// LastPowerActionCompletedDate*time.Time        `json:"LastPowerActionCompletedDate,omitempty"`
		// LastUpgradeState             *string          `json:"LastUpgradeState,omitempty"`
		// LastUpgradeStateChangeDate   *time.Time       `json:"LastUpgradeStateChangeDate,omitempty"`
		// Tags                         []string         `json:"Tags,omitempty"`
		// Hash                         string           `json:"Hash,omitempty"`
		MachineRole      int        `json:"MachineRole,omitempty"`
		CreatedDate      *time.Time `json:"CreatedDate,omitempty"`
		ModifiedDate     *time.Time `json:"ModifiedDate,omitempty"`
		CurrentLoadIndex struct {
			ID                 int        `json:"Id,omitempty"`
			EffectiveLoadIndex int        `json:"EffectiveLoadIndex,omitempty"`
			Cpu                int        `json:"Cpu,omitempty"`
			Memory             int        `json:"Memory,omitempty"`
			Disk               int        `json:"Disk,omitempty"`
			Network            int        `json:"Network,omitempty"`
			SessionCount       int        `json:"SessionCount,omitempty"`
			MachineId          string     `json:"MachineId,omitempty"`
			CreatedDate        *time.Time `json:"CreatedDate,omitempty"`
			// ModifiedDate      *time.Time `json:"ModifiedDate,omitempty"`
		} `json:"CurrentLoadIndex,omitempty"`
	} `json:"value,omitempty"`
}

type ServerOSDesktopSummaries_JSON struct {
	OdataContext string `json:"@odata.context,omitempty"`
	OdataCount   int    `json:"@odata.count,omitempty"`
	Message      string
	Value        []struct {
		ID                          int        `json:"Id,omitempty"`
		SummaryDate                 *time.Time `json:"SummaryDate,omitempty"`
		DesktopGroupId              string     `json:"DesktopGroupId,omitempty"`
		PeakConcurrentInstanceCount int        `json:"PeakConcurrentInstanceCount,omitempty"`
		TotalUsageDuration          int        `json:"TotalUsageDuration,omitempty"`
		TotalLaunchesCount          int        `json:"TotalLaunchesCount,omitempty"`
		StartingInstanceCount       int        `json:"StartingInstanceCount,omitempty"`
		Granularity                 int        `json:"Granularity,omitempty"`
		CreatedDate                 *time.Time `json:"CreatedDate,omitempty"`
		ModifiedDate                *time.Time `json:"ModifiedDate,omitempty"`
	} `json:"value,omitempty"`
}

type LoadIndexSummaries_JSON struct {
	OdataContext string `json:"@odata.context,omitempty"`
	OdataCount   int    `json:"@odata.count,omitempty"`
	Message      string
	Value        []struct {
		ID              int        `json:"Id,omitempty"`
		SummaryDate     *time.Time `json:"SummaryDate,omitempty"`
		MachineId       string     `json:"MachineId,omitempty"`
		SumCount        int        `json:"SumCount,omitempty"`
		SumLoadIndex    int        `json:"SumLoadIndex,omitempty"`
		SumCpu          int        `json:"SumCpu,omitempty"`
		SumNetwork      int        `json:"SumNetwork,omitempty"`
		SumDisk         int        `json:"SumDisk,omitempty"`
		SumMemory       int        `json:"SumMemory,omitempty"`
		SumSessionCount int        `json:"SumSessionCount,omitempty"`
		Granularity     int        `json:"Granularity,omitempty"`
		CreatedDate     *time.Time `json:"CreatedDate,omitempty"`
		ModifiedDate    *time.Time `json:"ModifiedDate,omitempty"`
	} `json:"value,omitempty"`
}

type MachineMetric_JSON struct {
	OdataContext string `json:"@odata.context,omitempty"`
	OdataCount   int    `json:"@odata.count,omitempty"`
	Message      string
	Value        []struct {
		MachineID     string     `json:"MachineId,omitempty"`
		CollectedDate *time.Time `json:"CollectedDate,omitempty"`
		Iops          float64    `json:"Iops,omitempty"`    // Changed to float64 to allow for decimal values
		Latency       float64    `json:"Latency,omitempty"` // Changed to float64 to allow for decimal values
	} `json:"value,omitempty"`
}

type SessionActivitySummaries_Agg1_JSON struct {
	OdataContext string `json:"@odata.context,omitempty"`
	OdataCount   int    `json:"@odata.count,omitempty"`
	Message      string
	Value        []struct {
		OdataID                         *string `json:"@odata.id,omitempty"`
		DesktopGroupId                  string  `json:"DesktopGroupId,omitempty"`
		SumTotalLogOnCount              int     `json:"SumTotalLogOnCount,omitempty"`
		AverageTotalLogOnDuration       float64 `json:"AverageTotalLogOnDuration,omitempty"`
		AverageConcurrentSessionCount   float64 `json:"AverageConcurrentSessionCount,omitempty"`
		AverageDisconnectedSessionCount float64 `json:"AverageDisconnectedSessionCount,omitempty"`
		AverageConnectedSessionCount    float64 `json:"AverageConnectedSessionCount,omitempty"`
	} `json:"value,omitempty"`
}

// AutoGenerated represents the structure of the JSON response from the Citrix CMS API for health data.
// type LogonMetricsDetails_JSON struct {
// 	OdataContext string `json:"@odata.context,omitempty"`
// 	OdataCount   int    `json:"@odata.count,omitempty"`
// 	Message      string `json:"Message,omitempty"`
// 	Value        []struct {
// 		ID                int        `json:"Id,omitempty"`
// 		SessionKey        string     `json:"SessionKey,omitempty"`
// 		UserInitStartDate *time.Time `json:"UserInitStartDate,omitempty"`
// 		UserInitEndDate   *time.Time `json:"UserInitEndDate,omitempty"`
// 		Session           struct {
// 			SessionKey                string     `json:"SessionKey,omitempty"`
// 			StartDate                 *time.Time `json:"StartDate,omitempty"`
// 			LogOnDuration             int        `json:"LogOnDuration,omitempty"`
// 			EndDate                   *time.Time `json:"EndDate,omitempty"`
// 			ExitCode                  int        `json:"ExitCode,omitempty"`
// 			FailureID                 int        `json:"FailureId,omitempty"`
// 			FailureDate               *time.Time `json:"FailureDate,omitempty"`
// 			ConnectionState           int        `json:"ConnectionState,omitempty"`
// 			SessionIdleTime           *time.Time `json:"SessionIdleTime,omitempty"`
// 			ConnectionStateChangeDate *time.Time `json:"ConnectionStateChangeDate,omitempty"`
// 			LifecycleState            int        `json:"LifecycleState,omitempty"`
// 			CurrentConnectionID       int        `json:"CurrentConnectionId,omitempty"`
// 			UserID                    int        `json:"UserId,omitempty"`
// 			MachineID                 string     `json:"MachineId,omitempty"`
// 			SessionType               int        `json:"SessionType,omitempty"`
// 			IsAnonymous               bool       `json:"IsAnonymous,omitempty"`
// 			PublishedDesktopID        int        `json:"PublishedDesktopId,omitempty"`
// 			CreatedDate               *time.Time `json:"CreatedDate,omitempty"`
// 			ModifiedDate              *time.Time `json:"ModifiedDate,omitempty"`
// 			Failure                   struct {
// 				ID                         int        `json:"Id,omitempty"`
// 				ConnectionFailureEnumValue int        `json:"ConnectionFailureEnumValue,omitempty"`
// 				Category                   int        `json:"Category,omitempty"`
// 				CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
// 				ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
// 			} `json:"Failure,omitempty"`
// 			CurrentConnection struct {
// 				ID                        int        `json:"Id,omitempty"`
// 				ClientName                string     `json:"ClientName,omitempty"`
// 				ClientAddress             string     `json:"ClientAddress,omitempty"`
// 				ClientPublicIP            string     `json:"ClientPublicIP,omitempty"`
// 				ClientVersion             string     `json:"ClientVersion,omitempty"`
// 				ClientPlatform            string     `json:"ClientPlatform,omitempty"`
// 				ClientISP                 string     `json:"ClientISP,omitempty"`
// 				ClientLocationCountry     string     `json:"ClientLocationCountry,omitempty"`
// 				ClientLocationCity        string     `json:"ClientLocationCity,omitempty"`
// 				ConnectedViaHostName      string     `json:"ConnectedViaHostName,omitempty"`
// 				ConnectedViaIPAddress     string     `json:"ConnectedViaIPAddress,omitempty"`
// 				LaunchedViaHostName       string     `json:"LaunchedViaHostName,omitempty"`
// 				LaunchedViaIPAddress      string     `json:"LaunchedViaIPAddress,omitempty"`
// 				IsReconnect               bool       `json:"IsReconnect,omitempty"`
// 				IsSecureIca               bool       `json:"IsSecureIca,omitempty"`
// 				Protocol                  string     `json:"Protocol,omitempty"`
// 				LogOnStartDate            *time.Time `json:"LogOnStartDate,omitempty"`
// 				LogOnEndDate              *time.Time `json:"LogOnEndDate,omitempty"`
// 				BrokeringDuration         int        `json:"BrokeringDuration,omitempty"`
// 				BrokeringDate             *time.Time `json:"BrokeringDate,omitempty"`
// 				DisconnectCode            int        `json:"DisconnectCode,omitempty"`
// 				DisconnectDate            *time.Time `json:"DisconnectDate,omitempty"`
// 				VMStartStartDate          *time.Time `json:"VMStartStartDate,omitempty"`
// 				VMPoweredOnDate           *time.Time `json:"VMPoweredOnDate,omitempty"`
// 				VMStartEndDate            *time.Time `json:"VMStartEndDate,omitempty"`
// 				ClientSessionValidateDate *time.Time `json:"ClientSessionValidateDate,omitempty"`
// 				ServerSessionValidateDate *time.Time `json:"ServerSessionValidateDate,omitempty"`
// 				EstablishmentDate         *time.Time `json:"EstablishmentDate,omitempty"`
// 				HdxStartDate              *time.Time `json:"HdxStartDate,omitempty"`
// 				HdxEndDate                *time.Time `json:"HdxEndDate,omitempty"`
// 				AuthenticationDuration    int        `json:"AuthenticationDuration,omitempty"`
// 				GpoStartDate              *time.Time `json:"GpoStartDate,omitempty"`
// 				GpoEndDate                *time.Time `json:"GpoEndDate,omitempty"`
// 				LogOnScriptsStartDate     *time.Time `json:"LogOnScriptsStartDate,omitempty"`
// 				LogOnScriptsEndDate       *time.Time `json:"LogOnScriptsEndDate,omitempty"`
// 				ProfileLoadStartDate      *time.Time `json:"ProfileLoadStartDate,omitempty"`
// 				ProfileLoadEndDate        *time.Time `json:"ProfileLoadEndDate,omitempty"`
// 				InteractiveStartDate      *time.Time `json:"InteractiveStartDate,omitempty"`
// 				InteractiveEndDate        *time.Time `json:"InteractiveEndDate,omitempty"`
// 				SessionKey                string     `json:"SessionKey,omitempty"`
// 				CreatedDate               *time.Time `json:"CreatedDate,omitempty"`
// 				ModifiedDate              *time.Time `json:"ModifiedDate,omitempty"`
// 				ConnectionFailureLog      struct {
// 					ID                         int        `json:"Id,omitempty"`
// 					SessionKey                 string     `json:"SessionKey,omitempty"`
// 					FailureDate                *time.Time `json:"FailureDate,omitempty"`
// 					UserID                     int        `json:"UserId,omitempty"`
// 					MachineID                  string     `json:"MachineId,omitempty"`
// 					ConnectionFailureEnumValue int        `json:"ConnectionFailureEnumValue,omitempty"`
// 					CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
// 					ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
// 				} `json:"ConnectionFailureLog,omitempty"`
// 			} `json:"CurrentConnection,omitempty"`
// 			User struct {
// 				ID           int        `json:"Id,omitempty"`
// 				Sid          string     `json:"Sid,omitempty"`
// 				Upn          string     `json:"Upn,omitempty"`
// 				UserName     string     `json:"UserName,omitempty"`
// 				FullName     string     `json:"FullName,omitempty"`
// 				Domain       string     `json:"Domain,omitempty"`
// 				CreatedDate  *time.Time `json:"CreatedDate,omitempty"`
// 				ModifiedDate *time.Time `json:"ModifiedDate,omitempty"`
// 			} `json:"User,omitempty"`
// 			Machine struct {
// 				ID                           string     `json:"Id,omitempty"`
// 				Sid                          string     `json:"Sid,omitempty"`
// 				Name                         string     `json:"Name,omitempty"`
// 				DNSName                      string     `json:"DnsName,omitempty"`
// 				LifecycleState               int        `json:"LifecycleState,omitempty"`
// 				IPAddress                    string     `json:"IPAddress,omitempty"`
// 				HostedMachineID              string     `json:"HostedMachineId,omitempty"`
// 				HostingServerName            string     `json:"HostingServerName,omitempty"`
// 				HostedMachineName            string     `json:"HostedMachineName,omitempty"`
// 				IsAssigned                   bool       `json:"IsAssigned,omitempty"`
// 				IsInMaintenanceMode          bool       `json:"IsInMaintenanceMode,omitempty"`
// 				IsPendingUpdate              bool       `json:"IsPendingUpdate,omitempty"`
// 				AgentVersion                 string     `json:"AgentVersion,omitempty"`
// 				AssociatedUserFullNames      string     `json:"AssociatedUserFullNames,omitempty"`
// 				AssociatedUserNames          string     `json:"AssociatedUserNames,omitempty"`
// 				AssociatedUserUPNs           string     `json:"AssociatedUserUPNs,omitempty"`
// 				CurrentRegistrationState     int        `json:"CurrentRegistrationState,omitempty"`
// 				RegistrationStateChangeDate  *time.Time `json:"RegistrationStateChangeDate,omitempty"`
// 				LastDeregisteredCode         int        `json:"LastDeregisteredCode,omitempty"`
// 				LastDeregisteredDate         *time.Time `json:"LastDeregisteredDate,omitempty"`
// 				CurrentPowerState            int        `json:"CurrentPowerState,omitempty"`
// 				CurrentSessionCount          int        `json:"CurrentSessionCount,omitempty"`
// 				ControllerDNSName            string     `json:"ControllerDnsName,omitempty"`
// 				PoweredOnDate                *time.Time `json:"PoweredOnDate,omitempty"`
// 				PowerStateChangeDate         *time.Time `json:"PowerStateChangeDate,omitempty"`
// 				FunctionalLevel              int        `json:"FunctionalLevel,omitempty"`
// 				FailureDate                  *time.Time `json:"FailureDate,omitempty"`
// 				WindowsConnectionSetting     int        `json:"WindowsConnectionSetting,omitempty"`
// 				IsPreparing                  bool       `json:"IsPreparing,omitempty"`
// 				FaultState                   int        `json:"FaultState,omitempty"`
// 				OSType                       string     `json:"OSType,omitempty"`
// 				CurrentLoadIndexID           int        `json:"CurrentLoadIndexId,omitempty"`
// 				CatalogID                    string     `json:"CatalogId,omitempty"`
// 				DesktopGroupID               string     `json:"DesktopGroupId,omitempty"`
// 				HypervisorID                 string     `json:"HypervisorId,omitempty"`
// 				LastPowerActionCompletedDate *time.Time `json:"LastPowerActionCompletedDate,omitempty"`
// 				LastUpgradeState             int        `json:"LastUpgradeState,omitempty"`
// 				LastUpgradeStateChangeDate   *time.Time `json:"LastUpgradeStateChangeDate,omitempty"`
// 				Hash                         string     `json:"Hash,omitempty"`
// 				MachineRole                  int        `json:"MachineRole,omitempty"`
// 				CreatedDate                  *time.Time `json:"CreatedDate,omitempty"`
// 				ModifiedDate                 *time.Time `json:"ModifiedDate,omitempty"`
// 				CurrentLoadIndex             struct {
// 					ID                 int        `json:"Id,omitempty"`
// 					EffectiveLoadIndex int        `json:"EffectiveLoadIndex,omitempty"`
// 					CPU                int        `json:"Cpu,omitempty"`
// 					Memory             int        `json:"Memory,omitempty"`
// 					Disk               int        `json:"Disk,omitempty"`
// 					Network            int        `json:"Network,omitempty"`
// 					SessionCount       int        `json:"SessionCount,omitempty"`
// 					MachineID          string     `json:"MachineId,omitempty"`
// 					CreatedDate        *time.Time `json:"CreatedDate,omitempty"`
// 					ModifiedDate       *time.Time `json:"ModifiedDate,omitempty"`
// 				} `json:"CurrentLoadIndex,omitempty"`
// 				Catalog struct {
// 					ID                    string     `json:"Id,omitempty"`
// 					Name                  string     `json:"Name,omitempty"`
// 					LifecycleState        int        `json:"LifecycleState,omitempty"`
// 					ProvisioningType      int        `json:"ProvisioningType,omitempty"`
// 					PersistentUserChanges int        `json:"PersistentUserChanges,omitempty"`
// 					IsMachinePhysical     bool       `json:"IsMachinePhysical,omitempty"`
// 					AllocationType        int        `json:"AllocationType,omitempty"`
// 					SessionSupport        int        `json:"SessionSupport,omitempty"`
// 					ProvisioningSchemeID  string     `json:"ProvisioningSchemeId,omitempty"`
// 					ZoneUID               string     `json:"ZoneUid,omitempty"`
// 					ZoneName              string     `json:"ZoneName,omitempty"`
// 					CreatedDate           *time.Time `json:"CreatedDate,omitempty"`
// 					ModifiedDate          *time.Time `json:"ModifiedDate,omitempty"`
// 				} `json:"Catalog,omitempty"`
// 				DesktopGroup struct {
// 					ID                  string     `json:"Id,omitempty"`
// 					Name                string     `json:"Name,omitempty"`
// 					IsRemotePC          bool       `json:"IsRemotePC,omitempty"`
// 					DesktopKind         int        `json:"DesktopKind,omitempty"`
// 					LifecycleState      int        `json:"LifecycleState,omitempty"`
// 					SessionSupport      int        `json:"SessionSupport,omitempty"`
// 					DeliveryType        int        `json:"DeliveryType,omitempty"`
// 					IsInMaintenanceMode bool       `json:"IsInMaintenanceMode,omitempty"`
// 					MachineCost         int        `json:"MachineCost,omitempty"`
// 					AutoscaleTagID      int        `json:"AutoscaleTagId,omitempty"`
// 					CreatedDate         *time.Time `json:"CreatedDate,omitempty"`
// 					ModifiedDate        *time.Time `json:"ModifiedDate,omitempty"`
// 				} `json:"DesktopGroup,omitempty"`
// 				Hypervisor struct {
// 					ID             string     `json:"Id,omitempty"`
// 					Name           string     `json:"Name,omitempty"`
// 					Type           string     `json:"Type,omitempty"`
// 					LifecycleState int        `json:"LifecycleState,omitempty"`
// 					CreatedDate    *time.Time `json:"CreatedDate,omitempty"`
// 					ModifiedDate   *time.Time `json:"ModifiedDate,omitempty"`
// 				} `json:"Hypervisor,omitempty"`
// 				MachineCost struct {
// 					MachineID                  string     `json:"MachineId,omitempty"`
// 					SpecID                     int        `json:"SpecId,omitempty"`
// 					CostPerHour                int        `json:"CostPerHour,omitempty"`
// 					PowerOnComputeCostPerHour  int        `json:"PowerOnComputeCostPerHour,omitempty"`
// 					PowerOnStorageCostPerHour  int        `json:"PowerOnStorageCostPerHour,omitempty"`
// 					PowerOffStorageCostPerHour int        `json:"PowerOffStorageCostPerHour,omitempty"`
// 					CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
// 					ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
// 				} `json:"MachineCost,omitempty"`
// 			} `json:"Machine,omitempty"`
// 			SessionRecordingServer struct {
// 				SessionKey                 string     `json:"SessionKey,omitempty"`
// 				SessionRecordingServerName string     `json:"SessionRecordingServerName,omitempty"`
// 				CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
// 			} `json:"SessionRecordingServer,omitempty"`
// 			PublishedDesktopName struct {
// 				ID            int    `json:"Id,omitempty"`
// 				PublishedName string `json:"PublishedName,omitempty"`
// 				Sessions      []any  `json:"Sessions,omitempty"`
// 			} `json:"PublishedDesktopName,omitempty"`
// 			SessionMetricsLatest struct {
// 				SessionKey         string     `json:"SessionKey,omitempty"`
// 				ClientL7Latency    float64    `json:"ClientL7Latency,omitempty"`
// 				CloudConnectorName string     `json:"CloudConnectorName,omitempty"`
// 				CreatedDate        *time.Time `json:"CreatedDate,omitempty"`
// 				EdtMtu             string     `json:"EdtMtu,omitempty"`
// 				GatewayPopName     string     `json:"GatewayPopName,omitempty"`
// 				HdxConnectionType  int        `json:"HdxConnectionType,omitempty"`
// 				HdxProtocolName    string     `json:"HdxProtocolName,omitempty"`
// 				ModifiedDate       *time.Time `json:"ModifiedDate,omitempty"`
// 				ServerL7Latency    float64    `json:"ServerL7Latency,omitempty"`
// 			} `json:"SessionMetricsLatest,omitempty"`
// 		} `json:"Session,omitempty"`
// 	} `json:"value,omitempty"`
// }

// AutoGenerated represents the structure of the JSON response from the Citrix CMS API for health data.
type MachineMetricDetails_JSON struct {
	OdataContext string `json:"@odata.context,omitempty"`
	OdataCount   int    `json:"@odata.count,omitempty"`
	Message      string
	Value        []struct {
		MachineID     string     `json:"MachineId,omitempty"`
		CollectedDate *time.Time `json:"CollectedDate,omitempty"`
		Iops          int        `json:"Iops,omitempty"`
		Latency       float64    `json:"Latency,omitempty"`
		Machine       struct {
			ID                           string     `json:"Id,omitempty"`
			Sid                          string     `json:"Sid,omitempty"`
			Name                         string     `json:"Name,omitempty"`
			DNSName                      string     `json:"DnsName,omitempty"`
			LifecycleState               int        `json:"LifecycleState,omitempty"`
			IPAddress                    string     `json:"IPAddress,omitempty"`
			HostedMachineID              string     `json:"HostedMachineId,omitempty"`
			HostingServerName            string     `json:"HostingServerName,omitempty"`
			HostedMachineName            string     `json:"HostedMachineName,omitempty"`
			IsAssigned                   bool       `json:"IsAssigned,omitempty"`
			IsInMaintenanceMode          bool       `json:"IsInMaintenanceMode,omitempty"`
			IsPendingUpdate              bool       `json:"IsPendingUpdate,omitempty"`
			AgentVersion                 string     `json:"AgentVersion,omitempty"`
			AssociatedUserFullNames      string     `json:"AssociatedUserFullNames,omitempty"`
			AssociatedUserNames          string     `json:"AssociatedUserNames,omitempty"`
			AssociatedUserUPNs           string     `json:"AssociatedUserUPNs,omitempty"`
			CurrentRegistrationState     int        `json:"CurrentRegistrationState,omitempty"`
			RegistrationStateChangeDate  *time.Time `json:"RegistrationStateChangeDate,omitempty"`
			LastDeregisteredCode         int        `json:"LastDeregisteredCode,omitempty"`
			LastDeregisteredDate         *time.Time `json:"LastDeregisteredDate,omitempty"`
			CurrentPowerState            int        `json:"CurrentPowerState,omitempty"`
			CurrentSessionCount          int        `json:"CurrentSessionCount,omitempty"`
			ControllerDNSName            string     `json:"ControllerDnsName,omitempty"`
			PoweredOnDate                *time.Time `json:"PoweredOnDate,omitempty"`
			PowerStateChangeDate         *time.Time `json:"PowerStateChangeDate,omitempty"`
			FunctionalLevel              int        `json:"FunctionalLevel,omitempty"`
			FailureDate                  *time.Time `json:"FailureDate,omitempty"`
			WindowsConnectionSetting     int        `json:"WindowsConnectionSetting,omitempty"`
			IsPreparing                  bool       `json:"IsPreparing,omitempty"`
			FaultState                   int        `json:"FaultState,omitempty"`
			OSType                       string     `json:"OSType,omitempty"`
			CurrentLoadIndexID           int        `json:"CurrentLoadIndexId,omitempty"`
			CatalogID                    string     `json:"CatalogId,omitempty"`
			DesktopGroupID               string     `json:"DesktopGroupId,omitempty"`
			HypervisorID                 string     `json:"HypervisorId,omitempty"`
			LastPowerActionCompletedDate *time.Time `json:"LastPowerActionCompletedDate,omitempty"`
			LastUpgradeState             int        `json:"LastUpgradeState,omitempty"`
			LastUpgradeStateChangeDate   *time.Time `json:"LastUpgradeStateChangeDate,omitempty"`
			Hash                         string     `json:"Hash,omitempty"`
			MachineRole                  int        `json:"MachineRole,omitempty"`
			CreatedDate                  *time.Time `json:"CreatedDate,omitempty"`
			ModifiedDate                 *time.Time `json:"ModifiedDate,omitempty"`
			CurrentLoadIndex             struct {
				ID                 int        `json:"Id,omitempty"`
				EffectiveLoadIndex int        `json:"EffectiveLoadIndex,omitempty"`
				CPU                int        `json:"Cpu,omitempty"`
				Memory             int        `json:"Memory,omitempty"`
				Disk               int        `json:"Disk,omitempty"`
				Network            int        `json:"Network,omitempty"`
				SessionCount       int        `json:"SessionCount,omitempty"`
				MachineID          string     `json:"MachineId,omitempty"`
				CreatedDate        *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate       *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"CurrentLoadIndex,omitempty"`
			Catalog struct {
				ID                    string     `json:"Id,omitempty"`
				Name                  string     `json:"Name,omitempty"`
				LifecycleState        int        `json:"LifecycleState,omitempty"`
				ProvisioningType      int        `json:"ProvisioningType,omitempty"`
				PersistentUserChanges int        `json:"PersistentUserChanges,omitempty"`
				IsMachinePhysical     bool       `json:"IsMachinePhysical,omitempty"`
				AllocationType        int        `json:"AllocationType,omitempty"`
				SessionSupport        int        `json:"SessionSupport,omitempty"`
				ProvisioningSchemeID  string     `json:"ProvisioningSchemeId,omitempty"`
				ZoneUID               string     `json:"ZoneUid,omitempty"`
				ZoneName              string     `json:"ZoneName,omitempty"`
				CreatedDate           *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate          *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"Catalog,omitempty"`
			DesktopGroup struct {
				ID                  string     `json:"Id,omitempty"`
				Name                string     `json:"Name,omitempty"`
				IsRemotePC          bool       `json:"IsRemotePC,omitempty"`
				DesktopKind         int        `json:"DesktopKind,omitempty"`
				LifecycleState      int        `json:"LifecycleState,omitempty"`
				SessionSupport      int        `json:"SessionSupport,omitempty"`
				DeliveryType        int        `json:"DeliveryType,omitempty"`
				IsInMaintenanceMode bool       `json:"IsInMaintenanceMode,omitempty"`
				MachineCost         float64    `json:"MachineCost,omitempty"`
				AutoscaleTagID      int        `json:"AutoscaleTagId,omitempty"`
				CreatedDate         *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate        *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"DesktopGroup,omitempty"`
			Hypervisor struct {
				ID             string     `json:"Id,omitempty"`
				Name           string     `json:"Name,omitempty"`
				Type           string     `json:"Type,omitempty"`
				LifecycleState int        `json:"LifecycleState,omitempty"`
				CreatedDate    *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate   *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"Hypervisor,omitempty"`
			MachineCost struct {
				MachineID                  string     `json:"MachineId,omitempty"`
				SpecID                     int        `json:"SpecId,omitempty"`
				CostPerHour                int        `json:"CostPerHour,omitempty"`
				PowerOnComputeCostPerHour  int        `json:"PowerOnComputeCostPerHour,omitempty"`
				PowerOnStorageCostPerHour  int        `json:"PowerOnStorageCostPerHour,omitempty"`
				PowerOffStorageCostPerHour int        `json:"PowerOffStorageCostPerHour,omitempty"`
				CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"MachineCost,omitempty"`
		} `json:"Machine,omitempty"`
	} `json:"value,omitempty"`
}

// AutoGenerated represents the structure of the JSON response from the Citrix CMS API for health data.
type SessionMetricsDetails_JSON struct {
	OdataContext string `json:"@odata.context,omitempty"`
	OdataCount   int    `json:"@odata.count,omitempty"`
	Message      string
	Value        []struct {
		ID              int        `json:"Id,omitempty"`
		CollectedDate   *time.Time `json:"CollectedDate,omitempty"`
		IcaRttMS        float64    `json:"IcaRttMS,omitempty"`
		IcaLatency      float64    `json:"IcaLatency,omitempty"`
		ClientL7Latency float64    `json:"ClientL7Latency,omitempty"`
		ServerL7Latency float64    `json:"ServerL7Latency,omitempty"`
		SessionID       string     `json:"SessionId,omitempty"`
		CreatedDate     *time.Time `json:"CreatedDate,omitempty"`
		ModifiedDate    *time.Time `json:"ModifiedDate,omitempty"`
		Session         struct {
			SessionKey                string     `json:"SessionKey,omitempty"`
			StartDate                 *time.Time `json:"StartDate,omitempty"`
			LogOnDuration             int        `json:"LogOnDuration,omitempty"`
			EndDate                   *time.Time `json:"EndDate,omitempty"`
			ExitCode                  int        `json:"ExitCode,omitempty"`
			FailureID                 int        `json:"FailureId,omitempty"`
			FailureDate               *time.Time `json:"FailureDate,omitempty"`
			ConnectionState           int        `json:"ConnectionState,omitempty"`
			SessionIdleTime           *time.Time `json:"SessionIdleTime,omitempty"`
			ConnectionStateChangeDate *time.Time `json:"ConnectionStateChangeDate,omitempty"`
			LifecycleState            int        `json:"LifecycleState,omitempty"`
			CurrentConnectionID       int        `json:"CurrentConnectionId,omitempty"`
			UserID                    int        `json:"UserId,omitempty"`
			MachineID                 string     `json:"MachineId,omitempty"`
			SessionType               int        `json:"SessionType,omitempty"`
			IsAnonymous               bool       `json:"IsAnonymous,omitempty"`
			PublishedDesktopID        int        `json:"PublishedDesktopId,omitempty"`
			CreatedDate               *time.Time `json:"CreatedDate,omitempty"`
			ModifiedDate              *time.Time `json:"ModifiedDate,omitempty"`
			Failure                   struct {
				ID                         int        `json:"Id,omitempty"`
				ConnectionFailureEnumValue int        `json:"ConnectionFailureEnumValue,omitempty"`
				Category                   int        `json:"Category,omitempty"`
				CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"Failure,omitempty"`
			CurrentConnection struct {
				ID                        int        `json:"Id,omitempty"`
				ClientName                string     `json:"ClientName,omitempty"`
				ClientAddress             string     `json:"ClientAddress,omitempty"`
				ClientPublicIP            string     `json:"ClientPublicIP,omitempty"`
				ClientVersion             string     `json:"ClientVersion,omitempty"`
				ClientPlatform            string     `json:"ClientPlatform,omitempty"`
				ClientISP                 string     `json:"ClientISP,omitempty"`
				ClientLocationCountry     string     `json:"ClientLocationCountry,omitempty"`
				ClientLocationCity        string     `json:"ClientLocationCity,omitempty"`
				ConnectedViaHostName      string     `json:"ConnectedViaHostName,omitempty"`
				ConnectedViaIPAddress     string     `json:"ConnectedViaIPAddress,omitempty"`
				LaunchedViaHostName       string     `json:"LaunchedViaHostName,omitempty"`
				LaunchedViaIPAddress      string     `json:"LaunchedViaIPAddress,omitempty"`
				IsReconnect               bool       `json:"IsReconnect,omitempty"`
				IsSecureIca               bool       `json:"IsSecureIca,omitempty"`
				Protocol                  string     `json:"Protocol,omitempty"`
				LogOnStartDate            *time.Time `json:"LogOnStartDate,omitempty"`
				LogOnEndDate              *time.Time `json:"LogOnEndDate,omitempty"`
				BrokeringDuration         int        `json:"BrokeringDuration,omitempty"`
				BrokeringDate             *time.Time `json:"BrokeringDate,omitempty"`
				DisconnectCode            int        `json:"DisconnectCode,omitempty"`
				DisconnectDate            *time.Time `json:"DisconnectDate,omitempty"`
				VMStartStartDate          *time.Time `json:"VMStartStartDate,omitempty"`
				VMPoweredOnDate           *time.Time `json:"VMPoweredOnDate,omitempty"`
				VMStartEndDate            *time.Time `json:"VMStartEndDate,omitempty"`
				ClientSessionValidateDate *time.Time `json:"ClientSessionValidateDate,omitempty"`
				ServerSessionValidateDate *time.Time `json:"ServerSessionValidateDate,omitempty"`
				EstablishmentDate         *time.Time `json:"EstablishmentDate,omitempty"`
				HdxStartDate              *time.Time `json:"HdxStartDate,omitempty"`
				HdxEndDate                *time.Time `json:"HdxEndDate,omitempty"`
				AuthenticationDuration    int        `json:"AuthenticationDuration,omitempty"`
				GpoStartDate              *time.Time `json:"GpoStartDate,omitempty"`
				GpoEndDate                *time.Time `json:"GpoEndDate,omitempty"`
				LogOnScriptsStartDate     *time.Time `json:"LogOnScriptsStartDate,omitempty"`
				LogOnScriptsEndDate       *time.Time `json:"LogOnScriptsEndDate,omitempty"`
				ProfileLoadStartDate      *time.Time `json:"ProfileLoadStartDate,omitempty"`
				ProfileLoadEndDate        *time.Time `json:"ProfileLoadEndDate,omitempty"`
				InteractiveStartDate      *time.Time `json:"InteractiveStartDate,omitempty"`
				InteractiveEndDate        *time.Time `json:"InteractiveEndDate,omitempty"`
				SessionKey                string     `json:"SessionKey,omitempty"`
				CreatedDate               *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate              *time.Time `json:"ModifiedDate,omitempty"`
				ConnectionFailureLog      struct {
					ID                         int        `json:"Id,omitempty"`
					SessionKey                 string     `json:"SessionKey,omitempty"`
					FailureDate                *time.Time `json:"FailureDate,omitempty"`
					UserID                     int        `json:"UserId,omitempty"`
					MachineID                  string     `json:"MachineId,omitempty"`
					ConnectionFailureEnumValue int        `json:"ConnectionFailureEnumValue,omitempty"`
					CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
					ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
				} `json:"ConnectionFailureLog,omitempty"`
			} `json:"CurrentConnection,omitempty"`
			User struct {
				ID           int        `json:"Id,omitempty"`
				Sid          string     `json:"Sid,omitempty"`
				Upn          string     `json:"Upn,omitempty"`
				UserName     string     `json:"UserName,omitempty"`
				FullName     string     `json:"FullName,omitempty"`
				Domain       string     `json:"Domain,omitempty"`
				CreatedDate  *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"User,omitempty"`
			Machine struct {
				ID                           string     `json:"Id,omitempty"`
				Sid                          string     `json:"Sid,omitempty"`
				Name                         string     `json:"Name,omitempty"`
				DNSName                      string     `json:"DnsName,omitempty"`
				LifecycleState               int        `json:"LifecycleState,omitempty"`
				IPAddress                    string     `json:"IPAddress,omitempty"`
				HostedMachineID              string     `json:"HostedMachineId,omitempty"`
				HostingServerName            string     `json:"HostingServerName,omitempty"`
				HostedMachineName            string     `json:"HostedMachineName,omitempty"`
				IsAssigned                   bool       `json:"IsAssigned,omitempty"`
				IsInMaintenanceMode          bool       `json:"IsInMaintenanceMode,omitempty"`
				IsPendingUpdate              bool       `json:"IsPendingUpdate,omitempty"`
				AgentVersion                 string     `json:"AgentVersion,omitempty"`
				AssociatedUserFullNames      string     `json:"AssociatedUserFullNames,omitempty"`
				AssociatedUserNames          string     `json:"AssociatedUserNames,omitempty"`
				AssociatedUserUPNs           string     `json:"AssociatedUserUPNs,omitempty"`
				CurrentRegistrationState     int        `json:"CurrentRegistrationState,omitempty"`
				RegistrationStateChangeDate  *time.Time `json:"RegistrationStateChangeDate,omitempty"`
				LastDeregisteredCode         int        `json:"LastDeregisteredCode,omitempty"`
				LastDeregisteredDate         *time.Time `json:"LastDeregisteredDate,omitempty"`
				CurrentPowerState            int        `json:"CurrentPowerState,omitempty"`
				CurrentSessionCount          int        `json:"CurrentSessionCount,omitempty"`
				ControllerDNSName            string     `json:"ControllerDnsName,omitempty"`
				PoweredOnDate                *time.Time `json:"PoweredOnDate,omitempty"`
				PowerStateChangeDate         *time.Time `json:"PowerStateChangeDate,omitempty"`
				FunctionalLevel              int        `json:"FunctionalLevel,omitempty"`
				FailureDate                  *time.Time `json:"FailureDate,omitempty"`
				WindowsConnectionSetting     int        `json:"WindowsConnectionSetting,omitempty"`
				IsPreparing                  bool       `json:"IsPreparing,omitempty"`
				FaultState                   int        `json:"FaultState,omitempty"`
				OSType                       string     `json:"OSType,omitempty"`
				CurrentLoadIndexID           int        `json:"CurrentLoadIndexId,omitempty"`
				CatalogID                    string     `json:"CatalogId,omitempty"`
				DesktopGroupID               string     `json:"DesktopGroupId,omitempty"`
				HypervisorID                 string     `json:"HypervisorId,omitempty"`
				LastPowerActionCompletedDate *time.Time `json:"LastPowerActionCompletedDate,omitempty"`
				LastUpgradeState             int        `json:"LastUpgradeState,omitempty"`
				LastUpgradeStateChangeDate   *time.Time `json:"LastUpgradeStateChangeDate,omitempty"`
				Hash                         string     `json:"Hash,omitempty"`
				MachineRole                  int        `json:"MachineRole,omitempty"`
				CreatedDate                  *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate                 *time.Time `json:"ModifiedDate,omitempty"`
				CurrentLoadIndex             struct {
					ID                 int        `json:"Id,omitempty"`
					EffectiveLoadIndex int        `json:"EffectiveLoadIndex,omitempty"`
					CPU                int        `json:"Cpu,omitempty"`
					Memory             int        `json:"Memory,omitempty"`
					Disk               int        `json:"Disk,omitempty"`
					Network            int        `json:"Network,omitempty"`
					SessionCount       int        `json:"SessionCount,omitempty"`
					MachineID          string     `json:"MachineId,omitempty"`
					CreatedDate        *time.Time `json:"CreatedDate,omitempty"`
					ModifiedDate       *time.Time `json:"ModifiedDate,omitempty"`
				} `json:"CurrentLoadIndex,omitempty"`
				Catalog struct {
					ID                    string     `json:"Id,omitempty"`
					Name                  string     `json:"Name,omitempty"`
					LifecycleState        int        `json:"LifecycleState,omitempty"`
					ProvisioningType      int        `json:"ProvisioningType,omitempty"`
					PersistentUserChanges int        `json:"PersistentUserChanges,omitempty"`
					IsMachinePhysical     bool       `json:"IsMachinePhysical,omitempty"`
					AllocationType        int        `json:"AllocationType,omitempty"`
					SessionSupport        int        `json:"SessionSupport,omitempty"`
					ProvisioningSchemeID  string     `json:"ProvisioningSchemeId,omitempty"`
					ZoneUID               string     `json:"ZoneUid,omitempty"`
					ZoneName              string     `json:"ZoneName,omitempty"`
					CreatedDate           *time.Time `json:"CreatedDate,omitempty"`
					ModifiedDate          *time.Time `json:"ModifiedDate,omitempty"`
				} `json:"Catalog,omitempty"`
				DesktopGroup struct {
					ID                  string     `json:"Id,omitempty"`
					Name                string     `json:"Name,omitempty"`
					IsRemotePC          bool       `json:"IsRemotePC,omitempty"`
					DesktopKind         int        `json:"DesktopKind,omitempty"`
					LifecycleState      int        `json:"LifecycleState,omitempty"`
					SessionSupport      int        `json:"SessionSupport,omitempty"`
					DeliveryType        int        `json:"DeliveryType,omitempty"`
					IsInMaintenanceMode bool       `json:"IsInMaintenanceMode,omitempty"`
					MachineCost         float64    `json:"MachineCost,omitempty"`
					AutoscaleTagID      int        `json:"AutoscaleTagId,omitempty"`
					CreatedDate         *time.Time `json:"CreatedDate,omitempty"`
					ModifiedDate        *time.Time `json:"ModifiedDate,omitempty"`
				} `json:"DesktopGroup,omitempty"`
				Hypervisor struct {
					ID             string     `json:"Id,omitempty"`
					Name           string     `json:"Name,omitempty"`
					Type           string     `json:"Type,omitempty"`
					LifecycleState int        `json:"LifecycleState,omitempty"`
					CreatedDate    *time.Time `json:"CreatedDate,omitempty"`
					ModifiedDate   *time.Time `json:"ModifiedDate,omitempty"`
				} `json:"Hypervisor,omitempty"`
				MachineCost struct {
					MachineID                  string     `json:"MachineId,omitempty"`
					SpecID                     int        `json:"SpecId,omitempty"`
					CostPerHour                int        `json:"CostPerHour,omitempty"`
					PowerOnComputeCostPerHour  int        `json:"PowerOnComputeCostPerHour,omitempty"`
					PowerOnStorageCostPerHour  int        `json:"PowerOnStorageCostPerHour,omitempty"`
					PowerOffStorageCostPerHour int        `json:"PowerOffStorageCostPerHour,omitempty"`
					CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
					ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
				} `json:"MachineCost,omitempty"`
			} `json:"Machine,omitempty"`
			SessionRecordingServer struct {
				SessionKey                 string     `json:"SessionKey,omitempty"`
				SessionRecordingServerName string     `json:"SessionRecordingServerName,omitempty"`
				CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
			} `json:"SessionRecordingServer,omitempty"`
			PublishedDesktopName struct {
				ID            int    `json:"Id,omitempty"`
				PublishedName string `json:"PublishedName,omitempty"`
			} `json:"PublishedDesktopName,omitempty"`
			SessionMetricsLatest struct {
				SessionKey         string     `json:"SessionKey,omitempty"`
				ClientL7Latency    float64    `json:"ClientL7Latency,omitempty"`
				CloudConnectorName string     `json:"CloudConnectorName,omitempty"`
				CreatedDate        *time.Time `json:"CreatedDate,omitempty"`
				EdtMtu             string     `json:"EdtMtu,omitempty"`
				GatewayPopName     string     `json:"GatewayPopName,omitempty"`
				HdxConnectionType  int        `json:"HdxConnectionType,omitempty"`
				HdxProtocolName    string     `json:"HdxProtocolName,omitempty"`
				ModifiedDate       *time.Time `json:"ModifiedDate,omitempty"`
				ServerL7Latency    float64    `json:"ServerL7Latency,omitempty"`
			} `json:"SessionMetricsLatest,omitempty"`
		} `json:"Session,omitempty"`
	} `json:"value,omitempty"`
}

// AutoGenerated represents the structure of the JSON response from the Citrix CMS API for session details.
type SessionsDetails_JSON struct {
	OdataContext string `json:"@odata.context,omitempty"`
	OdataCount   int    `json:"@odata.count,omitempty"`
	Message      string
	Value        []struct {
		SessionKey                string     `json:"SessionKey,omitempty"`
		StartDate                 *time.Time `json:"StartDate,omitempty"`
		LogOnDuration             int        `json:"LogOnDuration,omitempty"`
		EndDate                   *time.Time `json:"EndDate,omitempty"`
		ExitCode                  int        `json:"ExitCode,omitempty"`
		FailureID                 int        `json:"FailureId,omitempty"`
		FailureDate               *time.Time `json:"FailureDate,omitempty"`
		ConnectionState           int        `json:"ConnectionState,omitempty"`
		SessionIdleTime           *time.Time `json:"SessionIdleTime,omitempty"`
		ConnectionStateChangeDate *time.Time `json:"ConnectionStateChangeDate,omitempty"`
		LifecycleState            int        `json:"LifecycleState,omitempty"`
		CurrentConnectionID       int        `json:"CurrentConnectionId,omitempty"`
		UserID                    int        `json:"UserId,omitempty"`
		MachineID                 string     `json:"MachineId,omitempty"`
		SessionType               int        `json:"SessionType,omitempty"`
		IsAnonymous               bool       `json:"IsAnonymous,omitempty"`
		PublishedDesktopID        int        `json:"PublishedDesktopId,omitempty"`
		CreatedDate               *time.Time `json:"CreatedDate,omitempty"`
		ModifiedDate              *time.Time `json:"ModifiedDate,omitempty"`
		Failure                   struct {
			ID                         int        `json:"Id,omitempty"`
			ConnectionFailureEnumValue int        `json:"ConnectionFailureEnumValue,omitempty"`
			Category                   int        `json:"Category,omitempty"`
			CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
			ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
		} `json:"Failure,omitempty"`
		CurrentConnection struct {
			ID                        int        `json:"Id,omitempty"`
			ClientName                string     `json:"ClientName,omitempty"`
			ClientAddress             string     `json:"ClientAddress,omitempty"`
			ClientPublicIP            string     `json:"ClientPublicIP,omitempty"`
			ClientVersion             string     `json:"ClientVersion,omitempty"`
			ClientPlatform            string     `json:"ClientPlatform,omitempty"`
			ClientISP                 string     `json:"ClientISP,omitempty"`
			ClientLocationCountry     string     `json:"ClientLocationCountry,omitempty"`
			ClientLocationCity        string     `json:"ClientLocationCity,omitempty"`
			ConnectedViaHostName      string     `json:"ConnectedViaHostName,omitempty"`
			ConnectedViaIPAddress     string     `json:"ConnectedViaIPAddress,omitempty"`
			LaunchedViaHostName       string     `json:"LaunchedViaHostName,omitempty"`
			LaunchedViaIPAddress      string     `json:"LaunchedViaIPAddress,omitempty"`
			IsReconnect               bool       `json:"IsReconnect,omitempty"`
			IsSecureIca               bool       `json:"IsSecureIca,omitempty"`
			Protocol                  string     `json:"Protocol,omitempty"`
			LogOnStartDate            *time.Time `json:"LogOnStartDate,omitempty"`
			LogOnEndDate              *time.Time `json:"LogOnEndDate,omitempty"`
			BrokeringDuration         int        `json:"BrokeringDuration,omitempty"`
			BrokeringDate             *time.Time `json:"BrokeringDate,omitempty"`
			DisconnectCode            int        `json:"DisconnectCode,omitempty"`
			DisconnectDate            *time.Time `json:"DisconnectDate,omitempty"`
			VMStartStartDate          *time.Time `json:"VMStartStartDate,omitempty"`
			VMPoweredOnDate           *time.Time `json:"VMPoweredOnDate,omitempty"`
			VMStartEndDate            *time.Time `json:"VMStartEndDate,omitempty"`
			ClientSessionValidateDate *time.Time `json:"ClientSessionValidateDate,omitempty"`
			ServerSessionValidateDate *time.Time `json:"ServerSessionValidateDate,omitempty"`
			EstablishmentDate         *time.Time `json:"EstablishmentDate,omitempty"`
			HdxStartDate              *time.Time `json:"HdxStartDate,omitempty"`
			HdxEndDate                *time.Time `json:"HdxEndDate,omitempty"`
			AuthenticationDuration    int        `json:"AuthenticationDuration,omitempty"`
			GpoStartDate              *time.Time `json:"GpoStartDate,omitempty"`
			GpoEndDate                *time.Time `json:"GpoEndDate,omitempty"`
			LogOnScriptsStartDate     *time.Time `json:"LogOnScriptsStartDate,omitempty"`
			LogOnScriptsEndDate       *time.Time `json:"LogOnScriptsEndDate,omitempty"`
			ProfileLoadStartDate      *time.Time `json:"ProfileLoadStartDate,omitempty"`
			ProfileLoadEndDate        *time.Time `json:"ProfileLoadEndDate,omitempty"`
			InteractiveStartDate      *time.Time `json:"InteractiveStartDate,omitempty"`
			InteractiveEndDate        *time.Time `json:"InteractiveEndDate,omitempty"`
			SessionKey                string     `json:"SessionKey,omitempty"`
			CreatedDate               *time.Time `json:"CreatedDate,omitempty"`
			ModifiedDate              *time.Time `json:"ModifiedDate,omitempty"`
			ConnectionFailureLog      struct {
				ID                         int        `json:"Id,omitempty"`
				SessionKey                 string     `json:"SessionKey,omitempty"`
				FailureDate                *time.Time `json:"FailureDate,omitempty"`
				UserID                     int        `json:"UserId,omitempty"`
				MachineID                  string     `json:"MachineId,omitempty"`
				ConnectionFailureEnumValue int        `json:"ConnectionFailureEnumValue,omitempty"`
				CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"ConnectionFailureLog,omitempty"`
		} `json:"CurrentConnection,omitempty"`
		User struct {
			ID           int        `json:"Id,omitempty"`
			Sid          string     `json:"Sid,omitempty"`
			Upn          string     `json:"Upn,omitempty"`
			UserName     string     `json:"UserName,omitempty"`
			FullName     string     `json:"FullName,omitempty"`
			Domain       string     `json:"Domain,omitempty"`
			CreatedDate  *time.Time `json:"CreatedDate,omitempty"`
			ModifiedDate *time.Time `json:"ModifiedDate,omitempty"`
		} `json:"User,omitempty"`
		Machine struct {
			ID                           string     `json:"Id,omitempty"`
			Sid                          string     `json:"Sid,omitempty"`
			Name                         string     `json:"Name,omitempty"`
			DNSName                      string     `json:"DnsName,omitempty"`
			LifecycleState               int        `json:"LifecycleState,omitempty"`
			IPAddress                    string     `json:"IPAddress,omitempty"`
			HostedMachineID              string     `json:"HostedMachineId,omitempty"`
			HostingServerName            string     `json:"HostingServerName,omitempty"`
			HostedMachineName            string     `json:"HostedMachineName,omitempty"`
			IsAssigned                   bool       `json:"IsAssigned,omitempty"`
			IsInMaintenanceMode          bool       `json:"IsInMaintenanceMode,omitempty"`
			IsPendingUpdate              bool       `json:"IsPendingUpdate,omitempty"`
			AgentVersion                 string     `json:"AgentVersion,omitempty"`
			AssociatedUserFullNames      string     `json:"AssociatedUserFullNames,omitempty"`
			AssociatedUserNames          string     `json:"AssociatedUserNames,omitempty"`
			AssociatedUserUPNs           string     `json:"AssociatedUserUPNs,omitempty"`
			CurrentRegistrationState     int        `json:"CurrentRegistrationState,omitempty"`
			RegistrationStateChangeDate  *time.Time `json:"RegistrationStateChangeDate,omitempty"`
			LastDeregisteredCode         int        `json:"LastDeregisteredCode,omitempty"`
			LastDeregisteredDate         *time.Time `json:"LastDeregisteredDate,omitempty"`
			CurrentPowerState            int        `json:"CurrentPowerState,omitempty"`
			CurrentSessionCount          int        `json:"CurrentSessionCount,omitempty"`
			ControllerDNSName            string     `json:"ControllerDnsName,omitempty"`
			PoweredOnDate                *time.Time `json:"PoweredOnDate,omitempty"`
			PowerStateChangeDate         *time.Time `json:"PowerStateChangeDate,omitempty"`
			FunctionalLevel              int        `json:"FunctionalLevel,omitempty"`
			FailureDate                  *time.Time `json:"FailureDate,omitempty"`
			WindowsConnectionSetting     int        `json:"WindowsConnectionSetting,omitempty"`
			IsPreparing                  bool       `json:"IsPreparing,omitempty"`
			FaultState                   int        `json:"FaultState,omitempty"`
			OSType                       string     `json:"OSType,omitempty"`
			CurrentLoadIndexID           int        `json:"CurrentLoadIndexId,omitempty"`
			CatalogID                    string     `json:"CatalogId,omitempty"`
			DesktopGroupID               string     `json:"DesktopGroupId,omitempty"`
			HypervisorID                 string     `json:"HypervisorId,omitempty"`
			LastPowerActionCompletedDate *time.Time `json:"LastPowerActionCompletedDate,omitempty"`
			LastUpgradeState             int        `json:"LastUpgradeState,omitempty"`
			LastUpgradeStateChangeDate   *time.Time `json:"LastUpgradeStateChangeDate,omitempty"`
			Hash                         string     `json:"Hash,omitempty"`
			MachineRole                  int        `json:"MachineRole,omitempty"`
			CreatedDate                  *time.Time `json:"CreatedDate,omitempty"`
			ModifiedDate                 *time.Time `json:"ModifiedDate,omitempty"`
			CurrentLoadIndex             struct {
				ID                 int        `json:"Id,omitempty"`
				EffectiveLoadIndex int        `json:"EffectiveLoadIndex,omitempty"`
				CPU                int        `json:"Cpu,omitempty"`
				Memory             int        `json:"Memory,omitempty"`
				Disk               int        `json:"Disk,omitempty"`
				Network            int        `json:"Network,omitempty"`
				SessionCount       int        `json:"SessionCount,omitempty"`
				MachineID          string     `json:"MachineId,omitempty"`
				CreatedDate        *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate       *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"CurrentLoadIndex,omitempty"`
			Catalog struct {
				ID                    string     `json:"Id,omitempty"`
				Name                  string     `json:"Name,omitempty"`
				LifecycleState        int        `json:"LifecycleState,omitempty"`
				ProvisioningType      int        `json:"ProvisioningType,omitempty"`
				PersistentUserChanges int        `json:"PersistentUserChanges,omitempty"`
				IsMachinePhysical     bool       `json:"IsMachinePhysical,omitempty"`
				AllocationType        int        `json:"AllocationType,omitempty"`
				SessionSupport        int        `json:"SessionSupport,omitempty"`
				ProvisioningSchemeID  string     `json:"ProvisioningSchemeId,omitempty"`
				ZoneUID               string     `json:"ZoneUid,omitempty"`
				ZoneName              string     `json:"ZoneName,omitempty"`
				CreatedDate           *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate          *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"Catalog,omitempty"`
			DesktopGroup struct {
				ID                  string     `json:"Id,omitempty"`
				Name                string     `json:"Name,omitempty"`
				IsRemotePC          bool       `json:"IsRemotePC,omitempty"`
				DesktopKind         int        `json:"DesktopKind,omitempty"`
				LifecycleState      int        `json:"LifecycleState,omitempty"`
				SessionSupport      int        `json:"SessionSupport,omitempty"`
				DeliveryType        int        `json:"DeliveryType,omitempty"`
				IsInMaintenanceMode bool       `json:"IsInMaintenanceMode,omitempty"`
				MachineCost         float64    `json:"MachineCost,omitempty"`
				AutoscaleTagID      int        `json:"AutoscaleTagId,omitempty"`
				CreatedDate         *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate        *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"DesktopGroup,omitempty"`
			Hypervisor struct {
				ID             string     `json:"Id,omitempty"`
				Name           string     `json:"Name,omitempty"`
				Type           string     `json:"Type,omitempty"`
				LifecycleState int        `json:"LifecycleState,omitempty"`
				CreatedDate    *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate   *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"Hypervisor,omitempty"`
			MachineCost struct {
				MachineID                  string     `json:"MachineId,omitempty"`
				SpecID                     int        `json:"SpecId,omitempty"`
				CostPerHour                int        `json:"CostPerHour,omitempty"`
				PowerOnComputeCostPerHour  int        `json:"PowerOnComputeCostPerHour,omitempty"`
				PowerOnStorageCostPerHour  int        `json:"PowerOnStorageCostPerHour,omitempty"`
				PowerOffStorageCostPerHour int        `json:"PowerOffStorageCostPerHour,omitempty"`
				CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
				ModifiedDate               *time.Time `json:"ModifiedDate,omitempty"`
			} `json:"MachineCost,omitempty"`
		} `json:"Machine,omitempty"`
		SessionMetrics []struct {
			ID              int        `json:"Id,omitempty"`
			CollectedDate   *time.Time `json:"CollectedDate,omitempty"`
			IcaRttMS        int        `json:"IcaRttMS,omitempty"`
			IcaLatency      int        `json:"IcaLatency,omitempty"`
			ClientL7Latency int        `json:"ClientL7Latency,omitempty"`
			ServerL7Latency int        `json:"ServerL7Latency,omitempty"`
			SessionID       string     `json:"SessionId,omitempty"`
			CreatedDate     *time.Time `json:"CreatedDate,omitempty"`
			ModifiedDate    *time.Time `json:"ModifiedDate,omitempty"`
		} `json:"SessionMetrics,omitempty"`
		SessionRecordingServer struct {
			SessionKey                 string     `json:"SessionKey,omitempty"`
			SessionRecordingServerName string     `json:"SessionRecordingServerName,omitempty"`
			CreatedDate                *time.Time `json:"CreatedDate,omitempty"`
		} `json:"SessionRecordingServer,omitempty"`
		PublishedDesktopName struct {
			ID            int    `json:"Id,omitempty"`
			PublishedName string `json:"PublishedName,omitempty"`
		} `json:"PublishedDesktopName,omitempty"`
		SessionMetricsLatest struct {
			SessionKey         string      `json:"SessionKey,omitempty"`
			ClientL7Latency    int         `json:"ClientL7Latency,omitempty"`
			CloudConnectorName string      `json:"CloudConnectorName,omitempty"`
			CreatedDate        *time.Time  `json:"CreatedDate,omitempty"`
			EdtMtu             string      `json:"EdtMtu,omitempty"`
			GatewayPopName     string      `json:"GatewayPopName,omitempty"`
			HdxConnectionType  int         `json:"HdxConnectionType,omitempty"`
			HdxProtocolName    string      `json:"HdxProtocolName,omitempty"`
			ModifiedDate       **time.Time `json:"ModifiedDate,omitempty"`
			ServerL7Latency    int         `json:"ServerL7Latency,omitempty"`
		} `json:"SessionMetricsLatest,omitempty"`
	} `json:"value,omitempty"`
}
