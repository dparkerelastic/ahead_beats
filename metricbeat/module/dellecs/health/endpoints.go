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
	"strings"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// dashboard/nodes/{nodeid}/processes
func getNodeProcessEvents(m *MetricSet) ([]mb.Event, error) {
	field_prefix := "node_process"
	timestamp := time.Now().UTC()
	client := m.ecsClient
	endpoint, err := getEndpoint("Processes")
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
	for _, rawProcs := range response.Embedded.Instances {
		var proc NodeProcess
		err := json.Unmarshal(rawProcs, &proc)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal Storage Pool instance: %w", err)
		}
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				field_prefix: mapstr.M{
					"process_name":             proc.ProcessName,
					"process_id":               proc.ID,
					"cpu_utilization":          lastPercentValue(proc.CPUUtilization),
					"java_heap_utilization":    lastPercentValue(proc.JavaHeapUtilization),
					"max_java_heap_size":       lastBytesValue(proc.MaxJavaHeapSize),
					"memory_utilization_bytes": lastBytesValue(proc.MemoryUtilizationBytes),
					"memory_utilization":       lastPercentValue(proc.MemoryUtilization),
					"thread_count":             lastCountValue(proc.ThreadCount),
					"restart_time":             lastDummyValue(proc.RestartTime),
				},
			},
			RootFields: createECSFields(m),
		}
		events = append(events, event)

	}

	return events, nil

}

// "dashboard/nodes/{nodeid}"
func getNoteDetailsEvents(m *MetricSet) ([]mb.Event, error) {
	field_prefix := "node_details"
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
		if err != nil {
			return nil, fmt.Errorf("failed to get node endpoint: %w", err)
		}

		var nodeDetails NodeData
		err = json.Unmarshal([]byte(result), &nodeDetails)
		if err != nil {
			fmt.Printf("Failed to unmarshal node details for node %s: %v\n", node.ID, err)
			continue
		}

		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				field_prefix: mapstr.M{
					"node_id":                                 nodeDetails.ID,
					"storage_pool_id":                         nodeDetails.StoragePoolID,
					"name":                                    nodeDetails.DisplayName,
					"api_change":                              nodeDetails.APIChange,
					"num_bad_disks":                           nodeDetails.NumBadDisks,
					"storage_pool_name":                       nodeDetails.StoragePoolName,
					"display_name":                            nodeDetails.DisplayName,
					"num_ready_to_replace_disks":              nodeDetails.NumReadyToReplaceDisks,
					"num_maintenance_disks":                   nodeDetails.NumMaintenanceDisks,
					"health_status":                           nodeDetails.HealthStatus,
					"num_good_disks":                          nodeDetails.NumGoodDisks,
					"num_disks":                               nodeDetails.NumDisks,
					"disk_space_free_current_l1":              lastSpaceValue(node.DiskSpaceFreeCurrentL1),
					"allocated_capacity_forecast":             lastSpaceValue(node.AllocatedCapacityForecast),
					"disk_space_free_current_l2":              lastSpaceValue(node.DiskSpaceFreeCurrentL2),
					"disk_space_reserved_current":             lastSpaceValue(node.DiskSpaceReservedCurrent),
					"disk_space_free_l2":                      lastSpaceValue(node.DiskSpaceFreeL2),
					"disk_space_free_l1":                      lastSpaceValue(node.DiskSpaceFreeL1),
					"disk_space_total":                        lastSpaceValue(node.DiskSpaceTotal),
					"disk_space_total_current_l1":             lastSpaceValue(node.DiskSpaceTotalCurrentL1),
					"disk_space_total_current_l2":             lastSpaceValue(node.DiskSpaceTotalCurrentL2),
					"disk_space_offline_total_current":        lastSpaceValue(node.DiskSpaceOfflineTotalCurrent),
					"disk_space_allocated_l2":                 lastSpaceValue(node.DiskSpaceAllocatedL2),
					"disk_space_allocated_l1":                 lastSpaceValue(node.DiskSpaceAllocatedL1),
					"disk_space_free_current":                 lastSpaceValue(node.DiskSpaceFreeCurrent),
					"disk_space_total_current":                lastSpaceValue(node.DiskSpaceTotalCurrent),
					"disk_space_allocated_current":            lastSpaceValue(node.DiskSpaceAllocatedCurrent),
					"disk_space_allocated_current_l1":         lastSpaceValue(node.DiskSpaceAllocatedCurrentL1),
					"disk_space_allocated_current_l2":         lastSpaceValue(node.DiskSpaceAllocatedCurrentL2),
					"disk_space_allocated":                    lastSpaceValue(node.DiskSpaceAllocated),
					"disk_space_allocated_percent":            lastPercentValue(node.DiskSpaceAllocatedPercentage),
					"disk_space_allocated_percentage_current": lastPercentValue(node.DiskSpaceAllocatedPercentageCurrent),
				},
			},
			RootFields: createECSFields(m),
		}
		addSpaceSummaryValues(nodeDetails, &event, field_prefix)
		addPercentageSummaryValues(nodeDetails, &event, field_prefix)
		events = append(events, event)

	}

	return events, nil
}

