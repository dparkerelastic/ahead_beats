package cluster

import (
	"github.com/elastic/beats/v7/metricbeat/module/netapp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func getClusterFields(c Cluster) mapstr.M {

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
func createCounterFields(c CounterTableRow) mapstr.M {
	properties, err := netapp.ToJSONString(c.Properties)
	if err != nil {
		properties = ""
	}

	return mapstr.M{
		"counter_table": c.CounterTable.Name,
		"id":            c.ID,
		"properties":    properties,
	}
}
