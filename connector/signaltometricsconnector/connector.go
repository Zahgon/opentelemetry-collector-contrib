// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signaltometricsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/model"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

type signalToMetrics struct {
	next                  consumer.Metrics
	collectorInstanceInfo model.CollectorInstanceInfo
	logger                *zap.Logger
	errorMode             ottl.ErrorMode

	spanMetricDefs    []model.MetricDef[*ottlspan.TransformContext]
	dpMetricDefs      []model.MetricDef[*ottldatapoint.TransformContext]
	logMetricDefs     []model.MetricDef[*ottllog.TransformContext]
	profileMetricDefs []model.MetricDef[*ottlprofile.TransformContext]

	component.StartFunc
	component.ShutdownFunc
}

func (*signalToMetrics) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (sm *signalToMetrics) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// resAttrsCache lazily caches the filtered resource attributes per
// metric definition within a resource. Since resource attributes are
// constant for all signals within a resource, the result only needs
// to be computed once per metric definition per resource. The slice
// is allocated on the first match and reused across resources via
// clear(). A zero-value pcommon.Map entry indicates it has not been
// computed yet for the current resource.

func (sm *signalToMetrics) ConsumeMetrics(ctx context.Context, m pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// resAttrsCache lazily caches the filtered resource attributes per
// metric definition within a resource. Since resource attributes are
// constant for all signals within a resource, the result only needs
// to be computed once per metric definition per resource. The slice
// is allocated on the first match and reused across resources via
// clear(). A zero-value pcommon.Map entry indicates it has not been
// computed yet for the current resource.

//exhaustive:enforce

func (sm *signalToMetrics) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// resAttrsCache lazily caches the filtered resource attributes per
// metric definition within a resource. Since resource attributes are
// constant for all signals within a resource, the result only needs
// to be computed once per metric definition per resource. The slice
// is allocated on the first match and reused across resources via
// clear(). A zero-value pcommon.Map entry indicates it has not been
// computed yet for the current resource.

func (sm *signalToMetrics) ConsumeProfiles(ctx context.Context, profiles pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

// resAttrsCache lazily caches the filtered resource attributes per
// metric definition within a resource. Since resource attributes are
// constant for all signals within a resource, the result only needs
// to be computed once per metric definition per resource. The slice
// is allocated on the first match and reused across resources via
// clear(). A zero-value pcommon.Map entry indicates it has not been
// computed yet for the current resource.
