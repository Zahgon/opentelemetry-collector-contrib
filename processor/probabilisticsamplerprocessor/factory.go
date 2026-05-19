// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package probabilisticsamplerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

// The default precision is 4 hex digits, slightly more the original
// component logic's 14-bits of precision.
const defaultPrecision = 4

// NewFactory returns a new factory for the Probabilistic sampler processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createTracesProcessor creates a trace processor based on this config.
func createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

// createLogsProcessor creates a log processor based on this config.
func createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}
