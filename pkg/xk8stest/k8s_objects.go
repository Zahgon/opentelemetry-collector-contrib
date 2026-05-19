// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package xk8stest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/xk8stest"

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func CreateObject(client *K8sClient, manifest []byte) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cluster-scoped resources

func DeleteObject(client *K8sClient, obj *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

// cluster-scoped resources

func CreateObjects(client *K8sClient, dir string) ([]*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip directories

func DeleteObjects(client *K8sClient, objs []*unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}
