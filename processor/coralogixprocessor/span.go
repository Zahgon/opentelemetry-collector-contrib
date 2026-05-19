// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/coralogixprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"
)

type coralogixProcessor struct {
	config *Config
	component.StartFunc
	component.ShutdownFunc
	logger *zap.Logger
}

func newCoralogixProcessor(ctx context.Context, set processor.Settings, cfg *Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func (sp *coralogixProcessor) processTraces(_ context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	//nolint:staticcheck // QF1008: Keeping embedded field for clarity
	return *new(ptrace.Traces), nil
}
