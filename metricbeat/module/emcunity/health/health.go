package health

import (
	"fmt"
	"strconv"
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
	mb.Registry.MustAddMetricSet("emcunity", "health", New)
}

type config struct {
	Hosts     []string      `config:"hosts"`
	Period    time.Duration `config:"period"`
	DebugMode bool          `config:"debug"`
	UserName  string        `config:"username"`
	Password  string        `config:"password"`
	PerPage   int           `config:"per_page"`
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
	logger   *logp.Logger
	counter  int
	debug    bool
	username string
	password string
	perPage  int
}

// type MetricSet struct {
// 	mb.BaseMetricSet
// 	counter int
// }

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The emcunity health metricset is beta.")

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
		perPage:       config.PerPage,
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

	var resources_per_page int
	if m.perPage > 0 {
		resources_per_page = m.perPage
	} else {
		resources_per_page = 100
	}

	var metricData UnityData

	systemData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+System_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.system)
	//fmt.Printf("\n\nSystem Response: %+v", systemData)
	if err != nil {
		m.logger.Warnf("GetSystemMetrics failed; %v", err)
	} else {
		metricData.system = systemData.(System_JSON)
		metricData.system.Message = message

	}

	poolData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Pool_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.pool)
	//fmt.Printf("\n\nPool Response: %+v", poolData)
	if err != nil {
		m.logger.Warnf("GetPoolMetrics failed; %v", err)
	} else {
		metricData.pool = poolData.(Pool_JSON)
		metricData.pool.Message = message
	}

	poolUnitData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+PoolUnit_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.poolUnit)
	//fmt.Printf("\n\nPool Unit Response: %+v", poolUnitData)
	if err != nil {
		m.logger.Warnf("GetPoolUnitMetrics failed; %v", err)
	} else {
		metricData.poolUnit = poolUnitData.(PoolUnit_JSON)
		metricData.poolUnit.Message = message
	}

	lunData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Lun_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.lun)
	//fmt.Printf("\n\nLun Response: %+v", lunData)
	if err != nil {
		m.logger.Warnf("GetLunMetrics failed; %v", err)
	} else {
		metricData.lun = lunData.(Lun_JSON)
		metricData.lun.Message = message
	}

	storageProcessorData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+StorageProcessor_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.storageProcesser)
	//fmt.Printf("\n\nStorage Processor Response: %+v", storageProcessorData)
	if err != nil {
		m.logger.Warnf("GetStorageProcessorMetrics failed; %v", err)
	} else {
		metricData.storageProcesser = storageProcessorData.(StorageProcessor_JSON)
		metricData.storageProcesser.Message = message
	}

	storageResourceData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+StorageResource_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.storageResource)
	//fmt.Printf("\n\nStorage Resource Response: %+v", storageResourceData)
	if err != nil {
		m.logger.Warnf("GetStorageResourceMetrics failed; %v", err)
	} else {
		metricData.storageResource = storageResourceData.(StorageResource_JSON)
		metricData.storageResource.Message = message
	}

	storageTierData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+StorageTier_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.storageTier)
	//fmt.Printf("\n\nStorage Tier Response: %+v", storageTierData)
	if err != nil {
		m.logger.Warnf("GetStorageTierMetrics failed; %v", err)
	} else {
		metricData.storageTier = storageTierData.(StorageTier_JSON)
		metricData.storageTier.Message = message
	}

	licenseData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+License_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.license)
	//fmt.Printf("\n\nLicense Response: %+v", licenseData)
	if err != nil {
		m.logger.Warnf("GetStorageTierMetrics failed; %v", err)
	} else {
		metricData.license = licenseData.(License_JSON)
		metricData.license.Message = message
	}

	ethernetPortData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+EthernetPort_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.ethernetPort)
	//fmt.Printf("\n\nEthernet Port Response: %+v", ethernetPortData)
	if err != nil {
		m.logger.Warnf("GetEthernetPortMetrics failed; %v", err)
	} else {
		metricData.ethernetPort = ethernetPortData.(BasicEMCUnity_JSON)
		metricData.ethernetPort.Message = message
	}

	fileInterfaceData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+FileInterface_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.fileInterface)
	//fmt.Printf("\n\nFile Interface Response: %+v", fileInterfaceData)
	if err != nil {
		m.logger.Warnf("GetFileInterfaceMetrics failed; %v", err)
	} else {
		metricData.fileInterface = fileInterfaceData.(BasicEMCUnity_JSON)
		metricData.fileInterface.Message = message
	}

	remoteSystemData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+RemoteSystem_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.remoteSystem)
	//fmt.Printf("\n\nRemote System Response: %+v", remoteSystemData)
	if err != nil {
		m.logger.Warnf("GetRemoteSystemMetrics failed; %v", err)
	} else {
		metricData.remoteSystem = remoteSystemData.(BasicEMCUnity_JSON)
		metricData.remoteSystem.Message = message
	}

	diskData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Disk_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.disk)
	//fmt.Printf("\n\nDisk Response: %+v", diskData)
	if err != nil {

		m.logger.Warnf("GetDiskMetrics failed; %v", err)
	} else {
		metricData.disk = diskData.(Disk_JSON)
		metricData.disk.Message = message
	}

	datastoreData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+DataStore_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.datastore)
	//fmt.Printf("\n\nDatastore Response: %+v", datastoreData)
	if err != nil {
		m.logger.Warnf("GetDatastoreMetrics failed; %v", err)
	} else {

		metricData.datastore = datastoreData.(Datastore_JSON)
		metricData.datastore.Message = message
	}

	filesystemData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Filesystem_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.filesystem)
	//fmt.Printf("\n\nFilesystem Response: %+v", filesystemData)
	if err != nil {
		m.logger.Warnf("GetFilesystemMetrics failed; %v", err)
	} else {
		metricData.filesystem = filesystemData.(FileSystem_JSON)
		metricData.filesystem.Message = message
	}

	snapData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Snap_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.snap)
	//fmt.Printf("\n\nSnap Response: %+v", snapData)
	if err != nil {
		m.logger.Warnf("GetSnapMetrics failed; %v", err)
	} else {
		metricData.snap = snapData.(Snap_JSON)
		metricData.snap.Message = message
	}

	sasPortData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+SasPort_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.sasPort)
	//fmt.Printf("\n\nSasPort Response: %+v", sasPortData)
	if err != nil {
		m.logger.Warnf("GetSasPortMetrics failed; %v ", err)
	} else {
		metricData.sasPort = sasPortData.(SasPort_JSON)
		metricData.sasPort.Message = message
	}

	powerSupplyData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+PowerSupply_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.powerSupply)
	//fmt.Printf("\n\nPowerSupply Response: %+v", powerSupplyData)
	if err != nil {
		m.logger.Warnf("GetPowerSupplyMetrics failed; %v ", err)
	} else {
		metricData.powerSupply = powerSupplyData.(BasicEMCUnity_JSON)
		metricData.powerSupply.Message = message
	}

	fanData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Fan_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.fan)
	//fmt.Printf("\n\nFan Response: %+v", fanData)
	if err != nil {
		m.logger.Warnf("GetFanMetrics failed; %v ", err)
	} else {
		metricData.fan = fanData.(BasicEMCUnity_JSON)
		metricData.fan.Message = message
	}

	daeData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Dae_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.dae)
	//fmt.Printf("\n\nDAE (Disk Array Enclosure) Response: %+v", daeData)
	if err != nil {
		m.logger.Warnf("GetDaeMetrics (DiskArrayEnclosure) failed; %v ", err)
	} else {
		metricData.dae = daeData.(Dae_JSON)
		metricData.dae.Message = message
	}

	memoryModuleData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+MemoryModule_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.memoryModule)
	//fmt.Printf("\n\nMemory Module Response: %+v", memoryModuleData)
	if err != nil {
		m.logger.Warnf("GetMemoryModuleMetrics failed; %v ", err)
	} else {
		metricData.memoryModule = memoryModuleData.(BasicEMCUnity_JSON)
		metricData.memoryModule.Message = message
	}

	batteryData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Battery_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.battery)
	//fmt.Printf("\n\nBattery Response: %+v", batteryData)
	if err != nil {
		m.logger.Warnf("GetBatteryMetrics failed; %v ", err)
	} else {
		metricData.battery = batteryData.(BasicEMCUnity_JSON)
		metricData.battery.Message = message
	}

	ssdData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Ssd_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.ssd)
	//fmt.Printf("\n\nSsd Response: %+v", ssdData)
	if err != nil {
		m.logger.Warnf("GetSsdMetrics failed; %v ", err)
	} else {
		metricData.ssd = ssdData.(BasicEMCUnity_JSON)
		metricData.ssd.Message = message
	}

	raidGroupData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+RaidGroup_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.raidGroup)
	//fmt.Printf("\n\nRaidGroup Response: %+v", raidGroupData)
	if err != nil {
		m.logger.Warnf("GetRaidGroupMetrics failed; %v ", err)
	} else {
		metricData.raidGroup = raidGroupData.(BasicEMCUnity_JSON)
		metricData.raidGroup.Message = message
	}

	treeQuotaData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+TreeQuota_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.treeQuota)
	//fmt.Printf("\n\nTreeQuota Response: %+v", treeQuotaData)
	if err != nil {
		m.logger.Warnf("GetTreeQuotaMetrics failed; %v ", err)
	} else {
		metricData.treeQuota = treeQuotaData.(TreeQuota_JSON)
		metricData.treeQuota.Message = message
	}

	diskGroupData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+DiskGroup_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.diskGroup)
	//fmt.Printf("\n\nDiskGroup Response: %+v", diskGroupData)
	if err != nil {
		m.logger.Warnf("GetDiskGroupMetrics failed; %v ", err)
	} else {
		metricData.diskGroup = diskGroupData.(DiskGroup_JSON)
		metricData.diskGroup.Message = message
	}

	cifsServerData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+CifsServer_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.cifsServer)
	//fmt.Printf("\n\nCifsServer Response: %+v", cifsServerData)
	if err != nil {
		m.logger.Warnf("GetCifsServerMetrics failed; %v ", err)
	} else {
		metricData.cifsServer = cifsServerData.(BasicEMCUnity_JSON)
		metricData.cifsServer.Message = message
	}

	fastCacheData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+FastCache_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.fastCache)
	//fmt.Printf("\n\nFastCache Response: %+v", fastCacheData)
	if err != nil {
		m.logger.Warnf("GetFastCacheMetrics failed; %v ", err)
	} else {
		metricData.fastCache = fastCacheData.(FastCache_JSON)
		metricData.fastCache.Message = message
	}

	fastVPData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+FastVP_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.fastVP)
	//fmt.Printf("\n\nFastVP Response: %+v", fastVPData)
	if err != nil {
		m.logger.Warnf("GetFastVPMetrics failed; %v ", err)
	} else {
		metricData.fastVP = fastVPData.(FastVP_JSON)
		metricData.fastVP.Message = message
	}

	fcPortData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+FcPort_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.fcPort)
	//fmt.Printf("\n\nfcPort Response: %+v", fcPortData)
	if err != nil {
		m.logger.Warnf("GetFcPortPMetrics failed; %v ", err)
	} else {
		metricData.fcPort = fcPortData.(BasicEMCUnity_JSON)
		metricData.fcPort.Message = message
	}

	hostContainerData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+HostContainer_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.hostContainer)
	//fmt.Printf("\n\nHostContainer Response: %+v", hostContainerData)
	if err != nil {
		m.logger.Warnf("GetHostContainerMetrics failed; %v ", err)
	} else {
		metricData.hostContainer = hostContainerData.(BasicEMCUnity_JSON)
		metricData.hostContainer.Message = message
	}

	hostInitiatorData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+HostInitiator_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.hostInitiator)
	//fmt.Printf("\n\nHostInitiator Response: %+v", hostInitiatorData)
	if err != nil {
		m.logger.Warnf("GetHostInitiatorMetrics failed; %v ", err)
	} else {
		metricData.hostInitiator = hostInitiatorData.(BasicEMCUnity_JSON)
		metricData.hostInitiator.Message = message
	}

	hostData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Host_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.host)
	//fmt.Printf("\n\nHost Response: %+v", hostData)
	if err != nil {
		m.logger.Warnf("GetHostMetrics failed; %v ", err)
	} else {
		metricData.host = hostData.(BasicEMCUnity_JSON)
		metricData.host.Message = message
	}

	ioModuleData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+IoModule_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.ioModule)
	//fmt.Printf("\n\nIO Module Response: %+v", ioModuleData)
	if err != nil {
		m.logger.Warnf("GetIoModuleMetrics failed; %v ", err)
	} else {
		metricData.ioModule = ioModuleData.(BasicEMCUnity_JSON)
		metricData.ioModule.Message = message
	}

	lccData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+Lcc_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.lcc)
	//fmt.Printf("\n\nLcc (LinkedControlCards) Response: %+v", lccData)
	if err != nil {
		m.logger.Warnf("GetLccMetrics failed; %v ", err)
	} else {
		metricData.lcc = lccData.(BasicEMCUnity_JSON)
		metricData.lcc.Message = message
	}

	nasServerData, message, err := GetMetrics(hostInfo, hostInfo.baseurl+NasServer_API+"&per_page="+strconv.Itoa(resources_per_page), metricData.nasServer)
	//fmt.Printf("\n\nNasServer Response: %+v", nasServerData)
	if err != nil {
		m.logger.Warnf("GetNasServerMetrics failed; %v ", err)
	} else {
		metricData.nasServer = nasServerData.(BasicEMCUnity_JSON)
		metricData.nasServer.Message = message
	}

	reportMetrics(reporter, hostInfo.baseurl, metricData, m.debug)

	return nil
}

type Connection struct {
	baseurl  string
	username string
	password string
}
