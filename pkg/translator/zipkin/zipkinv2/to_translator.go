// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinv2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv2"

import (
	zipkinmodel "github.com/openzipkin/zipkin-go/model"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	conventionsv125 "go.opentelemetry.io/otel/semconv/v1.25.0"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/occonventions"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/internal/zipkin"
)

// ToTranslator converts from Zipkin data model to pdata.
type ToTranslator struct {
	// ParseStringTags should be set to true if tags should be converted to numbers when possible.
	ParseStringTags bool
}

// ToTraces translates Zipkin v2 spans into ptrace.Traces.
func (t ToTranslator) ToTraces(zipkinSpans []*zipkinmodel.SpanModel) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func getResourceSemanticConventionAttributeNames() []string { _ = "STUB: not implemented"; return nil }

var nonSpanAttributes = func() map[string]struct{} {
	attrs := make(map[string]struct{})
	for _, key := range getResourceSemanticConventionAttributeNames() {
		attrs[key] = struct{}{}
	}
	attrs[zipkin.TagServiceNameSource] = struct{}{}
	attrs[string(conventionsv125.OTelLibraryNameKey)] = struct{}{}
	attrs[string(conventionsv125.OTelLibraryVersionKey)] = struct{}{}
	attrs[occonventions.AttributeProcessStartTime] = struct{}{}
	attrs[occonventions.AttributeExporterVersion] = struct{}{}
	attrs[string(conventions.ProcessPIDKey)] = struct{}{}
	attrs[occonventions.AttributeResourceType] = struct{}{}
	return attrs
}()

// Custom Sort on
type byOTLPTypes []*zipkinmodel.SpanModel

func (b byOTLPTypes) Len() int { _ = "STUB: not implemented"; return 0 }

func (b byOTLPTypes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (b byOTLPTypes) Swap(i, j int) { _ = "STUB: not implemented"; return }

func zSpanToInternal(zspan *zipkinmodel.SpanModel, tags map[string]string, dest ptrace.Span, parseStringTags bool) error {
	_ = "STUB: not implemented"
	return nil
}

func populateSpanStatus(tags map[string]string, status ptrace.Status) {
	_ = "STUB: not implemented"
	return
}

func zipkinKindToSpanKind(kind zipkinmodel.Kind, tags map[string]string) ptrace.SpanKind {
	_ = "STUB: not implemented"
	return *new(ptrace.SpanKind)
}

func zTagsToSpanLinks(tags map[string]string, dest ptrace.SpanLinkSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert trace id.

// Convert span id.

func populateSpanEvents(zspan *zipkinmodel.SpanModel, events ptrace.SpanEventSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func jsonMapToAttributeMap(attrs map[string]any, dest pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func zTagsToInternalAttrs(zspan *zipkinmodel.SpanModel, tags map[string]string, dest pcommon.Map, parseStringTags bool) error {
	_ = "STUB: not implemented"
	return nil
}

func tagsToAttributeMap(tags map[string]string, dest pcommon.Map, parseStringTags bool) error {
	_ = "STUB: not implemented"
	return nil
}

func populateResourceFromZipkinSpan(tags map[string]string, localServiceName string, resource pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

func populateILFromZipkinSpan(tags map[string]string, instrLibName string, library pcommon.InstrumentationScope) {
	_ = "STUB: not implemented"
	return
}

func copySpanTags(tags map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func extractLocalServiceName(zspan *zipkinmodel.SpanModel) string {
	_ = "STUB: not implemented"
	return ""
}

func extractInstrumentationLibrary(zspan *zipkinmodel.SpanModel) string {
	_ = "STUB: not implemented"
	return ""
}

func setTimestampsV2(zspan *zipkinmodel.SpanModel, dest ptrace.Span, destAttrs pcommon.Map) {
	_ = "STUB: not implemented"
	// zipkin allows timestamp to be unset, but otel span expects startTimestamp to have a value.
	// unset gets converted to zero on the zspan object during json deserialization because
	// time.Time (the type of Timestamp field) cannot be nil.  If timestamp is zero, the
	// conversion from this internal format back to zipkin format in zipkin exporter fails.
	// Instead, set to *unix* time zero, and convert back in traces_to_zipkinv2.go
	return
}

// unmarshalJSON inflates trace id from hex string, possibly enclosed in quotes.
// TODO: Find a way to avoid this duplicate code. Consider to expose this in pdata.
func unmarshalJSON(dst, src []byte) error { _ = "STUB: not implemented"; return nil }

// TODO: Find a way to avoid this duplicate code. Consider to expose this in pdata.
var statusCodeValue = map[string]int32{
	"STATUS_CODE_UNSET": 0,
	"STATUS_CODE_OK":    1,
	"STATUS_CODE_ERROR": 2,
	// As reported in https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14965
	// The Zipkin exporter used a different set of names when serializing span state.
	"Unset": 0,
	"Ok":    1,
	"Error": 2,
}
