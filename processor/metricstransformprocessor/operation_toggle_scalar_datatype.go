// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// toggleScalarDataTypeOp translates the numeric value type to the opposite type, int -> double and double -> int.
// Applicable to sum and gauge metrics only.
func toggleScalarDataTypeOp(metric pmetric.Metric, f internalFilter) {
	_ = "STUB: not implemented"
	return
}
