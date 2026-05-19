// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinv1 // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv1"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type status struct {
	codePtr *ptrace.StatusCode
	message string
}

const (
	tagZipkinCensusCode    = "census.status_code"
	tagZipkinCensusMsg     = "census.status_description"
	tagZipkinOpenCensusMsg = "opencensus.status_description"
)

// statusMapper contains codes translated from different sources to OC status codes
type statusMapper struct {
	// oc status code extracted from "status.code" tags
	fromStatus status
	// oc status code extracted from "census.status_code" tags
	fromCensus status
	// oc status code extracted from "http.status_code" tags
	fromHTTP status
	// oc status code extracted from "error" tags
	fromErrorTag status
	// oc status code 'unknown' when the "error" tag exists but is invalid
	fromErrorTagUnknown status
}

// status fills the given ptrace.Status from the best possible extraction source.
// It'll first try to return status extracted from "census.status_code" to account for zipkin
// then fallback on code extracted from "status.code" tags
// and finally fallback on code extracted and translated from "http.status_code"
// status must be called after all tags/attributes are processed with the `fromAttribute` method.
func (m *statusMapper) status(dest ptrace.Status) { _ = "STUB: not implemented"; return }

func (m *statusMapper) fromAttribute(key string, attrib pcommon.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// Convert oc status (OK == 0) to otel (OK == 1), anything else is error.

// Keep the code as is, even if unknown. Since we are allowed to receive unknown values for enums.

// The status is stored with the "error" key, but otel does not care about the old census values.
// See https://github.com/census-instrumentation/opencensus-go/blob/1eb9a13c7dd02141e065a665f6bf5c99a090a16a/exporter/zipkin/zipkin.go#L160-L165

// attribToStatusCode maps an integer or string attribute value to a status code.
// The function return nil if the value is of another type or cannot be converted to an int32 value.
func attribToStatusCode(attr pcommon.Value) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func toInt32(i int64) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// statusCodeFromHTTP takes an HTTP status code and return the appropriate OpenTelemetry status code
// See: https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/http.md
func statusCodeFromHTTP(code int32) ptrace.StatusCode {
	_ = "STUB: not implemented"
	return *new(ptrace.StatusCode)
}
