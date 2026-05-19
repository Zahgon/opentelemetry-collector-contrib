// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processesscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processesscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

// NewFactory for Processes scraper.
func NewFactory() scraper.Factory { _ = "STUB: not implemented"; return *new(scraper.Factory) }

// createDefaultConfig creates the default configuration for the Scraper.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetricsScraper creates a scraper based on provided config.
func createMetricsScraper(
	ctx context.Context,
	settings scraper.Settings,
	config component.Config,
) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}
