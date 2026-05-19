// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expvarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultPath     = "/debug/vars"
	defaultEndpoint = "http://localhost:8000" + defaultPath
	defaultTimeout  = 3 * time.Second
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func newMetricsReceiver(
	_ context.Context,
	set receiver.Settings,
	rCfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func newDefaultConfig() component.Config { _ = "STUB: not implemented"; return *new(component.Config) }
