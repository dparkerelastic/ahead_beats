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
	"github.com/elastic/beats/v7/metricbeat/module/dellecs"
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
				field_prefix + ".process_name":             proc.ProcessName,
				field_prefix + ".process_id":               proc.ID,
				field_prefix + ".cpu_utilization":          lastPercentValue(proc.CPUUtilization),
				field_prefix + ".java_heap_utilization":    lastPercentValue(proc.JavaHeapUtilization),
				field_prefix + ".max_java_heap_size":       lastBytesValue(proc.MaxJavaHeapSize),
				field_prefix + ".memory_utilization_bytes": lastBytesValue(proc.MemoryUtilizationBytes),
				field_prefix + ".memory_utilization":       lastPercentValue(proc.MemoryUtilization),
				field_prefix + ".thread_count":             lastCountValue(proc.ThreadCount),
				field_prefix + ".restart_time":             lastDummyValue(proc.RestartTime),
			},
			RootFields: dellecs.MakeRootFields(m.config),
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
				field_prefix + ".id":                                      nodeDetails.ID,
				field_prefix + ".storage_pool_id":                         nodeDetails.StoragePoolID,
				field_prefix + ".name":                                    nodeDetails.DisplayName,
				field_prefix + ".api_change":                              nodeDetails.APIChange,
				field_prefix + ".num_bad_disks":                           nodeDetails.NumBadDisks,
				field_prefix + ".storage_pool_name":                       nodeDetails.StoragePoolName,
				field_prefix + ".display_name":                            nodeDetails.DisplayName,
				field_prefix + ".num_ready_to_replace_disks":              nodeDetails.NumReadyToReplaceDisks,
				field_prefix + ".num_maintenance_disks":                   nodeDetails.NumMaintenanceDisks,
				field_prefix + ".health_status":                           nodeDetails.HealthStatus,
				field_prefix + ".num_good_disks":                          nodeDetails.NumGoodDisks,
				field_prefix + ".num_disks":                               nodeDetails.NumDisks,
				field_prefix + ".disk_space_free_current_l1":              lastSpaceValue(node.DiskSpaceFreeCurrentL1),
				field_prefix + ".allocated_capacity_forecast":             lastSpaceValue(node.AllocatedCapacityForecast),
				field_prefix + ".disk_space_free_current_l2":              lastSpaceValue(node.DiskSpaceFreeCurrentL2),
				field_prefix + ".disk_space_reserved_current":             lastSpaceValue(node.DiskSpaceReservedCurrent),
				field_prefix + ".disk_space_free_l2":                      lastSpaceValue(node.DiskSpaceFreeL2),
				field_prefix + ".disk_space_free_l1":                      lastSpaceValue(node.DiskSpaceFreeL1),
				field_prefix + ".disk_space_total":                        lastSpaceValue(node.DiskSpaceTotal),
				field_prefix + ".disk_space_total_current_l1":             lastSpaceValue(node.DiskSpaceTotalCurrentL1),
				field_prefix + ".disk_space_total_current_l2":             lastSpaceValue(node.DiskSpaceTotalCurrentL2),
				field_prefix + ".disk_space_offline_total_current":        lastSpaceValue(node.DiskSpaceOfflineTotalCurrent),
				field_prefix + ".disk_space_allocated_l2":                 lastSpaceValue(node.DiskSpaceAllocatedL2),
				field_prefix + ".disk_space_allocated_l1":                 lastSpaceValue(node.DiskSpaceAllocatedL1),
				field_prefix + ".disk_space_free_current":                 lastSpaceValue(node.DiskSpaceFreeCurrent),
				field_prefix + ".disk_space_total_current":                lastSpaceValue(node.DiskSpaceTotalCurrent),
				field_prefix + ".disk_space_allocated_current":            lastSpaceValue(node.DiskSpaceAllocatedCurrent),
				field_prefix + ".disk_space_allocated_current_l1":         lastSpaceValue(node.DiskSpaceAllocatedCurrentL1),
				field_prefix + ".disk_space_allocated_current_l2":         lastSpaceValue(node.DiskSpaceAllocatedCurrentL2),
				field_prefix + ".disk_space_allocated":                    lastSpaceValue(node.DiskSpaceAllocated),
				field_prefix + ".disk_space_allocated_percent":            lastPercentValue(node.DiskSpaceAllocatedPercentage),
				field_prefix + ".disk_space_allocated_percentage_current": lastPercentValue(node.DiskSpaceAllocatedPercentageCurrent),
			},
			RootFields: dellecs.MakeRootFields(m.config),
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
					field_prefix + ".node_id":                                 node.ID,
					field_prefix + ".storage_pool_name":                       disk.StoragePoolName,
					field_prefix + ".display_name":                            disk.DisplayName,
					field_prefix + ".node_display_name":                       disk.NodeDisplayName,
					field_prefix + ".slot_id":                                 disk.SlotId,
					field_prefix + ".id":                                      disk.Id,
					field_prefix + ".storage_pool_id":                         disk.StoragePoolId,
					field_prefix + ".ssm_l2_status":                           disk.SsmL2Status,
					field_prefix + ".health_status":                           disk.HealthStatus,
					field_prefix + ".ssm_l1_status":                           disk.SsmL1Status,
					field_prefix + ".disk_space_free_current_l1":              lastSpaceValue(disk.DiskSpaceFreeCurrentL1),
					field_prefix + ".disk_space_free_current_l2":              lastSpaceValue(disk.DiskSpaceFreeCurrentL2),
					field_prefix + ".disk_space_free_l2":                      lastSpaceValue(disk.DiskSpaceFreeL2),
					field_prefix + ".disk_space_free_l1":                      lastSpaceValue(disk.DiskSpaceFreeL1),
					field_prefix + ".disk_space_total":                        lastSpaceValue(disk.DiskSpaceTotal),
					field_prefix + ".disk_space_total_current_l1":             lastSpaceValue(disk.DiskSpaceTotalCurrentL1),
					field_prefix + ".disk_space_total_current_l2":             lastSpaceValue(disk.DiskSpaceTotalCurrentL2),
					field_prefix + ".disk_space_allocated_l2":                 lastSpaceValue(disk.DiskSpaceAllocatedL2),
					field_prefix + ".disk_space_allocated_l1":                 lastSpaceValue(disk.DiskSpaceAllocatedL1),
					field_prefix + ".disk_space_free_current":                 lastSpaceValue(disk.DiskSpaceFreeCurrent),
					field_prefix + ".disk_space_total_current":                lastSpaceValue(disk.DiskSpaceTotalCurrent),
					field_prefix + ".disk_space_allocated_current":            lastSpaceValue(disk.DiskSpaceAllocatedCurrent),
					field_prefix + ".disk_space_allocated_current_l1":         lastSpaceValue(disk.DiskSpaceAllocatedCurrentL1),
					field_prefix + ".disk_space_allocated":                    lastSpaceValue(disk.DiskSpaceAllocated),
					field_prefix + ".disk_space_allocated_current_l2":         lastSpaceValue(disk.DiskSpaceAllocatedCurrentL2),
					field_prefix + ".disk_space_free":                         lastSpaceValue(disk.DiskSpaceFree),
					field_prefix + ".disk_space_allocated_percent":            lastPercentValue(disk.DiskSpaceAllocatedPercentage),
					field_prefix + ".disk_space_allocated_percentage_current": lastPercentValue(disk.DiskSpaceAllocatedPercentageCurrent),
				},
				RootFields: dellecs.MakeRootFields(m.config),
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
			field_prefix + ".total_provisioned_db": storage.TotalProvisionedGB,
			field_prefix + ".total_free_db":        storage.TotalFreeGB,
		},
		RootFields: dellecs.MakeRootFields(m.config),
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
				field_prefix + ".chunks_l1_journal_total_size":                  pool.ChunksL1JournalTotalSize,
				field_prefix + ".chunks_l1_btree_number":                        pool.ChunksL1BtreeNumber,
				field_prefix + ".chunks_l1_btree_total_size":                    pool.ChunksL1BtreeTotalSize,
				field_prefix + ".chunks_l1_journal_avg_size":                    pool.ChunksL1JournalAvgSize,
				field_prefix + ".chunks_l1_journal_number":                      pool.ChunksL1JournalNumber,
				field_prefix + ".chunks_l0_journal_total_size":                  pool.ChunksL0JournalTotalSize,
				field_prefix + ".chunks_l0_btree_number":                        pool.ChunksL0BtreeNumber,
				field_prefix + ".chunks_l0_btree_total_size":                    pool.ChunksL0BtreeTotalSize,
				field_prefix + ".chunks_l0_btree_avg_size":                      pool.ChunksL0BtreeAvgSize,
				field_prefix + ".chunks_l0_journal_avg_size":                    pool.ChunksL0JournalAvgSize,
				field_prefix + ".chunks_l0_journal_number":                      pool.ChunksL0JournalNumber,
				field_prefix + ".chunks_repo_total_seal_size":                   pool.ChunksRepoTotalSealSize,
				field_prefix + ".chunks_repo_number":                            pool.ChunksRepoNumber,
				field_prefix + ".chunks_xor_number":                             pool.ChunksXorNumber,
				field_prefix + ".chunks_xor_total_size":                         pool.ChunksXorTotalSize,
				field_prefix + ".chunks_geo_cache_total_size":                   pool.ChunksGeoCacheTotalSize,
				field_prefix + ".chunks_geo_cache_count":                        pool.ChunksGeoCacheCount,
				field_prefix + ".chunks_geo_copy_number":                        pool.ChunksGeoCopyNumber,
				field_prefix + ".chunks_geo_copy_total_size":                    pool.ChunksGeoCopyTotalSize,
				field_prefix + ".num_nodes":                                     pool.NumNodes,
				field_prefix + ".num_disks":                                     pool.NumDisks,
				field_prefix + ".num_bad_nodes":                                 pool.NumBadNodes,
				field_prefix + ".num_good_nodes":                                pool.NumGoodNodes,
				field_prefix + ".num_bad_disks":                                 pool.NumBadDisks,
				field_prefix + ".num_good_disks":                                pool.NumGoodDisks,
				field_prefix + ".num_maintenance_nodes":                         pool.NumMaintenanceNodes,
				field_prefix + ".num_maintenance_disks":                         pool.NumMaintenanceDisks,
				field_prefix + ".num_ready_to_replace_disks":                    pool.NumReadyToReplaceDisks,
				field_prefix + ".num_nodes_with_sufficient_disk_space":          pool.NumNodesWithSufficientDiskSpace,
				field_prefix + ".gc_user_data_is_enabled":                       pool.GcUserDataIsEnabled,
				field_prefix + ".gc_system_metadata_is_enabled":                 pool.GcSystemMetadataIsEnabled,
				field_prefix + ".recovery_complete_time_estimate":               pool.RecoveryCompleteTimeEstimate,
				field_prefix + ".chunks_ec_complete_time_estimate":              pool.ChunksEcCompleteTimeEstimate,
				field_prefix + ".id":                                            pool.ID,
				field_prefix + ".name":                                          pool.Name,
				field_prefix + ".gc_user_unreclaimable_current":                 lastCapacityValue(pool.GcUserUnreclaimableCurrent),
				field_prefix + ".gc_user_total_detected_current":                lastCapacityValue(pool.GcUserTotalDetectedCurrent),
				field_prefix + ".gc_system_reclaimed_current":                   lastCapacityValue(pool.GcSystemReclaimedCurrent),
				field_prefix + ".gc_system_reclaimed_per_interval":              lastCapacityValue(pool.GcSystemReclaimedPerInterval),
				field_prefix + ".gc_system_reclaimed_over_time_range":           lastCapacityValue(pool.GcSystemReclaimedOverTimeRange),
				field_prefix + ".gc_combined_reclaimed_current":                 lastCapacityValue(pool.GcCombinedReclaimedCurrent),
				field_prefix + ".gc_user_reclaimed_over_time_range":             lastCapacityValue(pool.GcUserReclaimedOverTimeRange),
				field_prefix + ".gc_system_pending_current":                     lastCapacityValue(pool.GcSystemPendingCurrent),
				field_prefix + ".gc_user_reclaimed_per_interval":                lastCapacityValue(pool.GcUserReclaimedPerInterval),
				field_prefix + ".gc_combined_total_detected_current":            lastCapacityValue(pool.GcCombinedTotalDetectedCurrent),
				field_prefix + ".gc_system_total_detected_current":              lastCapacityValue(pool.GcSystemTotalDetectedCurrent),
				field_prefix + ".gc_user_reclaimed_current":                     lastCapacityValue(pool.GcUserReclaimedCurrent),
				field_prefix + ".gc_user_pending_current":                       lastCapacityValue(pool.GcUserPendingCurrent),
				field_prefix + ".gc_system_unreclaimable_current":               lastCapacityValue(pool.GcSystemUnreclaimableCurrent),
				field_prefix + ".gc_combined_reclaimed_over_time_range":         lastCapacityValue(pool.GcCombinedReclaimedOverTimeRange),
				field_prefix + ".gc_combined_pending_current":                   lastCapacityValue(pool.GcCombinedPendingCurrent),
				field_prefix + ".gc_combined_unreclaimable_current":             lastCapacityValue(pool.GcCombinedUnreclaimableCurrent),
				field_prefix + ".disk_space_allocated_geo_cache_current":        lastCapacityValue(pool.DiskSpaceAllocatedGeoCacheCurrent),
				field_prefix + ".disk_space_allocated_local_protection_current": lastCapacityValue(pool.DiskSpaceAllocatedLocalProtectionCurrent),
				field_prefix + ".disk_space_allocated_system_metadata_current":  lastCapacityValue(pool.DiskSpaceAllocatedSystemMetadataCurrent),
				field_prefix + ".disk_space_allocated_user_data_current":        lastCapacityValue(pool.DiskSpaceAllocatedUserDataCurrent),
				field_prefix + ".allocated_capacity_forecast":                   lastCapacityValue(pool.AllocatedCapacityForecast),
				field_prefix + ".disk_space_allocated_percentage":               lastPercentValue(pool.DiskSpaceAllocatedPercentage),
				field_prefix + ".chunks_ec_coded_ratio_current":                 lastPercentValue(pool.ChunksEcCodedRatioCurrent),
				field_prefix + ".chunks_ec_coded_ratio":                         lastPercentValue(pool.ChunksEcCodedRatio),
				field_prefix + ".disk_space_allocated_percentage_current":       lastPercentValue(pool.DiskSpaceAllocatedPercentageCurrent),
			},
			RootFields: dellecs.MakeRootFields(m.config),
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
				field_prefix + ".name":                                          group.Name,
				field_prefix + ".id":                                            group.ID,
				field_prefix + ".num_zones":                                     group.NumZones,
				field_prefix + ".chunks_pending_xor_total_size":                 group.ChunksPendingXorTotalSize,
				field_prefix + ".chunks_repo_pending_replication_total_size":    group.ChunksRepoPendingReplicationTotalSize,
				field_prefix + ".chunks_journal_pending_replication_total_size": group.ChunksJournalPendingReplicationTotalSize,
				field_prefix + ".replication_egress_traffic":                    lastBandwidthValue(group.ReplicationEgressTraffic),
				field_prefix + ".replication_ingress_traffic":                   lastBandwidthValue(group.ReplicationIngressTraffic),
				field_prefix + ".replication_ingress_traffic_current":           lastBandwidthValue(group.ReplicationIngressTrafficCurrent),
				field_prefix + ".replication_egress_traffic_current":            lastBandwidthValue(group.ReplicationEgressTrafficCurrent),
			},
			RootFields: dellecs.MakeRootFields(m.config),
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
			prefix + ".chunks_ec_coded_total_seal_size_summary":      v.ChunksEcCodedTotalSealSizeSummary,
			prefix + ".chunks_ec_rate_summary":                       v.ChunksEcRateSummary,
			prefix + ".chunks_ec_coded_ratio_summary":                v.ChunksEcCodedRatioSummary,
			prefix + ".chunks_ec_applicable_total_seal_size_summary": v.ChunksEcApplicableTotalSealSizeSummary,
			prefix + ".recovery_rate_summary":                        v.RecoveryRateSummary,
			prefix + ".recovery_bad_chunks_total_size_summary":       v.RecoveryBadChunksTotalSizeSummary,
			prefix + ".disk_space_allocated_l1_summary":              v.DiskSpaceAllocatedL1Summary,
			prefix + ".disk_space_allocated_l2_summary":              v.DiskSpaceAllocatedL2Summary,
			prefix + ".disk_space_allocated_summary":                 v.DiskSpaceAllocatedSummary,

			prefix + ".disk_space_free_l1_summary": v.DiskSpaceFreeL1Summary,
			prefix + ".disk_space_free_l2_summary": v.DiskSpaceFreeL2Summary,
			prefix + ".disk_space_total_summary":   v.DiskSpaceTotalSummary,
			prefix + ".disk_space_free_summary":    v.DiskSpaceFreeSummary,
		}
	case ReplicationGroup:
		summaries = map[string]Summary{}
	}

	for key, summary := range summaries {
		event.MetricSetFields[key+".min"] = summary.Min[0].Space
		event.MetricSetFields[key+".max"] = summary.Max[0].Space
		event.MetricSetFields[key+".avg"] = summary.Avg
	}
}

