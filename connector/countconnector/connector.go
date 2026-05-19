// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package countconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

// count can count spans, span event, metrics, data points, log records or
// profiles and emit the counts onto a metrics pipeline.
type count struct {
	metricsConsumer consumer.Metrics
	component.StartFunc
	component.ShutdownFunc

	spansMetricDefs      map[string]metricDef[*ottlspan.TransformContext]
	spanEventsMetricDefs map[string]metricDef[*ottlspanevent.TransformContext]
	metricsMetricDefs    map[string]metricDef[*ottlmetric.TransformContext]
	dataPointsMetricDefs map[string]metricDef[*ottldatapoint.TransformContext]
	logsMetricDefs       map[string]metricDef[*ottllog.TransformContext]
	profilesMetricDefs   map[string]metricDef[*ottlprofile.TransformContext]
}

func (*count) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (c *count) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// don't add an empty resource

func (c *count) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

//exhaustive:enforce

// don't add an empty resource

func (c *count) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// don't add an empty resource

func (c *count) ConsumeProfiles(ctx context.Context, ld pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

// don't add an empty resource
