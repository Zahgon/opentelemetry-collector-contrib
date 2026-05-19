// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cronjob // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/cronjob"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	batchv1 "k8s.io/api/batch/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

const (
	// Keys for cronjob metadata.
	cronJobKeySchedule          = "schedule"
	cronJobKeyConcurrencyPolicy = "concurrency_policy"
)

func RecordMetrics(mb *metadata.MetricsBuilder, cj *batchv1.CronJob, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func GetMetadata(cj *batchv1.CronJob) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
