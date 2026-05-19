// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

// anyLogsExporter is an interface that satisfies both the default map logsExporter and the logsJSONExporter
type anyLogsExporter interface {
	start(context.Context, component.Host) error
	shutdown(context.Context) error
	pushLogsData(context.Context, plog.Logs) error
}

type logsJSONExporter struct {
	cfg            *Config
	logger         *zap.Logger
	db             driver.Conn
	insertSQL      string
	schemaFeatures struct {
		AttributeKeys bool
		EventName     bool
	}
}

func newLogsJSONExporter(logger *zap.Logger, cfg *Config) *logsJSONExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsJSONExporter) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	logsJSONColumnResourceAttributesKeys = "ResourceAttributesKeys"
	logsJSONColumnScopeAttributesKeys    = "ScopeAttributesKeys"
	logsJSONColumnLogAttributesKeys      = "LogAttributesKeys"
	logsJSONColumnEventName              = "EventName"
)

func (e *logsJSONExporter) detectSchemaFeatures(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsJSONExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *logsJSONExporter) pushLogsData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsJSONExporter) renderInsertLogsJSONSQL() error { _ = "STUB: not implemented"; return nil }

func renderCreateLogsJSONTableSQL(cfg *Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createLogsJSONTable(ctx context.Context, cfg *Config, db driver.Conn) error {
	_ = "STUB: not implemented"
	return nil
}
