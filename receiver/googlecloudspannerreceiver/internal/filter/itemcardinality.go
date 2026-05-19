// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filter // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filter"

import (
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"go.uber.org/zap"
)

type Item struct {
	SeriesKey string
	Timestamp time.Time
}

type ItemFilter interface {
	Filter(source []*Item) []*Item
	Shutdown() error
	TotalLimit() int
	LimitByTimestamp() int
	StartCache()
}

type ItemFilterResolver interface {
	Resolve(metricFullName string) (ItemFilter, error)
	Shutdown() error
}

type itemCardinalityFilter struct {
	metricName         string
	totalLimit         int
	limitByTimestamp   int
	itemActivityPeriod time.Duration
	logger             *zap.Logger
	cache              *ttlcache.Cache[string, struct{}]
	stopOnce           sync.Once
	startOnce          sync.Once
}

type currentLimitByTimestamp struct {
	limitByTimestamp int
}

func (f *currentLimitByTimestamp) dec() { _ = "STUB: not implemented"; return }

func (f *currentLimitByTimestamp) get() int { _ = "STUB: not implemented"; return 0 }

func NewItemCardinalityFilter(metricName string, totalLimit, limitByTimestamp int,
	itemActivityPeriod time.Duration, logger *zap.Logger,
) (ItemFilter, error) {
	_ = "STUB: not implemented"
	return *new(ItemFilter), nil
}

func (f *itemCardinalityFilter) TotalLimit() int { _ = "STUB: not implemented"; return 0 }

func (f *itemCardinalityFilter) LimitByTimestamp() int { _ = "STUB: not implemented"; return 0 }

func (f *itemCardinalityFilter) Filter(sourceItems []*Item) []*Item {
	_ = "STUB: not implemented"
	return nil
}

// StartCache explicitly starts the TTL cache cleanup loop.
// Idempotent: safe to call multiple times.
func (f *itemCardinalityFilter) StartCache() { _ = "STUB: not implemented"; return }

func (f *itemCardinalityFilter) filterItems(items []*Item) []*Item {
	_ = "STUB: not implemented"
	return nil
}

func (f *itemCardinalityFilter) includeItem(item *Item, limit *currentLimitByTimestamp) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *itemCardinalityFilter) canIncludeNewItem(currentLimitByTimestamp int) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *itemCardinalityFilter) Shutdown() error { _ = "STUB: not implemented"; return nil }

func groupByTimestamp(items []*Item) map[time.Time][]*Item { _ = "STUB: not implemented"; return nil }

func sortedKeys(groupedItems map[time.Time][]*Item) []time.Time {
	_ = "STUB: not implemented"
	return nil
}
