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
type pdataLogsUnmarshaler struct {
	plog.Unmarshaler
	encoding string
}

func (p pdataLogsUnmarshaler) Unmarshal(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (p pdataLogsUnmarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

func newPdataLogsUnmarshaler(unmarshaler plog.Unmarshaler, encoding string) LogsUnmarshaler {
	_ = "STUB: not implemented"
	return *new(LogsUnmarshaler)
}

type pdataTracesUnmarshaler struct {
	ptrace.Unmarshaler
	encoding string
}

func (p pdataTracesUnmarshaler) Unmarshal(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (p pdataTracesUnmarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

func newPdataTracesUnmarshaler(unmarshaler ptrace.Unmarshaler, encoding string) TracesUnmarshaler {
	_ = "STUB: not implemented"
	return *new(TracesUnmarshaler)
}

type pdataMetricsUnmarshaler struct {
	pmetric.Unmarshaler
	encoding string
}

func (p pdataMetricsUnmarshaler) Unmarshal(buf []byte) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (p pdataMetricsUnmarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

func newPdataMetricsUnmarshaler(unmarshaler pmetric.Unmarshaler, encoding string) MetricsUnmarshaler {
	_ = "STUB: not implemented"
	return *new(MetricsUnmarshaler)
}
