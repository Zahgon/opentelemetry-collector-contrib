// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	traceIDField       = "traceID"
	spanIDField        = "spanID"
	parentSpanIDField  = "parentSpanID"
	nameField          = "name"
	kindField          = "kind"
	linksField         = "links"
	timeField          = "time"
	startTimeField     = "start"
	endTimeField       = "end"
	traceStateField    = "traceState"
	durationField      = "duration"
	attributeField     = "attribute"
	statusCodeField    = "statusCode"
	statusMessageField = "statusMessage"
	logsField          = "logs"
)

// traceDataToLogService translates trace data into the LogService format.
func traceDataToLogServiceData(td ptrace.Traces) []*sls.Log { _ = "STUB: not implemented"; return nil }

func resourceSpansToLogServiceData(resourceSpans ptrace.ResourceSpans) []*sls.Log {
	_ = "STUB: not implemented"
	return nil
}

func spanToLogServiceData(span ptrace.Span, resourceContents, instrumentationLibraryContents []*sls.LogContent) *sls.Log {
	_ = "STUB: not implemented"
	return nil
}

// pre alloc, refine if logContent's len > 16

// if ParentSpanID is not valid, the return "", it is compatible for log service

func spanKindToShortString(kind ptrace.SpanKind) string { _ = "STUB: not implemented"; return "" }

func statusCodeToShortString(code ptrace.StatusCode) string { _ = "STUB: not implemented"; return "" }

func eventsToString(events ptrace.SpanEventSlice) string { _ = "STUB: not implemented"; return "" }

func spanLinksToString(spanLinkSlice ptrace.SpanLinkSlice) string {
	_ = "STUB: not implemented"
	return ""
}
