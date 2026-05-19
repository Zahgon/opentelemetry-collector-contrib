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

// TracesMarshaler marshals traces into Message array.
type TracesMarshaler interface {
	// Marshal serializes spans into sarama's ProducerMessages
	Marshal(traces ptrace.Traces, topic string) ([]*pulsar.ProducerMessage, error)

	// Encoding returns encoding name
	Encoding() string
}

// MetricsMarshaler marshals metrics into Message array
type MetricsMarshaler interface {
	// Marshal serializes metrics into sarama's ProducerMessages
	Marshal(metrics pmetric.Metrics, topic string) ([]*pulsar.ProducerMessage, error)

	// Encoding returns encoding name
	Encoding() string
}

// LogsMarshaler marshals logs into Message array
type LogsMarshaler interface {
	// Marshal serializes logs into sarama's ProducerMessages
	Marshal(logs plog.Logs, topic string) ([]*pulsar.ProducerMessage, error)

	// Encoding returns encoding name
	Encoding() string
}

// tracesMarshalers returns map of supported encodings with TracesMarshaler.
func tracesMarshalers() map[string]TracesMarshaler { _ = "STUB: not implemented"; return nil }

// metricsMarshalers returns map of supported encodings and MetricsMarshaler
func metricsMarshalers() map[string]MetricsMarshaler { _ = "STUB: not implemented"; return nil }

// logsMarshalers returns map of supported encodings and LogsMarshaler
func logsMarshalers() map[string]LogsMarshaler { _ = "STUB: not implemented"; return nil }
