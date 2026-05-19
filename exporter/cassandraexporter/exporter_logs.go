// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cassandraexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/cassandraexporter"

import (
	"context"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type logsExporter struct {
	client *gocql.Session
	logger *zap.Logger
	cfg    *Config
}

func newLogsExporter(logger *zap.Logger, cfg *Config) *logsExporter {
	_ = "STUB: not implemented"
	return nil
}

func initializeLogKernel(cfg *Config) error { _ = "STUB: not implemented"; return nil }

func newCluster(cfg *Config) (*gocql.ClusterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *logsExporter) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsExporter) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func parseCreateLogTableSQL(cfg *Config) string { _ = "STUB: not implemented"; return "" }

func (e *logsExporter) pushLogsData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}
