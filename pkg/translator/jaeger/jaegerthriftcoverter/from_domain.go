// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger/jaegerthriftcoverter"

import (
	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/jaeger"
)

// FromDomain takes an array of model.Span and returns
// an array of jaeger.Span.  If errors are found during
// conversion of tags, then error tags are appended.
func FromDomain(spans []*model.Span) []*jaeger.Span { _ = "STUB: not implemented"; return nil }

// FromDomainSpan takes a single model.Span and
// converts it into a jaeger.Span.  If errors are found
// during conversion of tags, then error tags are appended.
func FromDomainSpan(span *model.Span) *jaeger.Span { _ = "STUB: not implemented"; return nil }

type domainToJaegerTransformer struct{}

func (domainToJaegerTransformer) keyValueToTag(kv *model.KeyValue) *jaeger.Tag {
	_ = "STUB: not implemented"
	return nil
}

func (d domainToJaegerTransformer) convertKeyValuesToTags(kvs model.KeyValues) []*jaeger.Tag {
	_ = "STUB: not implemented"
	return nil
}

func (d domainToJaegerTransformer) convertLogs(logs []model.Log) []*jaeger.Log {
	_ = "STUB: not implemented"
	return nil
}

func (domainToJaegerTransformer) convertSpanRefs(refs []model.SpanRef) []*jaeger.SpanRef {
	_ = "STUB: not implemented"
	return nil
}

func (d domainToJaegerTransformer) transformSpan(span *model.Span) *jaeger.Span {
	_ = "STUB: not implemented"
	return nil
}
