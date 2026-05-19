// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package icmpcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/icmpcheckreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var errConfigNotICMPCheck = errors.New("config was not a ICMP check receiver config")

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
