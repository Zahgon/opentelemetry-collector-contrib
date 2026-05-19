// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"
)

type sumologicSubprocessor interface {
	processLogs(plog.Logs) error
	processMetrics(pmetric.Metrics) error
	processTraces(ptrace.Traces) error
	isEnabled() bool
	ConfigPropertyName() string
}

type sumologicProcessor struct {
	logger        *zap.Logger
	subprocessors []sumologicSubprocessor
}

func newsumologicProcessor(set processor.Settings, config *Config) *sumologicProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (processor *sumologicProcessor) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (processor *sumologicProcessor) shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (processor *sumologicProcessor) processLogs(_ context.Context, logs plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (processor *sumologicProcessor) processMetrics(_ context.Context, metrics pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (processor *sumologicProcessor) processTraces(_ context.Context, traces ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
