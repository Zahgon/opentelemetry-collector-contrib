// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelserializer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer"

import (
	"bytes"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
)

func (*Serializer) SerializeSpanEvent(resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, span ptrace.Span, spanEvent ptrace.SpanEvent, idx elasticsearch.Index, buf *bytes.Buffer) {
	_ = "STUB: not implemented"
	return
}

func (*Serializer) SerializeSpan(resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, span ptrace.Span, idx elasticsearch.Index, buf *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func writeStatus(w *jsonWriter, status ptrace.Status, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

func writeSpanLinks(w *jsonWriter, span ptrace.Span, first bool) bool {
	_ = "STUB: not implemented"
	return false
}
