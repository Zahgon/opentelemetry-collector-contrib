// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/xprocessor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

type filterProcessorFactory struct {
	resourceFunctions                   map[string]ottl.Factory[*ottlresource.TransformContext]
	dataPointFunctions                  map[string]ottl.Factory[*ottldatapoint.TransformContext]
	logFunctions                        map[string]ottl.Factory[*ottllog.TransformContext]
	metricFunctions                     map[string]ottl.Factory[*ottlmetric.TransformContext]
	spanEventFunctions                  map[string]ottl.Factory[*ottlspanevent.TransformContext]
	spanFunctions                       map[string]ottl.Factory[*ottlspan.TransformContext]
	profileFunctions                    map[string]ottl.Factory[*ottlprofile.TransformContext]
	defaultResourceFunctionsOverridden  bool
	defaultDataPointFunctionsOverridden bool
	defaultLogFunctionsOverridden       bool
	defaultMetricFunctionsOverridden    bool
	defaultSpanEventFunctionsOverridden bool
	defaultSpanFunctionsOverridden      bool
	defaultProfileFunctionsOverridden   bool
}

// FactoryOption applies changes to filterProcessorFactory.
type FactoryOption func(factory *filterProcessorFactory)

// WithResourceFunctions will override the default OTTL resource context functions with the provided resourceFunctions in resulting processor.
// Subsequent uses of WithResourceFunctions will merge the provided resourceFunctions with the previously registered functions.
func WithResourceFunctions(resourceFunctions []ottl.Factory[*ottlresource.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// Deprecated: [v0.152.0] Use WithResourceFunctions.
func WithResourceFunctionsNew(resourceFunctions []ottl.Factory[*ottlresource.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// WithDataPointFunctions will override the default OTTL datapoint context functions with the provided dataPointFunctions in resulting processor.
// Subsequent uses of WithDataPointFunctions will merge the provided dataPointFunctions with the previously registered functions.
func WithDataPointFunctions(dataPointFunctions []ottl.Factory[*ottldatapoint.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// Deprecated: [v0.152.0] Use WithDataPointFunctions.
func WithDataPointFunctionsNew(dataPointFunctions []ottl.Factory[*ottldatapoint.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// WithLogFunctions will override the default OTTL log context functions with the provided logFunctions in the resulting processor.
// Subsequent uses of WithLogFunctions will merge the provided logFunctions with the previously registered functions.
func WithLogFunctions(logFunctions []ottl.Factory[*ottllog.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// Deprecated: [v0.152.0] Use WithLogFunctions.
func WithLogFunctionsNew(logFunctions []ottl.Factory[*ottllog.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// WithMetricFunctions will override the default OTTL metric context functions with the provided metricFunctions in the resulting processor.
// Subsequent uses of WithMetricFunctions will merge the provided metricFunctions with the previously registered functions.
func WithMetricFunctions(metricFunctions []ottl.Factory[*ottlmetric.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// Deprecated: [v0.152.0] Use WithMetricFunctions.
func WithMetricFunctionsNew(metricFunctions []ottl.Factory[*ottlmetric.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// WithSpanEventFunctions will override the default OTTL spanevent context functions with the provided spanEventFunctions in the resulting processor.
// Subsequent uses of WithSpanEventFunctions will merge the provided spanEventFunctions with the previously registered functions.
func WithSpanEventFunctions(spanEventFunctions []ottl.Factory[*ottlspanevent.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// Deprecated: [v0.152.0] Use WithSpanEventFunctions.
func WithSpanEventFunctionsNew(spanEventFunctions []ottl.Factory[*ottlspanevent.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// WithSpanFunctions will override the default OTTL span context functions with the provided spanFunctions in the resulting processor.
// Subsequent uses of WithSpanFunctions will merge the provided spanFunctions with the previously registered functions.
func WithSpanFunctions(spanFunctions []ottl.Factory[*ottlspan.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// Deprecated: [v0.152.0] use WithSpanFunctions.
func WithSpanFunctionsNew(spanFunctions []ottl.Factory[*ottlspan.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// WithProfileFunctions will override the default OTTL profile context functions with the provided profileFunctions in the resulting processor.
// Subsequent uses of WithProfileFunctions will merge the provided profileFunctions with the previously registered functions.
func WithProfileFunctions(profileFunctions []ottl.Factory[*ottlprofile.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// Deprecated: [v0.152.0] use WithProfileFunctions.
func WithProfileFunctionsNew(profileFunctions []ottl.Factory[*ottlprofile.TransformContext]) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// NewFactory returns a new factory for the Filter processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// NewFactoryWithOptions can receive FactoryOption like With*Functions to register non-default OTTL functions in the resulting processor.
func NewFactoryWithOptions(options ...FactoryOption) processor.Factory {
	_ = "STUB: not implemented"
	return *new(processor.Factory)
}

func (f *filterProcessorFactory) createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (f *filterProcessorFactory) createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func (f *filterProcessorFactory) createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func (f *filterProcessorFactory) createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func (f *filterProcessorFactory) createProfilesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer xconsumer.Profiles,
) (xprocessor.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xprocessor.Profiles), nil
}
