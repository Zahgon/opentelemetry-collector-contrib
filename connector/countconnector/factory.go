// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package countconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/xconnector"
	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

// NewFactory returns a ConnectorFactory.
func NewFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

// createDefaultConfig creates the default configuration.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createTracesToMetrics creates a traces to metrics connector based on provided config.
func createTracesToMetrics(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

// Error checked in Config.Validate()

// Error checked in Config.Validate()

// createMetricsToMetrics creates a metricds to metrics connector based on provided config.
func createMetricsToMetrics(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

// Error checked in Config.Validate()

// Error checked in Config.Validate()

// createLogsToMetrics creates a logs to metrics connector based on provided config.
func createLogsToMetrics(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

// Error checked in Config.Validate()

// createProfilesToMetrics creates a profiles to metrics connector based on provided config.
func createProfilesToMetrics(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (xconnector.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xconnector.Profiles), nil
}

// Error checked in Config.Validate()

type metricDef[K any] struct {
	condition *ottl.ConditionSequence[K]
	desc      string
	attrs     []AttributeConfig
}
