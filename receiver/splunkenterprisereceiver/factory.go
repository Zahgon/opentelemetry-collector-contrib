// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkenterprisereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultInterval          = 10 * time.Minute
	defaultMaxSearchWaitTime = 60 * time.Second
)

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	// Default HttpClient settings
	return *new(component.Config)
}

// Default ScraperController settings

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	baseCfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
