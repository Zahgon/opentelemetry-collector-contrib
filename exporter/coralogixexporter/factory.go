// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/xexporter"
)

// NewFactory by Coralogix
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Traces GRPC client

// Default to gzip compression

func createTraceExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func createProfilesExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (xexporter.Profiles, error) {
	_ = "STUB: not implemented"
	return *

	// Validate that HTTP protocol is not used with profiles
	new(xexporter.Profiles), nil
}
