// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package schemaprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

// factory will store any of the precompiled schemas in future
type factory struct{}

// newDefaultConfiguration returns the configuration for schema processor
// with the default values being used throughout it
func newDefaultConfiguration() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func (factory) createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func (factory) createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func (factory) createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}
