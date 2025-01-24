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

		metrics = append(metrics, metric)
	}

	for _, poolData := range data.pool.Entries {
		metric := mapstr.M{}
		metric["pool.id"] = poolData.Content.ID
		metric["pool.name"] = poolData.Content.Name
		metric["pool.description"] = poolData.Content.Description
		metric["pool.size.free"] = poolData.Content.SizeFree
		metric["pool.size.total"] = poolData.Content.SizeTotal
		metric["pool.health.value"] = poolData.Content.Health.Value
		metric["pool.health.description.ids"] = poolData.Content.Health.DescriptionIds
		metric["pool.health.descriptions"] = poolData.Content.Health.Descriptions

		metrics = append(metrics, metric)
	}

	for _, poolUnitData := range data.poolUnit.Entries {
		metric := mapstr.M{}
		metric["poolunit.id"] = poolUnitData.Content.ID
		metric["poolunit.name"] = poolUnitData.Content.Name
		metric["poolunit.size.total"] = poolUnitData.Content.SizeTotal
		metric["poolunit.health.value"] = poolUnitData.Content.Health.Value
		metric["poolunit.health.description.ids"] = poolUnitData.Content.Health.DescriptionIds
		metric["poolunit.health.descriptions"] = poolUnitData.Content.Health.Descriptions

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

		metrics = append(metrics, metric)
	}

	for _, ethernetPortData := range data.ethernetPort.Entries {
		metric := mapstr.M{}
		metric["ethernet.port.id"] = ethernetPortData.Content.ID
		metric["ethernet.port.health.value"] = ethernetPortData.Content.Health.Value
		metric["ethernet.port.health.description.ids"] = ethernetPortData.Content.Health.DescriptionIds
		metric["ethernet.port.health.descriptions"] = ethernetPortData.Content.Health.Descriptions

		metrics = append(metrics, metric)
	}

	for _, fileInterfaceData := range data.fileInterface.Entries {
		metric := mapstr.M{}
		metric["file.interface.id"] = fileInterfaceData.Content.ID
		metric["file.interface.health.value"] = fileInterfaceData.Content.Health.Value
		metric["file.interface.health.description.ids"] = fileInterfaceData.Content.Health.DescriptionIds
		metric["file.interface.health.descriptions"] = fileInterfaceData.Content.Health.Descriptions

		metrics = append(metrics, metric)
	}

	for _, remoteSystemData := range data.remoteSystem.Entries {
		metric := mapstr.M{}
		metric["file.system.id"] = remoteSystemData.Content.ID
		metric["file.system.health.value"] = remoteSystemData.Content.Health.Value
		metric["file.system.health.description.ids"] = remoteSystemData.Content.Health.DescriptionIds
		metric["file.system.health.descriptions"] = remoteSystemData.Content.Health.Descriptions

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

		metrics = append(metrics, metric)
	}

	for _, datastoreData := range data.datastore.Entries {
		metric := mapstr.M{}
		metric["datastore.id"] = datastoreData.Content.ID
		metric["datastore.name"] = datastoreData.Content.Name
		metric["datastore.size.total"] = datastoreData.Content.SizeTotal
		metric["datastore.size.used"] = datastoreData.Content.SizeUsed

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

		metrics = append(metrics, metric)
	}

	// #TODO FIX THIS no org id
	reportMetricsForUnityStorage(reporter, baseURL, metrics)
}
