package health

import "time"

const Count_API = "$count=true"

const ServerOSDesktopSummaries_API = "/monitorodata/ServerOSDesktopSummaries"
const LoadIndexSummaries_API = "/monitorodata/LoadIndexSummaries"
const LoadIndexSummaries_API_PATH = "$filter=Granularity eq 60 and CreatedDate gt 2024-01-01T00:00:00.000Z&$top=1000&$count=true"
const MachineMetric_API = "/monitorodata/MachineMetric"
const SessionActivitySummaries_API = "/monitorodata/SessionActivitySummaries"
const SessionActivitySummaries_API_PATH = "$apply=filter(SummaryDate gt 2025-02-01T00:00:00Z)/groupby((DesktopGroupId),aggregate(ConnectedSessionCount with average as AverageConnectedSessionCount,DisconnectedSessionCount with average as AverageDisconnectedSessionCount,ConcurrentSessionCount with average as AverageConcurrentSessionCount,TotalLogOnDuration with average as AverageTotalLogOnDuration,TotalLogOnCount with sum as SumTotalLogOnCount))&$count=true&$top=1000"

// These are the details API endpoints for the metrics
const LogonMetric_Details_API = "/monitorodata/LogOnMetrics"
const LogonMetric_Details_API_PATH = "$expand=Session($expand=CurrentConnection),Session($expand=Machine),Session($expand=User),Session($expand=SessionMetricsLatest)&$count=true&%$top=1000"

const MachineMetric_Details_API = "/monitorodata/MachineMetric"
const MachineMetric_Details_API_PATH = "$expand=Machine($expand=CurrentLoadIndex),Machine($expand=Catalog),Machine($expand=DesktopGroup),Machine($expand=Hypervisor),Machine($expand=MachineCost)&$count=true&%$top=1000"

const SessionMetrics_Details_API = "/monitorodata/SessionMetrics"

// SessionMetrics_Details_API_PATH is dynamically generated in a function
func SessionMetrics_Details_API_PATH() string {
	return "$filter=CollectedDate gt " + time.Now().UTC().Add(-10*time.Minute).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

//const SessionMetrics_Details_API_PATH = "$expand=Session($expand=Failure),Session($expand=CurrentConnection),Session($expand=User),Session($expand=Machine),Session($expand=SessionMetricsLatest)&$count=true&%$top=1000"
//const SessionMetrics_Details_API_PATH = "$filter=Session/EndDate eq null&$count=true&%$top=1000"
//const SessionMetrics_Details_API_PATH = "$apply=filter(Granularity eq 60 and Session/EndDate eq null)/groupby((SessionId),aggregate(IcaRttMS with average as AverageIcaRttMS,IcaLatency with average as AverageIcaLatency,ClientL7Latency with average as AverageClientL7Latency,ServerL7Latency with average as AverageServerL7Latencyn))&$count=true&%$top=1000"

// const SessionMetrics_Details_API_PATH = "$count=true&%$top=1000"

//const SessionMetrics_Details_API_PATH = "$expand=Session($expand=Failure),Session($expand=CurrentConnection),Session($expand=User),Session($expand=Machine;$expand=Machine/DesktopGroup),Session($expand=SessionMetricsLatest)&$count=true&%$top=1000"

const Sessions_Details_API = "/monitorodata/Sessions"

// const SessionsActive_Details_API_PATH = "$filter=EndDate eq null&$expand=Session($expand=Failure)&$count=true&%$top=1000"
const SessionsActive_Details_API_PATH = "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,Machine($expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost),SessionMetrics,SessionMetricsLatest&$filter=EndDate eq null&$count=true&%$top=1000"

// var SessionsFailure_Details_API_PATH = "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,Machine($expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost),SessionMetrics,SessionMetricsLatest&$filter=FailureId ne null and FailureDate gt " + time.Now().UTC().Add(-24*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
var SessionsFailure_Details_API_PATH = "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,Machine($expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost),SessionMetrics,SessionMetricsLatest&$filter=FailureId ne null&$count=true&%$top=1000"

// var SessionsActive_Details_API_PATH = func() string {
// 	return "$filter=SessionMetrics/CollectedDate gt " + time.Now().UTC().Add(-10*time.Minute).Format("2006-01-02T15:04:05Z") + "&$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,Machine($expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost),SessionMetrics,SessionMetricsLatest&$filter=EndDate eq null&$count=true&%$top=1000"
// }()

//const SessionsActive_Details_API_PATH = "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,Machine($expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost),SessionMetricsLatest&$filter=EndDate eq null&$count=true&%$top=1000"

//const SessionsActive_Details_API_PATH = "$count=true&%$top=1000"

const Machines_Details_API = "/monitorodata/Machines"

// const SessionsActive_Details_API_PATH = "$filter=EndDate eq null&$expand=Session($expand=Failure)&$count=true&%$top=1000"
// const Machines_Details_API_PATH = "$expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost,LoadIndex,ResourceUtilization,ProcessUtilization,MachineMetric,MachoneHotfixLogs&$filter=LifecycleState eq 0&$count=true&%$top=1000"
const Machines_Details_API_PATH = "$expand=Catalog,DesktopGroup,Hypervisor,MachineCost&$filter=LifecycleState eq 0&$count=true&%$top=1000"

const ResourceUtilizationSummary_API = "/monitorodata/ResourceUtilizationSummary"

func ResourceUtilizationSummary_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const ResourceUtilization_API = "/monitorodata/ResourceUtilization"

func ResourceUtilization_API_PATH() string {
	return "$filter=CollectedDate gt " + time.Now().UTC().Add(-10*time.Minute).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const LogOnSummaries_API = "/monitorodata/LogOnSummaries"

func LogOnSummaries_API_PATH() string {
	return "$filter=Granularity eq 1440 and SummaryDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
	//return "$filter=Granularity eq 1440 and SummaryDate gt " + time.Now().UTC().Add(-720*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
	//return "$count=true&%$top=1000"
}

const MachineSummaries_API = "/monitorodata/MachineSummaries"

func MachineSummaries_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
	//return "$filter=Granularity eq 1440 and SummaryDate gt " + time.Now().UTC().Add(-720*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
	//return "$count=true&%$top=1000"
}

const MachineMetricSummary_API = "/monitorodata/MachineMetricSummary"

func MachineMetricSummary_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-48*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
	//return "$filter=Granularity eq 1440 and SummaryDate gt " + time.Now().UTC().Add(-720*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
	//return "$count=true&%$top=1000"
}

const SessionActivitySummaries_details_API = "/monitorodata/MachineSummaries"

func SessionActivitySummaries_details_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
	//return "$filter=Granularity eq 1440 and SummaryDate gt " + time.Now().UTC().Add(-720*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
	//return "$count=true&%$top=1000"
}

// var ResourceUtilizationSummary_API_PATH = func() string {
// 	// This API path filters session metrics collected in the last 10 minutes
// 	// to reduce the amount of data returned and improve performance.
// 	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-10*time.Minute).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
// }()
