package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
)

var timestamp time.Time

// {Name: "Disks", Endpoint: "/api/storage/disks", Fn: getDisks}
// custom endpoint handling for disks so that we can unroll the disk paths array
func getDisks(client *netapp.NetAppRestClient, endpoint netapp.Endpoint) ([]mb.Event, error) {

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records netapp.Records[Disk]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// unroll the disk paths array (for Kibana) into event per path.
	var events []mb.Event
	for _, record := range records.Records {
		pathEvents := createDiskPathEvents(record)
		events = append(events, pathEvents...)
	}

	return events, nil
}

// {Name: "Shelves", Endpoint: "/api/storage/shelves", Fn: getShelves},
// custom endpoint handling for shelves so that we can unroll the components
// (PSUs, fans, thermal, voltage, current, ports, acps) into individual events
func getShelves(client *netapp.NetAppRestClient, endpoint netapp.Endpoint) ([]mb.Event, error) {
	timestamp = time.Now().UTC()

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records netapp.Records[Shelf]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		psuEvents := createPSUEvents(record)
		events = append(events, psuEvents...)

		fanEvents := createFanEvents(record)
		events = append(events, fanEvents...)

		thermalEvents := createThermalEvents(record)
		events = append(events, thermalEvents...)

		voltageEvents := createVoltageEvents(record)
		events = append(events, voltageEvents...)

		currentEvents := createCurrentEvents(record)
		events = append(events, currentEvents...)

		portEvents := createPortEvents(record)
		events = append(events, portEvents...)

		acpEvents := createACPEvents(record)
		events = append(events, acpEvents...)
	}

	return events, nil
}
