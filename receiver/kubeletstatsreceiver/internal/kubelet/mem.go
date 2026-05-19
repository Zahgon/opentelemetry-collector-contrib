// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

func addMemoryMetrics(
	mb *metadata.MetricsBuilder,
	memoryMetrics metadata.MemoryMetrics,
	s *stats.MemoryStats,
	currentTime pcommon.Timestamp,
	r resources,
	nodeMemoryLimit float64,
) {
	_ = "STUB: not implemented"
	return
}
