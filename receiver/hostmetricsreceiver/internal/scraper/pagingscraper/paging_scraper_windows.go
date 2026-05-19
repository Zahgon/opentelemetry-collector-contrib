// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package pagingscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper/internal/metadata"
)

const (
	pagingUsageMetricsLen = 1
	pagingMetricsLen      = 1

	memory = "Memory"

	// The counters below are per second rates, but, instead of reading the calculated rates,
	// we read the raw values and post them as cumulative metrics.
	pageReadsPerSec  = "Page Reads/sec"
	pageWritesPerSec = "Page Writes/sec"
	pageFaultsPerSec = "Page Faults/sec" // All page faults, including minor and major, aka soft and hard faults
	pageMajPerSec    = "Pages/sec"       // Only major, aka hard, page faults.
)

// scraper for Paging Metrics
type pagingScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder

	pageReadsPerfCounter     winperfcounters.PerfCounterWatcher
	pageWritesPerfCounter    winperfcounters.PerfCounterWatcher
	pageFaultsPerfCounter    winperfcounters.PerfCounterWatcher
	pageMajFaultsPerfCounter winperfcounters.PerfCounterWatcher
	skipScrape               bool

	// for mocking
	bootTime           func(context.Context) (uint64, error)
	pageFileStats      func() ([]*pageFileStats, error)
	perfCounterFactory func(string, string, string) (winperfcounters.PerfCounterWatcher, error)
}

// newPagingScraper creates a Paging Scraper
func newPagingScraper(_ context.Context, settings scraper.Settings, cfg *Config) *pagingScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *pagingScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *pagingScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *pagingScraper) scrapePagingUsageMetric() error { _ = "STUB: not implemented"; return nil }

func (s *pagingScraper) recordPagingUsageDataPoints(now pcommon.Timestamp, pageFiles []*pageFileStats) {
	_ = "STUB: not implemented"
	return
}

func (s *pagingScraper) recordPagingUtilizationDataPoints(now pcommon.Timestamp, pageFiles []*pageFileStats) {
	_ = "STUB: not implemented"
	return
}

func (s *pagingScraper) scrapePagingOperationsMetric() error { _ = "STUB: not implemented"; return nil }

func (s *pagingScraper) scrapePagingFaultsMetric(errors *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Count is 2 since without major page faults none of the paging metrics will be recorded
