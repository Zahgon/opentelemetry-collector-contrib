// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/config"

import (
	"regexp"

	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	// defaultExponentialHistogramMaxSize is the default maximum number
	// of buckets per positive or negative number range. 160 buckets
	// default supports a high-resolution histogram able to cover a
	// long-tail latency distribution from 1ms to 100s with a relative
	// error of less than 5%.
	// Ref: https://opentelemetry.io/docs/specs/otel/metrics/sdk/#base2-exponential-bucket-histogram-aggregation
	defaultExponentialHistogramMaxSize = 160
)

var defaultHistogramBuckets = []float64{
	2, 4, 6, 8, 10, 50, 100, 200, 400, 800, 1000, 1400, 2000, 5000, 10_000, 15_000,
}

// Regex for [key] selector after ExtractGrokPatterns
var grokPatternKey = regexp.MustCompile(`ExtractGrokPatterns\([^)]*\)\s*\[[^\]]+\]`)

var _ confmap.Unmarshaler = (*Config)(nil)

// Config for the connector. Each configuration field describes the metrics
// to produce from a specific signal.
type Config struct {
	Spans      []MetricInfo `mapstructure:"spans"`
	Datapoints []MetricInfo `mapstructure:"datapoints"`
	Logs       []MetricInfo `mapstructure:"logs"`
	Profiles   []MetricInfo `mapstructure:"profiles"`
	// ErrorMode determines how the connector reacts to errors that occur while processing an OTTL
	// condition or statement during runtime data consumption. This setting does NOT affect errors
	// during OTTL statement parsing at configuration time - those will always cause startup failures.
	// Valid values are `propagate`, `ignore`, and `silent`.
	// `propagate` means the connector returns the error up the pipeline. This will result in the
	// payload being dropped from the collector.
	// `ignore` means the connector ignores errors returned by conditions and continues processing.
	// If an error occurs, the record is skipped and the error is logged.
	// `silent` means the connector ignores errors returned by conditions and continues processing.
	// If an error occurs, the record is skipped and the error is not logged.
	// The default value is `propagate`.
	ErrorMode ottl.ErrorMode `mapstructure:"error_mode"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// collect all errors at once

// Unmarshal implements the confmap.Unmarshaler interface. It allows
// unmarshaling the config with a custom logic to allow setting
// default values when/if required.
func (c *Config) Unmarshal(collectorCfg *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

type Attribute struct {
	Key string `mapstructure:"key"`
	// KeysExpression is an OTTL value expression that resolves to a list
	// of attribute keys at runtime. The expression must return a
	// pcommon.Slice or []string. Exactly one of Key or KeysExpression
	// must be set.
	KeysExpression string `mapstructure:"keys_expression"`
	Optional       bool   `mapstructure:"optional"`
	DefaultValue   any    `mapstructure:"default_value"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type Histogram struct {
	Buckets []float64 `mapstructure:"buckets"`
	Count   string    `mapstructure:"count"`
	Value   string    `mapstructure:"value"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type ExponentialHistogram struct {
	MaxSize int32  `mapstructure:"max_size"`
	Count   string `mapstructure:"count"`
	Value   string `mapstructure:"value"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type Sum struct {
	Value       string `mapstructure:"value"`
	IsMonotonic bool   `mapstructure:"monotonic"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type Gauge struct {
	Value string `mapstructure:"value"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// MetricInfo defines the structure of the metric produced by the connector.
type MetricInfo struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	// Unit, if not-empty, will set the unit associated with the metric.
	// See: https://github.com/open-telemetry/opentelemetry-collector/blob/b06236cc794982916cc956f20828b3e18eb33264/pdata/pmetric/generated_metric.go#L72-L81
	Unit string `mapstructure:"unit"`
	// IncludeResourceAttributes is a list of resource attributes that
	// needs to be included in the generated metric. If no resource
	// attribute is included in the list then all attributes are included.
	IncludeResourceAttributes []Attribute `mapstructure:"include_resource_attributes"`
	Attributes                []Attribute `mapstructure:"attributes"`
	// Conditions are a set of OTTL conditions which are ORed. Data is
	// processed into metrics only if the sequence evaluates to true.
	Conditions           []string                                      `mapstructure:"conditions"`
	Histogram            configoptional.Optional[Histogram]            `mapstructure:"histogram"`
	ExponentialHistogram configoptional.Optional[ExponentialHistogram] `mapstructure:"exponential_histogram"`
	Sum                  configoptional.Optional[Sum]                  `mapstructure:"sum"`
	Gauge                configoptional.Optional[Gauge]                `mapstructure:"gauge"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (mi *MetricInfo) ensureDefaults() { _ = "STUB: not implemented"; return }

// Add default buckets if explicit histogram is defined

// validateAttributeConfigs validates a list of Attribute configs. Each entry
// must have exactly one of Key or KeysExpression set. OTTL expressions are
// parsed to verify syntax. The label parameter is used in error messages.
func validateAttributeConfigs[K any](attrs []Attribute, pc *ottl.ParserCollection[*ottl.ValueExpression[K]], contextName, label string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mi *MetricInfo) validateHistogram() error { _ = "STUB: not implemented"; return nil }

func (mi *MetricInfo) validateSum() error { _ = "STUB: not implemented"; return nil }

func (mi *MetricInfo) validateGauge() error { _ = "STUB: not implemented"; return nil }

// validateMetricInfo is an utility method validate all supported metric
// types defined for the metric info including any ottl expressions.
// Condition parsing is handled by the caller because it needs a
// signal-specific filterottl helper that is not generic over K.
func validateMetricInfo[K any](mi *MetricInfo, pc *ottl.ParserCollection[*ottl.ValueExpression[K]], contextName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Exactly one metric should be defined. Also, validate OTTL expressions,
// note that, here we only evaluate if statements are valid. Check for
// required statements are left to the other validations.

// if ExtractGrokPatterns is used, validate the key selector

// Ensure a [key] selector is present after ExtractGrokPatterns
