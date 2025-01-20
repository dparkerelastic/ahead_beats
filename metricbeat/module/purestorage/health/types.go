package health

import "time"

// ArrayController represents the JSON structure for the array controllers endpoint.
type ArrayController struct {
	Status  string  `json:"status"`
	Name    string  `json:"name"`
	Version *string `json:"version"`
	Mode    string  `json:"mode"`
	Model   *string `json:"model"`
}

// ArrayMonitor represents the JSON structure for the array monitor endpoint.
type ArrayMonitor struct {
	WritesPerSec   int       `json:"writes_per_sec"`
	UsecPerWriteOp int       `json:"usec_per_write_op"`
	OutputPerSec   int       `json:"output_per_sec"`
	ReadsPerSec    int       `json:"reads_per_sec"`
	InputPerSec    int       `json:"input_per_sec"`
	Time           time.Time `json:"time"`
	UsecPerReadOp  int       `json:"usec_per_read_op"`
	QueueDepth     *int      `json:"queue_depth"`
}

// ArraySpace represents the JSON structure for the array space endpoint.
type ArraySpace struct {
	Capacity         int64   `json:"capacity"`
	Hostname         string  `json:"hostname"`
	System           int     `json:"system"`
	Snapshots        int     `json:"snapshots"`
	Volumes          int64   `json:"volumes"`
	DataReduction    float64 `json:"data_reduction"`
	Total            int64   `json:"total"`
	SharedSpace      int64   `json:"shared_space"`
	ThinProvisioning float64 `json:"thin_provisioning"`
	TotalReduction   float64 `json:"total_reduction"`
}

// Hardware represents the JSON structure for the hardware endpoint.
type Hardware struct {
	Status      string  `json:"status"`
	Slot        *int    `json:"slot"`
	Name        string  `json:"name"`
	Temperature *int    `json:"temperature"`
	Index       int     `json:"index"`
	Identify    *string `json:"identify"`
	Speed       *int    `json:"speed"`
	Details     *string `json:"details"`
}

// Drive represents the JSON structure for the drive endpoint.
type Drive struct {
	Status   string `json:"status"`
	Capacity int64  `json:"capacity"`
	Type     string `json:"type"`
	Name     string `json:"name"`
}

// PGroup represents the JSON structure for the protection group endpoint.
type PGroup struct {
	Name                 string     `json:"name"`
	PhysicalBytesWritten *int64     `json:"physical_bytes_written"`
	Started              *time.Time `json:"started"`
	Completed            *time.Time `json:"completed"`
	Created              time.Time  `json:"created"`
	Source               string     `json:"source"`
	TimeRemaining        *int64     `json:"time_remaining"`
	Progress             *int       `json:"progress"`
	DataTransferred      *int64     `json:"data_transferred"`
}

// VolumeSpace represents the JSON structure for the volume space endpoint.
type VolumeSpace struct {
	Size             int64   `json:"size"`
	Name             string  `json:"name"`
	System           *string `json:"system"`
	Snapshots        int64   `json:"snapshots"`
	Volumes          int64   `json:"volumes"`
	DataReduction    float64 `json:"data_reduction"`
	Total            int64   `json:"total"`
	SharedSpace      *int64  `json:"shared_space"`
	ThinProvisioning float64 `json:"thin_provisioning"`
	TotalReduction   float64 `json:"total_reduction"`
}

type VolumeMonitor struct {
	WritesPerSec   int    `json:"writes_per_sec"`
	Name           string `json:"name"`
	UsecPerWriteOp int    `json:"usec_per_write_op"`
	OutputPerSec   int    `json:"output_per_sec"`
	ReadsPerSec    int    `json:"reads_per_sec"`
	InputPerSec    int    `json:"input_per_sec"`
	Time           string `json:"time"`
	UsecPerReadOp  int    `json:"usec_per_read_op"`
}

// Array represents the JSON structure for the array endpoint.
type Array struct {
	Revision  string `json:"revision"`
	Version   string `json:"version"`
	ArrayName string `json:"array_name"`
	ID        string `json:"id"`
}

// Volume represents the JSON structure for the volume endpoint.
type Volume struct {
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Source  string    `json:"source"`
	Serial  string    `json:"serial"`
	Size    int64     `json:"size"`
}
