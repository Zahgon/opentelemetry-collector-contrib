// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func (c *prometheusConverter) addGaugeNumberDataPoints(dataPoints pmetric.NumberDataPointSlice,
	resource pcommon.Resource, scope pcommon.InstrumentationScope, settings Settings, name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// convert ns to ms

func (c *prometheusConverter) addSumNumberDataPoints(dataPoints pmetric.NumberDataPointSlice,
	resource pcommon.Resource, scope pcommon.InstrumentationScope, _ pmetric.Metric, settings Settings, name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// convert ns to ms
