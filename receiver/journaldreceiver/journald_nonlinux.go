// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package journaldreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
)

var _ adapter.LogReceiverType = (*receiverType)(nil)

// newFactoryAdapter creates a dummy factory.
func newFactoryAdapter() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createLogsReceiver(
	_ context.Context,
	_ receiver.Settings,
	_ component.Config,
	_ consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
