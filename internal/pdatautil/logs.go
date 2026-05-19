// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pdatautil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

// FlattenResourceLogs moves each LogRecord onto a dedicated ResourceLogs and ScopeLogs.
// Modifications are made in place. Order of LogRecords is preserved.
func FlattenLogs(rls plog.ResourceLogsSlice) { _ = "STUB: not implemented"; return }

// GroupByResourceLogs groups ScopeLogs by Resource. Modifications are made in place.
func GroupByResourceLogs(rls plog.ResourceLogsSlice) {
	_ = "STUB: not implemented"
	// Hash each ResourceLogs based on identifying information.
	return
}

// Find the first occurrence of each hash and note the index.

// Merge Resources with the same hash.

// This is the first occurrence of this hash.

// Remove the ResourceLogs which were merged onto others.

// Merge ScopeLogs within each ResourceLogs.

// GroupByScopeLogs groups LogRecords by scope. Modifications are made in place.
func GroupByScopeLogs(sls plog.ScopeLogsSlice) {
	_ = "STUB: not implemented"
	// Hash each ScopeLogs based on identifying information.
	return
}

// Find the first occurrence of each hash and note the index.

// Merge ScopeLogs with the same hash.

// This is the first occurrence of this hash.

// Remove the ScopeLogs which were merged onto others.

// Creates a hash based on the ScopeLogs attributes, name, and version
func HashScopeLogs(sl plog.ScopeLogs) [16]byte { _ = "STUB: not implemented"; return nil }
