// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// aggregateLabelValuesOp aggregates points that have the label values specified in aggregated_values
func aggregateLabelValuesOp(metric pmetric.Metric, mtpOp *internalOperation) {
	_ = "STUB: not implemented"
	return
}
