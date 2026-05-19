// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package intervalprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/intervalprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

// NewFactory returns a new factory for the Metrics Generation processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsProcessor(_ context.Context, set processor.Settings, cfg component.Config, nextConsumer consumer.Metrics) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}
