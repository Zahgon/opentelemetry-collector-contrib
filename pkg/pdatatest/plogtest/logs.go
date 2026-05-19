// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package plogtest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/plogtest"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

// CompareLogs compares each part of two given Logs and returns
// an error if they don't match. The error describes what didn't match.
func CompareLogs(expected, actual plog.Logs, options ...CompareLogsOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching resources so that each can only be matched once

// CompareResourceLogs compares each part of two given ResourceLogs and returns
// an error if they don't match. The error describes what didn't match.
func CompareResourceLogs(expected, actual plog.ResourceLogs) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching scope logs so that each record can only be matched once

// CompareScopeLogs compares each part of two given LogRecordSlices and returns
// an error if they don't match. The error describes what didn't match.
func CompareScopeLogs(expected, actual plog.ScopeLogs) error { _ = "STUB: not implemented"; return nil }

// Keep track of matching records so that each record can only be matched once

// CompareLogRecord compares each part of two given LogRecord and returns
// an error if they don't match. The error describes what didn't match.
func CompareLogRecord(expected, actual plog.LogRecord) error { _ = "STUB: not implemented"; return nil }
