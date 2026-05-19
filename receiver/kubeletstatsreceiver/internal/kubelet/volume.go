// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	v1 "k8s.io/api/core/v1"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

func addVolumeMetrics(mb *metadata.MetricsBuilder, volumeMetrics metadata.VolumeMetrics, s *stats.VolumeStats, currentTime pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func setResourcesFromVolume(rb *metadata.ResourceBuilder, volume *v1.Volume) {
	_ = "STUB: not implemented"

	// TODO: Support more types
	return
}

func SetPersistentVolumeLabels(rb *metadata.ResourceBuilder, pv v1.PersistentVolumeSource) {
	_ = "STUB: not implemented"
	// TODO: Support more types
	return
}

// pv.Glusterfs is a GlusterfsPersistentVolumeSource instead of GlusterfsVolumeSource,
// convert to GlusterfsVolumeSource so a single method can handle both structs. This
// can be broken out into separate methods if one is interested in different sets
// of labels from the two structs in the future.

func awsElasticBlockStoreDims(rb *metadata.ResourceBuilder, vs v1.AWSElasticBlockStoreVolumeSource) {
	_ = "STUB: not implemented"
	return
}

// AWS specific labels.

func gcePersistentDiskDims(rb *metadata.ResourceBuilder, vs v1.GCEPersistentDiskVolumeSource) {
	_ = "STUB: not implemented"
	return
}

// GCP specific labels.

func glusterfsDims(rb *metadata.ResourceBuilder, vs v1.GlusterfsVolumeSource) {
	_ = "STUB: not implemented"
	return
}

// GlusterFS specific labels.
