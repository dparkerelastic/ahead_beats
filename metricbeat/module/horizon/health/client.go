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
	"bytes"
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

	client := HorizonRestClient{
		config:  config,
		baseUrl: fmt.Sprintf("http://%s:%d/", config.Host, config.Port),
		client:  &http.Client{Transport: tr},
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return &client, nil
}
func (c *HorizonRestClient) login() error {
	apiToken, err := login(c.config)
	if err != nil {
		return err
	}

	// Update the Authorization header with the new token
	c.headers["Authorization"] = "Bearer " + apiToken
	return nil
}

func login(config *horizon.Config) (string, error) {
	url := fmt.Sprintf("http://%s:%d/rest/login", config.Host, config.Port)
	payload, err := buildLoginBody(config)
	if err != nil {
		return "", fmt.Errorf("failed to build login body: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call login URL: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		var login LoginToken
		if err := json.Unmarshal([]byte(body), &login); err != nil {
			return "", fmt.Errorf("failed to parse login response: %v", err)
		}

		// FIXME: might need to use the refresh token if doing login in Fetch does not work for some reason

		return login.AccessToken, nil
	} else {
		return "", fmt.Errorf("failed to login: %d", response.StatusCode)
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

func buildLoginBody(config *horizon.Config) ([]byte, error) {
	payload := map[string]string{
		"domain":   "AD-TEST-DOMAIN",
		"username": "Administrator",
		"password": config.Password,
	}
	return json.Marshal(payload)
}

func (c *HorizonRestClient) Post(endpoint string, body interface{}) (string, error) {
	payloadBytes, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.baseUrl+endpoint, io.NopCloser(bytes.NewReader(payloadBytes)))
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
