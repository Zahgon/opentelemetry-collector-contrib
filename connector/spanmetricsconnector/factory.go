// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package spanmetricsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

const (
	DefaultNamespace = "traces.span.metrics"
)

// NewFactory creates a factory for the spanmetrics connector.
func NewFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesToMetricsConnector(ctx context.Context, params connector.Settings, cfg component.Config, nextConsumer consumer.Metrics) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

// This never happens: the OpenTelemetry Collector automatically adds this attribute.
// See: https://github.com/open-telemetry/opentelemetry-collector/blob/main/service/internal/resource/config.go#L31
//
// The fallback logic below exists solely for lifecycle tests in generated_component_test.go,
// where the mocked telemetry setting does not include the service.instance.id attribute.
