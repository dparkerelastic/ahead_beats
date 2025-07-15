package protocols

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func getISCSISessions(client *netapp.NetAppRestClient, endpoint netapp.Endpoint) ([]mb.Event, error) {

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records netapp.Records[ISCSISession]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// unroll the disk paths array (for Kibana) into event per path.
	var events []mb.Event
	for _, record := range records.Records {
		connectEvents := createISCSIConnectionEvents(record)
		events = append(events, connectEvents...)
	}

	return events, nil
}

// createISCSIConnectionEvents creates an event for each iSCSI connection in a session.
func createISCSIConnectionEvents(record ISCSISession) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, conn := range record.Connections {
		sessionFields := createISCSISessionFields(record)
		sessionFields["connection"] = createISCSIConnectionFields(conn)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"iscsi_session": sessionFields,
			},
			RootFields: netapp.CreateECSFields(),
		}
		events = append(events, event)
	}
	return events
}

func getIGroups(client *netapp.NetAppRestClient, endpoint netapp.Endpoint) ([]mb.Event, error) {

	response, err := client.GetWithFields(endpoint.Endpoint, endpoint.QueryFields)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var records netapp.Records[IGroup]
	err = json.Unmarshal([]byte(response), &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var events []mb.Event
	for _, record := range records.Records {
		// unroll the initiator array (for Kibana) into event per path.
		initiatorEvents := createInitiatorEvents(record)
		events = append(events, initiatorEvents...)
	}

	return events, nil
}

func createInitiatorEvents(record IGroup) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, initiator := range record.Initiators {
		iGroupFields := createIGroupFields(record)
		iGroupFields["initiator"] = createInitiatorFields(initiator)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"igroup": iGroupFields,
			},
			RootFields: netapp.CreateECSFields(),
		}
		events = append(events, event)
	}
	return events
}
