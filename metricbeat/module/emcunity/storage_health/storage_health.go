package storage_health

import (
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
	mb.Registry.MustAddMetricSet("emcunity", "storage_health", New)
}

type config struct {
	Hosts     []string      `config:"hosts"`
	Period    time.Duration `config:"period"`
	DebugMode bool          `config:"debug"`
	UserName  string        `config:"username"`
	Password  string        `config:"password"`
}

// func defaultConfig() *config {
// 	return &config{
// 		Hosts:     []string{"localhost"},
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
	logger   *logp.Logger
	counter  int
	debug    bool
	username string
	password string
}

// type MetricSet struct {
// 	mb.BaseMetricSet
// 	counter int
// }

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The emcunity storage_health metricset is beta.")

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
		username:      config.UserName,
		password:      config.Password,
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
	hostInfo.username = m.username
	hostInfo.password = m.password

	//metricsData := make(map[Serial]*UnityData)
	//map[Serial]*UnityData

	var metricData UnityData

	systemData, err := GetSystemMetrics(hostInfo)
	fmt.Printf("\n\nSystem Response: %+v", systemData)
	if err != nil {
		m.logger.Errorf("GetSystemMetrics failed; %v", err)
	} else {
		metricData.system = systemData
	}

	poolData, _ := GetPoolMetrics(hostInfo)
	fmt.Printf("\n\nPool Response: %+v", poolData)
	if err != nil {
		m.logger.Errorf("GetPoolMetrics failed; %v", err)
	} else {
		metricData.pool = poolData
	}

	poolUnitData, _ := GetPoolUnitMetrics(hostInfo)
	fmt.Printf("\n\nPool Unit Response: %+v", poolUnitData)
	if err != nil {
		m.logger.Errorf("GetPoolUnitMetrics failed; %v", err)
	} else {
		metricData.poolUnit = poolUnitData
	}

	lunData, _ := GetLunMetrics(hostInfo)
	fmt.Printf("\n\nLun Response: %+v", lunData)
	if err != nil {
		m.logger.Errorf("GetLunMetrics failed; %v", err)
	} else {
		metricData.lun = lunData
	}

	storageProcessorData, _ := GetStorageProcessorMetrics(hostInfo)
	fmt.Printf("\n\nStorage Processor Response: %+v", storageProcessorData)
	if err != nil {
		m.logger.Errorf("GetStorageProcessorMetrics failed; %v", err)
	} else {
		metricData.storageProcesser = storageProcessorData
	}

	storageResourceData, _ := GetStorageResourceMetrics(hostInfo)
	fmt.Printf("\n\nStorage Resource Response: %+v", storageResourceData)
	if err != nil {
		m.logger.Errorf("GetStorageResourceMetrics failed; %v", err)
	} else {
		metricData.storageResource = storageResourceData
	}

	storageTierData, _ := GetStorageTierMetrics(hostInfo)
	fmt.Printf("\n\nStorage Tier Response: %+v", storageTierData)
	if err != nil {
		m.logger.Errorf("GetStorageTierMetrics failed; %v", err)
	} else {
		metricData.storageTier = storageTierData
	}

	licenseData, _ := GetLicenseMetrics(hostInfo)
	fmt.Printf("\n\nLicense Response: %+v", licenseData)
	if err != nil {
		m.logger.Errorf("GetStorageTierMetrics failed; %v", err)
	} else {
		metricData.license = licenseData
	}

	ethernetPortData, _ := GetEthernetPortMetrics(hostInfo)
	fmt.Printf("\n\nEthernet Port Response: %+v", ethernetPortData)
	if err != nil {
		m.logger.Errorf("GetEthernetPortMetrics failed; %v", err)
	} else {
		metricData.ethernetPort = ethernetPortData
	}

	fileInterfaceData, _ := GetFileInterfaceMetrics(hostInfo)
	fmt.Printf("\n\nFile Interface Response: %+v", fileInterfaceData)
	if err != nil {
		m.logger.Errorf("GetFileInterfaceMetrics failed; %v", err)
	} else {
		metricData.fileInterface = fileInterfaceData
	}

	remoteSystemData, _ := GetRemoteSystemMetrics(hostInfo)
	fmt.Printf("\n\nRemote System Response: %+v", remoteSystemData)
	if err != nil {
		m.logger.Errorf("GetRemoteSystemMetrics failed; %v", err)
	} else {
		metricData.remoteSystem = remoteSystemData
	}

	reportMetrics(reporter, hostInfo.baseurl, metricData)

	return nil
}

type Connection struct {
	baseurl  string
	username string
	password string
}
