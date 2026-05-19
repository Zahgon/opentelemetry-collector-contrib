// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clusterresourcequota // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/clusterresourcequota"

import (
	quotav1 "github.com/openshift/api/quota/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

func RecordMetrics(mb *metadata.MetricsBuilder, crq *quotav1.ClusterResourceQuota, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func extractValue(k v1.ResourceName, v resource.Quantity) int64 {
	_ = "STUB: not implemented"
	return 0
}
