// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package sematextexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sematextexporter"

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

var (
	metricsAppToken = uuid.NewString()
	logsAppToken    = uuid.NewString()
)

// NewFactory creates a factory for the Sematext metrics exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *

	// Initialize the logger for Sematext
	new(exporter.Metrics), nil
}

// Create a writer for sending metrics to Sematext

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
