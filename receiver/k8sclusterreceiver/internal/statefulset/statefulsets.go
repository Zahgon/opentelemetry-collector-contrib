// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statefulset // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/statefulset"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

const (
	// Keys for stateful set metadata.
	statefulSetCurrentVersion = "current_revision"
	statefulSetUpdateVersion  = "update_revision"
)

// Transform transforms the pod to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new statefulset fields.
func Transform(statefulset *appsv1.StatefulSet) *appsv1.StatefulSet {
	_ = "STUB: not implemented"
	return nil
}

func RecordMetrics(mb *metadata.MetricsBuilder, ss *appsv1.StatefulSet, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func GetMetadata(ss *appsv1.StatefulSet) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
