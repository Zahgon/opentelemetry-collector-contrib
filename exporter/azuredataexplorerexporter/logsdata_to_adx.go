// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuredataexplorerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type adxLog struct {
	Timestamp          string         // The timestamp of the occurrence. Formatted into string as RFC3339Nano
	ObservedTimestamp  string         // The timestamp of logs observed in opentelemetry collector.  Formatted into string as RFC3339Nano
	TraceID            string         // TraceId associated to the log
	SpanID             string         // SpanId associated to the log
	SeverityText       string         // The severity level of the log
	SeverityNumber     int32          // The severity number associated to the log
	Body               string         // The body/Text of the log
	ResourceAttributes map[string]any // JSON Resource attributes that can then be parsed.
	LogsAttributes     map[string]any // JSON attributes that can then be parsed.
}

// Convert the plog to the type ADXLog, this matches the scheme in the Log table in the database
func mapToAdxLog(resource pcommon.Resource, scope pcommon.InstrumentationScope, logData plog.LogRecord, _ *zap.Logger) *adxLog {
	_ = "STUB: not implemented"
	return nil
}
