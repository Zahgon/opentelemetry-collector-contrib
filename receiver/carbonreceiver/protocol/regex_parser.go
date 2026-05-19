// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package protocol // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"

import (
	"regexp"
)

const (
	metricNameCapturePrefix = "name_"
	keyCapturePrefix        = "key_"
)

// RegexParserConfig has the configuration for a parser that can breakdown a
// Carbon "metric path" and transform it in corresponding metric labels according
// to a series of regular expressions rules (see below for details).
//
// This is typically used to extract labels from a "naming hierarchy", see
// https://graphite.readthedocs.io/en/latest/feeding-carbon.html#step-1-plan-a-naming-hierarchy
//
// Examples:
//
// 1. Rule:
//   - regexp: "(?P<key_svc>[^.]+)\.(?P<key_host>[^.]+)\.cpu\.seconds"
//     name_prefix: cpu_seconds
//     labels:
//     k: v
//     Metric path: "service_name.host00.cpu.seconds"
//     Resulting metric:
//     name: cpu_seconds
//     label keys: {"svc", "host", "k"}
//     label values: {"service_name", "host00", "k"}
//
// 2. Rule:
//   - regexp: "^(?P<key_svc>[^.]+)\.(?P<key_host>[^.]+)\.(?P<name_0>[^.]+).(?P<name_1>[^.]+)$"
//     Metric path: "svc_02.host02.avg.duration"
//     Resulting metric:
//     name: avgduration
//     label keys: {"svc", "host"}
//     label values: {"svc_02", "host02"}
type RegexParserConfig struct {
	// Rules contains the regular expression rules to be used by the parser.
	// The first rule that matches and applies the transformations configured in
	// the respective RegexRule struct. If no rules match the metric is then
	// processed by the "plaintext" parser.
	Rules []*RegexRule `mapstructure:"rules"`

	// MetricNameSeparator is used when joining the name prefix of each individual
	// rule and the respective named captures that start with the prefix
	// "name_" (see RegexRule for more information).
	MetricNameSeparator string `mapstructure:"name_separator"`
}

// RegexRule describes how parts of the name of metric are going to be mapped
// to metric labels. The rule is only applied if the name matches the given
// regular expression.
type RegexRule struct {
	// Regular expression from which named matches are used to extract label
	// keys and values from Carbon metric paths.
	Regexp string `mapstructure:"regexp"`

	// NamePrefix is the prefix added to the metric name after extracting the
	// parts that will form labels and final metric name.
	NamePrefix string `mapstructure:"name_prefix"`

	// Labels are key-value pairs added as labels to the metrics that match this
	// rule.
	Labels map[string]string `mapstructure:"labels"`

	// MetricType selects the type of metric to be generated, supported values are
	// "gauge" (the default) and "cumulative".
	MetricType string `mapstructure:"type"`

	// Some fields cached after the compilation of the regular expression.
	compRegexp      *regexp.Regexp
	metricNameParts []string
}

var _ ParserConfig = (*RegexParserConfig)(nil)

// BuildParser builds the respective parser of the configuration instance.
func (rpc *RegexParserConfig) BuildParser() (Parser, error) {
	_ = "STUB: not implemented"
	return *new(Parser), nil
}

func compileRegexRules(rules []*RegexRule) error { _ = "STUB: not implemented"; return nil }

// Default capture.

// Correctly prefixed, nothing else to do.

type regexPathParser struct {
	rules []*RegexRule

	metricNameSeparator string

	// plaintextParser is used if no rule matches a given metric.
	plaintextPathParser plaintextPathParser
}

// parsePath converts the <metric_path> of a Carbon line (see pathParserHelper
// a full description of the line format) according to the RegexParserConfig
// settings.
func (rpp *regexPathParser) parsePath(path string, parsedPath *parsedPath) error {
	_ = "STUB: not implemented"
	return nil
}

// regexp pre-computes this slice.

// Skip unnamed groups.

// Skip unmatched groups.

func regexDefaultConfig() ParserConfig { _ = "STUB: not implemented"; return *new(ParserConfig) }
