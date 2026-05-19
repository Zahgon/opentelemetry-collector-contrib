// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// scaleValueOp scales the numeric metric value of sum and gauge metrics.
// For histograms it scales the value of the sum and the explicit bounds.
func scaleValueOp(metric pmetric.Metric, op *internalOperation, f internalFilter) {
	_ = "STUB: not implemented"
	return
}

func scaleHistogramOp(metric pmetric.Metric, op *internalOperation, f internalFilter) {
	_ = "STUB: not implemented"
	return
}

func scaleExpHistogramOp(metric pmetric.Metric, op *internalOperation, f internalFilter) {
	_ = "STUB: not implemented"
	return
}

// For the buckets, we only need to change the offset.
// The bucket counts and the scale remain the same.

func updateOffset(scale, offset int32, op *internalOperation) int32 {
	_ = "STUB: not implemented"
	// Take the middle of the first bucket.
	return 0
}

// Scale it according to the config.

// Find the new offset by mapping the scaled value.

// mapToIndex returns the index that corresponds to the given value on the scale.
// See https://opentelemetry.io/docs/specs/otel/metrics/data-model/#all-scales-use-the-logarithm-function.
func mapToIndex(value float64, scale int) int32 { _ = "STUB: not implemented"; return 0 }

func scaleExemplars(exemplars pmetric.ExemplarSlice, op *internalOperation) {
	_ = "STUB: not implemented"
	return
}
