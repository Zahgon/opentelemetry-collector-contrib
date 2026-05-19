// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type pubsubExporter struct {
	logger               *zap.Logger
	client               publisherClient
	cancel               context.CancelFunc
	userAgent            string
	ceSource             string
	ceCompression        compression
	config               *Config
	tracesMarshaler      ptrace.Marshaler
	tracesAttributes     map[string]string
	tracesWatermarkFunc  tracesWatermarkFunc
	metricsMarshaler     pmetric.Marshaler
	metricsAttributes    map[string]string
	metricsWatermarkFunc metricsWatermarkFunc
	logsMarshaler        plog.Marshaler
	logsAttributes       map[string]string
	logsWatermarkFunc    logsWatermarkFunc

	// To be overridden in tests
	makeUUID   func() (uuid.UUID, error)
	makeClient func(ctx context.Context, cfg *Config, userAgent string) (publisherClient, error)
}

type signal int

const (
	signalTrace  signal = iota
	signalMetric        = iota
	signalLog           = iota
)

type compression int

const (
	uncompressed compression = iota
	gZip                     = iota
)

type WatermarkBehavior int

const (
	current  WatermarkBehavior = iota
	earliest                   = iota
)

func (ex *pubsubExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (ex *pubsubExporter) initTraces(host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) initMetrics(host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) initLogs(host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) setTracesMarshalerFromExtension(host component.Host, extensionID component.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) setMetricsMarshalerFromExtension(host component.Host, extensionID component.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) setLogsMarshalerFromExtension(host component.Host, extensionID component.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) getMessageAttributes(signal signal, watermark time.Time) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ex *pubsubExporter) consumeTraces(ctx context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// No ordering key

func (ex *pubsubExporter) publishTraces(ctx context.Context, tracesForKey ptrace.Traces, orderingKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) consumeMetrics(ctx context.Context, metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// No ordering key

func (ex *pubsubExporter) publishMetrics(ctx context.Context, metricsForKey pmetric.Metrics, orderingKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) consumeLogs(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// No ordering key

func (ex *pubsubExporter) publishLogs(ctx context.Context, logs plog.Logs, orderingKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) publishMessage(ctx context.Context, data []byte, attributes map[string]string, orderingKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ex *pubsubExporter) compress(payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
