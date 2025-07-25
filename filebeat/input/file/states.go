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

package file

import (
	"sync"
	"time"

	"github.com/elastic/elastic-agent-libs/logp"
)

// States handles list of FileState. One must use NewStates to instantiate a
// file states registry. Using the zero-value is not safe.
type States struct {
	sync.RWMutex

	// states store
	states []State

	// idx maps state IDs to state indexes for fast lookup and modifications.
	idx map[string]int

	logger *logp.Logger
}

// NewStates generates a new states registry.
func NewStates(logger *logp.Logger) *States {
	return &States{
		states: nil,
		idx:    map[string]int{},
		logger: logger,
	}
}

// Update updates a state. If previous state didn't exist, new one is created
func (s *States) Update(newState State) {
	s.UpdateWithTs(newState, time.Now())
}

// UpdateWithTs updates a state, assigning the given timestamp.
// If previous state didn't exist, new one is created
func (s *States) UpdateWithTs(newState State, ts time.Time) {
	s.Lock()
	defer s.Unlock()

	id := newState.Id
	index := s.findPrevious(id)
	newState.Timestamp = ts

	if index >= 0 {
		s.states[index] = newState
	} else {
		// No existing state found, add new one
		s.idx[id] = len(s.states)
		s.states = append(s.states, newState)
		s.logger.Named("input").Debugf("New state added for %s", newState.Source)
	}
}

// FindPrevious lookups a registered state, that matching the new state.
// Returns a zero-state if no match is found.
func (s *States) FindPrevious(newState State) State {
	s.RLock()
	defer s.RUnlock()
	i := s.findPrevious(newState.Id)
	if i < 0 {
		return State{}
	}
	return s.states[i]
}

func (s *States) IsNew(state State) bool {
	s.RLock()
	defer s.RUnlock()
	i := s.findPrevious(state.Id)
	return i < 0
}

// findPrevious returns the previous state for the file.
// In case no previous state exists, index -1 is returned
func (s *States) findPrevious(id string) int {
	if i, exists := s.idx[id]; exists {
		return i
	}
	return -1
}

// Cleanup cleans up the state array. All states which are older then `older` are removed
// The number of states that were cleaned up and number of states that can be
// cleaned up in the future is returned.
func (s *States) Cleanup() (int, int) {
	return s.CleanupWith(nil)
}

// CleanupWith cleans up the state array. It calls `fn` with the state ID, for
// each entry to be removed.
func (s *States) CleanupWith(fn func(string)) (int, int) {
	s.Lock()
	defer s.Unlock()

	currentTime := time.Now()
	statesBefore := len(s.states)
	numCanExpire := 0

	L := len(s.states)
	for i := 0; i < L; {
		state := &s.states[i]
		canExpire := state.TTL > 0
		expired := (canExpire && currentTime.Sub(state.Timestamp) > state.TTL)

		if state.TTL == 0 || expired {
			if !state.Finished {
				s.logger.Errorf("State for %s should have been dropped, but couldn't as state is not finished.", state.Source)
				i++
				continue
			}

			delete(s.idx, state.Id)
			if fn != nil {
				fn(state.Id)
			}
			s.logger.Named("state").Debugf("State removed for %v because of older: %v", state.Source, state.TTL)

			L--
			if L != i {
				s.states[i] = s.states[L]
				s.idx[s.states[i].Id] = i
			}
		} else {
			i++
			if canExpire {
				numCanExpire++
			}
		}
	}

	s.states = s.states[:L]
	return statesBefore - L, numCanExpire
}

// Count returns number of states
func (s *States) Count() int {
	s.RLock()
	defer s.RUnlock()

	return len(s.states)
}

// GetStates creates copy of the file states.
func (s *States) GetStates() []State {
	s.RLock()
	defer s.RUnlock()

	newStates := make([]State, len(s.states))
	copy(newStates, s.states)

	return newStates
}

// SetStates overwrites all internal states with the given states array
func (s *States) SetStates(states []State) {
	s.Lock()
	defer s.Unlock()
	s.states = states

	// create new index
	s.idx = map[string]int{}
	for i := range states {
		s.idx[states[i].Id] = i
	}
}

// Copy create a new copy of the states object
func (s *States) Copy() *States {
	new := NewStates(s.logger)
	new.SetStates(s.GetStates())
	return new
}
