// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package groupbyattrsprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor/internal/metadata"
)

type groupByAttrsProcessor struct {
	logger           *zap.Logger
	groupByKeys      []string
	telemetryBuilder *metadata.TelemetryBuilder
}

// ProcessTraces process traces and groups traces by attribute.
func (gap *groupByAttrsProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// Some attributes are going to be moved from span to resource level,
// so we can delete those on the record level

// Lets combine the base resource attributes + the extracted (grouped) attributes
// and keep them in the grouping entry

// Copy the grouped data into output

func (gap *groupByAttrsProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// Some attributes are going to be moved from log record to resource level,
// so we can delete those on the record level

// Lets combine the base resource attributes + the extracted (grouped) attributes
// and keep them in the grouping entry

// Copy the grouped data into output

func (gap *groupByAttrsProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

//exhaustive:enforce

func deleteAttributes(attrsForRemoval, targetAttrs pcommon.Map) { _ = "STUB: not implemented"; return }

// extractGroupingAttributes extracts the keys and values of the specified Attributes
// that match with the attributes keys that is used for grouping
// Returns:
//   - whether any attribute matched (true) or none (false)
//   - the extracted AttributeMap of matching keys and their corresponding values
func (gap *groupByAttrsProcessor) extractGroupingAttributes(attrMap pcommon.Map) (bool, pcommon.Map) {
	_ = "STUB: not implemented"
	return false, *new(pcommon.Map)
}

// Searches for metric with same name in the specified InstrumentationLibrary and returns it. If nothing is found, create it.
func getMetricInInstrumentationLibrary(ilm pmetric.ScopeMetrics, searchedMetric pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	// Loop through all metrics and try to find the one that matches with the one we search for
	// (name and type)
	return *new(pmetric.Metric)
}

// We're here, which means that we haven't found our metric, so we need to create a new one, with the same name and type

// Move other special type specific values
//exhaustive:enforce

// Returns the Metric in the appropriate Resource matching with the specified Attributes
func (gap *groupByAttrsProcessor) getGroupedMetricsFromAttributes(
	ctx context.Context,
	mg *metricsGroup,
	originResourceMetrics pmetric.ResourceMetrics,
	ilm pmetric.ScopeMetrics,
	metric pmetric.Metric,
	attributes pcommon.Map,
) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// These attributes are going to be moved from datapoint to resource level,
// so we can delete those on the datapoint

// Get the ResourceMetrics matching with these attributes

// Get the corresponding instrumentation library

// Return the metric in this resource
