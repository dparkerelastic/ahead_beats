package health

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("citrix_cms", "health", New)
}

type config struct {
	Hosts        []string      `config:"hosts"`
	Period       time.Duration `config:"period"`
	DebugMode    bool          `config:"debug"`
	CustomerId   string        `config:"customer_id"`
	ClientId     string        `config:"client_id"`
	ClientSecret string        `config:"client_secret"`
	LimitResults int           `config:"limit_results"`
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	logger *logp.Logger
	//counter      int
	debug        bool
	customerId   string
	clientId     string
	clientSecret string
	authToken    string
	period       time.Duration
	previousTime time.Time
	limitResults int
	// machineMetricSummaryTime time.Time
	// machineMetricDetailsTime time.Time
	machineMetricLatest map[string]MachineMetric_Persist // MachineID as the key
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The citrix_cms health metricset is beta.")

	//config := struct{}{}
	//config := defaultConfig()
	config := &config{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}
	logger := logp.NewLogger(base.FullyQualifiedName())

	//machineMetricLatest := make(map[string]MachineMetric_Persist)

	return &MetricSet{
		BaseMetricSet: base,
		//counter:       1,
		debug:        config.DebugMode,
		logger:       logger,
		customerId:   config.CustomerId,
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
		period:       config.Period,
		limitResults: config.LimitResults,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {
	currentTime := time.Now()

	fmt.Println("Code is Running")
	fmt.Println("Current time in desired format:", currentTime.UTC().Format("2006-01-02T15:04:05Z"))

	if m.previousTime.IsZero() {
		m.previousTime = time.Now().Add(-m.period)
		fmt.Println("Previous time in desired format:", m.previousTime.UTC().Format("2006-01-02T15:04:05Z"))
	} else {
		fmt.Println("Previous time was already set:", m.previousTime.UTC().Format("2006-01-02T15:04:05Z"))

	}

	//Setup Connection Info for this Fetch
	hostInfo := Connection{}
	hostInfo.baseurl = m.Host()
	hostInfo.customerId = m.customerId
	hostInfo.clientId = m.clientId
	hostInfo.clientSecret = m.clientSecret

	if m.machineMetricLatest == nil {
		m.machineMetricLatest = make(map[string]MachineMetric_Persist)
	}

	var limit_results int
	if m.limitResults > 0 {
		limit_results = m.limitResults
	} else {
		limit_results = 1000
	}

	var metricData CMSData

	LoadIndexesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+LoadIndexes_API+"?"+url.PathEscape(LoadIndexes_API_PATH(m.previousTime, limit_results)), metricData.loadIndexes)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		fmt.Println("##############################")
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.loadIndexes = LoadIndexesData.(LoadIndexes_JSON)
		metricData.loadIndexes.Message = message
	}

	LoadIndexSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+LoadIndexSummaries_API+"?"+url.PathEscape(LoadIndexSummaries_API_PATH(m.previousTime, limit_results)), metricData.loadIndexSummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		fmt.Println("##############################")
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.loadIndexSummaries = LoadIndexSummariesData.(LoadIndexSummaries_JSON)
		metricData.loadIndexSummaries.Message = message
	}

	LogOnSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+LogOnSummaries_API+"?"+url.PathEscape(LogOnSummaries_API_PATH(m.previousTime, limit_results)), metricData.logOnSummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {

		metricData.logOnSummaries = LogOnSummariesData.(LogOnSummaries_JSON)
		metricData.logOnSummaries.Message = message
	}

	MachineDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+Machines_Details_API+"?"+url.PathEscape(Machines_Details_API_PATH(limit_results)), metricData.machineDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.machineDetails = MachineDetailsData.(MachineDetails_JSON)
		metricData.machineDetails.Message = message
	}

	MachineMetricDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+MachineMetric_Details_API+"?"+url.PathEscape(MachineMetric_Details_API_PATH(m.previousTime, limit_results)), metricData.machineMetricDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.machineMetricDetails = MachineMetricDetailsData.(MachineMetricDetails_JSON)
		metricData.machineMetricDetails.Message = message
		m.machineMetricLatest = RemoveMachineMetricDetailsByCollectedDate(&metricData.machineMetricDetails, m.machineMetricLatest)

	}

	MachineSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+MachineSummaries_API+"?"+url.PathEscape(MachineSummaries_API_PATH(m.previousTime, limit_results)), metricData.machineSummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {

		metricData.machineSummaries = MachineSummariesData.(MachineSummaries_JSON)
		metricData.machineSummaries.Message = message

	}

	ResourceUtilizationSummaryData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+ResourceUtilizationSummary_API+"?"+url.PathEscape(ResourceUtilizationSummary_API_PATH(m.previousTime, limit_results)), metricData.resourceUtilizationSummary)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.resourceUtilizationSummary = ResourceUtilizationSummaryData.(ResourceUtilizationSummary_JSON)
		metricData.resourceUtilizationSummary.Message = message
	}

	ResourceUtilizationData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+ResourceUtilization_API+"?"+url.PathEscape(ResourceUtilization_API_PATH(m.previousTime, limit_results)), metricData.resourceUtilization)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {

		metricData.resourceUtilization = ResourceUtilizationData.(ResourceUtilization_JSON)
		metricData.resourceUtilization.Message = message
	}

	SessionActivitySummariesDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+SessionActivitySummaries_details_API+"?"+url.PathEscape(SessionActivitySummaries_details_API_PATH(m.previousTime, limit_results)), metricData.sessionActivitySummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {

		metricData.sessionActivitySummaries = SessionActivitySummariesDetailsData.(SessionActivitySummaries_JSON)
		metricData.sessionActivitySummaries.Message = message
	}

	SessionsFailureData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+Sessions_Details_API+"?"+url.PathEscape(SessionsFailure_Details_API_PATH(m.previousTime, limit_results)), metricData.sessionFailureDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.sessionFailureDetails = SessionsFailureData.(SessionsDetails_JSON)
		metricData.sessionFailureDetails.Message = message
	}

	SessionMetricDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+SessionMetrics_Details_API+"?"+url.PathEscape(SessionMetrics_Details_API_PATH(m.previousTime, limit_results)), metricData.sessionMetricDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.sessionMetricDetails = SessionMetricDetailsData.(SessionMetricDetails_JSON)
		metricData.sessionMetricDetails.Message = message

	}
	SessionsDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+Sessions_Details_API+"?"+url.PathEscape(SessionsActive_Details_API_PATH(limit_results)), metricData.sessionDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.sessionDetails = SessionsDetailsData.(SessionsDetails_JSON)
		metricData.sessionDetails.Message = message
	}

	ServerOSDesktopSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+ServerOSDesktopSummaries_API+"?"+url.PathEscape(Count_API), metricData.serverOSDesktopSummaries)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		fmt.Println("##############################")
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.serverOSDesktopSummaries = ServerOSDesktopSummariesData.(ServerOSDesktopSummaries_JSON)
		metricData.serverOSDesktopSummaries.Message = message
	}

	reportMetrics(reporter, hostInfo.baseurl, metricData, m.debug)

	fmt.Println("Done - Current time in desired format:", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	m.previousTime = currentTime

	return nil
}

type Connection struct {
	baseurl      string
	customerId   string
	clientId     string
	clientSecret string
	authToken    string
}
