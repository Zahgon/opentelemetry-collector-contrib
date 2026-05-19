// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
)

type logsDataConsumer interface {
	consumeLogsJSON(ctx context.Context, json []byte) error
	setNextLogsConsumer(nextLogsConsumer consumer.Logs)
}

type tracesDataConsumer interface {
	consumeTracesJSON(ctx context.Context, json []byte) error
	setNextTracesConsumer(nextracesConsumer consumer.Traces)
}

type blobReceiver struct {
	blobEventHandler   eventHandler
	logger             *zap.Logger
	logsUnmarshaler    plog.Unmarshaler
	tracesUnmarshaler  ptrace.Unmarshaler
	nextLogsConsumer   consumer.Logs
	nextTracesConsumer consumer.Traces
	obsrecv            *receiverhelper.ObsReport
}

func (b *blobReceiver) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *blobReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *blobReceiver) setNextLogsConsumer(nextLogsConsumer consumer.Logs) {
	_ = "STUB: not implemented"
	return
}

func (b *blobReceiver) setNextTracesConsumer(nextTracesConsumer consumer.Traces) {
	_ = "STUB: not implemented"
	return
}

func (b *blobReceiver) consumeLogsJSON(ctx context.Context, json []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *blobReceiver) consumeTracesJSON(ctx context.Context, json []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Returns a new instance of the log receiver
func newReceiver(set receiver.Settings, eventHandler eventHandler) (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}
