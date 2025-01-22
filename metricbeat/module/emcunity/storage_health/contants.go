package storage_health

// url := "https://10.200.0.8:443/api/types/system/instances?fields=name,model,serialNumber,internalModel,platform,macAddress"
//const BasicSystemInfo_API = "/api/types/basicSystemInfo/instances?fields=id,name,model,softwareVersion,apiVersion"

const System_API = "/api/types/system/instances?fields=name,model,health,serialNumber,internalModel,platform,macAddress&compact=true"
const Pool_API = "/api/types/pool/instances?fields=id,name,sizeFree,sizeTotal,health,description&compact=true"
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
