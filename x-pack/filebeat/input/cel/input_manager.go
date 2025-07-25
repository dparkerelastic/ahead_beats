// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cel

import (
	"github.com/elastic/go-concert/unison"

	v2 "github.com/elastic/beats/v7/filebeat/input/v2"
	inputcursor "github.com/elastic/beats/v7/filebeat/input/v2/input-cursor"
	"github.com/elastic/beats/v7/libbeat/statestore"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
)

// inputManager wraps one stateless input manager
// and one cursor input manager. It will create one or the other
// based on the config that is passed.
type InputManager struct {
	cursor *inputcursor.InputManager
}

var _ v2.InputManager = InputManager{}

func NewInputManager(log *logp.Logger, store statestore.States) InputManager {
	return InputManager{
		cursor: &inputcursor.InputManager{
			Logger:     log,
			StateStore: store,
			Type:       inputName,
			Configure:  cursorConfigure,
		},
	}
}

func cursorConfigure(cfg *conf.C, logger *logp.Logger) ([]inputcursor.Source, inputcursor.Input, error) {
	src := &source{cfg: defaultConfig()}
	if err := cfg.Unpack(&src.cfg); err != nil {
		return nil, nil, err
	}
	src.cfg.checkUnsupportedParams(logger)
	return []inputcursor.Source{src}, input{}, nil
}

// checkUnsupportedParams checks if unsupported/deprecated/discouraged paramaters are set and logs a warning
func (c config) checkUnsupportedParams(logger *logp.Logger) {
	if c.RecordCoverage {
		logger.Named("cel").Warn("execution coverage enabled: " +
			"see documentation for details: https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#cel-record-coverage")
	}
	if c.Redact == nil {
		logger.Named("cel").Warn("missing recommended 'redact' configuration: " +
			"see documentation for details: https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#cel-state-redact")
	}
}

type source struct{ cfg config }

func (s *source) Name() string { return s.cfg.Resource.URL.String() }

// Init initializes both wrapped input managers.
func (m InputManager) Init(grp unison.Group) error {
	return m.cursor.Init(grp)
}

// Create creates a cursor input manager.
func (m InputManager) Create(cfg *conf.C) (v2.Input, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, err
	}
	return m.cursor.Create(cfg)
}
