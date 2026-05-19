// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

// ExceptionEventName the name of the exception event.
// TODO: Remove this when collector defines this semantic convention.
const ExceptionEventName = "exception"

func addCause(seg *awsxray.Segment, span ptrace.Span) { _ = "STUB: not implemented"; return }

// In the `rawExpectedSegmentForInstrumentedApp` X-Ray segment example in
// awsxray/tracesegment_test.go, you can see that sometimes the HTTP Response is not
// set but the Cause field is set for the root segment. And the actual HTTP status
// can only be found in one of the (nested) subsegmets. So we need to signal that
// in this case, the status of the span is not otlptrace.Status_Ok by
// temporarily setting the status to otlptrace.Status_UnknownError. This will be
// updated to a more specific error in the `segToSpans()` in translator.go once
// we traverse through all the subsegments.

// StatusCodeUnset is the default value for the span.Status().

// Right now the X-Ray exporter does not support the case where
// 1) CauseData is just a 16-char exception ID,
// 2) `WorkingDirectory` and `Paths`
// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/c615d2db351929b99e46f7b427f39c12afe15b54/exporter/awsxrayexporter/translator/cause.go#L107

// so we can only pass the cause exceptionID as the status message as a fallback mechanism

// not sure whether there are existing events, so
// append new empty events instead

// ID is a required field

func convertStackFramesToStackTraceStr(excp awsxray.Exception) string {
	_ = "STUB: not implemented"
	// resulting stacktrace looks like:
	// "<*excp.Type>: <*excp.Message>\n" +
	// "\tat <*frameN.Label>(<*frameN.Path>: <*frameN.Line>)\n"
	return ""
}

// the string representation of a frame looks like:
// <*frame.Label>(<*frame.Path>):line\n
