// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dpfilters // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation/dpfilters"

import sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"

// FilterSet is a collection of datapoint filters, any one of which must match
// for a datapoint to be matched.
type FilterSet struct {
	excludeFilters []*dataPointFilter
	includeFilters []*dataPointFilter
}

// Matches sends a datapoint through each of the filters in the set and returns
// true if at least one of them matches the datapoint.
func (fs *FilterSet) Matches(dp *sfxpb.DataPoint) bool { _ = "STUB: not implemented"; return false }

// If we match an exclusionary filter, run through each inclusion
// filter and see if anything includes the metrics.

func NewFilterSet(excludes, includes []MetricFilter) (*FilterSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDataPointFilters(metricFilters []MetricFilter) ([]*dataPointFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
