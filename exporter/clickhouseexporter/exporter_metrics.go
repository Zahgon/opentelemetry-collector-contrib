// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal/metrics"
)

type metricsExporter struct {
	db driver.Conn

	logger       *zap.Logger
	cfg          *Config
	tablesConfig metrics.MetricTablesConfigMapper
}

func newMetricsExporter(logger *zap.Logger, cfg *Config) *metricsExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func generateMetricTablesConfigMapper(cfg *Config) metrics.MetricTablesConfigMapper {
	_ = "STUB: not implemented"
	return *new(metrics.MetricTablesConfigMapper)
}

// shutdown will shut down the exporter.
func (e *metricsExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *metricsExporter) pushMetricsData(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}
