package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
)

var timestamp time.Time

// {Name: "SnapmirrorRelationships", Endpoint: "/api/snapmirror/relationships", Fn: getSnapmirrorRelationships},
// func getSnapmirrorRelationships(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[SnapMirrorRelationship]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}
// 	var events []mb.Event
// 	for _, record := range records.Records {
// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createSnapMirrorRelationshipFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// // {Name: "QuotaRules", Endpoint: "/api/storage/quota/rules", Fn: getQuotaRules},
// func getQuotaRules(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[QuotaRule]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {

// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createQuotaRuleFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// // {Name: "Aggregates", Endpoint: "/api/storage/aggregates", Fn: getAggregates}
// func getAggregates(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[Aggregate]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {
// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createAggregateFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// {Name: "Disks", Endpoint: "/api/storage/disks", Fn: getDisks}
func getDisks(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {

	client := m.netappClient

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[Disk]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// unroll the disk paths array (for Kibana) into event per path.
	var events []mb.Event
	for _, record := range records.Records {
		pathEvents := createDiskPathEvents(m, record)
		events = append(events, pathEvents...)
	}

	return events, nil
}

// {Name: "LUNs", Endpoint: "/api/storage/luns", Fn: getLUNs}
// func getLUNs(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[LUN]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {
// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createLUNFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// // {Name: "QosPolicies", Endpoint: "/api/storage/qos/policies", Fn: getQosPolicies}
// func getQosPolicies(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[QosPolicy]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {

// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createQosPolicyFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// // {Name: "Qtrees", Endpoint: "/api/storage/qtrees", Fn: getQtrees},
// func getQtrees(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[Qtree]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {

// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createQTreeFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// // {Name: "QuotaReports", Endpoint: "/api/storage/quota/reports", Fn: getQuotaReports},
// func getQuotaReports(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[QuotaReport]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {

// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createQuotaReportFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// {Name: "Shelves", Endpoint: "/api/storage/shelves", Fn: getShelves},
func getShelves(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
	timestamp = time.Now().UTC()
	client := m.netappClient

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
// func getVolumes(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[Volume]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {

// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createVolumeFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// // {Name: "SvmPeers", Endpoint: "/api/svm/peers", Fn: getSvmPeers}
// func getSvmPeers(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[SVMPeer]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {

// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createSVMPeerFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// // {Name: "Svms", Endpoint: "/api/svm/svms", Fn: getSvms}
// func getSvms(m *MetricSet, endpoint Endpoint) ([]mb.Event, error) {
// 	timestamp = time.Now().UTC()
// 	client := m.netappClient

// 	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	var records Records[SVM]
// 	err = json.Unmarshal([]byte(response), &records)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}

// 	var events []mb.Event
// 	for _, record := range records.Records {

// 		event := mb.Event{
// 			Timestamp:       timestamp,
// 			MetricSetFields: createSVMFields(record),
// 			RootFields:      createECSFields(m),
// 		}
// 		events = append(events, event)
// 	}

// 	return events, nil
// }

func ProcessEndpoint[T any](m *MetricSet, endpoint Endpoint, createFieldsFunc CreateFieldsFunc[T]) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.netappClient

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records Records[T]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {
		metricSetFields := createFieldsFunc(record)
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: metricSetFields,
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}
