// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/logstransformprocessor"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"
)

type logsTransformProcessor struct {
	set    component.TelemetrySettings
	config *Config

	consumer consumer.Logs

	pipe          *pipeline.DirectedPipeline
	firstOperator operator.Operator
	emitter       helper.LogEmitter
	fromConverter *adapter.FromPdataConverter
	shutdownFns   []component.ShutdownFunc
}

func newProcessor(config *Config, nextConsumer consumer.Logs, set component.TelemetrySettings) (*logsTransformProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*logsTransformProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (ltp *logsTransformProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// We call the shutdown functions in reverse order, so that the last thing we started
// is stopped first.

func (ltp *logsTransformProcessor) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// data flows in this order:
// ConsumeLogs: receives logs and forwards them for conversion to stanza format ->
// fromConverter: converts logs to stanza format ->
// converterLoop: forwards converted logs to the stanza pipeline ->
// pipeline: performs user configured operations on the logs ->
// transformProcessor: receives []*entry.Entries, converts them to plog.Logs and sends the converted OTLP logs to the next consumer
//
// We should start these components in reverse order of the data flow, then stop them in order of the data flow,
// in order to allow for pipeline draining.

func (ltp *logsTransformProcessor) startFromConverter() { _ = "STUB: not implemented"; return }

// startConverterLoop starts the converter loop, which reads all the logs translated by the fromConverter and then forwards
// them to pipeline
func (ltp *logsTransformProcessor) startConverterLoop(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (ltp *logsTransformProcessor) startPipeline() error {
	_ = "STUB: not implemented"
	// There is no need for this processor to use storage
	return nil
}

func (ltp *logsTransformProcessor) ConsumeLogs(_ context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	// Add the logs to the chain
	return nil
}

// converterLoop reads the log entries produced by the fromConverter and sends them
// into the pipeline
func (ltp *logsTransformProcessor) converterLoop(ctx context.Context, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// Add item to the first operator of the pipeline manually

func (ltp *logsTransformProcessor) consumeStanzaLogEntries(ctx context.Context, entries []*entry.Entry) {
	_ = "STUB: not implemented"
	return
}
