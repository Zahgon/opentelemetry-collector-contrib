// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tcpcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcpcheckreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// NewFactory creates a factory for tcpcheckreceiver receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func newDefaultConfig() component.Config { _ = "STUB: not implemented"; return *new(component.Config) }

func newReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
