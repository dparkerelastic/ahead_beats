package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

var timestamp time.Time

func createECSFields(ms *MetricSet) mapstr.M {

	return mapstr.M{
		"observer": mapstr.M{
			"hostname": ms.config.HostInfo.Hostname,
			"ip":       ms.config.HostInfo.IP,
			"type":     "storage",
			"vendor":   "NetApp",
		},
	}
}

// {Name: "SnapmirrorRelationships", Endpoint: "/api/snapmirror/relationships", Fn: getSnapmirrorRelationships},
func getSnapmirrorRelationships(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("SnapmirrorRelationships")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[SnapMirrorRelationship]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	var events []mb.Event
	for _, record := range records.Records {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createSnapMirrorRelationshipFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

func createSnapMirrorRelationshipFields(record SnapMirrorRelationship) mapstr.M {
	return mapstr.M{
		"backoff_level":              record.BackoffLevel,
		"consistency_group_failover": createGroupFailoverFields(record.ConsistencyGroupFailover),
		"destination":                createSnapMirrorEndpointFields(record.Destination),
		"exported_snapshot":          record.ExportedSnapshot,
		"group_type":                 record.GroupType,
		"healthy":                    record.Healthy,
		"identity_preservation":      record.IdentityPreservation,
		"io_serving_copy":            record.IOServingCopy,
		"lag_time":                   record.LagTime,
		"last_transfer_network_compression_ratio": record.LastTransferNetworkRatio,
		"last_transfer_type":                      record.LastTransferType,
		"master_bias_activated_site":              record.MasterBiasActivatedSite,
		"policy":                                  record.Policy,
		"preferred_site":                          record.PreferredSite,
		"restore":                                 record.Restore,
		"source":                                  createSnapMirrorEndpointFields(record.Source),
		"state":                                   record.State,
		"svmdr_volumes":                           record.SvmdrVolumes,
		"throttle":                                record.Throttle,
		"total_transfer_bytes":                    record.TotalTransferBytes,
		"total_transfer_duration":                 record.TotalTransferDuration,
		"transfer":                                record.Transfer,
		"transfer_schedule":                       record.TransferSchedule,
		"unhealthy_reason":                        record.UnhealthyReason,
		"uuid":                                    record.UUID,
	}
}

func createGroupFailoverFields(failover ConsistencyGroupFailover) mapstr.M {
	return mapstr.M{
		"error":  createStatusFields(failover.Error),
		"state":  failover.State,
		"status": createStatusFields(failover.Status),
		"type":   failover.Type,
	}
}

func createStatusFields(status StorageStatus) mapstr.M {
	return mapstr.M{
		"code":    status.Code,
		"message": status.Message,
	}
}

func createSnapMirrorEndpointFields(endpoint SnapMirrorEndpoint) mapstr.M {
	// Convert consistency group volumes
	volumes, err := netapp.ToJSONString(endpoint.ConsistencyGroupVolumes)
	if err != nil {
		volumes = ""
	}

	return mapstr.M{
		"cluster": mapstr.M{
			"name": endpoint.Cluster.Name,
			"uuid": endpoint.Cluster.UUID,
		},
		"svm": mapstr.M{
			"name": endpoint.SVM.Name,
			"uuid": endpoint.SVM.UUID,
		},
		"luns": mapstr.M{
			"name": endpoint.LUNs.Name,
			"uuid": endpoint.LUNs.UUID,
		},
		"path":                      endpoint.Path,
		"consistency_group_volumes": volumes,
	}
}

// {Name: "Aggregates", Endpoint: "/api/storage/aggregates", Fn: getAggregates}
func getAggregates(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Aggregates")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Aggregate]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "Disks", Endpoint: "/api/storage/disks", Fn: getDisks}
func getDisks(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Disks")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Disk]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "LUNs", Endpoint: "/api/storage/luns", Fn: getLUNs}
func getLUNs(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("LUNs")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[LUN]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "QosPolicies", Endpoint: "/api/storage/qos/policies", Fn: getQosPolicies}
func getQosPolicies(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("QosPolicies")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[QosPolicy]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "Qtrees", Endpoint: "/api/storage/qtrees", Fn: getQtrees},
func getQtrees(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Qtrees")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Qtree]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "QuotaReports", Endpoint: "/api/storage/quota/reports", Fn: getQuotaReports},
func getQuotaReports(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("QuotaReports")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[QuotaReport]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "QuotaRules", Endpoint: "/api/storage/quota/rules", Fn: getQuotaRules},
func getQuotaRules(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("QuotaRules")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[QuotaRule]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "Shelves", Endpoint: "/api/storage/shelves", Fn: getShelves},
func getShelves(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Shelves")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Shelf]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "Volumes", Endpoint: "/api/storage/volumes", Fn: getVolumes},
func getVolumes(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Volumes")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Volume]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "SvmPeers", Endpoint: "/api/svm/peers", Fn: getSvmPeers}
func getSvmPeers(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("SvmPeers")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[SVMPeer]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}

// {Name: "Svms", Endpoint: "/api/svm/svms", Fn: getSvms}
func getSvms(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Svms")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[SVM]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, nil
}
