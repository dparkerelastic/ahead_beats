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

package purestorage

import (
	"github.com/elastic/beats/v7/metricbeat/mb"
)

const (
	ModuleName = "purestorage"
)

type Config struct {
	HostIp        string `config:"host_ip" validate:"required"`
	ApiKey        string `config:"api_key" validate:"required"`
	ApiVersion    string `config:"api_version" validate:"required"`
	Port          uint   `config:"port"`
	DebugMode     string `config:"api_debug_mode"`
	SnmpBaseOID   string `config:"snmp_base_oid"`
	SnmpCommunity string `config:"snmp_community"`
	SnmpPort      uint16 `config:"snmp_port"`
}

func NewConfig(base mb.BaseMetricSet) (*Config, error) {
	config := Config{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil

}
