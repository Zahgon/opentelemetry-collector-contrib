// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logdedupprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/logdedupprocessor"

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
)

// logDedupProcessor is a logDedupProcessor that counts duplicate instances of logs.
type logDedupProcessor struct {
	emitInterval time.Duration
	conditions   *ottl.ConditionSequence[*ottllog.TransformContext]
	aggregator   *logAggregator
	remover      *fieldRemover
	nextConsumer consumer.Logs
	logger       *zap.Logger
	cancel       context.CancelFunc
	wg           sync.WaitGroup
	mux          sync.Mutex
}

func newProcessor(cfg *Config, nextConsumer consumer.Logs, settings processor.Settings) (*logDedupProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This should not happen due to config validation but we check anyways.

// Start starts the processor.
func (p *logDedupProcessor) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Capabilities returns the consumer's capabilities.
func (*logDedupProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// Shutdown stops the processor.
func (p *logDedupProcessor) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// Call cancel to stop the export interval goroutine and wait for it to finish.

// ConsumeLogs processes the logs.
func (p *logDedupProcessor) ConsumeLogs(ctx context.Context, pl plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// immediately consume any logs that didn't match any conditions

func (p *logDedupProcessor) aggregateLog(logRecord plog.LogRecord, scope pcommon.InstrumentationScope, resource pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

// handleExportInterval sends metrics at the configured interval.
func (p *logDedupProcessor) handleExportInterval(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Export any remaining logs

// exportLogs exports the logs to the next consumer.
func (p *logDedupProcessor) exportLogs(ctx context.Context) { _ = "STUB: not implemented"; return }

// Only send logs if we have some
