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
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/horizon"
)

type HorizonRestClient struct {
	config  *horizon.Config
	baseUrl string
	client  *http.Client
	headers map[string]string
}

func disableSSLVerification() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func GetClient(config *horizon.Config, base mb.BaseMetricSet) (*HorizonRestClient, error) {
	tr := disableSSLVerification()

	apiToken, err := login(config)
	if err != nil {
		return nil, err
	}

	client := HorizonRestClient{
		config:  config,
		baseUrl: fmt.Sprintf("https://%s:%d/", config.Host, config.Port),
		client:  &http.Client{Transport: tr},
		headers: map[string]string{
			"Authorizaion": "Bearer " + apiToken,
			"Content-Type": "application/json",
		},
	}

	return &client, nil
}

func login(config *horizon.Config) (string, error) {
	url := fmt.Sprintf("https://%s:%d/login.json", config.Host, config.Port)
	client := &http.Client{Transport: disableSSLVerification()}

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(config.Username, config.Password)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to login: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return "", fmt.Errorf("failed to parse login response: %v", err)
		}

		accessToken, ok := result["access_token"].(string)
		if !ok {
			return "", fmt.Errorf("access_token not found in login response")
		}

		// For now we will just login in the Fetch method, so we don't need to refresh the token
		// (token expires after 10 minutes)
		// refreshToken, ok := result["refresh_token"].(string)
		// if !ok {
		// 	return "", fmt.Errorf("refresh_token not found in login response")
		// }

		return accessToken, nil
	} else {
		return "", fmt.Errorf("failed to login: %d", resp.StatusCode)
	}

}

func (c *HorizonRestClient) Get(endpoint string) (string, error) {
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil

}
