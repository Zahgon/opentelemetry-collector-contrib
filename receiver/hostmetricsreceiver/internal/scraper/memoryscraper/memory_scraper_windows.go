// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package memoryscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/memoryscraper"

import (
	"github.com/shirou/gopsutil/v4/mem"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func (s *memoryScraper) recordMemoryUsageMetric(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	return
}

func (s *memoryScraper) recordMemoryUtilizationMetric(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	return
}

func (*memoryScraper) recordSystemSpecificMetrics(pcommon.Timestamp, *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	return
}
