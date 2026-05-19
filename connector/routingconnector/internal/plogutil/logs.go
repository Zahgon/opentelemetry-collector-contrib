// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package plogutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector/internal/plogutil"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

// MoveResourcesIf calls f sequentially for each ResourceLogs present in the first plog.Logs.
// If f returns true, the element is removed from the first plog.Logs and added to the second plog.Logs.
func MoveResourcesIf(from, to plog.Logs, f func(plog.ResourceLogs) bool) {
	_ = "STUB: not implemented"
	return
}

// CopyResourcesIf calls f sequentially for each ResourceLogs present in the first plog.Logs.
// If f returns true, the element is copied from the first plog.Logs to the second plog.Logs.
func CopyResourcesIf(from, to plog.Logs, f func(plog.ResourceLogs) bool) {
	_ = "STUB: not implemented"
	return
}

// MoveRecordsWithContextIf calls f sequentially for each LogRecord present in the first plog.Logs.
// If f returns true, the element is removed from the first plog.Logs and added to the second plog.Logs.
// Notably, the Resource and Scope associated with the LogRecord are created in the second plog.Logs only once.
// Resources or Scopes are removed from the original if they become empty. All ordering is preserved.
func MoveRecordsWithContextIf(from, to plog.Logs, f func(plog.ResourceLogs, plog.ScopeLogs, plog.LogRecord) bool) {
	_ = "STUB: not implemented"
	return
}

// CopyRecordsWithContextIf calls f sequentially for each LogRecord present in the first plog.Logs.
// If f returns true, the element is copied from the first plog.Logs to the second plog.Logs.
// Notably, the Resource and Scope associated with the LogRecord are created in the second plog.Logs only once.
func CopyRecordsWithContextIf(from, to plog.Logs, f func(plog.ResourceLogs, plog.ScopeLogs, plog.LogRecord) bool) {
	_ = "STUB: not implemented"
	return
}
