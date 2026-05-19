// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfxexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/featuregate"
)

const (
	defaultHTTPTimeout          = time.Second * 10
	defaultHTTP2ReadIdleTimeout = time.Second * 10
	defaultHTTP2PingTimeout     = time.Second * 10
	defaultMaxConns             = 100

	defaultDimMaxBuffered         = 10000
	defaultDimSendDelay           = 10 * time.Second
	defaultDimMaxConnsPerHost     = 20
	defaultDimMaxIdleConns        = 20
	defaultDimMaxIdleConnsPerHost = 20
)

var entityEventsFeatureGate = featuregate.GlobalRegistry().MustRegister(
	"exporter.signalfx.consumeEntityEvents",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("Process entity events from logs pipeline and convert to dimension property updates"),
	featuregate.WithRegisterFromVersion("v0.145.0"))

// NewFactory creates a factory for SignalFx exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	eCfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

// explicitly disable since we rely on http.Client timeout logic.

// If AccessTokenPassthrough enabled, split the incoming Metrics data by splunk.SFxAccessTokenLabel,
// this ensures that we get batches of data for the same token when pushing to the backend.

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// explicitly disable since we rely on http.Client timeout logic.

// If AccessTokenPassthrough enabled, split the incoming Metrics data by splunk.SFxAccessTokenLabel,
// this ensures that we get batches of data for the same token when pushing to the backend.
