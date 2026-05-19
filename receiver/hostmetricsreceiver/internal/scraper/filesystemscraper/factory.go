// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filesystemscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

// NewFactory for FileSystem scraper.
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

// Mounted by dockerd when starting a container by default
// Mounted by podman as described here: https://github.com/containers/podman/blob/ecbb52cb478309cfd59cc061f082702b69f0f4b7/docs/source/markdown/podman-run.1.md.in#L31
