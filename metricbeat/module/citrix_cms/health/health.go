package health

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created, Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("citrix_cms", "health", New)
}

// config defines the configuration options for the MetricSet.
type config struct {
	Hosts        []string      `config:"hosts"`         // List of hosts to fetch metrics from.
	Period       time.Duration `config:"period"`        // Fetch interval.
	DebugMode    bool          `config:"debug"`         // Enable debug mode.
	CustomerId   string        `config:"customer_id"`   // Customer ID for authentication.
	ClientId     string        `config:"client_id"`     // Client ID for authentication.
	ClientSecret string        `config:"client_secret"` // Client secret for authentication.
	LimitResults int           `config:"limit_results"` // Limit the number of results fetched.
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. This is achieved by embedding mb.BaseMetricSet.
type MetricSet struct {
	mb.BaseMetricSet
	logger                       *logp.Logger
	debug                        bool                             // Debug mode flag.
	customerId                   string                           // Customer ID for API requests.
	clientId                     string                           // Client ID for API requests.
	clientSecret                 string                           // Client secret for API requests.
	authToken                    string                           // Authentication token.
	period                       time.Duration                    // Fetch interval.
	previousTime                 time.Time                        // Timestamp of the previous fetch.
	limitResults                 int                              // Limit for the number of results.
	machineMetricLatest          map[string]MachineMetric_Persist // Cache for machine metrics.
	previousTimeLoadIndex        time.Time                        // Timestamp of the previous fetch.
	previousTimeLoadIndexSummary time.Time                        // Timestamp of the previous fetch.
	previousTimeLogOnSummary     time.Time                        // Timestamp of the previous fetch.
	//
	previousTimeMachineMetric   time.Time // Timestamp of the previous fetch.
	previousTimeMachineSummary  time.Time // Timestamp of the previous fetch.
	previousTimeResourceSummary time.Time // Timestamp of the previous fetch.
	previousTimeResourceUtil    time.Time // Timestamp of the previous fetch.
	previousTimeSessionActivity time.Time // Timestamp of the previous fetch.
	//
	previousTimeSessionFailure  time.Time // Timestamp of the previous fetch.
	previousTimeSessionMetric   time.Time // Timestamp of the previous fetch.
	previousTimeServerOSDesktop time.Time // Timestamp of the previous fetch.

}

// New creates a new instance of the MetricSet. It unpacks any MetricSet-specific
// configuration options and initializes the MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The citrix_cms health metricset is beta.")

	config := &config{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}
	logger := logp.NewLogger(base.FullyQualifiedName())

	return &MetricSet{
		BaseMetricSet: base,
		debug:         config.DebugMode,
		logger:        logger,
		customerId:    config.CustomerId,
		clientId:      config.ClientId,
		clientSecret:  config.ClientSecret,
		period:        config.Period,
		limitResults:  config.LimitResults,
	}, nil
}

