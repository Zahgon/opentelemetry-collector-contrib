// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dpfilters // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation/dpfilters"

import (
	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
)

type dimensionsFilter struct {
	filterMap map[string]*StringFilter
}

// newDimensionsFilter returns a filter that matches against a
// sfxpb.Dimension slice. The filter will return false if there's
// at least one dimension in the slice that fails to match. In case`
// there are no filters for any of the dimension keys in the slice,
// the filter will return false.
func newDimensionsFilter(m map[string][]string) (*dimensionsFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *dimensionsFilter) Matches(dimensions []*sfxpb.Dimension) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip if there are no filters associated with current dimension key.
