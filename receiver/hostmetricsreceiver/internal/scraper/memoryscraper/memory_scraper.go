// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package memoryscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/memoryscraper"

import (
	"context"
	"errors"

	"github.com/shirou/gopsutil/v4/mem"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/memoryscraper/internal/metadata"
)

const metricsLen = 2

var ErrInvalidTotalMem = errors.New("invalid total memory")

// scraper for Memory Metrics
type memoryScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder

	// for mocking gopsutil mem.VirtualMemory
	bootTime      func(context.Context) (uint64, error)
	virtualMemory func(context.Context) (*mem.VirtualMemoryStat, error)

	pageSize int64
}

// newMemoryScraper creates a Memory Scraper
func newMemoryScraper(_ context.Context, settings scraper.Settings, cfg *Config) *memoryScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *memoryScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *memoryScraper) recordMemoryLimitMetric(now pcommon.Timestamp, memInfo *mem.VirtualMemoryStat) {
	_ = "STUB: not implemented"
	return
}

func (s *memoryScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *memoryScraper) recordMemoryPageSizeMetric(now pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}
