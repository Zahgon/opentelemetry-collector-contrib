// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package lookupprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/lookupsource"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

var Type = metadata.Type

type FactoryOption func(*lookupProcessorFactory)

// WithSources adds custom source factories to the processor.
// This REPLACES the default sources on first call, then MERGES on subsequent calls.
// (Same pattern as transform processor's WithXxxFunctions)
//
// Example:
//
//	lookupprocessor.NewFactoryWithOptions(
//	    lookupprocessor.WithSources(httplookup.NewFactory()),
//	)
func WithSources(factories ...lookupsource.SourceFactory) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type lookupProcessorFactory struct {
	sources                  map[string]lookupsource.SourceFactory
	defaultSourcesOverridden bool
}

func defaultSources() map[string]lookupsource.SourceFactory { _ = "STUB: not implemented"; return nil }

func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// NewFactoryWithOptions creates a lookup processor factory with custom sources.
//
// Example (third-party HTTP source):
//
//	import (
//	    "github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor"
//	    "github.com/user/otel-lookup-http/httplookup"
//	)
//
//	factories.Processors[lookupprocessor.Type] = lookupprocessor.NewFactoryWithOptions(
//	    lookupprocessor.WithSources(httplookup.NewFactory()),
//	)
func NewFactoryWithOptions(options ...FactoryOption) processor.Factory {
	_ = "STUB: not implemented"
	return *new(processor.Factory)
}

func (*lookupProcessorFactory) createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (f *lookupProcessorFactory) createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

// parsedLookup holds a lookup config with its pre-parsed OTTL key expression.
type parsedLookup struct {
	keyExpr    *ottl.ValueExpression[*ottllog.TransformContext]
	context    ContextID
	attributes []AttributeMapping
}

func parseLookups(parser ottl.Parser[*ottllog.TransformContext], configs []LookupConfig) ([]parsedLookup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *lookupProcessorFactory) createSource(
	ctx context.Context,
	set processor.Settings,
	cfg *Config,
) (lookupsource.Source, error) {
	_ = "STUB: not implemented"
	return *new(lookupsource.Source), nil
}

// Decode the raw source config captured by mapstructure's ",remain" tag
// into the source's typed config struct. See SourceConfig for why this
// is deferred to factory time rather than config unmarshal time.
