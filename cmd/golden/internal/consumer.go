// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/cmd/golden/internal"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

var _ consumer.Metrics = (*MetricsSink)(nil)

// AttemptError tracks an error from a specific comparison attempt
type AttemptError struct {
	AttemptNumber int
	Error         error
}

type MetricsSink struct {
	cfg            *Config
	noExpected     bool
	expected       pmetric.Metrics
	DoneChan       chan struct{}
	Errors         []AttemptError
	attemptCounter int
}

func (*MetricsSink) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (m *MetricsSink) ConsumeMetrics(_ context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Clear errors on success

// Append error with attempt number

func NewConsumer(cfg *Config) (*MetricsSink, error) { _ = "STUB: not implemented"; return nil, nil }
