package health

import "time"

const Count_API = "$count=true"

const LoadIndexes_API = "/monitorodata/LoadIndexes"

func LoadIndexes_API_PATH() string {
	return "$count=true&%$top=1000"
}

const LoadIndexSummaries_API = "/monitorodata/LoadIndexSummaries"

func LoadIndexSummaries_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-10*time.Minute).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const LogOnSummaries_API = "/monitorodata/LogOnSummaries"

func LogOnSummaries_API_PATH() string {
	return "$filter=Granularity eq 1440 and SummaryDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const Machines_Details_API = "/monitorodata/Machines"
const Machines_Details_API_PATH = "$expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost&$filter=LifecycleState eq 0&$count=true&%$top=1000"

const MachineSummaries_API = "/monitorodata/MachineSummaries"

func MachineSummaries_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const MachineMetricSummary_API = "/monitorodata/MachineMetricSummary"

func MachineMetricSummary_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-48*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const MachineMetric_Details_API = "/monitorodata/MachineMetric"

func MachineMetric_Details_API_PATH() string {
	return "$filter=CollectedDate gt " + time.Now().UTC().Add(-24*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"

}

const ResourceUtilization_API = "/monitorodata/ResourceUtilization"

func ResourceUtilization_API_PATH() string {
	return "$filter=CollectedDate gt " + time.Now().UTC().Add(-10*time.Minute).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const ResourceUtilizationSummary_API = "/monitorodata/ResourceUtilizationSummary"

func ResourceUtilizationSummary_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const ServerOSDesktopSummaries_API = "/monitorodata/ServerOSDesktopSummaries"

const SessionActivitySummaries_details_API = "/monitorodata/MachineSummaries"

func SessionActivitySummaries_details_API_PATH() string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const SessionMetrics_Details_API = "/monitorodata/SessionMetrics"

func SessionMetrics_Details_API_PATH() string {
	return "$expand=Session&$filter=CollectedDate gt " + time.Now().UTC().Add(-10*time.Minute).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}

const Sessions_Details_API = "/monitorodata/Sessions"

const SessionsActive_Details_API_PATH = "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,SessionMetrics,SessionMetricsLatest&$filter=EndDate eq null&$count=true&%$top=1000"

func SessionsFailure_Details_API_PATH() string {
	return "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,SessionMetrics,SessionMetricsLatest&$filter=FailureId ne null and StartDate gt " + time.Now().UTC().Add(-24*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=1000"
}
