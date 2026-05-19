// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/service/hostcapabilities"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

var _ receiver.Metrics = (*receiverCreator)(nil)

// receiverCreator implements consumer.Metrics.
type receiverCreator struct {
	params               receiver.Settings
	cfg                  *Config
	nextLogsConsumer     consumer.Logs
	nextMetricsConsumer  consumer.Metrics
	nextTracesConsumer   consumer.Traces
	nextProfilesConsumer xconsumer.Profiles
	observerHandler      *observerHandler
	observables          []observer.Observable
}

func newReceiverCreator(params receiver.Settings, cfg *Config) receiver.Metrics {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics)
}

// host is an interface that the component.Host passed to receivercreator's Start function must implement
type host interface {
	component.Host
	hostcapabilities.ComponentFactory
}

// Start receiver_creator.
func (rc *receiverCreator) Start(_ context.Context, h component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Match all configured observables to the extensions that are running.

// Make sure all observables are present before starting any.

// Start all configured watchers.

// Shutdown stops the receiver_creator and all its receivers started at runtime.
func (rc *receiverCreator) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
