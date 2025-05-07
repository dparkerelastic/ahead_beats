package health

import (
	"strconv"
	"time"
)

const Count_API = "$count=true"

const LoadIndexes_API = "/monitorodata/LoadIndexes"

func LoadIndexes_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const LoadIndexSummaries_API = "/monitorodata/LoadIndexSummaries"

func LoadIndexSummaries_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const LogOnSummaries_API = "/monitorodata/LogOnSummaries"

func LogOnSummaries_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const MachineMetric_Details_API = "/monitorodata/MachineMetric"

// IMO this Citrix API has a bug, it does not have ModifiedDate field. Therefore, we are using CollectedDate field and have to hardcode the time
// to be 2 hours before the current time. This is not a good practice, but we have to do it for now.
func MachineMetric_Details_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=CollectedDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)

}

const MachineMetricSummary_API = "/monitorodata/MachineMetricSummary"

func MachineMetricSummary_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-48*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const Machines_Details_API = "/monitorodata/Machines"

func Machines_Details_API_PATH(limit_results int) string {
	return "$expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost&$filter=LifecycleState eq 0&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const MachineSummaries_API = "/monitorodata/MachineSummaries"

func MachineSummaries_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const ResourceUtilization_API = "/monitorodata/ResourceUtilization"

func ResourceUtilization_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const ResourceUtilizationSummary_API = "/monitorodata/ResourceUtilizationSummary"

func ResourceUtilizationSummary_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const ServerOSDesktopSummaries_API = "/monitorodata/ServerOSDesktopSummaries"

const SessionActivitySummaries_details_API = "/monitorodata/MachineSummaries"

func SessionActivitySummaries_details_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const SessionMetrics_Details_API = "/monitorodata/SessionMetrics"

func SessionMetrics_Details_API_PATH(previousTime time.Time, limit_results int) string {
	return "$expand=Session&$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

const Sessions_Details_API = "/monitorodata/Sessions"

func SessionsActive_Details_API_PATH(limit_results int) string {
	return "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,SessionMetrics,SessionMetricsLatest&$filter=EndDate eq null&$count=true&%$top=" + strconv.Itoa(limit_results)
}

func SessionsFailure_Details_API_PATH(previousTime time.Time, limit_results int) string {
	return "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,SessionMetrics,SessionMetricsLatest&$filter=FailureId ne null and FailureDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}
