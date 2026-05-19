// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logdedupprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/logdedupprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

// NewFactory creates a new factory for the processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// createLogsProcessor creates a log processor.
func createLogsProcessor(_ context.Context, settings processor.Settings, cfg component.Config, consumer consumer.Logs) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}
