// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package deltatocumulativeprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createMetricsProcessor(_ context.Context, set processor.Settings, cfg component.Config, next consumer.Metrics) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}
