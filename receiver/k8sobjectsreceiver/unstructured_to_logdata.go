// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobjectsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/watch"
)

type attrUpdaterFunc func(pcommon.Map)

func watchObjectsToLogData(event *watch.Event, observedAt time.Time, config *K8sObjectsConfig, version string) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func pullObjectsToLogData(event *unstructured.UnstructuredList, observedAt time.Time, config *K8sObjectsConfig, version string) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func unstructuredListToLogData(event *unstructured.UnstructuredList, observedAt time.Time, config *K8sObjectsConfig, version string, attrUpdaters ...attrUpdaterFunc) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

//nolint:errcheck

func getNamespace(e unstructured.Unstructured) string {
	_ = "STUB: not implemented"
	// first, try to use the GetNamespace() method, which checks for the metadata.namespace property
	return ""
}

// try to look up namespace in object.metadata.namespace (for objects reported via watch mode)
