// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver"

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
	defaultCloud              = azureCloud
)

var errConfigNotAzureMonitor = errors.New("Config was not a Azure Monitor receiver config")

// NewFactory creates a new receiver factory
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates the default configuration for the receiver
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(_ context.Context, params receiver.Settings, rConf component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
