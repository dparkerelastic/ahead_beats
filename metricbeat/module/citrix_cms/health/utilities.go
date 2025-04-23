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

	if strings.Contains(string(responseData), "Invalid bearer token") {
		fmt.Println("Invalid bearer token detected, fetching a new token...")
		newToken, tokenErr := m.fetchAuthToken()
		if tokenErr != nil {
			return jsonInfo, "", fmt.Errorf("failed to fetch new auth token: %w", tokenErr)
		}
		m.authToken = newToken
		fmt.Println("New Auth Token: ", m.authToken)

		// Retry fetching instance data with the new token
		responseData, err = GetInstanceData(m, hostInfo, url)
		if err != nil {
			fmt.Println("Error fetching instance data after refreshing token:", err)
			return jsonInfo, "", fmt.Errorf("failed to fetch instance data after refreshing token: %w", err)
		}
	}

	output, err := ExcludeNullValues(responseData)
	if err != nil {
		fmt.Println("Error excluding null values:", err)
		return jsonInfo, "", err
	}
	err = json.Unmarshal(output, &jsonInfo)

	//err = json.Unmarshal(responseData, &jsonInfo)

	if err != nil {
		fmt.Println("Error unmarshalling JSON response:", err)
		if strings.Contains(string(responseData), "Invalid bearer token") {
			m.authToken = ""
		}
		return jsonInfo, "", err
	}

	return jsonInfo, string(output), nil
	//return jsonInfo, string(responseData), nil

}

