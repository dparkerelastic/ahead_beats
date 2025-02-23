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
	"sort"
	"strings"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/dellecs"
)

// dashboard/nodes/{nodeid}/processes
// func getNodeProcessEvents(m *MetricSet) ([]mb.Event, error) {
// 	//timestamp := time.Now().UTC()
// 	client := m.ecsClient

// 	endpoint, err := getEndpoint("Processes")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var events []mb.Event
// 	var response Response

// 	for _, node := range m.nodes {
// 		nodeEndpoint := strings.Replace(endpoint.Endpoint, "{nodeid}", node.ID, -1)
// 		result, err := client.Get(nodeEndpoint)
// 		err = json.Unmarshal([]byte(result), &response)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to unmarshal local nodes response")
// 		}

// 		mb.Event{
// 			Timestamp: timestamp,
// 			MetricSetFields: map[string]interface{}{
// 				"array_monitor.writes_per_sec":    monitor.WritesPerSec,
// 				"array_monitor.usec_per_write_op": monitor.UsecPerWriteOp,
// 				"array_monitor.output_per_sec":    monitor.OutputPerSec,
// 				"array_monitor.reads_per_sec":     monitor.ReadsPerSec,
// 				"array_monitor.input_per_sec":     monitor.InputPerSec,
// 				"array_monitor.time":              monitor.Time,
// 				"array_monitor.usec_per_read_op":  monitor.UsecPerReadOp,
// 				"array_monitor.queue_depth":       monitor.QueueDepth,
// 			},
// 			RootFields: purestorage.MakeRootFields(m.config),
// 		})

// 		events = append(events, nodeEvents...)
// 	}
// 	return events, nil

// }

// "dashboard/nodes/{nodeid}"
func getNoteDetailsEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.ecsClient
	endpoint, err := getEndpoint("Node Details")
	if err != nil {
		return nil, err
	}
	var events []mb.Event
	for _, node := range m.nodes {

		nodeEndpoint := strings.Replace(endpoint.Endpoint, "{nodeid}", node.ID, -1)
		result, err := client.Get(nodeEndpoint)
		var nodeDetails NodeData
		err = json.Unmarshal([]byte(result), &nodeDetails)
		if err != nil {
			fmt.Printf("Failed to unmarshal node details for node %s: %v\n", node.ID, err)
			continue
		}

		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"node_details.id":                                      nodeDetails.ID,
				"node_details.storage_pool_id":                         nodeDetails.StoragePoolID,
				"node_details.name":                                    nodeDetails.DisplayName,
				"node_details.api_change":                              nodeDetails.APIChange,
				"node_details.num_bad_disks":                           nodeDetails.NumBadDisks,
				"node_details.storage_pool_name":                       nodeDetails.StoragePoolName,
				"node_details.display_name":                            nodeDetails.DisplayName,
				"node_details.num_ready_to_replace_disks":              nodeDetails.NumReadyToReplaceDisks,
				"node_details.num_maintenance_disks":                   nodeDetails.NumMaintenanceDisks,
				"node_details.disk_space_allocated_percentage_summary": nodeDetails.DiskSpaceAllocatedPercentageSummary,
				"node_details.health_status":                           nodeDetails.HealthStatus,
				"node_details.num_good_disks":                          nodeDetails.NumGoodDisks,
				"node_details.num_disks":                               nodeDetails.NumDisks,
				"node_details.disk_space_free_current_l1":              lastSpaceValue(node.DiskSpaceFreeCurrentL1),
				"node_details.allocated_capacity_forecast":             lastSpaceValue(node.AllocatedCapacityForecast),
				"node_details.disk_space_free_current_l2":              lastSpaceValue(node.DiskSpaceFreeCurrentL2),
				"node_details.disk_space_reserved_current":             lastSpaceValue(node.DiskSpaceReservedCurrent),
				"node_details.disk_space_free_l2":                      lastSpaceValue(node.DiskSpaceFreeL2),
				"node_details.disk_space_free_l1":                      lastSpaceValue(node.DiskSpaceFreeL1),
				"node_details.disk_space_total":                        lastSpaceValue(node.DiskSpaceTotal),
				"node_details.disk_space_total_current_l1":             lastSpaceValue(node.DiskSpaceTotalCurrentL1),
				"node_details.disk_space_total_current_l2":             lastSpaceValue(node.DiskSpaceTotalCurrentL2),
				"node_details.disk_space_offline_total_current":        lastSpaceValue(node.DiskSpaceOfflineTotalCurrent),
				"node_details.disk_space_allocated_l2":                 lastSpaceValue(node.DiskSpaceAllocatedL2),
				"node_details.disk_space_allocated_l1":                 lastSpaceValue(node.DiskSpaceAllocatedL1),
				"node_details.disk_space_free_current":                 lastSpaceValue(node.DiskSpaceFreeCurrent),
				"node_details.disk_space_total_current":                lastSpaceValue(node.DiskSpaceTotalCurrent),
				"node_details.disk_space_allocated_current":            lastSpaceValue(node.DiskSpaceAllocatedCurrent),
				"node_details.disk_space_allocated_current_l1":         lastSpaceValue(node.DiskSpaceAllocatedCurrentL1),
				"node_details.disk_space_allocated_current_l2":         lastSpaceValue(node.DiskSpaceAllocatedCurrentL2),
				"node_details.disk_space_allocated":                    lastSpaceValue(node.DiskSpaceAllocated),
				"node_details.disk_space_allocated_percent":            lastPercentValue(node.DiskSpaceAllocatedPercentage),
				"node_details.disk_space_allocated_percentage_current": lastPercentValue(node.DiskSpaceAllocatedPercentageCurrent),
			},
			RootFields: dellecs.MakeRootFields(m.config),
		}
		addDiskSpaceSummaryValues(nodeDetails, &event, "node_details")
		events = append(events, event)

	}

	return events, nil
}

