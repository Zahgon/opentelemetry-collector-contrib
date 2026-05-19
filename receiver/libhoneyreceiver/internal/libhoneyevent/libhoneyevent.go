// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package libhoneyevent // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/libhoneyevent"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	trc "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// FieldMapConfig is used to map the fields from the LibhoneyEvent to PData formats
type FieldMapConfig struct {
	Resources  ResourcesConfig  `mapstructure:"resources"`
	Scopes     ScopesConfig     `mapstructure:"scopes"`
	Attributes AttributesConfig `mapstructure:"attributes"`
}

// ResourcesConfig is used to map the fields from the LibhoneyEvent to PData formats
type ResourcesConfig struct {
	ServiceName string `mapstructure:"service_name"`
}

// ScopesConfig is used to map the fields from the LibhoneyEvent to PData formats
type ScopesConfig struct {
	LibraryName    string `mapstructure:"library_name"`
	LibraryVersion string `mapstructure:"library_version"`
}

// AttributesConfig is used to map the fields from the LibhoneyEvent to PData formats
type AttributesConfig struct {
	TraceID        string   `mapstructure:"trace_id"`
	ParentID       string   `mapstructure:"parent_id"`
	SpanID         string   `mapstructure:"span_id"`
	Name           string   `mapstructure:"name"`
	Error          string   `mapstructure:"error"`
	SpanKind       string   `mapstructure:"spankind"`
	DurationFields []string `mapstructure:"durationFields"`
}

// LibhoneyEvent is the event structure from libhoney
type LibhoneyEvent struct {
	Samplerate       int            `json:"samplerate" msgpack:"samplerate"`
	MsgPackTimestamp *time.Time     `msgpack:"time"`
	Time             string         `json:"time"` // should not be trusted. use MsgPackTimestamp
	Data             map[string]any `json:"data" msgpack:"data"`
}

// UnmarshalJSON overrides the unmarshall to make sure the MsgPackTimestamp is set
func (l *LibhoneyEvent) UnmarshalJSON(j []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalMsgpack overrides the unmarshall to make sure the MsgPackTimestamp is set
func (l *LibhoneyEvent) UnmarshalMsgpack(data []byte) error { _ = "STUB: not implemented"; return nil }

// Use a temporary struct to avoid recursion

// Ignore during msgpack unmarshal

// First unmarshal into the temp struct

// Copy fields to our tmp struct

// Check if Time field exists in Data and extract it

// DebugString returns a string representation of the LibhoneyEvent
func (l *LibhoneyEvent) DebugString() string { _ = "STUB: not implemented"; return "" }

// SignalType returns the type of signal this event represents. Only log is implemented for now.
func (l *LibhoneyEvent) SignalType(logger zap.Logger) string { _ = "STUB: not implemented"; return "" }

// GetService returns the service name from the event or the dataset name if no service name is found.
func (l *LibhoneyEvent) GetService(fields FieldMapConfig, seen *ServiceHistory, dataset string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetScope returns the scope key for the event. If the scope has not been seen before, it creates a new one.
func (l *LibhoneyEvent) GetScope(fields FieldMapConfig, seen *ScopeHistory, serviceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// if we've seen it, we don't expect it to be different right away so we'll just return it.

// otherwise, we need to make a new found scope

// Create a per-service default scope instead of using a global default

// Create a default scope for this specific service

func spanIDFrom(s string) trc.SpanID { _ = "STUB: not implemented"; return *new(trc.SpanID) }

func traceIDFrom(s string) trc.TraceID { _ = "STUB: not implemented"; return *new(trc.TraceID) }

func generateAnID(length int) []byte { _ = "STUB: not implemented"; return nil }

// SimpleScope is a simple struct to hold the scope data
type SimpleScope struct {
	ServiceName    string
	LibraryName    string
	LibraryVersion string
	ScopeSpans     ptrace.SpanSlice
	ScopeLogs      plog.LogRecordSlice
}

// ScopeHistory is a map of scope keys to the SimpleScope object
type ScopeHistory struct {
	Scope map[string]SimpleScope // key here is service.name+library.name
}

// ServiceHistory is a map of service names to the number of times they've been seen
type ServiceHistory struct {
	NameCount map[string]int
}

// ToPLogRecord converts a LibhoneyEvent to a Pdata LogRecord
func (l *LibhoneyEvent) ToPLogRecord(newLog *plog.LogRecord, alreadyUsedFields *[]string, logger zap.Logger) error {
	_ = "STUB: not implemented"
	// Handle cases where MsgPackTimestamp might be nil (e.g., JSON data from Refinery)
	return nil
}

// Parse time from Time field or use current time

// undoing this is gonna be complicated: https://github.com/honeycombio/husky/blob/91c0498333cd9f5eed1fdb8544ca486db7dea565/otlp/logs.go#L61

// GetParentID returns the parent id from the event or an error if it's not found
func (l *LibhoneyEvent) GetParentID(fieldName string) (trc.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(trc.SpanID), nil
}

// Extract 8 bytes for SpanID

// If it's a TraceID (16+ bytes), take the last 8 bytes as SpanID

// If it's already 8 bytes, use as-is

// ToPTraceSpan converts a LibhoneyEvent to a Pdata Span
func (l *LibhoneyEvent) ToPTraceSpan(newSpan *ptrace.Span, alreadyUsedFields *[]string, cfg FieldMapConfig, logger zap.Logger) error {
	_ = "STUB: not implemented"
	// Handle cases where MsgPackTimestamp might be nil (e.g., JSON data from Refinery)
	return nil
}

// Parse time from Time field or use current time

// Convert slice to [16]byte array

// Convert slice to [8]byte array
