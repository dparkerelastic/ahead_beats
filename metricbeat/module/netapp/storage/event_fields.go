package storage

import (
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func createECSFields(ms *MetricSet) mapstr.M {

	return mapstr.M{
		"observer": mapstr.M{
			"hostname": ms.config.HostInfo.Hostname,
			"ip":       ms.config.HostInfo.IP,
			"type":     "storage",
			"vendor":   "NetApp",
		},
	}
}

func createStorageStatisticsFields(stats StorageStatistics) mapstr.M {
	return mapstr.M{
		"timestamp":      stats.Timestamp,
		"status":         stats.Status,
		"throughput_raw": createIOLatencyFields(stats.ThroughputRaw),
		"iops_raw":       createIOLatencyFields(stats.IOPSRaw),
		"latency_raw":    createIOLatencyFields(stats.LatencyRaw),
	}
}

func createStorageMetricsFields(metric StorageMetrics) mapstr.M {
	return mapstr.M{
		"timestamp":  metric.Timestamp,
		"duration":   metric.Duration,
		"status":     metric.Status,
		"latency":    createIOLatencyFields(metric.Latency),
		"iops":       createIOLatencyFields(metric.IOPS),
		"throughput": createIOLatencyFields(metric.Throughput),
	}
}

func createSnapMirrorRelationshipFields(record SnapMirrorRelationship) mapstr.M {
	smdr, err := netapp.ToJSONString(record.SvmdrVolumes)
	if err != nil {
		smdr = ""
	}

	unhealthy_reason, err := netapp.ToJSONString(record.UnhealthyReason)
	if err != nil {
		unhealthy_reason = ""
	}

	return mapstr.M{
		"backoff_level":              record.BackoffLevel,
		"consistency_group_failover": createGroupFailoverFields(record.ConsistencyGroupFailover),
		"destination":                createSnapMirrorEndpointFields(record.Destination),
		"exported_snapshot":          record.ExportedSnapshot,
		"group_type":                 record.GroupType,
		"healthy":                    record.Healthy,
		"identity_preservation":      record.IdentityPreservation,
		"io_serving_copy":            record.IOServingCopy,
		"lag_time":                   record.LagTime,
		"last_transfer_network_compression_ratio": record.LastTransferNetworkRatio,
		"last_transfer_type":                      record.LastTransferType,
		"master_bias_activated_site":              record.MasterBiasActivatedSite,
		"policy":                                  record.Policy,
		"preferred_site":                          record.PreferredSite,
		"restore":                                 record.Restore,
		"source":                                  createSnapMirrorEndpointFields(record.Source),
		"state":                                   record.State,
		"svmdr_volumes":                           smdr,
		"throttle":                                record.Throttle,
		"total_transfer_bytes":                    record.TotalTransferBytes,
		"total_transfer_duration":                 record.TotalTransferDuration,
		"transfer":                                record.Transfer,
		"transfer_schedule":                       record.TransferSchedule,
		"unhealthy_reason":                        unhealthy_reason,
		"uuid":                                    record.UUID,
	}
}

func createGroupFailoverFields(failover ConsistencyGroupFailover) mapstr.M {
	return mapstr.M{
		"error":  createStatusFields(failover.Error),
		"state":  failover.State,
		"status": createStatusFields(failover.Status),
		"type":   failover.Type,
	}
}

func createStatusFields(status StorageStatus) mapstr.M {
	return mapstr.M{
		"code":    status.Code,
		"message": status.Message,
	}
}

func createSnapMirrorEndpointFields(snapEndpoint SnapMirrorEndpoint) mapstr.M {

	volumes, err := netapp.ToJSONString(snapEndpoint.ConsistencyGroupVolumes)
	if err != nil {
		volumes = ""
	}

	return mapstr.M{
		"cluster":                   createNamedObjectFields(snapEndpoint.Cluster),
		"svm":                       createNamedObjectFields(snapEndpoint.SVM),
		"luns":                      createNamedObjectFields(snapEndpoint.LUNs),
		"path":                      snapEndpoint.Path,
		"consistency_group_volumes": volumes,
	}
}

func createAggregateFields(record Aggregate) mapstr.M {

	fields := mapstr.M{
		"uuid":                    record.UUID,
		"name":                    record.Name,
		"node":                    createNamedObjectFields(record.Node),
		"home_node":               createNamedObjectFields(record.HomeNode),
		"snapshot":                createAggregateSnapshotFields(record.Snapshot),
		"space":                   createAggregateSpaceFields(record.Space),
		"state":                   record.State,
		"snaplock_type":           record.SnaplockType,
		"create_time":             record.CreateTime,
		"data_encryption":         createAggregateEncryptionFields(record.DataEncryption),
		"block_storage":           createAggregateBlockStorageFields(record.BlockStorage),
		"cloud_storage":           createAggregateCloudStorageFields(record.CloudStorage),
		"inactive_data_reporting": createInactiveDataReportFields(record.InactiveDataReport),
		"inode_attributes":        createInodeAttributesFields(record.InodeAttributes),
		"volume_count":            record.VolumeCount,
		"metrics":                 createStorageMetricsFields(record.Metrics),
		"statistics":              createStorageStatisticsFields(record.Statistics),
	}

	return fields
}

func createInodeAttributesFields(inodeAttributes InodeAttributes) mapstr.M {
	return mapstr.M{
		"files_total":         inodeAttributes.FilesTotal,
		"files_used":          inodeAttributes.FilesUsed,
		"max_files_available": inodeAttributes.MaxFilesAvailable,
		"max_files_possible":  inodeAttributes.MaxFilesPossible,
		"max_files_used":      inodeAttributes.MaxFilesUsed,
		"used_percent":        inodeAttributes.UsedPercent,
	}
}

func createInactiveDataReportFields(inactiveDataReport InactiveDataReport) mapstr.M {
	m := mapstr.M{
		"enabled": inactiveDataReport.Enabled,
	}
	if inactiveDataReport.Enabled {
		m["start_time"] = inactiveDataReport.StartTime
	}
	return m
}

func createAggregateCloudStorageFields(cloudStorage AggregateCloudStorage) mapstr.M {
	stores, err := netapp.ToJSONString(cloudStorage.Stores)
	if err != nil {
		stores = ""
	}
	return mapstr.M{
		"attach_eligible": cloudStorage.AttachEligible,
		"stores":          stores,
	}
}

func createAggregateEncryptionFields(encryption AggregateEncryption) mapstr.M {
	return mapstr.M{
		"software_encryption_enabled": encryption.SoftwareEncryptionEnabled,
		"drive_protection_enabled":    encryption.DriveProtectionEnabled,
	}
}

// createAggregateBlockStorageFields maps AggregateBlockStorage to mapstr.M
func createAggregateBlockStorageFields(blockStorage AggregateBlockStorage) mapstr.M {
	return mapstr.M{
		"uses_partitions": blockStorage.UsesPartitions,
		"storage_type":    blockStorage.StorageType,
		"primary":         blockStorage.Primary,
		"hybrid_cache":    blockStorage.HybridCache,
		"mirror":          blockStorage.Mirror,
		"plexes":          blockStorage.Plexes,
	}
}

func createNamedObjectFields(namedObject NamedObject) mapstr.M {
	return mapstr.M{
		"name": namedObject.Name,
		"uuid": namedObject.UUID,
	}
}

func createAggregateSnapshotFields(s AggregateSnapshot) mapstr.M {
	return mapstr.M{
		"files_total":         s.FilesTotal,
		"files_used":          s.FilesUsed,
		"max_files_available": s.MaxFilesAvailable,
		"max_files_used":      s.MaxFilesUsed,
	}
}

func createAggregateSpaceFields(s AggregateSpace) mapstr.M {
	return mapstr.M{
		"block_storage":                createBlockStorageSpaceFields(s.BlockStorage),
		"snapshot":                     createSnapshotSpaceFields(s.Snapshot),
		"cloud_storage":                createCloudStorageSpaceFields(s.CloudStorage),
		"efficiency":                   createEfficiencyFields(s.Efficiency),
		"efficiency_without_snapshots": createEfficiencySimpleFields(s.EfficiencyWithoutSnapshots),
		"efficiency_without_snapshots_flexclones": createEfficiencySimpleFields(s.EfficiencyWithoutSnapshotsFlexclones),
	}
}
func createEfficiencyFields(efficiency Efficiency) mapstr.M {
	return mapstr.M{
		"savings":                           efficiency.Savings,
		"ratio":                             efficiency.Ratio,
		"logical_used":                      efficiency.LogicalUsed,
		"cross_volume_background_dedupe":    efficiency.CrossVolumeBackgroundDedupe,
		"cross_volume_inline_dedupe":        efficiency.CrossVolumeInlineDedupe,
		"cross_volume_dedupe_savings":       efficiency.CrossVolumeDedupeSavings,
		"auto_adaptive_compression_savings": efficiency.AutoAdaptiveCompressionSavings,
		"enable_workload_informed_tsse":     efficiency.EnableWorkloadInformedTSSE,
		"wise_tsse_min_used_capacity_pct":   efficiency.WiseTSSEMinUsedCapacityPct,
	}
}

func createEfficiencySimpleFields(efficiencySimple EfficiencySimple) mapstr.M {
	return mapstr.M{
		"savings":      efficiencySimple.Savings,
		"ratio":        efficiencySimple.Ratio,
		"logical_used": efficiencySimple.LogicalUsed,
	}
}

func createCloudStorageSpaceFields(cloudStorageSpace CloudStorageSpace) mapstr.M {
	return mapstr.M{
		"used": cloudStorageSpace.Used,
	}
}

func createSnapshotSpaceFields(snapshotSpace SnapshotSpace) mapstr.M {
	return mapstr.M{
		"used_percent":    snapshotSpace.UsedPercent,
		"available":       snapshotSpace.Available,
		"total":           snapshotSpace.Total,
		"used":            snapshotSpace.Used,
		"reserve_percent": snapshotSpace.ReservePercent,
	}
}

func createBlockStorageSpaceFields(b BlockStorageSpace) mapstr.M {
	return mapstr.M{
		"size":                                     b.Size,
		"available":                                b.Available,
		"used":                                     b.Used,
		"used_percent":                             b.UsedPercent,
		"full_threshold_percent":                   b.FullThresholdPercent,
		"physical_used":                            b.PhysicalUsed,
		"physical_used_percent":                    b.PhysicalUsedPercent,
		"data_compacted_count":                     b.DataCompactedCount,
		"data_compaction_space_saved":              b.DataCompactionSpaceSaved,
		"data_compaction_space_saved_percent":      b.DataCompactionSpaceSavedPercent,
		"volume_deduplication_shared_count":        b.VolumeDeduplicationSharedCount,
		"volume_deduplication_space_saved":         b.VolumeDeduplicationSpaceSaved,
		"volume_deduplication_space_saved_percent": b.VolumeDeduplicationSpaceSavedPercent,
	}
}

func createDiskFields(record Disk) mapstr.M {
	aggregates := make([]mapstr.M, len(record.Aggregates))
	for i, agg := range record.Aggregates {
		aggregates[i] = createNamedObjectFields(agg)
	}

	return mapstr.M{
		"name":                    record.Name,
		"uid":                     record.UID,
		"serial_number":           record.SerialNumber,
		"model":                   record.Model,
		"vendor":                  record.Vendor,
		"firmware_version":        record.FirmwareVersion,
		"usable_size":             record.UsableSize,
		"rated_life_used_percent": record.RatedLifeUsedPercent,
		"type":                    record.Type,
		"effective_type":          record.EffectiveType,
		"class":                   record.Class,
		"container_type":          record.ContainerType,
		"pool":                    record.Pool,
		"state":                   record.State,
		"node":                    createNamedObjectFields(record.Node),
		"home_node":               createNamedObjectFields(record.HomeNode),
		"aggregates":              aggregates,
		"shelf":                   record.Shelf.UID,
		"local":                   record.Local,
		"bay":                     record.Bay,
		"self_encrypting":         record.SelfEncrypting,
		"fips_certified":          record.FipsCertified,
		"bytes_per_sector":        record.BytesPerSector,
		"sector_count":            record.SectorCount,
		"right_size_sector_count": record.RightSizeSectorCount,
		"physical_size":           record.PhysicalSize,
		"stats":                   createDiskStatsFields(record.Stats),
	}
}

func createDiskPathFields(path DiskPath) mapstr.M {
	return mapstr.M{
		"disk_path_name": path.DiskPathName,
		"initiator":      path.Initiator,
		"port_name":      path.PortName,
		"port_type":      path.PortType,
		"wwnn":           path.WWNN,
		"wwpn":           path.WWPN,
		"node":           createNamedObjectFields(path.Node),
	}
}

func createDiskStatsFields(stats DiskStats) mapstr.M {
	return mapstr.M{
		"average_latency":  stats.AverageLatency,
		"throughput":       stats.Throughput,
		"iops_total":       stats.IOPSTotal,
		"path_error_count": stats.PathErrorCount,
		"power_on_hours":   stats.PowerOnHours,
	}
}

func createLUNFields(record LUN) mapstr.M {
	return mapstr.M{
		"uuid":          record.UUID,
		"svm":           createNamedObjectFields(record.SVM),
		"name":          record.Name,
		"location":      createLunLocationFields(record.Location),
		"class":         record.Class,
		"create_time":   record.CreateTime,
		"enabled":       record.Enabled,
		"os_type":       record.OsType,
		"serial_number": record.SerialNumber,
		"space":         createLunSpaceFields(record.Space),
		"status":        createLunStatusFields(record.Status),
		"vvol":          createLunVVolFields(record.VVol),
	}
}

func createLunLocationFields(location LunLocation) mapstr.M {
	return mapstr.M{
		"logical_unit": location.LogicalUnit,
		"node":         createNamedObjectFields(location.Node),
		"volume":       createNamedObjectFields(location.Volume),
	}
}

func createLunSpaceFields(space LunSpace) mapstr.M {
	return mapstr.M{
		"scsi_thin_provisioning_support_enabled": space.SCSIThinProvisioningSupportEnabled,
		"size":                                   space.Size,
		"used":                                   space.Used,
		"guarantee":                              createLunGuaranteeFields(space.Guarantee),
	}
}

func createLunGuaranteeFields(guarantee LunGuarantee) mapstr.M {
	return mapstr.M{
		"requested": guarantee.Requested,
		"reserved":  guarantee.Reserved,
	}
}

func createLunStatusFields(status LunStatus) mapstr.M {
	return mapstr.M{
		"container_state": status.ContainerState,
		"mapped":          status.Mapped,
		"read_only":       status.ReadOnly,
		"state":           status.State,
	}
}

func createLunVVolFields(vvol LunVVol) mapstr.M {
	return mapstr.M{
		"is_bound": vvol.IsBound,
	}
}

func createQTreeFields(q Qtree) mapstr.M {
	return mapstr.M{
		"volume":           createNamedObjectFields(q.Volume),
		"id":               q.ID,
		"svm":              createNamedObjectFields(q.SVM),
		"name":             q.Name,
		"security_style":   q.SecurityStyle,
		"unix_permissions": q.UnixPermissions,
		"export_policy":    createExportPolicyIDFields(q.ExportPolicy),
		"path":             q.Path,
		"nas_path":         q.NAS.Path,
		"user_id":          q.User.ID,
		"group_id":         q.Group.ID,
		"metric":           createQtreeMetricsFields(q.Metrics),
	}
}

func createExportPolicyIDFields(e ExportPolicyID) mapstr.M {
	return mapstr.M{
		"name": e.Name,
		"id":   e.ID,
	}
}

func createQtreeMetricsFields(q QtreeMetrics) mapstr.M {
	return mapstr.M{
		"duration":   q.Duration,
		"iops":       createIOLatencyFields(q.IOPS),
		"latency":    createIOLatencyFields(q.Latency),
		"throughput": createIOLatencyFields(q.Throughput),
		"qtree":      createQtreeBriefFields(q.Qtree),
		"status":     q.Status,
		"svm":        createNamedObjectFields(q.SVM),
		"timestamp":  q.MetricTimestamp,
		"volume":     createNamedObjectFields(q.Volume),
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

func createQtreeBriefFields(q QtreeRef) mapstr.M {
	return mapstr.M{
		"id":   q.ID,
		"name": q.Name,
	}
}

func createQuotaReportFields(qr QuotaReport) mapstr.M {
	users, err := netapp.ToJSONString(qr.Users)
	if err != nil {
		users = ""
	}

	return mapstr.M{
		"svm":    createNamedObjectFields(qr.SVM),
		"volume": createNamedObjectFields(qr.Volume),
		"qtree":  createQtreeBriefFields(qr.Qtree),
		"type":   qr.Type,
		"index":  qr.Index,
		"group":  createNamedObjectFields(qr.Group),
		"users":  users,
		"files": mapstr.M{
			"hard_limit": qr.Files.HardLimit,
			"soft_limit": qr.Files.SoftLimit,
			"used":       qr.Files.Used,
		},
		"space": mapstr.M{
			"hard_limit": qr.Space.HardLimit,
			"soft_limit": qr.Space.SoftLimit,
			"used":       qr.Space.Used,
		},
	}
}

func createQuotaRuleFields(qr QuotaRule) mapstr.M {
	users, err := netapp.ToJSONString(qr.Users)
	if err != nil {
		users = ""
	}
	return mapstr.M{
		"files": mapstr.M{
			"hard_limit": qr.Files.HardLimit,
			"soft_limit": qr.Files.SoftLimit,
		},
		"space": mapstr.M{
			"hard_limit": qr.Space.HardLimit,
			"soft_limit": qr.Space.SoftLimit,
		},
		"qtree":        createQtreeBriefFields(qr.Qtree),
		"svm_name":     qr.SVM.Name,
		"type":         qr.Type,
		"user_mapping": qr.UserMapping,
		"users":        users,
		"uuid":         qr.UUID,
		"volume":       qr.Volume.Name,
	}
}

func createACPEvents(m *MetricSet, record Shelf) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, acp := range record.ACPs {
		shelfFields := createShelfFields(record)
		shelfFields["acp"] = createACPFields(acp)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"shelf": shelfFields,
			},
			RootFields: createECSFields(m),
		}
		events = append(events, event)
	}
	return events
}

