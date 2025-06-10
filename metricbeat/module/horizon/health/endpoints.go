package health

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func getConnectionServers(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("ConnectionServers")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	// Event per connection server per service/cs_replication/session_protocol
	var servers []ConnectionServer
	err = json.Unmarshal([]byte(response), &servers)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, server := range servers {

		for _, service := range server.Services {
			serverFields := createServerFields(server)
			serviceFields := mapstr.M{
				"service": mapstr.M{
					"name":   service.ServiceName,
					"status": service.Status,
				},
			}
			serverFields["service"] = serviceFields

			event := mb.Event{
				Timestamp:       timestamp,
				MetricSetFields: serverFields,
				RootFields:      createECSFields(m),
			}
			events = append(events, event)

		}

		for _, csReplication := range server.CSReplications {
			serverFields := createServerFields(server)
			csReplicationFields := mapstr.M{
				"cs_replication": mapstr.M{
					"server_name": csReplication.ServerName,
					"status":      csReplication.Status,
				},
			}
			serverFields["cs_replication"] = csReplicationFields

			event := mb.Event{
				Timestamp:       timestamp,
				MetricSetFields: serverFields,
				RootFields:      createECSFields(m),
			}
			events = append(events, event)

		}

		for _, sessionProtocol := range server.SessionProtocols {
			serverFields := createServerFields(server)
			sessionProtocolFields := mapstr.M{
				"session_protocol": mapstr.M{
					"name":   sessionProtocol.Protocol,
					"status": sessionProtocol.SessionCount,
				},
			}
			serverFields["session_protocol"] = sessionProtocolFields

			event := mb.Event{
				Timestamp:       timestamp,
				MetricSetFields: serverFields,
				RootFields:      createECSFields(m),
			}
			events = append(events, event)

		}
	}

	return events, nil
}

// createServerFields creates a reusable common.MapStr for a connection server
func createServerFields(server ConnectionServer) mapstr.M {
	return mapstr.M{
		"connection_server": mapstr.M{
			"name":                    server.Name,
			"status":                  server.Status,
			"connection_count":        server.ConnectionCount,
			"tunnel_connection_count": server.TunnelConnectionCount,
			"default_certificate":     server.DefaultCertificate,
			"certificate": mapstr.M{
				"valid":      server.Certificate.Valid,
				"valid_from": server.Certificate.ValidFrom,
				"valid_to":   server.Certificate.ValidTo,
			},
			"details": mapstr.M{
				"version": server.Details.Version,
				"build":   server.Details.Build,
			},
		},
	}
}

func getDesktopPools(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("ConnectionServers")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	// Event per connection server per service/cs_replication/session_protocol
	var pools []DesktopPool
	err = json.Unmarshal([]byte(response), &pools)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, pool := range pools {
		// Get installed applications for the pool
		installedAppEvents, err := getInstalledAppEvents(pool, m)
		if err != nil {
			m.Logger().Warnf("failed to get installed applications for pool %s: %v", pool.ID, err)
		}

		// If we had no installed application events, create a single event for the pool
		if len(installedAppEvents) == 0 {
			event := mb.Event{
				Timestamp:       timestamp,
				MetricSetFields: createDesktopPoolFields(pool),
				RootFields:      createECSFields(m),
			}
			events = append(events, event)
		} else {
			events = append(events, installedAppEvents...)
		}
	}

	return events, nil
}

