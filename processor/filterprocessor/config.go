// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset/regexp"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"
)

// Config defines configuration for Resource processor.
type Config struct {
	// ErrorMode determines how the processor reacts to errors that occur while processing an OTTL condition.
	// Valid values are `ignore` and `propagate`.
	// `ignore` means the processor ignores errors returned by conditions and continues on to the next condition. This is the recommended mode.
	// `propagate` means the processor returns the error up the pipeline.  This will result in the payload being dropped from the collector.
	// The default value is `ignore` when the `processor.filter.defaultErrorModeIgnore` feature gate is enabled (default), and `propagate` when disabled.
	ErrorMode ottl.ErrorMode `mapstructure:"error_mode"`

	// Deprecated: use TraceConditions instead.
	Spans filterconfig.MatchConfig `mapstructure:"spans"`
	// Deprecated: use MetricConditions instead.
	Metrics MetricFilters `mapstructure:"metrics"`
	// Deprecated: use LogConditions instead.
	Logs LogFilters `mapstructure:"logs"`
	// Deprecated: use TraceConditions instead.
	Traces TraceFilters `mapstructure:"traces"`
	// Deprecated: use ProfileConditions instead.
	Profiles ProfileFilters `mapstructure:"profiles"`

	MetricConditions  []condition.ContextConditions `mapstructure:"metric_conditions"`
	LogConditions     []condition.ContextConditions `mapstructure:"log_conditions"`
	TraceConditions   []condition.ContextConditions `mapstructure:"trace_conditions"`
	ProfileConditions []condition.ContextConditions `mapstructure:"profile_conditions"`

	resourceFunctions  map[string]ottl.Factory[*ottlresource.TransformContext]
	dataPointFunctions map[string]ottl.Factory[*ottldatapoint.TransformContext]
	logFunctions       map[string]ottl.Factory[*ottllog.TransformContext]
	metricFunctions    map[string]ottl.Factory[*ottlmetric.TransformContext]
	spanEventFunctions map[string]ottl.Factory[*ottlspanevent.TransformContext]
	spanFunctions      map[string]ottl.Factory[*ottlspan.TransformContext]
	profileFunctions   map[string]ottl.Factory[*ottlprofile.TransformContext]
}

// MetricFilters filters by Metric properties.
type MetricFilters struct {
	// Include match properties describe metrics that should be included in the Collector Service pipeline,
	// all other metrics should be dropped from further processing.
	// If both Include and Exclude are specified, Include filtering occurs first.
	Include *filterconfig.MetricMatchProperties `mapstructure:"include"`

	// Exclude match properties describe metrics that should be excluded from the Collector Service pipeline,
	// all other metrics should be included.
	// If both Include and Exclude are specified, Include filtering occurs first.
	Exclude *filterconfig.MetricMatchProperties `mapstructure:"exclude"`

	// RegexpConfig specifies options for the regexp match type
	RegexpConfig *regexp.Config `mapstructure:"regexp"`

	// ResourceConditions is a list of OTTL conditions for an ottlresource context.
	// If any condition resolves to true, the whole resource will be dropped.
	// Supports `and`, `or`, and `()`
	ResourceConditions []string `mapstructure:"resource"`

	// MetricConditions is a list of OTTL conditions for an ottlmetric context.
	// If any condition resolves to true, the metric will be dropped.
	// Supports `and`, `or`, and `()`
	MetricConditions []string `mapstructure:"metric"`

	// DataPointConditions is a list of OTTL conditions for an ottldatapoint context.
	// If any condition resolves to true, the datapoint will be dropped.
	// Supports `and`, `or`, and `()`
	DataPointConditions []string `mapstructure:"datapoint"`
}

// TraceFilters filters by OTTL conditions
type TraceFilters struct {
	// ResourceConditions is a list of OTTL conditions for an ottlresource context.
	// If any condition resolves to true, the whole resource will be dropped.
	// Supports `and`, `or`, and `()`
	ResourceConditions []string `mapstructure:"resource"`

	// SpanConditions is a list of OTTL conditions for an ottlspan context.
	// If any condition resolves to true, the span will be dropped.
	// Supports `and`, `or`, and `()`
	SpanConditions []string `mapstructure:"span"`

	// SpanEventConditions is a list of OTTL conditions for an ottlspanevent context.
	// If any condition resolves to true, the span event will be dropped.
	// Supports `and`, `or`, and `()`
	SpanEventConditions []string `mapstructure:"spanevent"`
}