func createACPFields(acp ACP) interface{} {
	return mapstr.M{
		"enabled":          acp.Enabled,
		"channel":          acp.Channel,
		"connection_state": acp.ConnectionState,
		"node":             createNamedObjectFields(acp.Node),
	}
}

func createPortEvents(m *MetricSet, record Shelf) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, port := range record.Ports {
		shelfFields := createShelfFields(record)
		shelfFields["port"] = createPortFields(port)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"shelf": shelfFields,
			},
			RootFields: createECSFields(m),
		}
		events = append(events, event)
	}
	return events
}

func createPortFields(port ShelfPort) interface{} {
	m := mapstr.M{
		"id":         port.ID,
		"module_id":  port.ModuleID,
		"designator": port.Designator,
		"state":      port.State,
		"internal":   port.Internal,
		"wwn":        port.WWN,
	}
	if port.Cable != nil {
		m["cable_id"] = port.Cable.Identifier
	}
	if port.Remote != nil {
		m["remote_wwn"] = port.Remote.WWN
	}
	return m
}

func createCurrentEvents(m *MetricSet, record Shelf) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, sensor := range record.CurrentSensors {
		shelfFields := createShelfFields(record)
		shelfFields["current_sensor"] = createCurrentFields(sensor)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"shelf": shelfFields,
			},
			RootFields: createECSFields(m),
		}
		events = append(events, event)
	}
	return events
}

