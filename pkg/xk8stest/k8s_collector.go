// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package xk8stest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/xk8stest"

import (
	"context"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

func CreateCollectorObjects(t *testing.T, client *K8sClient, testID, manifestsDir string, templateValues map[string]string, host string) []*unstructured.Unstructured {
	_ = "STUB: not implemented"
	return nil
}

func WaitForCollectorToStart(t *testing.T, client *K8sClient, podNamespace string, podLabels map[string]any) {
	_ = "STUB: not implemented"
	return
}

func collectorPodsReady(t *testing.T, client *K8sClient, namespace string, podGVR schema.GroupVersionResource, listOptions metav1.ListOptions) bool {
	_ = "STUB: not implemented"
	return false
}

func logCollectorPodDiagnostics(t *testing.T, client *K8sClient, namespace string, podGVR schema.GroupVersionResource, listOptions metav1.ListOptions) {
	_ = "STUB: not implemented"
	return
}

func podReady(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }

func logPodEvents(t *testing.T, client *K8sClient, namespace, podName string) {
	_ = "STUB: not implemented"
	return
}

func logRestartingContainers(t *testing.T, client *K8sClient, namespace, podName string, statuses []v1.ContainerStatus) {
	_ = "STUB: not implemented"
	return
}

func logContainerLogs(t *testing.T, client *K8sClient, namespace, podName, containerName string) {
	_ = "STUB: not implemented"
	return
}

func fetchContainerLogs(ctx context.Context, coreClient corev1client.CoreV1Interface, namespace, podName, containerName string, previous bool, tailLines *int64) string {
	_ = "STUB: not implemented"
	return ""
}

// FetchPodLogs returns the full log output for the first running pod matching
// the given labels in the specified namespace. It is intended for e2e test
// assertions that inspect collector behavior via its log output.
func FetchPodLogs(t *testing.T, client *K8sClient, namespace string, podLabels map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}
