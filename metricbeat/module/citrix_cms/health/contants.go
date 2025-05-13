package health

import (
	"strconv"
	"time"
)

// Count_API is a constant used to include a count of results in the API query.
const Count_API = "$count=true"

// LoadIndexes_API is the endpoint for fetching load indexes.
const LoadIndexes_API = "/monitorodata/LoadIndexes"

// LoadIndexes_API_PATH generates a query string for fetching load indexes
// filtered by the ModifiedDate greater than the provided previousTime
// and limits the results to the specified limit_results.
func LoadIndexes_API_PATH(previousTime time.Time, limit_results int) string {
	//return "$filter=ModifiedDate lt " + previousTime.Add(-5*time.Second).UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
	return "$filter=CreatedDate gt " + time.Now().UTC().Add(-24*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
	//return "$count=true&%$top=" + strconv.Itoa(limit_results)
}

// LoadIndexSummaries_API is the endpoint for fetching load index summaries.
const LoadIndexSummaries_API = "/monitorodata/LoadIndexSummaries"

// LoadIndexSummaries_API_PATH generates a query string for fetching load index summaries
// filtered by the ModifiedDate greater than the provided previousTime
// and limits the results to the specified limit_results.
func LoadIndexSummaries_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
	//return "$count=true&%$top=" + strconv.Itoa(limit_results)
}

// LogOnSummaries_API is the endpoint for fetching log-on summaries.
const LogOnSummaries_API = "/monitorodata/LogOnSummaries"

// LogOnSummaries_API_PATH generates a query string for fetching log-on summaries
// filtered by the ModifiedDate greater than the provided previousTime
// and limits the results to the specified limit_results.
func LogOnSummaries_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// MachineMetric_Details_API is the endpoint for fetching machine metric details.
const MachineMetric_Details_API = "/monitorodata/MachineMetric"

// MachineMetric_Details_API_PATH generates a query string for fetching machine metric details.
// Note: This API does not have a ModifiedDate field, so the CollectedDate field is used instead.
// The query filters results by CollectedDate greater than 2 hours before the current time
// and limits the results to the specified limit_results.
func MachineMetric_Details_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=CollectedDate gt " + time.Now().UTC().Add(-2*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// MachineMetricSummary_API is the endpoint for fetching machine metric summaries.
const MachineMetricSummary_API = "/monitorodata/MachineMetricSummary"

// MachineMetricSummary_API_PATH generates a query string for fetching machine metric summaries.
// The query filters results by SummaryDate greater than 48 hours before the current time
// and limits the results to the specified limit_results.
func MachineMetricSummary_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=SummaryDate gt " + time.Now().UTC().Add(-48*time.Hour).Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// Machines_Details_API is the endpoint for fetching machine details.
const Machines_Details_API = "/monitorodata/Machines"

// Machines_Details_API_PATH generates a query string for fetching machine details.
// The query expands related entities and filters results by LifecycleState equal to 0.
// It also limits the results to the specified limit_results.
func Machines_Details_API_PATH(limit_results int) string {
	return "$expand=CurrentLoadIndex,Catalog,DesktopGroup,Hypervisor,MachineCost&$filter=LifecycleState eq 0&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// MachineSummaries_API is the endpoint for fetching machine summaries.
const MachineSummaries_API = "/monitorodata/MachineSummaries"

// MachineSummaries_API_PATH generates a query string for fetching machine summaries
// filtered by the ModifiedDate greater than the provided previousTime
// and limits the results to the specified limit_results.
func MachineSummaries_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// ResourceUtilization_API is the endpoint for fetching resource utilization data.
const ResourceUtilization_API = "/monitorodata/ResourceUtilization"

// ResourceUtilization_API_PATH generates a query string for fetching resource utilization data
// filtered by the ModifiedDate greater than the provided previousTime
// and limits the results to the specified limit_results.
func ResourceUtilization_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// ResourceUtilizationSummary_API is the endpoint for fetching resource utilization summaries.
const ResourceUtilizationSummary_API = "/monitorodata/ResourceUtilizationSummary"

// ResourceUtilizationSummary_API_PATH generates a query string for fetching resource utilization summaries
// filtered by the ModifiedDate greater than the provided previousTime
// and limits the results to the specified limit_results.
func ResourceUtilizationSummary_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// SessionActivitySummaries_details_API is the endpoint for fetching session activity summaries.
const SessionActivitySummaries_details_API = "/monitorodata/MachineSummaries"

// SessionActivitySummaries_details_API_PATH generates a query string for fetching session activity summaries
// filtered by the ModifiedDate greater than the provided previousTime
// and limits the results to the specified limit_results.
func SessionActivitySummaries_details_API_PATH(previousTime time.Time, limit_results int) string {
	return "$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// SessionMetrics_Details_API is the endpoint for fetching session metrics details.
const SessionMetrics_Details_API = "/monitorodata/SessionMetrics"

// SessionMetrics_Details_API_PATH generates a query string for fetching session metrics details.
// The query expands the Session entity and filters results by the ModifiedDate greater than the provided previousTime.
// It also limits the results to the specified limit_results.
func SessionMetrics_Details_API_PATH(previousTime time.Time, limit_results int) string {
	return "$expand=Session&$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// Sessions_Details_API is the endpoint for fetching session details.
const Sessions_Details_API = "/monitorodata/Sessions"

// SessionsActive_Details_API_PATH generates a query string for fetching active session details.
// The query expands related entities and filters results by EndDate being null.
// It also limits the results to the specified limit_results.
func SessionsActive_Details_API_PATH(limit_results int) string {
	return "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,SessionMetrics,SessionMetricsLatest&$filter=EndDate eq null&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// SessionsFailure_Details_API_PATH generates a query string for fetching session failure details.
// The query expands related entities and filters results by FailureId not being null and FailureDate
// greater than the provided previousTime. It also limits the results to the specified limit_results.
func SessionsFailure_Details_API_PATH(previousTime time.Time, limit_results int) string {
	return "$expand=Failure,CurrentConnection,CurrentConnection($expand=ConnectionFailureLog),User,Machine,SessionMetrics,SessionMetricsLatest&$filter=FailureId ne null and FailureDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}

// ServerOSDesktopSummaries_API is the endpoint for fetching server OS desktop summaries.
const ServerOSDesktopSummaries_API = "/monitorodata/ServerOSDesktopSummaries"

// ServerOSDesktopSummaries_API_PATH generates a query string for fetching session metrics details.
// The query expands the Session entity and filters results by the ModifiedDate greater than the provided previousTime.
// It also limits the results to the specified limit_results.
func ServerOSDesktopSummaries_API_PATH(previousTime time.Time, limit_results int) string {
	return "$expand=Session&$filter=ModifiedDate gt " + previousTime.UTC().Format("2006-01-02T15:04:05Z") + "&$count=true&%$top=" + strconv.Itoa(limit_results)
}
