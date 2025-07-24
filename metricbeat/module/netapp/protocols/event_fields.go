package protocols

import (
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// endpoint: /api/protocols/san/iscsi/services
func createISCSIServiceFields(service ISCSIService) mapstr.M {
	return mapstr.M{
		"iscsi_service": mapstr.M{
			"svm":        netapp.CreateNamedObjectFields(service.SVM),
			"enabled":    service.Enabled,
			"target":     createTargetInfoFields(service.Target),
			"metric":     service.Metric,
			"statistics": service.Statistics,
		},
	}
}

func createTargetInfoFields(info TargetInfo) mapstr.M {
	return mapstr.M{
		"name":  info.Name,
		"alias": info.Alias,
	}
}

// endpoint: /api/protocols/san/iscsi/sessions
func createISCSISessionFields(session ISCSISession) mapstr.M {

	igroups, err := netapp.ToJSONString(session.Igroups)
	if err != nil {
		igroups = ""
	}

	return mapstr.M{
		"iscsi_session": mapstr.M{
			"igroups": igroups,
			"initiator": mapstr.M{
				"alias":   session.Initiator.Alias,
				"comment": session.Initiator.Comment,
				"name":    session.Initiator.Name,
			},
			"isid":                    session.ISID,
			"svm":                     netapp.CreateNamedObjectFields(session.SVM),
			"target_portal_group":     session.TargetPortalGroup,
			"target_portal_group_tag": session.TargetPortalGroupTag,
			"tsih":                    session.TSIH,
		},
	}
}

func createISCSIConnectionFields(conn ISCSIConnection) mapstr.M {
	return mapstr.M{
		"authentication_type": conn.AuthenticationType,
		"cid":                 conn.CID,
		"initiator_address": mapstr.M{
			"address": conn.InitiatorAddress.Address,
			"port":    conn.InitiatorAddress.Port,
		},
		"interface": mapstr.M{
			"ip": mapstr.M{
				"address": conn.Interface.IP.Address,
				"port":    conn.Interface.IP.Port,
			},
			"name": conn.Interface.Name,
			"uuid": conn.Interface.UUID,
		},
	}
}

// endpoint: /api/protocols/cifs/services
func createCIFSServicesFields(service CIFSService) mapstr.M {
	winservers, err := netapp.ToJSONString(service.Netbios.WinsServers)
	if err != nil {
		winservers = ""
	}
	return mapstr.M{
		"cifs_service": mapstr.M{
			"ad_domain": mapstr.M{
				"default_site":        service.AdDomain.DefaultSite,
				"fqdn":                service.AdDomain.FQDN,
				"organizational_unit": service.AdDomain.OrganizationalUnit,
			},
			"auth_style":                  service.AuthStyle,
			"auth_user_type":              service.AuthUserType,
			"authentication_method":       service.AuthenticationMethod,
			"client_id":                   service.ClientID,
			"comment":                     service.Comment,
			"default_unix_user":           service.DefaultUnixUser,
			"enabled":                     service.Enabled,
			"group_policy_object_enabled": service.GroupPolicyObjectEnabled,
			"key_vault_uri":               service.KeyVaultURI,
			"name":                        service.Name,
			"netbios": mapstr.M{
				"aliases":      service.Netbios.Aliases,
				"enabled":      service.Netbios.Enabled,
				"wins_servers": winservers,
			},
			"oauth_host":     service.OAuthHost,
			"options":        createCIFSOptionsFields(service.Options),
			"proxy_host":     service.ProxyHost,
			"proxy_port":     service.ProxyPort,
			"proxy_type":     service.ProxyType,
			"proxy_username": service.ProxyUsername,
			"security":       createCIFSSecurityFields(service.Security),
			"svm":            netapp.CreateNamedObjectFields(service.SVM),
			"tenant_id":      service.TenantID,
			"timeout":        service.Timeout,
			"verify_host":    service.VerifyHost,
			"workgroup":      service.Workgroup,
			"metric":         netapp.CreateMetricsFields(service.Metric),
			"statistics":     netapp.CreateStatisticsFields(service.Statistics),
		},
	}
}

func createCIFSSecurityFields(security CIFSSecurity) mapstr.M {
	return mapstr.M{
		"advertised_kdc_encryptions": security.AdvertisedKDCEncryptions,
		"aes_netlogon_enabled":       security.AESNetlogonEnabled,
		"encrypt_dc_connection":      security.EncryptDCConnection,
		"kdc_encryption":             security.KDCEncryption,
		"ldap_referral_enabled":      security.LDAPReferralEnabled,
		"lm_compatibility_level":     security.LMCompatibilityLevel,
		"restrict_anonymous":         security.RestrictAnonymous,
		"session_security":           security.SessionSecurity,
		"smb_encryption":             security.SMBEncryption,
		"smb_signing":                security.SMBSigning,
		"try_ldap_channel_binding":   security.TryLDAPChannelBinding,
		"use_ldaps":                  security.UseLDAPS,
		"use_start_tls":              security.UseStartTLS,
	}
}
func createCIFSOptionsFields(options CIFSOptions) mapstr.M {
	return mapstr.M{
		"admin_to_root_mapping":                 options.AdminToRootMapping,
		"advanced_sparse_file":                  options.AdvancedSparseFile,
		"backup_symlink_enabled":                options.BackupSymlinkEnabled,
		"client_dup_detection_enabled":          options.ClientDupDetectionEnabled,
		"client_version_reporting_enabled":      options.ClientVersionReportingEnabled,
		"copy_offload":                          options.CopyOffload,
		"dac_enabled":                           options.DacEnabled,
		"export_policy_enabled":                 options.ExportPolicyEnabled,
		"fake_open":                             options.FakeOpen,
		"fsctl_trim":                            options.FsctlTrim,
		"junction_reparse":                      options.JunctionReparse,
		"large_mtu":                             options.LargeMTU,
		"max_connections_per_session":           options.MaxConnectionsPerSession,
		"max_lifs_per_session":                  options.MaxLifsPerSession,
		"max_opens_same_file_per_tree":          options.MaxOpensSameFilePerTree,
		"max_same_tree_connect_per_session":     options.MaxSameTreeConnectPerSession,
		"max_same_user_sessions_per_connection": options.MaxSameUserSessionsPerConnection,
		"max_watches_set_per_tree":              options.MaxWatchesSetPerTree,
		"multichannel":                          options.Multichannel,
		"null_user_windows_name":                options.NullUserWindowsName,
		"path_component_cache":                  options.PathComponentCache,
		"referral":                              options.Referral,
		"shadowcopy":                            options.Shadowcopy,
		"shadowcopy_dir_depth":                  options.ShadowcopyDirDepth,
		"smb_credits":                           options.SmbCredits,
		"trusted_domain_enum_search_enabled":    options.TrustedDomainEnumSearchEnabled,
		"widelink_reparse_versions":             options.WidelinkReparseVersions,
	}
}

// endpoint: /api/protocols/cifs/shares
func createCIFSShareFields(share CIFSShare) mapstr.M {
	acls, err := netapp.ToJSONString(share.Acls)
	if err != nil {
		acls = ""
	}

	return mapstr.M{
		"cifs_share": mapstr.M{
			"access_based_enumeration":  share.AccessBasedEnumeration,
			"acls":                      acls,
			"allow_unencrypted_access":  share.AllowUnencryptedAccess,
			"attribute_cache":           share.AttributeCache,
			"browsable":                 share.Browsable,
			"change_notify":             share.ChangeNotify,
			"comment":                   share.Comment,
			"continuously_available":    share.ContinuouslyAvailable,
			"dir_umask":                 share.DirUmask,
			"encryption":                share.Encryption,
			"file_umask":                share.FileUmask,
			"force_group_for_create":    share.ForceGroupForCreate,
			"home_directory":            share.HomeDirectory,
			"max_connections_per_share": share.MaxConnectionsPerShare,
			"name":                      share.Name,
			"namespace_caching":         share.NamespaceCaching,
			"no_strict_security":        share.NoStrictSecurity,
			"offline_files":             share.OfflineFiles,
			"oplocks":                   share.Oplocks,
			"path":                      share.Path,
			"show_previous_versions":    share.ShowPreviousVersions,
			"show_snapshot":             share.ShowSnapshot,
			"svm":                       netapp.CreateNamedObjectFields(share.SVM),
			"unix_symlink":              share.UnixSymlink,
			"volume":                    netapp.CreateNamedObjectFields(share.Volume),
			"vscan_profile":             share.VscanProfile,
		},
	}
}

// endpoint: /api/protocols/san/igroups

func createIGroupFields(igroup IGroup) mapstr.M {

	lunMaps, err := netapp.ToJSONString(igroup.LunMaps)
	if err != nil {
		lunMaps = ""
	}

	portset, err := netapp.ToJSONString(igroup.Portset)
	if err != nil {
		portset = ""
	}

	replication := createIGroupReplicationFields(igroup.Replication)

	var target mapstr.M
	if igroup.Target != nil {
		target = mapstr.M{
			"firmware_revision": igroup.Target.FirmwareRevision,
			"product_id":        igroup.Target.ProductID,
			"vendor_id":         igroup.Target.VendorID,
		}
	}

	return mapstr.M{
		"igroup": mapstr.M{
			"comment":               igroup.Comment,
			"connectivity_tracking": createConnectivityTrackingFields(igroup.ConnectivityTracking),
			"delete_on_unmap":       igroup.DeleteOnUnmap,
			"lun_maps":              lunMaps,
			"name":                  igroup.Name,
			"os_type":               igroup.OsType,
			"portset":               portset,
			"protocol":              igroup.Protocol,
			"replication":           replication,
			"supports_igroups":      igroup.SupportsIgroups,
			"svm":                   netapp.CreateNamedObjectFields(igroup.SVM),
			"target":                target,
			"uuid":                  igroup.UUID,
		},
	}
}

func createIGroupReplicationFields(replication *IGroupReplication) mapstr.M {
	if replication == nil {
		return nil
	}

	var err mapstr.M
	if replication.Error != nil {
		err = mapstr.M{
			"igroup": mapstr.M{
				"local_svm": replication.Error.Igroup.LocalSVM,
				"name":      replication.Error.Igroup.Name,
				"uuid":      replication.Error.Igroup.UUID,
			},
			"summary": mapstr.M{
				"code":    replication.Error.Summary.Code,
				"message": replication.Error.Summary.Message,
			},
		}
	}

	var peerSvm mapstr.M
	if replication.PeerSVM != nil {
		peerSvm = netapp.CreateNamedObjectFields(*replication.PeerSVM)
	}

	return mapstr.M{
		"error":    err,
		"peer_svm": peerSvm,
		"state":    replication.State,
	}
}

func createInitiatorFields(initiator IGroupInitiator) mapstr.M {

	return mapstr.M{
		"comment":          initiator.Comment,
		"connection_state": initiator.ConnectivityTracking.ConnectionState,
		"igroup": mapstr.M{
			"comment": initiator.Igroup.Comment,
			"igroups": nil, // avoid deep recursion
			"name":    initiator.Igroup.Name,
			"uuid":    initiator.Igroup.UUID,
		},
		"name":      initiator.Name,
		"proximity": createProximityFields(initiator.Proximity),
	}
}

func createProximityFields(proximity IGroupInitiatorProximity) mapstr.M {
	peerSvms, err := netapp.ToJSONString(proximity.PeerSVMs)
	if err != nil {
		peerSvms = ""
	}
	return mapstr.M{
		"local_svm": proximity.LocalSVM,
		"peer_svms": peerSvms,
	}
}

func createConnectivityTrackingFields(ct IGroupConnectivity) mapstr.M {
	alerts, err := netapp.ToJSONString(ct.Alerts)
	if err != nil {
		alerts = ""
	}

	required, err := netapp.ToJSONString(ct.RequiredNodes)
	if err != nil {
		required = ""
	}

	return mapstr.M{
		"alerts":           alerts,
		"connection_state": ct.ConnectionState,
		"required_nodes":   required,
	}
}

// endpoint: /api/network/fc/interfaces
func createFCInterfaceFields(iface FCInterface) mapstr.M {
	return mapstr.M{
		"fc_interface": mapstr.M{
			"comment":       iface.Comment,
			"data_protocol": iface.DataProtocol,
			"enabled":       iface.Enabled,
			"location": mapstr.M{
				"home_node": netapp.CreateNamedObjectFields(iface.Location.HomeNode),
				"home_port": createLocationPortFields(iface.Location.HomePort),
				"is_home":   iface.Location.IsHome,
				"node":      netapp.CreateNamedObjectFields(iface.Location.Node),
				"port":      createLocationPortFields(iface.Location.Port),
			},
			"metric":       netapp.CreateMetricsFields(iface.Metric),
			"name":         iface.Name,
			"port_address": iface.PortAddress,
			"state":        iface.State,
			"statistics":   netapp.CreateStatisticsFields(iface.Statistics),
			"svm":          netapp.CreateNamedObjectFields(iface.SVM),
			"uuid":         iface.UUID,
			"wwnn":         iface.WWNN,
			"wwpn":         iface.WWPN,
		},
	}
}

func createLocationPortFields(port HomePort) mapstr.M {
	return mapstr.M{
		"name": port.Name,
		"node": mapstr.M{
			"name": port.Node.Name,
		},
		"uuid": port.UUID,
	}
}

// endpoint: /api/network/fc/ports

func createFCPortFields(port FCPort) mapstr.M {

	var transceiver mapstr.M
	if port.Transceiver != nil {
		transceiver = mapstr.M{
			"form_factor":  port.Transceiver.FormFactor,
			"manufacturer": port.Transceiver.Manufacturer,
			"part_number":  port.Transceiver.PartNumber,
		}
		if port.Transceiver.Capabilities != nil {
			transceiver["capabilities"] = netapp.IntArrayToString(port.Transceiver.Capabilities)
		}
	}

	supportedProtocols, err := netapp.ToJSONString(port.SupportedProtocols)
	if err != nil {
		supportedProtocols = ""
	}
	return mapstr.M{
		"fc_port": mapstr.M{
			"node":        netapp.CreateNamedObjectFields(port.Node),
			"name":        port.Name,
			"uuid":        port.UUID,
			"description": port.Description,
			"enabled":     port.Enabled,
			"fabric": mapstr.M{
				"connected":       port.Fabric.Connected,
				"connected_speed": port.Fabric.ConnectedSpeed,
				"port_address":    port.Fabric.PortAddress,
				"switch_port":     port.Fabric.SwitchPort,
			},
			"physical_protocol": port.PhysicalProtocol,
			"speed": mapstr.M{
				"maximum":    port.Speed.Maximum,
				"configured": port.Speed.Configured,
			},
			"state":               port.State,
			"supported_protocols": supportedProtocols,
			"transceiver":         transceiver,
			"wwnn":                port.WWNN,
			"wwpn":                port.WWPN,
			"metric":              netapp.CreateMetricsFields(port.Metric),
			"statistics":          netapp.CreateStatisticsFields(port.Statistics),
		},
	}
}

// endpoint: /api/protocols/san/fcp/services

func createFCPServiceFields(service FCPService) mapstr.M {
	return mapstr.M{
		"fcp_service": mapstr.M{
			"enabled":     service.Enabled,
			"metric":      netapp.CreateMetricsFields(service.Metric),
			"statistics":  netapp.CreateStatisticsFields(service.Statistics),
			"svm":         netapp.CreateNamedObjectFields(service.SVM),
			"target_name": service.Target.Name,
		},
	}
}

// endpoint: /api/protocols/nfs/services

func createNFSServiceFields(service NFSService) mapstr.M {
	return mapstr.M{
		"nfs_service": mapstr.M{
			"svm":     netapp.CreateNamedObjectFields(service.SVM),
			"enabled": service.Enabled,
			"state":   service.State,
			"transport": mapstr.M{
				"udp_enabled":  service.Transport.UDPEnabled,
				"tcp_enabled":  service.Transport.TCPEnabled,
				"rdma_enabled": service.Transport.RDMAEnabled,
			},
			"protocol": mapstr.M{
				"v3_enabled":                   service.Protocol.V3Enabled,
				"v3_64bit_identifiers_enabled": service.Protocol.V364bitIdentifiersEnabled,
				"v4_id_domain":                 service.Protocol.V4IDDomain,
				"v4_64bit_identifiers_enabled": service.Protocol.V464bitIdentifiersEnabled,
				"v40_enabled":                  service.Protocol.V40Enabled,
				"v41_enabled":                  service.Protocol.V41Enabled,
				"v4_grace_seconds":             service.Protocol.V4GraceSeconds,
				"v40_features": mapstr.M{
					"acl_enabled":              service.Protocol.V40Features.ACLEnabled,
					"read_delegation_enabled":  service.Protocol.V40Features.ReadDelegationEnabled,
					"write_delegation_enabled": service.Protocol.V40Features.WriteDelegationEnabled,
					"acl_preserve":             service.Protocol.V40Features.ACLPreserve,
				},
				"v41_features": mapstr.M{
					"acl_enabled":              service.Protocol.V41Features.ACLEnabled,
					"read_delegation_enabled":  service.Protocol.V41Features.ReadDelegationEnabled,
					"write_delegation_enabled": service.Protocol.V41Features.WriteDelegationEnabled,
					"pnfs_enabled":             service.Protocol.V41Features.PnfsEnabled,
				},
				"v3_features": mapstr.M{
					"mount_root_only": service.Protocol.V3Features.MountRootOnly,
				},
			},
			"vstorage_enabled":                 service.VstorageEnabled,
			"rquota_enabled":                   service.RquotaEnabled,
			"showmount_enabled":                service.ShowmountEnabled,
			"auth_sys_extended_groups_enabled": service.AuthSysExtendedGroupsEnabled,
			"extended_groups_limit":            service.ExtendedGroupsLimit,
			"credential_cache": mapstr.M{
				"positive_ttl": service.CredentialCache.PositiveTTL,
			},
			"qtree": mapstr.M{
				"export_enabled":  service.Qtree.ExportEnabled,
				"validate_export": service.Qtree.ValidateExport,
			},
			"access_cache_config": mapstr.M{
				"ttl_positive":       service.AccessCacheConfig.TTLPositive,
				"ttl_negative":       service.AccessCacheConfig.TTLNegative,
				"harvest_timeout":    service.AccessCacheConfig.HarvestTimeout,
				"is_dns_ttl_enabled": service.AccessCacheConfig.IsDnsTTLEnabled,
			},
			"file_session_io_grouping_count":    service.FileSessionIOGroupingCount,
			"file_session_io_grouping_duration": service.FileSessionIOGroupingDuration,
			"exports": mapstr.M{
				"name_service_lookup_protocol": service.Exports.NameServiceLookupProtocol,
			},
			"security": mapstr.M{
				"permitted_encryption_types": service.Security.PermittedEncryptionTypes,
			},
			"windows": mapstr.M{
				"v3_ms_dos_client_enabled": service.Windows.V3MsDosClientEnabled,
			},
			"metric": mapstr.M{
				"v3":  netapp.CreateMetricsFields(service.Metric.V3),
				"v4":  netapp.CreateMetricsFields(service.Metric.V4),
				"v41": netapp.CreateMetricsFields(service.Metric.V41),
			},
			"statistics": mapstr.M{
				"v3":  netapp.CreateStatisticsFields(service.Statistics.V3),
				"v4":  netapp.CreateStatisticsFields(service.Statistics.V4),
				"v41": netapp.CreateStatisticsFields(service.Statistics.V41),
			},
		},
	}
}

// endpoint: /api/protocols/nfs/export-policies
func createNFSExportPolicyFields(policy NFSExportPolicy) mapstr.M {
	return mapstr.M{
		"svm":  netapp.CreateNamedObjectFields(policy.SVM),
		"id":   policy.ID,
		"name": policy.Name,
	}
}

// endpoint: /api/network/ip/interfaces

func createIPInterfaceFields(iface IPInterface) mapstr.M {
	rdma_protocols, err := netapp.ToJSONString(iface.RDMAProtocols)
	if err != nil {
		rdma_protocols = ""
	}
	services, err := netapp.ToJSONString(iface.Services)
	if err != nil {
		services = ""
	}
	return mapstr.M{
		"ip_interface": mapstr.M{
			"ddns_enabled": iface.DDNSEnabled,
			"dns_zone":     iface.DNSZone,
			"enabled":      iface.Enabled,
			"ip": mapstr.M{
				"address": iface.IP.Address,
				"family":  iface.IP.Family,
				"netmask": iface.IP.Netmask,
			},
			"ipspace": netapp.CreateNamedObjectFields(iface.IPSpace),
			"location": mapstr.M{
				"auto_revert": iface.Location.AutoRevert,
				"failover":    iface.Location.Failover,
				"home_node":   netapp.CreateNamedObjectFields(iface.Location.HomeNode),
				"home_port":   createIPLocationPortFields(iface.Location.HomePort),
				"is_home":     iface.Location.IsHome,
				"node":        netapp.CreateNamedObjectFields(iface.Location.Node),
				"port":        createIPLocationPortFields(iface.Location.Port),
			},
			"metric":         netapp.CreateMetricsFields(iface.Metric),
			"name":           iface.Name,
			"probe_port":     iface.ProbePort,
			"rdma_protocols": rdma_protocols,
			"scope":          iface.Scope,
			"service_policy": netapp.CreateNamedObjectFields(iface.ServicePolicy),
			"services":       services,
			"state":          iface.State,
			"statistics":     netapp.CreateStatisticsFields(iface.Statistics),
			"subnet":         netapp.CreateNamedObjectFields(iface.Subnet),
			"svm":            netapp.CreateNamedObjectFields(iface.SVM),
			"uuid":           iface.UUID,
			"vip":            iface.VIP,
		},
	}
}

func createIPLocationPortFields(port IPLocationPort) mapstr.M {
	var node mapstr.M
	if port.Node != nil {
		node = netapp.CreateNamedObjectFields(*port.Node)
	}
	return mapstr.M{
		"name": port.Name,
		"node": node,
		"uuid": port.UUID,
	}
}
