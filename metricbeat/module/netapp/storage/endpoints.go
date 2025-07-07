package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
)

var timestamp time.Time

// {Name: "Disks", Endpoint: "/api/storage/disks", Fn: getDisks}
// custom endpoint handling for disks so that we can unroll the disk paths array
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

// {Name: "Shelves", Endpoint: "/api/storage/shelves", Fn: getShelves},
// custom endpoint handling for shelves so that we can unroll the components
// (PSUs, fans, thermal, voltage, current, ports, acps) into individual events
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

// ProcessEndpoint is a generic function to process basic endpoints.
// It takes a MetricSet, an Endpoint, and a function to create fields for the specific type T.
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
