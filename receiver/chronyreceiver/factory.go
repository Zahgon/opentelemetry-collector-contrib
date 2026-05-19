// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package chronyreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func newMetricsReceiver(
	ctx context.Context,
	set receiver.Settings,
	rCfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
