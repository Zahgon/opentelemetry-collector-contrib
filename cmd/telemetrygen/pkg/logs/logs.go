// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs

import (
	apilog "go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
)

// Start starts the log telemetry generator
func Start(cfg *Config) error { _ = "STUB: not implemented"; return nil }

// run executes the test scenario.
func run(c *Config, expF exporterFunc, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

type exporterFunc func() (sdklog.Exporter, error)

func exporterFactory(cfg *Config, logger *zap.Logger) exporterFunc {
	_ = "STUB: not implemented"
	return *new(exporterFunc)
}

func createExporter(cfg *Config, logger *zap.Logger) (sdklog.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdklog.Exporter), nil
}

func parseSeverity(severityText string, severityNumber int32) (string, apilog.Severity, error) {
	_ = "STUB: not implemented"
	return "", *new(apilog.Severity), nil
}

// severity number should match well-known severityText
