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

func CreateStatisticsFields(stats Statistics) mapstr.M {
	return mapstr.M{
		"timestamp":      stats.Timestamp,
		"status":         stats.Status,
		"throughput_raw": createIOLatencyFields(stats.ThroughputRaw),
		"iops_raw":       createIOLatencyFields(stats.IOPSRaw),
		"latency_raw":    createIOLatencyFields(stats.LatencyRaw),
	}
}

func CreateMetricsFields(metric Metrics) mapstr.M {
	return mapstr.M{
		"timestamp":  metric.Timestamp,
		"duration":   metric.Duration,
		"status":     metric.Status,
		"latency":    createIOLatencyFields(metric.Latency),
		"iops":       createIOLatencyFields(metric.IOPS),
		"throughput": createIOLatencyFields(metric.Throughput),
	}
}

func createIOLatencyFields(io IOLatency) mapstr.M {
	return mapstr.M{
		"read":  io.Read,
		"write": io.Write,
		"other": io.Other,
		"total": io.Total,
	}
}

// For processing basic endpoints we need a type-specific function to create the fields.
// ProcessEndpoint is a generic function that can be used for all basic endpoints,
// but it needs to know how to create the fields for the specific type.
type DispatchFunc func(*NetAppRestClient, Endpoint) ([]mb.Event, error)

// makeDispatchFunc creates a DispatchFunc for a specific type T. The create[type name]Fields functions
// passed to this function are defined with concrete types, so that's where we get our T type from.
// This allows us to use the same ProcessEndpoint function for all basic endpoints, while still being type-safe.
func MakeDispatchFunc[T any](createFunc CreateFieldsFunc[T]) DispatchFunc {
	return func(client *NetAppRestClient, e Endpoint) ([]mb.Event, error) {
		return ProcessEndpoint(*client, e, createFunc)
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
