// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// updateLabelOp updates labels and label values in metric based on given operation
func updateLabelOp(metric pmetric.Metric, mtpOp *internalOperation, f internalFilter) {
	_ = "STUB: not implemented"
	return
}