// createDesktopPoolFields creates a reusable mapstr.M for a desktop pool
func createDesktopPoolFields(pool DesktopPool) mapstr.M {
	return mapstr.M{
		"desktop_pool": mapstr.M{
			"id":           pool.ID,
			"name":         pool.Name,
			"display_name": pool.DisplayName,
			"description":  pool.Description,
			"type":         pool.Type,
			"source":       pool.Source,
			"enabled":      pool.Enabled,
			"settings": mapstr.M{
				"delete_in_progress":               pool.Settings.DeleteInProgress,
				"enable_client_restrictions":       pool.Settings.EnableClientRestrictions,
				"allow_multiple_sessions_per_user": pool.Settings.AllowMultipleSessionsPerUser,
				"session_type":                     pool.Settings.SessionType,
				"cloud_managed":                    pool.Settings.CloudManaged,
				"cloud_assigned":                   pool.Settings.CloudAssigned,
				"session_settings": mapstr.M{
					"power_policy":                           pool.Settings.SessionSettings.PowerPolicy,
					"disconnected_session_timeout_policy":    pool.Settings.SessionSettings.DisconnectedSessionTimeoutPolicy,
					"disconnected_session_timeout_minutes":   pool.Settings.SessionSettings.DisconnectedSessionTimeoutMinutes,
					"allow_users_to_reset_machines":          pool.Settings.SessionSettings.AllowUsersToResetMachines,
					"allow_multiple_sessions_per_user":       pool.Settings.SessionSettings.AllowMultipleSessionsPerUser,
					"delete_or_refresh_machine_after_logoff": pool.Settings.SessionSettings.DeleteOrRefreshMachineAfterLogoff,
					"refresh_os_disk_after_logoff":           pool.Settings.SessionSettings.RefreshOSDiskAfterLogoff,
				},
				"display_protocol_settings": mapstr.M{
					"display_protocols":                 pool.Settings.DisplayProtocolSettings.DisplayProtocols,
					"default_display_protocol":          pool.Settings.DisplayProtocolSettings.DefaultDisplayProtocol,
					"allow_users_to_choose_protocol":    pool.Settings.DisplayProtocolSettings.AllowUsersToChooseProtocol,
					"html_access_enabled":               pool.Settings.DisplayProtocolSettings.HTMLAccessEnabled,
					"session_collaboration_enabled":     pool.Settings.DisplayProtocolSettings.SessionCollaborationEnabled,
					"renderer3d":                        pool.Settings.DisplayProtocolSettings.Renderer3D,
					"grid_vgpus_enabled":                pool.Settings.DisplayProtocolSettings.GridVGPUsEnabled,
					"max_number_of_monitors":            pool.Settings.DisplayProtocolSettings.MaxNumberOfMonitors,
					"max_resolution_of_any_one_monitor": pool.Settings.DisplayProtocolSettings.MaxResolutionOfAnyOneMonitor,
				},
			},
		},
	}
}

