// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/parser"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/transport"
)

var _ receiver.Metrics = (*statsdReceiver)(nil)

// statsdReceiver implements the receiver.Metrics for StatsD protocol.
type statsdReceiver struct {
	settings receiver.Settings
	config   *Config

	server       transport.Server
	reporter     *reporter
	obsrecv      *receiverhelper.ObsReport
	parser       parser.Parser
	nextConsumer consumer.Metrics
	cancel       context.CancelFunc
}

// newReceiver creates the StatsD receiver with the given parameters.
func newReceiver(
	set receiver.Settings,
	config Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func buildTransportServer(config Config) (transport.Server, error) {
	_ = "STUB: not implemented"
	return *new(transport.Server), nil
}

// Start starts a UDP server that can process StatsD messages.
func (r *statsdReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Record every 100 to reduce overhead

// Shutdown stops the StatsD receiver.
func (r *statsdReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*statsdReceiver) Flush(ctx context.Context, metrics pmetric.Metrics, nextConsumer consumer.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}