func createCurrentFields(sensor CurrentSensor) mapstr.M {
	return mapstr.M{
		"id":        sensor.ID,
		"location":  sensor.Location,
		"current":   sensor.Current,
		"state":     sensor.State,
		"installed": sensor.Installed,
	}
}

func createVoltageEvents(m *MetricSet, record Shelf) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, sensor := range record.VoltageSensors {
		shelfFields := createShelfFields(record)
		shelfFields["voltage_sensor"] = createVoltageFields(sensor)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"shelf": shelfFields,
			},
			RootFields: createECSFields(m),
		}
		events = append(events, event)
	}
	return events
}

func createVoltageFields(sensor VoltageSensor) mapstr.M {
	return mapstr.M{
		"id":        sensor.ID,
		"location":  sensor.Location,
		"voltage":   sensor.Voltage,
		"state":     sensor.State,
		"installed": sensor.Installed,
	}
}

func createThermalEvents(m *MetricSet, record Shelf) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, sensor := range record.TempSensors {
		shelfFields := createShelfFields(record)
		shelfFields["temperature_sensor"] = createTempSensorFields(sensor)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"shelf": shelfFields,
			},
			RootFields: createECSFields(m),
		}
		events = append(events, event)
	}
	return events
}

func createTempSensorFields(sensor TemperatureSensor) mapstr.M {
	return mapstr.M{
		"id":          sensor.ID,
		"location":    sensor.Location,
		"temperature": sensor.Temperature,
		"ambient":     sensor.Ambient,
		"state":       sensor.State,
		"installed":   sensor.Installed,
		"threshold": mapstr.M{
			"high": mapstr.M{
				"critical": sensor.Threshold.High.Critical,
				"warning":  sensor.Threshold.High.Warning,
			},
			"low": mapstr.M{
				"critical": sensor.Threshold.Low.Critical,
				"warning":  sensor.Threshold.Low.Warning,
			},
		},
	}
}

