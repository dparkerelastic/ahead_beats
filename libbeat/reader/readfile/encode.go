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

package readfile

import (
	"bytes"
	"io"
	"time"

	"github.com/elastic/beats/v7/libbeat/reader"
	"github.com/elastic/beats/v7/libbeat/reader/readfile/encoding"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// Reader produces lines by reading lines from an io.Reader
// through a decoder converting the reader it's encoding to utf-8.
type EncoderReader struct {
	reader *LineReader
}

// Config stores the configuration for the readers required to read
// a file line by line
type Config struct {
	Codec      encoding.Encoding
	BufferSize int
	Terminator LineTerminator
	MaxBytes   int
	// If CollectOnEOF is set to true (default false) the line reader will return the buffer if EOF reached: this
	// will ensure full content including last line with no EOL will be returned for fully retrieved content that's
	// not appended anymore between reads.
	// If CollectOnEOF is set to false the line reader will return 0 content and keep the buffer at the current
	// state of appending data after temporarily EOF.
	CollectOnEOF bool
}

// NewEncodeReader creates a new Encode reader from input reader by applying
// the given codec.
func NewEncodeReader(r io.ReadCloser, config Config, logger *logp.Logger) (EncoderReader, error) {
	eReader, err := NewLineReader(r, config, logger)
	return EncoderReader{eReader}, err
}

// Next reads the next line from it's initial io.Reader
// This converts a io.Reader to a reader.reader
func (r EncoderReader) Next() (reader.Message, error) {
	c, sz, err := r.reader.Next()
	// Creating message object
	return reader.Message{
		Ts:      time.Now(),
		Content: bytes.Trim(c, "\xef\xbb\xbf"),
		Bytes:   sz,
		Fields:  mapstr.M{},
	}, err
}

func (r EncoderReader) Close() error {
	return r.reader.Close()
}
