package storage_health

// url := "https://10.200.0.8:443/api/types/system/instances?fields=name,model,serialNumber,internalModel,platform,macAddress"
//const BasicSystemInfo_API = "/api/types/basicSystemInfo/instances?fields=id,name,model,softwareVersion,apiVersion"

const System_API = "/api/types/system/instances?fields=name,model,health,serialNumber,internalModel,platform,macAddress&compact=true"
const Pool_API = "/api/types/pool/instances?fields=id,name,sizeFree,sizeTotal,health,description,harvestState,metadataSizeSubscribed,metadataSizeUsed,rebalanceProgress,sizeSubscribed,sizeUsed,snapSizeSubscribed,snapSizeUsed&compact=true"
const PoolUnit_API = "/api/types/poolUnit/instances?fields=name,health,sizeTotal&compact=true"
const Lun_API = "/api/types/lun/instances?fields=health,sizeTotal,sizeUsed,sizeAllocated,snapsSize,snapsSizeAllocated,metadataSize,metadataSizeAllocated&compact=true"
const StorageProcessor_API = "/api/types/storageProcessor/instances?fields=name,id,model,health&compact=true"
const StorageResource_API = "/api/types/storageResource/instances?fields=health,sizeTotal,sizeUsed,sizeAllocated,metadataSize,metadataSizeAllocated,snapsSizeTotal,snapsSizeAllocated,snapCount&compact=true"
const StorageTier_API = "/api/types/storageTier/instances?fields=disksTotal,disksUnused,virtualDisksTotal,virtualDisksUnused,sizeTotal,sizeFree&compact=true"
const License_API = "/api/types/license/instances?fields=id,name,isValid,expires,isPermanent,isInstalled,feature&compact=true"

// EthernetPort, FileInterface, RemoteSystem
const EthernetPort_API = "/api/types/ethernetPort/instances?fields=health&compact=true"
const FileInterface_API = "/api/types/fileInterface/instances?fields=health&compact=true"
const RemoteSystem_API = "/api/types/remoteSystem/instances?fields=health&compact=true"

const DataStore_API = "/api/types/datastore/instances?fields=sizeTotal,sizeUsed&compact=true"
const Disk_API = "/api/types/disk/instances?fields=health,size,rawSize,vendorSize&compact=true"
const Filesystem_API = "/api/types/filesystem/instances?fields=name,health,metadataSize,metadataSizeAllocated,sizeAllocated,sizeTotal,sizeUsed,snapCount,snapsSize,snapsSizeAllocated&compact=true"
const Snap_API = "/api/types/snap/instances?fields=name,size,state,expirationTime,creationTime&compact=true"
const SasPort_API = "/api/types/sasPort/instances?fields=name,needsReplacement,health&compact=true"

// Starting week of 1/27
const PowerSupply_API = "/api/types/powerSupply/instances?fields=name,needsReplacement,health&compact=true"
const Fan_API = "/api/types/fan/instances?fields=name,needsReplacement,health&compact=true"

// DAE = DiskArrayEnclosure
const Dae_API = "/api/types/dae/instances?fields=health,currentPower,avgPower,maxPower,currentTemperature,avgTemperature,maxTemperature&compact=true"
const MemoryModule_API = "/api/types/memoryModule/instances?fields=name,needsReplacement,health,size&compact=true"
const Battery_API = "/api/types/battery/instances?fields=name,needsReplacement,health&compact=true"
const Ssd_API = "/api/types/ssd/instances?fields=name,needsReplacement,health&compact=true"
const RaidGroup_API = "/api/types/raidGroup/instances?fields=name,health,sizeTotal&compact=true"
const TreeQuota_API = "/api/types/treeQuota/instances?fields=softLimit,hardLimit,sizeUsed,state&compact=true"

const DiskGroup_API = "/api/types/diskGroup/instances?fields=name,advertisedSize,diskSize,hotSparePolicyStatus,minHotSpareCandidates,rpm,speed,totalDisks,unconfiguredDisks&compact=true"

const CifsServer_API = "/api/types/cifsServer/instances?fields=name,health&compact=true"
const FastCache_API = "/api/types/fastCache/instances?fields=id,name,sizeFree,sizeTotal,health,numberOfDisks&compact=true"
const FastVP_API = "/api/types/fastVP/instances?fields=id,isScheduleEnabled,relocationRate,sizeMovingUp,sizeMovingDown,sizeMovingWithin,status,relocationDurationEstimate&compact=true"

// FiberChannelPorts (see below fcPort)
const FcPort_API = "/api/types/fcPort/instances?fields=name,health&compact=true"
const HostContainer_API = "/api/types/hostContainer/instances?fields=name,health&compact=true"
const HostInitiator_API = "/api/types/hostInitiator/instances?fields=health&compact=true"
const Host_API = "/api/types/host/instances?fields=name,health&compact=true"
const IoModule_API = "/api/types/ioModule/instances?fields=name,needsReplacement,health&compact=true"

// LinkControlCards (see below lcc)
const Lcc_API = "/api/types/lcc/instances?fields=name,needsReplacement,health&compact=true"
const NasServer_API = "/api/types/nasServer/instances?fields=name,health&compact=true"
