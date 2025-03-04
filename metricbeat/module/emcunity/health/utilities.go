package health

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"strconv"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
// 	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
// 	return nil
// }

// basicAuth converts the given username & password to Base64 encoded string.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func GetInstanceData(hostInfo Connection, url string) ([]byte, error) {

	req, _ := http.NewRequest("GET", url, nil)
	// add authorization header to the req
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-EMC-REST-CLIENT", "true")
	req.Header.Add("Authorization", "Basic "+basicAuth(hostInfo.username, hostInfo.password))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: nil})
	client := &http.Client{
		Transport: tr,
	}

	client.Jar = cookieJar
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var errorJson ErrorResponse
	err = json.Unmarshal(responseData, &errorJson)
	if err != nil {
		return nil, err
	} else if errorJson.Error.HTTPStatusCode != 0 {

		var errorMessage string = "Error: HTTPStatusCode: " + strconv.Itoa(errorJson.Error.HTTPStatusCode) + " Message: "
		for _, message := range errorJson.Error.Messages {
			errorMessage += message.EnUS + " "
		}
		return nil, errors.New(errorMessage)

	}

	return responseData, nil

}

func reportMetricsForUnityStorage(reporter mb.ReporterV2, baseURL string, metrics ...[]mapstr.M) {
	for _, metricSlice := range metrics {
		for _, metric := range metricSlice {
			event := mb.Event{ModuleFields: mapstr.M{"base_url": baseURL}}
			if ts, ok := metric["@timestamp"]; ok {
				t, err := time.Parse(time.RFC3339, ts.(string))
				if err == nil {
					// if the timestamp parsing fails, we just fall back to the event time
					// (and leave the additional timestamp in the event for posterity)
					event.Timestamp = t
					delete(metric, "@timestamp")
				}
			}

			for k, v := range metric {
				if !isEmpty(v) {
					//fmt.Println("k =" + k + " v=" + string(v))
					event.ModuleFields.Put(k, v)
				}
			}

			reporter.Event(event)
		}
	}
}

func isEmpty(value interface{}) bool {
	// we make use of the fact that all the dashboard API responses utilize
	// pointers for non-string types to filter out empty values from metric events.

	if value == nil {
		return true
	}

	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		return reflect.ValueOf(value).IsNil()
	}

	if t.Kind() == reflect.Slice || t.Kind() == reflect.String {
		return reflect.ValueOf(value).Len() == 0
	}

	return false
}

func GetMetrics[T any](hostInfo Connection, url string, jsonInfo T) (any, string, error) {
	responseData, err := GetInstanceData(hostInfo, url)

	if err != nil {
		return jsonInfo, "", err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		return jsonInfo, "", err
	}

	return jsonInfo, string(responseData), nil

}

