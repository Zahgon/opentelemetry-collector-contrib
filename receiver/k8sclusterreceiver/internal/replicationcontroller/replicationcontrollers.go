// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package replicationcontroller // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/replicationcontroller"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

func RecordMetrics(mb *metadata.MetricsBuilder, rc *corev1.ReplicationController, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func GetMetadata(rc *corev1.ReplicationController) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
