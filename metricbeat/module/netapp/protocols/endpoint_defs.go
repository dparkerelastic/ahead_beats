package protocols

import (
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
)

var basicEndpoints map[string]netapp.Endpoint
var customEndpoints map[string]netapp.Endpoint

// endpoint: /api/protocols/san/iscsi/services
// endpoint: /api/protocols/san/iscsi/sessions
// endpoint: /api/protocols/cifs/services
// endpoint: /api/protocols/cifs/shares
// endpoint: /api/protocols/san/igroups
// endpoint: /api/network/fc/interfaces
// endpoint: /api/network/fc/ports
// endpoint: /api/protocols/san/fcp/services
// endpoint: /api/protocols/nfs/services
// endpoint: /api/protocols/nfs/export-policies
// endpoint: /api/network/ip/interfaces

const (
	ISCSIServices     = "IscsiServices"
	ISCSISessions     = "IscsiSessions"
	CIFSServices      = "CifsServices"
	CIFSShares        = "CifsShares"
	SANIgroups        = "SanIgroups"
	FCInterfaces      = "FcInterfaces"
	FCPorts           = "FcPorts"
	FCPServices       = "FcpServices"
	NFSServices       = "NfsServices"
	NFSExportPolicies = "NfsExportPolicies"
	IPInterfaces      = "IpInterfaces"
)

func init() {

	basicEndpoints = map[string]netapp.Endpoint{
		ISCSIServices: {Name: ISCSIServices,
			Endpoint:    "/api/protocols/san/iscsi/services",
			QueryFields: ISCSIServiceFields,
		},
		ISCSISessions: {Name: ISCSISessions,
			Endpoint:    "/api/protocols/san/iscsi/sessions",
			QueryFields: ISCSISessionFields,
		},
		CIFSServices: {Name: CIFSServices,
			Endpoint:    "/api/protocols/cifs/services",
			QueryFields: CIFSServiceFields,
		},
		CIFSShares: {Name: CIFSShares,
			Endpoint:    "/api/protocols/cifs/shares",
			QueryFields: CIFSShareFields,
		},
		SANIgroups: {Name: SANIgroups,
			Endpoint:    "/api/protocols/san/igroups",
			QueryFields: IGroupFields,
		},
		FCInterfaces: {Name: FCInterfaces,
			Endpoint:    "/api/network/fc/interfaces",
			QueryFields: FCInterfaceFields,
		},
		FCPorts: {Name: FCPorts,
			Endpoint:    "/api/network/fc/ports",
			QueryFields: FCPortFields,
		},
		FCPServices: {Name: FCPServices,
			Endpoint:    "/api/protocols/san/fcp/services",
			QueryFields: FCPServiceFields,
		},
		NFSServices: {Name: NFSServices,
			Endpoint:    "/api/protocols/nfs/services",
			QueryFields: NFSServiceFields,
		},
		NFSExportPolicies: {Name: NFSExportPolicies,
			Endpoint:    "/api/protocols/nfs/export-policies",
			QueryFields: NFSExportPolicyFields,
		},
		IPInterfaces: {Name: IPInterfaces,
			Endpoint:    "/api/network/ip/interfaces",
			QueryFields: IPInterfaceFields,
		},
	}

}

// // For processing basic endpoints we need a type-specific function to create the fields.
// // ProcessEndpoint is a generic function that can be used for all basic endpoints,
// // but it needs to know how to create the fields for the specific type.

var endpointDispatchers = map[string]netapp.DispatchFunc{
	ISCSIServices:     netapp.MakeDispatchFunc(createISCSIServiceFields),
	ISCSISessions:     netapp.MakeDispatchFunc(createISCSISessionFields),
	CIFSServices:      netapp.MakeDispatchFunc(createCIFSServicesFields),
	CIFSShares:        netapp.MakeDispatchFunc(createCIFSShareFields),
	SANIgroups:        netapp.MakeDispatchFunc(createIGroupFields),
	FCInterfaces:      netapp.MakeDispatchFunc(createFCInterfaceFields),
	FCPorts:           netapp.MakeDispatchFunc(createFCPortFields),
	FCPServices:       netapp.MakeDispatchFunc(createFCPServiceFields),
	NFSServices:       netapp.MakeDispatchFunc(createNFSServiceFields),
	NFSExportPolicies: netapp.MakeDispatchFunc(createNFSExportPolicyFields),
	IPInterfaces:      netapp.MakeDispatchFunc(createIPInterfaceFields),
}

