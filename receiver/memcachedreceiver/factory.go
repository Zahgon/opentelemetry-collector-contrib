// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package memcachedreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultEndpoint           = "localhost:11211"
	defaultTimeout            = 10 * time.Second
	defaultCollectionInterval = 10 * time.Second
)

// NewFactory creates a factory for memcached receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
