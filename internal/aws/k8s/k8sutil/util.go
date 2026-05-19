// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sutil"

// CreatePodKey concatenates namespace and podName to get a pod key
func CreatePodKey(namespace, podName string) string { _ = "STUB: not implemented"; return "" }

// CreateContainerKey concatenates namespace, podName and containerName to get a container key
func CreateContainerKey(namespace, podName, containerName string) string {
	_ = "STUB: not implemented"
	return ""
}
