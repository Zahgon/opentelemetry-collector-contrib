// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

// This file implements factory for Splunk HEC receiver.

const (
	// Default endpoint to bind to.
	defaultEndpoint = "localhost:8088"
)

// NewFactory creates a factory for Splunk HEC receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// CreateDefaultConfig creates the default configuration for Splunk HEC receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// CreateMetrics creates a metrics receiver based on provided config.
func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// createLogsReceiver creates a logs receiver based on provided config.
func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

var receivers = sharedcomponent.NewSharedComponents()
