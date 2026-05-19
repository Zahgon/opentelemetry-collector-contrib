// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3receiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createTracesReceiver(ctx context.Context, settings receiver.Settings, cc component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func createMetricsReceiver(ctx context.Context, settings receiver.Settings, cc component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func createLogsReceiver(ctx context.Context, settings receiver.Settings, cc component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
