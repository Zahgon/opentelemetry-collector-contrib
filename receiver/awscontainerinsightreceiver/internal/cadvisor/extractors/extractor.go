// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extractors // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"

import (
	"time"

	cinfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"

	awsmetrics "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"
)

func GetStats(info *cinfo.ContainerInfo) *cinfo.ContainerStats {
	_ = "STUB: not implemented"
	return nil
}

// When there is more than one stats point, always use the last one

type CPUMemInfoProvider interface {
	GetNumCores() int64
	GetMemoryCapacity() int64
}

type MetricExtractor interface {
	HasValue(*cinfo.ContainerInfo) bool
	GetValue(info *cinfo.ContainerInfo, mInfo CPUMemInfoProvider, containerType string) []*CAdvisorMetric
	Shutdown() error
}

type CAdvisorMetric struct {
	// source of the metric for debugging merge conflict
	cgroupPath string
	// key/value pairs that are typed and contain the metric (numerical) data
	fields map[string]any
	// key/value string pairs that are used to identify the metrics
	tags map[string]string

	logger *zap.Logger
}

func newCadvisorMetric(mType string, logger *zap.Logger) *CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

func (c *CAdvisorMetric) GetTags() map[string]string { _ = "STUB: not implemented"; return nil }

func (c *CAdvisorMetric) GetFields() map[string]any { _ = "STUB: not implemented"; return nil }

func (c *CAdvisorMetric) GetMetricType() string { _ = "STUB: not implemented"; return "" }

func (c *CAdvisorMetric) AddTags(tags map[string]string) { _ = "STUB: not implemented"; return }

func (c *CAdvisorMetric) HasField(key string) bool { _ = "STUB: not implemented"; return false }

func (c *CAdvisorMetric) AddField(key string, val any) { _ = "STUB: not implemented"; return }

func (c *CAdvisorMetric) GetField(key string) any { _ = "STUB: not implemented"; return *new(any) }

func (c *CAdvisorMetric) HasTag(key string) bool { _ = "STUB: not implemented"; return false }

func (c *CAdvisorMetric) AddTag(key, val string) { _ = "STUB: not implemented"; return }

func (c *CAdvisorMetric) GetTag(key string) string { _ = "STUB: not implemented"; return "" }

func (c *CAdvisorMetric) RemoveTag(key string) { _ = "STUB: not implemented"; return }

func (c *CAdvisorMetric) Merge(src *CAdvisorMetric) {
	_ = "STUB: not implemented"
	// If there is any conflict, keep the fields with earlier timestamp
	return
}

func newFloat64RateCalculator() awsmetrics.MetricCalculator {
	_ = "STUB: not implemented"
	return *new(awsmetrics.MetricCalculator)
}

func assignRateValueToField(rateCalculator *awsmetrics.MetricCalculator, fields map[string]any, metricName string,
	cinfoName string, curVal any, curTime time.Time, multiplier float64,
) {
	_ = "STUB: not implemented"
	return
}

// MergeMetrics merges an array of cadvisor metrics based on common metric keys
func MergeMetrics(metrics []*CAdvisorMetric) []*CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

// this metric cannot be merged

// return MetricKey for merge-able metrics
func getMetricKey(metric *CAdvisorMetric) string { _ = "STUB: not implemented"; return "" }

// merge cpu, memory, net metric for type Instance

// merge cpu, memory, net metric for type Node

// merge cpu, memory, net metric for type Pod

// merge cpu, memory metric for type Container

// merge io_serviced, io_service_bytes for type InstanceDiskIO

// merge io_serviced, io_service_bytes for type NodeDiskIO
