// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package podmanreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func newMetricsReceiver(
	_ receiver.Settings,
	_ *Config,
	_ consumer.Metrics,
	_ any,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	config component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
