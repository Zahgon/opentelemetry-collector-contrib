// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsemfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// groupedMetric defines set of metrics with same namespace, timestamp and labels
type groupedMetric struct {
	labels   map[string]string
	metrics  map[string]*metricInfo
	metadata cWMetricMetadata
}

// metricInfo defines value and unit for OT Metrics
type metricInfo struct {
	value any
	unit  string
}

// addToGroupedMetric processes OT metrics and adds them into GroupedMetric buckets
func addToGroupedMetric(
	pmd pmetric.Metric,
	groupedMetrics map[any]*groupedMetric,
	metadata cWMetricMetadata,
	patternReplaceSucceeded bool,
	descriptor map[string]MetricDescriptor,
	config *Config,
	calculators *emfCalculators,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Drop stale or NaN metric values

// if patterns were found in config file and weren't replaced by resource attributes, replace those patterns with metric labels.
// if patterns are provided for a valid key and that key doesn't exist in the resource attributes, it is replaced with `undefined`.

// Extra params to use when grouping metrics

// if MetricName already exists in metrics map, print warning log

type kubernetesObj struct {
	ContainerName string                `json:"container_name,omitempty"`
	Docker        *internalDockerObj    `json:"docker,omitempty"`
	Host          string                `json:"host,omitempty"`
	Labels        *internalLabelsObj    `json:"labels,omitempty"`
	NamespaceName string                `json:"namespace_name,omitempty"`
	PodID         string                `json:"pod_id,omitempty"`
	PodName       string                `json:"pod_name,omitempty"`
	PodOwners     *internalPodOwnersObj `json:"pod_owners,omitempty"`
	ServiceName   string                `json:"service_name,omitempty"`
}

type internalDockerObj struct {
	ContainerID string `json:"container_id,omitempty"`
}

type internalLabelsObj struct {
	App             string `json:"app,omitempty"`
	PodTemplateHash string `json:"pod-template-hash,omitempty"`
}

type internalPodOwnersObj struct {
	OwnerKind string `json:"owner_kind,omitempty"`
	OwnerName string `json:"owner_name,omitempty"`
}

func addKubernetesWrapper(labels map[string]string) {
	_ = "STUB: not implemented"
	// fill in obj
	return
}

// handle nested empty object

func mapGetHelper(labels map[string]string, key string) string {
	_ = "STUB: not implemented"
	return ""
}

func translateUnit(metric pmetric.Metric, descriptor map[string]MetricDescriptor) string {
	_ = "STUB: not implemented"
	return ""
}

// CloudWatch doesn't support Nanoseconds
