// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

func addCPUMetrics(
	mb *metadata.MetricsBuilder,
	cpuMetrics metadata.CPUMetrics,
	s *stats.CPUStats,
	currentTime pcommon.Timestamp,
	r resources,
	nodeCPULimit float64,
) {
	_ = "STUB: not implemented"
	return
}

func addCPUUtilizationMetrics(
	mb *metadata.MetricsBuilder,
	cpuMetrics metadata.CPUMetrics,
	usageCores float64,
	currentTime pcommon.Timestamp,
	r resources,
	nodeCPULimit float64,
) {
	_ = "STUB: not implemented"
	return
}

func addCPUTimeMetric(mb *metadata.MetricsBuilder, recordDataPoint metadata.RecordDoubleDataPointFunc, s *stats.CPUStats, currentTime pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}
