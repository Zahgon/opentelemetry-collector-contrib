// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package simpleprometheusreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// This file implements factory for prometheus_simple receiver
const (
	defaultEndpoint    = "localhost:9090"
	defaultMetricsPath = "/metrics"
)

var defaultCollectionInterval = 10 * time.Second

// NewFactory creates a factory for "Simple" Prometheus receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