// Fields to be used in endpoint calls - ONTAP API only returns name and uuid unless fields are specified.
// Further, the NetApp doc discourages using fields=* for various reasons, but also metrics and statistics don't
// get retuned unless explicitly requested.

// endpoint: /api/protocols/san/iscsi/services
var ISCSIServiceFields = []string{
	"svm",
	"enabled",
	"target",
	"metric",
	"statistics",
}

// endpoint: /api/protocols/san/iscsi/sessions
var ISCSISessionFields = []string{
	"connections",
	"igroups",
	"initiator",
	"isid",
	"svm",
	"target_portal_group",
	"target_portal_group_tag",
	"tsih",
}

// endpoint: /api/protocols/cifs/services
var CIFSServiceFields = []string{
	"ad_domain",
	"auth-style",
	"auth_user_type",
	"authentication_method",
	"client_id",
	"comment",
	"default_unix_user",
	"enabled",
	"group_policy_object_enabled",
	"key_vault_uri",
	"metric",
	"name",
	"netbios",
	"oauth_host",
	"options",
	"proxy_host",
	"proxy_port",
	"proxy_type",
	"proxy_username",
	"security",
	"statistics",
	"svm",
	"tenant_id",
	"timeout",
	"verify_host",
	"workgroup",
}

// endpoint: /api/protocols/cifs/shares
var CIFSShareFields = []string{
	"access_based_enumeration",
	"acls",
	"allow_unencrypted_access",
	"attribute_cache",
	"browsable",
	"change_notify",
	"comment",
	"continuously_available",
	"dir_umask",
	"encryption",
	"file_umask",
	"force_group_for_create",
	"home_directory",
	"max_connections_per_share",
	"name",
	"namespace_caching",
	"no_strict_security",
	"offline_files",
	"oplocks",
	"path",
	"show_previous_versions",
	"show_snapshot",
	"svm",
	"unix_symlink",
	"volume",
	"vscan_profile",
}

// endpoint: /api/protocols/san/igroups
var IGroupFields = []string{
	"comment",
	"connectivity_tracking",
	"delete_on_unmap",
	"igroups",
	"initiators",
	"lun_maps",
	"name",
	"os_type",
	"parent_igroups",
	"portset",
	"protocol",
	"replication",
	"supports_igroups",
	"svm",
	"target",
	"uuid",
}

// endpoint: /api/network/fc/interfaces
var FCInterfaceFields = []string{
	"comment",
	"data_protocol",
	"enabled",
	"location",
	"metric",
	"name",
	"port_address",
	"state",
	"statistics",
	"svm",
	"uuid",
	"wwnn",
	"wwpn",
}

// endpoint: /api/network/fc/ports
var FCPortFields = []string{
	"node",
	"name",
	"uuid",
	"description",
	"enabled",
	"fabric",
	"physical_protocol",
	"speed",
	"state",
	"supported_protocols",
	"transceiver",
	"wwnn",
	"wwpn",
	"metric",
	"statistics",
}

// endpoint: /api/protocols/san/fcp/services
var FCPServiceFields = []string{
	"enabled",
	"metric",
	"statistics",
	"svm",
	"target",
}

// endpoint: /api/protocols/nfs/services
var NFSServiceFields = []string{
	"svm",
	"enabled",
	"state",
	"transport",
	"protocol",
	"vstorage_enabled",
	"rquota_enabled",
	"showmount_enabled",
	"auth_sys_extended_groups_enabled",
	"extended_groups_limit",
	"credential_cache",
	"qtree",
	"access_cache_config",
	"file_session_io_grouping_count",
	"file_session_io_grouping_duration",
	"exports",
	"security",
	"windows",
	"metric",
	"statistics",
}

// endpoint: /api/protocols/nfs/export-policies
var NFSExportPolicyFields = []string{
	"svm",
	"id",
	"name",
}

// endpoint: /api/network/ip/interfaces
var IPInterfaceFields = []string{
	"ddns_enabled",
	"dns_zone",
	"enabled",
	"ip",
	"ipspace",
	"location",
	"metric",
	"name",
	"probe_port",
	"rdma_protocols",
	"scope",
	"service_policy",
	"services",
	"state",
	"statistics",
	"subnet",
	"svm",
	"uuid",
	"vip",
}
