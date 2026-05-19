// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datadogreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

// NewFactory creates a factory for DataDog receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesReceiver(ctx context.Context, params receiver.Settings, cfg component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func createMetricsReceiver(ctx context.Context, params receiver.Settings, cfg component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func createLogsReceiver(ctx context.Context, params receiver.Settings, cfg component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

var receivers = sharedcomponent.NewSharedComponents()
