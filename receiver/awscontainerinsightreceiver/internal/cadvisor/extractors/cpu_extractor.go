// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extractors // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"

import (
	cInfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"

	awsmetrics "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"
)

const (
	decimalToMillicores = 1000
)

type CPUMetricExtractor struct {
	logger         *zap.Logger
	rateCalculator awsmetrics.MetricCalculator
}

func (*CPUMetricExtractor) HasValue(info *cInfo.ContainerInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CPUMetricExtractor) GetValue(info *cInfo.ContainerInfo, mInfo CPUMemInfoProvider, containerType string) []*CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

// Skip infra container and handle node, pod, other containers in pod

// When there is more than one stats point, always use the last one

func (c *CPUMetricExtractor) Shutdown() error { _ = "STUB: not implemented"; return nil }

func NewCPUMetricExtractor(logger *zap.Logger) *CPUMetricExtractor {
	_ = "STUB: not implemented"
	return nil
}
