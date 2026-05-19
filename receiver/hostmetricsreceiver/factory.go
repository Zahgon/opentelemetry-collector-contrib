// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hostmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/memoryscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/networkscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/nfsscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processesscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/systemscraper"
)

const (
	defaultMetadataCollectionInterval = 5 * time.Minute
)

// This file implements Factory for HostMetrics receiver.
var (
	scraperFactories = mustMakeFactories(
		cpuscraper.NewFactory(),
		diskscraper.NewFactory(),
		filesystemscraper.NewFactory(),
		loadscraper.NewFactory(),
		memoryscraper.NewFactory(),
		networkscraper.NewFactory(),
		nfsscraper.NewFactory(),
		pagingscraper.NewFactory(),
		processesscraper.NewFactory(),
		processscraper.NewFactory(),
		systemscraper.NewFactory(),
	)
)

func mustMakeFactories(factories ...scraper.Factory) map[component.Type]scraper.Factory {
	_ = "STUB: not implemented"
	return nil
}

// NewFactory creates a new factory for host metrics receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetricsReceiver creates a metrics receiver based on provided config.
func createMetricsReceiver(
	ctx context.Context,
	set receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func createLogsReceiver(
	_ context.Context, set receiver.Settings, cfg component.Config, consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func createAddScraperOptions(
	_ context.Context,
	cfg *Config,
	factories map[component.Type]scraper.Factory,
) ([]scraperhelper.ControllerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFactory(key component.Type, factories map[component.Type]scraper.Factory) (s scraper.Factory, err error) {
	_ = "STUB: not implemented"
	return *new(scraper.Factory), nil
}
