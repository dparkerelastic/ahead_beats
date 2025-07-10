package netapp

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

type CreateFieldsFunc[T any] func(T) mapstr.M
type EndpointFunc func(*NetAppRestClient, Endpoint) ([]mb.Event, error)

type Endpoint struct {
	Name        string
	Endpoint    string
	GetFunc     EndpointFunc // nil for basic endpoints
	QueryFields []string
}

func CreateStatusFields(status Status) mapstr.M {
	return mapstr.M{
		"code":    status.Code,
		"message": status.Message,
	}
}

func CreateECSFields() mapstr.M {
	return mapstr.M{
		"observer": mapstr.M{
			"type":   "storage",
			"vendor": "NetApp",
		},
	}
}

func CreateNamedObjectFields(namedObject NamedObject) mapstr.M {
	return mapstr.M{
		"name": namedObject.Name,
		"uuid": namedObject.UUID,
	}
}

// ProcessEndpoint is a generic function to process basic endpoints.
// It takes a MetricSet, an Endpoint, and a function to create fields for the specific type T.
func ProcessEndpoint[T any](client NetAppRestClient, endpoint Endpoint, createFieldsFunc CreateFieldsFunc[T]) ([]mb.Event, error) {
	timestamp := time.Now().UTC()

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
			RootFields:      CreateECSFields(),
		}
		events = append(events, event)
	}

	return events, nil
}
