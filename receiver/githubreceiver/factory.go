// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver"

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver/internal/scraper/githubscraper"
)

// This file implements a factory for the github receiver

const (
	defaultReadTimeout       = 500 * time.Millisecond
	defaultWriteTimeout      = 500 * time.Millisecond
	defaultPath              = "/events"
	defaultHealthPath        = "/health"
	defaultEndpoint          = "localhost:8080"
	defaultIncludeSpanEvents = false
)

var (
	scraperFactories = map[string]internal.ScraperFactory{
		githubscraper.TypeStr: &githubscraper.Factory{},
	}

	errConfigNotValid = errors.New("configuration is not valid for the github receiver")
)

// NewFactory creates a factory for the github receiver
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// Gets a factory for defined scraper.
func getScraperFactory(key string) (internal.ScraperFactory, bool) {
	_ = "STUB: not implemented"
	return *new(internal.ScraperFactory), false
}

// Create the default config based on the const(s) defined above.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Create the metrics receiver according to the OTEL conventions taking in the
// context, receiver params, configuration from the component, and consumer (process or exporter)
func createMetricsReceiver(
	ctx context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	// check that the configuration is valid
	return *new(receiver.Metrics), nil
}

func createTracesReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	// check that the configuration is valid
	return *new(receiver.Traces), nil
}

func createAddScraperOpts(
	ctx context.Context,
	params receiver.Settings,
	cfg *Config,
	factories map[string]internal.ScraperFactory,
) ([]scraperhelper.ControllerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createGitHubScraper(
	ctx context.Context,
	params receiver.Settings,
	key string,
	cfg internal.Config,
	factories map[string]internal.ScraperFactory,
) (s scraper.Metrics, err error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}
