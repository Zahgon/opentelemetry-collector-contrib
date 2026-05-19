// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package lokireceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultGRPCEndpoint = "localhost:3600"
	defaultHTTPEndpoint = "localhost:3500"
)

// NewFactory return a new receiver.Factory for loki receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createLogsReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
