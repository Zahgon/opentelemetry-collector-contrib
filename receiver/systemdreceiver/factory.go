// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package systemdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/systemdreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	errConfigNotValid = errors.New("config not valid for the systemd receiver")
	errNonLinux       = errors.New("systemd receiver is only supported on Linux")
)

// NewFactory creates a factory for systemd receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(_ context.Context, params receiver.Settings, rConf component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
