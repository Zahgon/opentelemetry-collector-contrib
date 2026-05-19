// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlquery // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

func rowToMetric(row StringMap, cfg *MetricCfg, dest pmetric.Metric, startTime, ts pcommon.Timestamp, scrapeCfg scraperhelper.ControllerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func setTimestamp(cfg *MetricCfg, dp pmetric.NumberDataPoint, startTime, ts pcommon.Timestamp, scrapeCfg scraperhelper.ControllerConfig) {
	_ = "STUB: not implemented"
	return

	// Cumulative sum should have a start time set to the beginning of the data points cumulation
}

// Non-cumulative sum should have a start time set to the previous endpoint

func setMetricFields(cfg *MetricCfg, dest pmetric.Metric) pmetric.NumberDataPointSlice {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPointSlice)
}

func cfgToAggregationTemporality(agg MetricAggregation) pmetric.AggregationTemporality {
	_ = "STUB: not implemented"
	return *new(pmetric.AggregationTemporality)
}

func setDataPointValue(cfg *MetricCfg, str string, dest pmetric.NumberDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}
