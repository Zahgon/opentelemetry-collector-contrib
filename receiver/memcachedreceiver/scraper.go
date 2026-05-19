// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package memcachedreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver/internal/metadata"
)

type memcachedScraper struct {
	logger    *zap.Logger
	config    *Config
	mb        *metadata.MetricsBuilder
	newClient newMemcachedClientFunc
}

func newMemcachedScraper(
	settings receiver.Settings,
	config *Config,
) memcachedScraper {
	_ = "STUB: not implemented"
	return *new(memcachedScraper)
}

func (r *memcachedScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	// Init client in scrape method in case there are transient errors in the
	// constructor.
	return *new(pmetric.Metrics), nil
}

// Calculated Metrics

func calculateHitRatio(misses, hits int64) float64 { _ = "STUB: not implemented"; return 0 }

// parseInt converts string to int64.
func (r *memcachedScraper) parseInt(key, value string) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// parseFloat converts string to float64.
func (r *memcachedScraper) parseFloat(key, value string) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (r *memcachedScraper) logInvalid(expectedType, key, value string) {
	_ = "STUB: not implemented"
	return
}
