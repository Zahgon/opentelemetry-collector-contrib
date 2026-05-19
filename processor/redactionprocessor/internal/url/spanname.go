// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package url // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor/internal/url"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	semconv125 "go.opentelemetry.io/otel/semconv/v1.25.0"
	semconv138 "go.opentelemetry.io/otel/semconv/v1.40.0"
)

// SanitizeSpanName sanitizes the span name if the span looks like an HTTP span.
// It returns the sanitized name and true when a change was made.
func SanitizeSpanName(span ptrace.Span, sanitizer *URLSanitizer) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// This means the full span name was replaced

func wasFullyRedacted(s string) bool { _ = "STUB: not implemented"; return false }

func shouldSanitizeSpan(span ptrace.Span) bool { _ = "STUB: not implemented"; return false }

var httpAttributeKeys = []string{
	string(semconv138.HTTPRouteKey),
	string(semconv138.HTTPRequestMethodKey),
	string(semconv138.HTTPRequestMethodOriginalKey),
	string(semconv138.HTTPResponseStatusCodeKey),
	string(semconv138.URLFullKey),
	string(semconv125.HTTPSchemeKey),
	string(semconv125.HTTPTargetKey),
	string(semconv125.HTTPMethodKey),
	string(semconv125.HTTPURLKey),
}

func hasHTTPAttributes(attrs pcommon.Map) bool { _ = "STUB: not implemented"; return false }
