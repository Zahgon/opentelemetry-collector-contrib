// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/faroreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

const (
	defaultFaroEndpoint = "localhost:8080"
)

// This is the map of already created Faro receivers for particular configurations.
// We maintain this map because the receiver.Factory is asked trace and log receivers separately
// when it gets createFaroReceiverTraces() and createFaroReceiverLogs() but they must not
// create separate objects, they must use one faroReceiver object per configuration.
// When the receiver is shutdown it should be removed from this map so the same configuration
// can be recreated successfully.
var receivers = sharedcomponent.NewSharedComponents()

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func newFaroReceiverFactory(fCfg *Config, set *receiver.Settings, err *error) func() component.Component {
	_ = "STUB: not implemented"
	return nil
}

func createFaroReceiverTraces(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextTraces consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func createFaroReceiverLogs(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextLogs consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}
