package health

import (
	"encoding/json"
	"fmt"
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

// func defaultConfig() *config {
// 	return &config{
// 		//Hosts:     []string{"localhost"},
// 		DebugMode: false,
// 		//Period:    time.Second * 10,
// 	}
// }

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

// type MetricSet struct {
// 	mb.BaseMetricSet
// 	counter int
// }

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

	// Uncomment the following lines to debug the connection info
	// fmt.Println("Host Info: ", hostInfo)
	// fmt.Println("hostinfo.baseurl: ", hostInfo.baseurl)
	// fmt.Println("hostinfo.customerId: ", hostInfo.customerId)
	// fmt.Println("hostinfo.clientId: ", hostInfo.clientId)
	// fmt.Println("hostinfo.clientSecret: ", hostInfo.clientSecret)

	var metricData CMSData

	machineCurrentLoadIndexData, message, err := GetMetrics(m, hostInfo, hostInfo.baseurl+MachineLoadIndex_API, metricData.machineCurrentLoadIndex)
	if err != nil {
		m.logger.Warnf("GetSystemMetrics failed; %v", err)
		fmt.Println("##############################")
		b, _ := json.MarshalIndent(message, "", "  ")
		fmt.Print(string(b))

	} else {
		metricData.machineCurrentLoadIndex = machineCurrentLoadIndexData.(MachinesResponse_JSON)
		//metricData.machineCurrentLoadIndex.Message = message
		metricData.machineCurrentLoadIndex.Message = "DansSuccess"

		//pprint(metricData.machineCurrentLoadIndex)
		//fmt.Println("\n\n##############################")
		//fmt.Printf("\n\nSystem Response: %+v", metricData.machineCurrentLoadIndex)

		// fmt.Println("\n\n##############################")
		// fmt.Printf("\n\nSystem Response: %+v", machineCurrentLoadIndexData)

		// fmt.Println("\n\n##############################")
		// b, _ := json.MarshalIndent(metricData.machineCurrentLoadIndex, "", "  ")
		// fmt.Print(string(b))

	}

	reportMetrics(reporter, hostInfo.baseurl, metricData, m.debug)

	fmt.Println("##############################")
	fmt.Println("Timestamp fetchMachineData at:", time.Now().Format(time.RFC3339))

	return nil
}

type Connection struct {
	baseurl      string
	customerId   string
	clientId     string
	clientSecret string
	authToken    string
}
