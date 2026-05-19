// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package replicaset // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/replicaset"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

// Transform transforms the replica set to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new replicaset fields.
func Transform(rs *appsv1.ReplicaSet) *appsv1.ReplicaSet { _ = "STUB: not implemented"; return nil }

func RecordMetrics(mb *metadata.MetricsBuilder, rs *appsv1.ReplicaSet, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func GetMetadata(rs *appsv1.ReplicaSet) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
