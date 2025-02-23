package health

import (
	"encoding/json"
)

type Response struct {
	APIChange string   `json:"apiChange"`
	Title     string   `json:"title"`
	Links     Links    `json:"_links"`
	Embedded  Embedded `json:"_embedded"`
}

type Embedded struct {
	Instances []json.RawMessage `json:"_instances"`
}

// Link represents an individual link with an href.
type Link struct {
	Href string `json:"href"`
}

// Links represents the collection of links.
type Links struct {
	Self      Link `json:"self"`
	Disks     Link `json:"disks"`
	Processes Link `json:"processes"`
}

type Summary struct {
	Min []TimestampedSpace `json:"Min"`
	Max []TimestampedSpace `json:"Max"`
	Avg float64            `json:"Avg,string"`
}

type StorageInfo struct {
	TotalProvisionedGB int `json:"totalProvisioned_gb"`
	TotalFreeGB        int `json:"totalFree_gb"`
}

type TimestampedRate struct {
	T    int64   `json:"t,string"`
	Rate float64 `json:"Rate,string"`
}

type TimestampedCapacity struct {
	T        int64 `json:"t,string"`
	Capacity int64 `json:"Capacity,string"`
}

type TimestampedSpace struct {
	T     int64 `json:"t,string"`
	Space int64 `json:"Space,string"`
}

type TimestampedPercent struct {
	T       int64   `json:"t,string"`
	Percent float64 `json:"Percent,string"`
}

type DiskSpaceSummary struct {
	Min []TimestampedSpace `json:"Min"`
	Max []TimestampedSpace `json:"Max"`
	Avg float64            `json:"Avg,string"`
}