func getInstalledAppEvents(pool DesktopPool, m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpointFormat := "rest/inventory/v1/desktop-pools/%s/installed-applications"
	endpoint := fmt.Sprintf(endpointFormat, pool.ID)

	response, err := client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var applications []InstalledApplication
	err = json.Unmarshal([]byte(response), &applications)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	// Create a separate event for each installed application. This works better with Kibana
	// than using a nested structure.
	for _, application := range applications {
		poolFields := createDesktopPoolFields(pool)
		poolFields["installed_application"] = createInstalledApplicationFields(application)
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: poolFields,
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

func createInstalledApplicationFields(application InstalledApplication) mapstr.M {
	return mapstr.M{
		"installed_application": mapstr.M{
			"name":            application.Name,
			"version":         application.Version,
			"publisher":       application.Publisher,
			"executable_path": application.ExecutablePath,
			"file_types":      flattenFileTypes(application.FileTypes),
			//"other_file_types": flattenFileTypes(application.OtherFileTypes),
		},
	}
}

func flattenFileTypes(types []FileType) string {
	var parts []string
	for _, filetype := range types {
		if filetype.Description != "" {
			parts = append(parts, fmt.Sprintf("%s: %s", filetype.Type, filetype.Description))
		} else {
			parts = append(parts, filetype.Type)
		}
	}

	return strings.Join(parts, "; ")
}

func getSessions(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("Sessions")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var sessions []Session
	err = json.Unmarshal([]byte(response), &sessions)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, session := range sessions {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createSessionFields(session),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

func createSessionFields(session Session) mapstr.M {
	return mapstr.M{
		"session": mapstr.M{
			"id":              session.ID,
			"user_id":         session.UserID,
			"broker_user_id":  session.BrokerUserID,
			"access_group_id": session.AccessGroupID,
			"machine_id":      session.MachineID,
			"desktop_pool_id": session.DesktopPoolID,
			"agent_version":   session.AgentVersion,
			"client_data": mapstr.M{
				"location_id": session.ClientData.LocationID,
				"type":        session.ClientData.Type,
				"address":     session.ClientData.Address,
				"name":        session.ClientData.Name,
				"version":     session.ClientData.Version,
			},
			"security_gateway_data": mapstr.M{
				"domain_name": session.SecurityGatewayData.DomainName,
				"address":     session.SecurityGatewayData.Address,
				"location":    session.SecurityGatewayData.Location,
			},
			"session_type":          session.SessionType,
			"session_protocol":      session.SessionProtocol,
			"session_state":         session.SessionState,
			"start_time":            session.StartTime,
			"disconnected_time":     session.DisconnectedTime,
			"last_session_duration": session.LastSessionDurationMs,
			"resourced_remotely":    session.ResourcedRemotely,
			"unauthenticated":       session.Unauthenticated,
			"idle_duration":         session.IdleDuration,
		},
	}
}

func getGateways(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("Gateways")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var gateways []Gateway
	err = json.Unmarshal([]byte(response), &gateways)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, gateway := range gateways {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createGatewayFields(gateway),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}
func createGatewayFields(gateway Gateway) mapstr.M {
	return mapstr.M{
		"gateway": mapstr.M{
			"id":                      gateway.ID,
			"name":                    gateway.Name,
			"status":                  gateway.Status,
			"active_connection_count": gateway.ActiveConnectionCount,
			"pcoip_connection_count":  gateway.PCoIPConnectionCount,
			"blast_connection_count":  gateway.BlastConnectionCount,
			"details": mapstr.M{
				"type":     gateway.Details.Type,
				"address":  gateway.Details.Address,
				"internal": gateway.Details.Internal,
				"version":  gateway.Details.Version,
			},
		},
	}
}

func getVirtualCenters(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("VirtualCenters")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var virtualCenters []VirtualCenter
	err = json.Unmarshal([]byte(response), &virtualCenters)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, vc := range virtualCenters {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createVirtualCenterFields(vc),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}
func createVirtualCenterFields(vc VirtualCenter) mapstr.M {
	return mapstr.M{
		"virtual_center": mapstr.M{
			"id":                            vc.ID,
			"version":                       vc.Version,
			"description":                   vc.Description,
			"instance_uuid":                 vc.InstanceUUID,
			"server_name":                   vc.ServerName,
			"port":                          vc.Port,
			"use_ssl":                       vc.UseSSL,
			"user_name":                     vc.UserName,
			"se_sparse_reclamation_enabled": vc.SeSparseReclamationEnabled,
			"enabled":                       vc.Enabled,
			"vmc_deployment":                vc.VMCDeployment,
			"limits": mapstr.M{
				"provisioning_limit":                      vc.Limits.ProvisioningLimit,
				"power_operations_limit":                  vc.Limits.PowerOperationsLimit,
				"instant_clone_engine_provisioning_limit": vc.Limits.InstantCloneEngineProvisioningLimit,
			},
			"storage_accelerator_data": mapstr.M{
				"enabled":               vc.StorageAcceleratorData.Enabled,
				"default_cache_size_mb": vc.StorageAcceleratorData.DefaultCacheSizeMB,
			},
			"certificate_override": mapstr.M{
				"certificate": vc.CertificateOverride.Certificate,
				"type":        vc.CertificateOverride.Type,
			},
		},
	}
}

func getMachines(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("Machines")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var machines []Machine
	err = json.Unmarshal([]byte(response), &machines)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, machine := range machines {
		// if we have virtual disks, create a separate event for each disk
		if len(machine.ManagedMachineData.VirtualDisks) > 0 {
			for _, disk := range machine.ManagedMachineData.VirtualDisks {
				diskFields := mapstr.M{
					"path":           disk.Path,
					"datastore_path": disk.DatastorePath,
					"capacity_mb":    disk.CapacityMB,
				}
				machineFields := createMachineFields(machine)
				managedData := createManagedMachineDataFields(machine.ManagedMachineData)
				managedData["virtual_disk"] = diskFields
				machineFields["managed_machine_data"] = managedData

				event := mb.Event{
					Timestamp:       timestamp,
					MetricSetFields: machineFields,
					RootFields:      createECSFields(m),
				}
				events = append(events, event)
			}
		} else {
			machineFields := createMachineFields(machine)
			managedData := createManagedMachineDataFields(machine.ManagedMachineData)
			machineFields["managed_machine_data"] = managedData

			event := mb.Event{
				Timestamp:       timestamp,
				MetricSetFields: machineFields,
				RootFields:      createECSFields(m),
			}
			events = append(events, event)
		}
	}

	return events, nil
}

func createMachineFields(machine Machine) mapstr.M {
	return mapstr.M{
		"machine": mapstr.M{
			"id":                                   machine.ID,
			"name":                                 machine.Name,
			"dns_name":                             machine.DNSName,
			"desktop_pool_id":                      machine.DesktopPoolID,
			"state":                                machine.State,
			"type":                                 machine.Type,
			"operating_system":                     machine.OperatingSystem,
			"operating_system_architecture":        machine.OperatingSystemArchitecture,
			"agent_version":                        machine.AgentVersion,
			"agent_build_number":                   machine.AgentBuildNumber,
			"remote_experience_agent_build_number": machine.RemoteExperienceAgentBuildNumber,
			"message_security_mode":                machine.MessageSecurityMode,
			"message_security_enhanced_mode_supported": machine.MessageSecurityEnhancedModeSupported,
			"pairing_state":                   machine.PairingState,
			"configured_by_connection_server": machine.ConfiguredByConnectionServer,
			"user_ids":                        machine.UserIDs,
		},
	}
}

func createManagedMachineDataFields(data ManagedMachineData) mapstr.M {
	return mapstr.M{

		"virtual_center_id":           data.VirtualCenterID,
		"host_name":                   data.HostName,
		"path":                        data.Path,
		"virtual_machine_power_state": data.VirtualMachinePowerState,
		"storage_accelerator_state":   data.StorageAcceleratorState,
		"memory_mb":                   data.MemoryMB,
		"missing_in_vcenter":          data.MissingInVCenter,
		"in_hold_customization":       data.InHoldCustomization,
		"create_time":                 data.CreateTime,
		"in_maintenance_mode":         data.InMaintenanceMode,
	}

}

func getRDSServers(m *MetricSet) ([]mb.Event, error) {

	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("RDSServers")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var rdsServers []RDSServer
	err = json.Unmarshal([]byte(response), &rdsServers)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, server := range rdsServers {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createRDSServerFields(server),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

func getFarms(m *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("Farms")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var farms []Farm
	err = json.Unmarshal([]byte(response), &farms)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, farm := range farms {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createFarmFields(farm),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

func createFarmFields(farm Farm) mapstr.M {
	return mapstr.M{
		"farm": mapstr.M{
			"id":           farm.ID,
			"name":         farm.Name,
			"display_name": farm.DisplayName,
			"description":  farm.Description,
			"enabled":      farm.Enabled,
			"source":       farm.Source,
			"type":         farm.Type,
			"settings": mapstr.M{
				"delete_in_progress": farm.Settings.DeleteInProgress,
				"desktop_id":         farm.Settings.DesktopID,
				"display_protocol_settings": mapstr.M{
					"allow_display_protocol_override": farm.Settings.DisplayProtocolSettings.AllowDisplayProtocolOverride,
					"default_display_protocol":        farm.Settings.DisplayProtocolSettings.DefaultDisplayProtocol,
					"grid_vgpus_enabled":              farm.Settings.DisplayProtocolSettings.GridVGPUsEnabled,
					"html_access_enabled":             farm.Settings.DisplayProtocolSettings.HTMLAccessEnabled,
					"session_collaboration_enabled":   farm.Settings.DisplayProtocolSettings.SessionCollaborationEnabled,
					"vgpu_grid_profile":               farm.Settings.DisplayProtocolSettings.VGPUGridProfile,
				},
				"load_balancer_settings": mapstr.M{
					"custom_script_in_use": farm.Settings.LoadBalancerSettings.CustomScriptInUse,
					"lb_metric_settings": mapstr.M{
						"cpu_threshold":                farm.Settings.LoadBalancerSettings.LBMetricSettings.CPUThreshold,
						"disk_queue_length_threshold":  farm.Settings.LoadBalancerSettings.LBMetricSettings.DiskQueueLengthThreshold,
						"disk_read_latency_threshold":  farm.Settings.LoadBalancerSettings.LBMetricSettings.DiskReadLatencyThreshold,
						"disk_write_latency_threshold": farm.Settings.LoadBalancerSettings.LBMetricSettings.DiskWriteLatencyThreshold,
						"include_session_count":        farm.Settings.LoadBalancerSettings.LBMetricSettings.IncludeSessionCount,
						"memory_threshold":             farm.Settings.LoadBalancerSettings.LBMetricSettings.MemoryThreshold,
					},
				},
				"server_error_threshold": farm.Settings.ServerErrorThreshold,
				"session_settings": mapstr.M{
					"disconnected_session_timeout_minutes": farm.Settings.SessionSettings.DisconnectedSessionTimeoutMinutes,
					"disconnected_session_timeout_policy":  farm.Settings.SessionSettings.DisconnectedSessionTimeoutPolicy,
					"empty_session_timeout_minutes":        farm.Settings.SessionSettings.EmptySessionTimeoutMinutes,
					"empty_session_timeout_policy":         farm.Settings.SessionSettings.EmptySessionTimeoutPolicy,
					"logoff_after_timeout":                 farm.Settings.SessionSettings.LogoffAfterTimeout,
					"pre_launch_session_timeout_minutes":   farm.Settings.SessionSettings.PreLaunchSessionTimeoutMinutes,
					"pre_launch_session_timeout_policy":    farm.Settings.SessionSettings.PreLaunchSessionTimeoutPolicy,
				},
			},
		},
	}
}

func createRDSServerFields(server RDSServer) mapstr.M {
	return mapstr.M{
		"rds_server": mapstr.M{
			"id":                                   server.ID,
			"name":                                 server.Name,
			"state":                                server.State,
			"enabled":                              server.Enabled,
			"farm_id":                              server.FarmID,
			"access_group_id":                      server.AccessGroupID,
			"description":                          server.Description,
			"dns_name":                             server.DNSName,
			"operating_system":                     server.OperatingSystem,
			"agent_version":                        server.AgentVersion,
			"agent_build_number":                   server.AgentBuildNumber,
			"remote_experience_agent_version":      server.RemoteExperienceAgentVersion,
			"remote_experience_agent_build_number": server.RemoteExperienceAgentBuildNumber,
			"message_security_mode":                server.MessageSecurityMode,
			"message_security_enhanced_mode_supported": server.MessageSecurityEnhancedModeSupported,
			"load_index":                         server.LoadIndex,
			"load_preference":                    server.LoadPreference,
			"session_count":                      server.SessionCount,
			"max_sessions_count":                 server.MaxSessionsCount,
			"max_sessions_count_configured":      server.MaxSessionsCountConfigured,
			"max_sessions_type":                  server.MaxSessionsType,
			"max_sessions_type_configured":       server.MaxSessionsTypeConfigured,
			"logoff_policy":                      server.LogoffPolicy,
			"operation":                          server.Operation,
			"operation_state":                    server.OperationState,
			"base_vm_id":                         server.BaseVMID,
			"base_vm_snapshot_id":                server.BaseVMSnapshotID,
			"pending_base_vm_id":                 server.PendingBaseVMID,
			"pending_base_vm_snapshot_id":        server.PendingBaseVMSnapshotID,
			"image_management_stream_id":         server.ImageManagementStreamID,
			"image_management_tag_id":            server.ImageManagementTagID,
			"pending_image_management_stream_id": server.PendingImageManagementStreamID,
			"pending_image_management_tag_id":    server.PendingImageManagementTagID,
		},
	}
}

func getCertificateData(m *MetricSet) ([]mb.Event, error) {

	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("CertificateData")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var certificates []CertificateData
	err = json.Unmarshal([]byte(response), &certificates)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, certificate := range certificates {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createCertificateDataFields(certificate),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

func createCertificateDataFields(certificate CertificateData) mapstr.M {
	return mapstr.M{
		"certificate_data": mapstr.M{
			"certificate_usage":             certificate.CertificateUsage,
			"dns_subject_alternative_names": certificate.DNSSubjectAlternativeNames,
			"in_use":                        certificate.InUse,
			"invalid_reasons":               certificate.InvalidReasons,
			"is_valid":                      certificate.IsValid,
			"issuer_name":                   certificate.IssuerName,
			"serial_number":                 certificate.SerialNumber,
			"sha1_thumbprint":               certificate.SHA1Thumbprint,
			"signature_algorithm":           certificate.SignatureAlgorithm,
			"subject_name":                  certificate.SubjectName,
			"valid_from":                    certificate.ValidFrom,
			"valid_until":                   certificate.ValidUntil,
		},
	}
}

func getLicenseData(m *MetricSet) ([]mb.Event, error) {

	timestamp := time.Now().UTC()
	client := m.horizonClient
	endpoint, err := getEndpoint("LicenseData")
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var licenses []LicenseData
	err = json.Unmarshal([]byte(response), &licenses)
	if err != nil {
		return nil, err
	}

	var events []mb.Event
	for _, license := range licenses {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createLicenseDataFields(license),
			RootFields:      createECSFields(m),
		}
		events = append(events, event)
	}

	return events, nil
}

func createLicenseDataFields(license LicenseData) mapstr.M {
	return mapstr.M{
		"license_data": mapstr.M{
			"application_pool_launch_enabled": license.ApplicationPoolLaunchEnabled,
			"desktop_pool_launch_enabled":     license.DesktopPoolLaunchEnabled,
			"expiration_time":                 license.ExpirationTime,
			"grace_period_days":               license.GracePeriodDays,
			"help_desk_enabled":               license.HelpDeskEnabled,
			"instant_clone_enabled":           license.InstantCloneEnabled,
			"license_edition":                 license.LicenseEdition,
			"license_health":                  license.LicenseHealth,
			"license_key":                     license.LicenseKey,
			"license_mode":                    license.LicenseMode,
			"licensed":                        license.Licensed,
			"session_collaboration_enabled":   license.SessionCollaborationEnabled,
			"subscription_slice_expiry":       license.SubscriptionSliceExpiry,
			"usage_model":                     license.UsageModel,
		},
	}
}

func createECSFields(ms *MetricSet) mapstr.M {
	//dataset := fmt.Sprintf("%s.%s", ms.Module().Name(), ms.Name())

	return mapstr.M{
		"observer": mapstr.M{
			"hostname": ms.config.HostInfo.Hostname,
			"ip":       ms.config.HostInfo.IP,
			"type":     "virtual-desktop-infrastructure",
			"vendor":   "VMWare",
		},
	}
}
