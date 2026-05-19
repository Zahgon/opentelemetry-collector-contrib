// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// fileExporter is the implementation of file exporter that writes telemetry data to a file
type fileExporter struct {
	conf       *Config
	marshaller *marshaller
	writer     *fileWriter
}

func (e *fileExporter) consumeTraces(_ context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *fileExporter) consumeMetrics(_ context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *fileExporter) consumeLogs(_ context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *fileExporter) consumeProfiles(_ context.Context, pd pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

// Start starts the flush timer if set.
func (e *fileExporter) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Optionally ensure the output directory exists.

// Shutdown stops the exporter and is invoked during shutdown.
// It stops the flush ticker if set.
func (e *fileExporter) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
