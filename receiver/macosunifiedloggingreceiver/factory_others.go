// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !darwin

package macosunifiedloggingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/macosunifiedloggingreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func newFactoryAdapter() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createLogsReceiverOthers(
	_ context.Context,
	_ receiver.Settings,
	_ component.Config,
	_ consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
