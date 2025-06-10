package health

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/nsxt"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func createECSFields(ms *MetricSet) mapstr.M {
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
	manager_role, err := nsxt.ToJSONString(node.ManagerRole)
	if err != nil {
		manager_role = ""
	}
	controller_role, err := nsxt.ToJSONString(node.ControllerRole)
	if err != nil {
		controller_role = ""
	}

	return mapstr.M{
		"cluster_node": mapstr.M{
			"id":                         node.ID,
			"display_name":               node.DisplayName,
			"external_id":                node.ExternalID,
			"appliance_mgmt_listen_addr": node.ApplianceMgmtListenAddr,
			"resource_type":              node.ResourceType,
			"manager_role":               manager_role,
			"controller_role":            controller_role,
			"create_time":                node.CreateTime,
			"create_user":                node.CreateUser,
			"last_modified_time":         node.LastModifiedTime,
			"last_modified_user":         node.LastModifiedUser,
			"protection":                 node.Protection,
			"revision":                   node.Revision,
			"system_owned":               node.SystemOwned,
		},
	}
}

// LM gets the Control Cluster Status (control_cluster_status), The Management Cluster Status (mgmt_cluster_status)
// and the count of Management Cluster Nodes.
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
	clusterId := clusterStatus.ClusterID
	clStatus := clusterStatus.ControlClusterStatus.Status
	mgrStatus := clusterStatus.MgmtClusterStatus.Status
	mgrNodeCount := len(clusterStatus.MgmtClusterStatus.OnlineNodes)
	for _, group := range clusterStatus.DetailedStatus.Groups {
		leaders, err := nsxt.ToJSONString(group.Leaders)
		if err != nil {
			leaders = ""
		}

		for _, member := range group.Members {
			memberFields := mapstr.M{
				"fqdn":   member.FQDN,
				"ip":     member.IP,
				"status": member.Status,
				"uuid":   member.UUID,
			}
			groupFields := mapstr.M{
				"cluster_status": mapstr.M{
					"cluster_id":              clusterId,
					"control_cluster_status":  clStatus,
					"mgmt_cluster_status":     mgrStatus,
					"mgmt_cluster_node_count": mgrNodeCount,
					"group_id":                group.GroupID,
					"group_status":            group.GroupStatus,
					"group_type":              group.GroupType,
					"leaders":                 leaders,
					"member_data":             memberFields,
				},
			}

			event := mb.Event{
				Timestamp:       timestamp,
				MetricSetFields: groupFields,
				RootFields:      createECSFields(ms),
			}

			events = append(events, event)
		}
	}

	return events, nil
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
	members, err := nsxt.ToJSONString(cluster.Members)
	if err != nil {
		members = ""
	}
	tags, err := nsxt.ToJSONString(cluster.Tags)
	if err != nil {
		tags = ""
	}

	cluster_profile_bindings, err := nsxt.ToJSONString(cluster.ClusterProfileBindings)
	if err != nil {
		cluster_profile_bindings = ""
	}
	return mapstr.M{
		"edge_cluster": mapstr.M{
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
			"members":                      members,
			"cluster_profile_bindings":     cluster_profile_bindings,
			"allocation_rules":             cluster.AllocationRules,
			"resource_type":                cluster.ResourceType,
			"tags":                         tags,
		},
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
	applied_tos, err := nsxt.ToJSONString(firewallSection.AppliedTos)
	if err != nil {
		applied_tos = ""
	}
	tags, err := nsxt.ToJSONString(firewallSection.Tags)
	if err != nil {
		tags = ""
	}
	return mapstr.M{
		"firewall_section": mapstr.M{
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
			"applied_tos":        applied_tos,
			"tags":               tags,
			"create_time":        firewallSection.CreateTime,
			"create_user":        firewallSection.CreateUser,
			"last_modified_time": firewallSection.LastModifiedTime,
			"last_modified_user": firewallSection.LastModifiedUser,
			"protection":         firewallSection.Protection,
			"revision":           firewallSection.Revision,
			"system_owned":       firewallSection.SystemOwned,
		},
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

	subnets, err := nsxt.ToJSONString(routerPort.Subnets)
	if err != nil {
		subnets = ""
	}

	service_bindings, err := nsxt.ToJSONString(routerPort.ServiceBindings)
	if err != nil {
		service_bindings = ""
	}
	tags, err := nsxt.ToJSONString(routerPort.Tags)
	if err != nil {
		tags = ""
	}
	routerPortFields := mapstr.M{
		"logical_router_port": mapstr.M{
			"id":                        routerPort.ID,
			"display_name":              routerPort.DisplayName,
			"description":               routerPort.Description,
			"resource_type":             routerPort.ResourceType,
			"logical_router_id":         routerPort.LogicalRouterID,
			"mac_address":               routerPort.MacAddress,
			"subnets":                   subnets,
			"edge_cluster_member_index": routerPort.EdgeClusterMemberIndex,
			"enable_multicast":          routerPort.EnableMulticast,
			"urpf_mode":                 routerPort.UrpFMode,
			"mode":                      routerPort.Mode,
			"mtu":                       routerPort.MTU,
			"tags":                      tags,
			"service_bindings":          service_bindings,
		},
	}

	// because linked_logical_router_port_id in the json can be either string or object
	if nil != routerPort.LinkedLogicalRouterPortID.Object {
		routerPortFields["linked_logical_router_port"] = mapstr.M{
			"is_valid":            routerPort.LinkedLogicalRouterPortID.Object.IsValid,
			"target_id":           routerPort.LinkedLogicalRouterPortID.Object.TargetID,
			"target_type":         routerPort.LinkedLogicalRouterPortID.Object.TargetType,
			"target_display_name": routerPort.LinkedLogicalRouterPortID.Object.TargetDisplayName,
		}
	} else {
		routerPortFields["linked_logical_router_port_id"] = routerPort.LinkedLogicalRouterPortID.ID
	}
	// because linked_logical_switch_port_id in the json can be either string or object
	if nil != routerPort.LinkedLogicalSwitchPortID.Object {
		routerPortFields["linked_logical_switch_port"] = mapstr.M{
			"is_valid":            routerPort.LinkedLogicalSwitchPortID.Object.IsValid,
			"target_id":           routerPort.LinkedLogicalSwitchPortID.Object.TargetID,
			"target_type":         routerPort.LinkedLogicalSwitchPortID.Object.TargetType,
			"target_display_name": routerPort.LinkedLogicalSwitchPortID.Object.TargetDisplayName,
		}
	} else {
		routerPortFields["inked_logical_switch_port_id"] = routerPort.LinkedLogicalSwitchPortID.ID
	}

	return routerPortFields
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
	ip_addresses, err := nsxt.ToJSONString(iface.IPAddresses)
	if err != nil {
		ip_addresses = ""
	}
	return mapstr.M{
		"network_interface": mapstr.M{
			"admin_status":      iface.AdminStatus,
			"broadcast_address": iface.BroadcastAddress,
			"default_gateway":   iface.DefaultGateway,
			"interface_id":      iface.InterfaceID,
			"ip_addresses":      ip_addresses,
			"ip_configuration":  iface.IPConfiguration,
			"link_status":       iface.LinkStatus,
			"mtu":               iface.MTU,
			"physical_address":  iface.PhysicalAddress,
		},
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
	subnets, err := nsxt.ToJSONString(pool.Subnets)
	if err != nil {
		subnets = "Subnets could not be parsed: " + err.Error()
	}
	tags, err := nsxt.ToJSONString(pool.Tags)
	if err != nil {
		tags = ""
	}
	return mapstr.M{
		"ip_pool": mapstr.M{
			"id":           pool.ID,
			"display_name": pool.DisplayName,
			"description":  pool.Description,
			"pool_usage": mapstr.M{
				"allocated_ids": pool.PoolUsage.AllocatedIDs,
				"free_ids":      pool.PoolUsage.FreeIDs,
				"total_ids":     pool.PoolUsage.TotalIDs,
			},
			"resource_type":      pool.ResourceType,
			"subnets":            subnets,
			"tags":               tags,
			"create_time":        pool.CreateTime,
			"create_user":        pool.CreateUser,
			"last_modified_time": pool.LastModifiedTime,
			"last_modified_user": pool.LastModifiedUser,
			"protection":         pool.Protection,
			"revision":           pool.Revision,
			"system_owned":       pool.SystemOwned,
		},
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
		// create an event for every host switch in the node
		for _, hostSwitch := range node.HostSwitchSpec.HostSwitches {

			nodeFields := createTransportNodeFields(node)
			hostSwitchFields := createHostSwitchFields(hostSwitch, node.HostSwitchSpec.ResourceType)

			// Safely access "transport_node" as mapstr.M and set "host_switch"
			if transportNode, ok := nodeFields["transport_node"].(mapstr.M); ok {
				transportNode["host_switch"] = hostSwitchFields
			} else {
				return nil, fmt.Errorf("expected transport_node to be mapstr.M, got %T", nodeFields["transport_node"])
			}

			event := mb.Event{
				Timestamp:       timestamp,
				MetricSetFields: nodeFields,
				RootFields:      createECSFields(ms),
			}
			events = append(events, event)
		}
	}

	return events, nil
}

func createHostSwitchFields(hostSwitch HostSwitch, resourceType string) mapstr.M {
	// As of this writing, there is no data availble to determine the types of these arrays,
	// so just put them in a string for now.
	cpu_config, err := nsxt.ToJSONString(hostSwitch.CPUConfig)
	if err != nil {
		cpu_config = "CPUConfig could not be parsed: " + err.Error()
	}

	pnics_uninstall_migration, err := nsxt.ToJSONString(hostSwitch.PnicsUninstallMigration)
	if err != nil {
		pnics_uninstall_migration = "PnicsUninstallMigration could not be parsed: " + err.Error()
	}

	vmk_install_migration, err := nsxt.ToJSONString(hostSwitch.VmkInstallMigration)
	if err != nil {
		vmk_install_migration = "VmkInstallMigration could not be parsed: " + err.Error()
	}
	vmk_uninstall_migration, err := nsxt.ToJSONString(hostSwitch.VmkUninstallMigration)
	if err != nil {
		vmk_uninstall_migration = "VmkUninstallMigration could not be parsed: " + err.Error()
	}

	host_switch_profile_ids, err := nsxt.ToJSONString(hostSwitch.HostSwitchProfileIDs)
	if err != nil {
		host_switch_profile_ids = ""
	}

	ip_list, err := nsxt.ToJSONString(hostSwitch.IPAssignmentSpec.IPList)
	if err != nil {
		ip_list = ""
	}

	pnics, err := nsxt.ToJSONString(hostSwitch.Pnics)
	if err != nil {
		pnics = ""
	}

	transport_zone_endpoints, err := nsxt.ToJSONString(hostSwitch.TransportZoneEndpoints)
	if err != nil {
		transport_zone_endpoints = ""
	}

	uplinks, err := nsxt.ToJSONString(hostSwitch.Uplinks)
	if err != nil {
		uplinks = ""
	}

	ipAssignmentSpec := mapstr.M{

		"resource_type":   hostSwitch.IPAssignmentSpec.ResourceType,
		"ip_pool_id":      hostSwitch.IPAssignmentSpec.IPPoolID,
		"default_gateway": hostSwitch.IPAssignmentSpec.DefaultGateway,
		"ip_list":         ip_list,
		"subnet_mask":     hostSwitch.IPAssignmentSpec.SubnetMask,
	}

	return mapstr.M{

		"resource_type":             resourceType,
		"cpu_config":                cpu_config,
		"host_switch_id":            hostSwitch.HostSwitchID,
		"host_switch_mode":          hostSwitch.HostSwitchMode,
		"host_switch_name":          hostSwitch.HostSwitchName,
		"host_switch_profile_ids":   host_switch_profile_ids,
		"host_switch_type":          hostSwitch.HostSwitchType,
		"ip_assignment_spec":        ipAssignmentSpec,
		"is_migrate_pnics":          hostSwitch.IsMigratePnics,
		"not_ready":                 hostSwitch.NotReady,
		"pnics":                     pnics,
		"pnics_uninstall_migration": pnics_uninstall_migration,
		"vmk_install_migration":     vmk_install_migration,
		"vmk_uninstall_migration":   vmk_uninstall_migration,
		"transport_zone_endpoints":  transport_zone_endpoints,
		"uplinks":                   uplinks,
	}
}
func createTransportNodeFields(node TransportNode) mapstr.M {
	deploymentInfo := createNodeDeploymentInfoFields(node.NodeDeploymentInfo)
	tags, err := nsxt.ToJSONString(node.Tags)
	if err != nil {
		tags = ""
	}
	transport_zone_endpoints, err := nsxt.ToJSONString(node.TransportZoneEndpoints)
	if err != nil {
		transport_zone_endpoints = ""
	}
	return mapstr.M{
		"transport_node": mapstr.M{
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
			"tags":                     tags,
			"transport_zone_endpoints": transport_zone_endpoints,
			"node_deployment_info":     deploymentInfo,
		},
	}
}

func createNodeDeploymentInfoFields(nodeDeploymentInfo NodeDeploymentInfo) mapstr.M {
	deploymentConfig, err := nsxt.ToJSONString(nodeDeploymentInfo.DeploymentConfig)
	if err != nil {
		deploymentConfig = "DeploymentConfig could not be parsed: " + err.Error()
	}

	nodeSettings, err := nsxt.ToJSONString(nodeDeploymentInfo.NodeSettings)
	if err != nil {
		nodeSettings = "NodeSettings could not be parsed: " + err.Error()
	}

	ip_addresses, err := nsxt.ToJSONString(nodeDeploymentInfo.IPAddresses)
	if err != nil {
		ip_addresses = "IPAddresses could not be parsed: " + err.Error()
	}

	discovered_ip_addresses, err := nsxt.ToJSONString(nodeDeploymentInfo.DiscoveredIPs)
	if err != nil {
		discovered_ip_addresses = ""
	}

	return mapstr.M{
		"create_time":             nodeDeploymentInfo.CreateTime,
		"create_user":             nodeDeploymentInfo.CreateUser,
		"last_modified_time":      nodeDeploymentInfo.LastModifiedTime,
		"last_modified_user":      nodeDeploymentInfo.LastModifiedUser,
		"protection":              nodeDeploymentInfo.Protection,
		"revision":                nodeDeploymentInfo.Revision,
		"system_owned":            nodeDeploymentInfo.SystemOwned,
		"resource_type":           nodeDeploymentInfo.ResourceType,
		"deployment_type":         nodeDeploymentInfo.DeploymentType,
		"deployment_config":       deploymentConfig,
		"display_name":            nodeDeploymentInfo.DisplayName,
		"description":             nodeDeploymentInfo.Description,
		"external_id":             nodeDeploymentInfo.ExternalID,
		"id":                      nodeDeploymentInfo.ID,
		"ip_addresses":            ip_addresses,
		"node_settings":           nodeSettings,
		"discovered_node_id":      nodeDeploymentInfo.DiscoveredNodeID,
		"fqdn":                    nodeDeploymentInfo.FQDN,
		"managed_by_server":       nodeDeploymentInfo.ManagedByServer,
		"discovered_ip_addresses": discovered_ip_addresses,
		"os_type":                 nodeDeploymentInfo.OSType,
		"os_version":              nodeDeploymentInfo.OSVersion,
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
	transport_zone_profile_ids, err := nsxt.ToJSONString(zone.TransportZoneProfileIDs)
	if err != nil {
		transport_zone_profile_ids = "TransportZoneProfileIDs could not be parsed: " + err.Error()
	}
	uplink_teaming_policy_names, err := nsxt.ToJSONString(zone.UplinkTeamingPolicyNames)
	if err != nil {
		uplink_teaming_policy_names = "UplinkTeamingPolicyNames could not be parsed: " + err.Error()
	}
	tags, err := nsxt.ToJSONString(zone.Tags)
	if err != nil {
		tags = ""
	}
	return mapstr.M{
		"transport_zone": mapstr.M{
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
			"tags":                        tags,
			"transport_zone_profile_ids":  transport_zone_profile_ids,
			"uplink_teaming_policy_names": uplink_teaming_policy_names,
		},
	}
}

// Returns events for the latest backup status of cluster, node and inventory
// The latest status is determined by the latest start time
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

	// Find the latest backup status for cluster, node and inventory
	clusterBackupStatus := latestStatusValue(backupHistory.ClusterBackupStatuses)
	nodeBackupStatus := latestStatusValue(backupHistory.NodeBackupStatuses)
	inventoryBackupStatus := latestStatusValue(backupHistory.InventoryBackupStatuses)

	clusterFields := createBackupStatusFields(clusterBackupStatus, "cluster")
	nodeFields := createBackupStatusFields(nodeBackupStatus, "node")
	inventoryFields := createBackupStatusFields(inventoryBackupStatus, "inventory")

	event := mb.Event{
		Timestamp: timestamp,
		MetricSetFields: mapstr.M{
			"backup_status": mapstr.M{
				"cluster":   clusterFields,
				"node":      nodeFields,
				"inventory": inventoryFields,
			},
		},
		RootFields: createECSFields(ms),
	}
	events = append(events, event)

	return events, nil
}

func createBackupStatusFields(backupStatus BackupStatus, backupType string) mapstr.M {
	return mapstr.M{
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
	internal_transit_subnets, err := nsxt.ToJSONString(tier0.InternalTransitSubnets)
	if err != nil {
		internal_transit_subnets = "InternalTransitSubnets could not be parsed: " + err.Error()
	}

	transit_subnets, err := nsxt.ToJSONString(tier0.TransitSubnets)
	if err != nil {
		transit_subnets = "TransitSubnets could not be parsed: " + err.Error()
	}

	ipv6_profile_paths, err := nsxt.ToJSONString(tier0.IPv6ProfilePaths)
	if err != nil {
		ipv6_profile_paths = "IPv6ProfilePaths could not be parsed: " + err.Error()
	}

	tags, err := nsxt.ToJSONString(tier0.Tags)
	if err != nil {
		tags = ""
	}
	advanced_config := mapstr.M{
		"connectivity":        tier0.AdvancedConfig.Connectivity,
		"forwarding_up_timer": tier0.AdvancedConfig.ForwardingUpTimer,
	}

	return mapstr.M{
		"tier0": mapstr.M{
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
			"advanced_config":          advanced_config,
			"internal_transit_subnets": internal_transit_subnets,
			"transit_subnets":          transit_subnets,
			"ipv6_profile_paths":       ipv6_profile_paths,
			"tags":                     tags,
		},
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
	ipv6_profile_paths, err := nsxt.ToJSONString(tier1.IPv6ProfilePaths)
	if err != nil {
		ipv6_profile_paths = "IPv6ProfilePaths could not be parsed: " + err.Error()
	}
	route_advertisement_types, err := nsxt.ToJSONString(tier1.RouteAdvertisementTypes)
	if err != nil {
		route_advertisement_types = "RouteAdvertisementTypes could not be parsed: " + err.Error()
	}
	tags, err := nsxt.ToJSONString(tier1.Tags)
	if err != nil {
		tags = ""
	}

	return mapstr.M{
		"tier1": mapstr.M{
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
			"ipv6_profile_paths":        ipv6_profile_paths,
			"route_advertisement_types": route_advertisement_types,
			"tier0_path":                tier1.Tier0Path,
			"unique_id":                 tier1.UniqueID,
			"tags":                      tags,
		},
	}
}

func latestStatusValue(stats []BackupStatus) BackupStatus {
	if len(stats) == 0 {
		return BackupStatus{}
	}

	// Sort the stats by EndTime in descending order and return the first one
	sort.Slice(stats, func(i, j int) bool {
		return time.Unix(stats[i].EndTime, 0).After(time.Unix(stats[j].EndTime, 0))
	})
	return stats[0]
}
