// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunk // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/splunk"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	// Keys are taken from https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/overview.md#trace-context-in-legacy-formats.
	// spanIDFieldKey is the key used in log event for the span id (if any).
	spanIDFieldKey = "span_id"
	// traceIDFieldKey is the key used in the log event for the trace id (if any).
	traceIDFieldKey = "trace_id"
)

func LogToSplunkEvent(res pcommon.Resource, lr plog.LogRecord, toOtelAttrs HecToOtelAttrs, toHecAttrs OtelToHecFields, source, sourceType, index string) *Event {
	_ = "STUB: not implemented"
	return nil
}

// events with no body are rejected by

// ignore

// ignore

func mergeValue(dst map[string]any, k string, v pcommon.Value) { _ = "STUB: not implemented"; return }

func isArrayFlat(array pcommon.Slice) bool { _ = "STUB: not implemented"; return false }

func flattenAndMergeMap(src pcommon.Map, dst map[string]any, key string) {
	_ = "STUB: not implemented"
	return
}
