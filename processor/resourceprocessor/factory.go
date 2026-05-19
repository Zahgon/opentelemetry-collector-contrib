// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package resourceprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/xprocessor"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

// NewFactory returns a new factory for the Resource processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// Note: This isn't a valid configuration because the processor would do no work.
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

func createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func createProfilesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer xconsumer.Profiles,
) (xprocessor.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xprocessor.Profiles), nil
}
