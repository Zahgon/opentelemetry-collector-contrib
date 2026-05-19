// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetry // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/telemetry"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/metadata"
)

func New(set component.TelemetrySettings) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

type Metrics struct {
	*metadata.TelemetryBuilder

	tracked *func() int
}

func (m *Metrics) Datapoints() Counter { _ = "STUB: not implemented"; return *new(Counter) }

func (m *Metrics) WithTracked(streams func() int) { _ = "STUB: not implemented"; return }

func Error(msg string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func Cause(err error) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Counter struct{ metric.Int64Counter }

func (c Counter) Inc(ctx context.Context, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}
