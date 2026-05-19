// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

var consumerCapabilities = consumer.Capabilities{MutatesData: true}

// NewFactory returns a new factory for the Metrics transform processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

// validateConfiguration validates the input configuration has all of the required fields for the processor
// An error is returned if there are any invalid inputs.
func validateConfiguration(config *Config) error { _ = "STUB: not implemented"; return nil }

// buildHelperConfig constructs the maps that will be useful for the operations
func buildHelperConfig(config *Config, version string) ([]internalTransform, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createFilter(filterConfig filterConfig) (internalFilter, error) {
	_ = "STUB: not implemented"
	return *new(internalFilter), nil
}

// createLabelValueMapping creates the labelValue rename mappings based on the valueActions
func createLabelValueMapping(valueActions []valueAction, version string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// sliceToSet converts slice of strings to set of strings
// Returns the set of strings
func sliceToSet(slice []string) map[string]bool { _ = "STUB: not implemented"; return nil }

func getMatcherMap(strMap map[string]string, ctor func(string) (StringMatcher, error)) (map[string]StringMatcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
