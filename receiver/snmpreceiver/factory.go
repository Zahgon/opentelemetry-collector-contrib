// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package snmpreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var errConfigNotSNMP = errors.New("config was not a SNMP receiver config")

// NewFactory creates a new receiver factory for SNMP
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates a config for SNMP with as many default values as possible
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetricsReceiver creates the metric receiver for SNMP
func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	config component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// addMissingConfigDefaults adds any missing config parameters that have defaults
func addMissingConfigDefaults(cfg *Config) error {
	_ = "STUB: not implemented"
	// Add the schema prefix to the endpoint if it doesn't contain one
	return nil
}

// Add default port to endpoint if it doesn't contain one

// Set defaults for metric configs
