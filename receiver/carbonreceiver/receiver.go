// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package carbonreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/internal/transport"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"
)

var errEmptyEndpoint = errors.New("empty endpoint")

// carbonreceiver implements a receiver.Metrics for Carbon plaintext, aka "line", protocol.
// see https://graphite.readthedocs.io/en/latest/feeding-carbon.html#the-plaintext-protocol.
type carbonReceiver struct {
	settings receiver.Settings
	config   *Config

	server       transport.Server
	reporter     transport.Reporter
	parser       protocol.Parser
	nextConsumer consumer.Metrics
}

var _ receiver.Metrics = (*carbonReceiver)(nil)

// newMetricsReceiver creates the Carbon receiver with the given configuration.
func newMetricsReceiver(
	set receiver.Settings,
	config Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// Set the defaults

func buildTransportServer(config Config) (transport.Server, error) {
	_ = "STUB: not implemented"
	return *new(transport.Server), nil
}

// Start tells the receiver to start its processing.
// By convention the consumer of the received data is set when the receiver
// instance is created.
func (r *carbonReceiver) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown tells the receiver that should stop reception,
// giving it a chance to perform any necessary clean-up.
func (r *carbonReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
