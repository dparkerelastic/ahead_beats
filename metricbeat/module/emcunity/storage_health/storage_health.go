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
		m.logger.Warnf("GetSystemMetrics failed; %v", err)
	} else {
		metricData.system = systemData
	}

	poolData, err := GetPoolMetrics(hostInfo)
	fmt.Printf("\n\nPool Response: %+v", poolData)
	if err != nil {
		m.logger.Warnf("GetPoolMetrics failed; %v", err)
	} else {
		metricData.pool = poolData
	}

	poolUnitData, err := GetPoolUnitMetrics(hostInfo)
	fmt.Printf("\n\nPool Unit Response: %+v", poolUnitData)
	if err != nil {
		m.logger.Warnf("GetPoolUnitMetrics failed; %v", err)
	} else {
		metricData.poolUnit = poolUnitData
	}

	lunData, err := GetLunMetrics(hostInfo)
	fmt.Printf("\n\nLun Response: %+v", lunData)
	if err != nil {
		m.logger.Warnf("GetLunMetrics failed; %v", err)
	} else {
		metricData.lun = lunData
	}

	storageProcessorData, err := GetStorageProcessorMetrics(hostInfo)
	fmt.Printf("\n\nStorage Processor Response: %+v", storageProcessorData)
	if err != nil {
		m.logger.Warnf("GetStorageProcessorMetrics failed; %v", err)
	} else {
		metricData.storageProcesser = storageProcessorData
	}

	storageResourceData, err := GetStorageResourceMetrics(hostInfo)
	fmt.Printf("\n\nStorage Resource Response: %+v", storageResourceData)
	if err != nil {
		m.logger.Warnf("GetStorageResourceMetrics failed; %v", err)
	} else {
		metricData.storageResource = storageResourceData
	}

	storageTierData, err := GetStorageTierMetrics(hostInfo)
	fmt.Printf("\n\nStorage Tier Response: %+v", storageTierData)
	if err != nil {
		m.logger.Warnf("GetStorageTierMetrics failed; %v", err)
	} else {
		metricData.storageTier = storageTierData
	}

	licenseData, err := GetLicenseMetrics(hostInfo)
	fmt.Printf("\n\nLicense Response: %+v", licenseData)
	if err != nil {
		m.logger.Warnf("GetStorageTierMetrics failed; %v", err)
	} else {
		metricData.license = licenseData
	}

	ethernetPortData, err := GetEthernetPortMetrics(hostInfo)
	fmt.Printf("\n\nEthernet Port Response: %+v", ethernetPortData)
	if err != nil {
		m.logger.Warnf("GetEthernetPortMetrics failed; %v", err)
	} else {
		metricData.ethernetPort = ethernetPortData
	}

	fileInterfaceData, err := GetFileInterfaceMetrics(hostInfo)
	fmt.Printf("\n\nFile Interface Response: %+v", fileInterfaceData)
	if err != nil {
		m.logger.Warnf("GetFileInterfaceMetrics failed; %v", err)
	} else {
		metricData.fileInterface = fileInterfaceData
	}

	remoteSystemData, err := GetRemoteSystemMetrics(hostInfo)
	fmt.Printf("\n\nRemote System Response: %+v", remoteSystemData)
	if err != nil {
		m.logger.Warnf("GetRemoteSystemMetrics failed; %v", err)
	} else {
		metricData.remoteSystem = remoteSystemData
	}

	diskData, err := GetDiskMetrics(hostInfo)
	fmt.Printf("\n\nDisk Response: %+v", diskData)
	if err != nil {

		m.logger.Warnf("GetDiskMetrics failed; %v", err)
	} else {
		metricData.disk = diskData
	}

	datastoreData, err := GetDatastoreMetrics(hostInfo)
	fmt.Printf("\n\nDatastore Response: %+v", datastoreData)
	if err != nil {
		m.logger.Warnf("GetDatastoreMetrics failed; %v", err)
	} else {
		metricData.datastore = datastoreData
	}

	filesystemData, err := GetFilesystemMetrics(hostInfo)
	fmt.Printf("\n\nFilesystem Response: %+v", filesystemData)
	if err != nil {
		m.logger.Warnf("GetFilesystemMetrics failed; %v", err)
	} else {
		metricData.filesystem = filesystemData
	}

	snapData, err := GetSnapMetrics(hostInfo)
	fmt.Printf("\n\nSnap Response: %+v", snapData)
	if err != nil {
		m.logger.Warnf("GetSnapMetrics failed; %v", err)
	} else {
		metricData.snap = snapData
	}

	sasPortData, err := GetSasPortMetrics(hostInfo)
	fmt.Printf("\n\nSasPort Response: %+v", sasPortData)
	if err != nil {
		m.logger.Warnf("GetSasPortMetrics failed; %v ", err)
	} else {
		metricData.sasPort = sasPortData
	}

	reportMetrics(reporter, hostInfo.baseurl, metricData)

	return nil
}

type Connection struct {
	baseurl  string
	username string
	password string
}
