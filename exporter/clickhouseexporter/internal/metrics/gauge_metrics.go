// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal/metrics"

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type gaugeModel struct {
	metricName        string
	metricDescription string
	metricUnit        string
	metadata          *MetricsMetaData
	gauge             pmetric.Gauge
}

type gaugeMetrics struct {
	gaugeModels []*gaugeModel
	insertSQL   string
	count       int
}

func (g *gaugeMetrics) insert(ctx context.Context, db driver.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *gaugeMetrics) Add(resAttr pcommon.Map, resURL string, scopeInstr pcommon.InstrumentationScope, scopeURL string, metrics pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}