// "dashboard/nodes/{nodeid}/disks"
func getDiskEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.ecsClient
	endpoint, err := getEndpoint("Disks")
	if err != nil {
		return nil, err
	}
	var events []mb.Event
	for _, node := range m.nodes {

		diskEndpoint := strings.Replace(endpoint.Endpoint, "{nodeid}", node.ID, -1)
		result, err := client.Get(diskEndpoint)

		var response Response
		err = json.Unmarshal([]byte(result), &response)
		if err != nil {
			fmt.Printf("failed to unmarshal Disks response")
			continue
		}

		for _, rawDisk := range response.Embedded.Instances {
			var disk DiskData
			err := json.Unmarshal(rawDisk, &disk)
			if err != nil {
				fmt.Printf("Failed to unmarshal disk instance for node %s: %v\n", node.ID, err)
				continue
			}
			event := mb.Event{
				Timestamp: timestamp,
				MetricSetFields: map[string]interface{}{
					"disks.node_id":                                 node.ID,
					"disks.storage_pool_name":                       disk.StoragePoolName,
					"disks.display_name":                            disk.DisplayName,
					"disks.node_display_name":                       disk.NodeDisplayName,
					"disks.slot_id":                                 disk.SlotId,
					"disks.id":                                      disk.Id,
					"disks.storage_pool_id":                         disk.StoragePoolId,
					"disks.ssm_l2_status":                           disk.SsmL2Status,
					"disks.health_status":                           disk.HealthStatus,
					"disks.ssm_l1_status":                           disk.SsmL1Status,
					"disks.disk_space_free_current_l1":              lastSpaceValue(disk.DiskSpaceFreeCurrentL1),
					"disks.disk_space_free_current_l2":              lastSpaceValue(disk.DiskSpaceFreeCurrentL2),
					"disks.disk_space_free_l2":                      lastSpaceValue(disk.DiskSpaceFreeL2),
					"disks.disk_space_free_l1":                      lastSpaceValue(disk.DiskSpaceFreeL1),
					"disks.disk_space_total":                        lastSpaceValue(disk.DiskSpaceTotal),
					"disks.disk_space_total_current_l1":             lastSpaceValue(disk.DiskSpaceTotalCurrentL1),
					"disks.disk_space_total_current_l2":             lastSpaceValue(disk.DiskSpaceTotalCurrentL2),
					"disks.disk_space_allocated_l2":                 lastSpaceValue(disk.DiskSpaceAllocatedL2),
					"disks.disk_space_allocated_l1":                 lastSpaceValue(disk.DiskSpaceAllocatedL1),
					"disks.disk_space_free_current":                 lastSpaceValue(disk.DiskSpaceFreeCurrent),
					"disks.disk_space_total_current":                lastSpaceValue(disk.DiskSpaceTotalCurrent),
					"disks.disk_space_allocated_current":            lastSpaceValue(disk.DiskSpaceAllocatedCurrent),
					"disks.disk_space_allocated_current_l1":         lastSpaceValue(disk.DiskSpaceAllocatedCurrentL1),
					"disks.disk_space_allocated":                    lastSpaceValue(disk.DiskSpaceAllocated),
					"disks.disk_space_allocated_current_l2":         lastSpaceValue(disk.DiskSpaceAllocatedCurrentL2),
					"disks.disk_space_free":                         lastSpaceValue(disk.DiskSpaceFree),
					"disks.disk_space_allocated_percent":            lastPercentValue(disk.DiskSpaceAllocatedPercentage),
					"disks.disk_space_allocated_percentage_current": lastPercentValue(disk.DiskSpaceAllocatedPercentageCurrent),
				},
				RootFields: dellecs.MakeRootFields(m.config),
			}
			addDiskSpaceSummaryValues(disk, &event, "disks")
			events = append(events, event)
		}
	}
	return events, nil
}

func getCapacityEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.ecsClient
	endpoint, err := getEndpoint("Capacity")
	if err != nil {
		return nil, err
	}
	var events []mb.Event
	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var storage StorageInfo
	err = json.Unmarshal([]byte(response), &storage)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal capacity data: %v\n", err)
	}

	event := mb.Event{
		Timestamp: timestamp,
		MetricSetFields: map[string]interface{}{
			"capacity.total_provisioned_db": storage.TotalProvisionedGB,
			"capacity.total_free_db":        storage.TotalFreeGB,
		},
		RootFields: dellecs.MakeRootFields(m.config),
	}
	events = append(events, event)
	return events, nil
}

func getStoragePoolsEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.ecsClient
	endpoint, err := getEndpoint("Storage Pools")
	if err != nil {
		return nil, err
	}

	var response Response
	result, err := client.Get(endpoint.Endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Pools response: %w", err)

	}

	var events []mb.Event
	for _, rawPool := range response.Embedded.Instances {
		var pool StoragePoolData
		err := json.Unmarshal(rawPool, &pool)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal Storage Pool instance: %w", err)
		}

		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"storage_pool.chunks_l1_journal_total_size":         pool.ChunksL1JournalTotalSize,
				"storage_pool.chunks_l1_btree_number":               pool.ChunksL1BtreeNumber,
				"storage_pool.chunks_l1_btree_total_size":           pool.ChunksL1BtreeTotalSize,
				"storage_pool.chunks_l1_journal_avg_size":           pool.ChunksL1JournalAvgSize,
				"storage_pool.chunks_l1_journal_number":             pool.ChunksL1JournalNumber,
				"storage_pool.chunks_l0_journal_total_size":         pool.ChunksL0JournalTotalSize,
				"storage_pool.chunks_l0_btree_number":               pool.ChunksL0BtreeNumber,
				"storage_pool.chunks_l0_btree_total_size":           pool.ChunksL0BtreeTotalSize,
				"storage_pool.chunks_l0_btree_avg_size":             pool.ChunksL0BtreeAvgSize,
				"storage_pool.chunks_l0_journal_avg_size":           pool.ChunksL0JournalAvgSize,
				"storage_pool.chunks_l0_journal_number":             pool.ChunksL0JournalNumber,
				"storage_pool.chunks_repo_total_seal_size":          pool.ChunksRepoTotalSealSize,
				"storage_pool.chunks_repo_number":                   pool.ChunksRepoNumber,
				"storage_pool.chunks_xor_number":                    pool.ChunksXorNumber,
				"storage_pool.chunks_xor_total_size":                pool.ChunksXorTotalSize,
				"storage_pool.chunks_geo_cache_total_size":          pool.ChunksGeoCacheTotalSize,
				"storage_pool.chunks_geo_cache_count":               pool.ChunksGeoCacheCount,
				"storage_pool.chunks_geo_copy_number":               pool.ChunksGeoCopyNumber,
				"storage_pool.chunks_geo_copy_total_size":           pool.ChunksGeoCopyTotalSize,
				"storage_pool.num_nodes":                            pool.NumNodes,
				"storage_pool.num_disks":                            pool.NumDisks,
				"storage_pool.num_bad_nodes":                        pool.NumBadNodes,
				"storage_pool.num_good_nodes":                       pool.NumGoodNodes,
				"storage_pool.num_bad_disks":                        pool.NumBadDisks,
				"storage_pool.num_good_disks":                       pool.NumGoodDisks,
				"storage_pool.num_maintenance_nodes":                pool.NumMaintenanceNodes,
				"storage_pool.num_maintenance_disks":                pool.NumMaintenanceDisks,
				"storage_pool.num_ready_to_replace_disks":           pool.NumReadyToReplaceDisks,
				"storage_pool.num_nodes_with_sufficient_disk_space": pool.NumNodesWithSufficientDiskSpace,
				"storage_pool.gc_user_data_is_enabled":              pool.GcUserDataIsEnabled,
				"storage_pool.gc_system_metadata_is_enabled":        pool.GcSystemMetadataIsEnabled,
				"storage_pool.recovery_complete_time_estimate":      pool.RecoveryCompleteTimeEstimate,
				"storage_pool.chunks_ec_complete_time_estimate":     pool.ChunksEcCompleteTimeEstimate,
				"storage_pool.id":                                   pool.ID,
				"storage_pool.name":                                 pool.Name,
			},
			RootFields: dellecs.MakeRootFields(m.config),
		}

		addDiskSpaceSummaryValues(pool, &event, "storage_pools")
		events = append(events, event)
	}

	return events, nil

}

