// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclusterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/collection"
)

const (
	transport = "http"

	defaultInitialSyncTimeout = 10 * time.Minute
)

var _ receiver.Metrics = (*kubernetesReceiver)(nil)

type kubernetesReceiver struct {
	dataCollector   *collection.DataCollector
	resourceWatcher *resourceWatcher

	config          *Config
	settings        receiver.Settings
	metricsConsumer consumer.Metrics
	cancel          context.CancelFunc
	obsrecv         *receiverhelper.ObsReport
}

type getExporters interface {
	GetExporters() map[pipeline.Signal]map[component.ID]component.Component
}

func (kr *kubernetesReceiver) startReceiver(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Setup metadata exporters before initializing watchers to avoid concurrent access

// Wait till either the initial cache sync times out or until the cancel method
// corresponding to this context is called.

// If the context times out, set initialSyncTimedOut and report a fatal error. Currently
// this timeout is 10 minutes, which appears to be long enough.

func (kr *kubernetesReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// if extension is defined start with k8s leader elector

func (kr *kubernetesReceiver) stopReceiver() { _ = "STUB: not implemented"; return }

func (kr *kubernetesReceiver) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (kr *kubernetesReceiver) dispatchMetrics(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Metric collection is not enabled.

// newMetricsReceiver creates the Kubernetes cluster receiver with the given configuration.
func newMetricsReceiver(
	ctx context.Context, set receiver.Settings, cfg component.Config, consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// newMetricsReceiver creates the Kubernetes cluster receiver with the given configuration.
func newLogsReceiver(
	ctx context.Context, set receiver.Settings, cfg component.Config, consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// newMetricsReceiver creates the Kubernetes cluster receiver with the given configuration.
func newReceiver(_ context.Context, set receiver.Settings, cfg component.Config) (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}
