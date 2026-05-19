// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gogo/protobuf/jsonpb"
	jaegerproto "github.com/jaegertracing/jaeger-idl/model/v1"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type jaegerMarshaler struct {
	marshaler jaegerBatchMarshaler
}

var _ TracesMarshaler = (*jaegerMarshaler)(nil)

func (j jaegerMarshaler) Marshal(traces ptrace.Traces, _ string) ([]*pulsar.ProducerMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jaegerMarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

type jaegerBatchMarshaler interface {
	marshal(batch *jaegerproto.Batch) ([]byte, error)
	encoding() string
}

type jaegerProtoBatchMarshaler struct{}

var _ jaegerBatchMarshaler = (*jaegerProtoBatchMarshaler)(nil)

func (jaegerProtoBatchMarshaler) marshal(batch *jaegerproto.Batch) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jaegerProtoBatchMarshaler) encoding() string { _ = "STUB: not implemented"; return "" }

type jaegerJSONBatchMarshaler struct {
	pbMarshaler *jsonpb.Marshaler
}

var _ jaegerBatchMarshaler = (*jaegerJSONBatchMarshaler)(nil)

func newJaegerJSONMarshaler() *jaegerJSONBatchMarshaler { _ = "STUB: not implemented"; return nil }

func (p jaegerJSONBatchMarshaler) marshal(batch *jaegerproto.Batch) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jaegerJSONBatchMarshaler) encoding() string { _ = "STUB: not implemented"; return "" }
