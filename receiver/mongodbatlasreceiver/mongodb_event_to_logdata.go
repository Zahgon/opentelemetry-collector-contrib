// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/model"
)

const (
	// Number of log attributes to add to the plog.LogRecordSlice for host logs.
	totalLogAttributes = 10
	// Number of log attributes to add to the plog.LogRecordSlice for audit logs.
	totalAuditLogAttributes = 16

	// Number of resource attributes to add to the plog.ResourceLogs.
	totalResourceAttributes = 4
)

// jsonTimestampLayout for the timestamp format in the plog.Logs structure
const (
	jsonTimestampLayout    = "2006-01-02T15:04:05.000-07:00"
	consoleTimestampLayout = "2006-01-02T15:04:05.000-0700"
)

// Severity mapping of the mongodb atlas logs
var severityMap = map[string]plog.SeverityNumber{
	"F":  plog.SeverityNumberFatal,
	"E":  plog.SeverityNumberError,
	"W":  plog.SeverityNumberWarn,
	"I":  plog.SeverityNumberInfo,
	"D":  plog.SeverityNumberDebug,
	"D1": plog.SeverityNumberDebug,
	"D2": plog.SeverityNumberDebug2,
	"D3": plog.SeverityNumberDebug3,
	"D4": plog.SeverityNumberDebug4,
	"D5": plog.SeverityNumberDebug4,
}

// mongoAuditEventToLogRecord converts model.AuditLog event to plog.LogRecordSlice and adds the resource attributes.
func mongodbAuditEventToLogData(logger *zap.Logger, logs []model.AuditLog, pc projectContext, hostname, logName string, clusterInfo clusterInfo) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// Attributes related to the object causing the event.

// Insert Raw Log message into Body of LogRecord

// Since Audit Logs don't have a severity/level
// Set the "SeverityNumber" and "SeverityText" to INFO

// mongoEventToLogRecord converts model.LogEntry event to plog.LogRecordSlice and adds the resource attributes.
func mongodbEventToLogData(logger *zap.Logger, logs []model.LogEntry, pc projectContext, hostname, logName string, clusterInfo clusterInfo) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// Attributes related to the object causing the event.

// Insert Raw Log message into Body of LogRecord

// Set the "SeverityNumber" and "SeverityText" if a known type of
// severity is found.

//nolint:errcheck

// log ID is not present on MongoDB 4.2 systems

func tsLayout(clusterVersion string) string { _ = "STUB: not implemented"; return "" }
