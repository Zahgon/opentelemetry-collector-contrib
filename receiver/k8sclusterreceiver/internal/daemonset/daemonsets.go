// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package daemonset // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/daemonset"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

// Transform transforms the pod to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new daemonset fields.
func Transform(ds *appsv1.DaemonSet) *appsv1.DaemonSet { _ = "STUB: not implemented"; return nil }

func RecordMetrics(mb *metadata.MetricsBuilder, ds *appsv1.DaemonSet, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func GetMetadata(ds *appsv1.DaemonSet) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
