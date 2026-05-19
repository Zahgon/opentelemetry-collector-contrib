// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasetexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"

import (
	"context"
	"time"

	"github.com/scalyr/dataset-go/pkg/api/add_events"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// We define it here so we can easily mock it inside tests
var now = time.Now

// Prefix which is added to all the special / internal DataSet fields
const specialDataSetFieldNamePrefix string = "sca:"

// If a LogRecord doesn't contain severity or we can't map it to a valid DataSet severity, we use
// this value (3 - INFO) instead
const defaultDataSetSeverityLevel int = dataSetLogLevelInfo

// Constants for valid DataSet log levels (aka Event.sev int field value)
const (
	dataSetLogLevelFinest = 0
	dataSetLogLevelTrace  = 1
	dataSetLogLevelDebug  = 2
	dataSetLogLevelInfo   = 3
	dataSetLogLevelWarn   = 4
	dataSetLogLevelError  = 5
	dataSetLogLevelFatal  = 6
)

func createLogsExporter(ctx context.Context, set exporter.Settings, config component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func buildBody(settings LogsSettings, attrs map[string]any, value pcommon.Value) string {
	_ = "STUB: not implemented"
	// The message / body is stored as part of the "message" field on the DataSet event.
	return ""
}

// Additionally, we support de-composing complex message value (e.g. map / dictionary) into
// multiple event attributes.
//
// This functionality is behind a config option / feature flag and not enabled by default
// since no other existing DataSet integrations handle it in this manner (aka for out of
// the box consistency reasons).
// If user wants to achieve something like that, they usually handle that on the client
// (e.g. attribute processor or similar) or on the server (DataSet server side JSON parser
// for the message field).

// Function maps OTel severity on the LogRecord to DataSet severity level (number)
func mapOtelSeverityToDataSetSeverity(log plog.LogRecord) int {
	_ = "STUB: not implemented"
	// This function maps OTel severity level to DataSet severity levels
	//
	// Valid OTel levels - https://opentelemetry.io/docs/specs/otel/logs/data-model/#field-severitynumber
	// and valid DataSet ones - https://github.com/scalyr/logstash-output-scalyr/blob/master/lib/logstash/outputs/scalyr.rb#L70
	return 0
}

// Per docs, SeverityNumber is optional so if it's not present we fall back to SeverityText
// https://opentelemetry.io/docs/specs/otel/logs/data-model/#field-severitytext

func mapLogRecordSevNumToDataSetSeverity(sevNum plog.SeverityNumber) int {
	_ = "STUB: not implemented"
	// Maps LogRecord.SeverityNumber field value to DataSet severity value.
	return 0
}

// See https://opentelemetry.io/docs/specs/otel/logs/data-model/#field-severitynumber
// for OTEL mappings

// TRACE

// DEBUG

// INFO

// WARN

// ERROR

// FATAL / CRITICAL / EMERGENCY

func mapLogRecordSeverityTextToDataSetSeverity(sevText string) int {
	_ = "STUB: not implemented"
	// Maps LogRecord.SeverityText field value to DataSet severity value.
	return 0
}

func buildEventFromLog(
	log plog.LogRecord,
	resource pcommon.Resource,
	scope pcommon.InstrumentationScope,
	serverHost string,
	logSettings LogsSettings,
) *add_events.EventBundle {
	_ = "STUB: not implemented"
	return nil
}

// Event needs to always have timestamp set otherwise it will get set to unix epoch start time

// ObservedTimestamp should always be set, but in case if it's not, we fall back to
// current time

func (e *datasetExporter) consumeLogs(_ context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}
