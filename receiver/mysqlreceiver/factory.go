// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mysqlreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"

import (
	"context"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// Shared query-plan cache so that a plan fetched by scrapeTopQueries() can be
// reused by scrapeQuerySamples() without a second EXPLAIN call.
// Only create when at least one event scraper is enabled; creating unconditionally
// spawns a background goroutine even when no scraping occurs.

// we have 2 updated only attributes. so we set the cache size accordingly.
// TODO: parameterize this cache size.

// newCache creates a new cache with the given size.
// If the size is less or equal to 0, it will be set to 1.
// It will never return an error.
func newCache[v any](size int) *lru.Cache[string, v] { _ = "STUB: not implemented"; return nil }

// Ignore returned error as lru will only return an error when the size is less than 0

func newTTLCache[v any](size int, ttl time.Duration) *expirable.LRU[string, v] {
	_ = "STUB: not implemented"
	return nil
}
