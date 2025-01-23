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
	"strconv"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/purestorage"
	"github.com/elastic/elastic-agent-libs/logp"
)

const (
	metricsetName = "arrays"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet(purestorage.ModuleName, metricsetName, New)

}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	config   *purestorage.Config
	logger   *logp.Logger
	psClient *PureSnmpClient
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The purestorage arrays metricset is beta.")
	config, err := purestorage.NewConfig(base)
	if err != nil {
		return nil, err
	}

	logger := logp.NewLogger(base.FullyQualifiedName())

	// Get the session cookie
	psClient, err := GetSnmpClient(config, base)

	if err != nil {
		logger.Errorf("Failed to get session cookie: %v", err)
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		config:        config,
		logger:        logger,
		psClient:      psClient,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	timestamp := time.Now().UTC()

	results, err := m.psClient.Get()
	if err != nil {
		errstr := fmt.Sprintf("failed to get SNMP data: %v", err)
		m.logger.Errorf(errstr)
		return fmt.Errorf("%s", errstr)
	}

	for _, result := range results {
		// All of the OID values are supposed to be integers, according to the MIB. If the conversion fails,
		// we log an error, report and event with error.message, and continue to the next OID.
		value, err := strconv.Atoi(result.Value)
		if err != nil {
			errstr := fmt.Sprintf("failed to convert SNMP value to integer: %v for oid %s", err, result.OIDName)
			m.logger.Errorf(errstr)
			errevent := mb.Event{
				Timestamp: timestamp,
				MetricSetFields: map[string]interface{}{
					"snmp.oid_name": result.OIDName,
					"snmp.oid":      result.OID,
				},
				RootFields: purestorage.MakeErrorFields(errstr, m.config.HostIp),
			}
			report.Event(errevent)
			continue
		}

		event := mb.Event{
			Timestamp: timestamp,
			MetricSetFields: map[string]interface{}{
				"snmp.oid_name": result.OIDName,
				"snmp.oid":      result.OID,
				"snmp.value":    value,
			},
			RootFields: purestorage.MakeRootFields(m.config.HostIp),
		}

		report.Event(event)
	}
	return nil
}
