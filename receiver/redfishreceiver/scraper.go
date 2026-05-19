// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redfishreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redfishreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/redfish"
)

// scraperClient is a struct containing the RedfishClient
// and the resources it needs to collect.
type scraperClient struct {
	*redfish.Client
	ResourceSet map[Resource]bool
}

// redfishScraper is a scraper responsible for scraping redfish metrics using multiple
// scraperClients to scrape each redfish server in the given otel config.
type redfishScraper struct {
	clients  []*scraperClient
	cfg      *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
	logger   *zap.Logger
}

func newScraper(conf *Config, settings receiver.Settings) *redfishScraper {
	_ = "STUB: not implemented"
	return nil
}

// start is a method to initialize our redfishScraper scraperClients
func (s *redfishScraper) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// create redfish clients

// create redfish client

// create resource set for each scraper client

// append new scraper client

// scrape is a method invoked periodically to scrape all server redfish apis
// and add their metrics to a metrics buffer
func (s *redfishScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// only record computer system metrics if it exists in the scraperClient's resourceSet

// Chassis metrics depend on ComputerSystem data

// only record chassis metrics if it exists in the scraperClient's resourceSet

// only scrape Fans and Temperatures if they exist in the scraperClient's resourceSet

// only record Fans metrics if it exists in the scraperClient's resourceSet

// only record Temperatures metrics if it exists in the scraperClient's resourceSet

// Always present - resource attributes
