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
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	//config  *config
	logger       *logp.Logger
	counter      int
	debug        bool
	customerId   string
	clientId     string
	clientSecret string
	authToken    string
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

	return &MetricSet{
		BaseMetricSet: base,
		counter:       1,
		debug:         config.DebugMode,
		logger:        logger,
		customerId:    config.CustomerId,
		clientId:      config.ClientId,
		clientSecret:  config.ClientSecret,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {
	fmt.Println("Code is Running")

	//Setup Connection Info for this Fetch
	hostInfo := Connection{}
	hostInfo.baseurl = m.Host()
	hostInfo.customerId = m.customerId
	hostInfo.clientId = m.clientId
	hostInfo.clientSecret = m.clientSecret

	var metricData CMSData

	// ServerOSDesktopSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+ServerOSDesktopSummaries_API+"?"+url.PathEscape(Count_API), metricData.serverOSDesktopSummaries)
	// if err != nil {
	// 	m.logger.Warnf("GetMetrics failed; %v", err)
	// 	fmt.Println("##############################")
	// 	b, _ := json.MarshalIndent(message, "", "  ")
	// 	fmt.Print(string(b))

	// } else {
	// 	metricData.serverOSDesktopSummaries = ServerOSDesktopSummariesData.(ServerOSDesktopSummaries_JSON)
	// 	metricData.serverOSDesktopSummaries.Message = message
	// }
	// // ##### Harvested via Machine Details API
	// LoadIndexSummariesData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+LoadIndexSummaries_API+"?"+url.PathEscape(LoadIndexSummaries_API_PATH), metricData.loadIndexSummaries)
	// if err != nil {
	// 	m.logger.Warnf("GetMetrics failed; %v", err)
	// 	fmt.Println("##############################")
	// 	b, _ := json.MarshalIndent(message, "", "  ")
	// 	fmt.Print(string(b))

	// } else {
	// 	metricData.loadIndexSummaries = LoadIndexSummariesData.(LoadIndexSummaries_JSON)
	// 	metricData.loadIndexSummaries.Message = message
	// }

	// MachineMetricData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+MachineMetric_API+"?"+url.PathEscape(Count_API), metricData.machineMetric)
	// if err != nil {
	// 	m.logger.Warnf("GetMetrics failed; %v", err)
	// 	b, _ := json.MarshalIndent(message, "", "  ")
	// 	fmt.Print(string(b))

	// } else {
	// 	metricData.machineMetric = MachineMetricData.(MachineMetric_JSON)
	// 	metricData.machineMetric.Message = message
	// }

	// SessionActivitySummaries_Agg1_Data, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+SessionActivitySummaries_API+"?"+url.PathEscape(SessionActivitySummaries_API_PATH), metricData.sessionActivitySummaries_Agg1)
	// if err != nil {
	// 	m.logger.Warnf("GetMetrics failed; %v", err)
	// 	b, _ := json.MarshalIndent(message, "", "  ")
	// 	fmt.Print(string(b))

	// } else {
	// 	metricData.sessionActivitySummaries_Agg1 = SessionActivitySummaries_Agg1_Data.(SessionActivitySummaries_Agg1_JSON)
	// 	metricData.sessionActivitySummaries_Agg1.Message = message
	// }

	// MachineMetricDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+MachineMetric_Details_API+"?"+url.PathEscape(MachineMetric_Details_API_PATH), metricData.machineMetricDetails)
	// if err != nil {
	// 	m.logger.Warnf("GetMetrics failed; %v", err)
	// 	b, _ := json.MarshalIndent(message, "", "  ")
	// 	fmt.Print(string(b))

	// } else {
	// 	metricData.machineMetricDetails = MachineMetricDetailsData.(MachineMetricDetails_JSON)
	// 	metricData.machineMetricDetails.Message = message
	// }

	// LogonMetricDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+LogonMetric_Details_API+"?"+url.PathEscape(LogonMetric_Details_API_PATH), metricData.logonMetricDetails)
	// if err != nil {
	// 	m.logger.Warnf("GetMetrics failed; %v", err)
	// 	b, _ := json.MarshalIndent(message, "", "  ")
	// 	fmt.Print(string(b))

	// } else {
	// 	metricData.logonMetricDetails = LogonMetricDetailsData.(LogonMetricsDetails_JSON)
	// 	metricData.logonMetricDetails.Message = message
	// }

	//SessionMetricsDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+SessionMetrics_Details_API+"?"+url.PathEscape("$filter=CollectedDate  gt "+time.Now().UTC().Add(-10*time.Minute).Format("2006-01-02T15:04:05Z")+"&$count=true&%$top=1000"), metricData.sessionMetricsDetails)
	// SessionMetricsDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+SessionMetrics_Details_API+"?"+url.PathEscape(SessionMetrics_Details_API_PATH), metricData.sessionMetricsDetails)

	// if err != nil {
	// 	m.logger.Warnf("GetMetrics failed; %v", err)
	// 	b, _ := json.MarshalIndent(message, "", "  ")
	// 	fmt.Print(string(b))

	// } else {
	// 	metricData.sessionMetricsDetails = SessionMetricsDetailsData.(SessionMetricsDetails_JSON)
	// 	metricData.sessionMetricsDetails.Message = message
	// }

	SessionsDetailsData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+Sessions_Details_API+"?"+url.PathEscape(SessionsActive_Details_API_PATH), metricData.sessionDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.sessionDetails = SessionsDetailsData.(SessionsDetails_JSON)
		metricData.sessionDetails.Message = message
	}

	SessionsFailureData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+Sessions_Details_API+"?"+url.PathEscape(SessionsFailure_Details_API_PATH), metricData.sessionFailureDetails)
	if err != nil {
		m.logger.Warnf("GetMetrics failed; %v", err)
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.sessionFailureDetails = SessionsFailureData.(SessionsDetails_JSON)
		metricData.sessionFailureDetails.Message = message
	}

	reportMetrics(reporter, hostInfo.baseurl, metricData, m.debug)

	fmt.Println("Timestamp fetchMachineData at:", time.Now().Format(time.RFC3339))

	fmt.Println("Current time in desired format:", time.Now().UTC().Format("2006-01-02T15:04:05Z"))

	fmt.Println("Time 5 minutes ago in desired format:", time.Now().UTC().Add(-5*time.Minute).Format("2006-01-02T15:04:05Z"))

	return nil
}

type Connection struct {
	baseurl      string
	customerId   string
	clientId     string
	clientSecret string
	authToken    string
}
