// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package deployment // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/deployment"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

// Transform transforms the pod to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new deployment fields.
func Transform(deployment *appsv1.Deployment) *appsv1.Deployment {
	_ = "STUB: not implemented"
	return nil
}

func RecordMetrics(mb *metadata.MetricsBuilder, dep *appsv1.Deployment, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func GetMetadata(dep *appsv1.Deployment) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
