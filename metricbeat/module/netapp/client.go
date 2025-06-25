package netapp

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/elastic/beats/v7/metricbeat/mb"
)

type NetAppRestClient struct {
	config  *Config
	baseUrl string
	client  *http.Client
	headers map[string]string
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
	}

	return &client, nil
}

func (c *NetAppRestClient) Get(endpoint string) (string, error) {
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
