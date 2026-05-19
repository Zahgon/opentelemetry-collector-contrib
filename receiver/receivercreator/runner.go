// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	rcvr "go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/xreceiver"
	"go.uber.org/zap"
)

// runner starts and stops receiver instances.
type runner interface {
	// start a metrics receiver instance from its static config and discovered config.
	start(receiver receiverConfig, discoveredConfig userConfigMap, consumer *enhancingConsumer) (component.Component, error)
	// shutdown a receiver.
	shutdown(rcvr component.Component) error
}

// receiverRunner handles starting/stopping of a concrete subreceiver instance.
type receiverRunner struct {
	logger      *zap.Logger
	params      rcvr.Settings
	idNamespace component.ID
	host        host
	receivers   map[string]*wrappedReceiver
	lock        *sync.Mutex
}

func newReceiverRunner(params rcvr.Settings, host host) *receiverRunner {
	_ = "STUB: not implemented"
	return nil
}

var _ runner = (*receiverRunner)(nil)

func (run *receiverRunner) start(
	receiver receiverConfig,
	discoveredConfig userConfigMap,
	consumer *enhancingConsumer,
) (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}

// Sets dynamically created receiver to something like receiver_creator/1/redis{endpoint="localhost:6380"}/<EndpointID>.

// shutdown the given receiver.
func (*receiverRunner) shutdown(rcvr component.Component) error {
	_ = "STUB: not implemented"
	return nil
}

// loadRuntimeReceiverConfig loads the given receiverTemplate merged with config values
// that may have been discovered at runtime.
func (*receiverRunner) loadRuntimeReceiverConfig(
	factory rcvr.Factory,
	receiver receiverConfig,
	discoveredConfig userConfigMap,
) (component.Config, string, error) {
	_ = "STUB: not implemented"
	// remove dynamically added "endpoint" field if not supported by receiver
	return *new(component.Config), "", nil
}

// mergeTemplateAndDiscoveredConfigs will unify the templated and discovered configs,
// setting the `endpoint` field from the discovered one if 1. not specified by the user
// and 2. determined to be supported (by trial and error of unmarshalling a temp intermediary).
func mergeTemplatedAndDiscoveredConfigs(factory rcvr.Factory, templated, discovered userConfigMap) (*confmap.Conf, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// confirm the endpoint we've added is supported, removing if not

// we assume that the error is due to unused keys in the config, so we need to remove endpoint key

// Merge in discoveredConfig containing values discovered at runtime.

// createLogsRuntimeReceiver creates a receiver that is discovered at runtime.
func (run *receiverRunner) createLogsRuntimeReceiver(
	factory rcvr.Factory,
	id component.ID,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (rcvr.Logs, error) {
	_ = "STUB: not implemented"
	return *new(rcvr.Logs), nil
}

// createMetricsRuntimeReceiver creates a receiver that is discovered at runtime.
func (run *receiverRunner) createMetricsRuntimeReceiver(
	factory rcvr.Factory,
	id component.ID,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (rcvr.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(rcvr.Metrics), nil
}

// createTracesRuntimeReceiver creates a receiver that is discovered at runtime.
func (run *receiverRunner) createTracesRuntimeReceiver(
	factory rcvr.Factory,
	id component.ID,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (rcvr.Traces, error) {
	_ = "STUB: not implemented"
	return *new(rcvr.Traces), nil
}

// createProfilesRuntimeReceiver creates a receiver that is discovered at runtime.
func (run *receiverRunner) createProfilesRuntimeReceiver(
	factory rcvr.Factory,
	id component.ID,
	cfg component.Config,
	nextConsumer xconsumer.Profiles,
) (xreceiver.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xreceiver.Profiles), nil
}

var _ component.Component = (*wrappedReceiver)(nil)

type wrappedReceiver struct {
	logs     rcvr.Logs
	metrics  rcvr.Metrics
	traces   rcvr.Traces
	profiles xreceiver.Profiles
}

func (w *wrappedReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