func reportMetrics(reporter mb.ReporterV2, baseURL string, data UnityData, debug bool) {
	metrics := []mapstr.M{}

	for _, systemData := range data.system.Entries {
		metric := mapstr.M{}
		metric["health.system.id"] = systemData.Content.ID
		metric["health.system.name"] = systemData.Content.Name
		metric["health.system.model"] = systemData.Content.Model
		metric["health.system.serial_number"] = systemData.Content.SerialNumber
		metric["health.system.internal_model"] = systemData.Content.InternalModel
		metric["health.system.platform"] = systemData.Content.Platform
		metric["health.system.macaddress"] = systemData.Content.MacAddress
		//metric["health.system.system_uuid"] = systemData.Content.SystemUUID
		metric["health.system.health.value"] = systemData.Content.Health.Value
		metric["health.system.health.description.ids"] = systemData.Content.Health.DescriptionIds
		metric["health.system.health.descriptions"] = systemData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.system.Message
		}

		metrics = append(metrics, metric)
	}

	for _, poolData := range data.pool.Entries {
		metric := mapstr.M{}
		metric["health.pool.id"] = poolData.Content.ID
		metric["health.pool.name"] = poolData.Content.Name
		metric["health.pool.description"] = poolData.Content.Description
		metric["health.pool.size.free"] = poolData.Content.SizeFree
		metric["health.pool.size.used"] = poolData.Content.SizeUsed
		metric["health.pool.size.total"] = poolData.Content.SizeTotal
		metric["health.pool.size.subscribed"] = poolData.Content.SizeSubscribed
		metric["health.pool.health.value"] = poolData.Content.Health.Value
		metric["health.pool.health.description.ids"] = poolData.Content.Health.DescriptionIds
		metric["health.pool.health.descriptions"] = poolData.Content.Health.Descriptions
		metric["health.pool.harvest.state"] = poolData.Content.HarvestState
		metric["health.pool.metadata.size.subscribed"] = poolData.Content.MetadataSizeSubscribed
		metric["health.pool.metadata.size.used"] = poolData.Content.MetadataSizeUsed
		metric["health.pool.snap.size.subscribed"] = poolData.Content.SnapSizeSubscribed
		metric["health.pool.snap.size.used"] = poolData.Content.SnapSizeUsed
		metric["health.pool.rebalance.progress"] = poolData.Content.RebalanceProgress

		//calculation - if divide by zero is not going to occur
		if poolData.Content.SizeTotal > 0 {
			metric["health.pool.size.percent.used"] = int((float64(poolData.Content.SizeUsed) / float64(poolData.Content.SizeTotal)) * float64(100))
		}

		if debug {
			metric["health.message"] = data.pool.Message
		}

		metrics = append(metrics, metric)
	}

	for _, poolUnitData := range data.poolUnit.Entries {
		metric := mapstr.M{}
		metric["health.pool.unit.id"] = poolUnitData.Content.ID
		metric["health.pool.unit.name"] = poolUnitData.Content.Name
		metric["health.pool.unit.size.total"] = poolUnitData.Content.SizeTotal
		metric["health.pool.unit.health.value"] = poolUnitData.Content.Health.Value
		metric["health.pool.unit.health.description.ids"] = poolUnitData.Content.Health.DescriptionIds
		metric["health.pool.unit.health.descriptions"] = poolUnitData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.poolUnit.Message
		}

		metrics = append(metrics, metric)
	}

	for _, lunData := range data.lun.Entries {
		metric := mapstr.M{}
		metric["health.lun.id"] = lunData.Content.ID
		metric["health.lun.size.total"] = lunData.Content.SizeTotal
		metric["health.lun.size.allocated"] = lunData.Content.SizeAllocated
		metric["health.lun.snap.size.total"] = lunData.Content.SnapsSize
		metric["health.lun.snap.size.allocated"] = lunData.Content.SnapsSizeAllocated
		metric["health.lun.metadata.size.total"] = lunData.Content.MetadataSize
		metric["health.lun.metadata.size.allocated"] = lunData.Content.MetadataSizeAllocated
		metric["health.lun.health.value"] = lunData.Content.Health.Value
		metric["health.lun.health.description.ids"] = lunData.Content.Health.DescriptionIds
		metric["health.lun.health.descriptions"] = lunData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.lun.Message
		}

		metrics = append(metrics, metric)
	}

	for _, storageProcessorData := range data.storageProcesser.Entries {
		metric := mapstr.M{}
		metric["health.storage.processor.id"] = storageProcessorData.Content.ID
		metric["health.storage.processor.model"] = storageProcessorData.Content.Model
		metric["health.storage.processor.name"] = storageProcessorData.Content.Name
		metric["health.storage.processor.health.value"] = storageProcessorData.Content.Health.Value
		metric["health.storage.processor.health.description.ids"] = storageProcessorData.Content.Health.DescriptionIds
		metric["health.storage.processor.health.descriptions"] = storageProcessorData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.storageProcesser.Message
		}

		metrics = append(metrics, metric)
	}

	for _, storageResourceData := range data.storageResource.Entries {
		metric := mapstr.M{}
		metric["health.storage.resource.id"] = storageResourceData.Content.ID
		metric["health.storage.resource.health.value"] = storageResourceData.Content.Health.Value
		metric["health.storage.resource.health.description.ids"] = storageResourceData.Content.Health.DescriptionIds
		metric["health.storage.resource.health.descriptions"] = storageResourceData.Content.Health.Descriptions
		metric["health.storage.resource.size.total"] = storageResourceData.Content.SizeTotal
		metric["health.storage.resource.size.allocated"] = storageResourceData.Content.SizeAllocated
		metric["health.storage.resource.snap.count"] = storageResourceData.Content.SnapCount
		metric["health.storage.resource.snaps.size.total"] = storageResourceData.Content.SnapsSizeTotal
		metric["health.storage.resource.snaps.size.allocated"] = storageResourceData.Content.SnapsSizeAllocated
		metric["health.storage.resource.metadata.size.total"] = storageResourceData.Content.MetadataSize
		metric["health.storage.resource.metadata.size.allocated"] = storageResourceData.Content.MetadataSizeAllocated

		if debug {
			metric["health.message"] = data.storageResource.Message
		}

		metrics = append(metrics, metric)
	}

	for _, storageTierData := range data.storageTier.Entries {
		metric := mapstr.M{}
		metric["health.storage.tier.id"] = storageTierData.Content.ID
		metric["health.storage.tier.size.total"] = storageTierData.Content.SizeTotal
		metric["health.storage.tier.size.free"] = storageTierData.Content.SizeFree
		metric["health.storage.tier.disk.total"] = storageTierData.Content.DisksTotal
		metric["health.storage.tier.disk.unused"] = storageTierData.Content.DisksUnused
		metric["health.storage.tier.virtual.disk.total"] = storageTierData.Content.VirtualDisksTotal
		metric["health.storage.tier.virtual.disk.unused"] = storageTierData.Content.VirtualDisksUnused

		if storageTierData.Content.DisksTotal > 0 {
			metric["health.storage.tier.disk.used"] = storageTierData.Content.DisksTotal - storageTierData.Content.DisksUnused
		}

		if storageTierData.Content.VirtualDisksTotal > 0 {
			metric["health.storage.tier.virtual.disk.used"] = storageTierData.Content.VirtualDisksTotal - storageTierData.Content.VirtualDisksUnused
		}

		if debug {
			metric["health.message"] = data.storageTier.Message
		}

		metrics = append(metrics, metric)
	}

	for _, licenseData := range data.license.Entries {
		metric := mapstr.M{}
		metric["health.license.id"] = licenseData.Content.ID
		metric["health.license.name"] = licenseData.Content.Name
		metric["health.license.isvalid"] = licenseData.Content.IsValid
		metric["health.license.ispermanent"] = licenseData.Content.IsPermanent
		metric["health.license.isinstalled"] = licenseData.Content.IsInstalled
		metric["health.license.expires"] = licenseData.Content.Expires
		metric["health.license.feature.id"] = licenseData.Content.Feature.ID

		if debug {
			metric["health.message"] = data.license.Message
		}

		metrics = append(metrics, metric)
	}

	for _, ethernetPortData := range data.ethernetPort.Entries {
		metric := mapstr.M{}
		metric["health.ethernet.port.id"] = ethernetPortData.Content.ID
		metric["health.ethernet.port.health.value"] = ethernetPortData.Content.Health.Value
		metric["health.ethernet.port.health.description.ids"] = ethernetPortData.Content.Health.DescriptionIds
		metric["health.ethernet.port.health.descriptions"] = ethernetPortData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.ethernetPort.Message
		}

		metrics = append(metrics, metric)
	}

	for _, fileInterfaceData := range data.fileInterface.Entries {
		metric := mapstr.M{}
		metric["health.file.interface.id"] = fileInterfaceData.Content.ID
		metric["health.file.interface.health.value"] = fileInterfaceData.Content.Health.Value
		metric["health.file.interface.health.description.ids"] = fileInterfaceData.Content.Health.DescriptionIds
		metric["health.file.interface.health.descriptions"] = fileInterfaceData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.fileInterface.Message
		}

		metrics = append(metrics, metric)
	}

	for _, remoteSystemData := range data.remoteSystem.Entries {
		metric := mapstr.M{}
		metric["health.remote.system.id"] = remoteSystemData.Content.ID
		metric["health.remote.system.health.value"] = remoteSystemData.Content.Health.Value
		metric["health.remote.system.health.description.ids"] = remoteSystemData.Content.Health.DescriptionIds
		metric["health.remote.system.health.descriptions"] = remoteSystemData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.remoteSystem.Message
		}

		metrics = append(metrics, metric)
	}

	for _, diskData := range data.disk.Entries {
		metric := mapstr.M{}
		metric["health.disk.id"] = diskData.Content.ID
		metric["health.disk.name"] = diskData.Content.Name
		metric["health.disk.raw.size"] = diskData.Content.RawSize
		metric["health.disk.size"] = diskData.Content.Size
		metric["health.disk.vendor.size"] = diskData.Content.VendorSize
		metric["health.disk.health.value"] = diskData.Content.Health.Value
		metric["health.disk.health.desciption.ids"] = diskData.Content.Health.DescriptionIds
		metric["health.disk.health.desciption.descriptions"] = diskData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.disk.Message
		}

		metrics = append(metrics, metric)
	}

	for _, datastoreData := range data.datastore.Entries {
		metric := mapstr.M{}
		metric["health.datastore.id"] = datastoreData.Content.ID
		metric["health.datastore.name"] = datastoreData.Content.Name
		metric["health.datastore.size.total"] = datastoreData.Content.SizeTotal
		metric["health.datastore.size.used"] = datastoreData.Content.SizeUsed

		//calculation if divide by zero is not going to occur
		if datastoreData.Content.SizeTotal > 0 {
			metric["health.datastore.size.percent.used"] = int((float64(datastoreData.Content.SizeUsed) / float64(datastoreData.Content.SizeTotal)) * float64(100))
		}

		if debug {
			metric["health.message"] = data.datastore.Message
		}

		metrics = append(metrics, metric)
	}

	for _, filesystemData := range data.filesystem.Entries {
		metric := mapstr.M{}
		metric["health.filesystem.id"] = filesystemData.Content.ID
		metric["health.filesystem.name"] = filesystemData.Content.Name
		metric["health.filesystem.health.value"] = filesystemData.Content.Health.Value
		metric["health.filesystem.health.description.ids"] = filesystemData.Content.Health.DescriptionIds
		metric["health.filesystem.health.descriptions"] = filesystemData.Content.Health.Descriptions
		metric["health.filesystem.size.total"] = filesystemData.Content.SizeTotal
		metric["health.filesystem.size.used"] = filesystemData.Content.SizeUsed
		metric["health.filesystem.size.allocated"] = filesystemData.Content.SizeAllocated
		metric["health.filesystem.metadata.size.total"] = filesystemData.Content.MetadataSize
		metric["health.filesystem.metadata.size.allocated"] = filesystemData.Content.MetadataSizeAllocated
		metric["health.filesystem.snap.count"] = filesystemData.Content.SnapCount
		metric["health.filesystem.snaps.size.total"] = filesystemData.Content.SnapsSize
		metric["health.filesystem.snaps.size.allocated"] = filesystemData.Content.SnapsSizeAllocated

		//calculation if divide by zero is not going to occur
		if filesystemData.Content.SizeTotal > 0 {
			metric["health.filesystem.size.percent.used"] = int((float64(filesystemData.Content.SizeUsed) / float64(filesystemData.Content.SizeTotal)) * float64(100))
		}

		if debug {
			metric["health.message"] = data.filesystem.Message
		}

		metrics = append(metrics, metric)
	}

	for _, snapData := range data.snap.Entries {
		metric := mapstr.M{}
		metric["health.snap.id"] = snapData.Content.ID
		metric["health.snap.name"] = snapData.Content.Name
		metric["health.snap.size.total"] = snapData.Content.Size
		metric["health.snap.state"] = snapData.Content.State
		metric["health.snap.creation.time"] = snapData.Content.CreationTime
		metric["health.snap.expiration.time"] = snapData.Content.ExpirationTime

		if debug {
			metric["health.message"] = data.snap.Message
		}

		metrics = append(metrics, metric)
	}

	for _, sasPortData := range data.sasPort.Entries {
		metric := mapstr.M{}
		metric["health.sas.port.id"] = sasPortData.Content.ID
		metric["health.sas.port.name"] = sasPortData.Content.Name
		metric["health.sas.port.needs_replacement"] = sasPortData.Content.NeedsReplacment
		metric["health.sas.port.health.value"] = sasPortData.Content.Health.Value
		metric["health.sas.port.health.desciption.ids"] = sasPortData.Content.Health.DescriptionIds
		metric["health.sas.port.health.desciption.descriptions"] = sasPortData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.sasPort.Message
		}

		metrics = append(metrics, metric)
	}

	for _, powerSupplyData := range data.powerSupply.Entries {
		metric := mapstr.M{}
		metric["health.power.supply.id"] = powerSupplyData.Content.ID
		metric["health.power.supply.name"] = powerSupplyData.Content.Name
		metric["health.power.supply.needs_replacement"] = powerSupplyData.Content.NeedsReplacment
		metric["health.power.supply.health.value"] = powerSupplyData.Content.Health.Value
		metric["health.power.supply.health.desciption.ids"] = powerSupplyData.Content.Health.DescriptionIds
		metric["health.power.supply.health.desciption.descriptions"] = powerSupplyData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.powerSupply.Message
		}

		metrics = append(metrics, metric)
	}

	for _, fanData := range data.fan.Entries {
		metric := mapstr.M{}
		metric["health.fan.id"] = fanData.Content.ID
		metric["health.fan.name"] = fanData.Content.Name
		metric["health.fan.needs_replacement"] = fanData.Content.NeedsReplacment
		metric["health.fan.health.value"] = fanData.Content.Health.Value
		metric["health.fan.health.desciption.ids"] = fanData.Content.Health.DescriptionIds
		metric["health.fan.health.desciption.descriptions"] = fanData.Content.Health.Descriptions
		if debug {
			metric["health.message"] = data.fan.Message
		}

		metrics = append(metrics, metric)
	}

	for _, daeData := range data.dae.Entries {
		metric := mapstr.M{}
		metric["health.dae.id"] = daeData.Content.ID
		metric["health.dae.name"] = daeData.Content.Name
		metric["health.dae.current.power"] = daeData.Content.CurrentPower
		metric["health.dae.avg.power"] = daeData.Content.AvgPower
		metric["health.dae.max.power"] = daeData.Content.MaxPower
		metric["health.dae.current.temperature"] = daeData.Content.CurrentTemperature
		metric["health.dae.avg.temperature"] = daeData.Content.AvgTemperature
		metric["health.dae.max.temperature"] = daeData.Content.MaxTemperature
		metric["health.dae.needs_replacement"] = daeData.Content.NeedsReplacment
		metric["health.dae.health.value"] = daeData.Content.Health.Value
		metric["health.dae.health.desciption.ids"] = daeData.Content.Health.DescriptionIds
		metric["health.dae.health.desciption.descriptions"] = daeData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.dae.Message
		}

		metrics = append(metrics, metric)
	}

	for _, memoryModuleData := range data.memoryModule.Entries {
		metric := mapstr.M{}
		metric["health.memory.module.id"] = memoryModuleData.Content.ID
		metric["health.memory.module.name"] = memoryModuleData.Content.Name
		metric["health.memory.module.size"] = memoryModuleData.Content.Size
		metric["health.memory.module.needs_replacement"] = memoryModuleData.Content.NeedsReplacment
		metric["health.memory.module.health.value"] = memoryModuleData.Content.Health.Value
		metric["health.memory.module.health.desciption.ids"] = memoryModuleData.Content.Health.DescriptionIds
		metric["health.memory.module.health.desciption.descriptions"] = memoryModuleData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.memoryModule.Message
		}

		metrics = append(metrics, metric)
	}

	for _, batteryData := range data.battery.Entries {
		metric := mapstr.M{}
		metric["health.battery.id"] = batteryData.Content.ID
		metric["health.battery.name"] = batteryData.Content.Name
		metric["health.battery.needs_replacement"] = batteryData.Content.NeedsReplacment
		metric["health.battery.health.value"] = batteryData.Content.Health.Value
		metric["health.battery.health.desciption.ids"] = batteryData.Content.Health.DescriptionIds
		metric["health.battery.health.desciption.descriptions"] = batteryData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.battery.Message
		}

		metrics = append(metrics, metric)
	}

	for _, ssdData := range data.ssd.Entries {
		metric := mapstr.M{}
		metric["health.ssd.id"] = ssdData.Content.ID
		metric["health.ssd.name"] = ssdData.Content.Name
		metric["health.ssd.needs_replacement"] = ssdData.Content.NeedsReplacment
		metric["health.ssd.health.value"] = ssdData.Content.Health.Value
		metric["health.ssd.health.desciption.ids"] = ssdData.Content.Health.DescriptionIds
		metric["health.ssd.health.desciption.descriptions"] = ssdData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.ssd.Message
		}

		metrics = append(metrics, metric)
	}

	for _, raidGroupData := range data.raidGroup.Entries {
		metric := mapstr.M{}
		metric["health.raid.group.id"] = raidGroupData.Content.ID
		metric["health.raid.group.name"] = raidGroupData.Content.Name
		metric["health.raid.group.size.total"] = raidGroupData.Content.SizeTotal
		metric["health.raid.group.health.value"] = raidGroupData.Content.Health.Value
		metric["health.raid.group.health.desciption.ids"] = raidGroupData.Content.Health.DescriptionIds
		metric["health.raid.group.health.desciption.descriptions"] = raidGroupData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.raidGroup.Message
		}

		metrics = append(metrics, metric)
	}

	for _, treeQuotaData := range data.treeQuota.Entries {
		metric := mapstr.M{}
		metric["health.tree.quota.id"] = treeQuotaData.Content.ID
		metric["health.tree.quota.soft.limit"] = treeQuotaData.Content.SoftLimit
		metric["health.tree.quota.hard.limit"] = treeQuotaData.Content.HardLimit
		metric["health.tree.quota.size.used"] = treeQuotaData.Content.SizeUsed
		metric["health.tree.quota.size.state"] = treeQuotaData.Content.State

		if debug {
			metric["health.message"] = data.treeQuota.Message
		}

		metrics = append(metrics, metric)
	}

	for _, diskGroupData := range data.diskGroup.Entries {
		metric := mapstr.M{}
		metric["health.disk.group.id"] = diskGroupData.Content.ID
		metric["health.disk.group.name"] = diskGroupData.Content.Name
		metric["health.disk.group.advertised.size"] = diskGroupData.Content.AdvertisedSize
		metric["health.disk.group.disk.size"] = diskGroupData.Content.DiskSize
		metric["health.disk.group.hot.spare.policy.status"] = diskGroupData.Content.HotSparePolicyStatus
		metric["health.disk.group.min.spare.policy.status"] = diskGroupData.Content.MinHotSpareCandidates
		metric["health.disk.group.rpm"] = diskGroupData.Content.Rpm
		metric["health.disk.group.speed"] = diskGroupData.Content.Speed
		metric["health.disk.group.total.disks"] = diskGroupData.Content.TotalDisks
		metric["health.disk.group.unconfigured.disks"] = diskGroupData.Content.UnconfiguredDisks

		if debug {
			metric["health.message"] = data.diskGroup.Message
		}

		metrics = append(metrics, metric)
	}

	for _, cifsServerData := range data.cifsServer.Entries {
		metric := mapstr.M{}
		metric["health.cifs.server.id"] = cifsServerData.Content.ID
		metric["health.cifs.server.name"] = cifsServerData.Content.Name
		metric["health.cifs.server.health.value"] = cifsServerData.Content.Health.Value
		metric["health.cifs.server.health.desciption.ids"] = cifsServerData.Content.Health.DescriptionIds
		metric["health.cifs.server.health.desciption.descriptions"] = cifsServerData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.cifsServer.Message
		}

		metrics = append(metrics, metric)
	}

	for _, fastCacheData := range data.fastCache.Entries {
		metric := mapstr.M{}
		metric["health.fast.cache.id"] = fastCacheData.Content.ID
		metric["health.fast.cache.name"] = fastCacheData.Content.Name
		metric["health.fast.cache.number_of_disks"] = fastCacheData.Content.NumberOfDisks
		metric["health.fast.cache.size.free"] = fastCacheData.Content.SizeFree
		metric["health.fast.cache.size.total"] = fastCacheData.Content.SizeTotal
		metric["health.fast.cache.health.value"] = fastCacheData.Content.Health.Value
		metric["health.fast.cache.health.description.ids"] = fastCacheData.Content.Health.DescriptionIds
		metric["health.fast.cache.health.descriptions"] = fastCacheData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.fastCache.Message
		}

		metrics = append(metrics, metric)
	}

	for _, fastVPData := range data.fastVP.Entries {
		metric := mapstr.M{}
		metric["health.fast.vp.id"] = fastVPData.Content.ID
		metric["health.fast.vp.status"] = fastVPData.Content.Status
		metric["health.fast.vp.is_schedule_enabled"] = fastVPData.Content.IsScheduleEnabled
		metric["health.fast.vp.relocation.duration.estimate"] = fastVPData.Content.RelocationDurationEstimate
		metric["health.fast.vp.relocation.rate"] = fastVPData.Content.RelocationRate
		metric["health.fast.vp.size.moving.down"] = fastVPData.Content.SizeMovingDown
		metric["health.fast.vp.size.moving.up"] = fastVPData.Content.SizeMovingUp
		metric["health.fast.vp.size.moving.within"] = fastVPData.Content.SizeMovingWithin

		if debug {
			metric["health.message"] = data.fastVP.Message
		}

		metrics = append(metrics, metric)
	}

	for _, fcPortData := range data.fcPort.Entries {
		metric := mapstr.M{}
		metric["health.fc.port.id"] = fcPortData.Content.ID
		metric["health.fc.port.name"] = fcPortData.Content.Name
		metric["health.fc.port.health.value"] = fcPortData.Content.Health.Value
		metric["health.fc.port.health.desciption.ids"] = fcPortData.Content.Health.DescriptionIds
		metric["health.fc.port.health.desciption.descriptions"] = fcPortData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.fcPort.Message
		}

		metrics = append(metrics, metric)
	}

	for _, hostContainerData := range data.hostContainer.Entries {
		metric := mapstr.M{}
		metric["health.host.container.id"] = hostContainerData.Content.ID
		metric["health.host.container.name"] = hostContainerData.Content.Name
		metric["health.host.container.health.value"] = hostContainerData.Content.Health.Value
		metric["health.host.container.health.desciption.ids"] = hostContainerData.Content.Health.DescriptionIds
		metric["health.host.container.health.desciption.descriptions"] = hostContainerData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.hostContainer.Message
		}

		metrics = append(metrics, metric)
	}

	for _, hostInitiatorData := range data.hostInitiator.Entries {
		metric := mapstr.M{}
		metric["health.host.initiator.id"] = hostInitiatorData.Content.ID
		metric["health.host.initiator.health.value"] = hostInitiatorData.Content.Health.Value
		metric["health.host.initiator.health.desciption.ids"] = hostInitiatorData.Content.Health.DescriptionIds
		metric["health.host.initiator.health.desciption.descriptions"] = hostInitiatorData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.hostInitiator.Message
		}

		metrics = append(metrics, metric)
	}

	for _, hostData := range data.host.Entries {
		metric := mapstr.M{}
		metric["health.host.id"] = hostData.Content.ID
		metric["health.host.name"] = hostData.Content.Name
		metric["health.host.health.value"] = hostData.Content.Health.Value
		metric["health.host.health.desciption.ids"] = hostData.Content.Health.DescriptionIds
		metric["health.host.health.desciption.descriptions"] = hostData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.host.Message
		}

		metrics = append(metrics, metric)
	}

	for _, ioModuleData := range data.ioModule.Entries {
		metric := mapstr.M{}
		metric["health.io.module.id"] = ioModuleData.Content.ID
		metric["health.io.module.name"] = ioModuleData.Content.Name
		metric["health.io.module.needs_replacement"] = ioModuleData.Content.NeedsReplacment
		metric["health.io.module.health.value"] = ioModuleData.Content.Health.Value
		metric["health.io.module.health.desciption.ids"] = ioModuleData.Content.Health.DescriptionIds
		metric["health.io.module.health.desciption.descriptions"] = ioModuleData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.ioModule.Message
		}

		metrics = append(metrics, metric)
	}

	for _, lccData := range data.lcc.Entries {
		metric := mapstr.M{}
		metric["health.lcc.id"] = lccData.Content.ID
		metric["health.lcc.name"] = lccData.Content.Name
		metric["health.lcc.needs_replacement"] = lccData.Content.NeedsReplacment
		metric["health.lcc.health.value"] = lccData.Content.Health.Value
		metric["health.lcc.health.desciption.ids"] = lccData.Content.Health.DescriptionIds
		metric["health.lcc.health.desciption.descriptions"] = lccData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.lcc.Message
		}

		metrics = append(metrics, metric)
	}

	for _, nasServerData := range data.nasServer.Entries {
		metric := mapstr.M{}
		metric["health.nas.server.id"] = nasServerData.Content.ID
		metric["health.nas.server.name"] = nasServerData.Content.Name
		metric["health.nas.server.health.value"] = nasServerData.Content.Health.Value
		metric["health.nas.server.health.desciption.ids"] = nasServerData.Content.Health.DescriptionIds
		metric["health.nas.server.health.desciption.descriptions"] = nasServerData.Content.Health.Descriptions

		if debug {
			metric["health.message"] = data.nasServer.Message
		}

		metrics = append(metrics, metric)
	}

	// #TODO FIX THIS no org id
	reportMetricsForUnityStorage(reporter, baseURL, metrics)
}
