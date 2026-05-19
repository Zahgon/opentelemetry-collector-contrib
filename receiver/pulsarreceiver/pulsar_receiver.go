// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"

import (
	"context"
	"errors"

	"github.com/apache/pulsar-client-go/pulsar"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

var errUnrecognizedEncoding = errors.New("unrecognized encoding")

const alreadyClosedError = "AlreadyClosedError"

const transport = "pulsar"

type pulsarTracesConsumer struct {
	tracesConsumer  consumer.Traces
	topic           string
	client          pulsar.Client
	cancel          context.CancelFunc
	consumer        pulsar.Consumer
	unmarshaler     TracesUnmarshaler
	settings        receiver.Settings
	consumerOptions pulsar.ConsumerOptions
	obsrecv         *receiverhelper.ObsReport
}

func newTracesReceiver(config Config, set receiver.Settings, unmarshalers map[string]TracesUnmarshaler, nextConsumer consumer.Traces) (*pulsarTracesConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pulsarTracesConsumer) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func consumerTracesLoop(ctx context.Context, c *pulsarTracesConsumer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Ensure returned errors are handled

func (c *pulsarTracesConsumer) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type pulsarMetricsConsumer struct {
	metricsConsumer consumer.Metrics
	unmarshaler     MetricsUnmarshaler
	topic           string
	client          pulsar.Client
	consumer        pulsar.Consumer
	cancel          context.CancelFunc
	settings        receiver.Settings
	consumerOptions pulsar.ConsumerOptions
	obsrecv         *receiverhelper.ObsReport
}

func newMetricsReceiver(config Config, set receiver.Settings, unmarshalers map[string]MetricsUnmarshaler, nextConsumer consumer.Metrics) (*pulsarMetricsConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pulsarMetricsConsumer) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func consumeMetricsLoop(ctx context.Context, c *pulsarMetricsConsumer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Ensure returned errors are handled

func (c *pulsarMetricsConsumer) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type pulsarLogsConsumer struct {
	logsConsumer    consumer.Logs
	unmarshaler     LogsUnmarshaler
	topic           string
	client          pulsar.Client
	consumer        pulsar.Consumer
	cancel          context.CancelFunc
	settings        receiver.Settings
	consumerOptions pulsar.ConsumerOptions
	obsrecv         *receiverhelper.ObsReport
}

func newLogsReceiver(config Config, set receiver.Settings, unmarshalers map[string]LogsUnmarshaler, nextConsumer consumer.Logs) (*pulsarLogsConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pulsarLogsConsumer) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func consumeLogsLoop(ctx context.Context, c *pulsarLogsConsumer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *pulsarLogsConsumer) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
