// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zookeeperscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/scraper/zookeeperscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

const (
	defaultEndpoint = "localhost:2181"
)

func NewFactory() scraper.Factory { _ = "STUB: not implemented"; return *new(scraper.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetrics creates zookeeper (metrics) scraper.
func createMetrics(
	_ context.Context,
	params scraper.Settings,
	config component.Config,
) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}
