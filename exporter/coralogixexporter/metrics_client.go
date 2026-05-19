// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
)

func newMetricsExporter(cfg component.Config, set exporter.Settings) (*metricsExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type metricsExporter struct {
	grpcMetricsExporter pmetricotlp.GRPCClient
	httpMetricsExporter httpMetricsExporter
	*signalExporter
}

func (e *metricsExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) enhanceContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
