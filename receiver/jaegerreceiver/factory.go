// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaegerreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"

// This file implements factory for Jaeger receiver.

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	// Default endpoints to bind to.
	defaultGRPCEndpoint          = "localhost:14250"
	defaultHTTPEndpoint          = "localhost:14268"
	defaultThriftCompactEndpoint = "localhost:6831"
	defaultThriftBinaryEndpoint  = "localhost:6832"
)

// NewFactory creates a new Jaeger receiver factory.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// CreateDefaultConfig creates the default configuration for Jaeger receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createTracesReceiver creates a trace receiver based on provided config.
func createTracesReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	// Convert settings in the source config to configuration struct
	// that Jaeger receiver understands.
	// Error handling for the conversion is done in the Validate function from the Config object itself.
	return *new(receiver.Traces), nil
}

// Create the receiver.
