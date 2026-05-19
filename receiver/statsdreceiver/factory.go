// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"

import (
	"context"
	"os"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/protocol"
)

const (
	defaultBindEndpoint        = "localhost:8125"
	defaultAggregationInterval = 60 * time.Second
	defaultEnableMetricType    = false
	defaultIsMonotonicCounter  = false
	defaultSocketPermissions   = os.FileMode(0o622)
)

var defaultTimerHistogramMapping = []protocol.TimerHistogramMapping{{StatsdType: "timer", ObserverType: "gauge"}, {StatsdType: "histogram", ObserverType: "gauge"}, {StatsdType: "distribution", ObserverType: "gauge"}}

// NewFactory creates a factory for the StatsD receiver.
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
