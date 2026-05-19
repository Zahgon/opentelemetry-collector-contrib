// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricsaslogsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/metricsaslogsconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

func NewFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

func createMetricsToLogs(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}
