// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package resourcequota // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/resourcequota"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

func RecordMetrics(mb *metadata.MetricsBuilder, rq *corev1.ResourceQuota, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}
