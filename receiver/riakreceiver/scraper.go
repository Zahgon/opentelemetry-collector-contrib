// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package riakreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver/internal/model"
)

var errClientNotInit = errors.New("client not initialized")

// riakScraper handles scraping of Riak metrics
type riakScraper struct {
	logger   *zap.Logger
	cfg      *Config
	settings component.TelemetrySettings
	client   client
	mb       *metadata.MetricsBuilder
}

// newScraper creates a new scraper
func newScraper(logger *zap.Logger, cfg *Config, settings receiver.Settings) *riakScraper {
	_ = "STUB: not implemented"
	return nil
}

// start starts the scraper by creating a new HTTP Client on the scraper
func (r *riakScraper) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// scrape collects metrics from the Riak API
func (r *riakScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	// Validate we don't attempt to scrape without initializing the client
	return *new(pmetric.Metrics), nil
}

// Get stats for processing

// collectStats collects metrics
func (r *riakScraper) collectStats(stat *model.Stats) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// scrape node.operation.count metric

// scrape node.operation.time.mean metric

// scrape node.read_repair.count metric

// scrape node.memory.limit metric

// scrape vnode.operation.count metric

// scrape vnode.index.operation.count metric
