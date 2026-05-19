// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mezmoexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/mezmoexporter"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type mezmoExporter struct {
	config          *Config
	settings        component.TelemetrySettings
	client          *http.Client
	userAgentString string
	log             *zap.Logger
	wg              sync.WaitGroup
}

type mezmoLogLine struct {
	Timestamp int64             `json:"timestamp"`
	Line      string            `json:"line"`
	App       string            `json:"app"`
	Level     string            `json:"level"`
	Meta      map[string]string `json:"meta"`
}

type mezmoLogBody struct {
	Lines []mezmoLogLine `json:"lines"`
}

func newLogsExporter(config *Config, settings component.TelemetrySettings, buildInfo component.BuildInfo, logger *zap.Logger) *mezmoExporter {
	_ = "STUB: not implemented"
	return nil
}

func (m *mezmoExporter) pushLogData(_ context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mezmoExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *mezmoExporter) stop(context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (m *mezmoExporter) logDataToMezmo(ld plog.Logs) error { _ = "STUB: not implemented"; return nil }

// Convert the log resources to mezmo lines...

// Convert Attributes to meta fields being mindful of the maxMetaDataSize restriction

// Send them to Mezmo in batches < 10MB in size

func (m *mezmoExporter) sendLinesToMezmo(post string) (errs error) {
	_ = "STUB: not implemented"
	return nil
}
