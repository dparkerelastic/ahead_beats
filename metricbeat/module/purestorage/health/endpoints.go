package health

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/beats/v7/metricbeat/mb"
)

type Endpoint struct {
	Name       string
	Endpoint   string
	ReturnType interface{}
}

var endpoints = map[string]Endpoint{
	"ArrayControllers": {Name: "ArrayControllers", Endpoint: "array?controllers=true", ReturnType: ArrayController{}},
	"ArrayMonitor":     {Name: "ArrayMonitor", Endpoint: "array?action=monitor", ReturnType: ArrayMonitor{}},
	"ArraySpace":       {Name: "ArraySpace", Endpoint: "array?space=true", ReturnType: ArraySpace{}},
	"Hardware":         {Name: "Hardware", Endpoint: "hardware", ReturnType: Hardware{}},
	"Drive":            {Name: "Drive", Endpoint: "drive", ReturnType: Drive{}},
	"PGroup":           {Name: "PGroup", Endpoint: "pgroup?snap=true&transfer=true&pending=true", ReturnType: PGroup{}},
	"VolumeMonitor":    {Name: "VolumeMonitor", Endpoint: "volume?action=monitor", ReturnType: VolumeMonitor{}},
	"VolumeSpace":      {Name: "VolumeSpace", Endpoint: "volume?space=true", ReturnType: VolumeSpace{}},
	"Array":            {Name: "Array", Endpoint: "array", ReturnType: Array{}},
	"Volume":           {Name: "Volume", Endpoint: "volume", ReturnType: Volume{}},
}

func getEndpoint(name string) (Endpoint, error) {
	endpoint, ok := endpoints[name]
	if !ok {
		return Endpoint{}, fmt.Errorf("%s not found in the map", name)
	}
	return endpoint, nil
}

func getArrayControllersEvents(m *MetricSet) ([]mb.Event, error) {
	client := m.psClient
	endpoint, err := getEndpoint("ArrayControllers")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, err
	}

	var arrayControllers []ArrayController
	err = json.Unmarshal([]byte(response), &arrayControllers)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, controller := range arrayControllers {
		events = append(events, mb.Event{
			MetricSetFields: map[string]interface{}{
				"status":  controller.Status,
				"name":    controller.Name,
				"version": controller.Version,
				"mode":    controller.Mode,
				"model":   controller.Model,
			},
		})
	}

	return events, nil
}

func getArrayMonitorEvents(m *MetricSet) ([]mb.Event, error) {

	client := m.psClient
	endpoint, err := getEndpoint("ArrayControllers")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, err
	}

	var arrayMonitors []ArrayMonitor
	err = json.Unmarshal([]byte(response), &arrayMonitors)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, monitor := range arrayMonitors {
		events = append(events, mb.Event{
			MetricSetFields: map[string]interface{}{
				"writes_per_sec":    monitor.WritesPerSec,
				"usec_per_write_op": monitor.UsecPerWriteOp,
				"output_per_sec":    monitor.OutputPerSec,
				"reads_per_sec":     monitor.ReadsPerSec,
				"input_per_sec":     monitor.InputPerSec,
				"time":              monitor.Time,
				"usec_per_read_op":  monitor.UsecPerReadOp,
				"queue_depth":       monitor.QueueDepth,
			},
		})
	}

	return events, nil
}

func getArraySpaceEvents(m *MetricSet) ([]mb.Event, error) {
	return nil, nil
}

func getHardwareEvents(m *MetricSet) ([]mb.Event, error) {
	return nil, nil
}
func getDriveEvents(m *MetricSet) ([]mb.Event, error) {
	return nil, nil
}
func getPGroupEvents(m *MetricSet) ([]mb.Event, error) {
	return nil, nil
}
func getVolumeMonitorEvents(m *MetricSet) ([]mb.Event, error) {
	return nil, nil
}

func getVolumeSpaceEvents(m *MetricSet) ([]mb.Event, error) {
	return nil, nil
}
func getArrayEvents(m *MetricSet) ([]mb.Event, error) {
	return nil, nil
}

func getVolumeEvents(m *MetricSet) ([]mb.Event, error) {
	return nil, nil
}
