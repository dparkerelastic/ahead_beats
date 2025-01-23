package arrays

import (
	"fmt"
	"log"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/gosnmp/gosnmp"

	"github.com/elastic/beats/v7/metricbeat/module/purestorage"
)

type PureSnmpClient struct {
	config    *purestorage.Config
	target    string
	client    *gosnmp.GoSNMP
	community string
	baseOID   string
}

type SNMPResult struct {
	OIDName string `json:"OIDName"`
	OID     string `json:"OID"`
	Value   string `json:"Value"`
}

var oidToName = map[string]string{
	".1.3.6.1.4.1.40482.4.1.0": "PureArrayReadBandwidth",
	".1.3.6.1.4.1.40482.4.3.0": "PureArrayReadIOPS",
	".1.3.6.1.4.1.40482.4.5.0": "PureArrayReadLatency",
	".1.3.6.1.4.1.40482.4.2.0": "PureArrayWriteBandwidth",
	".1.3.6.1.4.1.40482.4.4.0": "PureArrayWriteIOPS",
	".1.3.6.1.4.1.40482.4.6.0": "PureArrayWriteLatency",
}

func GetSnmpClient(config *purestorage.Config, base mb.BaseMetricSet) (*PureSnmpClient, error) {
	snmp := &gosnmp.GoSNMP{
		Target:    config.HostIp,
		Port:      config.SnmpPort,
		Community: config.SnmpCommunity,
		Version:   gosnmp.Version2c,
		Timeout:   gosnmp.Default.Timeout,
	}

	client := PureSnmpClient{
		config:    config,
		target:    config.HostIp,
		client:    snmp,
		community: config.SnmpCommunity,
		baseOID:   config.SnmpBaseOID,
	}

	return &client, nil
}

func (c *PureSnmpClient) Get() ([]SNMPResult, error) {

	if err := c.client.Connect(); err != nil {

		return nil, fmt.Errorf("error connecting to SNMP target %s: %v", c.target, err)
	}
	defer c.client.Conn.Close()

	var results []SNMPResult
	err := c.client.Walk(c.baseOID, func(variable gosnmp.SnmpPDU) error {
		name := oidToName[variable.Name] // Lookup the human-readable name
		if name == "" {
			name = "Unknown OID"
		}
		value := fmt.Sprintf("%v", variable.Value)
		results = append(results, SNMPResult{
			OIDName: name,
			OID:     variable.Name,
			Value:   value,
		})
		return nil
	})
	if err != nil {
		log.Fatalf("error performing SNMP Walk: %v", err)
		return nil, fmt.Errorf("error performing SNMP Walk: %v", err)
	}

	return results, nil
}
