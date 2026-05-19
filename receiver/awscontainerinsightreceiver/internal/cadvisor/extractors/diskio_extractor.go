// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extractors // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"

import (
	"time"

	cInfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"

	awsmetrics "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"
)

type DiskIOMetricExtractor struct {
	logger         *zap.Logger
	rateCalculator awsmetrics.MetricCalculator
}

func (*DiskIOMetricExtractor) HasValue(info *cInfo.ContainerInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DiskIOMetricExtractor) GetValue(info *cInfo.ContainerInfo, _ CPUMemInfoProvider, containerType string) []*CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

func (d *DiskIOMetricExtractor) extractIoMetrics(curStatsSet []cInfo.PerDiskStats, namePrefix, containerType, infoName string, curTime time.Time) []*CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

func (d *DiskIOMetricExtractor) Shutdown() error { _ = "STUB: not implemented"; return nil }

func ioMetricName(prefix, key string) string { _ = "STUB: not implemented"; return "" }

func devName(dStats cInfo.PerDiskStats) string { _ = "STUB: not implemented"; return "" }

func NewDiskIOMetricExtractor(logger *zap.Logger) *DiskIOMetricExtractor {
	_ = "STUB: not implemented"
	return nil
}

func getDiskIOMetricType(containerType string, logger *zap.Logger) string {
	_ = "STUB: not implemented"
	return ""
}
