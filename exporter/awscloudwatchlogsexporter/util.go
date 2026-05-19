// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscloudwatchlogsexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"

import (
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
	"InstanceId":           "service.instance.id",
	"FaasName":             "faas.name",
	"FaasVersion":          "faas.version",
}

func isPatternValid(s string) (bool, string) { _ = "STUB: not implemented"; return false, "" }

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

// getLogInfo retrieves the log group and log stream names from a given set of metrics.
func getLogInfo(resourceAttrs map[string]any, config *Config) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// Convert to map[string]string

// Override log group/stream if specified in config. However, in this case, customer won't have correlation experience

func anyMapToStringMap(resourceAttrs map[string]any) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