// "dashboard/nodes/{nodeid}/disks"
func getDiskEvents(m *MetricSet) ([]mb.Event, error) {
	field_prefix := "disks"
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
		if err != nil {
			return nil, fmt.Errorf("failed to get disk endpoint: %w", err)
		}

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
					field_prefix: mapstr.M{
						"node_id":                                 node.ID,
						"storage_pool_name":                       disk.StoragePoolName,
						"display_name":                            disk.DisplayName,
						"node_display_name":                       disk.NodeDisplayName,
						"slot_id":                                 disk.SlotId,
						"disk_id":                                 disk.Id,
						"storage_pool_id":                         disk.StoragePoolId,
						"ssm_l2_status":                           disk.SsmL2Status,
						"health_status":                           disk.HealthStatus,
						"ssm_l1_status":                           disk.SsmL1Status,
						"disk_space_free_current_l1":              lastSpaceValue(disk.DiskSpaceFreeCurrentL1),
						"disk_space_free_current_l2":              lastSpaceValue(disk.DiskSpaceFreeCurrentL2),
						"disk_space_free_l2":                      lastSpaceValue(disk.DiskSpaceFreeL2),
						"disk_space_free_l1":                      lastSpaceValue(disk.DiskSpaceFreeL1),
						"disk_space_total":                        lastSpaceValue(disk.DiskSpaceTotal),
						"disk_space_total_current_l1":             lastSpaceValue(disk.DiskSpaceTotalCurrentL1),
						"disk_space_total_current_l2":             lastSpaceValue(disk.DiskSpaceTotalCurrentL2),
						"disk_space_allocated_l2":                 lastSpaceValue(disk.DiskSpaceAllocatedL2),
						"disk_space_allocated_l1":                 lastSpaceValue(disk.DiskSpaceAllocatedL1),
						"disk_space_free_current":                 lastSpaceValue(disk.DiskSpaceFreeCurrent),
						"disk_space_total_current":                lastSpaceValue(disk.DiskSpaceTotalCurrent),
						"disk_space_allocated_current":            lastSpaceValue(disk.DiskSpaceAllocatedCurrent),
						"disk_space_allocated_current_l1":         lastSpaceValue(disk.DiskSpaceAllocatedCurrentL1),
						"disk_space_allocated":                    lastSpaceValue(disk.DiskSpaceAllocated),
						"disk_space_allocated_current_l2":         lastSpaceValue(disk.DiskSpaceAllocatedCurrentL2),
						"disk_space_free":                         lastSpaceValue(disk.DiskSpaceFree),
						"disk_space_allocated_percent":            lastPercentValue(disk.DiskSpaceAllocatedPercentage),
						"disk_space_allocated_percentage_current": lastPercentValue(disk.DiskSpaceAllocatedPercentageCurrent),
					},
				},
				RootFields: createECSFields(m),
			}
			addSpaceSummaryValues(disk, &event, field_prefix)
			addPercentageSummaryValues(disk, &event, field_prefix)
			events = append(events, event)
		}
	}
	return events, nil
}

