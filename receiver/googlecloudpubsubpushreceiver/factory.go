// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubpushreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubpushreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createLogsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextLogs consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
