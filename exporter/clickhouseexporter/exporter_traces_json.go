// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

// anyTracesExporter is an interface that satisfies both the default map tracesExporter and the tracesJSONExporter
type anyTracesExporter interface {
	start(context.Context, component.Host) error
	shutdown(context.Context) error
	pushTraceData(ctx context.Context, td ptrace.Traces) error
}

type tracesJSONExporter struct {
	cfg            *Config
	logger         *zap.Logger
	db             driver.Conn
	insertSQL      string
	schemaFeatures struct {
		AttributeKeys bool
	}
}

func newTracesJSONExporter(logger *zap.Logger, cfg *Config) *tracesJSONExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesJSONExporter) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	tracesJSONColumnResourceAttributesKeys = "ResourceAttributesKeys"
	tracesJSONColumnSpanAttributesKeys     = "SpanAttributesKeys"
)

func (e *tracesJSONExporter) detectSchemaFeatures(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesJSONExporter) shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesJSONExporter) pushTraceData(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func convertEventsJSON(events ptrace.SpanEventSlice) (times []time.Time, names, attrs []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func convertLinksJSON(links ptrace.SpanLinkSlice) (traceIDs, spanIDs, states, attrs []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func (e *tracesJSONExporter) renderInsertTracesJSONSQL() { _ = "STUB: not implemented"; return }

func renderCreateTracesJSONTableSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func createTraceJSONTables(ctx context.Context, cfg *Config, db driver.Conn) error {
	_ = "STUB: not implemented"
	return nil
}
