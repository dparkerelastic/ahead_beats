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
	"os"
	"strings"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/purestorage"
)

type PureRestClient struct {
	config  *purestorage.Config
	baseUrl string
	client  *http.Client
	headers map[string]string
}

func GetClient(config *purestorage.Config, base mb.BaseMetricSet) (*PureRestClient, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	psClient := PureRestClient{
		config:  config,
		baseUrl: fmt.Sprintf("https://%s/api/%s/", config.Host, config.ApiVersion),
		client:  &http.Client{Transport: tr},
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	// // Get the session cookie
	// cookie, err := getSessionCookie(config.Host, config.ApiVersion, config.ApiKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("error initializing PureStorage client: %w", err)
	// }
	// psClient.headers = map[string]string{
	// 	"Cookie":       cookie,
	// 	"Content-Type": "application/json",
	// }

	return &psClient, nil
}

func (c *PureRestClient) login() error {
	cookie, err := getSessionCookie(c.config.Host, c.config.ApiVersion, c.config.ApiKey)
	if err != nil {
		return fmt.Errorf("error initializing PureStorage client: %w", err)
	}
	c.headers["Cookie"] = cookie
	return nil
}

func getSessionCookie(hostName, apiVersion, apiToken string) (string, error) {
	var cookie string
	fileName := fmt.Sprintf("tmp/%s-sessionCookie.dat", hostName)

	// first try to get cookie data from cache file
	os.MkdirAll("tmp", os.ModePerm)
	sessionFile := fileName

	// If there is a session file, return the contents
	if _, err := os.Stat(sessionFile); err == nil {
		data, err := os.ReadFile(sessionFile)
		if err != nil {
			return "", err
		}
		cookie = strings.TrimSpace(string(data))
		if isCookieValid(cookie) {
			return cookie, nil
		}
	}

	// cached cookie is invalid, so get a new session from the server
	apiURL := fmt.Sprintf("https://%s/api/%s/auth/session", hostName, apiVersion)
	headers := map[string]string{"Content-Type": "application/json"}
	data := map[string]string{"api_token": apiToken}
	jsonData, _ := json.Marshal(data)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error %d: %s", resp.StatusCode, string(body))
	}

	// loop through http header fields
	cookie = resp.Header.Get("Set-Cookie")

	// is our new cookie valid (it must be)?
	if !isCookieValid(cookie) {
		return "", fmt.Errorf("error retrieving valid session cookie")
	}

	// replace the session cache file contents & return the cookie
	err = os.WriteFile(sessionFile, []byte(cookie), 0644)
	if err != nil {
		return "", err
	}
	return cookie, nil
}

func isCookieValid(cookie string) bool {
	dateFormat := "Mon, 02-Jan-2006 15:04:05 MST"
	dateFormatNoDash := "Mon, 02 Jan 2006 15:04:05 MST"
	now := time.Now()

	if cookie == "" || !strings.Contains(cookie, "; ") {
		return false
	}
	cookieContents := strings.Split(cookie, "; ")
	expireObj := cookieContents[1]
	if !strings.Contains(expireObj, "=") {
		return false
	}

	expireDateElement := strings.Split(expireObj, "=")
	expireDate := expireDateElement[1]
	var parsedExpireDate time.Time
	var err error
	if strings.Contains(expireDate, "-") {
		parsedExpireDate, err = time.Parse(dateFormat, expireDate)
	} else {
		parsedExpireDate, err = time.Parse(dateFormatNoDash, expireDate)
	}
	if err != nil {
		return false
	}
	return now.Before(parsedExpireDate)
}

func (c *PureRestClient) Get(endpoint string) (string, error) {
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