func getReplicationGroupsEvents(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.ecsClient
	endpoint, err := getEndpoint("Replication Groups")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return nil, nil
}

func getLocalNodes(client *ECSRestClient) ([]NodeData, error) {

	nodesResponse, err := client.Get("dashboard/zones/localzone/nodes")
	if err != nil {
		return nil, fmt.Errorf("failed to get local nodes")
	}
	var response Response
	err = json.Unmarshal([]byte(nodesResponse), &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal local nodes response")
	}

	var nodeInstances []NodeData
	for _, rawInstance := range response.Embedded.Instances {
		var nodeInstance NodeData
		err := json.Unmarshal(rawInstance, &nodeInstance)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal node instance: %v", err)
		}
		nodeInstances = append(nodeInstances, nodeInstance)
	}
	return nodeInstances, nil
}

func latestSpaceValue(spaces []TimestampedSpace) (TimestampedSpace, error) {
	if len(spaces) == 0 {
		return TimestampedSpace{}, fmt.Errorf("empty array")
	}

	latest := spaces[0]
	sort.Slice(spaces, func(i, j int) bool {
		return time.Unix(0, spaces[i].T).After(time.Unix(0, spaces[j].T))
	})
	latest = spaces[0]

	return latest, nil
}

func lastSpaceValue(spaces []TimestampedSpace) int {
	if len(spaces) == 0 {
		return -1
	}
	return int(spaces[len(spaces)-1].Space)
}

