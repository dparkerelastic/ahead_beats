package purestorage

import (
	"github.com/elastic/beats/v7/metricbeat/mb"
)

const (
	ModuleName = "purestorage"
)

type Config struct {
	HostIp     string `config:"host_ip" validate:"required"`
	ApiKey     string `config:"api_key" validate:"required"`
	ApiVersion string `config:"api_key" validate:"required"`
	Port       uint   `config:"port"`
	DebugMode  string `config:"api_debug_mode"`
}

func NewConfig(base mb.BaseMetricSet) (*Config, error) {
	config := Config{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil

}
