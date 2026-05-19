// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterfactory // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filterfactory"

import (
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

type filterBuilder struct {
	logger *zap.Logger
	config *ItemFilterFactoryConfig
}

func (b filterBuilder) buildFilterByMetricZeroTotalLimit() map[string]filter.ItemFilter {
	_ = "STUB: not implemented"
	return nil
}

func (b filterBuilder) buildFilterByMetricPositiveTotalLimit() (map[string]filter.ItemFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle metric groups with low cardinality

// Handle metric groups with high cardinality

func (b filterBuilder) handleLowCardinalityGroups(groups []*metadata.MetricsMetadata, remainingTotalLimit int,
	filterByMetric map[string]filter.ItemFilter,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// For low cardinality metrics total limit is equal to limit by timestamp

func (b filterBuilder) handleHighCardinalityGroups(groups []*metadata.MetricsMetadata, remainingTotalLimit int,
	filterByMetric map[string]filter.ItemFilter,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b filterBuilder) constructFiltersForGroups(totalLimitPerMetric, limitPerMetricByTimestamp int,
	groups []*metadata.MetricsMetadata, remainingTotalLimit int, filterByMetric map[string]filter.ItemFilter,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func countMetricsInGroups(metadataItems []*metadata.MetricsMetadata) (amount int) {
	_ = "STUB: not implemented"
	return 0
}

func groupByCardinality(metadataItems []*metadata.MetricsMetadata) map[bool][]*metadata.MetricsMetadata {
	_ = "STUB: not implemented"
	return nil
}
