// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package snowflakereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snowflakereceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultInterval = 30 * time.Minute
	defaultRole     = "ACCOUNTADMIN"
	defaultDB       = "SNOWFLAKE"
	defaultSchema   = "ACCOUNT_USAGE"
)

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	baseCfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}
