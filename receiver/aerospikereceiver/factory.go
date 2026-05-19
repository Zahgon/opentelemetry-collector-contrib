// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aerospikereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultEndpoint              = "localhost:3000"
	defaultTimeout               = 20 * time.Second
	defaultCollectClusterMetrics = false
)

// NewFactory creates a new ReceiverFactory with default configuration
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createMetricsReceiver creates a new MetricsReceiver using scraperhelper
func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
