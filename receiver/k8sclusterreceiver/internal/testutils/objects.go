// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/testutils"

import (
	quotav1 "github.com/openshift/api/quota/v1"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewService(id string) *corev1.Service { _ = "STUB: not implemented"; return nil }

func NewEndpointSlice(id string) *discoveryv1.EndpointSlice { _ = "STUB: not implemented"; return nil }

func NewLoadBalancerService(id string) *corev1.Service { _ = "STUB: not implemented"; return nil }

func NewHPA(id string) *autoscalingv2.HorizontalPodAutoscaler {
	_ = "STUB: not implemented"
	return nil
}

func NewJob(id string) *batchv1.Job { _ = "STUB: not implemented"; return nil }

func NewClusterResourceQuota(id string) *quotav1.ClusterResourceQuota {
	_ = "STUB: not implemented"
	return nil
}

func NewDaemonset(id string) *appsv1.DaemonSet { _ = "STUB: not implemented"; return nil }

func NewDeployment(id string) *appsv1.Deployment { _ = "STUB: not implemented"; return nil }

func NewReplicaSet(id string) *appsv1.ReplicaSet { _ = "STUB: not implemented"; return nil }

func NewNode(id string) *corev1.Node { _ = "STUB: not implemented"; return nil }

func NewPodWithContainer(id string, spec *corev1.PodSpec, status *corev1.PodStatus) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func NewPodSpecWithContainer(containerName string) *corev1.PodSpec {
	_ = "STUB: not implemented"
	return nil
}

func NewPodStatusWithContainer(containerName, containerID string) *corev1.PodStatus {
	_ = "STUB: not implemented"
	return nil
}

func NewEvictedTerminatedPodStatusWithContainer(containerName, containerID string) *corev1.PodStatus {
	_ = "STUB: not implemented"
	return nil
}

func WithOwnerReferences(or []v1.OwnerReference, obj any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func NewNamespace(id string) *corev1.Namespace { _ = "STUB: not implemented"; return nil }

func NewReplicationController(id string) *corev1.ReplicationController {
	_ = "STUB: not implemented"
	return nil
}

func NewResourceQuota(id string) *corev1.ResourceQuota { _ = "STUB: not implemented"; return nil }

func NewStatefulset(id string) *appsv1.StatefulSet { _ = "STUB: not implemented"; return nil }

func NewCronJob(id string) *batchv1.CronJob { _ = "STUB: not implemented"; return nil }

func NewPersistentVolume(id string) *corev1.PersistentVolume { _ = "STUB: not implemented"; return nil }

func NewPersistentVolumeClaim(id string) *corev1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return nil
}
