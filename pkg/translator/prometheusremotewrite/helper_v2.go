// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// addResourceTargetInfoV2 converts the resource to the target info metric.
func (c *prometheusConverterV2) addResourceTargetInfoV2(resource pcommon.Resource, settings Settings, timestamp pcommon.Timestamp) error {
	_ = "STUB: not implemented"
	return nil
}

// If we only have job + instance, then target_info isn't useful, so don't add it.

// TODO what to do with this in case of full utf-8 support?

// We need at least one identifying label to generate target_info.

// convert ns to ms

// addSampleWithLabels is a helper function to create and add a sample with labels
func (c *prometheusConverterV2) addSampleWithLabels(sampleValue float64, timestamp int64, noRecordedValue bool,
	baseName string, baseLabels []prompb.Label, labelName, labelValue string, metadata metadata,
) {
	_ = "STUB: not implemented"
	return
}

func (c *prometheusConverterV2) addSummaryDataPoints(dataPoints pmetric.SummaryDataPointSlice, resource pcommon.Resource, scope pcommon.InstrumentationScope,
	settings Settings, baseName string, metadata metadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Add sum and count samples

// Process quantiles

func (c *prometheusConverterV2) addHistogramDataPoints(dataPoints pmetric.HistogramDataPointSlice,
	resource pcommon.Resource, scope pcommon.InstrumentationScope, settings Settings, baseName string, metadata metadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the sum is unset, it indicates the _sum metric point should be
// omitted

// treat count as a sample in an individual TimeSeries

// cumulative count for conversion to cumulative histogram

// process each bound, based on histograms proto definition, # of buckets = # of explicit bounds + 1

// add le=+Inf bucket

// TODO implement exemplars support
