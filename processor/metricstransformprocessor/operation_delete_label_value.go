// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// deleteLabelValueOp deletes a label value and all data associated with it
func deleteLabelValueOp(metric pmetric.Metric, mtpOp *internalOperation) {
	_ = "STUB: not implemented"
	return
}

//exhaustive:enforce

func hasAttr(attrs pcommon.Map, k, v string) bool { _ = "STUB: not implemented"; return false }
