// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/remotetapprocessor"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"golang.org/x/net/websocket"
	"golang.org/x/time/rate"
)

type wsprocessor struct {
	config            *Config
	telemetrySettings component.TelemetrySettings
	server            *http.Server
	shutdownWG        sync.WaitGroup
	cs                *channelSet
	limiter           *rate.Limiter
}

var (
	logMarshaler    = &plog.JSONMarshaler{}
	metricMarshaler = &pmetric.JSONMarshaler{}
	traceMarshaler  = &ptrace.JSONMarshaler{}
)

func newProcessor(settings processor.Settings, config *Config) *wsprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (w *wsprocessor) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wsprocessor) handleConn(conn *websocket.Conn) { _ = "STUB: not implemented"; return }

func (w *wsprocessor) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// The processor's channelset is only modified by its server, so once
// it's completely shutdown it's safe to shutdown the channelset itself.

func (w *wsprocessor) ConsumeMetrics(_ context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (w *wsprocessor) ConsumeLogs(_ context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (w *wsprocessor) ConsumeTraces(_ context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
