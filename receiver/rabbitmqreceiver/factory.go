// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package rabbitmqreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var errConfigNotRabbit = errors.New("config was not a RabbitMQ receiver config")

// NewFactory creates a new receiver factory
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(_ context.Context, params receiver.Settings, rConf component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