// LogFilters filters by Log properties.
type LogFilters struct {
	// Include match properties describe logs that should be included in the Collector Service pipeline,
	// all other logs should be dropped from further processing.
	// If both Include and Exclude are specified, Include filtering occurs first.
	Include *LogMatchProperties `mapstructure:"include"`
	// Exclude match properties describe logs that should be excluded from the Collector Service pipeline,
	// all other logs should be included.
	// If both Include and Exclude are specified, Include filtering occurs first.
	Exclude *LogMatchProperties `mapstructure:"exclude"`

	// ResourceConditions is a list of OTTL conditions for an ottlresource context.
	// If any condition resolves to true, the whole resource will be dropped.
	// Supports `and`, `or`, and `()`
	ResourceConditions []string `mapstructure:"resource"`

	// LogConditions is a list of OTTL conditions for an ottllog context.
	// If any condition resolves to true, the log event will be dropped.
	// Supports `and`, `or`, and `()`
	LogConditions []string `mapstructure:"log_record"`
}

// LogMatchType specifies the strategy for matching against `plog.Log`s.
type LogMatchType string

// These are the MatchTypes that users can specify for filtering
// `plog.Log`s.
const (
	strictType = LogMatchType(filterset.Strict)
	regexpType = LogMatchType(filterset.Regexp)
)

var severityToNumber = map[string]plog.SeverityNumber{
	"1":      plog.SeverityNumberTrace,
	"2":      plog.SeverityNumberTrace2,
	"3":      plog.SeverityNumberTrace3,
	"4":      plog.SeverityNumberTrace4,
	"5":      plog.SeverityNumberDebug,
	"6":      plog.SeverityNumberDebug2,
	"7":      plog.SeverityNumberDebug3,
	"8":      plog.SeverityNumberDebug4,
	"9":      plog.SeverityNumberInfo,
	"10":     plog.SeverityNumberInfo2,
	"11":     plog.SeverityNumberInfo3,
	"12":     plog.SeverityNumberInfo4,
	"13":     plog.SeverityNumberWarn,
	"14":     plog.SeverityNumberWarn2,
	"15":     plog.SeverityNumberWarn3,
	"16":     plog.SeverityNumberWarn4,
	"17":     plog.SeverityNumberError,
	"18":     plog.SeverityNumberError2,
	"19":     plog.SeverityNumberError3,
	"20":     plog.SeverityNumberError4,
	"21":     plog.SeverityNumberFatal,
	"22":     plog.SeverityNumberFatal2,
	"23":     plog.SeverityNumberFatal3,
	"24":     plog.SeverityNumberFatal4,
	"TRACE":  plog.SeverityNumberTrace,
	"TRACE2": plog.SeverityNumberTrace2,
	"TRACE3": plog.SeverityNumberTrace3,
	"TRACE4": plog.SeverityNumberTrace4,
	"DEBUG":  plog.SeverityNumberDebug,
	"DEBUG2": plog.SeverityNumberDebug2,
	"DEBUG3": plog.SeverityNumberDebug3,
	"DEBUG4": plog.SeverityNumberDebug4,
	"INFO":   plog.SeverityNumberInfo,
	"INFO2":  plog.SeverityNumberInfo2,
	"INFO3":  plog.SeverityNumberInfo3,
	"INFO4":  plog.SeverityNumberInfo4,
	"WARN":   plog.SeverityNumberWarn,
	"WARN2":  plog.SeverityNumberWarn2,
	"WARN3":  plog.SeverityNumberWarn3,
	"WARN4":  plog.SeverityNumberWarn4,
	"ERROR":  plog.SeverityNumberError,
	"ERROR2": plog.SeverityNumberError2,
	"ERROR3": plog.SeverityNumberError3,
	"ERROR4": plog.SeverityNumberError4,
	"FATAL":  plog.SeverityNumberFatal,
	"FATAL2": plog.SeverityNumberFatal2,
	"FATAL3": plog.SeverityNumberFatal3,
	"FATAL4": plog.SeverityNumberFatal4,
}

var errInvalidSeverity = errors.New("not a valid severity")

// logSeverity is a type that represents a SeverityNumber as a string
type logSeverity string

// validate checks that the logSeverity is valid
func (l logSeverity) validate() error {
	_ = "STUB: not implemented"

	// No severity specified, which means to ignore this field.
	return nil
}

// severityNumber returns the severity number that the logSeverity represents
func (l logSeverity) severityNumber() plog.SeverityNumber {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber)
}

