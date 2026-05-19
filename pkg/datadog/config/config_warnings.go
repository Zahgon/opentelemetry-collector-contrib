// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"

import (
	"go.opentelemetry.io/collector/confmap"
)

var _ error = (*deprecatedError)(nil)

// deprecatedError is an error related to a renamed setting.
type deprecatedError struct {
	// oldName of the configuration option.
	oldName string
	// newName of the configuration option.
	newName string
	// updateFn updates the configuration to map the old value into the new one.
	// It must only be called when the old value is set and is not the default.
	updateFn func(*Config)
}

// List of settings that are deprecated but not yet removed.
var renamedSettings = []deprecatedError{
	{
		oldName: "metrics::histograms::send_count_sum_metrics",
		newName: "metrics::histograms::send_aggregation_metrics",
		updateFn: func(c *Config) {
			c.Metrics.HistConfig.SendAggregations = c.Metrics.HistConfig.SendCountSum
		},
	},
	{
		oldName: "traces::peer_service_aggregation",
		newName: "traces::peer_tags_aggregation",
		updateFn: func(c *Config) {
			c.Traces.PeerTagsAggregation = c.Traces.PeerServiceAggregation
		},
	},
}

// Error implements the error interface.
func (e deprecatedError) Error() string { _ = "STUB: not implemented"; return "" }

// Check if the deprecated option is being used.
// Error out if both the old and new options are being used.
func (e deprecatedError) Check(configMap *confmap.Conf) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// UpdateCfg to move the old configuration value into the new one.
func (e deprecatedError) UpdateCfg(cfg *Config) {
	_ = "STUB: not implemented"

	// handleRenamedSettings for a given configuration map.
	// Error out if any pair of old-new options are set at the same time.
	return
}

func handleRenamedSettings(configMap *confmap.Conf, cfg *Config) (warnings []error, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only update config if old name is in use
