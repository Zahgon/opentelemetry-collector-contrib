// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cassandraexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/cassandraexporter"

import (
	"context"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type tracesExporter struct {
	client *gocql.Session
	logger *zap.Logger
	cfg    *Config
}

func newTracesExporter(logger *zap.Logger, cfg *Config) *tracesExporter {
	_ = "STUB: not implemented"
	return nil
}

func initializeTraceKernel(cfg *Config) error { _ = "STUB: not implemented"; return nil }

func parseCreateSpanTableSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func parseCreateEventsTypeSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func parseCreateLinksTypeSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func parseCreateDatabaseSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func (e *tracesExporter) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *tracesExporter) pushTraceData(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}
