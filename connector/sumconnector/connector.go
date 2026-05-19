// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/sumconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

// sum can sum attribute values from spans, span event, metrics, data points, or log records
// and emit the sums onto a metrics pipeline.
type sum struct {
	metricsConsumer consumer.Metrics
	component.StartFunc
	component.ShutdownFunc

	spansMetricDefs      map[string]metricDef[*ottlspan.TransformContext]
	spanEventsMetricDefs map[string]metricDef[*ottlspanevent.TransformContext]
	metricsMetricDefs    map[string]metricDef[*ottlmetric.TransformContext]
	dataPointsMetricDefs map[string]metricDef[*ottldatapoint.TransformContext]
	logsMetricDefs       map[string]metricDef[*ottllog.TransformContext]
}

func (*sum) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (c *sum) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// don't add an empty resource

func (c *sum) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

//exhaustive:enforce
//  For metric types each must be handled in exactly the same way
//  Switch case required because each type calls DataPoints() differently

// don't add an empty resource

func (c *sum) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// don't add an empty resource
