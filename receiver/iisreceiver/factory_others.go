// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package iisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// createMetricsReceiver creates a metrics receiver based on provided config.
func createMetricsReceiver(
	_ context.Context,
	_ receiver.Settings,
	_ component.Config,
	_ consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
