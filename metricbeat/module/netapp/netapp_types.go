package netapp

import "time"

// General API response objects
type Records[T any] struct {
	NumRecords int    `json:"num_records"`
	Records    []T    `json:"records"`
	Error      Status `json:"error"`
}

type Status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// General types
type NamedObject struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type Metrics struct {
	Timestamp  time.Time `json:"timestamp"`
	Duration   string    `json:"duration"`
	Status     string    `json:"status"`
	Throughput IOLatency `json:"throughput"`
	IOPS       IOLatency `json:"iops"`
	Latency    IOLatency `json:"latency"`
}

type Statistics struct {
	Timestamp     time.Time `json:"timestamp"`
	Status        string    `json:"status"`
	ThroughputRaw IOLatency `json:"throughput_raw"`
	IOPSRaw       IOLatency `json:"iops_raw"`
	LatencyRaw    IOLatency `json:"latency_raw"`
}

type IOLatency struct {
	Read  int `json:"read"`
	Write int `json:"write"`
	Other int `json:"other"`
	Total int `json:"total"`
}

type Link struct {
	Href string `json:"href"`
}

type IPAddress struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type IPInterface struct {
	IP   IPAddress `json:"ip"`
	Name string    `json:"name"`
	UUID string    `json:"uuid"`
}
