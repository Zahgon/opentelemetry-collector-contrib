// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logicmonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter"

import (
	"context"
	"net/http"

	lmsdklogs "github.com/logicmonitor/lm-data-sdk-go/api/logs"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	logs "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter/internal/logs"
)

// These are logicmonitor specific constants needed to map the resource with the logs on logicmonitor platform.
const (
	hostname         = "hostname"
	hostnameProperty = "system.hostname"
)

type logExporter struct {
	config   *Config
	sender   *logs.Sender
	settings component.TelemetrySettings
	cancel   context.CancelFunc
}

// Create new logicmonitor logs exporter
func newLogsExporter(_ context.Context, cfg component.Config, set exporter.Settings) *logExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *logExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logExporter) PushLogData(ctx context.Context, lg plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func buildLogIngestOpts(config *Config, client *http.Client) []lmsdklogs.Option {
	_ = "STUB: not implemented"
	return nil
}

func timestampFromLogRecord(lr plog.LogRecord) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}