// Fetch gathers data and converts it to the appropriate format. It publishes
// the event, which is then forwarded to the output. In case of an error, it
// sets the Error field of mb.Event or calls report.Error().
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {

	// Initialize the previous fetch time if it's the first fetch.
	// Initialize to time period subtracted from the current time.
	if m.previousTime.IsZero() {
		m.previousTime = time.Now().Add(-m.period)
		m.previousTimeLoadIndex = time.Now().Add(-m.period)
		m.previousTimeLoadIndexSummary = time.Now().Add(-m.period)
		m.previousTimeLogOnSummary = time.Now().Add(-m.period)
		//m.previousTimeMachineDetails = time.Now().Add(-m.period)
		m.previousTimeMachineMetric = time.Now().Add(-m.period)
		m.previousTimeMachineSummary = time.Now().Add(-m.period)
		m.previousTimeResourceSummary = time.Now().Add(-m.period)
		m.previousTimeResourceUtil = time.Now().Add(-m.period)
		m.previousTimeSessionActivity = time.Now().Add(-m.period)
		//m.previousTimeSessionDetails = time.Now().Add(-m.period)
		m.previousTimeSessionFailure = time.Now().Add(-m.period)
		m.previousTimeSessionMetric = time.Now().Add(-m.period)
		m.previousTimeServerOSDesktop = time.Now().Add(-m.period)
	}

	// Setup connection information for the current fetch.
	hostInfo := Connection{
		baseurl:      m.Host(),
		customerId:   m.customerId,
		clientId:     m.clientId,
		clientSecret: m.clientSecret,
	}

	// Initialize the machine metrics cache if it's nil.
	if m.machineMetricLatest == nil {
		m.machineMetricLatest = make(map[string]MachineMetric_Persist)
	}

	// Determine the limit for the number of results.
	limit_results := m.limitResults
	if limit_results <= 0 {
		limit_results = 1000
	}

	// Initialize a structure to hold all metric data.
	var metricData CMSData

	currentTime := time.Now()
	// Update the previous fetch time.
	m.previousTime = currentTime

	// Fetch various metrics using the GetMetrics function.
	// Each API call is wrapped in error handling and logging.
	// Note: I am suspicious about this API endpoint. It seems to always return a ModifiedDate every time you call it.
	// However I cannot see the load changing, so I am not sure if it is correct. It might only be measured once a day.
	// Else this might be a running average that is updated every time the API is called. Because modified date must not be called until
	// the API is called, when I set filter by ModifiedDate to be greater than the previous time, it does not return any data. So I am
	// calling for 24 hours of data and then removing the duplicates.
	currentTime = time.Now()
	LoadIndexesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+LoadIndexes_API+"?"+url.PathEscape(LoadIndexes_API_PATH(m.previousTimeLoadIndex, limit_results)), metricData.loadIndexes)
	if err != nil {
		m.logger.Warnf("GetMetrics LoadIndex failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.loadIndexes = LoadIndexesData.(LoadIndexes_JSON)
		metricData.loadIndexes.Message = message
		RemoveDuplicateLoadIndexEntries(&metricData.loadIndexes)
	}
	m.previousTimeLoadIndex = currentTime

	// Fetch Load Index Summaries metrics.
	currentTime = time.Now()
	LoadIndexSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+LoadIndexSummaries_API+"?"+url.PathEscape(LoadIndexSummaries_API_PATH(m.previousTimeLoadIndexSummary, limit_results)), metricData.loadIndexSummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics LoadIndexSummaries failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.loadIndexSummaries = LoadIndexSummariesData.(LoadIndexSummaries_JSON)
		metricData.loadIndexSummaries.Message = message
	}
	m.previousTimeLoadIndexSummary = currentTime

	// Fetch Log On Summaries metrics.
	currentTime = time.Now()
	LogOnSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+LogOnSummaries_API+"?"+url.PathEscape(LogOnSummaries_API_PATH(m.previousTimeLogOnSummary, limit_results)), metricData.logOnSummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics LogOnSummaries failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.logOnSummaries = LogOnSummariesData.(LogOnSummaries_JSON)
		metricData.logOnSummaries.Message = message
	}
	m.previousTimeLogOnSummary = currentTime

	// Fetch Machine Details metrics.
	MachineDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+Machines_Details_API+"?"+url.PathEscape(Machines_Details_API_PATH(limit_results)), metricData.machineDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics MachineDetails failed; %v", err)
	} else {
		metricData.machineDetails = MachineDetailsData.(MachineDetails_JSON)
		metricData.machineDetails.Message = message
	}

	// Fetch Machine Metric Details metrics.
	currentTime = time.Now()
	MachineMetricDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+MachineMetric_Details_API+"?"+url.PathEscape(MachineMetric_Details_API_PATH(m.previousTimeMachineMetric, limit_results)), metricData.machineMetricDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics MachineMetricDetails failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.machineMetricDetails = MachineMetricDetailsData.(MachineMetricDetails_JSON)
		metricData.machineMetricDetails.Message = message
		m.machineMetricLatest = RemoveMachineMetricDetailsByCollectedDate(&metricData.machineMetricDetails, m.machineMetricLatest)
	}
	m.previousTimeMachineMetric = currentTime

	// Fetch Machine Summaries metrics.
	currentTime = time.Now()
	MachineSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+MachineSummaries_API+"?"+url.PathEscape(MachineSummaries_API_PATH(m.previousTimeMachineSummary, limit_results)), metricData.machineSummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics MachineSummaries failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.machineSummaries = MachineSummariesData.(MachineSummaries_JSON)
		metricData.machineSummaries.Message = message
	}
	m.previousTimeMachineSummary = currentTime

	// Fetch Resource Utilization Summary metrics.
	currentTime = time.Now()
	ResourceUtilizationSummaryData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+ResourceUtilizationSummary_API+"?"+url.PathEscape(ResourceUtilizationSummary_API_PATH(m.previousTimeResourceSummary, limit_results)), metricData.resourceUtilizationSummary)
	if err != nil {
		m.logger.Warnf("GetMetrics ResourceUtilizationSummary failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.resourceUtilizationSummary = ResourceUtilizationSummaryData.(ResourceUtilizationSummary_JSON)
		metricData.resourceUtilizationSummary.Message = message
	}
	m.previousTimeResourceSummary = currentTime

	// Fetch Resource Utilization metrics.
	currentTime = time.Now()
	ResourceUtilizationData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+ResourceUtilization_API+"?"+url.PathEscape(ResourceUtilization_API_PATH(m.previousTimeResourceUtil, limit_results)), metricData.resourceUtilization)
	if err != nil {
		m.logger.Warnf("GetMetrics ResourceUtilization failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.resourceUtilization = ResourceUtilizationData.(ResourceUtilization_JSON)
		metricData.resourceUtilization.Message = message
	}
	m.previousTimeResourceUtil = currentTime

	// Fetch Session Activity Summaries metrics.
	currentTime = time.Now()
	SessionActivitySummariesDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+SessionActivitySummaries_details_API+"?"+url.PathEscape(SessionActivitySummaries_details_API_PATH(m.previousTimeSessionActivity, limit_results)), metricData.sessionActivitySummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics SessionActivitySummaries failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.sessionActivitySummaries = SessionActivitySummariesDetailsData.(SessionActivitySummaries_JSON)
		metricData.sessionActivitySummaries.Message = message
	}
	m.previousTimeSessionActivity = currentTime

	// Fetch Session Details metrics.
	SessionsDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+Sessions_Details_API+"?"+url.PathEscape(SessionsActive_Details_API_PATH(limit_results)), metricData.sessionDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics SessionDetails failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.sessionDetails = SessionsDetailsData.(SessionsDetails_JSON)
		metricData.sessionDetails.Message = message
	}

	// Fetch Session Failure metrics.
	currentTime = time.Now()
	SessionsFailureData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+Sessions_Details_API+"?"+url.PathEscape(SessionsFailure_Details_API_PATH(m.previousTimeSessionFailure, limit_results)), metricData.sessionFailureDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics SessionsFailure failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.sessionFailureDetails = SessionsFailureData.(SessionsDetails_JSON)
		metricData.sessionFailureDetails.Message = message
	}
	m.previousTimeSessionFailure = currentTime

	// Fetch Session Metric Details metrics.
	currentTime = time.Now()
	SessionMetricDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+SessionMetrics_Details_API+"?"+url.PathEscape(SessionMetrics_Details_API_PATH(m.previousTimeSessionMetric, limit_results)), metricData.sessionMetricDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics SessionMetricDetails failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.sessionMetricDetails = SessionMetricDetailsData.(SessionMetricDetails_JSON)
		metricData.sessionMetricDetails.Message = message
	}
	m.previousTimeSessionMetric = currentTime

	// Fetch Server OS Desktop Summaries metrics.
	// I have been unable to test this API endpoint, so I am not sure if it is correct.
	// On test server it never returns any data to test against.
	currentTime = time.Now()
	ServerOSDesktopSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+ServerOSDesktopSummaries_API+"?"+url.PathEscape(ServerOSDesktopSummaries_API_PATH(m.previousTimeServerOSDesktop, limit_results)), metricData.serverOSDesktopSummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics ServerOSDesktopSummaries failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		m.logger.Warn(string(b))
	} else {
		metricData.serverOSDesktopSummaries = ServerOSDesktopSummariesData.(ServerOSDesktopSummaries_JSON)
		metricData.serverOSDesktopSummaries.Message = message
	}
	m.previousTimeServerOSDesktop = currentTime

	// Report the collected metrics.
	reportMetrics(reporter, hostInfo.baseurl, metricData, m.debug)

	return nil
}

// Connection holds the connection details for API requests.
type Connection struct {
	baseurl      string // Base URL of the API.
	customerId   string // Customer ID for authentication.
	clientId     string // Client ID for authentication.
	clientSecret string // Client secret for authentication.
	authToken    string // Authentication token.
}
