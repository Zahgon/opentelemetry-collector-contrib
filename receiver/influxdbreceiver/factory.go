// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package influxdbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const defaultEndpoint = "localhost:8086"

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(_ context.Context, params receiver.Settings, cfg component.Config, nextConsumer consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
