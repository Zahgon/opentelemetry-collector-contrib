// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func (c *prometheusConverterV2) addGaugeNumberDataPoints(dataPoints pmetric.NumberDataPointSlice,
	resource pcommon.Resource, scope pcommon.InstrumentationScope, settings Settings, name string, metadata metadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

// convert ns to ms

func (c *prometheusConverterV2) addSumNumberDataPoints(dataPoints pmetric.NumberDataPointSlice,
	resource pcommon.Resource, scope pcommon.InstrumentationScope, _ pmetric.Metric, settings Settings, name string, metadata metadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

// convert ns to ms

// getPromExemplarsV2 returns a slice of writev2.Exemplar from pdata exemplars.
func getPromExemplarsV2[T exemplarType](pt T, symbolize func(string) uint32) []writev2.Exemplar {
	_ = "STUB: not implemented"
	return nil
}

// Symbolize labels and store references
