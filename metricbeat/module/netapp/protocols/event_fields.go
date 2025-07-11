package protocols

import (
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// endpoint: /api/protocols/san/iscsi/services
func createISCSIServiceFields(service ISCSIService) mapstr.M {
	return mapstr.M{
		"svm":        netapp.CreateNamedObjectFields(service.SVM),
		"enabled":    service.Enabled,
		"target":     createTargetInfoFields(service.Target),
		"metric":     service.Metric,
		"statistics": service.Statistics,
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
	connections := make([]mapstr.M, len(session.Connections))
	for i, conn := range session.Connections {
		connections[i] = mapstr.M{
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

	igroups := make([]mapstr.M, len(session.Igroups))
	for i, ig := range session.Igroups {
		igroups[i] = netapp.CreateNamedObjectFields(ig)
	}

	initiator := mapstr.M{
		"alias":   session.Initiator.Alias,
		"comment": session.Initiator.Comment,
		"name":    session.Initiator.Name,
	}

	return mapstr.M{
		"connections":             connections,
		"igroups":                 igroups,
		"initiator":               initiator,
		"isid":                    session.ISID,
		"svm":                     netapp.CreateNamedObjectFields(session.SVM),
		"target_portal_group":     session.TargetPortalGroup,
		"target_portal_group_tag": session.TargetPortalGroupTag,
		"tsih":                    session.TSIH,
	}
}

// endpoint: /api/protocols/cifs/services
func createCIFSServicesFields(service CIFSService) mapstr.M {
	return mapstr.M{
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
		"metric":                      service.Metric,
		"name":                        service.Name,
		"netbios": mapstr.M{
			"aliases":      service.Netbios.Aliases,
			"enabled":      service.Netbios.Enabled,
			"wins_servers": service.Netbios.WinsServers,
		},
		"oauth_host": service.OAuthHost,
		"options": mapstr.M{
			"admin_to_root_mapping":                 service.Options.AdminToRootMapping,
			"advanced_sparse_file":                  service.Options.AdvancedSparseFile,
			"backup_symlink_enabled":                service.Options.BackupSymlinkEnabled,
			"client_dup_detection_enabled":          service.Options.ClientDupDetectionEnabled,
			"client_version_reporting_enabled":      service.Options.ClientVersionReportingEnabled,
			"copy_offload":                          service.Options.CopyOffload,
			"dac_enabled":                           service.Options.DacEnabled,
			"export_policy_enabled":                 service.Options.ExportPolicyEnabled,
			"fake_open":                             service.Options.FakeOpen,
			"fsctl_trim":                            service.Options.FsctlTrim,
			"junction_reparse":                      service.Options.JunctionReparse,
			"large_mtu":                             service.Options.LargeMTU,
			"max_connections_per_session":           service.Options.MaxConnectionsPerSession,
			"max_lifs_per_session":                  service.Options.MaxLifsPerSession,
			"max_opens_same_file_per_tree":          service.Options.MaxOpensSameFilePerTree,
			"max_same_tree_connect_per_session":     service.Options.MaxSameTreeConnectPerSession,
			"max_same_user_sessions_per_connection": service.Options.MaxSameUserSessionsPerConnection,
			"max_watches_set_per_tree":              service.Options.MaxWatchesSetPerTree,
			"multichannel":                          service.Options.Multichannel,
			"null_user_windows_name":                service.Options.NullUserWindowsName,
			"path_component_cache":                  service.Options.PathComponentCache,
			"referral":                              service.Options.Referral,
			"shadowcopy":                            service.Options.Shadowcopy,
			"shadowcopy_dir_depth":                  service.Options.ShadowcopyDirDepth,
			"smb_credits":                           service.Options.SmbCredits,
			"trusted_domain_enum_search_enabled":    service.Options.TrustedDomainEnumSearchEnabled,
			"widelink_reparse_versions":             service.Options.WidelinkReparseVersions,
		},
		"proxy_host":     service.ProxyHost,
		"proxy_port":     service.ProxyPort,
		"proxy_type":     service.ProxyType,
		"proxy_username": service.ProxyUsername,
		"security": mapstr.M{
			"advertised_kdc_encryptions": service.Security.AdvertisedKDCEncryptions,
			"aes_netlogon_enabled":       service.Security.AESNetlogonEnabled,
			"encrypt_dc_connection":      service.Security.EncryptDCConnection,
			"kdc_encryption":             service.Security.KDCEncryption,
			"ldap_referral_enabled":      service.Security.LDAPReferralEnabled,
			"lm_compatibility_level":     service.Security.LMCompatibilityLevel,
			"restrict_anonymous":         service.Security.RestrictAnonymous,
			"session_security":           service.Security.SessionSecurity,
			"smb_encryption":             service.Security.SMBEncryption,
			"smb_signing":                service.Security.SMBSigning,
			"try_ldap_channel_binding":   service.Security.TryLDAPChannelBinding,
			"use_ldaps":                  service.Security.UseLDAPS,
			"use_start_tls":              service.Security.UseStartTLS,
		},
		"statistics":  service.Statistics,
		"svm":         netapp.CreateNamedObjectFields(service.SVM),
		"tenant_id":   service.TenantID,
		"timeout":     service.Timeout,
		"verify_host": service.VerifyHost,
		"workgroup":   service.Workgroup,
	}
}

// endpoint: /api/protocols/cifs/shares
func createCIFSShareFields(share CIFSShare) mapstr.M {
	acls := make([]mapstr.M, len(share.Acls))
	for i, acl := range share.Acls {
		acls[i] = mapstr.M{
			"permission":      acl.Permission,
			"type":            acl.Type,
			"user_or_group":   acl.UserOrGroup,
			"win_sid_unix_id": acl.WinSidUnixID,
		}
	}

	return mapstr.M{
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
	}
}

// endpoint: /api/protocols/san/igroups

func createIGroupFields(igroup IGroup) mapstr.M {
	var connectivityTracking mapstr.M
	if igroup.ConnectivityTracking != nil {
		alerts := make([]mapstr.M, len(igroup.ConnectivityTracking.Alerts))
		for i, alert := range igroup.ConnectivityTracking.Alerts {
			args := make([]mapstr.M, len(alert.Summary.Arguments))
			for j, arg := range alert.Summary.Arguments {
				args[j] = mapstr.M{
					"code":    arg.Code,
					"message": arg.Message,
				}
			}
			alerts[i] = mapstr.M{
				"summary": mapstr.M{
					"arguments": args,
					"code":      alert.Summary.Code,
					"message":   alert.Summary.Message,
				},
			}
		}
		requiredNodes := make([]mapstr.M, len(igroup.ConnectivityTracking.RequiredNodes))
		for i, node := range igroup.ConnectivityTracking.RequiredNodes {
			requiredNodes[i] = netapp.CreateNamedObjectFields(node)
		}
		connectivityTracking = mapstr.M{
			"alerts":           alerts,
			"connection_state": igroup.ConnectivityTracking.ConnectionState,
			"required_nodes":   requiredNodes,
		}
	}

	igroups := make([]mapstr.M, len(igroup.Igroups))
	for i, ig := range igroup.Igroups {
		nested := make([]mapstr.M, len(ig.Igroups))
		for j, n := range ig.Igroups {
			nested[j] = mapstr.M{
				"comment": n.Comment,
				"igroups": nil, // avoid deep recursion
				"name":    n.Name,
				"uuid":    n.UUID,
			}
		}
		igroups[i] = mapstr.M{
			"comment": ig.Comment,
			"igroups": nested,
			"name":    ig.Name,
			"uuid":    ig.UUID,
		}
	}

	initiators := make([]mapstr.M, len(igroup.Initiators))
	for i, ini := range igroup.Initiators {
		var conn mapstr.M
		if ini.ConnectivityTracking != nil {
			conn = mapstr.M{
				"connection_state": ini.ConnectivityTracking.ConnectionState,
			}
		}
		var ig mapstr.M
		if ini.Igroup != nil {
			ig = mapstr.M{
				"comment": ini.Igroup.Comment,
				"igroups": nil, // avoid deep recursion
				"name":    ini.Igroup.Name,
				"uuid":    ini.Igroup.UUID,
			}
		}
		var prox mapstr.M
		if ini.Proximity != nil {
			peerSvms := make([]mapstr.M, len(ini.Proximity.PeerSVMs))
			for j, svm := range ini.Proximity.PeerSVMs {
				peerSvms[j] = netapp.CreateNamedObjectFields(svm)
			}
			prox = mapstr.M{
				"local_svm": ini.Proximity.LocalSVM,
				"peer_svms": peerSvms,
			}
		}
		initiators[i] = mapstr.M{
			"comment":               ini.Comment,
			"connectivity_tracking": conn,
			"igroup":                ig,
			"name":                  ini.Name,
			"proximity":             prox,
		}
	}

	lunMaps := make([]mapstr.M, len(igroup.LunMaps))
	for i, lm := range igroup.LunMaps {
		var node mapstr.M
		if lm.Lun.Node != nil {
			node = netapp.CreateNamedObjectFields(*lm.Lun.Node)
		}
		lunMaps[i] = mapstr.M{
			"logical_unit_number": lm.LogicalUnitNumber,
			"lun": mapstr.M{
				"name": lm.Lun.Name,
				"node": node,
				"uuid": lm.Lun.UUID,
			},
		}
	}

	parentIgroups := make([]mapstr.M, len(igroup.ParentIgroups))
	for i, p := range igroup.ParentIgroups {
		nested := make([]mapstr.M, len(p.ParentIgroups))
		for j, n := range p.ParentIgroups {
			nested[j] = mapstr.M{
				"comment":        n.Comment,
				"name":           n.Name,
				"parent_igroups": nil, // avoid deep recursion
				"uuid":           n.UUID,
			}
		}
		parentIgroups[i] = mapstr.M{
			"comment":        p.Comment,
			"name":           p.Name,
			"parent_igroups": nested,
			"uuid":           p.UUID,
		}
	}

	var portset mapstr.M
	if igroup.Portset != nil {
		portset = netapp.CreateNamedObjectFields(*igroup.Portset)
	}

	var replication mapstr.M
	if igroup.Replication != nil {
		var err mapstr.M
		if igroup.Replication.Error != nil {
			err = mapstr.M{
				"igroup": mapstr.M{
					"local_svm": igroup.Replication.Error.Igroup.LocalSVM,
					"name":      igroup.Replication.Error.Igroup.Name,
					"uuid":      igroup.Replication.Error.Igroup.UUID,
				},
				"summary": mapstr.M{
					"arguments": func() []mapstr.M {
						args := make([]mapstr.M, len(igroup.Replication.Error.Summary.Arguments))
						for i, arg := range igroup.Replication.Error.Summary.Arguments {
							args[i] = mapstr.M{
								"code":    arg.Code,
								"message": arg.Message,
							}
						}
						return args
					}(),
					"code":    igroup.Replication.Error.Summary.Code,
					"message": igroup.Replication.Error.Summary.Message,
				},
			}
		}
		var peerSvm mapstr.M
		if igroup.Replication.PeerSVM != nil {
			peerSvm = netapp.CreateNamedObjectFields(*igroup.Replication.PeerSVM)
		}
		replication = mapstr.M{
			"error":    err,
			"peer_svm": peerSvm,
			"state":    igroup.Replication.State,
		}
	}

	var target mapstr.M
	if igroup.Target != nil {
		target = mapstr.M{
			"firmware_revision": igroup.Target.FirmwareRevision,
			"product_id":        igroup.Target.ProductID,
			"vendor_id":         igroup.Target.VendorID,
		}
	}

	return mapstr.M{
		"comment":               igroup.Comment,
		"connectivity_tracking": connectivityTracking,
		"delete_on_unmap":       igroup.DeleteOnUnmap,
		"igroups":               igroups,
		"initiators":            initiators,
		"lun_maps":              lunMaps,
		"name":                  igroup.Name,
		"os_type":               igroup.OsType,
		"parent_igroups":        parentIgroups,
		"portset":               portset,
		"protocol":              igroup.Protocol,
		"replication":           replication,
		"supports_igroups":      igroup.SupportsIgroups,
		"svm":                   netapp.CreateNamedObjectFields(igroup.SVM),
		"target":                target,
		"uuid":                  igroup.UUID,
	}
}

// endpoint: /api/network/fc/interfaces
func createFCInterfaceFields(iface FCInterface) mapstr.M {
	return mapstr.M{
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
		"metric":       iface.Metric,
		"name":         iface.Name,
		"port_address": iface.PortAddress,
		"state":        iface.State,
		"statistics":   iface.Statistics,
		"svm":          netapp.CreateNamedObjectFields(iface.SVM),
		"uuid":         iface.UUID,
		"wwnn":         iface.WWNN,
		"wwpn":         iface.WWPN,
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
			"capabilities": port.Transceiver.Capabilities,
			"part_number":  port.Transceiver.PartNumber,
		}
	}

	return mapstr.M{
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
		"supported_protocols": port.SupportedProtocols,
		"transceiver":         transceiver,
		"wwnn":                port.WWNN,
		"wwpn":                port.WWPN,
		"metric":              port.Metric,
		"statistics":          port.Statistics,
	}
}

// endpoint: /api/protocols/san/fcp/services

func createFCPServiceFields(service FCPService) mapstr.M {
	return mapstr.M{
		"enabled":    service.Enabled,
		"metric":     service.Metric,
		"statistics": service.Statistics,
		"svm":        netapp.CreateNamedObjectFields(service.SVM),
		"target": mapstr.M{
			"name": service.Target.Name,
		},
	}
}

// endpoint: /api/protocols/nfs/services

func createNFSServiceFields(service NFSService) mapstr.M {
	return mapstr.M{
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
	return mapstr.M{
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
			"home_port": mapstr.M{
				"name": iface.Location.HomePort.Name,
				"node": func() mapstr.M {
					if iface.Location.HomePort.Node != nil {
						return netapp.CreateNamedObjectFields(*iface.Location.HomePort.Node)
					}
					return nil
				}(),
				"uuid": iface.Location.HomePort.UUID,
			},
			"is_home": iface.Location.IsHome,
			"node":    netapp.CreateNamedObjectFields(iface.Location.Node),
			"port": mapstr.M{
				"name": iface.Location.Port.Name,
				"node": func() mapstr.M {
					if iface.Location.Port.Node != nil {
						return netapp.CreateNamedObjectFields(*iface.Location.Port.Node)
					}
					return nil
				}(),
				"uuid": iface.Location.Port.UUID,
			},
		},
		"metric":         netapp.CreateMetricsFields(iface.Metric),
		"name":           iface.Name,
		"probe_port":     iface.ProbePort,
		"rdma_protocols": iface.RDMAProtocols,
		"scope":          iface.Scope,
		"service_policy": netapp.CreateNamedObjectFields(iface.ServicePolicy),
		"services":       iface.Services,
		"state":          iface.State,
		"statistics":     netapp.CreateStatisticsFields(iface.Statistics),
		"subnet":         netapp.CreateNamedObjectFields(iface.Subnet),
		"svm":            netapp.CreateNamedObjectFields(iface.SVM),
		"uuid":           iface.UUID,
		"vip":            iface.VIP,
	}
}
