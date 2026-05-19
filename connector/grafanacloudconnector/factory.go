// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package grafanacloudconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/grafanacloudconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

func NewFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesToMetricsConnector(_ context.Context, params connector.Settings, cfg component.Config, next consumer.Metrics) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}
