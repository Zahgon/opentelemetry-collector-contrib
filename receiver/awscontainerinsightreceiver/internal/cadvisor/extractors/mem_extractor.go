// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extractors // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"

import (
	cinfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"

	awsmetrics "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"
)

type MemMetricExtractor struct {
	logger         *zap.Logger
	rateCalculator awsmetrics.MetricCalculator
}

func (*MemMetricExtractor) HasValue(info *cinfo.ContainerInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *MemMetricExtractor) GetValue(info *cinfo.ContainerInfo, mInfo CPUMemInfoProvider, containerType string) []*CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

func (m *MemMetricExtractor) Shutdown() error { _ = "STUB: not implemented"; return nil }

func NewMemMetricExtractor(logger *zap.Logger) *MemMetricExtractor {
	_ = "STUB: not implemented"
	return nil
}
