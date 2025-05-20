package health

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func createECSFields(ms *MetricSet) mapstr.M {
	//dataset := fmt.Sprintf("%s.%s", ms.Module().Name(), ms.Name())

	return mapstr.M{
		"observer": mapstr.M{
			"hostname": ms.config.HostInfo.Hostname,
			"ip":       ms.config.HostInfo.IP,
			"type":     "software-defined-networking",
			"vendor":   "VMWare",
			"product":  "NSX-T",
		},
	}
}

func getClusterNodesEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Cluster Nodes")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var clusterResponse ClusterNodesResponse
	err = json.Unmarshal([]byte(response), &clusterResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, node := range clusterResponse.Results {
		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createClusterNodeFields(node),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createClusterNodeFields(node ClusterNode) mapstr.M {
	return mapstr.M{
		"id":                         node.ID,
		"display_name":               node.DisplayName,
		"external_id":                node.ExternalID,
		"appliance_mgmt_listen_addr": node.ApplianceMgmtListenAddr,
		"resource_type":              node.ResourceType,
		"manager_role":               node.ManagerRole,
		"controller_role":            node.ControllerRole,
		"create_time":                node.CreateTime,
		"create_user":                node.CreateUser,
		"last_modified_time":         node.LastModifiedTime,
		"last_modified_user":         node.LastModifiedUser,
		"protection":                 node.Protection,
		"revision":                   node.Revision,
		"system_owned":               node.SystemOwned,
	}
}

func getClusterStatusEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Cluster Status")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var clusterStatus ClusterStatus
	err = json.Unmarshal([]byte(response), &clusterStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	event := mb.Event{
		Timestamp:       timestamp,
		MetricSetFields: createClusterStatusFields(clusterStatus),
		RootFields:      createECSFields(ms),
	}
	events = append(events, event)
	// TODO: Add logic to create events from cluster status
	return events, nil
}

func createClusterStatusFields(clusterStatus ClusterStatus) mapstr.M {
	return mapstr.M{
		"cluster_id": clusterStatus.ClusterID,
		// TODO
	}
}

func getEdgeClustersEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Edge Clusters")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var clusterResponse EdgeClustersResponse
	err = json.Unmarshal([]byte(response), &clusterResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, cluster := range clusterResponse.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createEdgeClusterFields(cluster),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createEdgeClusterFields(cluster EdgeCluster) mapstr.M {
	return mapstr.M{
		"create_time":                  cluster.CreateTime,
		"create_user":                  cluster.CreateUser,
		"last_modified_time":           cluster.LastModifiedTime,
		"last_modified_user":           cluster.LastModifiedUser,
		"protection":                   cluster.Protection,
		"revision":                     cluster.Revision,
		"system_owned":                 cluster.SystemOwned,
		"id":                           cluster.ID,
		"display_name":                 cluster.DisplayName,
		"description":                  cluster.Description,
		"deployment_type":              cluster.DeploymentType,
		"enable_inter_site_forwarding": cluster.EnableInterSiteForwarding,
		"member_node_type":             cluster.MemberNodeType,
		// FIXME members array
		"members":                  cluster.Members,
		"cluster_profile_bindings": cluster.ClusterProfileBindings,
		"allocation_rules":         cluster.AllocationRules,
		"resource_type":            cluster.ResourceType,
		"tags":                     cluster.Tags,
	}
}

func getFirewallSectionsEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Firewall Sections")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var firewalls FirewallSectionList
	err = json.Unmarshal([]byte(response), &firewalls)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, firewallSection := range firewalls.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createFirewallSectionFields(firewallSection),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createFirewallSectionFields(firewallSection FirewallSection) mapstr.M {
	return mapstr.M{
		"id":                 firewallSection.ID,
		"display_name":       firewallSection.DisplayName,
		"description":        firewallSection.Description,
		"comments":           firewallSection.Comments,
		"resource_type":      firewallSection.ResourceType,
		"category":           firewallSection.Category,
		"section_type":       firewallSection.SectionType,
		"enforced_on":        firewallSection.EnforcedOn,
		"is_default":         firewallSection.IsDefault,
		"locked":             firewallSection.Locked,
		"lock_modified_by":   firewallSection.LockModifiedBy,
		"lock_modified_time": firewallSection.LockModifiedTime,
		"stateful":           firewallSection.Stateful,
		"tcp_strict":         firewallSection.TcpStrict,
		"rule_count":         firewallSection.RuleCount,
		"priority":           firewallSection.Priority,
		"autoplumbed":        firewallSection.AutoPlumbed,
		"applied_tos":        firewallSection.AppliedTos,
		"tags":               firewallSection.Tags,
		"create_time":        firewallSection.CreateTime,
		"create_user":        firewallSection.CreateUser,
		"last_modified_time": firewallSection.LastModifiedTime,
		"last_modified_user": firewallSection.LastModifiedUser,
		"protection":         firewallSection.Protection,
		"revision":           firewallSection.Revision,
		"system_owned":       firewallSection.SystemOwned,
	}
}

func getLogicalRouterPortsEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Logical Router Ports")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var routerPorts LogicalRouterPortList
	err = json.Unmarshal([]byte(response), &routerPorts)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, routerPort := range routerPorts.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createRouterPortFields(routerPort),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createRouterPortFields(routerPort LogicalRouterPort) mapstr.M {
	return mapstr.M{
		"id":                            routerPort.ID,
		"display_name":                  routerPort.DisplayName,
		"description":                   routerPort.Description,
		"resource_type":                 routerPort.ResourceType,
		"logical_router_id":             routerPort.LogicalRouterID,
		"mac_address":                   routerPort.MacAddress,
		"subnets":                       routerPort.Subnets,
		"linked_logical_router_port_id": routerPort.LinkedLogicalRouterPortID,
		"linked_logical_switch_port_id": routerPort.LinkedLogicalSwitchPortID,
		"edge_cluster_member_index":     routerPort.EdgeClusterMemberIndex,
		"enable_multicast":              routerPort.EnableMulticast,
		"urpf_mode":                     routerPort.UrpFMode,
		"mode":                          routerPort.Mode,
		"mtu":                           routerPort.MTU,
		"tags":                          routerPort.Tags,
		"service_bindings":              routerPort.ServiceBindings,
	}
}

func getNodeNetworkInterfacesEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Node Network Interfaces")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var interfaces NetworkInterfaceList
	err = json.Unmarshal([]byte(response), &interfaces)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, iface := range interfaces.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createNetworkInterfaceFields(iface),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createNetworkInterfaceFields(iface NetworkInterface) mapstr.M {
	return mapstr.M{
		"schema":            iface.Schema,
		"self":              iface.Self,
		"admin_status":      iface.AdminStatus,
		"broadcast_address": iface.BroadcastAddress,
		"default_gateway":   iface.DefaultGateway,
		"interface_id":      iface.InterfaceID,
		"ip_addresses":      iface.IPAddresses,
		"ip_configuration":  iface.IPConfiguration,
		"link_status":       iface.LinkStatus,
		"mtu":               iface.MTU,
		"physical_address":  iface.PhysicalAddress,
	}
}

func getIPPoolsEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("IP Pools")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var pools IPPoolsResponse
	err = json.Unmarshal([]byte(response), &pools)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, pool := range pools.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createIPPoolFields(pool),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createIPPoolFields(pool IPPool) mapstr.M {
	return mapstr.M{
		"id":                 pool.ID,
		"display_name":       pool.DisplayName,
		"description":        pool.Description,
		"pool_usage":         pool.PoolUsage,
		"resource_type":      pool.ResourceType,
		"subnets":            pool.Subnets,
		"tags":               pool.Tags,
		"create_time":        pool.CreateTime,
		"create_user":        pool.CreateUser,
		"last_modified_time": pool.LastModifiedTime,
		"last_modified_user": pool.LastModifiedUser,
		"protection":         pool.Protection,
		"revision":           pool.Revision,
		"system_owned":       pool.SystemOwned,
	}
}

func getTransportNodesEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Transport Nodes")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var nodes TransportNodesResponse
	err = json.Unmarshal([]byte(response), &nodes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, node := range nodes.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createTransportNodeFields(node),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createTransportNodeFields(node TransportNode) mapstr.M {
	return mapstr.M{
		"create_time":              node.CreateTime,
		"create_user":              node.CreateUser,
		"last_modified_time":       node.LastModifiedTime,
		"last_modified_user":       node.LastModifiedUser,
		"protection":               node.Protection,
		"revision":                 node.Revision,
		"system_owned":             node.SystemOwned,
		"id":                       node.ID,
		"node_id":                  node.NodeID,
		"display_name":             node.DisplayName,
		"description":              node.Description,
		"failure_domain_id":        node.FailureDomainID,
		"is_overridden":            node.IsOverridden,
		"maintenance_mode":         node.MaintenanceMode,
		"resource_type":            node.ResourceType,
		"tags":                     node.Tags,
		"host_switch_spec":         node.HostSwitchSpec,
		"transport_zone_endpoints": node.TransportZoneEndpoints,
		"node_deployment_info":     node.NodeDeploymentInfo,
	}
}

func getTransportZonesEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Transport Zones")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var zones TransportZonesResponse
	err = json.Unmarshal([]byte(response), &zones)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, zone := range zones.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createTransportZoneFields(zone),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createTransportZoneFields(zone TransportZone) mapstr.M {
	return mapstr.M{
		"create_time":                 zone.CreateTime,
		"create_user":                 zone.CreateUser,
		"last_modified_time":          zone.LastModifiedTime,
		"last_modified_user":          zone.LastModifiedUser,
		"protection":                  zone.Protection,
		"revision":                    zone.Revision,
		"schema":                      zone.Schema,
		"system_owned":                zone.SystemOwned,
		"id":                          zone.ID,
		"display_name":                zone.DisplayName,
		"host_switch_id":              zone.HostSwitchID,
		"host_switch_name":            zone.HostSwitchName,
		"host_switch_mode":            zone.HostSwitchMode,
		"is_default":                  zone.IsDefault,
		"nested_nsx":                  zone.NestedNSX,
		"resource_type":               zone.ResourceType,
		"transport_type":              zone.TransportType,
		"tags":                        zone.Tags,
		"transport_zone_profile_ids":  zone.TransportZoneProfileIDs,
		"uplink_teaming_policy_names": zone.UplinkTeamingPolicyNames,
	}
}

func getClusterBackupHistoryEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Cluster Backup History")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var backupHistory BackupHistory
	err = json.Unmarshal([]byte(response), &backupHistory)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, backup := range backupHistory.ClusterBackupStatuses {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createBackupStatusFields(backup, "cluster"),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}
	for _, backup := range backupHistory.NodeBackupStatuses {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createBackupStatusFields(backup, "node"),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}
	for _, backup := range backupHistory.InventoryBackupStatuses {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createBackupStatusFields(backup, "inventory"),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createBackupStatusFields(backupStatus BackupStatus, backupType string) mapstr.M {
	return mapstr.M{
		"backup_type":   backupType,
		"backup_id":     backupStatus.BackupID,
		"start_time":    backupStatus.StartTime,
		"end_time":      backupStatus.EndTime,
		"success":       backupStatus.Success,
		"error_code":    backupStatus.ErrorCode,
		"error_message": backupStatus.ErrorMessage,
		"type":          backupType,
	}
}

func getInfraTier0sEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Infrastructure Tier-0s")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var tier0s Tier0ListResponse
	err = json.Unmarshal([]byte(response), &tier0s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, tier0 := range tier0s.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createTier0Fields(tier0),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}

func createTier0Fields(tier0 Tier0) mapstr.M {
	return mapstr.M{
		"create_time":              tier0.CreateTime,
		"create_user":              tier0.CreateUser,
		"last_modified_time":       tier0.LastModifiedTime,
		"last_modified_user":       tier0.LastModifiedUser,
		"protection":               tier0.Protection,
		"revision":                 tier0.Revision,
		"system_owned":             tier0.SystemOwned,
		"id":                       tier0.ID,
		"display_name":             tier0.DisplayName,
		"description":              tier0.Description,
		"resource_type":            tier0.ResourceType,
		"path":                     tier0.Path,
		"parent_path":              tier0.ParentPath,
		"relative_path":            tier0.RelativePath,
		"marked_for_delete":        tier0.MarkedForDelete,
		"overridden":               tier0.Overridden,
		"default_rule_logging":     tier0.DefaultRuleLogging,
		"disable_firewall":         tier0.DisableFirewall,
		"force_whitelisting":       tier0.ForceWhitelisting,
		"failover_mode":            tier0.FailoverMode,
		"ha_mode":                  tier0.HAMode,
		"unique_id":                tier0.UniqueID,
		"advanced_config":          tier0.AdvancedConfig,
		"internal_transit_subnets": tier0.InternalTransitSubnets,
		"transit_subnets":          tier0.TransitSubnets,
		"ipv6_profile_paths":       tier0.IPv6ProfilePaths,
		"tags":                     tier0.Tags,
	}
}

func getInfraTier1sEvents(ms *MetricSet) ([]mb.Event, error) {
	timestamp := time.Now().UTC()
	client := ms.nsxtClient
	endpoint, err := getEndpoint("Infrastructure Tier-1s")
	if err != nil {
		return nil, err
	}

	response, err := client.Get(endpoint.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var tier1s Tier1ListResponse
	err = json.Unmarshal([]byte(response), &tier1s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster nodes response")
	}

	var events []mb.Event
	for _, tier1 := range tier1s.Results {

		event := mb.Event{
			Timestamp:       timestamp,
			MetricSetFields: createTier1Fields(tier1),
			RootFields:      createECSFields(ms),
		}
		events = append(events, event)
	}

	return events, nil
}
func createTier1Fields(tier1 Tier1) mapstr.M {
	return mapstr.M{
		"create_time":               tier1.CreateTime,
		"create_user":               tier1.CreateUser,
		"last_modified_time":        tier1.LastModifiedTime,
		"last_modified_user":        tier1.LastModifiedUser,
		"protection":                tier1.Protection,
		"revision":                  tier1.Revision,
		"system_owned":              tier1.SystemOwned,
		"id":                        tier1.ID,
		"display_name":              tier1.DisplayName,
		"description":               tier1.Description,
		"resource_type":             tier1.ResourceType,
		"path":                      tier1.Path,
		"parent_path":               tier1.ParentPath,
		"relative_path":             tier1.RelativePath,
		"marked_for_delete":         tier1.MarkedForDelete,
		"overridden":                tier1.Overridden,
		"default_rule_logging":      tier1.DefaultRuleLogging,
		"disable_firewall":          tier1.DisableFirewall,
		"force_whitelisting":        tier1.ForceWhitelisting,
		"failover_mode":             tier1.FailoverMode,
		"enable_standby_relocation": tier1.EnableStandbyRelocation,
		"pool_allocation":           tier1.PoolAllocation,
		"ipv6_profile_paths":        tier1.IPv6ProfilePaths,
		"route_advertisement_types": tier1.RouteAdvertisementTypes,
		"tier0_path":                tier1.Tier0Path,
		"unique_id":                 tier1.UniqueID,
		"tags":                      tier1.Tags,
	}
}
