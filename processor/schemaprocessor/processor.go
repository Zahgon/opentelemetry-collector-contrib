// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package schemaprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/translation"
)

type schemaProcessor struct {
	telemetry   component.TelemetrySettings
	componentID component.ID
	config      *Config

	log *zap.Logger

	manager           translation.Manager
	telemetryBuilder  *metadata.TelemetryBuilder
	storageClient     storage.Client
	migrationFromURLs map[string]string // target schema URL → migration from URL (for metrics)
}

func newSchemaProcessor(_ context.Context, conf component.Config, set processor.Settings) (*schemaProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t schemaProcessor) recordTranslation(ctx context.Context, fromSchemaURL, toSchemaURL string) {
	_ = "STUB: not implemented"
	return
}

func (t schemaProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (t schemaProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (t schemaProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// start will add HTTP provider to the manager and prefetch schemas
func (t *schemaProcessor) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *schemaProcessor) shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func getStorageClient(ctx context.Context, host component.Host, storageID, componentID component.ID) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}
