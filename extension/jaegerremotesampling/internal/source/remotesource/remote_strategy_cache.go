// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotesource // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source/remotesource"

import (
	"context"
	"sync"
	"time"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
)

type serviceStrategyCache interface {
	get(ctx context.Context, serviceName string) (*api_v2.SamplingStrategyResponse, bool)
	put(ctx context.Context, serviceName string, response *api_v2.SamplingStrategyResponse)
	Close() error
}

// serviceStrategyCacheEntry is a timestamped sampling strategy response
type serviceStrategyCacheEntry struct {
	retrievedAt      time.Time
	strategyResponse *api_v2.SamplingStrategyResponse
}

// serviceStrategyTTLCache is a naive in-memory TTL serviceStrategyTTLCache of service-specific sampling strategies
// returned from the remote source. Each cached item has its own TTL used to determine whether it is valid for read
// usage (based on the time of write).
type serviceStrategyTTLCache struct {
	itemTTL time.Duration

	stopCh chan struct{}
	rw     sync.RWMutex
	items  map[string]serviceStrategyCacheEntry
}

// Initial size of cache's underlying map
const initialRemoteResponseCacheSize = 32

func newServiceStrategyCache(itemTTL time.Duration) serviceStrategyCache {
	_ = "STUB: not implemented"
	return *new(serviceStrategyCache)
}

// Launches a "cleaner" goroutine that naively blows away stale items with a frequency equal to the item TTL.
// Note that this is for memory usage and not for correctness (the get() function checks item validity).

// get returns a cached sampling strategy if one is present and is no older than the serviceStrategyTTLCache's per-item TTL.
func (c *serviceStrategyTTLCache) get(
	ctx context.Context,
	serviceName string,
) (*api_v2.SamplingStrategyResponse, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// put unconditionally overwrites the given service's serviceStrategyTTLCache item entry and resets its timestamp used for TTL checks.
func (c *serviceStrategyTTLCache) put(
	ctx context.Context,
	serviceName string,
	response *api_v2.SamplingStrategyResponse,
) {
	_ = "STUB: not implemented"
	return
}

// periodicallyClearCache periodically clears expired items from the cache and replaces the backing map with only
// valid (fresh) items. Note that this is not necessary for correctness, just preferred for memory usage hygiene.
// Client request activity drives the replacement of stale items with fresh items upon cache misses for any service.
func (c *serviceStrategyTTLCache) periodicallyClearCache(
	ctx context.Context,
	schedulingPeriod time.Duration,
) {
	_ = "STUB: not implemented"
	return
}

// Notice that we swap the map rather than using map's delete (which doesn't reduce its allocated size).

func (c *serviceStrategyTTLCache) Close() error { _ = "STUB: not implemented"; return nil }

func (c *serviceStrategyTTLCache) staleItem(ctx context.Context, item serviceStrategyCacheEntry) bool {
	_ = "STUB: not implemented"
	return false
}

type noopStrategyCache struct{}

func (*noopStrategyCache) get(context.Context, string) (*api_v2.SamplingStrategyResponse, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (*noopStrategyCache) put(context.Context, string, *api_v2.SamplingStrategyResponse) {
	_ = "STUB: not implemented"
	return
}

func (*noopStrategyCache) Close() error { _ = "STUB: not implemented"; return nil }

func newNoopStrategyCache() serviceStrategyCache {
	_ = "STUB: not implemented"
	return *new(serviceStrategyCache)
}
