// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package diskscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper"

import (
	"github.com/shirou/gopsutil/v4/disk"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const systemSpecificMetricsLen = 2

func (s *diskScraper) recordSystemSpecificDataPoints(now pcommon.Timestamp, ioCounters map[string]disk.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *diskScraper) recordDiskWeightedIOTimeMetric(now pcommon.Timestamp, ioCounters map[string]disk.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *diskScraper) recordDiskMergedMetric(now pcommon.Timestamp, ioCounters map[string]disk.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}
