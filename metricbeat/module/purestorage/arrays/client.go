// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package arrays

import (
	"fmt"

	"github.com/gosnmp/gosnmp"

	"github.com/elastic/beats/v7/metricbeat/mb"

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

type OidField struct {
	OID          string
	OIDName      string
	OIDFieldName string
}

var OidToName = map[string]OidField{
	".1.3.6.1.4.1.40482.4.1.0": {".1.3.6.1.4.1.40482.4.1.0", "PureArrayReadBandwidth", "read_bandwidth"},
	".1.3.6.1.4.1.40482.4.3.0": {".1.3.6.1.4.1.40482.4.3.0", "PureArrayReadIOPS", "read_iops"},
	".1.3.6.1.4.1.40482.4.5.0": {".1.3.6.1.4.1.40482.4.5.0", "PureArrayReadLatency", "read_latency"},
	".1.3.6.1.4.1.40482.4.2.0": {".1.3.6.1.4.1.40482.4.2.0", "PureArrayWriteBandwidth", "write_bandwidth"},
	".1.3.6.1.4.1.40482.4.4.0": {".1.3.6.1.4.1.40482.4.4.0", "PureArrayWriteIOPS", "write_iops"},
	".1.3.6.1.4.1.40482.4.6.0": {".1.3.6.1.4.1.40482.4.6.0", "PureArrayWriteLatency", "write_latency"},
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
		name := OidToName[variable.Name].OIDName // Lookup the human-readable name
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
		return nil, fmt.Errorf("error performing SNMP Walk: %v", err)
	}

	return results, nil
}