func getCapacityEvents(m *MetricSet) ([]mb.Event, error) {
	field_prefix := "capacity"
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
		return nil, fmt.Errorf("failed to unmarshal capacity data: %v", err)
	}

	event := mb.Event{
		Timestamp: timestamp,
		MetricSetFields: map[string]interface{}{
			field_prefix + ".total_provisioned_gb": storage.TotalProvisionedGB,
			field_prefix + ".total_free_gb":        storage.TotalFreeGB,
		},
		RootFields: createECSFields(m),
	}
	events = append(events, event)
	return events, nil
}

func getStoragePoolsEvents(m *MetricSet) ([]mb.Event, error) {
	field_prefix := "storage_pool"
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
				field_prefix: mapstr.M{
					"chunks_l1_journal_total_size":                  pool.ChunksL1JournalTotalSize,
					"chunks_l1_btree_number":                        pool.ChunksL1BtreeNumber,
					"chunks_l1_btree_total_size":                    pool.ChunksL1BtreeTotalSize,
					"chunks_l1_journal_avg_size":                    pool.ChunksL1JournalAvgSize,
					"chunks_l1_journal_number":                      pool.ChunksL1JournalNumber,
					"chunks_l0_journal_total_size":                  pool.ChunksL0JournalTotalSize,
					"chunks_l0_btree_number":                        pool.ChunksL0BtreeNumber,
					"chunks_l0_btree_total_size":                    pool.ChunksL0BtreeTotalSize,
					"chunks_l0_btree_avg_size":                      pool.ChunksL0BtreeAvgSize,
					"chunks_l0_journal_avg_size":                    pool.ChunksL0JournalAvgSize,
					"chunks_l0_journal_number":                      pool.ChunksL0JournalNumber,
					"chunks_repo_total_seal_size":                   pool.ChunksRepoTotalSealSize,
					"chunks_repo_number":                            pool.ChunksRepoNumber,
					"chunks_xor_number":                             pool.ChunksXorNumber,
					"chunks_xor_total_size":                         pool.ChunksXorTotalSize,
					"chunks_geo_cache_total_size":                   pool.ChunksGeoCacheTotalSize,
					"chunks_geo_cache_count":                        pool.ChunksGeoCacheCount,
					"chunks_geo_copy_number":                        pool.ChunksGeoCopyNumber,
					"chunks_geo_copy_total_size":                    pool.ChunksGeoCopyTotalSize,
					"num_nodes":                                     pool.NumNodes,
					"num_disks":                                     pool.NumDisks,
					"num_bad_nodes":                                 pool.NumBadNodes,
					"num_good_nodes":                                pool.NumGoodNodes,
					"num_bad_disks":                                 pool.NumBadDisks,
					"num_good_disks":                                pool.NumGoodDisks,
					"num_maintenance_nodes":                         pool.NumMaintenanceNodes,
					"num_maintenance_disks":                         pool.NumMaintenanceDisks,
					"num_ready_to_replace_disks":                    pool.NumReadyToReplaceDisks,
					"num_nodes_with_sufficient_disk_space":          pool.NumNodesWithSufficientDiskSpace,
					"gc_user_data_is_enabled":                       pool.GcUserDataIsEnabled,
					"gc_system_metadata_is_enabled":                 pool.GcSystemMetadataIsEnabled,
					"recovery_complete_time_estimate":               pool.RecoveryCompleteTimeEstimate,
					"chunks_ec_complete_time_estimate":              pool.ChunksEcCompleteTimeEstimate,
					"pool_id":                                       pool.ID,
					"name":                                          pool.Name,
					"gc_user_unreclaimable_current":                 lastCapacityValue(pool.GcUserUnreclaimableCurrent),
					"gc_user_total_detected_current":                lastCapacityValue(pool.GcUserTotalDetectedCurrent),
					"gc_system_reclaimed_current":                   lastCapacityValue(pool.GcSystemReclaimedCurrent),
					"gc_system_reclaimed_per_interval":              lastCapacityValue(pool.GcSystemReclaimedPerInterval),
					"gc_system_reclaimed_over_time_range":           lastCapacityValue(pool.GcSystemReclaimedOverTimeRange),
					"gc_combined_reclaimed_current":                 lastCapacityValue(pool.GcCombinedReclaimedCurrent),
					"gc_user_reclaimed_over_time_range":             lastCapacityValue(pool.GcUserReclaimedOverTimeRange),
					"gc_system_pending_current":                     lastCapacityValue(pool.GcSystemPendingCurrent),
					"gc_user_reclaimed_per_interval":                lastCapacityValue(pool.GcUserReclaimedPerInterval),
					"gc_combined_total_detected_current":            lastCapacityValue(pool.GcCombinedTotalDetectedCurrent),
					"gc_system_total_detected_current":              lastCapacityValue(pool.GcSystemTotalDetectedCurrent),
					"gc_user_reclaimed_current":                     lastCapacityValue(pool.GcUserReclaimedCurrent),
					"gc_user_pending_current":                       lastCapacityValue(pool.GcUserPendingCurrent),
					"gc_system_unreclaimable_current":               lastCapacityValue(pool.GcSystemUnreclaimableCurrent),
					"gc_combined_reclaimed_over_time_range":         lastCapacityValue(pool.GcCombinedReclaimedOverTimeRange),
					"gc_combined_pending_current":                   lastCapacityValue(pool.GcCombinedPendingCurrent),
					"gc_combined_unreclaimable_current":             lastCapacityValue(pool.GcCombinedUnreclaimableCurrent),
					"disk_space_allocated_geo_cache_current":        lastCapacityValue(pool.DiskSpaceAllocatedGeoCacheCurrent),
					"disk_space_allocated_local_protection_current": lastCapacityValue(pool.DiskSpaceAllocatedLocalProtectionCurrent),
					"disk_space_allocated_system_metadata_current":  lastCapacityValue(pool.DiskSpaceAllocatedSystemMetadataCurrent),
					"disk_space_allocated_user_data_current":        lastCapacityValue(pool.DiskSpaceAllocatedUserDataCurrent),
					"allocated_capacity_forecast":                   lastCapacityValue(pool.AllocatedCapacityForecast),
					"disk_space_allocated_percentage":               lastPercentValue(pool.DiskSpaceAllocatedPercentage),
					"chunks_ec_coded_ratio_current":                 lastPercentValue(pool.ChunksEcCodedRatioCurrent),
					"chunks_ec_coded_ratio":                         lastPercentValue(pool.ChunksEcCodedRatio),
					"disk_space_allocated_percentage_current":       lastPercentValue(pool.DiskSpaceAllocatedPercentageCurrent),
				},
			},
			RootFields: createECSFields(m),
		}

		addSpaceSummaryValues(pool, &event, field_prefix)
		addPercentageSummaryValues(pool, &event, field_prefix)
		events = append(events, event)
	}

	return events, nil

}

