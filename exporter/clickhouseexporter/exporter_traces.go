// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type tracesExporter struct {
	db        driver.Conn
	insertSQL string

	logger *zap.Logger
	cfg    *Config
}

func newTracesExporter(logger *zap.Logger, cfg *Config) *tracesExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *tracesExporter) pushTraceData(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func convertEvents(events ptrace.SpanEventSlice) (times []time.Time, names []string, attrs []column.IterableOrderedMap) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func convertLinks(links ptrace.SpanLinkSlice) (traceIDs, spanIDs, states []string, attrs []column.IterableOrderedMap) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func renderInsertTracesSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func renderCreateTracesTableSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func renderCreateTraceIDTsTableSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func renderTraceIDTsMaterializedViewSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func createTraceTables(ctx context.Context, cfg *Config, db driver.Conn) error {
	_ = "STUB: not implemented"
	return nil
}
