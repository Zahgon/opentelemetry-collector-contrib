// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetrictest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// ValidateMetrics reports semantic errors in md (Metrics Data).
// Currently it checks that no two datapoints within the same metric share
// identical attribute sets (duplicate datapoint identity).
// It returns nil if no violations are found.
func ValidateMetrics(md pmetric.Metrics) error { _ = "STUB: not implemented"; return nil }

// validateDatapointUniqueness checks that no two datapoints within the given
// metric have identical attribute sets.
func validateDatapointUniqueness(m pmetric.Metric) error {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return nil
}

// extractAttributeMaps builds a slice of attribute maps from any datapoint slice
// using a callback, since the various datapoint slice types share no common interface.
func extractAttributeMaps(count int, getAttrs func(int) pcommon.Map) []pcommon.Map {
	_ = "STUB: not implemented"
	return nil
}

// checkDuplicateDatapointAttrs reports an error for every pair of datapoints
// whose attribute maps hash to the same value.
func checkDuplicateDatapointAttrs(attrs []pcommon.Map) error { _ = "STUB: not implemented"; return nil }

// hash to first index