func createFanEvents(m *MetricSet, record Shelf) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, fan := range record.Fans {
		shelfFields := createShelfFields(record)
		shelfFields["fan"] = createFansFields(fan)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"shelf": shelfFields,
			},
			RootFields: createECSFields(m),
		}
		events = append(events, event)
	}
	return events
}

func createFansFields(fan Fan) mapstr.M {
	return mapstr.M{
		"id":        fan.ID,
		"location":  fan.Location,
		"rpm":       fan.RPM,
		"state":     fan.State,
		"installed": fan.Installed,
	}
}

func createPSUEvents(m *MetricSet, record Shelf) []mb.Event {
	timestamp := time.Now().UTC()

	var events []mb.Event
	for _, fru := range record.FRUs {
		if fru.Type == "psu" && fru.PSU != nil {
			shelfFields := createShelfFields(record)
			shelfFields["fru"] = createPSUFields(fru)
			event := mb.Event{
				Timestamp: timestamp,
				MetricSetFields: mapstr.M{
					"shelf": shelfFields,
				},
				RootFields: createECSFields(m),
			}
			events = append(events, event)
		}
	}
	return events
}

func createDiskPathEvents(m *MetricSet, record Disk) []mb.Event {
	timestamp := time.Now().UTC()
	var events []mb.Event
	for _, path := range record.Paths {
		diskFields := createDiskFields(record)
		diskFields["path"] = createDiskPathFields(path)
		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: mapstr.M{
				"disk": diskFields,
			},
			RootFields: createECSFields(m),
		}
		events = append(events, event)
	}
	return events
}

