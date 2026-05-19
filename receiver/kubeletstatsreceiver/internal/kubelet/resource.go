// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

func getContainerResource(rb *metadata.ResourceBuilder, sPod *stats.PodStats, sContainer *stats.ContainerStats,
	k8sMetadata Metadata,
) (pcommon.Resource, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), nil
}

func getVolumeResourceOptions(rb *metadata.ResourceBuilder, sPod *stats.PodStats, vs *stats.VolumeStats,
	k8sMetadata Metadata,
) (pcommon.Resource, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), nil
}
