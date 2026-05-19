// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsecscontainermetrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver/internal/awsecscontainermetrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"
)

// metricDataAccumulator defines the accumulator
type metricDataAccumulator struct {
	mds []pmetric.Metrics
}

// getMetricsData generates OT Metrics data from task metadata and docker stats
func (acc *metricDataAccumulator) getMetricsData(containerStatsMap map[string]*ContainerStats, metadata ecsutil.TaskMetadata, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

func (acc *metricDataAccumulator) accumulate(md pmetric.Metrics) { _ = "STUB: not implemented"; return }

func isEmptyStats(stats *ContainerStats) bool { _ = "STUB: not implemented"; return false }

func convertContainerMetrics(stats *ContainerStats, logger *zap.Logger, containerMetadata *ecsutil.ContainerMetadata) ECSMetrics {
	_ = "STUB: not implemented"
	return *new(ECSMetrics)
}

func overrideWithTaskLevelLimit(taskMetrics *ECSMetrics, metadata ecsutil.TaskMetadata) {
	_ = "STUB: not implemented"
	// Overwrite Memory limit with task level limit
	return
}

// Overwrite CPU limit with task level limit

// taskMetrics.CPUReserved cannot be zero. In ECS, user needs to set CPU limit
// at least in one place (either in task level or in container level). If the
// task level CPULimit is not present, we calculate it from the summation of
// all container CPU limits.

func calculateDuration(startTime, endTime string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
