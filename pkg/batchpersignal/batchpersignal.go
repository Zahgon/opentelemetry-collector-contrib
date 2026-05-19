// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batchpersignal // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// SplitTraces returns one ptrace.Traces for each trace in the given ptrace.Traces input. Each of the resulting ptrace.Traces contains exactly one trace.
func SplitTraces(batch ptrace.Traces) []ptrace.Traces {
	_ = "STUB: not implemented"
	// for each span in the resource spans, we group them into batches of rs/ils/traceID.
	// if the same traceID exists in different ils, they land in different batches.
	return nil
}

// the batches for this ILS

// for the first traceID in the ILS, initialize the map entry
// and add the singleTraceBatch to the result list

// currently, the ResourceSpans implementation has only a Resource and an ILS. We'll copy the Resource
// and set our own ILS

// currently, the ILS implementation has only an InstrumentationLibrary and spans. We'll copy the library
// and set our own spans

// there is only one instrumentation library per batch

// SplitLogs returns one plog.Logs for each trace in the given plog.Logs input. Each of the resulting plog.Logs contains exactly one log.
func SplitLogs(batch plog.Logs) []plog.Logs {
	_ = "STUB: not implemented"
	// for each log in the resource logs, we group them into batches of rl/sl/traceID.
	// if the same traceID exists in different sl, they land in different batches.
	return nil
}

// the batches for this ILL

// for the first traceID in the ILL, initialize the map entry
// and add the singleTraceBatch to the result list

// currently, the ResourceLogs implementation has only a Resource and an ILL. We'll copy the Resource
// and set our own ILL

// currently, the ILL implementation has only an InstrumentationLibrary and logs. We'll copy the library
// and set our own logs

// there is only one instrumentation library per batch

// SplitMetrics returns one pmetric.Metrics for each metric in the given pmetric.Metrics input. Each of the resulting pmetric.Metrics contains exactly one metric.
func SplitMetrics(batch pmetric.Metrics) []pmetric.Metrics {
	_ = "STUB: not implemented"
	// for each label in the resource labels, we group them into batches of rs/ils/metricName.
	// if the same metricName exists in different ils, they land in different batches.
	return nil
}

// the batches for this ILS

// key := pcommon.NewByteSlice()
// key.FromRaw([]byte(metric.Name()))

// for the first metric in the ILS, initialize the map entry
// and add the singleMetricBatch to the result list

// currently, the ResourceMetrics implementation has only a Resource and an ILS. We'll copy the Resource
// and set our own ILS

// currently, the ILS implementation has only an InstrumentationLibrary and metrics. We'll copy the library
// and set our own metrics

// there is only one instrumentation library per batch
