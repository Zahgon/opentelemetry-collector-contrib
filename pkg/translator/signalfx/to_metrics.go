// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfx // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx"

import (
	"github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const numMetricTypes = 4

// ToTranslator converts from SignalFx proto data model to pdata.
type ToTranslator struct {
	// prevent unkeyed literal initialization
	_ struct{}
}

// ToMetrics converts SignalFx proto data points to pmetric.Metrics.
func (*ToTranslator) ToMetrics(sfxDataPoints []*model.DataPoint) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// This is a map from [metric_name, metric_type] -> index + 1 in the Metrics slice. Used to combine datapoints together.

// TODO: Log or metric for this odd ball?

func setDataTypeAndPoints(sfxDataPoint *model.DataPoint, ms pmetric.MetricSlice, datapointToMetric map[string][4]int) error {
	_ = "STUB: not implemented"
	return nil
}

// Only emit gauge and sum.

// Numerical: Periodic, instantaneous measurement of some state.

func fillNumberDataPoint(sfxDataPoint *model.DataPoint, dps pmetric.NumberDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func fillInAttributes(dimensions []*model.Dimension, attributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// TODO: Log or metric for this odd ball?
