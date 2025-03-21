package health

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
// 	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
// 	return nil
// }

// basicAuth converts the given username & password to Base64 encoded string.
// func basicAuth(username, password string) string {
// 	auth := username + ":" + password
// 	return base64.StdEncoding.EncodeToString([]byte(auth))
// }

func GetInstanceData(m *MetricSet, hostInfo Connection, url string) ([]byte, error) {

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	if m.authToken == "" {
		fmt.Println("Auth Token is empty, fetching a new token...")
		newToken, err := m.fetchAuthToken()
		if err != nil {
			return nil, fmt.Errorf("failed to fetch auth token: %w", err)
		}
		m.authToken = newToken
		fmt.Println("New Auth Token: ", m.authToken)
	} else {
		fmt.Println("Auth Token is already set, using existing token...")
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Citrix-CustomerId", hostInfo.customerId)
	req.Header.Set("Authorization", fmt.Sprintf("CWSAuth bearer=%s", m.authToken))

	// cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: nil})
	client := &http.Client{
		Timeout: 90 * time.Second,
	}

	// client.Jar = cookieJar
	// Execute the HTTP request
	response, err := client.Do(req)
	// Close the response body to avoid resource leaks
	if err != nil {
		fmt.Println("Error executing HTTP request:", err)
		if strings.Contains(err.Error(), "401") {
			fmt.Println("Token expired, fetching a new token...")
			newToken, tokenErr := m.fetchAuthToken()
			if tokenErr != nil {
				return nil, fmt.Errorf("failed to fetch new auth token: %w", tokenErr)
			}
			m.authToken = newToken
			fmt.Println("New Auth Token: ", m.authToken)

			// Retry fetching machine data with the new token
			response, err = client.Do(req)
			if err != nil {
				fmt.Println("Error executing HTTP request after refreshing token:", err)
				return nil, fmt.Errorf("failed to fetch machine data after refreshing token: %w", err)
			}
		} else {
			fmt.Println("Error executing HTTP request:", err)
			return nil, err
		}
	}
	defer response.Body.Close()
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	return responseData, nil

}

func reportMetricsForCitrixCMS(reporter mb.ReporterV2, baseURL string, metrics ...[]mapstr.M) {
	for _, metricSlice := range metrics {
		for _, metric := range metricSlice {
			event := mb.Event{ModuleFields: mapstr.M{"base_url": baseURL}}
			if ts, ok := metric["@timestamp"]; ok {
				t, err := time.Parse(time.RFC3339, ts.(string))
				if err == nil {
					// if the timestamp parsing fails, we just fall back to the event time
					// (and leave the additional timestamp in the event for posterity)
					event.Timestamp = t
					delete(metric, "@timestamp")
				}
			}

			for k, v := range metric {
				if !isEmpty(v) {
					//fmt.Println("k =" + k + " v=" + string(v))
					event.ModuleFields.Put(k, v)
				}
			}

			reporter.Event(event)
		}
	}
}

func isEmpty(value interface{}) bool {
	// we make use of the fact that all the dashboard API responses utilize
	// pointers for non-string types to filter out empty values from metric events.

	if value == nil {
		return true
	}

	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		return reflect.ValueOf(value).IsNil()
	}

	if t.Kind() == reflect.Slice || t.Kind() == reflect.String {
		return reflect.ValueOf(value).Len() == 0
	}

	return false
}

func GetMetrics[T any](m *MetricSet, hostInfo Connection, url string, jsonInfo T) (any, string, error) {
	responseData, err := GetInstanceData(m, hostInfo, url)

	if err != nil {
		fmt.Println("Error fetching instance data:", err)
		return jsonInfo, "", err
	}

	err = json.Unmarshal(responseData, &jsonInfo)
	if err != nil {
		fmt.Println("Error unmarshalling JSON response:", err)
		fmt.Println("responseData:", string(responseData))
		// Close the response body to avoid resource leaks
		return jsonInfo, "", err
	}

	return jsonInfo, string(responseData), nil

}

func reportMetrics(reporter mb.ReporterV2, baseURL string, data CMSData, debug bool) {
	metrics := []mapstr.M{}

	for _, machineLoadData := range data.machineCurrentLoadIndex.Value {
		metric := mapstr.M{}
		//metric["health.machine.id"] = machineLoadData.ID
		v := reflect.ValueOf(machineLoadData)
		t := reflect.TypeOf(machineLoadData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if !isEmpty(fieldValue.Interface()) {
				metricKey := fmt.Sprintf("health.machine.%s", field.Name)
				metric[metricKey] = fieldValue.Interface()
			}
		}
		// Add the message field if debug is enabled
		// This is useful for debugging purposes to see the message returned by the API
		// when the machine load index is fetched
		// if debug {
		// 	metric["health.message"] = data.system.Message
		// }

		metrics = append(metrics, metric)
	}

	// #TODO FIX THIS no org id
	reportMetricsForCitrixCMS(reporter, baseURL, metrics)
}

func (m *MetricSet) fetchAuthToken() (string, error) {
	apiURL := fmt.Sprintf("%s/cctrustoauth2/%s/tokens/clients", m.Host(), m.customerId)

	// Prepare form data
	formData := map[string]string{
		"grant_type":    "client_credentials",
		"client_secret": m.clientSecret,
		"client_id":     m.clientId,
	}

	// Encode form data
	form := ""
	for key, value := range formData {
		form += fmt.Sprintf("%s=%s&", key, value)
	}
	form = form[:len(form)-1] // Remove trailing '&'

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(form))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Extract the token from the response
	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", fmt.Errorf("failed to parse response JSON: %w", err)
	}

	token, ok := responseData["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	return token, nil
}
