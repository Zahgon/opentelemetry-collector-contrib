// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	reportTransport      = "pubsub"
	reportFormatProtobuf = "protobuf"
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

type pubsubReceiverFactory struct {
	receivers map[*Config]*pubsubReceiver
}

func (*pubsubReceiverFactory) CreateDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (factory *pubsubReceiverFactory) ensureReceiver(settings receiver.Settings, config component.Config) (*pubsubReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (factory *pubsubReceiverFactory) CreateTraces(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func (factory *pubsubReceiverFactory) CreateMetrics(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func (factory *pubsubReceiverFactory) CreateLogs(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
