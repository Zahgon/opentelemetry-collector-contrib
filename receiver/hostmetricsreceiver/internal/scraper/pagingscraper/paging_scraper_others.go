// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package pagingscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper"

import (
	"context"

	"github.com/shirou/gopsutil/v4/mem"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper/internal/metadata"
)

const (
	pagingUsageMetricsLen = 2
	pagingMetricsLen      = 2
)

// scraper for Paging Metrics
type pagingScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder

	// for mocking
	bootTime         func(context.Context) (uint64, error)
	getPageFileStats func() ([]*pageFileStats, error)
	swapMemory       func(context.Context) (*mem.SwapMemoryStat, error)
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

func (s *pagingScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *pagingScraper) scrapePagingUsageMetric() error { _ = "STUB: not implemented"; return nil }

func (s *pagingScraper) recordPagingUsageDataPoints(now pcommon.Timestamp, pageFileStats []*pageFileStats) {
	_ = "STUB: not implemented"
	return
}

func (s *pagingScraper) recordPagingUtilizationDataPoints(now pcommon.Timestamp, pageFileStats []*pageFileStats) {
	_ = "STUB: not implemented"
	return
}

func (s *pagingScraper) scrapePagingMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *pagingScraper) recordPagingOperationsDataPoints(now pcommon.Timestamp, swap *mem.SwapMemoryStat) {
	_ = "STUB: not implemented"
	return
}

func (s *pagingScraper) recordPageFaultsDataPoints(now pcommon.Timestamp, swap *mem.SwapMemoryStat) {
	_ = "STUB: not implemented"
	return
}
