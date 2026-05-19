// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver/internal/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	common "skywalking.apache.org/repo/goapi/collect/common/v3"
	agent "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const (
	jvmScopeName            = "runtime_metrics"
	gcCountMetricName       = "sw.jvm.gc.count"
	gcDurationMetricName    = "sw.jvm.gc.duration"
	MemoryPoolInitName      = "jvm.memory.init"
	MemoryPoolMaxName       = "jvm.memory.max"
	MemoryPoolUsedName      = "jvm.memory.used"
	MemoryPoolCommittedName = "jvm.memory.committed"
	ThreadCountName         = "jvm.thread.count"
	CPUUtilizationName      = "jvm.cpu.recent_utilization"
)

func SwMetricsToMetrics(collection *agent.JVMMetricCollection) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func jvmMetricToResource(serviceName, serviceInstance string, resource pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

func jvmMetricToResourceMetrics(jvmMetric *agent.JVMMetric, sm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	// gc metric to otlp metric
	return
}

// memory pool metric to otlp metric

// thread metric to otlp metric

// cpu metric to otpl metric

// gcMetricToMetrics translate gc metrics
func gcMetricToMetrics(timestamp int64, gcList []*agent.GC, dest pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	// gc count and gc duration is not
	return
}

func buildGCAttrs(gc *agent.GC) pcommon.Map { _ = "STUB: not implemented"; return *new(pcommon.Map) }

// memoryPoolMetricToMetrics translate memoryPool metrics
func memoryPoolMetricToMetrics(timestamp int64, memoryPools []*agent.MemoryPool, sm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func buildMemoryPoolAttrs(pool *agent.MemoryPool) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func threadMetricToMetrics(timestamp int64, thread *agent.Thread, dest pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func buildThreadTypeAttrs(typeValue string) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func cpuMetricToMetrics(timestamp int64, cpu *common.CPU, dest pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func fillNumberDataPointIntValue(timestamp, value int64, point pmetric.NumberDataPoint, attrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func fillNumberDataPointDoubleValue(timestamp int64, value float64, point pmetric.NumberDataPoint, attrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}
