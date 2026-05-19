// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk"
)

const (
	defaultMaxIdleCons          = 100
	defaultHTTPTimeout          = 10 * time.Second
	defaultHTTP2ReadIdleTimeout = time.Second * 10
	defaultHTTP2PingTimeout     = time.Second * 10
	defaultIdleConnTimeout      = 10 * time.Second
	defaultSplunkAppName        = "OpenTelemetry Collector Contrib"
)

// TODO: Find a place for this to be shared.
type baseMetricsExporter struct {
	component.Component
	consumer.Metrics
}

// TODO: Find a place for this to be shared.
type baseLogsExporter struct {
	component.Component
	consumer.Logs
}

// TODO: Find a place for this to be shared.
type baseTracesExporter struct {
	component.Component
	consumer.Traces
}

// NewFactory creates a factory for Splunk HEC exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// explicitly disable since we rely on http.Client timeout logic.

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

// explicitly disable since we rely on http.Client timeout logic.

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter exporter.Logs, err error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// explicitly disable since we rely on http.Client timeout logic.

// hecRequiredMetadataKeys are the context metadata keys the exporterhelper
// batcher must use to partition requests so different HEC routing targets
// are never merged into the same batch.
var hecRequiredMetadataKeys = []string{splunk.HecTokenLabel, splunk.DefaultIndexLabel}

// hecQueueSettings returns a copy of qs with hecRequiredMetadataKeys
// guaranteed to be present in Batch.Partition.MetadataKeys. Keys already
// set by the user are preserved; required keys are appended only when absent
// (case-insensitive). Returns qs unchanged when batching is not configured.
func hecQueueSettings(qs configoptional.Optional[exporterhelper.QueueBatchConfig]) configoptional.Optional[exporterhelper.QueueBatchConfig] {
	_ = "STUB: not implemented"
	return nil
}
