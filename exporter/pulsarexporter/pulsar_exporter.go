// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"

import (
	"context"
	"errors"

	"github.com/apache/pulsar-client-go/pulsar"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

var errUnrecognizedEncoding = errors.New("unrecognized encoding")

type PulsarTracesProducer struct {
	cfg       Config
	client    pulsar.Client
	producer  pulsar.Producer
	topic     string
	marshaler TracesMarshaler
	logger    *zap.Logger
}

func (e *PulsarTracesProducer) tracesPusher(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *PulsarTracesProducer) Close(context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *PulsarTracesProducer) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

type PulsarMetricsProducer struct {
	cfg       Config
	client    pulsar.Client
	producer  pulsar.Producer
	topic     string
	marshaler MetricsMarshaler
	logger    *zap.Logger
}

func (e *PulsarMetricsProducer) metricsDataPusher(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *PulsarMetricsProducer) Close(context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *PulsarMetricsProducer) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

type PulsarLogsProducer struct {
	cfg       Config
	client    pulsar.Client
	producer  pulsar.Producer
	topic     string
	marshaler LogsMarshaler
	logger    *zap.Logger
}

func (e *PulsarLogsProducer) logsDataPusher(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *PulsarLogsProducer) Close(context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *PulsarLogsProducer) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func newPulsarProducer(config Config) (pulsar.Client, pulsar.Producer, error) {
	_ = "STUB: not implemented"
	return *new(pulsar.Client), *new(pulsar.Producer), nil
}

func newMetricsExporter(config Config, set exporter.Settings, marshalers map[string]MetricsMarshaler) (*PulsarMetricsProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTracesExporter(config Config, set exporter.Settings, marshalers map[string]TracesMarshaler) (*PulsarTracesProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newLogsExporter(config Config, set exporter.Settings, marshalers map[string]LogsMarshaler) (*PulsarLogsProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
