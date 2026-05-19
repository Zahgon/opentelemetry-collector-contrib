// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gitlabreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/gitlabreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func createTracesReceiver(_ context.Context, params receiver.Settings, cfg component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	// check that the configuration is valid
	return *new(receiver.Traces), nil
}

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }
