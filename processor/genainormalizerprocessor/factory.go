// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package genainormalizerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/genainormalizerprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

// NewFactory returns a new factory for the genainormalizer processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// createDefaultConfig returns the default configuration. Sources must be
// explicitly specified by the user; there are no built-in source defaults.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}
