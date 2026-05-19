// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func makeHTTP(span ptrace.Span) (map[string]pcommon.Value, *awsxray.HTTPData) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prefer HTTP forwarded information (AttributeHTTPClientIP) when present.

// Prefer HTTP forwarded information (AttributeHTTPClientIP) when present.

// Didn't have any HTTP-specific information so don't need to fill it in segment

func extractResponseSizeFromEvents(span ptrace.Span) int64 {
	_ = "STUB: not implemented"
	// Support instrumentation that sets response size in span or as an event.
	return 0
}

func extractResponseSizeFromAttributes(attributes pcommon.Map) int64 {
	_ = "STUB: not implemented"
	return 0
}

func constructClientURL(urlParts map[string]string) string {
	_ = "STUB: not implemented"
	// follows OpenTelemetry specification-defined combinations for client spans described in
	// https://github.com/open-telemetry/semantic-conventionsv112/blob/main/docs/http/http-spans.md#http-client
	return ""
}

// full URL available so no need to assemble

func constructServerURL(urlParts map[string]string) string {
	_ = "STUB: not implemented"
	// follows OpenTelemetry specification-defined combinations for server spans described in
	// https://github.com/open-telemetry/semantic-conventionsv112/blob/main/docs/http/http-spans.md#http-server
	return ""
}

// full URL available so no need to assemble