type StoragePoolData struct {
	Links                                    Links                 `json:"_links"`
	ChunksL1JournalTotalSize                 int64                 `json:"chunksL1JournalTotalSize,string"`
	RecoveryRateCurrent                      []TimestampedRate     `json:"recoveryRateCurrent"`
	ChunksRepoTotalSealSize                  int64                 `json:"chunksRepoTotalSealSize,string"`
	ChunksL1BtreeNumber                      int64                 `json:"chunksL1BtreeNumber,string"`
	GcUserUnreclaimableCurrent               []TimestampedCapacity `json:"gcUserUnreclaimableCurrent"`
	GcUserTotalDetectedCurrent               []TimestampedCapacity `json:"gcUserTotalDetectedCurrent"`
	GcSystemReclaimedCurrent                 []TimestampedCapacity `json:"gcSystemReclaimedCurrent"`
	NumMaintenanceDisks                      int64                 `json:"numMaintenanceDisks,string"`
	ChunksGeoCopyNumber                      int64                 `json:"chunksGeoCopyNumber,string"`
	ChunksEcApplicableTotalSealSizeCurrent   []TimestampedSpace    `json:"chunksEcApplicableTotalSealSizeCurrent"`
	ChunksEcCodedTotalSealSizeSummary        Summary               `json:"chunksEcCodedTotalSealSizeSummary"`
	GcUserDataIsEnabled                      bool                  `json:"gcUserDataIsEnabled,string"`
	ID                                       string                `json:"id"`
	GcSystemReclaimedPerInterval             []TimestampedCapacity `json:"gcSystemReclaimedPerInterval"`
	ChunksEcRateSummary                      Summary               `json:"chunksEcRateSummary"`
	ChunksEcCodedTotalSealSizeCurrent        []TimestampedSpace    `json:"chunksEcCodedTotalSealSizeCurrent"`
	DiskSpaceAllocatedL2                     []TimestampedSpace    `json:"diskSpaceAllocatedL2"`
	GcSystemReclaimedOverTimeRange           []TimestampedCapacity `json:"gcSystemReclaimedOverTimeRange"`
	DiskSpaceAllocatedL1                     []TimestampedSpace    `json:"diskSpaceAllocatedL1"`
	GcCombinedReclaimedCurrent               []TimestampedCapacity `json:"gcCombinedReclaimedCurrent"`
	GcUserReclaimedOverTimeRange             []TimestampedCapacity `json:"gcUserReclaimedOverTimeRange"`
	ChunksL1BtreeTotalSize                   int64                 `json:"chunksL1BtreeTotalSize,string"`
	ChunksL0JournalAvgSize                   float64               `json:"chunksL0JournalAvgSize,string"`
	DiskSpaceAllocatedCurrent                []TimestampedSpace    `json:"diskSpaceAllocatedCurrent"`
	NumBadNodes                              int64                 `json:"numBadNodes,string"`
	DiskSpaceFreeSummary                     DiskSpaceSummary      `json:"diskSpaceFreeSummary"`
	RecoveryBadChunksTotalSizeCurrent        []TimestampedSpace    `json:"recoveryBadChunksTotalSizeCurrent"`
	DiskSpaceAllocated                       []TimestampedSpace    `json:"diskSpaceAllocated"`
	DiskSpaceFreeCurrentL1                   []TimestampedSpace    `json:"diskSpaceFreeCurrentL1"`
	DiskSpaceFreeCurrentL2                   []TimestampedSpace    `json:"diskSpaceFreeCurrentL2"`
	DiskSpaceAllocatedL2Summary              DiskSpaceSummary      `json:"diskSpaceAllocatedL2Summary"`
	DiskSpaceAllocatedPercentage             []TimestampedPercent  `json:"diskSpaceAllocatedPercentage"`
	DiskSpaceFreeL2                          []TimestampedSpace    `json:"diskSpaceFreeL2"`
	DiskSpaceFreeL1                          []TimestampedSpace    `json:"diskSpaceFreeL1"`
	DiskSpaceAllocatedSummary                DiskSpaceSummary      `json:"diskSpaceAllocatedSummary"`
	DiskSpaceAllocatedGeoCacheCurrent        []TimestampedCapacity `json:"diskSpaceAllocatedGeoCacheCurrent"`
	ChunksEcCodedRatioCurrent                []TimestampedPercent  `json:"chunksEcCodedRatioCurrent"`
	DiskSpaceAllocatedLocalProtectionCurrent []TimestampedCapacity `json:"diskSpaceAllocatedLocalProtectionCurrent"`
	ChunksXorNumber                          int64                 `json:"chunksXorNumber,string"`
	NumMaintenanceNodes                      int64                 `json:"numMaintenanceNodes,string"`
	ChunksEcApplicableTotalSealSize          []TimestampedSpace    `json:"chunksEcApplicableTotalSealSize"`
	GcCombinedTotalDetectedCurrent           []TimestampedCapacity `json:"gcCombinedTotalDetectedCurrent"`
	DiskSpaceTotal                           []TimestampedSpace    `json:"diskSpaceTotal"`
	ChunksL0BtreeNumber                      int64                 `json:"chunksL0BtreeNumber,string"`
	ChunksGeoCacheTotalSize                  int64                 `json:"chunksGeoCacheTotalSize,string"`
	DiskSpaceOfflineTotalCurrent             []TimestampedSpace    `json:"diskSpaceOfflineTotalCurrent"`
	GcSystemTotalDetectedCurrent             []TimestampedCapacity `json:"gcSystemTotalDetectedCurrent"`
	DiskSpaceFreeCurrent                     []TimestampedSpace    `json:"diskSpaceFreeCurrent"`
	ChunksRepoNumber                         int64                 `json:"chunksRepoNumber,string"`
	ChunksL0JournalTotalSize                 int64                 `json:"chunksL0JournalTotalSize,string"`
	RecoveryRateSummary                      Summary               `json:"recoveryRateSummary"`
	GcSystemPendingCurrent                   []TimestampedCapacity `json:"gcSystemPendingCurrent"`
	RecoveryRate                             []TimestampedRate     `json:"recoveryRate"`
	ChunksL0BtreeAvgSize                     float64               `json:"chunksL0BtreeAvgSize,string"`
	DiskSpaceAllocatedSystemMetadataCurrent  []TimestampedCapacity `json:"diskSpaceAllocatedSystemMetadataCurrent"`
	NumDisks                                 int64                 `json:"numDisks,string"`
	GcUserReclaimedPerInterval               []TimestampedCapacity `json:"gcUserReclaimedPerInterval"`
	DiskSpaceAllocatedL1Summary              DiskSpaceSummary      `json:"diskSpaceAllocatedL1Summary"`
	ChunksEcCodedRatio                       []TimestampedPercent  `json:"chunksEcCodedRatio"`
	DiskSpaceFree                            []TimestampedSpace    `json:"diskSpaceFree"`
	AllocatedCapacityForecast                []TimestampedCapacity `json:"allocatedCapacityForecast"`
	DiskSpaceFreeL1Summary                   DiskSpaceSummary      `json:"diskSpaceFreeL1Summary"`
	DiskSpaceTotalSummary                    DiskSpaceSummary      `json:"diskSpaceTotalSummary"`
	DiskSpaceFreeL2Summary                   DiskSpaceSummary      `json:"diskSpaceFreeL2Summary"`
	GcUserReclaimedCurrent                   []TimestampedCapacity `json:"gcUserReclaimedCurrent"`
	NumReadyToReplaceDisks                   int64                 `json:"numReadyToReplaceDisks,string"`
	ChunksEcCodedTotalSealSize               []TimestampedSpace    `json:"chunksEcCodedTotalSealSize"`
	ChunksEcRate                             []TimestampedRate     `json:"chunksEcRate"`
	ChunksL1JournalAvgSize                   float64               `json:"chunksL1JournalAvgSize,string"`
	ChunksL0JournalNumber                    int64                 `json:"chunksL0JournalNumber,string"`
	NumNodesWithSufficientDiskSpace          int64                 `json:"numNodesWithSufficientDiskSpace,string"`
	GcSystemUnreclaimableCurrent             []TimestampedCapacity `json:"gcSystemUnreclaimableCurrent"`
	ChunksEcCompleteTimeEstimate             float64               `json:"chunksEcCompleteTimeEstimate,string"`
	DiskSpaceAllocatedPercentageSummary      PercentageSummary     `json:"diskSpaceAllocatedPercentageSummary"`
	RecoveryCompleteTimeEstimate             float64               `json:"recoveryCompleteTimeEstimate,string"`
	DiskSpaceTotalCurrent                    []TimestampedSpace    `json:"diskSpaceTotalCurrent"`
	NumNodes                                 int64                 `json:"numNodes,string"`
	Name                                     string                `json:"name"`
	GcUserPendingCurrent                     []TimestampedCapacity `json:"gcUserPendingCurrent"`
	NumGoodNodes                             int64                 `json:"numGoodNodes,string"`
	DiskSpaceAllocatedCurrentL1              []TimestampedSpace    `json:"diskSpaceAllocatedCurrentL1"`
	DiskSpaceAllocatedCurrentL2              []TimestampedSpace    `json:"diskSpaceAllocatedCurrentL2"`
	ChunksGeoCacheCount                      int64                 `json:"chunksGeoCacheCount,string"`
	DiskSpaceAllocatedUserDataCurrent        []TimestampedCapacity `json:"diskSpaceAllocatedUserDataCurrent"`
	NumBadDisks                              int64                 `json:"numBadDisks,string"`
	GcCombinedReclaimedOverTimeRange         []TimestampedCapacity `json:"gcCombinedReclaimedOverTimeRange"`
	GcCombinedPendingCurrent                 []TimestampedCapacity `json:"gcCombinedPendingCurrent"`
	DiskSpaceAllocatedGeoCopyCurrent         []TimestampedCapacity `json:"diskSpaceAllocatedGeoCopyCurrent"`
	ChunksEcCodedRatioSummary                Summary               `json:"chunksEcCodedRatioSummary"`
	ChunksL0BtreeTotalSize                   int64                 `json:"chunksL0BtreeTotalSize,string"`
	DiskSpaceReservedCurrent                 []TimestampedSpace    `json:"diskSpaceReservedCurrent"`
	ChunksEcRateCurrent                      []TimestampedRate     `json:"chunksEcRateCurrent"`
	GcCombinedUnreclaimableCurrent           []TimestampedCapacity `json:"gcCombinedUnreclaimableCurrent"`
	RecoveryBadChunksTotalSize               []TimestampedSpace    `json:"recoveryBadChunksTotalSize"`
	ChunksL1JournalNumber                    int64                 `json:"chunksL1JournalNumber,string"`
	GcSystemMetadataIsEnabled                bool                  `json:"gcSystemMetadataIsEnabled,string"`
	DiskSpaceAllocatedPercentageCurrent      []TimestampedPercent  `json:"diskSpaceAllocatedPercentageCurrent"`
	ChunksEcApplicableTotalSealSizeSummary   Summary               `json:"chunksEcApplicableTotalSealSizeSummary"`
	ChunksXorTotalSize                       int64                 `json:"chunksXorTotalSize,string"`
	NumGoodDisks                             int64                 `json:"numGoodDisks,string"`
	RecoveryBadChunksTotalSizeSummary        Summary               `json:"recoveryBadChunksTotalSizeSummary"`
	ChunksGeoCopyTotalSize                   int64                 `json:"chunksGeoCopyTotalSize,string"`
}

