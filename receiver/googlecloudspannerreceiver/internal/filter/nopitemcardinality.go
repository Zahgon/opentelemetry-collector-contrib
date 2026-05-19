// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filter // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filter"

type nopItemCardinalityFilter struct {
	// No fields here
}

type nopItemFilterResolver struct {
	nopFilter *nopItemCardinalityFilter
}

func NewNopItemCardinalityFilter() ItemFilter { _ = "STUB: not implemented"; return *new(ItemFilter) }

func NewNopItemFilterResolver() ItemFilterResolver {
	_ = "STUB: not implemented"
	return *new(ItemFilterResolver)
}

func (*nopItemCardinalityFilter) Filter(sourceItems []*Item) []*Item {
	_ = "STUB: not implemented"
	return nil
}

func (*nopItemCardinalityFilter) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (*nopItemCardinalityFilter) TotalLimit() int { _ = "STUB: not implemented"; return 0 }

func (*nopItemCardinalityFilter) LimitByTimestamp() int { _ = "STUB: not implemented"; return 0 }

func (r *nopItemFilterResolver) Resolve(string) (ItemFilter, error) {
	_ = "STUB: not implemented"
	return *new(ItemFilter), nil
}

func (*nopItemFilterResolver) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (*nopItemCardinalityFilter) StartCache() { _ = "STUB: not implemented"; return }
