// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package diskscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper/internal/metadata"
)

const (
	metricsLen = 5

	logicalDisk = "LogicalDisk"

	readsPerSec  = "Disk Reads/sec"
	writesPerSec = "Disk Writes/sec"

	readBytesPerSec  = "Disk Read Bytes/sec"
	writeBytesPerSec = "Disk Write Bytes/sec"

	idleTime = "% Idle Time"

	avgDiskSecsPerRead  = "Avg. Disk sec/Read"
	avgDiskSecsPerWrite = "Avg. Disk sec/Write"

	queueLength = "Current Disk Queue Length"
)

var counterNames = []string{
	readBytesPerSec,
	writeBytesPerSec,
	readsPerSec,
	writesPerSec,
	idleTime,
	avgDiskSecsPerRead,
	avgDiskSecsPerWrite,
	queueLength,
}

// diskScraper for Disk Metrics
type diskScraper struct {
	settings  scraper.Settings
	config    *Config
	startTime pcommon.Timestamp
	mb        *metadata.MetricsBuilder
	includeFS filterset.FilterSet
	excludeFS filterset.FilterSet

	perfCounters []winperfcounters.PerfCounterWatcher
	skipScrape   bool

	// for mocking
	bootTime           func(ctx context.Context) (uint64, error)
	perfCounterFactory func(string, string, string) (winperfcounters.PerfCounterWatcher, error)
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

// Initialize the performance counter watchers

func (s *diskScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// For each counter and respective set of values record the metrics

// NOTE: int64-to-uint64 cast is safe because perf counter values are non-negative.

func includeDevice(deviceName string, includeFS, excludeFS filterset.FilterSet) bool {
	_ = "STUB: not implemented"
	return false
}
