package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
)

var timestamp time.Time

// {Name: "SnapmirrorRelationships", Endpoint: "/api/snapmirror/relationships", Fn: getSnapmirrorRelationships},
func getSnapmirrorRelationships(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("SnapmirrorRelationships")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
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

// {Name: "Aggregates", Endpoint: "/api/storage/aggregates", Fn: getAggregates}
func getAggregates(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Aggregates")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Aggregate]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createAggregateFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

// {Name: "Disks", Endpoint: "/api/storage/disks", Fn: getDisks}
func getDisks(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Disks")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Disk]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createDiskFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

// {Name: "LUNs", Endpoint: "/api/storage/luns", Fn: getLUNs}
func getLUNs(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("LUNs")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[LUN]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createLUNFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

// {Name: "QosPolicies", Endpoint: "/api/storage/qos/policies", Fn: getQosPolicies}
func getQosPolicies(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("QosPolicies")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
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

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Qtree]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createQTreeFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

// {Name: "QuotaReports", Endpoint: "/api/storage/quota/reports", Fn: getQuotaReports},
func getQuotaReports(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("QuotaReports")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[QuotaReport]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createQuotaReportFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

// {Name: "Shelves", Endpoint: "/api/storage/shelves", Fn: getShelves},
func getShelves(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Shelves")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Shelf]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		psuEvents := createPSUEvents(m, record)
		events = append(events, psuEvents...)

		fanEvents := createFanEvents(m, record)
		events = append(events, fanEvents...)

		thermalEvents := createThermalEvents(m, record)
		events = append(events, thermalEvents...)

		voltageEvents := createVoltageEvents(m, record)
		events = append(events, voltageEvents...)

		currentEvents := createCurrentEvents(m, record)
		events = append(events, currentEvents...)

		portEvents := createPortEvents(m, record)
		events = append(events, portEvents...)

		acpEvents := createACPEvents(m, record)
		events = append(events, acpEvents...)
	}

	return events, nil
}

// {Name: "Volumes", Endpoint: "/api/storage/volumes", Fn: getVolumes},
func getVolumes(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Volumes")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Volume]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createVolumeFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

// {Name: "SvmPeers", Endpoint: "/api/svm/peers", Fn: getSvmPeers}
func getSvmPeers(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("SvmPeers")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[SVMPeer]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createSVMPeerFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

// {Name: "Svms", Endpoint: "/api/svm/svms", Fn: getSvms}
func getSvms(m *MetricSet) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient
	endpoint, err := getEndpoint("Svms")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[SVM]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createSVMFields(record),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}
