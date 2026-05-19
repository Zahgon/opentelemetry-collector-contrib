// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package entry // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"

import (
	"sync"
	"time"
)

var timeNow = time.Now

// Entry is a flexible representation of log data associated with a timestamp.
type Entry struct {
	ObservedTimestamp time.Time      `json:"observed_timestamp"      yaml:"observed_timestamp"`
	Timestamp         time.Time      `json:"timestamp"               yaml:"timestamp"`
	Body              any            `json:"body"                    yaml:"body"`
	Attributes        map[string]any `json:"attributes,omitempty"    yaml:"attributes,omitempty"`
	Resource          map[string]any `json:"resource,omitempty"      yaml:"resource,omitempty"`
	SeverityText      string         `json:"severity_text,omitempty" yaml:"severity_text,omitempty"`
	SpanID            []byte         `json:"span_id,omitempty"       yaml:"span_id,omitempty"`
	TraceID           []byte         `json:"trace_id,omitempty"      yaml:"trace_id,omitempty"`
	TraceFlags        []byte         `json:"trace_flags,omitempty"   yaml:"trace_flags,omitempty"`
	Severity          Severity       `json:"severity"                yaml:"severity"`
	ScopeName         string         `json:"scope_name"              yaml:"scope_name"`
}

var entriesPool = sync.Pool{
	New: func() any {
		return &Entry{}
	},
}

var zeroE = &Entry{}

// New will create a new log entry with current timestamp and an empty body.
func New() *Entry { _ = "STUB: not implemented"; return nil }

// Put releases the entry back to the pool
func Put(e *Entry) {
	_ = "STUB: not implemented"

	// AddAttribute will add a key/value pair to the entry's attributes.
	return
}

func (entry *Entry) AddAttribute(key, value string) { _ = "STUB: not implemented"; return }

// AddResourceKey wil add a key/value pair to the entry's resource.
func (entry *Entry) AddResourceKey(key, value string) { _ = "STUB: not implemented"; return }

// Get will return the value of a field on the entry, including a boolean indicating if the field exists.
func (entry *Entry) Get(field FieldInterface) (any, bool) {
	_ = "STUB: not implemented"
	return *

	// Set will set the value of a field on the entry.
	new(any), false
}

func (entry *Entry) Set(field FieldInterface, val any) error { _ = "STUB: not implemented"; return nil }

// Delete will delete a field from the entry.
func (entry *Entry) Delete(field FieldInterface) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Read will read the value of a field into a designated interface.
func (entry *Entry) Read(field FieldInterface, dest any) error {
	_ = "STUB: not implemented"
	return nil
}

// readToInterface reads a field to a designated interface pointer.
func (entry *Entry) readToInterface(field FieldInterface, dest *any) error {
	_ = "STUB: not implemented"
	return nil
}

// readToString reads a field to a designated string pointer.
func (entry *Entry) readToString(field FieldInterface, dest *string) error {
	_ = "STUB: not implemented"
	return nil
}

// readToInterfaceMap reads a field to a designated map interface pointer.
func (entry *Entry) readToInterfaceMap(field FieldInterface, dest *map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// readToStringMap reads a field to a designated map string pointer.
func (entry *Entry) readToStringMap(field FieldInterface, dest *map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Copy will return a deep copy of the entry.
func (entry *Entry) Copy() *Entry { _ = "STUB: not implemented"; return nil }