func createPSUFields(fru FRU) interface{} {

	return mapstr.M{
		"type":             fru.Type,
		"id":               fru.ID,
		"state":            fru.State,
		"part_number":      fru.PartNumber,
		"serial_number":    fru.SerialNumber,
		"firmware_version": fru.FirmwareVersion,
		"installed":        fru.Installed,
		"psu": mapstr.M{
			"model":        fru.PSU.Model,
			"power_drawn":  fru.PSU.PowerDrawn,
			"power_rating": fru.PSU.PowerRating,
			"crest_factor": fru.PSU.CrestFactor,
		},
	}
}

func createShelfFields(record Shelf) mapstr.M {

	paths, err := netapp.ToJSONString(record.Paths)
	if err != nil {
		paths = ""
	}

	bays, err := netapp.ToJSONString(record.Bays)
	if err != nil {
		bays = ""
	}

	return mapstr.M{
		"uid":             record.UID,
		"name":            record.Name,
		"id":              record.ID,
		"serial_number":   record.SerialNumber,
		"model":           record.Model,
		"module_type":     record.ModuleType,
		"internal":        record.Internal,
		"local":           record.Local,
		"manufacturer":    record.Manufacturer.Name,
		"state":           record.State,
		"connection_type": record.ConnectionType,
		"disk_count":      record.DiskCount,
		"location_led":    record.LocationLED,
		"paths":           paths,
		"bays":            bays,
	}
}

