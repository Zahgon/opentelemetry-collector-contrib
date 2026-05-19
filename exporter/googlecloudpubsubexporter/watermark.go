// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type (
	metricsWatermarkFunc func(metrics pmetric.Metrics, processingTime time.Time, allowedDrift time.Duration) time.Time
	logsWatermarkFunc    func(logs plog.Logs, processingTime time.Time, allowedDrift time.Duration) time.Time
	tracesWatermarkFunc  func(traces ptrace.Traces, processingTime time.Time, allowedDrift time.Duration) time.Time
)

type collectFunc func(timestamp pcommon.Timestamp) bool

// collector helps traverse the OTLP tree to calculate the final time to set to the ce-time attribute
type collector struct {
	// the current system clock time, set at the start of the tree traversal
	processingTime time.Time
	// maximum allowed difference for the processingTime
	allowedDrift time.Duration
	// calculated time, that can be set each time a timestamp is given to a calculation function
	calculatedTime time.Time
}

// add a new timestamp, and set the calculated time if it's earlier then the current calculated,
// taking into account the allowedDrift
func (c *collector) earliest(timestamp pcommon.Timestamp) bool {
	_ = "STUB: not implemented"
	return false
}

// function that doesn't traverse the metric data, return the processingTime
func currentMetricsWatermark(_ pmetric.Metrics, processingTime time.Time, _ time.Duration) time.Time {
	_ = "STUB: not implemented"
	return *

	// function that traverse the metric data, and returns the earliest timestamp (within limits of the allowedDrift)
	new(time.Time)
}

func earliestMetricsWatermark(metrics pmetric.Metrics, processingTime time.Time, allowedDrift time.Duration) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// traverse the metric data, with a collectFunc
func traverseMetrics(metrics pmetric.Metrics, collect collectFunc) {
	_ = "STUB: not implemented"
	return
}

//exhaustive:enforce

// function that doesn't traverse the log data, return the processingTime
func currentLogsWatermark(_ plog.Logs, processingTime time.Time, _ time.Duration) time.Time {
	_ = "STUB: not implemented"
	return *

	// function that traverse the log data, and returns the earliest timestamp (within limits of the allowedDrift)
	new(time.Time)
}

func earliestLogsWatermark(logs plog.Logs, processingTime time.Time, allowedDrift time.Duration) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// traverse the log data, with a collectFunc
func traverseLogs(logs plog.Logs, collect collectFunc) { _ = "STUB: not implemented"; return }

// function that doesn't traverse the trace data, return the processingTime
func currentTracesWatermark(_ ptrace.Traces, processingTime time.Time, _ time.Duration) time.Time {
	_ = "STUB: not implemented"
	return *

	// function that traverse the trace data, and returns the earliest timestamp (within limits of the allowedDrift)
	new(time.Time)
}

func earliestTracesWatermark(traces ptrace.Traces, processingTime time.Time, allowedDrift time.Duration) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// traverse the trace data, with a collectFunc
func traverseTraces(traces ptrace.Traces, collect collectFunc) { _ = "STUB: not implemented"; return }
