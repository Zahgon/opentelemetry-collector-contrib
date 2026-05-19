// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package container // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/container"

import (
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

const (
	// Keys for container metadata used for entity attributes.
	containerKeyStatus         = "container.status"
	containerKeyStatusReason   = "container.status.reason"
	containerCreationTimestamp = "container.creation_timestamp"
	containerName              = "k8s.container.name"
	containerImageName         = "container.image.name"
	containerImageTag          = "container.image.tag"

	// Values for container metadata
	containerStatusRunning    = "running"
	containerStatusWaiting    = "waiting"
	containerStatusTerminated = "terminated"
)

var allContainerStatusReasons = []metadata.AttributeK8sContainerStatusReason{
	metadata.AttributeK8sContainerStatusReasonContainerCreating,
	metadata.AttributeK8sContainerStatusReasonCrashLoopBackOff,
	metadata.AttributeK8sContainerStatusReasonCreateContainerConfigError,
	metadata.AttributeK8sContainerStatusReasonErrImagePull,
	metadata.AttributeK8sContainerStatusReasonImagePullBackOff,
	metadata.AttributeK8sContainerStatusReasonOOMKilled,
	metadata.AttributeK8sContainerStatusReasonCompleted,
	metadata.AttributeK8sContainerStatusReasonError,
	metadata.AttributeK8sContainerStatusReasonContainerCannotRun,
}

// RecordSpecMetrics metricizes values from the container spec.
// This includes values like resource requests and limits.
func RecordSpecMetrics(logger *zap.Logger, mb *metadata.MetricsBuilder, c corev1.Container, pod *corev1.Pod, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

//exhaustive:ignore

//exhaustive:ignore

// Record k8s.container.status.reason metric: for each known reason emit 1 for the current one, 0 otherwise.

// Emit in deterministic order for test stability.

func GetMetadata(pod *corev1.Pod, cs corev1.ContainerStatus, logger *zap.Logger) *metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}

func boolToInt64(b bool) int64 { _ = "STUB: not implemented"; return 0 }

var re = regexp.MustCompile(`^[\w_-]+://`)

// stripContainerID returns a pure container id without the runtime scheme://.
func stripContainerID(id string) string { _ = "STUB: not implemented"; return "" }
