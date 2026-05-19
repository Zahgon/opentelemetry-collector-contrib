// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package systemscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/systemscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/systemscraper/internal/metadata"
)

// scraper for System Metrics
type systemsScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder

	// for mocking
	bootTime func(context.Context) (uint64, error)
	uptime   func(context.Context) (uint64, error)
}

// newSystemScraper creates an Uptime related metric
func newSystemScraper(_ context.Context, settings scraper.Settings, cfg *Config) *systemsScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *systemsScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *systemsScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}