// LogMatchProperties specifies the set of properties in a log to match against and the
// type of string pattern matching to use.
type LogMatchProperties struct {
	// LogMatchType specifies the type of matching desired
	LogMatchType LogMatchType `mapstructure:"match_type"`

	// ResourceAttributes defines a list of possible resource attributes to match logs against.
	// A match occurs if any resource attribute matches all expressions in this given list.
	ResourceAttributes []filterconfig.Attribute `mapstructure:"resource_attributes"`

	// RecordAttributes defines a list of possible record attributes to match logs against.
	// A match occurs if any record attribute matches at least one expression in this given list.
	RecordAttributes []filterconfig.Attribute `mapstructure:"record_attributes"`

	// SeverityTexts is a list of strings that the LogRecord's severity text field must match
	// against.
	SeverityTexts []string `mapstructure:"severity_texts"`

	// SeverityNumberProperties defines how to match against a log record's SeverityNumber, if defined.
	SeverityNumberProperties *LogSeverityNumberMatchProperties `mapstructure:"severity_number"`

	// LogBodies is a list of strings that the LogRecord's body field must match
	// against.
	LogBodies []string `mapstructure:"bodies"`
}

// validate checks that the LogMatchProperties is valid
func (lmp LogMatchProperties) validate() error { _ = "STUB: not implemented"; return nil }

// isEmpty returns true if the properties is "empty" (meaning, there are no filters specified)
// if this is the case, the filter should be ignored.
func (lmp LogMatchProperties) isEmpty() bool { _ = "STUB: not implemented"; return false }

// matchProperties converts the LogMatchProperties to a corresponding filterconfig.MatchProperties
func (lmp LogMatchProperties) matchProperties() *filterconfig.MatchProperties {
	_ = "STUB: not implemented"
	return nil
}

// Include SeverityNumberProperties if defined

type LogSeverityNumberMatchProperties struct {
	// Min is the minimum severity needed for the log record to match.
	// This corresponds to the short names specified here:
	// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md#displaying-severity
	// this field is case-insensitive ("INFO" == "info")
	Min logSeverity `mapstructure:"min"`

	// MatchUndefined lets logs records with "unknown" severity match.
	// If MinSeverity is not set, this field is ignored, as fields are not matched based on severity.
	MatchUndefined bool `mapstructure:"match_undefined"`
}

// validate checks that the LogMatchProperties is valid
func (lmp LogSeverityNumberMatchProperties) validate() error { _ = "STUB: not implemented"; return nil }

// ProfileFilters filters by OTTL conditions
type ProfileFilters struct {
	_ struct{} // prevent unkeyed literals

	// ResourceConditions is a list of OTTL conditions for an ottlresource context.
	// If any condition resolves to true, the whole resource will be dropped.
	// Supports `and`, `or`, and `()`
	ResourceConditions []string `mapstructure:"resource"`

	// ProfileConditions is a list of OTTL conditions for an ottlprofile context.
	// If any condition resolves to true, the profile will be dropped.
	// Supports `and`, `or`, and `()`
	ProfileConditions []string `mapstructure:"profile"`
}

// Unmarshal is used internally by mapstructure to parse the filterprocessor configuration (Config),
// adding support to structured and flat configuration styles.
// When the flat configuration style is used, all conditions are grouped into a common.ContextConditions
// object, with empty [common.ContextConditions.Context] value.
// On the other hand, structured configurations are parsed following the mapstructure Config format.
//
// Example of flat configuration:
//
//	log_conditions:
//	  - resource.attributes["key1"] == "value"
//	  - resource.attributes["key2"] == "value"
//
// Example of structured configuration:
//
//	log_conditions:
//	  - context: "resource"
//	    conditions:
//	      - attributes["key1"] == "value"
//	      - attributes["key2"] == "value"
func (cfg *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// Array of strings means it's a basic configuration style

var _ component.Config = (*Config)(nil)

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) validateExplicitContextConfig() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) validateInferredContextConfig() error {
	_ = "STUB: not implemented"
	// Remove the old format.
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/41176
	return nil
}

func (cfg *Config) newTraceParserCollection(telemetrySettings component.TelemetrySettings) (*condition.TraceParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cfg *Config) newMetricParserCollection(telemetrySettings component.TelemetrySettings) (*condition.MetricParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cfg *Config) newLogParserCollection(telemetrySettings component.TelemetrySettings) (*condition.LogParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cfg *Config) newProfileParserCollection(telemetrySettings component.TelemetrySettings) (*condition.ProfileParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
