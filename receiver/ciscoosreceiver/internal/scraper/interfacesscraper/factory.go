// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package interfacesscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/interfacesscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
	"go.uber.org/zap"
)

// NewFactory creates a factory for interfaces scraper.
func NewFactory() scraper.Factory { _ = "STUB: not implemented"; return *new(scraper.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsScraper(_ context.Context, settings scraper.Settings, cfg component.Config) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}

func newInterfacesScraper(logger *zap.Logger, config *Config) *interfacesScraper {
	_ = "STUB: not implemented"
	return nil
}