type TrafficSummary struct {
	Min []interface{} `json:"Min"`
	Max []interface{} `json:"Max"`
	Avg float64       `json:"Avg,string"`
}

type ReplicationGroup struct {
	Links                                    Links          `json:"_links"`
	ReplicationEgressTrafficSummary          TrafficSummary `json:"replicationEgressTrafficSummary"`
	NumZones                                 int            `json:"numZones,string"`
	ReplicationEgressTraffic                 []interface{}  `json:"replicationEgressTraffic"`
	ReplicationIngressTraffic                []interface{}  `json:"replicationIngressTraffic"`
	ReplicationIngressTrafficCurrent         []interface{}  `json:"replicationIngressTrafficCurrent"`
	Name                                     string         `json:"name"`
	ChunksPendingXorTotalSize                int            `json:"chunksPendingXorTotalSize,string"`
	ChunksRepoPendingReplicationTotalSize    int            `json:"chunksRepoPendingReplicationTotalSize,string"`
	ID                                       string         `json:"id"`
	ReplicationEgressTrafficCurrent          []interface{}  `json:"replicationEgressTrafficCurrent"`
	ChunksJournalPendingReplicationTotalSize int            `json:"chunksJournalPendingReplicationTotalSize,string"`
	ReplicationIngressTrafficSummary         TrafficSummary `json:"replicationIngressTrafficSummary"`
}

