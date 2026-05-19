// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package saphanareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver/internal/metadata"
)

// Runs intermittently, fetching info from SAP HANA, creating metrics/datapoints,
// and feeding them to a metricsConsumer.
type sapHanaScraper struct {
	settings receiver.Settings
	cfg      *Config
	mbs      map[string]*metadata.MetricsBuilder
	factory  sapHanaConnectionFactory
}

func newSapHanaScraper(settings receiver.Settings, cfg *Config, factory sapHanaConnectionFactory) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}

func (s *sapHanaScraper) getMetricsBuilder(resourceAttributes map[string]string) (*metadata.MetricsBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Scrape is called periodically, querying SAP HANA and building Metrics to send to
// the next consumer.
func (s *sapHanaScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}
