// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudfoundryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// This file implements factory for Cloud Foundry receiver.

const (
	defaultUAAUsername       = "admin"
	defaultRLPGatewayShardID = "opentelemetry"
	defaultURL               = "https://localhost"
)

// NewFactory creates a factory for collectd receiver.
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

func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
