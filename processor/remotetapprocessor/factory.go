// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/remotetapprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

var processors = sharedcomponent.NewSharedComponents()

func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createMetricsProcessor(ctx context.Context, params processor.Settings, cfg component.Config, c consumer.Metrics) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func createLogsProcessor(ctx context.Context, params processor.Settings, cfg component.Config, c consumer.Logs) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func createTraceProcessor(ctx context.Context, params processor.Settings, cfg component.Config, c consumer.Traces) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}
