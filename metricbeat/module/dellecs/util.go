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

package dellecs

import (
	"fmt"
	"net"
	"strings"

	"github.com/elastic/elastic-agent-libs/mapstr"
)

func StringToBool(s string) (bool, error) {
	s = strings.ToLower(s)
	switch s {
	case "yes":
		return true, nil
	case "true":
		return true, nil
	case "no":
		return false, nil
	case "false":
		return false, nil
	}

	// Default to false
	return false, fmt.Errorf("invalid value: %s", s)
}

func MakeRootFields(config *Config) mapstr.M {
	return mapstr.M{
		"host.ip":   config.HostInfo.IP,
		"host.name": config.HostInfo.Hostname,
	}
}

func CreateArray[T any](size int, defaultValue T) []T {
	array := make([]T, size)
	for i := range array {
		array[i] = defaultValue
	}
	return array
}

type HostInfo struct {
	IP       string
	Hostname string
}

func GetHostInfo(input string) (HostInfo, error) {
	var hostInfo HostInfo

	// Try to parse the input as an IP address
	ip := net.ParseIP(input)
	if ip != nil {
		hostInfo.IP = ip.String()
		// Perform a reverse lookup to get the hostname
		names, err := net.LookupAddr(ip.String())
		if err != nil {
			// If the reverse lookup fails, set the hostname to "hostname not found" and let the calling
			// function handle the error by logging a warning - we don't want to quit over this
			hostInfo.Hostname = "hostname not found"
			return hostInfo, err
		}
		if len(names) > 0 {
			hostInfo.Hostname = names[0]
		} else {
			hostInfo.Hostname = "hostname not found"
		}

	} else {
		// Try to resolve the input as a hostname
		addrs, err := net.LookupHost(input)
		if err != nil {
			return hostInfo, fmt.Errorf("failed to lookup IP for hostname %s: %v", input, err)
		}
		if len(addrs) > 0 {
			hostInfo.IP = addrs[0]
			hostInfo.Hostname = input
		}
	}

	return hostInfo, nil
}

func convertToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
