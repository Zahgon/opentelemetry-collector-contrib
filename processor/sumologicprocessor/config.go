// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"

import (
	"go.opentelemetry.io/collector/component"
)

type Config struct {
	AddCloudNamespace           bool                     `mapstructure:"add_cloud_namespace"`
	TranslateAttributes         bool                     `mapstructure:"translate_attributes"`
	TranslateTelegrafAttributes bool                     `mapstructure:"translate_telegraf_attributes"`
	NestAttributes              NestingProcessorConfig   `mapstructure:"nest_attributes"`
	AggregateAttributes         []AggregationPair        `mapstructure:"aggregate_attributes"`
	LogFieldsAttributes         LogFieldAttributesConfig `mapstructure:"field_attributes"`
	TranslateDockerMetrics      bool                     `mapstructure:"translate_docker_metrics"`
}

type AggregationPair struct {
	Attribute string   `mapstructure:"attribute"`
	Prefixes  []string `mapstructure:"prefixes"`
}

const (
	defaultAddCloudNamespace           = true
	defaultTranslateAttributes         = true
	defaultTranslateTelegrafAttributes = true
	defaultTranslateDockerMetrics      = false

	// Nesting processor default config
	defaultNestingEnabled            = false
	defaultNestingSeparator          = "."
	defaultNestingSquashSingleValues = false

	defaultAddSeverityNumberAttribute = false
	defaultAddSeverityTextAttribute   = false
	defaultAddSpanIDAttribute         = false
	defaultAddTraceIDAttribute        = false
)

var _ component.Config = (*Config)(nil)

func defaultNestingInclude() []string { _ = "STUB: not implemented"; return nil }

func defaultNestingExclude() []string { _ = "STUB: not implemented"; return nil }

func defaultAggregateAttributes() []AggregationPair { _ = "STUB: not implemented"; return nil }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Validate config
func (*Config) Validate() error { _ = "STUB: not implemented"; return nil }