func getReplicationGroupsEvents(m *MetricSet) ([]mb.Event, error) {
	field_prefix := "rep_group"
	timestamp := time.Now().UTC()
	client := m.ecsClient
	endpoint, err := getEndpoint("Replication Groups")
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
		return nil, fmt.Errorf("failed to unmarshal ReplicationGroup response: %w", err)

	}

	var events []mb.Event
	for _, rawGroup := range response.Embedded.Instances {
		var group ReplicationGroup
		err := json.Unmarshal(rawGroup, &group)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal Replication Group instance: %w", err)
		}

		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				field_prefix: mapstr.M{
					"name":                          group.Name,
					"group_id":                      group.ID,
					"num_zones":                     group.NumZones,
					"chunks_pending_xor_total_size": group.ChunksPendingXorTotalSize,
					"chunks_repo_pending_replication_total_size":    group.ChunksRepoPendingReplicationTotalSize,
					"chunks_journal_pending_replication_total_size": group.ChunksJournalPendingReplicationTotalSize,
					"replication_egress_traffic":                    lastBandwidthValue(group.ReplicationEgressTraffic),
					"replication_ingress_traffic":                   lastBandwidthValue(group.ReplicationIngressTraffic),
					"replication_ingress_traffic_current":           lastBandwidthValue(group.ReplicationIngressTrafficCurrent),
					"replication_egress_traffic_current":            lastBandwidthValue(group.ReplicationEgressTrafficCurrent),
				},
			},
			RootFields: createECSFields(m),
		}

		addSpaceSummaryValues(group, &event, field_prefix)
		addPercentageSummaryValues(group, &event, field_prefix)
		addTrafficSummaryValues(group, &event, field_prefix)
		events = append(events, event)
	}

	return events, nil
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

