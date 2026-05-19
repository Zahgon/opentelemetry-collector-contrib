// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cache // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/cache"

import (
	lru "github.com/hashicorp/golang-lru/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// lruDecisionCache implements Cache as a simple LRU cache.
// It holds trace IDs that had sampling decisions made on them.
// It does not specify the type of sampling decision that was made, only that
// a decision was made for an ID. You need separate DecisionCaches for caching
// sampled and not sampled trace IDs.
type lruDecisionCache struct {
	cache *lru.Cache[uint64, DecisionMetadata]
}

var _ Cache = (*lruDecisionCache)(nil)

// NewLRUDecisionCache returns a new lruDecisionCache.
// The size parameter indicates the amount of keys the cache will hold before it
// starts evicting the least recently used key.
func NewLRUDecisionCache(size int) (Cache, error) {
	_ = "STUB: not implemented"
	return *new(Cache), nil
}

func (c *lruDecisionCache) Get(id pcommon.TraceID) (DecisionMetadata, bool) {
	_ = "STUB: not implemented"
	return *new(DecisionMetadata), false
}

func (c *lruDecisionCache) Put(id pcommon.TraceID, metadata DecisionMetadata) {
	_ = "STUB: not implemented"
	return
}

func rightHalfTraceID(id pcommon.TraceID) uint64 { _ = "STUB: not implemented"; return 0 }
