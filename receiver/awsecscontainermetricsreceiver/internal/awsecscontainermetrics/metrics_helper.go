// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsecscontainermetrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver/internal/awsecscontainermetrics"

import (
	"go.uber.org/zap"
)

// getContainerMetrics generate ECS Container metrics from Container stats
func getContainerMetrics(stats *ContainerStats, logger *zap.Logger) ECSMetrics {
	_ = "STUB: not implemented"
	return *new(ECSMetrics)
}

// Followed ECS Agent calculations
// https://github.com/aws/amazon-ecs-agent/blob/1ebf0604c13013596cfd4eb239574a85890b13e8/agent/stats/utils.go#L30
func getNetworkStats(stats map[string]NetworkStats) [8]uint64 {
	_ = "STUB: not implemented"
	return nil
}

// Followed ECS Agent calculations
// https://github.com/aws/amazon-ecs-agent/blob/1ebf0604c13013596cfd4eb239574a85890b13e8/agent/stats/utils_unix.go#L48
func extractStorageUsage(stats *DiskStats) (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

// ignoring "Async", "Total", "Sum", etc

func aggregateTaskMetrics(taskMetrics *ECSMetrics, conMetrics ECSMetrics) {
	_ = "STUB: not implemented"
	return
}
