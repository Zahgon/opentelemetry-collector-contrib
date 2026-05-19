// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultGranularity   = "PT1M" // 1-minute, as per https://docs.atlas.mongodb.com/reference/api/process-measurements/
	defaultAlertsEnabled = false
	defaultLogsEnabled   = false
)

// NewFactory creates a factory for MongoDB Atlas receiver
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func createCombinedLogReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// Confirm at least one project is enabled for access logs before adding

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// reset default of 1 minute to be 3 minutes in order to avoid null values for some metrics that do not publish
// more frequently
