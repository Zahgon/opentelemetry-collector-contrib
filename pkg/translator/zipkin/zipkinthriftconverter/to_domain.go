// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package zipkin // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinthriftconverter"

import (
	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/zipkincore"
	"go.opentelemetry.io/otel/trace"
)

const (
	// UnknownServiceName is serviceName we give to model.Process if we cannot find it anywhere in a Zipkin span
	UnknownServiceName = "unknown-service-name"
	component          = "component"
	peerservice        = "peer.service"
	peerHostIPv4       = "peer.ipv4"
	peerHostIPv6       = "peer.ipv6"
	peerPort           = "peer.port"
)

var (
	coreAnnotations = map[string]string{
		zipkincore.SERVER_RECV: trace.SpanKindServer.String(),
		zipkincore.SERVER_SEND: trace.SpanKindServer.String(),
		zipkincore.CLIENT_RECV: trace.SpanKindClient.String(),
		zipkincore.CLIENT_SEND: trace.SpanKindClient.String(),
	}

	// Some tags on Zipkin spans really describe the process emitting them rather than an individual span.
	// Once all clients are upgraded to use native Jaeger model, this won't be happenning, but for now
	// we remove these tags from the span and store them in the Process.
	processTagAnnotations = map[string]string{
		"jaegerClient":    "jaeger.version", // transform this tag name to client.version
		"jaeger.hostname": "hostname",       // transform this tag name to hostname
		"jaeger.version":  "jaeger.version", // keep this key as is
	}

	trueByteSlice = []byte{1}

	// DefaultLogFieldKey is the log field key which translates directly into Zipkin's Annotation.Value,
	// provided it's the only field in the log.
	// In all other cases the fields are encoded into Annotation.Value as JSON string.
	// TODO move to domain model
	DefaultLogFieldKey = "event"

	// IPTagName is the Jaeger tag name for an IPv4/IPv6 IP address.
	// TODO move to domain model
	IPTagName = "ip"
)

// ToDomain transforms a trace in zipkin.thrift format into model.Trace.
// The transformation assumes that all spans have the same Trace ID.
// A valid model.Trace is always returned, even when there are errors.
// The errors are more of an "fyi", describing issues in the data.
// TODO consider using different return type instead of `error`.
func ToDomain(zSpans []*zipkincore.Span) (*model.Trace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToDomainSpan transforms a span in zipkin.thrift format into model.Span.
// A valid model.Span is always returned, even when there are errors.
// The errors are more of an "fyi", describing issues in the data.
// TODO consider using different return type instead of `error`.
func ToDomainSpan(zSpan *zipkincore.Span) ([]*model.Span, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type toDomain struct{}

func (td toDomain) ToDomain(zSpans []*zipkincore.Span) (*model.Trace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// remove duplicate Process instances

func (td toDomain) ToDomainSpans(zSpan *zipkincore.Span) ([]*model.Span, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (toDomain) findAnnotation(zSpan *zipkincore.Span, value string) *zipkincore.Annotation {
	_ = "STUB: not implemented"
	return nil
}

// transformSpan transforms a zipkin span into a Jaeger span
func (td toDomain) transformSpan(zSpan *zipkincore.Span) []*model.Span {
	_ = "STUB: not implemented"
	return nil
}

// if the span is client and server we split it into two separate spans

// if the first span is a client span we create server span and vice-versa.

// getFlags takes a Zipkin Span and deduces the proper flags settings
func (toDomain) getFlags(zSpan *zipkincore.Span) model.Flags {
	_ = "STUB: not implemented"
	return *new(model.Flags)
}

// Get a correct start time to use for the span if it's not set directly
func (td toDomain) getStartTimeAndDuration(zSpan *zipkincore.Span) (timestamp, duration int64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// generateProcess takes a Zipkin Span and produces a model.Process.
// An optional error may also be returned, but it is not fatal.
func (td toDomain) generateProcess(zSpan *zipkincore.Span) (*model.Process, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the ip process tag already exists, don't add it again

func (td toDomain) findServiceNameAndIP(zSpan *zipkincore.Span) (string, int32, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// If no core annotations exist, use the service name from any annotation

// Tracer can also report a span with just binary annotation/s

func (toDomain) isCoreAnnotation(annotation *zipkincore.Annotation) bool {
	_ = "STUB: not implemented"
	return false
}

func (toDomain) isProcessTag(binaryAnnotation *zipkincore.BinaryAnnotation) bool {
	_ = "STUB: not implemented"
	return false
}

func (td toDomain) isSpanTag(binaryAnnotation *zipkincore.BinaryAnnotation) bool {
	_ = "STUB: not implemented"
	return false
}

type tagPredicate func(*zipkincore.BinaryAnnotation) bool

func (td toDomain) getTags(binAnnotations []*zipkincore.BinaryAnnotation, tagInclude tagPredicate) []model.KeyValue {
	_ = "STUB: not implemented"
	// this will be memory intensive due to how slices work, and it's specifically because we have to filter out
	// some binary annotations. improvement here would be just collecting the indices in binAnnotations we want.
	return nil
}

func (toDomain) transformBinaryAnnotation(binaryAnnotation *zipkincore.BinaryAnnotation) (model.KeyValue, error) {
	_ = "STUB: not implemented"
	return *new(model.KeyValue), nil
}

func bytesToNumber(b []byte, number any) error { _ = "STUB: not implemented"; return nil }

func (td toDomain) getLogs(annotations []*zipkincore.Annotation) []model.Log {
	_ = "STUB: not implemented"
	return nil
}

// If the annotation has no value, throw it out

// skip core annotations

func (toDomain) getLogFields(annotation *zipkincore.Annotation) []model.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// Since Zipkin format does not support kv-logging, some clients encode those Logs
// as annotations with JSON value. Therefore, we try JSON decoding first.

func (toDomain) getSpanKindTag(annotations []*zipkincore.Annotation) (model.KeyValue, bool) {
	_ = "STUB: not implemented"
	return *new(model.KeyValue), false
}

func (toDomain) getPeerTags(endpoint *zipkincore.Endpoint, tags []model.KeyValue) []model.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// Zipkin defines Ipv6 field as: "IPv6 host address packed into 16 bytes. Ex Inet6Address.getBytes()".
// https://github.com/openzipkin/zipkin-api/blob/master/thrift/zipkinCore.thrift#L305
