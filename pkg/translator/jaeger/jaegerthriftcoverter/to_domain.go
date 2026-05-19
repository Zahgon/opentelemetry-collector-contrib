// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger/jaegerthriftcoverter"

import (
	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/jaeger"
)

// ToDomain transforms a set of spans and a process in jaeger.thrift format into a slice of model.Span.
// A valid []*model.Span is always returned, even when there are errors.
// Errors are presented as tags on spans
func ToDomain(jSpans []*jaeger.Span, jProcess *jaeger.Process) []*model.Span {
	_ = "STUB: not implemented"
	return nil
}

// ToDomainSpan transforms a span in jaeger.thrift format into model.Span.
// A valid model.Span is always returned, even when there are errors.
// Errors are presented as tags on spans
func ToDomainSpan(jSpan *jaeger.Span, jProcess *jaeger.Process) *model.Span {
	_ = "STUB: not implemented"
	return nil
}

// ToDomainProcess transforms a process in jaeger.thrift format to model.Span.
func ToDomainProcess(jProcess *jaeger.Process) *model.Process {
	_ = "STUB: not implemented"
	return nil
}

type toDomain struct{}

func (td toDomain) ToDomain(jSpans []*jaeger.Span, jProcess *jaeger.Process) []*model.Span {
	_ = "STUB: not implemented"
	return nil
}

func (td toDomain) ToDomainSpan(jSpan *jaeger.Span, jProcess *jaeger.Process) *model.Span {
	_ = "STUB: not implemented"
	return nil
}

func (td toDomain) transformSpan(jSpan *jaeger.Span, mProcess *model.Process) *model.Span {
	_ = "STUB: not implemented"
	return nil
}

// allocate extra space for future append operation

// We no longer store ParentSpanID in the domain model, but the data in Thrift model
// might still have these IDs without representing them in the References, so we
// convert it back into child-of reference.

func (toDomain) getReferences(jRefs []*jaeger.SpanRef) []model.SpanRef {
	_ = "STUB: not implemented"
	return nil
}

// getProcess takes a jaeger.thrift process and produces a model.Process.
// Any errors are presented as tags
func (td toDomain) getProcess(jProcess *jaeger.Process) *model.Process {
	_ = "STUB: not implemented"
	return nil
}

// convert the jaeger.Tag slice to domain KeyValue slice
// zipkin/to_domain.go does not give a default slice size since it has to filter annotations, jaeger conversion is more predictable
// thus to avoid future full array copy when using append, pre-allocate extra space as an optimization
func (td toDomain) getTags(tags []*jaeger.Tag, extraSpace int) model.KeyValues {
	_ = "STUB: not implemented"
	return *new(model.KeyValues)
}

func (toDomain) getTag(tag *jaeger.Tag) model.KeyValue {
	_ = "STUB: not implemented"
	return *new(model.KeyValue)
}

func (td toDomain) getLogs(logs []*jaeger.Log) []model.Log { _ = "STUB: not implemented"; return nil }
