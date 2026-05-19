// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package snowflakereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snowflakereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snowflakereceiver/internal/metadata"
)

type snowflakeMetricsScraper struct {
	client   *snowflakeClient
	settings component.TelemetrySettings
	conf     *Config
	mb       *metadata.MetricsBuilder
}

func newSnowflakeMetricsScraper(settings receiver.Settings, conf *Config) *snowflakeMetricsScraper {
	_ = "STUB: not implemented"
	return nil
}

// for use with receiver.scraperhelper
func (s *snowflakeMetricsScraper) start(_ context.Context, _ component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *snowflakeMetricsScraper) shutdown(_ context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func errorListener(eQueue <-chan error, eOut chan<- *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// wrapper for all of the sub-scraping tasks, implements the scraper interface for
// snowflakeMetricsScraper
func (s *snowflakeMetricsScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *snowflakeMetricsScraper) scrapeBillingMetrics(ctx context.Context, t pcommon.Timestamp, errs chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (s *snowflakeMetricsScraper) scrapeWarehouseBillingMetrics(ctx context.Context, t pcommon.Timestamp, errs chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (s *snowflakeMetricsScraper) scrapeLoginMetrics(ctx context.Context, t pcommon.Timestamp, errs chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (s *snowflakeMetricsScraper) scrapeHighLevelQueryMetrics(ctx context.Context, t pcommon.Timestamp, errs chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (s *snowflakeMetricsScraper) scrapeDBMetrics(ctx context.Context, t pcommon.Timestamp, errs chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (s *snowflakeMetricsScraper) scrapeSessionMetrics(ctx context.Context, t pcommon.Timestamp, errs chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (s *snowflakeMetricsScraper) scrapeSnowpipeMetrics(ctx context.Context, t pcommon.Timestamp, errs chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (s *snowflakeMetricsScraper) scrapeStorageMetrics(ctx context.Context, t pcommon.Timestamp, errs chan<- error) {
	_ = "STUB: not implemented"
	return
}
