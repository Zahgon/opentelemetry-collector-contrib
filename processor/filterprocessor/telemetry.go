// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/otel/metric"
)

type filterTelemetry struct {
	attr    metric.MeasurementOption
	counter metric.Int64Counter
}

func newFilterTelemetry(set processor.Settings, signal pipeline.Signal) (*filterTelemetry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fpt *filterTelemetry) record(ctx context.Context, dropped int64) {
	_ = "STUB: not implemented"
	return
}
