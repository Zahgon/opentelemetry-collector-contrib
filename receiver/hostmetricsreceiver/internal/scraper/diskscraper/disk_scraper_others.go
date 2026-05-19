// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package diskscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper"

import (
	"context"

	"github.com/shirou/gopsutil/v4/disk"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper/internal/metadata"
)

const (
	standardMetricsLen = 5
	metricsLen         = standardMetricsLen + systemSpecificMetricsLen
)

// scraper for Disk Metrics
type diskScraper struct {
	settings  scraper.Settings
	config    *Config
	startTime pcommon.Timestamp
	mb        *metadata.MetricsBuilder
	includeFS filterset.FilterSet
	excludeFS filterset.FilterSet

	// for mocking
	bootTime   func(context.Context) (uint64, error)
	ioCounters func(ctx context.Context, names ...string) (map[string]disk.IOCountersStat, error)
}

// newDiskScraper creates a Disk Scraper
func newDiskScraper(_ context.Context, settings scraper.Settings, cfg *Config) (*diskScraper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *diskScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *diskScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// filter devices by name

func (s *diskScraper) recordDiskIOMetric(now pcommon.Timestamp, ioCounters map[string]disk.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *diskScraper) recordDiskOperationsMetric(now pcommon.Timestamp, ioCounters map[string]disk.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *diskScraper) recordDiskIOTimeMetric(now pcommon.Timestamp, ioCounters map[string]disk.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *diskScraper) recordDiskOperationTimeMetric(now pcommon.Timestamp, ioCounters map[string]disk.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *diskScraper) recordDiskPendingOperationsMetric(now pcommon.Timestamp, ioCounters map[string]disk.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *diskScraper) filterByDevice(ioCounters map[string]disk.IOCountersStat) map[string]disk.IOCountersStat {
	_ = "STUB: not implemented"
	return nil
}

func (s *diskScraper) includeDevice(deviceName string) bool {
	_ = "STUB: not implemented"
	return false
}
