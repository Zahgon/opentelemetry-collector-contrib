// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinv2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv2"

import (
	zipkinmodel "github.com/openzipkin/zipkin-go/model"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	spanEventDataFormat = "%s|%s|%d"
	spanLinkDataFormat  = "%s|%s|%s|%s|%d"
)

var sampled = true

// FromTranslator converts from pdata to Zipkin data model.
type FromTranslator struct{}

// FromTraces translates internal trace data into Zipkin v2 spans.
// Returns a slice of Zipkin SpanModel's.
func (FromTranslator) FromTraces(td ptrace.Traces) ([]*zipkinmodel.SpanModel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resourceSpansToZipkinSpans(rs ptrace.ResourceSpans, estSpanCount int) ([]*zipkinmodel.SpanModel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractScopeTags(il pcommon.InstrumentationScope, zTags map[string]string) {
	_ = "STUB: not implemented"
	return
}

func spanToZipkinSpan(
	span ptrace.Span,
	localServiceName string,
	zTags map[string]string,
) (*zipkinmodel.SpanModel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// leave timestamp unset on zs (zipkin span) if
// otel span startTime is zero.  Zipkin has a
// case where startTime is not set on the span.
// See handling of this (and setting of otel span
// to unix time zero) in zipkinv2_to_traces.go

func populateStatus(status ptrace.Status, zs *zipkinmodel.SpanModel, tags map[string]string) {
	_ = "STUB: not implemented"
	return
}

// The error tag should only be set if Status is Error. If a boolean version
// ({"error":false} or {"error":"false"}) is present, it SHOULD be removed.
// Zipkin will treat any span with error sent as failed.

// Per specs, Span Status MUST be reported as a key-value pair in tags to Zipkin, unless it is UNSET.
// In the latter case it MUST NOT be reported.
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/sdk_exporters/zipkin.md#status

func aggregateSpanTags(span ptrace.Span, zTags map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func spanEventsToZipkinAnnotations(events ptrace.SpanEventSlice, zs *zipkinmodel.SpanModel) error {
	_ = "STUB: not implemented"
	return nil
}

func spanLinksToZipkinTags(links ptrace.SpanLinkSlice, zTags map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func attributeMapToStringMap(attrMap pcommon.Map) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func removeRedundantTags(redundantKeys map[string]bool, zTags map[string]string) {
	_ = "STUB: not implemented"
	return
}

func resourceToZipkinEndpointServiceNameAndAttributeMap(
	resource pcommon.Resource,
) (serviceName string, zTags map[string]string) {
	_ = "STUB: not implemented"
	return "", nil
}

func extractZipkinServiceName(zTags map[string]string) string { _ = "STUB: not implemented"; return "" }

func spanKindToZipkinKind(kind ptrace.SpanKind) zipkinmodel.Kind {
	_ = "STUB: not implemented"
	return *new(zipkinmodel.Kind)
}

func zipkinEndpointFromTags(
	zTags map[string]string,
	localServiceName string,
	remoteEndpoint bool,
	redundantKeys map[string]bool,
) (endpoint *zipkinmodel.Endpoint) {
	_ = "STUB: not implemented"
	return nil
}

func isIPv6Address(ipStr string) bool { _ = "STUB: not implemented"; return false }

func convertTraceID(t pcommon.TraceID) zipkinmodel.TraceID {
	_ = "STUB: not implemented"
	return *new(zipkinmodel.TraceID)
}

func convertSpanID(s pcommon.SpanID) zipkinmodel.ID {
	_ = "STUB: not implemented"
	return *new(zipkinmodel.ID)
}