func lastPercentValue(spaces []TimestampedPercent) int {
	if len(spaces) == 0 {
		return -1
	}
	return int(spaces[len(spaces)-1].Percent)
}

func addDiskSpaceSummaryValues(data interface{}, event *mb.Event, prefix string) {
	var summaries map[string]Summary

	switch v := data.(type) {
	case NodeData:
		summaries = map[string]Summary{
			prefix + ".disk_space_free_l1_summary":      v.DiskSpaceFreeL1Summary,
			prefix + ".disk_space_allocated_l2_summary": v.DiskSpaceAllocatedL2Summary,
			prefix + ".disk_space_total_summary":        v.DiskSpaceTotalSummary,
			prefix + ".disk_space_free_l2_summary":      v.DiskSpaceFreeL2Summary,
			prefix + ".disk_space_allocated_summary":    v.DiskSpaceAllocatedSummary,
			prefix + ".disk_space_free_summary":         v.DiskSpaceFreeSummary,
			prefix + ".disk_space_allocated_l1_summary": v.DiskSpaceAllocatedL1Summary,
		}
	case DiskData:
		summaries = map[string]Summary{
			prefix + ".disk_space_free_l1_summary":      v.DiskSpaceFreeL1Summary,
			prefix + ".disk_space_allocated_l2_summary": v.DiskSpaceAllocatedL2Summary,
			prefix + ".disk_space_total_summary":        v.DiskSpaceTotalSummary,
			prefix + ".disk_space_free_l2_summary":      v.DiskSpaceFreeL2Summary,
			prefix + ".disk_space_allocated_summary":    v.DiskSpaceAllocatedSummary,
			prefix + ".disk_space_free_summary":         v.DiskSpaceFreeSummary,
			prefix + ".disk_space_allocated_l1_summary": v.DiskSpaceAllocatedL1Summary,
		}
	case StoragePoolData:
		summaries = map[string]Summary{
			prefix + ".chunks_ec_coded_total_seal_size_summary":       v.ChunksEcCodedTotalSealSizeSummary,
			prefix + ".chunks_ec_rate_summary":                        v.ChunksEcRateSummary,
			prefix + ".chunks_ec_coded_ratio_summary":                 v.ChunksEcCodedRatioSummary,
			prefix + ".chunks_ec_applicable_total_seal_size_summary":  v.ChunksEcApplicableTotalSealSizeSummary,
			prefix + ".recovery_rate_summary":                         v.RecoveryRateSummary,
			prefix + ".recovery_bad_chunks_total_size_summary":        v.RecoveryBadChunksTotalSizeSummary,
			prefix + ".disk_space_allocated_l1_summary":               v.DiskSpaceAllocatedL1Summary,
			prefix + ".disk_space_allocated_l2_summary":               v.DiskSpaceAllocatedL2Summary,
			prefix + ".disk_space_allocated_summary":                  v.DiskSpaceAllocatedSummary,
			prefix + ".disk_space_allocated_percentage_summary":       v.DiskSpaceAllocatedPercentageSummary,
			prefix + ".disk_space_free_l1_summary":                    v.DiskSpaceFreeL1Summary,
			prefix + ".disk_space_free_l2_summary":                    v.DiskSpaceFreeL2Summary,
			prefix + ".disk_space_total_summary":                      v.DiskSpaceTotalSummary,
			prefix + ".disk_space_free_summary":                       v.DiskSpaceFreeSummary,
		}

	for key, summary := range summaries {
		event.MetricSetFields[key+".min"] = summary.Min[0].Space
		event.MetricSetFields[key+".max"] = summary.Max[0].Space
		event.MetricSetFields[key+".avg"] = summary.Avg
	}
}
