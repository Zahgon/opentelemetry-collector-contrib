// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package envoyalsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/envoyalsreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultGRPCEndpoint = "localhost:19001"
)

// NewFactory creates a new ALS receiver factory.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func newReceiver(
	_ context.Context,
	st receiver.Settings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
