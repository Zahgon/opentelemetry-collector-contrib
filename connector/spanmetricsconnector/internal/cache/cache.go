// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cache // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector/internal/cache"

import (
	"github.com/hashicorp/golang-lru/v2/simplelru"
)

// Cache consists of an LRU cache and the evicted items from the LRU cache.
// This data structure makes sure all the cached items can be retrieved either from the LRU cache or the evictedItems
// map. In spanmetricsconnector's use case, we need to hold all the items during the current processing step for
// building the metrics. The evicted items can/should be safely removed once the metrics are built from the current
// batch of spans.
//
// Important: This implementation is non-thread safe.
type Cache[K comparable, V any] struct {
	lru          *simplelru.LRU[K, V]
	evictedItems map[K]V
}

// NewCache creates a Cache.
func NewCache[K comparable, V any](size int) (*Cache[K, V], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveEvictedItems cleans all the evicted items.
func (c *Cache[K, V]) RemoveEvictedItems() {
	_ = "STUB: not implemented"
	// we need to keep the original pointer to evictedItems map as it is used in the closure of lru.NewWithEvict
	return
}

// Add a value to the cache, returns true if an eviction occurred and updates the "recently used"-ness of the key.
func (c *Cache[K, V]) Add(key K, value V) bool { _ = "STUB: not implemented"; return false }

// Get an item from the LRU cache or evicted items.
func (c *Cache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// Revive from evicted items back into the main cache if a fetch was attempted.

// Remove removes a key from the cache if it exists.
func (c *Cache[K, V]) Remove(key K) bool { _ = "STUB: not implemented"; return false }

// Len returns the number of items in the cache.
func (c *Cache[K, V]) Len() int {
	_ = "STUB: not implemented"

	// Purge removes all the items from the LRU cache and evicted items.
	return 0
}

func (c *Cache[K, V]) Purge() { _ = "STUB: not implemented"; return }

// ForEach iterates over all the items within the cache, as well as the evicted items (if any).
func (c *Cache[K, V]) ForEach(fn func(k K, v V)) { _ = "STUB: not implemented"; return }
