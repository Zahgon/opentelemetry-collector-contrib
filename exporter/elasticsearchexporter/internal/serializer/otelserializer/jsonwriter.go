// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelserializer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer"

import (
	"bytes"
	"unicode/utf8"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
)

// jsonWriter is a low-overhead JSON writer that writes directly to a
// bytes.Buffer, avoiding the interface-dispatch and state-machine overhead
// of go-structform/json.Visitor.
type jsonWriter struct {
	buf *bytes.Buffer
}

func newJSONWriter(buf *bytes.Buffer) jsonWriter {
	_ = "STUB: not implemented"
	return *new(jsonWriter)
}

func (w *jsonWriter) startObject() { _ = "STUB: not implemented"; return }

func (w *jsonWriter) endObject() { _ = "STUB: not implemented"; return }

func (w *jsonWriter) startArray() { _ = "STUB: not implemented"; return }

func (w *jsonWriter) endArray() { _ = "STUB: not implemented"; return }

// key writes a JSON object key with a preceding comma if first is false.
// Returns false (the new value for the caller's "first" tracking variable).
func (w *jsonWriter) key(k string, first bool) bool { _ = "STUB: not implemented"; return false }

func (w *jsonWriter) boolVal(b bool) { _ = "STUB: not implemented"; return }

func (w *jsonWriter) nullVal() { _ = "STUB: not implemented"; return }

func (w *jsonWriter) int64Val(n int64) { _ = "STUB: not implemented"; return }

func (w *jsonWriter) uint64Val(n uint64) { _ = "STUB: not implemented"; return }

// float64Val writes a float64, always including a radix point (e.g. 1.0 not 1)
// to preserve type information for ES dynamic mapping.
func (w *jsonWriter) float64Val(val float64) { _ = "STUB: not implemented"; return }

// Insert ".0" before exponent.
// Copy tail for reuse below. Any write to buf would overwrite the
// remaining b content, leading to a corruption in the tail part.
// tail length is based on IEEE 754 max exponent of +308 or min exponent of -324, padded
// for alignment.

func (w *jsonWriter) arrayComma(first bool) bool { _ = "STUB: not implemented"; return false }

// jsonString writes a JSON-escaped string (with surrounding quotes).
// Uses the same HTML-safe escaping as go-structform and go-fastjson.
func (w *jsonWriter) jsonString(s string) { _ = "STUB: not implemented"; return }

const hexChars = "0123456789abcdef"

// htmlSafeSet matches go-structform's htmlEscapeSet (inverted): true means safe (no escape needed).
var htmlSafeSet [utf8.RuneSelf]bool

func init() {
	for i := range htmlSafeSet {
		htmlSafeSet[i] = true
	}
	for i := range 32 {
		htmlSafeSet[i] = false
	}
	for _, c := range `\"` {
		htmlSafeSet[c] = false
	}
	for _, c := range "&<>" {
		htmlSafeSet[c] = false
	}
}

// writeTimestampField writes "@timestamp" or similar with msec.nsec format.
func (w *jsonWriter) writeTimestampField(key string, timestamp pcommon.Timestamp, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeTimestampEpochMillisField(key string, timestamp pcommon.Timestamp, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeUIntField(key string, val uint64, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeStringFieldSkipDefault(key, value string, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeIntFieldSkipDefault(key string, val int64, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeTraceIDField(id pcommon.TraceID, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeSpanIDField(key string, id pcommon.SpanID, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeDataStream(idx elasticsearch.Index, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeResource(resource pcommon.Resource, resourceSchemaURL string, stringifyMapAttributes, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeScope(scope pcommon.InstrumentationScope, scopeSchemaURL string, stringifyMapAttributes, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeAttributes(attributes pcommon.Map, stringifyMapValues, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *jsonWriter) writeValue(val pcommon.Value, stringifyMaps bool) {
	_ = "STUB: not implemented"
	return
}

func (w *jsonWriter) writeMap(m pcommon.Map) { _ = "STUB: not implemented"; return }

func (w *jsonWriter) writeGeolocationAttributes(attributes pcommon.Map, first bool) bool {
	_ = "STUB: not implemented"
	return false
}
