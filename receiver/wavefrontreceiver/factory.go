// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wavefrontreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	// tcpIdleTimeoutDefault is the default timeout for idle TCP connections.
	tcpIdleTimeoutDefault = 30 * time.Second
)

// This file implements factory for the Wavefront receiver.

// NewFactory creates a factory for WaveFront receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
