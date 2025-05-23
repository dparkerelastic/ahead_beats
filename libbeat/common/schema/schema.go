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

package schema

import (
	"errors"

	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// Schema describes how a map[string]interface{} object can be parsed and converted into
// an event. The conversions can be described using an (optionally nested) mapstr.M
// that contains Conv objects.
type Schema map[string]Mapper

// Mapper interface represents a valid type to be used in a schema.
type Mapper interface {
	// Map applies the Mapper conversion on the data and adds the result
	// to the event on the key.
	Map(key string, event mapstr.M, data map[string]interface{}) []error

	HasKey(key string) bool
}

// A Conv object represents a conversion mechanism from the data map to the event map.
type Conv struct {
	Func            Converter // Convertor function
	Key             string    // The key in the data map
	Optional        bool      // Whether to ignore errors if the key is not found
	Required        bool      // Whether to provoke errors if the key is not found
	IgnoreAllErrors bool      // Ignore any value conversion error
}

// Converter function type
type Converter func(key string, data map[string]interface{}) (interface{}, error)

// Map applies the conversion on the data and adds the result
// to the event on the key.
func (conv Conv) Map(key string, event mapstr.M, data map[string]interface{}) []error {
	value, err := conv.Func(conv.Key, data)
	if err != nil {
		var keyErr *KeyNotFoundError
		if errors.As(err, &keyErr) {
			keyErr.Optional = conv.Optional
			keyErr.Required = conv.Required
		}
		if conv.IgnoreAllErrors {
			logp.Debug("schema", "ignoring error for key %q: %s", key, err)
			return nil
		}
		return []error{err}
	}
	event[key] = value
	return nil
}

func (conv Conv) HasKey(key string) bool {
	return conv.Key == key
}

// implements Mapper interface for structure
type Object map[string]Mapper

// Map applies the schema for an object
func (o Object) Map(key string, event mapstr.M, data map[string]interface{}) []error {
	subEvent := mapstr.M{}
	errs := applySchemaToEvent(subEvent, data, o)
	event[key] = subEvent
	return errs
}

func (o Object) HasKey(key string) bool {
	return hasKey(key, o)
}

// ApplyTo adds the fields extracted from data, converted using the schema, to the
// event map.
func (s Schema) ApplyTo(event mapstr.M, data map[string]interface{}, opts ...ApplyOption) (mapstr.M, []error) {
	if len(opts) == 0 {
		opts = DefaultApplyOptions
	}
	errors := applySchemaToEvent(event, data, s)
	for _, opt := range opts {
		event, errors = opt(event, errors)
	}
	return event, errors
}

// Apply converts the fields extracted from data, using the schema, into a new map and reports back the errors.
func (s Schema) Apply(data map[string]interface{}, opts ...ApplyOption) (mapstr.M, error) {
	event, errs := s.ApplyTo(mapstr.M{}, data, opts...)
	return event, errors.Join(errs...)
}

// HasKey checks if the key is part of the schema
func (s Schema) HasKey(key string) bool {
	return hasKey(key, s)
}

func hasKey(key string, mappers map[string]Mapper) bool {
	for _, mapper := range mappers {
		if mapper.HasKey(key) {
			return true
		}
	}
	return false
}

func applySchemaToEvent(event mapstr.M, data map[string]interface{}, conversions map[string]Mapper) []error {
	var errs []error
	for key, mapper := range conversions {
		if err := mapper.Map(key, event, data); err != nil {
			errs = append(errs, err...)
		}
	}
	return errs
}

// SchemaOption is for adding optional parameters to the conversion
// functions
type SchemaOption func(c Conv) Conv

// Optional sets the optional flag, that suppresses the error in case
// the key doesn't exist
func Optional(c Conv) Conv {
	c.Optional = true
	return c
}

// Required sets the required flag, that provokes an error in case the key
// doesn't exist, even if other missing keys can be ignored
func Required(c Conv) Conv {
	c.Required = true
	return c
}

// IgnoreAllErrors set the enable all errors flag
func IgnoreAllErrors(c Conv) Conv {
	c.IgnoreAllErrors = true
	return c
}

// setOptions adds the optional flags to the Conv object
func SetOptions(c Conv, opts []SchemaOption) Conv {
	for _, opt := range opts {
		c = opt(c)
	}
	return c
}
