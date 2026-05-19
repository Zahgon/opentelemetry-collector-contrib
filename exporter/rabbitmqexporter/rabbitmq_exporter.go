// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package rabbitmqexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/rabbitmqexporter"

import (
	"context"
	"crypto/tls"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/rabbitmqexporter/internal/publisher"
)

type rabbitmqExporter struct {
	config *Config
	tlsFactory
	settings       component.TelemetrySettings
	routingKey     string
	connectionName string
	*marshaler
	publisherFactory
	publisher publisher.Publisher
}

type (
	publisherFactory = func(publisher.DialConfig) (publisher.Publisher, error)
	tlsFactory       = func(context.Context) (*tls.Config, error)
)

func newRabbitmqExporter(cfg *Config, set component.TelemetrySettings, publisherFactory publisherFactory, tlsFactory tlsFactory, routingKey, connectionName string) *rabbitmqExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *rabbitmqExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *rabbitmqExporter) publishTraces(context context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *rabbitmqExporter) publishMetrics(context context.Context, metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *rabbitmqExporter) publishLogs(context context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *rabbitmqExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }
