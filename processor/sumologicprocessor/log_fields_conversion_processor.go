// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	SeverityNumberAttributeName = "loglevel"
	SeverityTextAttributeName   = "severitytext"
	SpanIDAttributeName         = "spanid"
	TraceIDAttributeName        = "traceid"
)

type LogFieldAttribute struct {
	Enabled bool   `mapstructure:"enabled"`
	Name    string `mapstructure:"name"`
}

type LogFieldAttributesConfig struct {
	SeverityNumberAttribute *LogFieldAttribute `mapstructure:"severity_number"`
	SeverityTextAttribute   *LogFieldAttribute `mapstructure:"severity_text"`
	SpanIDAttribute         *LogFieldAttribute `mapstructure:"span_id"`
	TraceIDAttribute        *LogFieldAttribute `mapstructure:"trace_id"`
}

// spanIDToHexOrEmptyString returns a hex string from SpanID.
// An empty string is returned, if SpanID is empty.
func spanIDToHexOrEmptyString(id pcommon.SpanID) string { _ = "STUB: not implemented"; return "" }

// traceIDToHexOrEmptyString returns a hex string from TraceID.
// An empty string is returned, if TraceID is empty.
func traceIDToHexOrEmptyString(id pcommon.TraceID) string { _ = "STUB: not implemented"; return "" }

var severityNumberToLevel = map[string]string{
	plog.SeverityNumberUnspecified.String(): "UNSPECIFIED",
	plog.SeverityNumberTrace.String():       "TRACE",
	plog.SeverityNumberTrace2.String():      "TRACE2",
	plog.SeverityNumberTrace3.String():      "TRACE3",
	plog.SeverityNumberTrace4.String():      "TRACE4",
	plog.SeverityNumberDebug.String():       "DEBUG",
	plog.SeverityNumberDebug2.String():      "DEBUG2",
	plog.SeverityNumberDebug3.String():      "DEBUG3",
	plog.SeverityNumberDebug4.String():      "DEBUG4",
	plog.SeverityNumberInfo.String():        "INFO",
	plog.SeverityNumberInfo2.String():       "INFO2",
	plog.SeverityNumberInfo3.String():       "INFO3",
	plog.SeverityNumberInfo4.String():       "INFO4",
	plog.SeverityNumberWarn.String():        "WARN",
	plog.SeverityNumberWarn2.String():       "WARN2",
	plog.SeverityNumberWarn3.String():       "WARN3",
	plog.SeverityNumberWarn4.String():       "WARN4",
	plog.SeverityNumberError.String():       "ERROR",
	plog.SeverityNumberError2.String():      "ERROR2",
	plog.SeverityNumberError3.String():      "ERROR3",
	plog.SeverityNumberError4.String():      "ERROR4",
	plog.SeverityNumberFatal.String():       "FATAL",
	plog.SeverityNumberFatal2.String():      "FATAL2",
	plog.SeverityNumberFatal3.String():      "FATAL3",
	plog.SeverityNumberFatal4.String():      "FATAL4",
}

// logFieldsConversionProcessor converts specific log entries to attributes which leads to presenting them as fields
// in the backend
type logFieldsConversionProcessor struct {
	LogFieldsAttributes *LogFieldAttributesConfig
}

func newLogFieldConversionProcessor(logFieldsAttributes *LogFieldAttributesConfig) *logFieldsConversionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (proc *logFieldsConversionProcessor) addAttributes(log plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

func (proc *logFieldsConversionProcessor) processLogs(logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*logFieldsConversionProcessor) processMetrics(pmetric.Metrics) error {
	_ = "STUB: not implemented"
	// No-op. Metrics should not be translated.
	return nil
}

func (*logFieldsConversionProcessor) processTraces(ptrace.Traces) error {
	_ = "STUB: not implemented"
	// No-op. Traces should not be translated.
	return nil
}

func (proc *logFieldsConversionProcessor) isEnabled() bool { _ = "STUB: not implemented"; return false }

func (*logFieldsConversionProcessor) ConfigPropertyName() string {
	_ = "STUB: not implemented"
	return ""
}
