// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tlscheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tlscheckreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var errConfigNotTLSCheck = errors.New(`invalid config`)

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
