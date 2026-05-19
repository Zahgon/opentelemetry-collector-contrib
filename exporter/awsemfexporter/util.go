// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsemfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var patternKeyToAttributeMap = map[string]string{
	"ClusterName":          "aws.ecs.cluster.name",
	"TaskId":               "aws.ecs.task.id",
	"NodeName":             "k8s.node.name",
	"PodName":              "pod",
	"ServiceName":          "service.name",
	"ContainerInstanceId":  "aws.ecs.container.instance.id",
	"TaskDefinitionFamily": "aws.ecs.task.family",
}

func replacePatterns(s string, attrMap map[string]string, logger *zap.Logger) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func replacePatternWithAttrValue(s, patternKey string, attrMap map[string]string, logger *zap.Logger) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func replace(s, pattern, value string, logger *zap.Logger) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// getNamespace retrieves namespace for given set of metrics from user config.
func getNamespace(rm pmetric.ResourceMetrics, namespace string) string {
	_ = "STUB: not implemented"
	return ""
}

// getLogInfo retrieves the log group and log stream names from a given set of metrics.
func getLogInfo(rm pmetric.ResourceMetrics, cWNamespace string, config *Config) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// Override log group/stream if specified in config. However, in this case, customer won't have correlation experience

// dedupDimensions removes duplicated dimension sets from the given dimensions.
// Prerequisite: each dimension set is already sorted
func dedupDimensions(dimensions [][]string) (deduped [][]string) {
	_ = "STUB: not implemented"
	return nil
}

// Only add dimension set if not a duplicate

// dimensionRollup creates rolled-up dimensions from the metric's label set.
// The returned dimensions are sorted in alphabetical order within each dimension set
func dimensionRollup(dimensionRollupOption string, labels map[string]string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// Empty dimension must be always present in a roll up.

// If OTel key exists in labels, add it as a zero dimension but remove it
// temporarily from labels as it is not an original label

// "Zero" dimension rollup

// "One" dimension rollup

// Add back OTel key to labels if it was removed

// unixNanoToMilliseconds converts a timestamp in nanoseconds to milliseconds.
func unixNanoToMilliseconds(timestamp pcommon.Timestamp) int64 { _ = "STUB: not implemented"; return 0 }

// attrMaptoStringMap converts a pcommon.Map to a map[string]string
func attrMaptoStringMap(attrMap pcommon.Map) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
