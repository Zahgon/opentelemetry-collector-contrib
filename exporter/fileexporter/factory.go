// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/xexporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

const (
	// the number of old log files to retain
	defaultMaxBackups = 100

	// the format of encoded telemetry data
	formatTypeJSON  = "json"
	formatTypeProto = "proto"

	// the type of compression codec
	compressionZSTD = "zstd"

	defaultMaxOpenFiles = 100

	defaultResourceAttribute = "fileexporter.path_segment"
)

type FileExporter interface {
	component.Component
	consumeTraces(_ context.Context, td ptrace.Traces) error
	consumeMetrics(_ context.Context, md pmetric.Metrics) error
	consumeLogs(_ context.Context, ld plog.Logs) error
	consumeProfiles(_ context.Context, pd pprofile.Profiles) error
}

// NewFactory creates a factory for OTLP exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func createProfilesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (xexporter.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xexporter.Profiles), nil
}

// getOrCreateFileExporter creates a FileExporter and caches it for a particular configuration,
// or returns the already cached one. Caching is required because the factory is asked trace and
// metric receivers separately when it gets CreateTraces() and CreateMetrics()
// but they must not create separate objects, they must use one Exporter object per configuration.
func getOrCreateFileExporter(cfg component.Config, logger *zap.Logger) FileExporter {
	_ = "STUB: not implemented"
	return *new(FileExporter)
}

func newFileExporter(conf *Config, logger *zap.Logger) FileExporter {
	_ = "STUB: not implemented"
	return *new(FileExporter)
}

func newFileWriter(path string, shouldAppend bool, rotation *Rotation, flushInterval time.Duration, export exportFunc) (*fileWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is the map of already created File exporters for particular configurations.
// We maintain this map because the Factory is asked trace and metric receivers separately
// when it gets CreateTraces() and CreateMetrics() but they must not
// create separate objects, they must use one Exporter object per configuration.
var exporters = sharedcomponent.NewSharedComponents()
