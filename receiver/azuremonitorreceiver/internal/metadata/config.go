// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver/internal/metadata"

// MetricsBuilderConfig is a structural subset of an otherwise 1-1 copy of metadata.yaml
type MetricsBuilderConfig struct {
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	_ = "STUB: not implemented"
	return *new(MetricsBuilderConfig)
}

// Don't use DefaultResourceAttributesConfig() here to be able to set true by default.
