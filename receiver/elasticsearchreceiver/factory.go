// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver"

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultCollectionInterval = 10 * time.Second
	defaultHTTPClientTimeout  = 10 * time.Second
)

// NewFactory creates a factory for elasticsearch receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates the default elasticsearchreceiver config.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

var errConfigNotES = errors.New("config was not an elasticsearch receiver config")

// createMetricsReceiver creates a metrics receiver for scraping elasticsearch metrics.
func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
