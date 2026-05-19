// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterfactory // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filterfactory"

import (
	"time"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

const (
	defaultMetricDataPointsAmountInPeriod = 24 * 60
	defaultItemActivityPeriod             = 24 * time.Hour
)

type itemFilterFactory struct {
	filterByMetric map[string]filter.ItemFilter
}

type ItemFilterFactoryConfig struct {
	MetadataItems  []*metadata.MetricsMetadata
	TotalLimit     int
	ProjectAmount  int
	InstanceAmount int
	DatabaseAmount int
}

func NewItemFilterResolver(logger *zap.Logger, config *ItemFilterFactoryConfig) (filter.ItemFilterResolver, error) {
	_ = "STUB: not implemented"
	return *new(filter.ItemFilterResolver), nil
}

func (config *ItemFilterFactoryConfig) validate() error { _ = "STUB: not implemented"; return nil }

func (f *itemFilterFactory) Resolve(metricFullName string) (filter.ItemFilter, error) {
	_ = "STUB: not implemented"
	return *new(filter.ItemFilter), nil
}

func (f *itemFilterFactory) Shutdown() error { _ = "STUB: not implemented"; return nil }

// StartAllCaches starts the TTL caches for all filters inside the resolver.
// It is a no-op if the resolver is not the concrete factory from this package.
func StartAllCaches(resolver filter.ItemFilterResolver) { _ = "STUB: not implemented"; return }
