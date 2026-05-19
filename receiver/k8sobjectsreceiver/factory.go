// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobjectsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
