// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processesscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processesscraper"

import (
	"context"

	"github.com/shirou/gopsutil/v4/load"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processesscraper/internal/metadata"
)

var metricsLength = func() int {
	n := 0
	if enableProcessesCount {
		n++
	}
	if enableProcessesCreated {
		n++
	}
	return n
}()

// scraper for Processes Metrics
type processesScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder

	// for mocking gopsutil
	getMiscStats func(context.Context) (*load.MiscStat, error)
	getProcesses func(context.Context) ([]proc, error)
	bootTime     func(context.Context) (uint64, error)
}

// for mocking out gopsutil process.Process
type proc interface {
	Status() ([]string, error)
}

type processesMetadata struct {
	countByStatus    map[metadata.AttributeStatus]int64 // ignored if enableProcessesCount is false
	processesCreated *int64                             // ignored if enableProcessesCreated is false
}

// newProcessesScraper creates a set of Processes related metrics
func newProcessesScraper(_ context.Context, settings scraper.Settings, cfg *Config) *processesScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *processesScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processesScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}
