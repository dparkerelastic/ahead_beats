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

package nsxt

import (
	"errors"
	"net"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"
)

const (
	ModuleName  = "nsxt"
	DefaultPort = 443
)

type Config struct {
	Host      string `config:"host" validate:"required"`
	Username  string `config:"username" validate:"required"`
	Password  string `config:"password" validate:"required"`
	Port      uint   `config:"port" default:"443"`
	Protocol  string `config:"protocol" default:"https"`
	DebugMode string `config:"api_debug_mode"`
	HostInfo  HostInfo
}

func NewConfig(base mb.BaseMetricSet, logger *logp.Logger) (*Config, error) {
	config := Config{}
	var err error

	if err = base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	if config.Port == 0 {
		config.Port = DefaultPort
	}

	if config.Host == "" {
		return nil, errors.New("host is required")
	} else {
		config.HostInfo, err = GetHostInfo(config.Host)
		if err != nil {
			var dnsErr *net.DNSError
			if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
				logger.Warnf("Ignoring: host not found for ip: %s", config.Host)
			} else {
				return nil, err
			}
		}
		return &config, nil
	}

}
