// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fluentforwardreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"

import (
	"context"
	"net"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
)

// Give the event channel a bit of buffer to help reduce backpressure on
// FluentBit and increase throughput.
const eventChannelLength = 100

type fluentReceiver struct {
	collector *collector
	listener  net.Listener
	conf      *Config
	logger    *zap.Logger
	server    *server
	cancel    context.CancelFunc
}

func newFluentReceiver(set receiver.Settings, conf *Config, next consumer.Logs) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func (r *fluentReceiver) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *fluentReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// Close the listener to stop accepting new connections
