// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type logsExporter struct {
	db             driver.Conn
	insertSQL      string
	schemaFeatures struct {
		EventName bool
	}

	logger *zap.Logger
	cfg    *Config
}

func newLogsExporter(logger *zap.Logger, cfg *Config) *logsExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsExporter) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	logsColumnEventName = "EventName"
)

func (e *logsExporter) detectSchemaFeatures(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *logsExporter) pushLogsData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// 16 matches the max number of columns in the insert statement.
// If you add or remove columns, update this value.

func (e *logsExporter) renderInsertLogsSQL() error { _ = "STUB: not implemented"; return nil }

// versionFullTextSearch is the minimum ClickHouse version that supports TYPE text() indexes.
var versionFullTextSearch = proto.Version{Major: 26, Minor: 2}

func renderCreateLogsTableSQL(cfg *Config, hasFullTextSearch bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createLogsTable(ctx context.Context, cfg *Config, db driver.Conn, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}
