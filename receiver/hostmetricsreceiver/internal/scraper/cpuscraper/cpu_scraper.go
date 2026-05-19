// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cpuscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper"

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper/ucal"
)

const (
	metricsLen = 2
	hzInAMHz   = 1_000_000
)

// cpuScraper for CPU Metrics
type cpuScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder
	ucal     *ucal.CPUUtilizationCalculator

	// for mocking
	bootTime func(context.Context) (uint64, error)
	times    func(context.Context, bool) ([]cpu.TimesStat, error)
	now      func() time.Time
}

type cpuInfo struct {
	frequency float64
	processor uint
}

// newCPUScraper creates a set of CPU related metrics
func newCPUScraper(_ context.Context, settings scraper.Settings, cfg *Config) *cpuScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *cpuScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *cpuScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

/*percpu=*/