// PercentageSummary represents a summary with percentage values.
type PercentageSummary struct {
	Min []TimestampedPercent `json:"Min"`
	Max []TimestampedPercent `json:"Max"`
	Avg float64              `json:"Avg,string"`
}

// NodeData represents the root structure of the JSON.
type NodeData struct {
	Links                               Links                `json:"_links"`
	APIChange                           int                  `json:"apiChange,string"`
	NumBadDisks                         int                  `json:"numBadDisks,string"`
	StoragePoolName                     string               `json:"storagePoolName"`
	DiskSpaceFreeCurrentL1              []TimestampedSpace   `json:"diskSpaceFreeCurrentL1"`
	AllocatedCapacityForecast           []TimestampedSpace   `json:"allocatedCapacityForecast"`
	DiskSpaceFreeCurrentL2              []TimestampedSpace   `json:"diskSpaceFreeCurrentL2"`
	DiskSpaceFreeL1Summary              Summary              `json:"diskSpaceFreeL1Summary"`
	DiskSpaceAllocatedL2Summary         Summary              `json:"diskSpaceAllocatedL2Summary"`
	DiskSpaceTotalSummary               Summary              `json:"diskSpaceTotalSummary"`
	DiskSpaceFreeL2Summary              Summary              `json:"diskSpaceFreeL2Summary"`
	DiskSpaceAllocatedPercentage        []TimestampedPercent `json:"diskSpaceAllocatedPercentage"`
	DisplayName                         string               `json:"displayName"`
	DiskSpaceReservedCurrent            []TimestampedSpace   `json:"diskSpaceReservedCurrent"`
	DiskSpaceFreeL2                     []TimestampedSpace   `json:"diskSpaceFreeL2"`
	NumReadyToReplaceDisks              int                  `json:"numReadyToReplaceDisks,string"`
	DiskSpaceFreeL1                     []TimestampedSpace   `json:"diskSpaceFreeL1"`
	DiskSpaceAllocatedSummary           Summary              `json:"diskSpaceAllocatedSummary"`
	NumMaintenanceDisks                 int                  `json:"numMaintenanceDisks,string"`
	DiskSpaceTotal                      []TimestampedSpace   `json:"diskSpaceTotal"`
	ID                                  string               `json:"id"`
	StoragePoolID                       string               `json:"storagePoolId"`
	DiskSpaceTotalCurrentL1             []TimestampedSpace   `json:"diskSpaceTotalCurrentL1"`
	DiskSpaceTotalCurrentL2             []TimestampedSpace   `json:"diskSpaceTotalCurrentL2"`
	DiskSpaceOfflineTotalCurrent        []TimestampedSpace   `json:"diskSpaceOfflineTotalCurrent"`
	DiskSpaceAllocatedL2                []TimestampedSpace   `json:"diskSpaceAllocatedL2"`
	DiskSpaceAllocatedL1                []TimestampedSpace   `json:"diskSpaceAllocatedL1"`
	DiskSpaceAllocatedPercentageSummary PercentageSummary    `json:"diskSpaceAllocatedPercentageSummary"`
	DiskSpaceFreeCurrent                []TimestampedSpace   `json:"diskSpaceFreeCurrent"`
	DiskSpaceAllocatedPercentageCurrent []TimestampedPercent `json:"diskSpaceAllocatedPercentageCurrent"`
	DiskSpaceTotalCurrent               []TimestampedSpace   `json:"diskSpaceTotalCurrent"`
	DiskSpaceAllocatedCurrent           []TimestampedSpace   `json:"diskSpaceAllocatedCurrent"`
	HealthStatus                        string               `json:"healthStatus"`
	NumGoodDisks                        int                  `json:"numGoodDisks,string"`
	DiskSpaceFreeSummary                Summary              `json:"diskSpaceFreeSummary"`
	NumDisks                            int                  `json:"numDisks,string"`
	DiskSpaceAllocatedCurrentL1         []TimestampedSpace   `json:"diskSpaceAllocatedCurrentL1"`
	DiskSpaceAllocatedCurrentL2         []TimestampedSpace   `json:"diskSpaceAllocatedCurrentL2"`
	DiskSpaceAllocated                  []TimestampedSpace   `json:"diskSpaceAllocated"`
	DiskSpaceAllocatedL1Summary         Summary              `json:"diskSpaceAllocatedL1Summary"`
	DiskSpaceFree                       []TimestampedSpace   `json:"diskSpaceFree"`
}

