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

package health

import (
	"errors"
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/dellecs"
	"github.com/elastic/elastic-agent-libs/logp"
)

const (
	metricsetName = "health"
)

type Endpoint struct {
	Name     string
	Endpoint string
	Fn       func(*MetricSet) ([]mb.Event, error)
}

var endpoints map[string]Endpoint

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	endpoints = map[string]Endpoint{
		"Processes":          {Name: "Processes", Endpoint: "dashboard/nodes/{nodeid}/processes", Fn: getNodeProcessEvents},
		"Node Details":       {Name: "Node Details", Endpoint: "dashboard/nodes/{nodeid}", Fn: getNoteDetailsEvents},
		"Disks":              {Name: "Disks", Endpoint: "dashboard/nodes/{nodeid}/disks", Fn: getDiskEvents},
		"Capacity":           {Name: "Capacity", Endpoint: "object/capacity.json", Fn: getCapacityEvents},
		"Storage Pools":      {Name: "Storage Pools", Endpoint: "dashboard/zones/localzone/storagepools", Fn: getStoragePoolsEvents},
		"Replication Groups": {Name: "Replication Groups", Endpoint: "dashboard/zones/localzone/replicationgroups", Fn: getReplicationGroupsEvents},
	}

	mb.Registry.MustAddMetricSet(dellecs.ModuleName, metricsetName, New)
}

func getEndpoint(name string) (Endpoint, error) {
	endpoint, ok := endpoints[name]
	if !ok {
		return Endpoint{}, fmt.Errorf("%s not found in the map", name)
	}
	return endpoint, nil
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	config    *dellecs.Config
	logger    *logp.Logger
	ecsClient *ECSRestClient
	nodes     []NodeData
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The dellecs health metricset is beta.")

	config, err := dellecs.NewConfig(base)
	if err != nil {
		return nil, err
	}

	logger := logp.NewLogger(base.FullyQualifiedName())

	// Get the session cookie
	ecsClient, err := GetClient(config, base)
	if err != nil {
		logger.Errorf("Failed to get session cookie: %v", err)
		return nil, err
	}
	nodes, err := getLocalNodes(ecsClient)

	if err != nil {
		logger.Errorf("Failed to get session cookie: %v", err)
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		config:        config,
		logger:        logger,
		ecsClient:     ecsClient,
		nodes:         nodes,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	// accumulate errs and report them all at the end so that we don't
	// stop processing events if one of the fetches fails
	var errs []error

	for _, endpoint := range endpoints {
		m.logger.Debugf("Calling endpoint %s ....", endpoint.Name)
		events, err := endpoint.Fn(m)
		m.logger.Debugf("Fetched %d %s events", len(events), endpoint.Name)

		if err != nil {
			m.logger.Errorf("Error getting %s events: %s", endpoint.Name, err)
			errs = append(errs, err)
		} else {
			for _, event := range events {
				report.Event(event)
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("error while fetching system metrics: %w", errors.Join(errs...))
	}

	return nil

}
