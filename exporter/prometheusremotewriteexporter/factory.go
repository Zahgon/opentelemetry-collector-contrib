// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a new Prometheus Remote Write exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createMetricsExporter(ctx context.Context, set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// We almost read 0 bytes, so no need to tune ReadBufferSize.

// To set this as default once `exporter.prometheusremotewritexporter.EnableMultipleWorkers` is removed
// MaxBatchRequestParallelism: 5,

// TODO: Set TranslationStrategy to UnderscoreEscapingWithSuffixes once AddMetricSuffixes is removed.

// TODO(jbd): Adjust the default queue size.