type DiskData struct {
	Links                               Links                `json:"_links"`
	StoragePoolName                     string               `json:"storagePoolName"`
	DiskSpaceFreeCurrentL1              []TimestampedSpace   `json:"diskSpaceFreeCurrentL1"`
	DiskSpaceFreeCurrentL2              []TimestampedSpace   `json:"diskSpaceFreeCurrentL2"`
	DiskSpaceFreeL1Summary              Summary              `json:"diskSpaceFreeL1Summary"`
	DiskSpaceAllocatedL2Summary         Summary              `json:"diskSpaceAllocatedL2Summary"`
	DiskSpaceTotalSummary               Summary              `json:"diskSpaceTotalSummary"`
	DiskSpaceFreeL2Summary              Summary              `json:"diskSpaceFreeL2Summary"`
	DiskSpaceAllocatedPercentage        []TimestampedPercent `json:"diskSpaceAllocatedPercentage"`
	DisplayName                         string               `json:"displayName"`
	NodeDisplayName                     string               `json:"nodeDisplayName"`
	DiskSpaceFreeL2                     []TimestampedSpace   `json:"diskSpaceFreeL2"`
	DiskSpaceFreeL1                     []TimestampedSpace   `json:"diskSpaceFreeL1"`
	DiskSpaceAllocatedSummary           Summary              `json:"diskSpaceAllocatedSummary"`
	DiskSpaceTotal                      []TimestampedSpace   `json:"diskSpaceTotal"`
	SlotId                              string               `json:"slotId"`
	Id                                  string               `json:"id"`
	StoragePoolId                       string               `json:"storagePoolId"`
	DiskSpaceTotalCurrentL1             []TimestampedSpace   `json:"diskSpaceTotalCurrentL1"`
	DiskSpaceTotalCurrentL2             []TimestampedSpace   `json:"diskSpaceTotalCurrentL2"`
	DiskSpaceAllocatedL2                []TimestampedSpace   `json:"diskSpaceAllocatedL2"`
	SsmL2Status                         string               `json:"ssmL2Status"`
	DiskSpaceAllocatedL1                []TimestampedSpace   `json:"diskSpaceAllocatedL1"`
	DiskSpaceAllocatedPercentageSummary Summary              `json:"diskSpaceAllocatedPercentageSummary"`
	DiskSpaceFreeCurrent                []TimestampedSpace   `json:"diskSpaceFreeCurrent"`
	DiskSpaceAllocatedPercentageCurrent []TimestampedPercent `json:"diskSpaceAllocatedPercentageCurrent"`
	DiskSpaceTotalCurrent               []TimestampedSpace   `json:"diskSpaceTotalCurrent"`
	DiskSpaceAllocatedCurrent           []TimestampedSpace   `json:"diskSpaceAllocatedCurrent"`
	HealthStatus                        string               `json:"healthStatus"`
	DiskSpaceFreeSummary                Summary              `json:"diskSpaceFreeSummary"`
	DiskSpaceAllocatedCurrentL1         []TimestampedSpace   `json:"diskSpaceAllocatedCurrentL1"`
	DiskSpaceAllocated                  []TimestampedSpace   `json:"diskSpaceAllocated"`
	DiskSpaceAllocatedCurrentL2         []TimestampedSpace   `json:"diskSpaceAllocatedCurrentL2"`
	DiskSpaceAllocatedL1Summary         Summary              `json:"diskSpaceAllocatedL1Summary"`
	NodeId                              string               `json:"nodeId"`
	SsmL1Status                         string               `json:"ssmL1Status"`
	DiskSpaceFree                       []TimestampedSpace   `json:"diskSpaceFree"`
}
