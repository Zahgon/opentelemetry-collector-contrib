// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vcrreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcrreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/xreceiver"
)

/*type rawFile struct {
	path       string
	signalType string
	sequence   int
}*/

type tapeFile struct {
	path    string
	startNs int64
	endNs   int64
}

type tapeLoader struct {
	dir         []string
	signalType  string
	loadedFiles map[string]tapeFile
}

type Config struct {
	IncludeRaw  []string `mapstructure:"include_raw"`
	ExcludeRaw  []string `mapstructure:"exclude_raw"`
	IncludeTape []string `mapstructure:"include_tape"`
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// default to current directory

type vcrReceiver struct {
	loader *tapeLoader
}

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func (*vcrReceiver) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*vcrReceiver) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func newTapeLoader(dir []string, signalType string) (*tapeLoader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *tapeLoader) loadTapeFiles() error { _ = "STUB: not implemented"; return nil }

// Skip files that don't match the pattern

// matches[0] = full string
// matches[1] = signal type (metrics|traces|logs|profiles)
// matches[2] = start timestamp
// matches[3] = end timestamp

func createMetricsReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	_ consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// start playback tape loop here

func createTracesReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	_ consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

// start playback tape loop here

func createLogsReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	_ consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// start playback tape loop here

func createProfilesReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	_ xconsumer.Profiles,
) (xreceiver.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xreceiver.Profiles), nil
}

// start playback tape loop here
