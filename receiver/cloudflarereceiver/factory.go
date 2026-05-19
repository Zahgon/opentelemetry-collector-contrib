// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudflarereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudflarereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// NewFactory returns the component factory for the cloudflarereceiver
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
