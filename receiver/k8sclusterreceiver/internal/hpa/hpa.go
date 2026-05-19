// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hpa // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/hpa"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	autoscalingv2 "k8s.io/api/autoscaling/v2"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

func RecordMetrics(mb *metadata.MetricsBuilder, hpa *autoscalingv2.HorizontalPodAutoscaler, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func GetMetadata(hpa *autoscalingv2.HorizontalPodAutoscaler) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
