// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudflarereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudflarereceiver"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	rcvr "go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
)

type logsReceiver struct {
	logger            *zap.Logger
	cfg               *LogsConfig
	serverConfig      *confighttp.ServerConfig
	server            *http.Server
	consumer          consumer.Logs
	wg                *sync.WaitGroup
	telemetrySettings component.TelemetrySettings
	id                component.ID // ID of the receiver component
	obsrecv           *receiverhelper.ObsReport
}

const secretHeaderName = "X-CF-Secret"

func newLogsReceiver(params rcvr.Settings, cfg *Config, consumer consumer.Logs) (*logsReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *logsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *logsReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (l *logsReceiver) startListening(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *logsReceiver) handleRequest(rw http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Limit request body size

// Read the decompressed response body

func parsePayload(payload []byte) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *logsReceiver) processLogs(now pcommon.Timestamp, logs []map[string]any) plog.Logs {
	_ = "STUB: not implemented"
	return *

	// Group logs by ZoneName field if it was configured so it can be used as a resource attribute
	new(plog.Logs)
}

// Only process fields that are in the config mapping

// Skip fields not in mapping when we have a config

// else if l.cfg.Attributes is empty, default to processing all fields with no renaming

// Flatten the map and add each field with a prefixed key

// severityFromStatusCode translates HTTP status code to OpenTelemetry severity number.
func severityFromStatusCode(statusCode int64) plog.SeverityNumber {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber)
}

// flattenMap recursively flattens a map[string]any into a single level map
// with keys joined by the specified separator
func flattenMap(input map[string]any, prefix, separator string, result map[string]any) {
	_ = "STUB: not implemented"
	return
}

// Replace hyphens with underscores in the key. Content-Type becomes Content_Type

// Recursively flatten nested maps
