// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"

import (
	"context"

	_ "github.com/prometheus/prometheus/plugins" // init() of this package registers service discovery impl.
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
)

// NewFactory creates a new Prometheus receiver factory.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func configWarnings(logger *zap.Logger, cfg *Config) { _ = "STUB: not implemented"; return }
