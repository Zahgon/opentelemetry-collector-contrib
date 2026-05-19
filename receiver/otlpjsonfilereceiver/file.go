// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpjsonfilereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otlpjsonfilereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/xreceiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"
)

const (
	transport = "file"
)

// NewFactory creates a factory for file receiver
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

type Config struct {
	fileconsumer.Config `mapstructure:",squash"`
	StorageID           *component.ID `mapstructure:"storage"`
	ReplayFile          bool          `mapstructure:"replay_file"`
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

type otlpjsonfilereceiver struct {
	input     *fileconsumer.Manager
	id        component.ID
	storageID *component.ID
}

func (f *otlpjsonfilereceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *otlpjsonfilereceiver) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func createLogsReceiver(_ context.Context, settings receiver.Settings, configuration component.Config, logs consumer.Logs) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// Appends token.Attributes

func createMetricsReceiver(_ context.Context, settings receiver.Settings, configuration component.Config, metrics consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// Appends token.Attributes

func createTracesReceiver(_ context.Context, settings receiver.Settings, configuration component.Config, traces consumer.Traces) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

// Appends token.Attributes

func createProfilesReceiver(_ context.Context, settings receiver.Settings, configuration component.Config, profiles xconsumer.Profiles) (xreceiver.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xreceiver.Profiles), nil
}

// TODO Append token.Attributes

func appendToMap(attributes map[string]any, attr pcommon.Map) { _ = "STUB: not implemented"; return }
