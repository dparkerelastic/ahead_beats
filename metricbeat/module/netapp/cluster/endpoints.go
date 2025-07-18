package cluster

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// {Name: "Shelves", Endpoint: "/api/storage/shelves", Fn: getShelves},
// custom endpoint handling for shelves so that we can unroll the components
// (PSUs, fans, thermal, voltage, current, ports, acps) into individual events
func getClusterNodes(client *netapp.NetAppRestClient, endpoint netapp.Endpoint) ([]mb.Event, error) {

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records netapp.Records[ClusterNode]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		haEvents := createHAEvents(record)
		events = append(events, haEvents...)

	}

	return events, nil
}

func createHAEvents(record ClusterNode) []mb.Event {
	timestamp := time.Now().UTC()

	var events []mb.Event
	for _, aggStat := range record.HA.Giveback.Status {

		nodeFields := createNodeFields(record)
		haFields := createNodeHAFields(record.HA)
		haFields["giveback"] = createGivebackStatusFields(aggStat)
		nodeFields["ha"] = haFields
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"node": nodeFields,
			},
			RootFields: netapp.CreateECSFields(),
		}
		events = append(events, event)
	}
	return events
}

func getSensors(client *netapp.NetAppRestClient, endpoint netapp.Endpoint) ([]mb.Event, error) {
	query := map[string][]string{
		"type": {"battery-life", "fan", "thermal", "voltage"},
	}

	response, err := client.GetWithQuery(endpoint.Endpoint, endpoint.QueryFields, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records netapp.Records[Sensor]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {
		fields := createSensorFields(record)
		event := mb.Event{
			Timestamp: time.Now().UTC(),
			MetricSetFields: mapstr.M{
				"sensor": fields,
			},
			RootFields: netapp.CreateECSFields(),
		}
		events = append(events, event)
	}

	return events, nil
}

func getSystemCounters(client *netapp.NetAppRestClient, endpoint netapp.Endpoint) ([]mb.Event, error) {

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records netapp.Records[CounterTableRow]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {

		counterEvents := createCounterEvents(record)
		events = append(events, counterEvents...)

	}

	return events, nil
}

func createCounterEvents(record CounterTableRow) []mb.Event {
	var events []mb.Event
	timestamp := time.Now().UTC()

	for _, counter := range record.Counters {
		fields := createCounterTableFields(record)
		fields["counter_name"] = counter.Name
		fields["counter_value"] = counter.Value

		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"counter": fields,
			},
			RootFields: netapp.CreateECSFields(),
		}
		events = append(events, event)
	}

	return events
}
