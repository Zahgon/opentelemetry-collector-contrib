// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal/metadata"

import "go.opentelemetry.io/collector/pdata/pcommon"

func (mb *MetricsBuilder) RecordVcenterResourcePoolMemoryUsageDataPointWithoutTypeAttribute(ts pcommon.Timestamp, val int64) {
	_ = "STUB: not implemented"
	return
}

func (m *metricVcenterResourcePoolMemoryUsage) recordDataPointWithoutType(ts pcommon.Timestamp, val int64) {
	_ = "STUB: not implemented"
	return
}
