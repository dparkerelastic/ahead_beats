// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package httpjson

import (
	inputcursor "github.com/elastic/beats/v7/filebeat/input/v2/input-cursor"
	"github.com/elastic/beats/v7/libbeat/management/status"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

type cursor struct {
	cfg cursorConfig

	state mapstr.M

	status status.StatusReporter
	log    *logp.Logger
}

func newCursor(cfg cursorConfig, stat status.StatusReporter, log *logp.Logger) *cursor {
	return &cursor{cfg: cfg, status: stat, log: log}
}

func (c *cursor) load(cursor *inputcursor.Cursor) {
	if c == nil || cursor == nil || cursor.IsNew() {
		c.log.Debug("new cursor: nothing loaded")
		return
	}

	if c.state == nil {
		c.state = mapstr.M{}
	}

	if err := cursor.Unpack(&c.state); err != nil {
		c.log.Errorf("Reset cursor state. Failed to read from registry: %v", err)
		c.status.UpdateStatus(status.Degraded, "failed to load cursor: "+err.Error())
		return
	}

	c.log.Debugf("cursor loaded: %v", c.state)
}

func (c *cursor) update(trCtx *transformContext) {
	if c.cfg == nil {
		return
	}

	if c.state == nil {
		c.state = mapstr.M{}
	}

	for k, cfg := range c.cfg {
		stat := c.status
		if cfg.mustIgnoreEmptyValue() {
			stat = ignoreEmptyValueReporter{stat}
		}
		v, _ := cfg.Value.Execute(trCtx, transformable{}, k, cfg.Default, stat, c.log)
		if v != "" || !cfg.mustIgnoreEmptyValue() {
			_, _ = c.state.Put(k, v)
			c.log.Debugf("cursor.%s stored with %s", k, v)
		}
	}
}

// ignoreEmptyValueReporter is an abuse of the type system to allow the cursor
// update mechanism to signal to valueTpl.Execute not to report empty values
// as health degraded.
type ignoreEmptyValueReporter struct {
	status.StatusReporter
}

func (c *cursor) clone() mapstr.M {
	if c == nil || c.state == nil {
		return mapstr.M{}
	}
	return c.state.Clone()
}
