// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type pdataLogsMarshaler struct {
	marshaler plog.Marshaler
	encoding  string
}

func (p pdataLogsMarshaler) Marshal(ld plog.Logs, _ string) ([]*pulsar.ProducerMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p pdataLogsMarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

func newPdataLogsMarshaler(marshaler plog.Marshaler, encoding string) LogsMarshaler {
	_ = "STUB: not implemented"
	return *new(LogsMarshaler)
}

type pdataMetricsMarshaler struct {
	marshaler pmetric.Marshaler
	encoding  string
}

func (p pdataMetricsMarshaler) Marshal(ld pmetric.Metrics, _ string) ([]*pulsar.ProducerMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p pdataMetricsMarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

func newPdataMetricsMarshaler(marshaler pmetric.Marshaler, encoding string) MetricsMarshaler {
	_ = "STUB: not implemented"
	return *new(MetricsMarshaler)
}

type pdataTracesMarshaler struct {
	marshaler ptrace.Marshaler
	encoding  string
}

func (p pdataTracesMarshaler) Marshal(td ptrace.Traces, _ string) ([]*pulsar.ProducerMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p pdataTracesMarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

func newPdataTracesMarshaler(marshaler ptrace.Marshaler, encoding string) TracesMarshaler {
	_ = "STUB: not implemented"
	return *new(TracesMarshaler)
}
