package storage_health

import (
	"encoding/json"
)

func GetSystemMetrics(hostInfo Connection) (System_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+System_API)

	var jsonInfo System_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}

func GetPoolMetrics(hostInfo Connection) (Pool_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+Pool_API)

	var jsonInfo Pool_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}

func GetPoolUnitMetrics(hostInfo Connection) (PoolUnit_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+PoolUnit_API)

	var jsonInfo PoolUnit_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}

func GetLunMetrics(hostInfo Connection) (Lun_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+Lun_API)

	var jsonInfo Lun_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}

func GetStorageProcessorMetrics(hostInfo Connection) (StorageProcessor_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+StorageProcessor_API)

	var jsonInfo StorageProcessor_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}

func GetStorageResourceMetrics(hostInfo Connection) (StorageResource_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+StorageResource_API)

	var jsonInfo StorageResource_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}
func GetStorageTierMetrics(hostInfo Connection) (StorageTier_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+StorageTier_API)

	var jsonInfo StorageTier_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}

func GetLicenseMetrics(hostInfo Connection) (License_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+License_API)

	var jsonInfo License_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil
}

// EthernetPort, FileInterface, RemoteSystem
func GetEthernetPortMetrics(hostInfo Connection) (BasicEMCUnity_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+EthernetPort_API)

	var jsonInfo BasicEMCUnity_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}

func GetFileInterfaceMetrics(hostInfo Connection) (BasicEMCUnity_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+FileInterface_API)

	var jsonInfo BasicEMCUnity_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}

func GetRemoteSystemMetrics(hostInfo Connection) (BasicEMCUnity_JSON, error) {

	responseData, err := GetInstanceData(hostInfo, hostInfo.baseurl+RemoteSystem_API)

	var jsonInfo BasicEMCUnity_JSON
	if err != nil {
		return jsonInfo, err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, err
	}

	return jsonInfo, nil

}
