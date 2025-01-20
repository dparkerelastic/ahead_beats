// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package health

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/purestorage"
)

type Endpoint struct {
	Name       string
	Endpoint   string
	ReturnType interface{}
	Fn         func(*MetricSet) ([]mb.Event, error)
}

func getArrayControllersEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("ArrayControllers")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var arrayControllers []ArrayController
	err = json.Unmarshal([]byte(response), &arrayControllers)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, controller := range arrayControllers {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"array_controller.status":  controller.Status,
				"array_controller.name":    controller.Name,
				"array_controller.version": controller.Version,
				"array_controller.mode":    controller.Mode,
				"array_controller.model":   controller.Model,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil
}

func getArrayMonitorEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("ArrayControllers")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var arrayMonitors []ArrayMonitor
	err = json.Unmarshal([]byte(response), &arrayMonitors)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, monitor := range arrayMonitors {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"array_monitor.writes_per_sec":    monitor.WritesPerSec,
				"array_monitor.usec_per_write_op": monitor.UsecPerWriteOp,
				"array_monitor.output_per_sec":    monitor.OutputPerSec,
				"array_monitor.reads_per_sec":     monitor.ReadsPerSec,
				"array_monitor.input_per_sec":     monitor.InputPerSec,
				"array_monitor.time":              monitor.Time,
				"array_monitor.usec_per_read_op":  monitor.UsecPerReadOp,
				"array_monitor.queue_depth":       monitor.QueueDepth,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil
}

func getArraySpaceEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("ArraySpace")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var arraySpaces []ArraySpace
	err = json.Unmarshal([]byte(response), &arraySpaces)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, space := range arraySpaces {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"array_space.capacity":          space.Capacity,
				"array_space.hostname":          space.Hostname,
				"array_space.system":            space.System,
				"array_space.snapshots":         space.Snapshots,
				"array_space.volumes":           space.Volumes,
				"array_space.data_reduction":    space.DataReduction,
				"array_space.total":             space.Total,
				"array_space.shared_space":      space.SharedSpace,
				"array_space.thin_provisioning": space.ThinProvisioning,
				"array_space.total_reduction":   space.TotalReduction,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil

}

func getHardwareEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("Hardware")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var hardwareItems []Hardware
	err = json.Unmarshal([]byte(response), &hardwareItems)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, item := range hardwareItems {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"hardware.status":      item.Status,
				"hardware.slot":        item.Slot,
				"hardware.name":        item.Name,
				"hardware.temperature": item.Temperature,
				"hardware.index":       item.Index,
				"hardware.identify":    item.Identify,
				"hardware.speed":       item.Speed,
				"hardware.details":     item.Details,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil
}

func getDriveEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("Drive")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var drives []Drive
	err = json.Unmarshal([]byte(response), &drives)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, drive := range drives {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"drive.status":   drive.Status,
				"drive.capacity": drive.Capacity,
				"drive.type":     drive.Type,
				"drive.name":     drive.Name,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil
}

func getPGroupEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("PGroup")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var pGroups []PGroup
	err = json.Unmarshal([]byte(response), &pGroups)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, pGroup := range pGroups {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"pgroup.name":                   pGroup.Name,
				"pgroup.physical_bytes_written": pGroup.PhysicalBytesWritten,
				"pgroup.started":                pGroup.Started,
				"pgroup.completed":              pGroup.Completed,
				"pgroup.created":                pGroup.Created,
				"pgroup.source":                 pGroup.Source,
				"pgroup.time_remaining":         pGroup.TimeRemaining,
				"pgroup.progress":               pGroup.Progress,
				"pgroup.data_transferred":       pGroup.DataTransferred,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil
}

func getVolumeMonitorEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("VolumeMonitor")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var volumeMonitors []VolumeMonitor
	err = json.Unmarshal([]byte(response), &volumeMonitors)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, space := range volumeMonitors {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"volume_monitor.writes_per_sec":    space.WritesPerSec,
				"volume_monitor.name":              space.Name,
				"volume_monitor.usec_per_write_op": space.UsecPerWriteOp,
				"volume_monitor.output_per_sec":    space.OutputPerSec,
				"volume_monitor.reads_per_sec":     space.ReadsPerSec,
				"volume_monitor.input_per_sec":     space.InputPerSec,
				"volume_monitor.time":              space.Time,
				"volume_monitor.usec_per_read_op":  space.UsecPerReadOp,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil
}

func getVolumeSpaceEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("VolumeSpace")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var volumeSpaces []VolumeSpace
	err = json.Unmarshal([]byte(response), &volumeSpaces)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, space := range volumeSpaces {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"volume_space.size":              space.Size,
				"volume_space.name":              space.Name,
				"volume_space.system":            space.System,
				"volume_space.snapshots":         space.Snapshots,
				"volume_space.volumes":           space.Volumes,
				"volume_space.data_reduction":    space.DataReduction,
				"volume_space.total":             space.Total,
				"volume_space.shared_space":      space.SharedSpace,
				"volume_space.thin_provisioning": space.ThinProvisioning,
				"volume_space.total_reduction":   space.TotalReduction,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil
}
func getArrayEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("Array")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var array Array
	err = json.Unmarshal([]byte(response), &array)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	events = append(events, mb.Event{
		Timestamp: timestamp,
		MetricSetFields: map[string]interface{}{
			"array.revision":   array.Revision,
			"array.version":    array.Version,
			"array.array_name": array.ArrayName,
			"array.id":         array.ID,
		},
		RootFields: purestorage.MakeRootFields(m.config.HostIp),
	})

	return events, nil
}

func getVolumeEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.psClient
	endpoint, err := getEndpoint("Volume")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var volumes []Volume
	err = json.Unmarshal([]byte(response), &volumes)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, volume := range volumes {
		events = append(events, mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"volume.name":    volume.Name,
				"volume.created": volume.Created,
				"volume.source":  volume.Source,
				"volume.serial":  volume.Serial,
				"volume.size":    volume.Size,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		})
	}

	return events, nil
}