func GetInstanceData(m *MetricSet, hostInfo Connection, url string) ([]byte, error) {

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

	client := &http.Client{
		Timeout: 90 * time.Second,
	}

	response, err := client.Do(req)
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

func reportMetrics(reporter mb.ReporterV2, baseURL string, data CMSData, debug bool) {
	metrics := []mapstr.M{}

	for _, metricData := range data.machineCurrentLoadIndex.Value {
		metric := mapstr.M{}

		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

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

	for _, metricData := range data.serverOSDesktopSummaries.Value {
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if !isEmpty(fieldValue.Interface()) {
				metricKey := fmt.Sprintf("health.server.os.desktop.summary.%s", field.Name)
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

	for _, metricData := range data.loadIndexSummaries.Value {
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {
				metricKey := fmt.Sprintf("health.load.index.summary.%s", field.Name)
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

	for _, metricData := range data.machineMetric.Value {
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if !isEmpty(fieldValue.Interface()) {
				metricKey := fmt.Sprintf("health.machine.metric.%s", field.Name)
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

	for _, metricData := range data.sessionActivitySummaries_Agg1.Value {
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if !isEmpty(fieldValue.Interface()) {
				metricKey := fmt.Sprintf("health.session.activity.summaries.aggregate.%s", field.Name)
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

	// for _, metricData := range data.logonMetricDetails.Value {
	// 	metric := mapstr.M{}
	// 	//metric["health.machine.id"] = metricData.ID
	// 	v := reflect.ValueOf(metricData)
	// 	t := reflect.TypeOf(metricData)

	// 	for i := 0; i < t.NumField(); i++ {
	// 		field := t.Field(i)
	// 		fieldValue := v.Field(i)

	// 		if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {
	// 			metricKey := fmt.Sprintf("health.logon.metric.details.%s", field.Name)
	// 			metric[metricKey] = fieldValue.Interface()
	// 		}
	// 	}
	// 	// Add the message field if debug is enabled
	// 	// This is useful for debugging purposes to see the message returned by the API
	// 	// when the machine load index is fetched
	// 	// if debug {
	// 	// 	metric["health.message"] = data.system.Message
	// 	// }

	// 	metrics = append(metrics, metric)
	// }

	for _, metricData := range data.machineMetricDetails.Value {
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {
				metricKey := fmt.Sprintf("health.machine.metric.details.%s", field.Name)
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

	// for _, metricData := range data.sessionMetricsDetails.Value {
	// 	metric := mapstr.M{}
	// 	//metric["health.machine.id"] = metricData.ID
	// 	v := reflect.ValueOf(metricData)
	// 	t := reflect.TypeOf(metricData)

	// 	for i := 0; i < t.NumField(); i++ {
	// 		field := t.Field(i)
	// 		fieldValue := v.Field(i)

	// 		if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {
	// 			metricKey := fmt.Sprintf("health.session.metric.details.%s", field.Name)
	// 			metric[metricKey] = fieldValue.Interface()
	// 		}
	// 	}
	// 	// Add the message field if debug is enabled
	// 	// This is useful for debugging purposes to see the message returned by the API
	// 	// when the machine load index is fetched
	// 	// if debug {
	// 	// 	metric["health.message"] = data.system.Message
	// 	// }

	// 	metrics = append(metrics, metric)
	// }

	for _, metricData := range data.sessionDetails.Value {
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				//fmt.Println("Field Name:", field.Name)
				if field.Name == "SessionMetrics" {
					if fieldValue.Kind() == reflect.Slice && fieldValue.Len() > 0 {
						lastValue := fieldValue.Index(fieldValue.Len() - 1).Interface()
						metricKey := fmt.Sprintf("health.session.details.%s.last", field.Name)
						metric[metricKey] = lastValue
					}
				} else {
					metricKey := fmt.Sprintf("health.session.details.%s", field.Name)
					metric[metricKey] = fieldValue.Interface()
				}
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

	for _, metricData := range data.sessionFailureDetails.Value {

		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				//fmt.Println("Field Name:", field.Name)
				if field.Name == "SessionMetrics" {
					if fieldValue.Kind() == reflect.Slice && fieldValue.Len() > 0 {
						lastValue := fieldValue.Index(fieldValue.Len() - 1).Interface()
						metricKey := fmt.Sprintf("health.session.failure.details.%s.last", field.Name)
						metric[metricKey] = lastValue
					}
				} else {
					metricKey := fmt.Sprintf("health.session.failure.details.%s", field.Name)
					metric[metricKey] = fieldValue.Interface()
				}
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

	for _, metricData := range data.machineDetails.Value {

		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				//fmt.Println("Field Name:", field.Name)
				if field.Name == "SessionMetrics" {
					if fieldValue.Kind() == reflect.Slice && fieldValue.Len() > 0 {
						lastValue := fieldValue.Index(fieldValue.Len() - 1).Interface()
						metricKey := fmt.Sprintf("health.machine.details.%s.last", field.Name)
						metric[metricKey] = lastValue
					}
				} else {
					metricKey := fmt.Sprintf("health.machine.details.%s", field.Name)
					metric[metricKey] = fieldValue.Interface()
				}
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

	for _, metricData := range data.resourceUtilizationSummary.ResourceUtilizationSummaryEntries {

		//fmt.Println("Current UTC Time (Truncated to Hour, Minus 1 Hour):", time.Now().UTC().Truncate(time.Hour).Add(-1*time.Hour).Format("2006-01-02T15:04:05Z"))
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				//fmt.Println("Field Name:", field.Name)
				metricKey := fmt.Sprintf("health.resource.utilization.summary.one.hour.%s", field.Name)
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

	for _, metricData := range data.resourceUtilization.ResourceUtilizationEntries {

		//fmt.Println("Current UTC Time (Truncated to Hour, Minus 1 Hour):", time.Now().UTC().Truncate(time.Hour).Add(-1*time.Hour).Format("2006-01-02T15:04:05Z"))
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				//fmt.Println("Field Name:", field.Name)
				metricKey := fmt.Sprintf("health.resource.utilization.five.minute.%s", field.Name)
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

	for _, metricData := range data.logOnSummaries.LogOnSummariesEntries {

		//fmt.Println("Current UTC Time (Truncated to Hour, Minus 1 Hour):", time.Now().UTC().Truncate(time.Hour).Add(-1*time.Hour).Format("2006-01-02T15:04:05Z"))
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				fmt.Println("Field Name:", field.Name)
				metricKey := fmt.Sprintf("health.logon.summaries.one.hour.%s", field.Name)
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

	for _, metricData := range data.machineSummaries.MachineSummariesEntries {

		//fmt.Println("Current UTC Time (Truncated to Hour, Minus 1 Hour):", time.Now().UTC().Truncate(time.Hour).Add(-1*time.Hour).Format("2006-01-02T15:04:05Z"))
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				fmt.Println("Field Name:", field.Name)
				metricKey := fmt.Sprintf("health.machine.summaries.%s", field.Name)
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

	for _, metricData := range data.machineMetricSummary.MachineMetricSummaryEntries {

		//fmt.Println("Current UTC Time (Truncated to Hour, Minus 1 Hour):", time.Now().UTC().Truncate(time.Hour).Add(-1*time.Hour).Format("2006-01-02T15:04:05Z"))
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				fmt.Println("Field Name:", field.Name)
				metricKey := fmt.Sprintf("health.machine.metric.summary.%s", field.Name)
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

	for _, metricData := range data.sessionActivitySummaries.SessionActivitySummariesEntries {

		//fmt.Println("Current UTC Time (Truncated to Hour, Minus 1 Hour):", time.Now().UTC().Truncate(time.Hour).Add(-1*time.Hour).Format("2006-01-02T15:04:05Z"))
		metric := mapstr.M{}
		//metric["health.machine.id"] = metricData.ID
		v := reflect.ValueOf(metricData)
		t := reflect.TypeOf(metricData)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if fieldValue.IsValid() && fieldValue.CanInterface() && !isEmpty(fieldValue.Interface()) {

				fmt.Println("Field Name:", field.Name)
				metricKey := fmt.Sprintf("health.session.activity.summaries.%s", field.Name)
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

	reportMetricsForCitrixCMS(reporter, baseURL, metrics)
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

func ExcludeNullValues(input []byte) ([]byte, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(input, &raw); err != nil {
		return nil, err
	}

	filtered := filterNullValues(raw)
	return json.Marshal(filtered)
}

func filterNullValues(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			if value != nil {
				filteredValue := filterNullValues(value)
				if filteredValue != nil {
					result[key] = filteredValue
				}
			}
		}
		return result
	case []interface{}:
		var result []interface{}
		for _, value := range v {
			if value != nil {
				filteredValue := filterNullValues(value)
				if filteredValue != nil {
					result = append(result, filteredValue)
				}
			}
		}
		return result
	default:
		return data
	}
}

func RemoveResourceUtilizationDuplicatesByMachineID(resourceUtilization *ResourceUtilization_JSON) {
	uniqueEntries := make(map[string]ResourceUtilizationEntry)
	for _, entry := range resourceUtilization.ResourceUtilizationEntries {
		if existing, found := uniqueEntries[entry.MachineID]; found {
			// Keep the entry with the latest CollectedDate
			if entry.CollectedDate.After(*existing.CollectedDate) {
				uniqueEntries[entry.MachineID] = entry
			}
		} else {
			uniqueEntries[entry.MachineID] = entry
		}
	}

	// Convert the map back to a slice
	resourceUtilization.ResourceUtilizationEntries = make([]ResourceUtilizationEntry, 0, len(uniqueEntries))
	for _, entry := range uniqueEntries {
		resourceUtilization.ResourceUtilizationEntries = append(resourceUtilization.ResourceUtilizationEntries, entry)
	}
}

func RemoveResourceUtilizationSummaryDuplicatesByMachineID(resourceUtilizationSummary *ResourceUtilizationSummary_JSON) {
	uniqueEntries := make(map[string]ResourceUtilizationSummaryEntry)
	for _, entry := range resourceUtilizationSummary.ResourceUtilizationSummaryEntries {
		if existing, found := uniqueEntries[entry.MachineID]; found {
			// Keep the entry with the latest SummaryDate
			if entry.SummaryDate.After(*existing.SummaryDate) {
				uniqueEntries[entry.MachineID] = entry
			}
		} else {
			uniqueEntries[entry.MachineID] = entry
		}
	}

	// Convert the map back to a slice
	resourceUtilizationSummary.ResourceUtilizationSummaryEntries = make([]ResourceUtilizationSummaryEntry, 0, len(uniqueEntries))
	for _, entry := range uniqueEntries {
		resourceUtilizationSummary.ResourceUtilizationSummaryEntries = append(resourceUtilizationSummary.ResourceUtilizationSummaryEntries, entry)
	}
}

func RemoveMachineSummariesDuplicatesByDesktopGroupID(machineSummaries *MachineSummaries_JSON) {
	uniqueEntries := make(map[string]MachineSummariesEntry)
	for _, entry := range machineSummaries.MachineSummariesEntries {
		if existing, found := uniqueEntries[entry.DesktopGroupID]; found {
			// Keep the entry with the latest SummaryDate
			if entry.SummaryDate.After(*existing.SummaryDate) {
				uniqueEntries[entry.DesktopGroupID] = entry
			}
		} else {
			uniqueEntries[entry.DesktopGroupID] = entry
		}
	}

	// Convert the map back to a slice
	machineSummaries.MachineSummariesEntries = make([]MachineSummariesEntry, 0, len(uniqueEntries))
	for _, entry := range uniqueEntries {
		machineSummaries.MachineSummariesEntries = append(machineSummaries.MachineSummariesEntries, entry)
	}
}

func RemoveMachineMetricSummaryDuplicatesByMachineID(machineMetricSummary *MachineMetricSummary_JSON) {
	uniqueEntries := make(map[string]MachineMetricSummaryEntry)
	for _, entry := range machineMetricSummary.MachineMetricSummaryEntries {
		if existing, found := uniqueEntries[entry.MachineID]; found {
			// Keep the entry with the latest SummaryDate
			if entry.SummaryDate.After(*existing.SummaryDate) {
				uniqueEntries[entry.MachineID] = entry
			}
		} else {
			uniqueEntries[entry.MachineID] = entry
		}
	}

	// Convert the map back to a slice
	machineMetricSummary.MachineMetricSummaryEntries = make([]MachineMetricSummaryEntry, 0, len(uniqueEntries))
	for _, entry := range uniqueEntries {
		machineMetricSummary.MachineMetricSummaryEntries = append(machineMetricSummary.MachineMetricSummaryEntries, entry)
	}
}

func RemoveSessionActivitySummariesDuplicatesByDesktopGroupID(sessionActivitySummaries *SessionActivitySummaries_JSON) {
	uniqueEntries := make(map[string]SessionActivitySummariesEntry)
	for _, entry := range sessionActivitySummaries.SessionActivitySummariesEntries {
		if existing, found := uniqueEntries[entry.DesktopGroupID]; found {
			// Keep the entry with the latest SummaryDate
			if entry.SummaryDate.After(*existing.SummaryDate) {
				uniqueEntries[entry.DesktopGroupID] = entry
			}
		} else {
			uniqueEntries[entry.DesktopGroupID] = entry
		}
	}

	// Convert the map back to a slice
	sessionActivitySummaries.SessionActivitySummariesEntries = make([]SessionActivitySummariesEntry, 0, len(uniqueEntries))
	for _, entry := range uniqueEntries {
		sessionActivitySummaries.SessionActivitySummariesEntries = append(sessionActivitySummaries.SessionActivitySummariesEntries, entry)
	}
}

func RemoveLogOnSummariesDuplicatesByDesktopGroupID(logOnSummaries *LogOnSummaries_JSON) {
	uniqueEntries := make(map[string]LogOnSummariesEntry)
	for _, entry := range logOnSummaries.LogOnSummariesEntries {
		if existing, found := uniqueEntries[entry.DesktopGroupID]; found {
			// Keep the entry with the latest SummaryDate
			if entry.SummaryDate.After(*existing.SummaryDate) {
				uniqueEntries[entry.DesktopGroupID] = entry
			}
		} else {
			uniqueEntries[entry.DesktopGroupID] = entry
		}
	}

	// Convert the map back to a slice
	logOnSummaries.LogOnSummariesEntries = make([]LogOnSummariesEntry, 0, len(uniqueEntries))
	for _, entry := range uniqueEntries {
		logOnSummaries.LogOnSummariesEntries = append(logOnSummaries.LogOnSummariesEntries, entry)
	}
}