func createVolumeFields(record Volume) mapstr.M {
	aggregates, err := netapp.ToJSONString(record.Aggregates)
	if err != nil {
		aggregates = ""
	}

	fields := mapstr.M{
		"uuid":                             record.UUID,
		"comment":                          record.Comment,
		"create_time":                      record.CreateTime,
		"language":                         record.Language,
		"name":                             record.Name,
		"size":                             record.Size,
		"state":                            record.State,
		"style":                            record.Style,
		"tiering_policy":                   record.Tiering.Policy,
		"cloud_retrieval_policy":           record.CloudRetrievalPolicy,
		"type":                             record.Type,
		"aggregates":                       aggregates,
		"snapshot_count":                   record.SnapshotCount,
		"msid":                             record.MSID,
		"scheduled_snapshot_naming_scheme": record.ScheduledSnapshotNamingScheme,
		"clone":                            record.Clone,
		"nas":                              record.NAS,
		"snapshot_locking_enabled":         record.SnapshotLockingEnabled,
		"snapshot_policy":                  createNamedObjectFields(record.NamedObject),
		"svm":                              createNamedObjectFields(record.SVM),
		"space":                            createVolumeSpaceFields(record.Space),
		"metrics":                          createStorageMetricsFields(record.Metrics),
		"snapmirror":                       record.Snapmirror,
		"activity_tracking":                record.ActivityTracking,
		"granular_data":                    record.GranularData,
		"granular_data_mode":               record.GranularDataMode,
		// "analytics":                        record.Analytics,
	}

	return fields
}