func addPercentageSummaryValues(data interface{}, event *mb.Event, prefix string) {
	var summaries map[string]PercentageSummary
	switch v := data.(type) {
	case NodeData:
		summaries = map[string]PercentageSummary{
			prefix + ".disk_space_allocated_percentage_summary": v.DiskSpaceAllocatedPercentageSummary,
		}
	case StoragePoolData:
		summaries = map[string]PercentageSummary{
			prefix + ".disk_space_allocated_percentage_summary": v.DiskSpaceAllocatedPercentageSummary,
		}
	}

	for key, summary := range summaries {
		event.MetricSetFields[key+".min"] = summary.Min[0].Percent
		event.MetricSetFields[key+".max"] = summary.Max[0].Percent
		event.MetricSetFields[key+".avg"] = summary.Avg
	}
}

func addTrafficSummaryValues(data interface{}, event *mb.Event, prefix string) {
	var summaries map[string]TrafficSummary
	switch v := data.(type) {
	case ReplicationGroup:
		summaries = map[string]TrafficSummary{
			prefix + "replication_ingress_traffic_summary": v.ReplicationIngressTrafficSummary,
			prefix + "replication_egress_traffic_summary":  v.ReplicationEgressTrafficSummary,
		}
	}

	for key, summary := range summaries {
		event.MetricSetFields[key+".min"] = summary.Min[0].Bandwidth
		event.MetricSetFields[key+".max"] = summary.Max[0].Bandwidth
		event.MetricSetFields[key+".avg"] = summary.Avg
	}
}
