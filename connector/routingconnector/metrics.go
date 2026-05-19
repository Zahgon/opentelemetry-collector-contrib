// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package routingconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricsConnector struct {
	component.StartFunc
	component.ShutdownFunc

	logger *zap.Logger
	config *Config
	router *router[consumer.Metrics]
}

func newMetricsConnector(
	set connector.Settings,
	config component.Config,
	metrics consumer.Metrics,
) (*metricsConnector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*metricsConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (c *metricsConnector) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// all metrics are routed

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// anything left wasn't matched by any route. Send to default consumer

func groupAllMetrics(
	groups map[consumer.Metrics]pmetric.Metrics,
	cons consumer.Metrics,
	metrics pmetric.Metrics,
) {
	_ = "STUB: not implemented"
	return
}