func createVolumeSpaceFields(volumeSpace VolumeSpace) mapstr.M {

	return mapstr.M{
		"size":                                 volumeSpace.Size,
		"available":                            volumeSpace.Available,
		"used":                                 volumeSpace.Used,
		"is_used_stale":                        volumeSpace.IsUsedStale,
		"block_storage_inactive_user_data":     volumeSpace.BlockStorageInactiveUserData,
		"local_tier_footprint":                 volumeSpace.LocalTierFootprint,
		"footprint":                            volumeSpace.Footprint,
		"over_provisioned":                     volumeSpace.OverProvisioned,
		"metadata":                             volumeSpace.Metadata,
		"total_footprint":                      volumeSpace.TotalFootprint,
		"delayed_free_footprint":               volumeSpace.DelayedFreeFootprint,
		"file_operation_metadata":              volumeSpace.FileOperationMetadata,
		"volume_guarantee_footprint":           volumeSpace.VolumeGuaranteeFootprint,
		"effective_total_footprint":            volumeSpace.EffectiveTotalFootprint,
		"user_data":                            volumeSpace.UserData,
		"used_by_afs":                          volumeSpace.UsedByAFS,
		"available_percent":                    volumeSpace.AvailablePercent,
		"afs_total":                            volumeSpace.AFSTotal,
		"full_threshold_percent":               volumeSpace.FullThresholdPercent,
		"nearly_full_threshold_percent":        volumeSpace.NearlyFullThresholdPercent,
		"overwrite_reserve":                    volumeSpace.OverwriteReserve,
		"overwrite_reserve_used":               volumeSpace.OverwriteReserveUsed,
		"size_available_for_snapshots":         volumeSpace.SizeAvailableForSnapshots,
		"percent_used":                         volumeSpace.PercentUsed,
		"fractional_reserve":                   volumeSpace.FractionalReserve,
		"block_storage_inactive_user_data_pct": volumeSpace.BlockStorageInactiveUserDataPct,
		"physical_used_percent":                volumeSpace.PhysicalUsedPercent,
		"physical_used":                        volumeSpace.PhysicalUsed,
		"expected_available":                   volumeSpace.ExpectedAvailable,
		"filesystem_size":                      volumeSpace.FilesystemSize,
		"filesystem_size_fixed":                volumeSpace.FilesystemSizeFixed,
		"large_size_enabled":                   volumeSpace.LargeSizeEnabled,
		"total_metadata":                       volumeSpace.TotalMetadata,
		"total_metadata_footprint":             volumeSpace.TotalMetadataFootprint,
		"logical_space": mapstr.M{
			"reporting":         volumeSpace.LogicalSpace.Reporting,
			"enforcement":       volumeSpace.LogicalSpace.Enforcement,
			"used_by_afs":       volumeSpace.LogicalSpace.UsedByAFS,
			"used_percent":      volumeSpace.LogicalSpace.UsedPercent,
			"used":              volumeSpace.LogicalSpace.Used,
			"used_by_snapshots": volumeSpace.LogicalSpace.UsedBySnapshots,
		},
	}
}

