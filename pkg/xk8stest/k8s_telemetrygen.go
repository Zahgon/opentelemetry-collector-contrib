// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package xk8stest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/xk8stest"

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type TelemetrygenObjInfo struct {
	Namespace         string
	PodLabelSelectors map[string]any
	DataType          string
	Workload          string
}

type TelemetrygenCreateOpts struct {
	TestID       string
	ManifestsDir string
	OtlpEndpoint string
	DataTypes    []string
}

// getPodLabelSelectors returns labels used to select pods created by the workload.
// - Deployment/StatefulSet/DaemonSet: spec.selector.matchLabels (fallback to template.metadata.labels)
// - Job: spec.template.metadata.labels
// - CronJob: spec.jobTemplate.spec.template.metadata.labels
func getPodLabelSelectors(obj *unstructured.Unstructured) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fallback — uncommon but robust

// last resort if API server already defaulted it

func CreateTelemetryGenObjects(t *testing.T, client *K8sClient, createOpts *TelemetrygenCreateOpts) ([]*unstructured.Unstructured, []*TelemetrygenObjInfo) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WaitForTelemetryGenToStart(t *testing.T, client *K8sClient, podNamespace string, podLabels map[string]any, workload, dataType string) {
	_ = "STUB: not implemented"
	return
}

func telemetrygenPodsStatus(t *testing.T, client *K8sClient, podNamespace string, listOptions metav1.ListOptions) (bool, bool, string) {
	_ = "STUB: not implemented"
	return false, false, ""
}

func podWaitMessage(pod *v1.Pod) string { _ = "STUB: not implemented"; return "" }
