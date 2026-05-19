// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"

import (
	"context"
	"sync"

	"github.com/hashicorp/golang-lru/v2/simplelru"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type groupingFileExporter struct {
	conf          *Config
	logger        *zap.Logger
	marshaller    *marshaller
	pathPrefix    string
	pathSuffix    string
	attribute     string
	maxOpenFiles  int
	newFileWriter func(path string) (*fileWriter, error)

	mutex   sync.Mutex
	writers *simplelru.LRU[string, *fileWriter]
}

func (e *groupingFileExporter) consumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *groupingFileExporter) consumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *groupingFileExporter) consumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *groupingFileExporter) consumeProfiles(ctx context.Context, pd pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *groupingFileExporter) write(_ context.Context, pathSegment string, buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *groupingFileExporter) getWriter(pathSegment string) (*fileWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cleanPathPrefix(pathPrefix string) string { _ = "STUB: not implemented"; return "" }

func (e *groupingFileExporter) fullPath(pathSegment string) string {
	_ = "STUB: not implemented"
	return ""
}

// avoid path traversal vulnerability

func (e *groupingFileExporter) onEvict(_ string, writer *fileWriter) {
	_ = "STUB: not implemented"
	return
}

func group[T any](e *groupingFileExporter, groups map[string][]T, resource pcommon.Resource, resourceEntries T) {
	_ = "STUB: not implemented"
	return
}

// Start initializes and starts the exporter.
func (e *groupingFileExporter) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown stops the exporter and is invoked during shutdown.
// It stops flushes and closes all underlying writers.
func (e *groupingFileExporter) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
