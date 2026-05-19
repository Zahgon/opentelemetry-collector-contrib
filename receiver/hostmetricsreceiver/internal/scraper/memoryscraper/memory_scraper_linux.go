// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package memoryscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/memoryscraper"

import (
	"github.com/shirou/gopsutil/v4/mem"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func (s *memoryScraper) recordMemoryUsageMetric(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	// TODO: rely on memInfo.Used value once https://github.com/shirou/gopsutil/pull/1882 is released
	// gopsutil formula: https://github.com/shirou/gopsutil/pull/1882/files#diff-5af8322731595fb792b48f3c38f31ddb24f596cf11a74a9c37b19734597baef6R321
	return
}

// gopsutil legacy "Used" memory formula = Total - Free - Buffers - Cache

func (s *memoryScraper) recordMemoryLinuxSharedMetric(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	return
}

func (s *memoryScraper) recordMemoryUtilizationMetric(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	// TODO: rely on memInfo.Used value once https://github.com/shirou/gopsutil/pull/1882 is released
	// gopsutil formula: https://github.com/shirou/gopsutil/pull/1882/files#diff-5af8322731595fb792b48f3c38f31ddb24f596cf11a74a9c37b19734597baef6R321
	return
}

// gopsutil legacy "Used" memory formula = Total - Free - Buffers - Cache

func (s *memoryScraper) recordLinuxMemoryAvailableMetric(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	return
}

func (s *memoryScraper) recordLinuxMemoryDirtyMetric(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	// This value is collected from /proc/meminfo and converted from kB to bytes in gopsutil:
	// https://github.com/shirou/gopsutil/blob/d8750909ba41f2de9750c90a6d2074c68dfc677e/mem/mem_linux.go#L148
	return
}

func (s *memoryScraper) recordLinuxHugePagesMetrics(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	// Record page size in bytes (converted from kB to bytes in gopsutil):
	// https://github.com/shirou/gopsutil/blob/v4.25.12/mem/mem_linux.go#L303
	return
}

// Record limit (total hugepages available) in number of pages

// Calculate used pages

// Record usage with state attributes in number of pages

// Record reserved as a separate metric (not a state, since reserved pages are included in free_huge_pages
// but cannot be used for non-reserved allocations)

// Record surplus as a separate metric (not a state, since surplus pages can also be in used/free states)

// Record utilization with state attributes as percentage

func (s *memoryScraper) recordSystemSpecificMetrics(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	return
}
