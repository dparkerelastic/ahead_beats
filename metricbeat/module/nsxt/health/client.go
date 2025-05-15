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
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/nsxt"
)

type NsxtRestClient struct {
	config  *nsxt.Config
	baseUrl string
	client  *http.Client
	headers map[string]string
}

func disableSSLVerification() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func GetClient(config *nsxt.Config, base mb.BaseMetricSet) (*NsxtRestClient, error) {
	tr := disableSSLVerification()
	auth := base64.StdEncoding.EncodeToString([]byte(config.Username + ":" + config.Password))

	client := NsxtRestClient{
		config:  config,
		baseUrl: fmt.Sprintf("https://%s:%d/", config.Host, config.Port),
		client:  &http.Client{Transport: tr},
		headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Basic " + auth,
		},
	}

	return &client, nil
}

func (c *NsxtRestClient) Get(endpoint string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+endpoint, nil)
	if err != nil {
		return "", err
	}
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the HTTP status code indicates an error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body) // Read the body for additional error details
		return "", fmt.Errorf("server error: %s (status code: %d)", string(body), resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil

}
