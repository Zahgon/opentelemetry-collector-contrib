// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// copy from kafka receiver

// TracesUnmarshaler deserializes the message body.
type TracesUnmarshaler interface {
	// Unmarshal deserializes the message body into tracesConsumer.
	Unmarshal([]byte) (ptrace.Traces, error)

	// Encoding of the serialized messages.
	Encoding() string
}

// MetricsUnmarshaler deserializes the message body
type MetricsUnmarshaler interface {
	// Unmarshal deserializes the message body into tracesConsumer
	Unmarshal([]byte) (pmetric.Metrics, error)

	// Encoding of the serialized messages
	Encoding() string
}

// LogsUnmarshaler deserializes the message body.
type LogsUnmarshaler interface {
	// Unmarshal deserializes the message body into tracesConsumer.
	Unmarshal([]byte) (plog.Logs, error)

	// Encoding of the serialized messages.
	Encoding() string
}

// defaultTracesUnmarshalers returns map of supported encodings with TracesUnmarshaler.
func defaultTracesUnmarshalers() map[string]TracesUnmarshaler {
	_ = "STUB: not implemented"
	return nil
}

func defaultMetricsUnmarshalers() map[string]MetricsUnmarshaler {
	_ = "STUB: not implemented"
	return nil
}

func defaultLogsUnmarshalers() map[string]LogsUnmarshaler { _ = "STUB: not implemented"; return nil }
