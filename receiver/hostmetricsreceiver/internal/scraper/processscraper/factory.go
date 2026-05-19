// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

// NewFactory for Process scraper.
func NewFactory() scraper.Factory { _ = "STUB: not implemented"; return *new(scraper.Factory) }

// createDefaultConfig creates the default configuration for the Scraper.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetricsScraper creates a resource scraper based on provided config.
func createMetricsScraper(
	_ context.Context,
	settings scraper.Settings,
	cfg component.Config,
) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	// Hardcoded supported-OS list pending mdatagen `supported_os` annotation support.
	// See https://github.com/open-telemetry/opentelemetry-collector/issues/15020;
	// migrate to a metadata.yaml declaration and drop the runtime.GOOS check once it lands.
	return *new(scraper.Metrics), nil
}
