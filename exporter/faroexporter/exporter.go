// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/faroexporter"

import (
	"context"
	"net/http"

	faro "github.com/grafana/faro/pkg/go"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type faroExporter struct {
	config    *Config
	client    *http.Client
	logger    *zap.Logger
	settings  component.TelemetrySettings
	userAgent string
}

const (
	headerRetryAfter = "Retry-After"
	jsonContentType  = "application/json"
)

func newExporter(cfg component.Config, set exporter.Settings) (*faroExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fe *faroExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (fe *faroExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (fe *faroExporter) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (fe *faroExporter) export(ctx context.Context, fp *faro.Payload) error {
	_ = "STUB: not implemented"
	return nil
}

func (fe *faroExporter) consume(ctx context.Context, fp []faro.Payload) error {
	_ = "STUB: not implemented"
	return nil
}

func (*faroExporter) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}
