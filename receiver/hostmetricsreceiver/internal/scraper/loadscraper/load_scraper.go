// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper"

import (
	"context"
	"errors"

	"github.com/shirou/gopsutil/v4/load"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper/internal/metadata"
)

const metricsLen = 3

// errPreventScrape is used to indicate that skip scrape should be set to true
// when the sampler fails to start.
var errPreventScrape = errors.New("cannot scrape load metrics")

// scraper for Load Metrics
type loadScraper struct {
	settings   scraper.Settings
	config     *Config
	mb         *metadata.MetricsBuilder
	skipScrape bool

	// for mocking
	bootTime func(context.Context) (uint64, error)
	load     func(context.Context) (*load.AvgStat, error)
}

// newLoadScraper creates a set of Load related metrics
func newLoadScraper(_ context.Context, settings scraper.Settings, cfg *Config) *loadScraper {
	_ = "STUB: not implemented"
	return nil
}

// start
func (s *loadScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// shutdown
func (s *loadScraper) shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"

	// We skipped scraping because the sampler failed to start,
	// so it doesn't need to be shut down.
	return nil
}

// scrape
func (s *loadScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}
