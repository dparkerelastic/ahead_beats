package storage_health

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

func reportMetrics(reporter mb.ReporterV2, baseURL string, data UnityData) {
	metrics := []mapstr.M{}

	for _, systemData := range data.system.Entries {
		metric := mapstr.M{}
		metric["system.id"] = systemData.Content.ID
		metric["system.name"] = systemData.Content.Name
		metric["system.model"] = systemData.Content.Model
		metric["system.serial_number"] = systemData.Content.SerialNumber
		metric["system.internal_model"] = systemData.Content.InternalModel
		metric["system.platform"] = systemData.Content.Platform
		metric["system.macaddress"] = systemData.Content.MacAddress
		//metric["system.system_uuid"] = systemData.Content.SystemUUID
		metric["system.health.value"] = systemData.Content.Health.Value
		metric["system.health.description.ids"] = systemData.Content.Health.DescriptionIds
		metric["system.health.descriptions"] = systemData.Content.Health.Descriptions

		metric["message"] = data.system.Message
		metrics = append(metrics, metric)
	}

	for _, poolData := range data.pool.Entries {
		metric := mapstr.M{}
		metric["pool.id"] = poolData.Content.ID
		metric["pool.name"] = poolData.Content.Name
		metric["pool.description"] = poolData.Content.Description
		metric["pool.size.free"] = poolData.Content.SizeFree
		metric["pool.size.used"] = poolData.Content.SizeUsed
		metric["pool.size.total"] = poolData.Content.SizeTotal
		metric["pool.health.value"] = poolData.Content.Health.Value
		metric["pool.health.description.ids"] = poolData.Content.Health.DescriptionIds
		metric["pool.health.descriptions"] = poolData.Content.Health.Descriptions
		metric["pool.harvest.state"] = poolData.Content.HarvestState
		metric["pool.metadata.size.subscribed"] = poolData.Content.MetadataSizeSubscribed
		metric["pool.metadata.size.used"] = poolData.Content.MetadataSizeUsed
		metric["pool.rebalance.progress"] = poolData.Content.RebalanceProgress
		metric["pool.size.subscribed"] = poolData.Content.SizeSubscribed
		metric["pool.snap.size.subscribed"] = poolData.Content.SnapSizeSubscribed
		metric["pool.snap.size.used"] = poolData.Content.SnapSizeUsed

		//calculation - if divide by zero is not going to occur
		if poolData.Content.SizeTotal > 0 {
			metric["pool.size.percent.used"] = int((float64(poolData.Content.SizeUsed) / float64(poolData.Content.SizeTotal)) * float64(100))
		}

		metric["message"] = data.pool.Message

		metrics = append(metrics, metric)
	}

	for _, poolUnitData := range data.poolUnit.Entries {
		metric := mapstr.M{}
		metric["pool.unit.id"] = poolUnitData.Content.ID
		metric["pool.unit.name"] = poolUnitData.Content.Name
		metric["pool.unit.size.total"] = poolUnitData.Content.SizeTotal
		metric["pool.unit.health.value"] = poolUnitData.Content.Health.Value
		metric["pool.unit.health.description.ids"] = poolUnitData.Content.Health.DescriptionIds
		metric["pool.unit.health.descriptions"] = poolUnitData.Content.Health.Descriptions

		metric["message"] = data.poolUnit.Message

		metrics = append(metrics, metric)
	}

	for _, lunData := range data.lun.Entries {
		metric := mapstr.M{}
		metric["lun.id"] = lunData.Content.ID
		metric["lun.size.total"] = lunData.Content.SizeTotal
		metric["lun.size.allocated"] = lunData.Content.SizeAllocated
		metric["lun.snap.size"] = lunData.Content.SnapsSize
		metric["lun.snap.size.allocated"] = lunData.Content.SnapsSizeAllocated
		metric["lun.metadata.size"] = lunData.Content.MetadataSize
		metric["lun.metadata.size.allocated"] = lunData.Content.MetadataSizeAllocated
		metric["lun.health.value"] = lunData.Content.Health.Value
		metric["lun.health.description.ids"] = lunData.Content.Health.DescriptionIds
		metric["lun.health.descriptions"] = lunData.Content.Health.Descriptions

		metric["message"] = data.lun.Message

		metrics = append(metrics, metric)
	}

	for _, storageProcessorData := range data.storageProcesser.Entries {
		metric := mapstr.M{}
		metric["storage.processor.id"] = storageProcessorData.Content.ID
		metric["storage.processor.model"] = storageProcessorData.Content.Model
		metric["storage.processor.name"] = storageProcessorData.Content.Name
		metric["storage.processor.health.value"] = storageProcessorData.Content.Health.Value
		metric["storage.processor.health.description.ids"] = storageProcessorData.Content.Health.DescriptionIds
		metric["storage.processor.health.descriptions"] = storageProcessorData.Content.Health.Descriptions

		metric["message"] = data.storageProcesser.Message

		metrics = append(metrics, metric)
	}

	for _, storageResourceData := range data.storageResource.Entries {
		metric := mapstr.M{}
		metric["storage.resource.id"] = storageResourceData.Content.ID
		metric["storage.resource.health.value"] = storageResourceData.Content.Health.Value
		metric["storage.resource.health.description.ids"] = storageResourceData.Content.Health.DescriptionIds
		metric["storage.resource.health.descriptions"] = storageResourceData.Content.Health.Descriptions
		metric["storage.resource.size.total"] = storageResourceData.Content.SizeTotal
		metric["storage.resource.size.allocated"] = storageResourceData.Content.SizeAllocated
		metric["storage.resource.snap.count"] = storageResourceData.Content.SnapCount
		metric["storage.resource.snaps.size.total"] = storageResourceData.Content.SnapsSizeTotal
		metric["storage.resource.snaps.size.allocated"] = storageResourceData.Content.SnapsSizeAllocated
		metric["storage.resource.metadata.size"] = storageResourceData.Content.MetadataSize
		metric["storage.resource.metadata.size.allocated"] = storageResourceData.Content.MetadataSizeAllocated

		metric["message"] = data.storageResource.Message

		metrics = append(metrics, metric)
	}

	for _, storageTierData := range data.storageTier.Entries {
		metric := mapstr.M{}
		metric["storage.tier.id"] = storageTierData.Content.ID
		metric["storage.tier.size.total"] = storageTierData.Content.SizeTotal
		metric["storage.tier.size.free"] = storageTierData.Content.SizeFree
		metric["storage.tier.disk.total"] = storageTierData.Content.DisksTotal
		metric["storage.tier.disk.unused"] = storageTierData.Content.DisksUnused
		metric["storage.tier.virtual.disk.total"] = storageTierData.Content.VirtualDisksTotal
		metric["storage.tier.virtual.disk.unused"] = storageTierData.Content.VirtualDisksUnused

		if storageTierData.Content.DisksTotal > 0 {
			metric["storage.tier.disk.used"] = storageTierData.Content.DisksTotal - storageTierData.Content.DisksUnused
		}

		if storageTierData.Content.VirtualDisksTotal > 0 {
			metric["storage.tier.virtual.disk.used"] = storageTierData.Content.VirtualDisksTotal - storageTierData.Content.VirtualDisksUnused
		}

		metric["message"] = data.storageTier.Message

		metrics = append(metrics, metric)
	}

	for _, licenseData := range data.license.Entries {
		metric := mapstr.M{}
		metric["license.id"] = licenseData.Content.ID
		metric["license.name"] = licenseData.Content.Name
		metric["license.isvalid"] = licenseData.Content.IsValid
		metric["license.ispermanent"] = licenseData.Content.IsPermanent
		metric["license.isinstalled"] = licenseData.Content.IsInstalled
		metric["license.expires"] = licenseData.Content.Expires
		metric["license.feature.id"] = licenseData.Content.Feature.ID

		metric["message"] = data.license.Message

		metrics = append(metrics, metric)
	}

	for _, ethernetPortData := range data.ethernetPort.Entries {
		metric := mapstr.M{}
		metric["ethernet.port.id"] = ethernetPortData.Content.ID
		metric["ethernet.port.health.value"] = ethernetPortData.Content.Health.Value
		metric["ethernet.port.health.description.ids"] = ethernetPortData.Content.Health.DescriptionIds
		metric["ethernet.port.health.descriptions"] = ethernetPortData.Content.Health.Descriptions

		metric["message"] = data.ethernetPort.Message

		metrics = append(metrics, metric)
	}

	for _, fileInterfaceData := range data.fileInterface.Entries {
		metric := mapstr.M{}
		metric["file.interface.id"] = fileInterfaceData.Content.ID
		metric["file.interface.health.value"] = fileInterfaceData.Content.Health.Value
		metric["file.interface.health.description.ids"] = fileInterfaceData.Content.Health.DescriptionIds
		metric["file.interface.health.descriptions"] = fileInterfaceData.Content.Health.Descriptions

		metric["message"] = data.fileInterface.Message

		metrics = append(metrics, metric)
	}

	for _, remoteSystemData := range data.remoteSystem.Entries {
		metric := mapstr.M{}
		metric["remote.system.id"] = remoteSystemData.Content.ID
		metric["remote.system.health.value"] = remoteSystemData.Content.Health.Value
		metric["remote.system.health.description.ids"] = remoteSystemData.Content.Health.DescriptionIds
		metric["remote.system.health.descriptions"] = remoteSystemData.Content.Health.Descriptions

		metric["message"] = data.remoteSystem.Message

		metrics = append(metrics, metric)
	}

	for _, diskData := range data.disk.Entries {
		metric := mapstr.M{}
		metric["disk.id"] = diskData.Content.ID
		metric["disk.name"] = diskData.Content.Name
		metric["disk.raw.size"] = diskData.Content.RawSize
		metric["disk.size"] = diskData.Content.Size
		metric["disk.vendor.size"] = diskData.Content.VendorSize
		metric["disk.health.value"] = diskData.Content.Health.Value
		metric["disk.health.desciption.ids"] = diskData.Content.Health.DescriptionIds
		metric["disk.health.desciption.descriptions"] = diskData.Content.Health.Descriptions

		metric["message"] = data.disk.Message

		metrics = append(metrics, metric)
	}

	for _, datastoreData := range data.datastore.Entries {
		metric := mapstr.M{}
		metric["datastore.id"] = datastoreData.Content.ID
		metric["datastore.name"] = datastoreData.Content.Name
		metric["datastore.size.total"] = datastoreData.Content.SizeTotal
		metric["datastore.size.used"] = datastoreData.Content.SizeUsed

		//calculation if divide by zero is not going to occur
		if datastoreData.Content.SizeTotal > 0 {
			metric["datastore.size.percent.used"] = int((float64(datastoreData.Content.SizeUsed) / float64(datastoreData.Content.SizeTotal)) * float64(100))
		}

		metric["message"] = data.datastore.Message

		metrics = append(metrics, metric)
	}

	for _, filesystemData := range data.filesystem.Entries {
		metric := mapstr.M{}
		metric["filesystem.id"] = filesystemData.Content.ID
		metric["filesystem.name"] = filesystemData.Content.Name
		metric["filesystem.health.value"] = filesystemData.Content.Health.Value
		metric["filesystem.health.description.ids"] = filesystemData.Content.Health.DescriptionIds
		metric["filesystem.health.descriptions"] = filesystemData.Content.Health.Descriptions
		metric["filesystem.metadata.size"] = filesystemData.Content.MetadataSize
		metric["filesystem.metadata.size.allocated"] = filesystemData.Content.MetadataSizeAllocated
		metric["filesystem.size.total"] = filesystemData.Content.SizeTotal
		metric["filesystem.size.used"] = filesystemData.Content.SizeUsed
		metric["filesystem.size.allocated"] = filesystemData.Content.SizeAllocated
		metric["filesystem.snap.count"] = filesystemData.Content.SnapCount
		metric["filesystem.snaps.size"] = filesystemData.Content.SnapsSize
		metric["filesystem.snaps.size.allocated"] = filesystemData.Content.SnapsSizeAllocated

		//calculation if divide by zero is not going to occur
		if filesystemData.Content.SizeTotal > 0 {
			metric["filesystem.size.percent.used"] = int((float64(filesystemData.Content.SizeUsed) / float64(filesystemData.Content.SizeTotal)) * float64(100))
		}

		metric["message"] = data.filesystem.Message

		metrics = append(metrics, metric)
	}

	for _, snapData := range data.snap.Entries {
		metric := mapstr.M{}
		metric["snap.id"] = snapData.Content.ID
		metric["snap.name"] = snapData.Content.Name
		metric["snap.size"] = snapData.Content.Size
		metric["snap.state"] = snapData.Content.State
		metric["snap.creation.time"] = snapData.Content.CreationTime
		metric["snap.expiration.time"] = snapData.Content.ExpirationTime

		metric["message"] = data.snap.Message

		metrics = append(metrics, metric)
	}

	for _, sasPortData := range data.sasPort.Entries {
		metric := mapstr.M{}
		metric["sas.port.id"] = sasPortData.Content.ID
		metric["sas.port.name"] = sasPortData.Content.Name
		metric["sas.port.needs_replacement"] = sasPortData.Content.NeedsReplacment
		metric["sas.port.health.value"] = sasPortData.Content.Health.Value
		metric["sas.port.health.desciption.ids"] = sasPortData.Content.Health.DescriptionIds
		metric["sas.port.health.desciption.descriptions"] = sasPortData.Content.Health.Descriptions

		metric["message"] = data.sasPort.Message

		metrics = append(metrics, metric)
	}

	for _, powerSupplyData := range data.powerSupply.Entries {
		metric := mapstr.M{}
		metric["power.supply.id"] = powerSupplyData.Content.ID
		metric["power.supply.name"] = powerSupplyData.Content.Name
		metric["power.supply.needs_replacement"] = powerSupplyData.Content.NeedsReplacment
		metric["power.supply.health.value"] = powerSupplyData.Content.Health.Value
		metric["power.supply.health.desciption.ids"] = powerSupplyData.Content.Health.DescriptionIds
		metric["power.supply.health.desciption.descriptions"] = powerSupplyData.Content.Health.Descriptions

		metric["message"] = data.powerSupply.Message

		metrics = append(metrics, metric)
	}

	for _, fanData := range data.fan.Entries {
		metric := mapstr.M{}
		metric["fan.id"] = fanData.Content.ID
		metric["fan.name"] = fanData.Content.Name
		metric["fan.needs_replacement"] = fanData.Content.NeedsReplacment
		metric["fan.health.value"] = fanData.Content.Health.Value
		metric["fan.health.desciption.ids"] = fanData.Content.Health.DescriptionIds
		metric["fan.health.desciption.descriptions"] = fanData.Content.Health.Descriptions

		metric["message"] = data.fan.Message

		metrics = append(metrics, metric)
	}

	for _, daeData := range data.dae.Entries {
		metric := mapstr.M{}
		metric["dae.id"] = daeData.Content.ID
		metric["dae.name"] = daeData.Content.Name
		metric["dae.current.power"] = daeData.Content.CurrentPower
		metric["dae.avg.power"] = daeData.Content.AvgPower
		metric["dae.max.power"] = daeData.Content.MaxPower
		metric["dae.current.temperature"] = daeData.Content.CurrentTemperature
		metric["dae.avg.temperature"] = daeData.Content.AvgTemperature
		metric["dae.max.temperature"] = daeData.Content.MaxTemperature
		metric["dae.needs_replacement"] = daeData.Content.NeedsReplacment
		metric["dae.health.value"] = daeData.Content.Health.Value
		metric["dae.health.desciption.ids"] = daeData.Content.Health.DescriptionIds
		metric["dae.health.desciption.descriptions"] = daeData.Content.Health.Descriptions

		metric["message"] = data.dae.Message

		metrics = append(metrics, metric)
	}

	for _, memoryModuleData := range data.memoryModule.Entries {
		metric := mapstr.M{}
		metric["memory.module.id"] = memoryModuleData.Content.ID
		metric["memory.module.name"] = memoryModuleData.Content.Name
		metric["memory.module.size"] = memoryModuleData.Content.Size
		metric["memory.module.needs_replacement"] = memoryModuleData.Content.NeedsReplacment
		metric["memory.module.health.value"] = memoryModuleData.Content.Health.Value
		metric["memory.module.health.desciption.ids"] = memoryModuleData.Content.Health.DescriptionIds
		metric["memory.module.health.desciption.descriptions"] = memoryModuleData.Content.Health.Descriptions

		metric["message"] = data.memoryModule.Message

		metrics = append(metrics, metric)
	}

	for _, batteryData := range data.battery.Entries {
		metric := mapstr.M{}
		metric["battery.id"] = batteryData.Content.ID
		metric["battery.name"] = batteryData.Content.Name
		metric["battery.needs_replacement"] = batteryData.Content.NeedsReplacment
		metric["battery.health.value"] = batteryData.Content.Health.Value
		metric["battery.health.desciption.ids"] = batteryData.Content.Health.DescriptionIds
		metric["battery.health.desciption.descriptions"] = batteryData.Content.Health.Descriptions

		metric["message"] = data.battery.Message

		metrics = append(metrics, metric)
	}

	for _, ssdData := range data.ssd.Entries {
		metric := mapstr.M{}
		metric["ssd.id"] = ssdData.Content.ID
		metric["ssd.name"] = ssdData.Content.Name
		metric["ssd.needs_replacement"] = ssdData.Content.NeedsReplacment
		metric["ssd.health.value"] = ssdData.Content.Health.Value
		metric["ssd.health.desciption.ids"] = ssdData.Content.Health.DescriptionIds
		metric["ssd.health.desciption.descriptions"] = ssdData.Content.Health.Descriptions

		metric["message"] = data.ssd.Message

		metrics = append(metrics, metric)
	}

	for _, raidGroupData := range data.raidGroup.Entries {
		metric := mapstr.M{}
		metric["raid.group.id"] = raidGroupData.Content.ID
		metric["raid.group.name"] = raidGroupData.Content.Name
		metric["raid.group.size.total"] = raidGroupData.Content.SizeTotal
		metric["raid.group.health.value"] = raidGroupData.Content.Health.Value
		metric["raid.group.health.desciption.ids"] = raidGroupData.Content.Health.DescriptionIds
		metric["raid.group.health.desciption.descriptions"] = raidGroupData.Content.Health.Descriptions

		metric["message"] = data.raidGroup.Message

		metrics = append(metrics, metric)
	}

	for _, treeQuotaData := range data.treeQuota.Entries {
		metric := mapstr.M{}
		metric["tree.quota.id"] = treeQuotaData.Content.ID
		metric["tree.quota.soft.limit"] = treeQuotaData.Content.SoftLimit
		metric["tree.quota.hard.limit"] = treeQuotaData.Content.HardLimit
		metric["tree.quota.size.used"] = treeQuotaData.Content.SizeUsed
		metric["tree.quota.size.state"] = treeQuotaData.Content.State

		metric["message"] = data.treeQuota.Message

		metrics = append(metrics, metric)
	}

	for _, diskGroupData := range data.diskGroup.Entries {
		metric := mapstr.M{}
		metric["disk.group.id"] = diskGroupData.Content.ID
		metric["disk.group.name"] = diskGroupData.Content.Name
		metric["disk.group.advertised.size"] = diskGroupData.Content.AdvertisedSize
		metric["disk.group.disk.size"] = diskGroupData.Content.DiskSize
		metric["disk.group.hot.spare.policy.status"] = diskGroupData.Content.HotSparePolicyStatus
		metric["disk.group.min.spare.policy.status"] = diskGroupData.Content.MinHotSpareCandidates
		metric["disk.group.rpm"] = diskGroupData.Content.Rpm
		metric["disk.group.speed"] = diskGroupData.Content.Speed
		metric["disk.group.total.disks"] = diskGroupData.Content.TotalDisks
		metric["disk.group.unconfigured.disks"] = diskGroupData.Content.UnconfiguredDisks

		metric["message"] = data.diskGroup.Message

		metrics = append(metrics, metric)
	}

	for _, cifsServerData := range data.cifsServer.Entries {
		metric := mapstr.M{}
		metric["cifs.server.id"] = cifsServerData.Content.ID
		metric["cifs.server.name"] = cifsServerData.Content.Name
		metric["cifs.server.health.value"] = cifsServerData.Content.Health.Value
		metric["cifs.server.health.desciption.ids"] = cifsServerData.Content.Health.DescriptionIds
		metric["cifs.server.health.desciption.descriptions"] = cifsServerData.Content.Health.Descriptions

		metric["message"] = data.cifsServer.Message

		metrics = append(metrics, metric)
	}

	for _, fastCacheData := range data.fastCache.Entries {
		metric := mapstr.M{}
		metric["fast.cache.id"] = fastCacheData.Content.ID
		metric["fast.cache.name"] = fastCacheData.Content.Name
		metric["fast.cache.number_of_disks"] = fastCacheData.Content.NumberOfDisks
		metric["fast.cache.size.free"] = fastCacheData.Content.SizeFree
		metric["fast.cache.size.total"] = fastCacheData.Content.SizeTotal
		metric["fast.cache.health.value"] = fastCacheData.Content.Health.Value
		metric["fast.cache.health.description.ids"] = fastCacheData.Content.Health.DescriptionIds
		metric["fast.cache.health.descriptions"] = fastCacheData.Content.Health.Descriptions

		metric["message"] = data.fastCache.Message

		metrics = append(metrics, metric)
	}

	for _, fastVPData := range data.fastVP.Entries {
		metric := mapstr.M{}
		metric["fast.vp.id"] = fastVPData.Content.ID
		metric["fast.vp.status"] = fastVPData.Content.Status
		metric["fast.vp.is_schedule_enabled"] = fastVPData.Content.IsScheduleEnabled
		metric["fast.vp.relocation.duration.estimate"] = fastVPData.Content.RelocationDurationEstimate
		metric["fast.vp.relocation.rate"] = fastVPData.Content.RelocationRate
		metric["fast.vp.size.moving.down"] = fastVPData.Content.SizeMovingDown
		metric["fast.vp.size.moving.up"] = fastVPData.Content.SizeMovingUp
		metric["fast.vp.size.moving.within"] = fastVPData.Content.SizeMovingWithin

		metric["message"] = data.fastVP.Message

		metrics = append(metrics, metric)
	}

	for _, fcPortData := range data.fcPort.Entries {
		metric := mapstr.M{}
		metric["fc.port.id"] = fcPortData.Content.ID
		metric["fc.port.name"] = fcPortData.Content.Name
		metric["fc.port.health.value"] = fcPortData.Content.Health.Value
		metric["fc.port.health.desciption.ids"] = fcPortData.Content.Health.DescriptionIds
		metric["fc.port.health.desciption.descriptions"] = fcPortData.Content.Health.Descriptions

		metric["message"] = data.fcPort.Message

		metrics = append(metrics, metric)
	}

	for _, hostContainerData := range data.hostContainer.Entries {
		metric := mapstr.M{}
		metric["host.container.id"] = hostContainerData.Content.ID
		metric["host.container.name"] = hostContainerData.Content.Name
		metric["host.container.health.value"] = hostContainerData.Content.Health.Value
		metric["host.container.health.desciption.ids"] = hostContainerData.Content.Health.DescriptionIds
		metric["host.container.health.desciption.descriptions"] = hostContainerData.Content.Health.Descriptions

		metric["message"] = data.hostContainer.Message

		metrics = append(metrics, metric)
	}

	for _, hostInitiatorData := range data.hostInitiator.Entries {
		metric := mapstr.M{}
		metric["host.initiator.id"] = hostInitiatorData.Content.ID
		metric["host.initiator.health.value"] = hostInitiatorData.Content.Health.Value
		metric["host.initiator.health.desciption.ids"] = hostInitiatorData.Content.Health.DescriptionIds
		metric["host.initiator.health.desciption.descriptions"] = hostInitiatorData.Content.Health.Descriptions

		metric["message"] = data.hostInitiator.Message

		metrics = append(metrics, metric)
	}

	for _, hostData := range data.host.Entries {
		metric := mapstr.M{}
		metric["host.id"] = hostData.Content.ID
		metric["host.name"] = hostData.Content.Name
		metric["host.health.value"] = hostData.Content.Health.Value
		metric["host.health.desciption.ids"] = hostData.Content.Health.DescriptionIds
		metric["host.health.desciption.descriptions"] = hostData.Content.Health.Descriptions

		metric["message"] = data.host.Message

		metrics = append(metrics, metric)
	}

	for _, ioModuleData := range data.ioModule.Entries {
		metric := mapstr.M{}
		metric["io.module.id"] = ioModuleData.Content.ID
		metric["io.module.name"] = ioModuleData.Content.Name
		metric["io.module.needs_replacement"] = ioModuleData.Content.NeedsReplacment
		metric["io.module.health.value"] = ioModuleData.Content.Health.Value
		metric["io.module.health.desciption.ids"] = ioModuleData.Content.Health.DescriptionIds
		metric["io.module.health.desciption.descriptions"] = ioModuleData.Content.Health.Descriptions

		metric["message"] = data.ioModule.Message

		metrics = append(metrics, metric)
	}

	for _, lccData := range data.lcc.Entries {
		metric := mapstr.M{}
		metric["lcc.id"] = lccData.Content.ID
		metric["lcc.name"] = lccData.Content.Name
		metric["lcc.needs_replacement"] = lccData.Content.NeedsReplacment
		metric["lcc.health.value"] = lccData.Content.Health.Value
		metric["lcc.health.desciption.ids"] = lccData.Content.Health.DescriptionIds
		metric["lcc.health.desciption.descriptions"] = lccData.Content.Health.Descriptions

		metric["message"] = data.lcc.Message

		metrics = append(metrics, metric)
	}

	for _, nasServerData := range data.nasServer.Entries {
		metric := mapstr.M{}
		metric["nas.server.id"] = nasServerData.Content.ID
		metric["nas.server.name"] = nasServerData.Content.Name
		metric["nas.server.health.value"] = nasServerData.Content.Health.Value
		metric["nas.server.health.desciption.ids"] = nasServerData.Content.Health.DescriptionIds
		metric["nas.server.health.desciption.descriptions"] = nasServerData.Content.Health.Descriptions

		metric["message"] = data.nasServer.Message

		metrics = append(metrics, metric)
	}

	// #TODO FIX THIS no org id
	reportMetricsForUnityStorage(reporter, baseURL, metrics)
}
