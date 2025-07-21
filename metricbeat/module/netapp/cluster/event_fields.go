package cluster

import (
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func createClusterFields(c Cluster) mapstr.M {

	dns_domains, err := netapp.ToJSONString(c.DNSDomains)
	if err != nil {
		dns_domains = ""
	}
	name_servers, err := netapp.ToJSONString(c.NameServers)
	if err != nil {
		name_servers = ""
	}

	ntp_servers, err := netapp.ToJSONString(c.NTPServers)
	if err != nil {
		ntp_servers = ""
	}

	management_interfaces, err := netapp.ToJSONString(c.ManagementInterfaces)
	if err != nil {
		management_interfaces = ""
	}

	fields := mapstr.M{
		"name":                          c.Name,
		"uuid":                          c.UUID,
		"location":                      c.Location,
		"contact":                       c.Contact,
		"version":                       c.Version,
		"dns_domains":                   dns_domains,
		"name_servers":                  name_servers,
		"ntp_servers":                   ntp_servers,
		"management_interfaces":         management_interfaces,
		"metric":                        c.Metric,
		"statistics":                    c.Statistics,
		"san_optimized":                 c.SANOptimized,
		"disaggregated":                 c.Disaggregated,
		"auto_enable_analytics":         c.AutoEnableAnalytics,
		"auto_enable_activity_tracking": c.AutoEnableActivityTrack,
	}

	if c.Timezone != nil {
		fields["timezone"] = c.Timezone
	}
	if c.Certificate != nil {
		fields["certificate"] = c.Certificate
	}
	if c.PeeringPolicy != nil {
		fields["peering_policy"] = c.PeeringPolicy
	}

	return fields
}

func createSensorFields(s Sensor) mapstr.M {
	fields := mapstr.M{
		"node":            netapp.CreateNamedObjectFields(s.Node),
		"index":           s.Index,
		"name":            s.Name,
		"type":            s.Type,
		"value":           s.Value,
		"value_units":     s.ValueUnits,
		"threshold_state": s.ThresholdState,
	}

	if s.CriticalLowThreshold != nil {
		fields["critical_low_threshold"] = *s.CriticalLowThreshold
	}
	if s.WarningLowThreshold != nil {
		fields["warning_low_threshold"] = *s.WarningLowThreshold
	}
	if s.WarningHighThreshold != nil {
		fields["warning_high_threshold"] = *s.WarningHighThreshold
	}
	if s.CriticalHighThreshold != nil {
		fields["critical_high_threshold"] = *s.CriticalHighThreshold
	}

	return fields
}

func createNodeFields(n ClusterNode) mapstr.M {
	managementInterfaces, err := netapp.ToJSONString(n.ManagementInterfaces)
	if err != nil {
		managementInterfaces = ""
	}
	clusterInterfaces, err := netapp.ToJSONString(n.ClusterInterfaces)
	if err != nil {
		clusterInterfaces = ""
	}

	fields := mapstr.M{
		"uuid":          n.UUID,
		"name":          n.Name,
		"serial_number": n.SerialNumber,
		"location":      n.Location,
		"owner":         n.Owner,
		"model":         n.Model,
		"system_id":     n.SystemID,
		"version": mapstr.M{
			"full":       n.Version.Full,
			"generation": n.Version.Generation,
			"major":      n.Version.Major,
			"minor":      n.Version.Minor,
		},
		"date":                  n.Date,
		"uptime":                n.Uptime,
		"state":                 n.State,
		"membership":            n.Membership,
		"management_interfaces": managementInterfaces,
		"cluster_interfaces":    clusterInterfaces,
		"storage_configuration": n.StorageConfiguration,
		"system_aggregate": mapstr.M{
			"uuid": n.SystemAggregate.UUID,
			"name": n.SystemAggregate.Name,
		},
		"controller":        createControllerFields(n.Controller),
		"service_processor": createServiceProcessorFields(n.ServiceProcessor),
		"nvram": mapstr.M{
			"id":            n.NVRAM.ID,
			"battery_state": n.NVRAM.BatteryState,
		},
		"external_cache": mapstr.M{
			"is_enabled":        n.ExternalCache.IsEnabled,
			"is_hya_enabled":    n.ExternalCache.IsHYAEnabled,
			"is_rewarm_enabled": n.ExternalCache.IsRewarmEnabled,
			"pcs_size":          n.ExternalCache.PCSSize,
		},
		"hw_assist": mapstr.M{
			"status": mapstr.M{
				"enabled": n.HWAssist.Status.Enabled,
				"local": mapstr.M{
					"state": n.HWAssist.Status.Local.State,
					"ip":    n.HWAssist.Status.Local.IP,
					"port":  n.HWAssist.Status.Local.Port,
				},
				"partner": mapstr.M{
					"state": n.HWAssist.Status.Partner.State,
					"ip":    n.HWAssist.Status.Partner.IP,
					"port":  n.HWAssist.Status.Partner.Port,
				},
			},
		},
		"anti_ransomware_version": n.AntiRansomwareVersion,
		"metric":                  n.Metric,
		"statistics":              n.Statistics,
	}

	return fields
}

func createServiceProcessorFields(sp NodeServiceProcessor) mapstr.M {

	fields := mapstr.M{
		"dhcp_enabled":       sp.DHCPEnabled,
		"state":              sp.State,
		"mac_address":        sp.MACAddress,
		"firmware_version":   sp.FirmwareVersion,
		"link_status":        sp.LinkStatus,
		"type":               sp.Type,
		"is_ip_configured":   sp.IsIPConfigured,
		"autoupdate_enabled": sp.AutoupdateEnabled,
		"last_update_state":  sp.LastUpdateState,
	}

	if sp.IPv4Interface != nil {
		fields["ipv4_interface"] = mapstr.M{
			"address":     sp.IPv4Interface.Address,
			"netmask":     sp.IPv4Interface.Netmask,
			"gateway":     sp.IPv4Interface.Gateway,
			"enabled":     sp.IPv4Interface.Enabled,
			"setup_state": sp.IPv4Interface.SetupState,
		}
	}
	if sp.IPv6Interface != nil {
		fields["ipv6_interface"] = mapstr.M{
			"enabled": sp.IPv6Interface.Enabled,
		}
	}
	if sp.SSHInfo != nil {
		fields["ssh_info"] = mapstr.M{
			"allowed_addresses": sp.SSHInfo.AllowedAddresses,
		}
	}
	if sp.Primary != nil {
		fields["primary"] = mapstr.M{
			"is_current": sp.Primary.IsCurrent,
			"state":      sp.Primary.State,
			"version":    sp.Primary.Version,
		}
	}
	if sp.Backup != nil {
		fields["backup"] = mapstr.M{
			"is_current": sp.Backup.IsCurrent,
			"state":      sp.Backup.State,
			"version":    sp.Backup.Version,
		}
	}
	if sp.APIService != nil {
		fields["api_service"] = mapstr.M{
			"enabled":      sp.APIService.Enabled,
			"limit_access": sp.APIService.LimitAccess,
			"port":         sp.APIService.Port,
		}
	}
	if sp.WebService != nil {
		fields["web_service"] = mapstr.M{
			"enabled":      sp.WebService.Enabled,
			"limit_access": sp.WebService.LimitAccess,
		}
	}

	return fields
}

func createControllerFields(c NodeController) mapstr.M {
	return mapstr.M{
		"board":            c.Board,
		"memory_size":      c.MemorySize,
		"over_temperature": c.OverTemperature,
		"failed_fan": mapstr.M{
			"count": c.FailedFan.Count,
			"message": mapstr.M{
				"message": c.FailedFan.Message.Message,
				"code":    c.FailedFan.Message.Code,
			},
		},
		"failed_power_supply": mapstr.M{
			"count": c.FailedPowerSupply.Count,
			"message": mapstr.M{
				"message": c.FailedPowerSupply.Message.Message,
				"code":    c.FailedPowerSupply.Message.Code,
			},
		},
		"cpu": mapstr.M{
			"firmware_release": c.CPU.FirmwareRelease,
			"processor":        c.CPU.Processor,
			"count":            c.CPU.Count,
		},
	}
}

func createNodeHAFields(ha NodeHA) mapstr.M {
	partners, err := netapp.ToJSONString(ha.Partners)
	if err != nil {
		partners = ""
	}
	ports, err := netapp.ToJSONString(ha.Ports)
	if err != nil {
		ports = ""
	}
	return mapstr.M{
		"enabled":       ha.Enabled,
		"auto_giveback": ha.AutoGiveback,
		"partners":      partners,

		"takeover": mapstr.M{
			"state": ha.Takeover.State,
		},
		"interconnect": mapstr.M{
			"adapter": ha.Interconnect.Adapter,
			"state":   ha.Interconnect.State,
		},
		"ports": ports,
	}
}

func createGivebackStatusFields(s NodeHAGBStatus) mapstr.M {
	return mapstr.M{
		"state":          s.State,
		"aggregate_name": s.Aggregate.Name,
	}
}

func getPropertyValue(props []CounterProperty, key string) string {
	for _, prop := range props {
		if prop.Name == key {
			return prop.Value
		}
	}
	return "unknown"
}

func createCounterTableFields(row CounterTableRow) mapstr.M {

	node_name := getPropertyValue(row.Properties, "node.name")
	properties, err := netapp.ToJSONString(row.Properties)
	if err != nil {
		properties = ""
	}

	return mapstr.M{
		"counter_table": row.CounterTable.Name,
		"node_name":     node_name,
		"id":            row.ID,
		"properties":    properties,
	}
}

func createPeerFields(p ClusterPeer) mapstr.M {
	initialAllowedSVMs, err := netapp.ToJSONString(p.InitialAllowedSVMs)
	if err != nil {
		initialAllowedSVMs = ""
	}
	peerApplications, err := netapp.ToJSONString(p.PeerApplications)
	if err != nil {
		peerApplications = ""
	}
	ipAddresses, err := netapp.ToJSONString(p.Remote.IPAddresses)
	if err != nil {
		ipAddresses = ""
	}

	return mapstr.M{
		"authentication": mapstr.M{
			"expiry_time": p.Authentication.ExpiryTime,
			"in_use":      p.Authentication.InUse,
			"passphrase":  p.Authentication.Passphrase,
			"state":       p.Authentication.State,
		},
		"encryption": mapstr.M{
			"proposed": p.Encryption.Proposed,
			"state":    p.Encryption.State,
		},
		"initial_allowed_svms": initialAllowedSVMs,
		"ip_address":           p.IPAddress,
		"ipspace":              netapp.CreateNamedObjectFields(p.Ipspace),
		"name":                 p.Name,
		"peer_applications":    peerApplications,
		"remote": mapstr.M{
			"ip_addresses":  ipAddresses,
			"name":          p.Remote.Name,
			"serial_number": p.Remote.SerialNumber,
		},
		"status": mapstr.M{
			"state":       p.Status.State,
			"update_time": p.Status.UpdateTime,
		},
		"uuid": p.UUID,
		"version": mapstr.M{
			"full":       p.Version.Full,
			"generation": p.Version.Generation,
			"major":      p.Version.Major,
			"minor":      p.Version.Minor,
		},
	}
}
