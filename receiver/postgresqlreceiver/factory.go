// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package postgresqlreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"

import (
	"context"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// newCache creates a new cache with the given size.
// If the size is less or equal to 0, it will be set to 1.
// It will never return an error.
func newCache(size int) *lru.Cache[string, float64] { _ = "STUB: not implemented"; return nil }

// lru will only return error when the size is less than 0

func newTTLCache[v any](size int, ttl time.Duration) *expirable.LRU[string, v] {
	_ = "STUB: not implemented"
	return nil
}

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

// createLogsReceiver create a logs receiver based on provided config.
func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	receiverCfg component.Config,
	logsConsumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// query sample collection does not need cache, but we do not want to make it
// nil, so create one size 1 cache as a placeholder.

// we have 10 updated only attributes. so we set the cache size accordingly.
