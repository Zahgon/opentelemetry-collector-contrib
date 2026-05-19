// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solacereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	// default value for max unacked messages
	defaultMaxUnacked int32 = 1000
	// default value for host
	defaultHost string = "localhost:5671"
)

// NewFactory creates a factory for Solace receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// CreateTraces creates a trace receiver based on provided config. Component is not shared
func createTracesReceiver(
	_ context.Context,
	params receiver.Settings,
	receiverConfig component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

// pass cfg, params and next consumer through
