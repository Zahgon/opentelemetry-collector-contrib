// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalkingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver"

// This file implements factory for skywalking receiver.

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

const (
	// Protocol values.
	protoGRPC = "grpc"
	protoHTTP = "http"

	// Default endpoints to bind to.
	defaultGRPCEndpoint = "localhost:11800"
	defaultHTTPEndpoint = "localhost:12800"
)

// NewFactory creates a new Skywalking receiver factory.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// CreateDefaultConfig creates the default configuration for Skywalking receiver.
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
	// Convert settings in the source c to configuration struct
	// that Skywalking receiver understands.
	return *new(receiver.Traces), nil
}

// createMetricsReceiver creates a metrics receiver based on provided config.
func createMetricsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	// Convert settings in the source c to configuration struct
	// that Skywalking receiver understands.
	return *new(receiver.Metrics), nil
}

// create the config that Skywalking receiver will use.
func createConfiguration(rCfg *Config) (*configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set ports

// extract the port number from string in "address:port" format. If the
// port number cannot be extracted returns an error.
func extractPortFromEndpoint(endpoint string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var receivers = sharedcomponent.NewSharedComponents()
