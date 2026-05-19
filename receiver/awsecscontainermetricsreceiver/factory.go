// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsecscontainermetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// Factory for awscontainermetrics
const (
	// Default collection interval. Every 20s the receiver will collect metrics from Amazon ECS Task Metadata Endpoint
	defaultCollectionInterval = 20 * time.Second
)

// NewFactory creates a factory for AWS ECS Container Metrics receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig returns a default config for the receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// CreateMetrics creates an AWS ECS Container Metrics receiver.
func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	baseCfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
