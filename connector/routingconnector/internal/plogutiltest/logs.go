// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package plogutiltest // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector/internal/plogutiltest"

import "go.opentelemetry.io/collector/pdata/plog"

// NewLogs returns a plog.Logs with a uniform structure where resources, scopes, and
// log records are identical across all instances, except for one identifying field.
//
// Identifying fields:
// - Resources have an attribute called "resourceName" with a value of "resourceN".
// - Scopes have a name with a value of "scopeN".
// - LogRecords have a body with a value of "logN".
//
// Example: NewLogs("AB", "XYZ", "1234") returns:
//
//	resourceA, resourceB
//	    each with scopeX, scopeY, scopeZ
//	        each with log1, log2, log3, log4
//
// Each byte in the input string is a unique ID for the corresponding element.
func NewLogs(resourceIDs, scopeIDs, logRecordIDs string) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func NewLogsFromOpts(resources ...plog.ResourceLogs) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func Resource(id string, scopes ...plog.ScopeLogs) plog.ResourceLogs {
	_ = "STUB: not implemented"
	return *new(plog.ResourceLogs)
}

func Scope(id string, logs ...plog.LogRecord) plog.ScopeLogs {
	_ = "STUB: not implemented"
	return *new(plog.ScopeLogs)
}

func LogRecord(id string) plog.LogRecord { _ = "STUB: not implemented"; return *new(plog.LogRecord) }
