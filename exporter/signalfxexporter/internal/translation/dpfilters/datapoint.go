// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dpfilters // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation/dpfilters"

import (
	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
)

type dataPointFilter struct {
	metricFilter     *StringFilter
	dimensionsFilter *dimensionsFilter
}

// newDataPointFilter returns a new dataPointFilter filter with the given configuration.
func newDataPointFilter(metricNames []string, dimSet map[string][]string) (*dataPointFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Matches tests a datapoint to see whether it is excluded by this
func (f *dataPointFilter) Matches(dp *sfxpb.DataPoint) bool {
	_ = "STUB: not implemented"
	return false
}
