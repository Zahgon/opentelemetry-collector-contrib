// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter"

import (
	"context"
	"crypto/tls"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type syslogexporter struct {
	config    *Config
	logger    *zap.Logger
	tlsConfig *tls.Config
	formatter formatter
}

func initExporter(cfg *Config, createSettings exporter.Settings) (*syslogexporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newLogsExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg *Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func (se *syslogexporter) pushLogsData(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (se *syslogexporter) exportBatch(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (se *syslogexporter) exportNonBatch(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}