func lastSpaceValue(spaces []TimestampedSpace) int64 {
	if len(spaces) == 0 {
		return -1
	}
	return spaces[len(spaces)-1].Space
}

func lastPercentValue(spaces []TimestampedPercent) int64 {
	if len(spaces) == 0 {
		return -1
	}
	return int64(spaces[len(spaces)-1].Percent)
}

func lastCapacityValue(spaces []TimestampedCapacity) int64 {
	if len(spaces) == 0 {
		return -1
	}
	return spaces[len(spaces)-1].Capacity
}

func lastBandwidthValue(spaces []TimestampedBandwidth) int64 {
	if len(spaces) == 0 {
		return -1
	}
	return spaces[len(spaces)-1].Bandwidth
}

func lastBytesValue(tbytes []TimestampedBytes) int64 {
	if len(tbytes) == 0 {
		return -1
	}
	return tbytes[len(tbytes)-1].Bytes
}

func lastCountValue(count []TimestampedCount) int64 {
	if len(count) == 0 {
		return -1
	}
	return count[len(count)-1].Count
}

func lastDummyValue(timestampedDummy []TimestampedDummy) int {
	if len(timestampedDummy) == 0 {
		return -1
	}
	return int(timestampedDummy[len(timestampedDummy)-1].T)
}

func addSpaceSummaryValues(data interface{}, event *mb.Event, prefix string) {
	var summaries map[string]Summary

	switch v := data.(type) {
	case NodeData:
		summaries = map[string]Summary{
			"disk_space_free_l1_summary":      v.DiskSpaceFreeL1Summary,
			"disk_space_allocated_l2_summary": v.DiskSpaceAllocatedL2Summary,
			"disk_space_total_summary":        v.DiskSpaceTotalSummary,
			"disk_space_free_l2_summary":      v.DiskSpaceFreeL2Summary,
			"disk_space_allocated_summary":    v.DiskSpaceAllocatedSummary,
			"disk_space_free_summary":         v.DiskSpaceFreeSummary,
			"disk_space_allocated_l1_summary": v.DiskSpaceAllocatedL1Summary,
		}
	case StoragePoolData:
		summaries = map[string]Summary{
			"chunks_ec_coded_total_seal_size_summary":      v.ChunksEcCodedTotalSealSizeSummary,
			"chunks_ec_rate_summary":                       v.ChunksEcRateSummary,
			"chunks_ec_coded_ratio_summary":                v.ChunksEcCodedRatioSummary,
			"chunks_ec_applicable_total_seal_size_summary": v.ChunksEcApplicableTotalSealSizeSummary,
			"recovery_rate_summary":                        v.RecoveryRateSummary,
			"recovery_bad_chunks_total_size_summary":       v.RecoveryBadChunksTotalSizeSummary,
			"disk_space_allocated_l1_summary":              v.DiskSpaceAllocatedL1Summary,
			"disk_space_allocated_l2_summary":              v.DiskSpaceAllocatedL2Summary,
			"disk_space_allocated_summary":                 v.DiskSpaceAllocatedSummary,
			"disk_space_free_l1_summary":                   v.DiskSpaceFreeL1Summary,
			"disk_space_free_l2_summary":                   v.DiskSpaceFreeL2Summary,
			"disk_space_total_summary":                     v.DiskSpaceTotalSummary,
			"disk_space_free_summary":                      v.DiskSpaceFreeSummary,
		}
	case ReplicationGroup:
		summaries = map[string]Summary{}
	case DiskData:
		summaries = map[string]Summary{
			"disk_space_free_l1_summary":      v.DiskSpaceFreeL1Summary,
			"disk_space_allocated_l2_summary": v.DiskSpaceAllocatedL2Summary,
			"disk_space_total_summary":        v.DiskSpaceTotalSummary,
			"disk_space_free_l2_summary":      v.DiskSpaceFreeL2Summary,
			"disk_space_allocated_summary":    v.DiskSpaceAllocatedSummary,
			"disk_space_free_summary":         v.DiskSpaceFreeSummary,
			"disk_space_allocated_l1_summary": v.DiskSpaceAllocatedL1Summary,
		}
	}

	prefixMap, ok := event.MetricSetFields[prefix].(mapstr.M)
	if !ok {
		prefixMap = mapstr.M{}
		event.MetricSetFields[prefix] = prefixMap
	}

	for key, summary := range summaries {
		prefixMap[key] = mapstr.M{}
		if prefixMapKey, ok := prefixMap[key].(mapstr.M); ok {
			if len(summary.Min) > 0 {
				prefixMapKey["min"] = summary.Min[0].Space
			}
			if len(summary.Max) > 0 {
				prefixMapKey["max"] = summary.Max[0].Space
			}
			prefixMapKey["avg"] = summary.Avg
		}
	}
}

