// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package groupbyattrsprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type tracesGroup struct {
	traces         ptrace.Traces
	resourceHashes [][16]byte
}

func newTracesGroup() *tracesGroup { _ = "STUB: not implemented"; return nil }

// findOrCreateResource searches for a Resource with matching attributes and returns it. If nothing is found, it is being created
func (tg *tracesGroup) findOrCreateResourceSpans(originResource pcommon.Resource, requiredAttributes pcommon.Map) ptrace.ResourceSpans {
	_ = "STUB: not implemented"
	return *new(ptrace.ResourceSpans)
}

type metricsGroup struct {
	metrics        pmetric.Metrics
	resourceHashes [][16]byte
}

func newMetricsGroup() *metricsGroup { _ = "STUB: not implemented"; return nil }

// findOrCreateResourceMetrics searches for a Resource with matching attributes and returns it. If nothing is found, it is being created
func (mg *metricsGroup) findOrCreateResourceMetrics(originResource pcommon.Resource, requiredAttributes pcommon.Map) pmetric.ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ResourceMetrics)
}

type logsGroup struct {
	logs           plog.Logs
	resourceHashes [][16]byte
}

// newLogsGroup returns new logsGroup with predefined capacity
func newLogsGroup() *logsGroup { _ = "STUB: not implemented"; return nil }

// findOrCreateResourceLogs searches for a Resource with matching attributes and returns it. If nothing is found, it is being created
func (lg *logsGroup) findOrCreateResourceLogs(originResource pcommon.Resource, requiredAttributes pcommon.Map) plog.ResourceLogs {
	_ = "STUB: not implemented"
	return *new(plog.ResourceLogs)
}

func instrumentationLibrariesEqual(il1, il2 pcommon.InstrumentationScope) bool {
	_ = "STUB: not implemented"
	return false
}

// matchingScopeSpans searches for a ptrace.ScopeSpans instance matching
// given InstrumentationScope. If nothing is found, it creates a new one
func matchingScopeSpans(rl ptrace.ResourceSpans, library pcommon.InstrumentationScope) ptrace.ScopeSpans {
	_ = "STUB: not implemented"
	return *new(ptrace.ScopeSpans)
}

// matchingScopeLogs searches for a plog.ScopeLogs instance matching
// given InstrumentationScope. If nothing is found, it creates a new one
func matchingScopeLogs(rl plog.ResourceLogs, library pcommon.InstrumentationScope) plog.ScopeLogs {
	_ = "STUB: not implemented"
	return *new(plog.ScopeLogs)
}

// matchingScopeMetrics searches for a pmetric.ScopeMetrics instance matching
// given InstrumentationScope. If nothing is found, it creates a new one
func matchingScopeMetrics(rm pmetric.ResourceMetrics, library pcommon.InstrumentationScope) pmetric.ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ScopeMetrics)
}

// buildReferenceResource returns a new resource that we'll be looking for in existing Resources
// as a merge of the Attributes of the original Resource with the requested Attributes.
func buildReferenceResource(originResource pcommon.Resource, requiredAttributes pcommon.Map) pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}
