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

	powerSupplyData, err := GetPowerSupplyMetrics(hostInfo)
	fmt.Printf("\n\nPowerSupply Response: %+v", powerSupplyData)
	if err != nil {
		m.logger.Warnf("GetPowerSupplyMetrics failed; %v ", err)
	} else {
		metricData.powerSupply = powerSupplyData
	}

	fanData, err := GetFanMetrics(hostInfo)
	fmt.Printf("\n\nFan Response: %+v", fanData)
	if err != nil {
		m.logger.Warnf("GetFanMetrics failed; %v ", err)
	} else {
		metricData.fan = fanData
	}

	daeData, err := GetDaeMetrics(hostInfo)
	fmt.Printf("\n\nDAE (Disk Array Enclosure) Response: %+v", daeData)
	if err != nil {
		m.logger.Warnf("GetDaeMetrics (DiskArrayEnclosure) failed; %v ", err)
	} else {
		metricData.dae = daeData
	}

	memoryModuleData, err := GetMemoryModuleMetrics(hostInfo)
	fmt.Printf("\n\nMemory Module Response: %+v", memoryModuleData)
	if err != nil {
		m.logger.Warnf("GetMemoryModuleMetrics failed; %v ", err)
	} else {
		metricData.memoryModule = memoryModuleData
	}

	batteryData, err := GetBatteryMetrics(hostInfo)
	fmt.Printf("\n\nBattery Response: %+v", batteryData)
	if err != nil {
		m.logger.Warnf("GetBatteryMetrics failed; %v ", err)
	} else {
		metricData.battery = batteryData
	}

	ssdData, err := GetSsdMetrics(hostInfo)
	fmt.Printf("\n\nSsd Response: %+v", ssdData)
	if err != nil {
		m.logger.Warnf("GetSsdMetrics failed; %v ", err)
	} else {
		metricData.ssd = ssdData
	}

	raidGroupData, err := GetRaidGroupMetrics(hostInfo)
	fmt.Printf("\n\nRaidGroup Response: %+v", raidGroupData)
	if err != nil {
		m.logger.Warnf("GetRaidGroupMetrics failed; %v ", err)
	} else {
		metricData.raidGroup = raidGroupData
	}

	treeQuotaData, err := GetTreeQuotaMetrics(hostInfo)
	fmt.Printf("\n\nTreeQuota Response: %+v", treeQuotaData)
	if err != nil {
		m.logger.Warnf("GetTreeQuotaMetrics failed; %v ", err)
	} else {
		metricData.treeQuota = treeQuotaData
	}

	diskGroupData, err := GetDiskGroupMetrics(hostInfo)
	fmt.Printf("\n\nDiskGroup Response: %+v", diskGroupData)
	if err != nil {
		m.logger.Warnf("GetDiskGroupMetrics failed; %v ", err)
	} else {
		metricData.diskGroup = diskGroupData
	}

	cifsServerData, err := GetCifsServerMetrics(hostInfo)
	fmt.Printf("\n\nCifsServer Response: %+v", cifsServerData)
	if err != nil {
		m.logger.Warnf("GetCifsServerMetrics failed; %v ", err)
	} else {
		metricData.cifsServer = cifsServerData
	}

	fastCacheData, err := GetFastCacheMetrics(hostInfo)
	fmt.Printf("\n\nFastCache Response: %+v", fastCacheData)
	if err != nil {
		m.logger.Warnf("GetFastCacheMetrics failed; %v ", err)
	} else {
		metricData.fastCache = fastCacheData
	}

	fastVPData, err := GetFastVPMetrics(hostInfo)
	fmt.Printf("\n\nFastVP Response: %+v", fastVPData)
	if err != nil {
		m.logger.Warnf("GetFastVPMetrics failed; %v ", err)
	} else {
		metricData.fastVP = fastVPData
	}

	fcPortData, err := GetFcPortMetrics(hostInfo)
	fmt.Printf("\n\nfcPort Response: %+v", fcPortData)
	if err != nil {
		m.logger.Warnf("GetFcPortPMetrics failed; %v ", err)
	} else {
		metricData.fcPort = fcPortData
	}

	hostContainerData, err := GetHostContainerMetrics(hostInfo)
	fmt.Printf("\n\nHostContainer Response: %+v", hostContainerData)
	if err != nil {
		m.logger.Warnf("GetHostContainerMetrics failed; %v ", err)
	} else {
		metricData.hostContainer = hostContainerData
	}

	hostInitiatorData, err := GetHostInitiatorMetrics(hostInfo)
	fmt.Printf("\n\nHostInitiator Response: %+v", hostInitiatorData)
	if err != nil {
		m.logger.Warnf("GetHostInitiatorMetrics failed; %v ", err)
	} else {
		metricData.hostInitiator = hostInitiatorData
	}

	hostData, err := GetHostMetrics(hostInfo)
	fmt.Printf("\n\nHost Response: %+v", hostData)
	if err != nil {
		m.logger.Warnf("GetHostMetrics failed; %v ", err)
	} else {
		metricData.host = hostData
	}

	ioModuleData, err := GetIoModuleMetrics(hostInfo)
	fmt.Printf("\n\nIO Module Response: %+v", ioModuleData)
	if err != nil {
		m.logger.Warnf("GetIoModuleMetrics failed; %v ", err)
	} else {
		metricData.ioModule = ioModuleData
	}

	lccData, err := GetLccMetrics(hostInfo)
	fmt.Printf("\n\nLcc (LinkedControlCards) Response: %+v", lccData)
	if err != nil {
		m.logger.Warnf("GetLccMetrics failed; %v ", err)
	} else {
		metricData.lcc = lccData
	}

	nasServerData, err := GetNasServerMetrics(hostInfo)
	fmt.Printf("\n\nNasServer Response: %+v", nasServerData)
	if err != nil {
		m.logger.Warnf("GetNasServerMetrics failed; %v ", err)
	} else {
		metricData.nasServer = nasServerData
	}

	reportMetrics(reporter, hostInfo.baseurl, metricData)

	return nil
}

type Connection struct {
	baseurl  string
	username string
	password string
}
