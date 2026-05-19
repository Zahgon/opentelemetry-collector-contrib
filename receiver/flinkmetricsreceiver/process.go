// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package flinkmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver/internal/models"
)

func (s *flinkmetricsScraper) processJobmanagerMetrics(now pcommon.Timestamp, jobmanagerMetrics *models.JobmanagerMetrics) {
	_ = "STUB: not implemented"
	return
}

func (s *flinkmetricsScraper) processTaskmanagerMetrics(now pcommon.Timestamp, taskmanagerMetricInstances []*models.TaskmanagerMetrics) {
	_ = "STUB: not implemented"
	return
}

func (s *flinkmetricsScraper) processJobsMetrics(now pcommon.Timestamp, jobsMetricsInstances []*models.JobMetrics) {
	_ = "STUB: not implemented"
	return
}

func (s *flinkmetricsScraper) processSubtaskMetrics(now pcommon.Timestamp, subtaskMetricsInstances []*models.SubtaskMetrics) {
	_ = "STUB: not implemented"
	return
}

// record task metrics

// record operator metrics
