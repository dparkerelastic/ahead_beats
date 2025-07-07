package netapp

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/elastic/beats/v7/metricbeat/mb"
)

type NetAppRestClient struct {
	config        *Config
	baseUrl       string
	client        *http.Client
	headers       map[string]string
	returnTimeout int
}

func disableSSLVerification() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func GetClient(config *Config, base mb.BaseMetricSet) (*NetAppRestClient, error) {
	tr := disableSSLVerification()

	username := config.Username
	password := config.Password

	// Encode credentials
	creds := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	client := NetAppRestClient{
		config:  config,
		baseUrl: fmt.Sprintf("%s://%s:%d", config.Protocol, config.Host, config.Port),
		client:  &http.Client{Transport: tr},
		headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Basic " + creds,
		},
		returnTimeout: config.ReturnTimeout,
	}

	return &client, nil
}

func (c *NetAppRestClient) CreateNetAppRequest(endpoint string, fields []string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if len(fields) > 0 {
		q.Add("fields", strings.Join(fields, ","))
	}
	q.Add("return_records", "true")
	q.Add("return_timeout", fmt.Sprintf("%d", c.returnTimeout))
	req.URL.RawQuery = q.Encode()

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (c *NetAppRestClient) GetWithFields(endpoint string, fields []string) (string, error) {
	req, err := c.CreateNetAppRequest(endpoint, fields)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("server error: %s (status code: %d)", string(body), resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
