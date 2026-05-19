// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package node // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/node"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

const (
	// Keys for node metadata and entity attributes. These are NOT used by resource attributes.
	nodeCreationTime       = "node.creation_timestamp"
	k8sNodeConditionPrefix = "k8s.node.condition"
)

// Transform transforms the node to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new node fields.
func Transform(node *corev1.Node) *corev1.Node { _ = "STUB: not implemented"; return nil }

func RecordMetrics(mb *metadata.MetricsBuilder, node *corev1.Node, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func CustomMetrics(set receiver.Settings, rb *metadata.ResourceBuilder, node *corev1.Node, nodeConditionTypesToReport,
	allocatableTypesToReport []string, ts pcommon.Timestamp,
) pmetric.ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ResourceMetrics)
}

// Adding 'node condition type' metrics

// Adding 'node allocatable type' metrics

// TODO: Generate a schema URL for the node metrics in the metadata package and use them here.

var nodeConditionValues = map[corev1.ConditionStatus]int64{
	corev1.ConditionTrue:    1,
	corev1.ConditionFalse:   0,
	corev1.ConditionUnknown: -1,
}

func nodeConditionValue(node *corev1.Node, condType corev1.NodeConditionType) int64 {
	_ = "STUB: not implemented"
	return 0
}

func GetMetadata(node *corev1.Node) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}

// Node can have many additional conditions (gke has 18 on v1.29). Bad thresholds/implementations
// of custom conditions can cause value to oscillate between true/false frequently. So, only sending the node
// pressure conditions that are set by kubelet to avoid noise.
// https://pkg.go.dev/k8s.io/api/core/v1#NodeConditionType

func getContainerRuntimeInfo(rawInfo string) (runtime, version string) {
	_ = "STUB: not implemented"
	// Kubelet reports container runtime version in the following format:
	// <runtime-name>://<version>
	return "", ""
}

func getNodeConditionMetric(nodeConditionTypeValue string) string {
	_ = "STUB: not implemented"
	return ""
}

func getNodeAllocatableUnit(res corev1.ResourceName) string { _ = "STUB: not implemented"; return "" }

func setNodeAllocatableValue(dp pmetric.NumberDataPoint, res corev1.ResourceName, q resource.Quantity) {
	_ = "STUB: not implemented"
	return
}

func getNodeAllocatableMetric(nodeAllocatableTypeValue string) string {
	_ = "STUB: not implemented"
	return ""
}
