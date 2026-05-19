// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package apachesparkreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var errConfigNotSpark = errors.New("config was not a Spark receiver config")

// NewFactory creates a new receiver factory for Spark
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates a config for Spark with as many default values as possible
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetricsReceiver creates the metric receiver for Spark
func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	config component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
