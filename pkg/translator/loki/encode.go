// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loki // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// JSON representation of the LogRecord as described by https://developers.google.com/protocol-buffers/docs/proto3#json
type lokiEntry struct {
	Name                 string                `json:"name,omitempty"`
	Body                 json.RawMessage       `json:"body,omitempty"`
	TraceID              string                `json:"traceid,omitempty"`
	SpanID               string                `json:"spanid,omitempty"`
	Severity             string                `json:"severity,omitempty"`
	Flags                uint32                `json:"flags,omitempty"`
	Attributes           map[string]any        `json:"attributes,omitempty"`
	Resources            map[string]any        `json:"resources,omitempty"`
	InstrumentationScope *instrumentationScope `json:"instrumentation_scope,omitempty"`
}

type instrumentationScope struct {
	Name       string         `json:"name,omitempty"`
	Version    string         `json:"version,omitempty"`
	Attributes map[string]any `json:"attributes,omitempty"`
}

// Encode converts an OTLP log record and its resource attributes into a JSON
// string representing a Loki entry. An error is returned when the record can't
// be marshaled into JSON.
func Encode(lr plog.LogRecord, res pcommon.Resource, scope pcommon.InstrumentationScope) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// EncodeLogfmt converts an OTLP log record and its resource attributes into a logfmt
// string representing a Loki entry. An error is returned when the record can't
// be marshaled into logfmt.
func EncodeLogfmt(lr plog.LogRecord, res pcommon.Resource, scope pcommon.InstrumentationScope) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// todo handle maps, slices

func serializeBodyJSON(body pcommon.Value) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no body

func bodyToKeyvals(body pcommon.Value) []any { _ = "STUB: not implemented"; return nil }

// try to parse record body as logfmt, but failing that assume it's plain text

func valueToKeyvals(key string, value pcommon.Value) []any { _ = "STUB: not implemented"; return nil }

// if given key:value pair already exists in keyvals, replace value. Otherwise append
func keyvalsReplaceOrAppend(keyvals []any, key string, value any) []any {
	_ = "STUB: not implemented"
	return nil
}

func parseLogfmtLine(line string) (*[]any, error) { _ = "STUB: not implemented"; return nil, nil }
