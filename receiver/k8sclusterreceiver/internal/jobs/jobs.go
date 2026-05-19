// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jobs // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/jobs"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	batchv1 "k8s.io/api/batch/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

func RecordMetrics(mb *metadata.MetricsBuilder, j *batchv1.Job, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// Transform transforms the job to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new job fields.
func Transform(job *batchv1.Job) *batchv1.Job { _ = "STUB: not implemented"; return nil }

func GetMetadata(j *batchv1.Job) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
