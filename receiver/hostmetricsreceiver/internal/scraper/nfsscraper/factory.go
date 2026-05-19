// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nfsscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/nfsscraper"

import (
	"context"
	"errors"
	"runtime"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

var (
	supportedOS      = runtime.GOOS == "linux"
	errUnsupportedOS = errors.New("the nfs scraper is only available on Linux")
)

// NewFactory for NFS scraper.
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
	return *new(scraper.Metrics), nil
}
