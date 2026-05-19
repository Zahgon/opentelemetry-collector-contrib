// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package routingconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type tracesConnector struct {
	component.StartFunc
	component.ShutdownFunc

	logger *zap.Logger
	config *Config
	router *router[consumer.Traces]
}

func newTracesConnector(
	set connector.Settings,
	config component.Config,
	traces consumer.Traces,
) (*tracesConnector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*tracesConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (c *tracesConnector) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// all traces are routed

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// anything left wasn't matched by any route. Send to default consumer

func groupAllTraces(
	groups map[consumer.Traces]ptrace.Traces,
	cons consumer.Traces,
	traces ptrace.Traces,
) {
	_ = "STUB: not implemented"
	return
}
