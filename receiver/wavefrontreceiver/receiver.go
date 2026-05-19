// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wavefrontreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var _ receiver.Metrics = (*metricsReceiver)(nil)

type metricsReceiver struct {
	cfg            *Config
	set            receiver.Settings
	nextConsumer   consumer.Metrics
	carbonReceiver receiver.Metrics
}

func newMetricsReceiver(cfg *Config, set receiver.Settings, nextConsumer consumer.Metrics) *metricsReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (r *metricsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Wavefront is very similar to Carbon: it is TCP based in which each received
// text line represents a single metric data point. They differ on the format
// of their textual representation.
//
// The Wavefront receiver leverages the Carbon receiver code by implementing
// a dedicated parser for its format.

// TODO: update after other parsers are implemented for Carbon receiver.

func (r *metricsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
