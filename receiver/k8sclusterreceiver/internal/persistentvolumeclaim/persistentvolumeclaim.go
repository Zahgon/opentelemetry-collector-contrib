// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package persistentvolumeclaim // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/persistentvolumeclaim"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

const (
	k8sPVCCreationTime = "k8s.persistentvolumeclaim.creation_timestamp"
	k8sBoundPVName     = "k8s.persistentvolume.name"
)

// Transform transforms the PersistentVolumeClaim to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new PersistentVolumeClaim fields.
func Transform(pvc *corev1.PersistentVolumeClaim) *corev1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return nil
}

func shouldSkipAnnotation(key string) bool { _ = "STUB: not implemented"; return false }

// RecordMetrics records metrics for the PersistentVolumeClaim.
func RecordMetrics(mb *metadata.MetricsBuilder, pvc *corev1.PersistentVolumeClaim, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// GetMetadata returns the metadata for the PersistentVolumeClaim.
func GetMetadata(pvc *corev1.PersistentVolumeClaim) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
