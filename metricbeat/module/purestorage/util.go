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

package purestorage

import (
	"fmt"
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

func MakeRootFields(HostIp string) mapstr.M {
	return mapstr.M{
		"host.ip": HostIp,
	}
}

func CreateArray[T any](size int, defaultValue T) []T {
	array := make([]T, size)
	for i := range array {
		array[i] = defaultValue
	}
	return array
}
