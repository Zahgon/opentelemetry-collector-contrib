// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package redactionprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

// NewFactory creates a factory for the redaction processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"

	// createTracesProcessor creates an instance of redaction for processing traces
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

// TODO: Placeholder for an error metric in the next PR

// createLogsProcessor creates an instance of redaction for processing logs
func createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

// Attributes are defined for metrics and traces:
// https://opentelemetry.io/docs/specs/semconv/database/
// For logs, we don't rely on the "db.system.name" attribute to
// do the sanitization.

// createMetricsProcessor creates an instance of redaction for processing metrics
func createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}
