// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package slowsqlconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/slowsqlconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

// NewFactory creates a factory for the slowsql connector.
func NewFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesToLogsConnector(_ context.Context, params connector.Settings, cfg component.Config, nextConsumer consumer.Logs) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}
