// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package plogtest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/plogtest"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// CompareLogsOption can be used to mutate expected and/or actual logs before comparing.
type CompareLogsOption interface {
	applyOnLogs(expected, actual plog.Logs)
}

type compareLogsOptionFunc func(expected, actual plog.Logs)

func (f compareLogsOptionFunc) applyOnLogs(expected, actual plog.Logs) {
	_ = "STUB: not implemented"
	return

	// IgnoreResourceAttributeValue is a CompareLogsOption that removes a resource attribute
	// from all resources.
}

func IgnoreResourceAttributeValue(attributeName string) CompareLogsOption {
	_ = "STUB: not implemented"
	return *new(CompareLogsOption)
}

type ignoreResourceAttributeValue struct {
	attributeName string
}

func (opt ignoreResourceAttributeValue) applyOnLogs(expected, actual plog.Logs) {
	_ = "STUB: not implemented"
	return
}

func (opt ignoreResourceAttributeValue) maskLogsResourceAttributeValue(logs plog.Logs) {
	_ = "STUB: not implemented"
	return
}

// IgnoreResourceEntityRefs is a CompareLogsOption that clears entity references
// on all resources.
func IgnoreResourceEntityRefs() CompareLogsOption {
	_ = "STUB: not implemented"
	return *new(CompareLogsOption)
}

func maskLogsResourceEntityRefs(logs plog.Logs) { _ = "STUB: not implemented"; return }

// IgnoreLogRecordAttributeValue is a CompareLogsOption that sets the value of an attribute
// to empty bytes for every log record
func IgnoreLogRecordAttributeValue(attributeName string) CompareLogsOption {
	_ = "STUB: not implemented"
	return *new(CompareLogsOption)
}

type ignoreLogRecordAttributeValue struct {
	attributeName string
}

func (opt ignoreLogRecordAttributeValue) applyOnLogs(expected, actual plog.Logs) {
	_ = "STUB: not implemented"
	return
}

func (opt ignoreLogRecordAttributeValue) maskLogRecordAttributeValue(logs plog.Logs) {
	_ = "STUB: not implemented"
	return
}

func IgnoreTimestamp() CompareLogsOption { _ = "STUB: not implemented"; return *new(CompareLogsOption) }

func maskTimestamp(logs plog.Logs, ts pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func IgnoreObservedTimestamp() CompareLogsOption {
	_ = "STUB: not implemented"
	return *new(CompareLogsOption)
}

func maskObservedTimestamp(logs plog.Logs, ts pcommon.Timestamp) { _ = "STUB: not implemented"; return }

// IgnoreResourceLogsOrder is a CompareLogsOption that ignores the order of resource traces/metrics/logs.
func IgnoreResourceLogsOrder() CompareLogsOption {
	_ = "STUB: not implemented"
	return *new(CompareLogsOption)
}

func sortResourceLogsSlice(rls plog.ResourceLogsSlice) { _ = "STUB: not implemented"; return }

// IgnoreScopeLogsOrder is a CompareLogsOption that ignores the order of instrumentation scope traces/metrics/logs.
func IgnoreScopeLogsOrder() CompareLogsOption {
	_ = "STUB: not implemented"
	return *new(CompareLogsOption)
}

func sortScopeLogsSlices(ls plog.Logs) { _ = "STUB: not implemented"; return }

// IgnoreLogRecordsOrder is a CompareLogsOption that ignores the order of log records.
func IgnoreLogRecordsOrder() CompareLogsOption {
	_ = "STUB: not implemented"
	return *new(CompareLogsOption)
}

func sortLogRecordSlices(ls plog.Logs) { _ = "STUB: not implemented"; return }

// IgnoreScopeLogsVersion is a CompareLogsOption that ignores the version of scope logs.
func IgnoreScopeLogsVersion() CompareLogsOption {
	_ = "STUB: not implemented"
	return *new(CompareLogsOption)
}

func maskScopeLogsVersion(logs plog.Logs, version string) { _ = "STUB: not implemented"; return }