func addPercentageSummaryValues(data interface{}, event *mb.Event, prefix string) {
	var summaries map[string]PercentageSummary
	switch v := data.(type) {
	case NodeData:
		summaries = map[string]PercentageSummary{
			"disk_space_allocated_percentage_summary": v.DiskSpaceAllocatedPercentageSummary,
		}
	case StoragePoolData:
		summaries = map[string]PercentageSummary{
			"disk_space_allocated_percentage_summary": v.DiskSpaceAllocatedPercentageSummary,
		}
	}

	prefixMap, ok := event.MetricSetFields[prefix].(mapstr.M)
	if !ok {
		prefixMap = mapstr.M{}
		event.MetricSetFields[prefix] = prefixMap
	}

	for key, summary := range summaries {
		prefixMap[key] = mapstr.M{}
		if prefixMapKey, ok := prefixMap[key].(mapstr.M); ok {
			if len(summary.Min) > 0 {
				prefixMapKey["min"] = summary.Min[0].Percent
			}
			if len(summary.Max) > 0 {
				prefixMapKey["max"] = summary.Max[0].Percent
			}
			prefixMapKey["avg"] = summary.Avg
		}
	}
}

func addTrafficSummaryValues(data interface{}, event *mb.Event, prefix string) {
	var summaries map[string]TrafficSummary
	switch v := data.(type) {
	case ReplicationGroup:
		summaries = map[string]TrafficSummary{
			"replication_ingress_traffic_summary": v.ReplicationIngressTrafficSummary,
			"replication_egress_traffic_summary":  v.ReplicationEgressTrafficSummary,
		}
	}

	prefixMap, ok := event.MetricSetFields[prefix].(mapstr.M)
	if !ok {
		prefixMap = mapstr.M{}
		event.MetricSetFields[prefix] = prefixMap
	}

	for key, summary := range summaries {

		prefixMap[key] = mapstr.M{}
		if prefixMapKey, ok := prefixMap[key].(mapstr.M); ok {
			if len(summary.Min) > 0 {
				prefixMapKey["min"] = summary.Min[0].Bandwidth
			}
			if len(summary.Max) > 0 {
				prefixMapKey["max"] = summary.Max[0].Bandwidth
			}
			prefixMapKey["avg"] = summary.Avg
		}
	}
}

func createECSFields(ms *MetricSet) mapstr.M {
	dataset := fmt.Sprintf("%s.%s", ms.Module().Name(), ms.Name())

	return mapstr.M{
		"event": mapstr.M{
			"dataset": dataset,
		},
		"observer": mapstr.M{
			"hostname": ms.config.HostInfo.Hostname,
			"ip":       ms.config.HostInfo.IP,
			"type":     "cloud-storage",
			"vendor":   "Dell",
		},
	}
}
