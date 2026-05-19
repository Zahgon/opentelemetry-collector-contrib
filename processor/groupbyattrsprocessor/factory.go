// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package groupbyattrsprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

var consumerCapabilities = consumer.Capabilities{MutatesData: true}

// NewFactory returns a new factory for the Filter processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// createDefaultConfig creates the default configuration for the processor.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createGroupByAttrsProcessor(set processor.Settings, attributes []string) (*groupByAttrsProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
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

// createLogsProcessor creates a logs processor based on this config.
func createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

// createMetricsProcessor creates a metrics processor based on this config.
func createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}
