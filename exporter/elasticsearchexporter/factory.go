// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/exporterhelper/xexporterhelper"
	"go.opentelemetry.io/collector/exporter/xexporter"
)

// NewFactory creates a factory for Elastic exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// default is set in exporter code

// createLogsExporter creates a new exporter for logs.
//
// Logs are directly indexed into Elasticsearch.
func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createTracesExporter(ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// createProfilesExporter creates a new exporter for profiles.
//
// Profiles are directly indexed into Elasticsearch.
func createProfilesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (xexporter.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xexporter.Profiles), nil
}

func exporterhelperOptions(
	cfg *Config,
	start component.StartFunc,
	shutdown component.ShutdownFunc,
	qbs xexporterhelper.QueueBatchSettings,
) []exporterhelper.Option {
	_ = "STUB: not implemented"
	// not setting capabilities as they will default to non-mutating but will be updated
	// by the base-exporter to mutating if batching is enabled.
	return nil
}

// Effectively disable timeout_sender because timeout is enforced in bulk indexer.
