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

package add_cloud_metadata

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v7/libbeat/beat"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/logp/logptest"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func initQCloudTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/meta-data/instance-id" {
			_, _ = w.Write([]byte("ins-qcloudv5"))
			return
		}
		if r.RequestURI == "/meta-data/placement/region" {
			_, _ = w.Write([]byte("china-south-gz"))
			return
		}
		if r.RequestURI == "/meta-data/placement/zone" {
			_, _ = w.Write([]byte("gz-azone2"))
			return
		}

		http.Error(w, "not found", http.StatusNotFound)
	}))
}

func TestRetrieveQCloudMetadata(t *testing.T) {
	logp.TestingSetup()

	server := initQCloudTestServer()
	defer server.Close()

	config, err := conf.NewConfigFrom(map[string]interface{}{
		"providers": []string{"tencent"},
		"host":      server.Listener.Addr().String(),
	})

	if err != nil {
		t.Fatal(err)
	}

	p, err := New(config, logptest.NewTestingLogger(t, ""))
	if err != nil {
		t.Fatal(err)
	}

	actual, err := p.Run(&beat.Event{Fields: mapstr.M{}})
	if err != nil {
		t.Fatal(err)
	}

	expected := mapstr.M{
		"cloud": mapstr.M{
			"provider": "qcloud",
			"instance": mapstr.M{
				"id": "ins-qcloudv5",
			},
			"region":            "china-south-gz",
			"availability_zone": "gz-azone2",
			"service": mapstr.M{
				"name": "CVM",
			},
		},
	}
	assert.Equal(t, expected, actual.Fields)
}
