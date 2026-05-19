// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal/metrics"

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal/sqltemplates"
)

var supportedMetricTypes = map[pmetric.MetricType]string{
	pmetric.MetricTypeGauge:                sqltemplates.MetricsGaugeCreateTable,
	pmetric.MetricTypeSum:                  sqltemplates.MetricsSumCreateTable,
	pmetric.MetricTypeHistogram:            sqltemplates.MetricsHistogramCreateTable,
	pmetric.MetricTypeExponentialHistogram: sqltemplates.MetricsExpHistogramCreateTable,
	pmetric.MetricTypeSummary:              sqltemplates.MetricsSummaryCreateTable,
}

var logger *zap.Logger

type MetricTablesConfigMapper map[pmetric.MetricType]MetricTypeConfig

type MetricTypeConfig struct {
	Name string `mapstructure:"name"`
}

// MetricsModel is used to group metric data and insert into clickhouse
// any type of metrics need implement it.
type MetricsModel interface {
	// Add used to bind MetricsMetaData to a specific metric then put them into a slice
	Add(resAttr pcommon.Map, resURL string, scopeInstr pcommon.InstrumentationScope, scopeURL string, metrics pmetric.Metric)

	// insert is used to insert metric data to clickhouse
	insert(ctx context.Context, db driver.Conn) error
}

// MetricsMetaData contain specific metric data
type MetricsMetaData struct {
	ResAttr    pcommon.Map
	ResURL     string
	ScopeURL   string
	ScopeInstr pcommon.InstrumentationScope
}

// SetLogger set a logger instance
func SetLogger(l *zap.Logger) {
	_ = "STUB: not implemented"

	// NewMetricsTable create metric tables with an expiry time to storage metric telemetry data
	return
}

func NewMetricsTable(ctx context.Context, tablesConfig MetricTablesConfigMapper, database, cluster, engine, ttlExpr string, db driver.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMetricsModel create a model for contain different metric data
func NewMetricsModel(tablesConfig MetricTablesConfigMapper, database string) map[pmetric.MetricType]MetricsModel {
	_ = "STUB: not implemented"
	return nil
}

// InsertMetrics insert metric data into clickhouse concurrently
func InsertMetrics(ctx context.Context, db driver.Conn, metricsMap map[pmetric.MetricType]MetricsModel) error {
	_ = "STUB: not implemented"
	return nil
}

func convertExemplars(exemplars pmetric.ExemplarSlice) (clickhouse.ArraySet, clickhouse.ArraySet, clickhouse.ArraySet, clickhouse.ArraySet, clickhouse.ArraySet) {
	_ = "STUB: not implemented"
	return *new(clickhouse.ArraySet), *new(clickhouse.ArraySet), *new(clickhouse.ArraySet), *new(clickhouse.ArraySet), *new(clickhouse.ArraySet)
}

// https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/metrics/v1/metrics.proto#L358
// define two types for one datapoint value, clickhouse only use one value of float64 to store them
func getValue(intValue int64, floatValue float64, dataType any) float64 {
	_ = "STUB: not implemented"
	return 0
}

func AttributesToMap(attributes pcommon.Map) column.IterableOrderedMap {
	_ = "STUB: not implemented"
	return *new(column.IterableOrderedMap)
}

func GetServiceName(resAttr pcommon.Map) string { _ = "STUB: not implemented"; return "" }

func convertSliceToArraySet[T any](slice []T) clickhouse.ArraySet {
	_ = "STUB: not implemented"
	return *new(clickhouse.ArraySet)
}

func convertValueAtQuantile(valueAtQuantile pmetric.SummaryDataPointValueAtQuantileSlice) (clickhouse.ArraySet, clickhouse.ArraySet) {
	_ = "STUB: not implemented"
	return *new(clickhouse.ArraySet), *new(clickhouse.ArraySet)
}
