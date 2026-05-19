// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package persistentvolume // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/persistentvolume"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

const (
	k8sPVCreationTime = "k8s.persistentvolume.creation_timestamp"
)

// Transform transforms the PersistentVolume to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new PersistentVolume fields.
func Transform(pv *corev1.PersistentVolume) *corev1.PersistentVolume {
	_ = "STUB: not implemented"
	return nil
}

func shouldSkipAnnotation(key string) bool { _ = "STUB: not implemented"; return false }

// RecordMetrics records metrics for the PersistentVolume.
func RecordMetrics(mb *metadata.MetricsBuilder, pv *corev1.PersistentVolume, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// GetMetadata returns the metadata for the PersistentVolume.
func GetMetadata(pv *corev1.PersistentVolume) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
