// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extractors // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"

import (
	cinfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"

	awsmetrics "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"
)

type NetMetricExtractor struct {
	logger         *zap.Logger
	rateCalculator awsmetrics.MetricCalculator
}

func getInterfacesStats(stats *cinfo.ContainerStats) []cinfo.InterfaceStats {
	_ = "STUB: not implemented"
	return nil
}

func (*NetMetricExtractor) HasValue(info *cinfo.ContainerInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *NetMetricExtractor) GetValue(info *cinfo.ContainerInfo, _ CPUMemInfoProvider, containerType string) []*CAdvisorMetric {
	_ = "STUB: not implemented"
	// Just a protection here, there is no Container level Net metrics
	return nil
}

// Rename type to pod so the metric name prefix is pod_

// used for aggregation

// used to identify the network interface

func (n *NetMetricExtractor) Shutdown() error { _ = "STUB: not implemented"; return nil }

func NewNetMetricExtractor(logger *zap.Logger) *NetMetricExtractor {
	_ = "STUB: not implemented"
	return nil
}

func getNetMetricType(containerType string, logger *zap.Logger) string {
	_ = "STUB: not implemented"
	return ""
}
