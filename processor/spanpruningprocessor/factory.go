// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package spanpruningprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanpruningprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

// NewFactory returns a new factory for the Span Pruning processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}
