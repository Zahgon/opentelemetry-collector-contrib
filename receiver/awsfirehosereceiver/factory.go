// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultEndpoint = "localhost:4433"
)

// NewFactory creates a receiver factory for awsfirehose. Currently, only
// available in metrics pipelines.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates a default config with the endpoint set
// to port 8443 and the record type set to the CloudWatch metric stream.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetricsReceiver implements the CreateMetrics function type.
func createMetricsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// createMetricsReceiver implements the CreateMetricsReceiver function type.
func createLogsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