func createSVMPeerFields(record SVMPeer) mapstr.M {
	applications, err := netapp.ToJSONString(record.Applications)
	if err != nil {
		applications = ""
	}
	return mapstr.M{
		"applications": applications,
		"name":         record.Name,
		"peer": mapstr.M{
			"cluster": createNamedObjectFields(record.Peer.Cluster),
			"svm":     createNamedObjectFields(record.Peer.SVM),
		},
		"state": record.State,
		"svm":   createNamedObjectFields(record.SVM),
		"uuid":  record.UUID,
	}
}

func createSVMFields(record SVM) mapstr.M {
	aggregates, err := netapp.ToJSONString(record.Aggregates)
	if err != nil {
		aggregates = ""
	}
	ipInterfaces, err := netapp.ToJSONString(record.IPInterfaces)
	if err != nil {
		ipInterfaces = ""
	}

	return mapstr.M{
		"uuid":                                 record.UUID,
		"name":                                 record.Name,
		"subtype":                              record.Subtype,
		"language":                             record.Language,
		"aggregates":                           aggregates,
		"state":                                record.State,
		"comment":                              record.Comment,
		"ipspace":                              createNamedObjectFields(record.IPSpace),
		"ip_interfaces":                        ipInterfaces,
		"snapshot_policy":                      createNamedObjectFields(record.SnapshotPolicy),
		"nis_enabled":                          record.NIS.Enabled,
		"ldap_enabled":                         record.LDAP.Enabled,
		"nfs":                                  createProtocolStatusFields(record.NFS),
		"cifs":                                 createProtocolStatusFields(record.CIFS),
		"iscsi":                                createProtocolStatusFields(record.ISCSI),
		"fcp":                                  createProtocolStatusFields(record.FCP),
		"nvme":                                 createProtocolStatusFields(record.NVMe),
		"ndmp_allowed":                         record.NDMP.Allowed,
		"s3":                                   createProtocolStatusFields(record.S3),
		"certificate":                          record.Certificate.UUID,
		"aggregates_delegated":                 record.AggregatesDelegated,
		"retention_period":                     record.RetentionPeriod,
		"max_volumes":                          record.MaxVolumes,
		"anti_ransomware_default_volume_state": record.AntiRansomwareDefaultVolumeState,
		"is_space_reporting_logical":           record.IsSpaceReportingLogical,
		"is_space_enforcement_logical":         record.IsSpaceEnforcementLogical,
		"auto_enable_analytics":                record.AutoEnableAnalytics,
		"auto_enable_activity_tracking":        record.AutoEnableActivityTracking,
		"anti_ransomware_auto_switch_enabled":  record.AntiRansomwareAutoSwitchEnabled,
		"anti_ransomware_auto_switch_data_percent": record.AntiRansomwareAutoSwitchDataPercent,
		"anti_ransomware_auto_switch_no_ext_days":  record.AntiRansomwareAutoSwitchNoExtDays,
		"anti_ransomware_auto_switch_min_period":   record.AntiRansomwareAutoSwitchMinPeriod,
		"anti_ransomware_auto_switch_min_files":    record.AntiRansomwareAutoSwitchMinFiles,
		"anti_ransomware_auto_switch_min_exts":     record.AntiRansomwareAutoSwitchMinExts,
		// "nsswitch":                             record.NSSwitch,

	}
}

func createProtocolStatusFields(status ProtocolStatus) mapstr.M {
	return mapstr.M{
		"allowed": status.Allowed,
		"enabled": status.Enabled,
	}
}

func createQosPolicyFields(policy QosPolicy) mapstr.M {
	fields := mapstr.M{
		"name":         policy.Name,
		"object_count": policy.ObjectCount,
		"pgid":         policy.Pgid,
		"policy_class": policy.PolicyClass,
		"scope":        policy.Scope,
		"svm":          createNamedObjectFields(policy.SVM),
		"uuid":         policy.UUID,
	}
	if policy.Adaptive != nil {
		fields["adaptive"] = policy.Adaptive
	}
	if policy.Fixed != nil {
		fields["fixed"] = policy.Fixed
	}
	return fields
}
