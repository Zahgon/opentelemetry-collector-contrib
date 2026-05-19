// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var (
	_ encoding.TracesUnmarshalerExtension  = (*azureExtension)(nil)
	_ encoding.LogsUnmarshalerExtension    = (*azureExtension)(nil)
	_ encoding.MetricsUnmarshalerExtension = (*azureExtension)(nil)
)

type azureExtension struct {
	config            *Config
	logger            *zap.Logger
	logUnmarshaler    plog.Unmarshaler
	traceUnmarshaler  ptrace.Unmarshaler
	metricUnmarshaler pmetric.Unmarshaler
}

func (ex *azureExtension) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (ex *azureExtension) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (ex *azureExtension) UnmarshalMetrics(buf []byte) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (ex *azureExtension) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*azureExtension) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
