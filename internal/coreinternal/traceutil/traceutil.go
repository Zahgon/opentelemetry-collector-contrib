// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traceutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/traceutil"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// SpanKindStr returns a string representation of the SpanKind as it's defined in the proto.
// The function provides old behavior of ptrace.SpanKind.String() to support graceful adoption of
// https://github.com/open-telemetry/opentelemetry-collector/pull/6250.
func SpanKindStr(sk ptrace.SpanKind) string { _ = "STUB: not implemented"; return "" }

// StatusCodeStr returns a string representation of the StatusCode as it's defined in the proto.
// The function provides old behavior of ptrace.StatusCode.String() to support graceful adoption of
// https://github.com/open-telemetry/opentelemetry-collector/pull/6250.
func StatusCodeStr(sk ptrace.StatusCode) string { _ = "STUB: not implemented"; return "" }

// SpanIDToHexOrEmptyString returns a hex string from SpanID.
// An empty string is returned, if SpanID is empty.
func SpanIDToHexOrEmptyString(id pcommon.SpanID) string { _ = "STUB: not implemented"; return "" }

// TraceIDToHexOrEmptyString returns a hex string from TraceID.
// An empty string is returned, if TraceID is empty.
func TraceIDToHexOrEmptyString(id pcommon.TraceID) string { _ = "STUB: not implemented"; return "" }
