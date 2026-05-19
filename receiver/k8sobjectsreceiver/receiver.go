// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobjectsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"k8s.io/client-go/dynamic"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sinventory"
)

type k8sobjectsreceiver struct {
	setting         receiver.Settings
	config          *Config
	objects         []*K8sObjectsConfig
	stopperChanList []chan struct{}
	client          dynamic.Interface
	consumer        consumer.Logs
	obsrecv         *receiverhelper.ObsReport
	storageClient   storage.Client
	mu              sync.Mutex
	cancel          context.CancelFunc
	observerFunc    func(ctx context.Context, object *K8sObjectsConfig) (k8sinventory.Observer, error)
	wg              sync.WaitGroup
}

func newReceiver(params receiver.Settings, config *Config, consumer consumer.Logs) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// Set default interval if in PullMode and interval is 0

func getObserverFunc(kr *k8sobjectsreceiver) func(ctx context.Context, object *K8sObjectsConfig) (k8sinventory.Observer, error) {
	_ = "STUB: not implemented"
	return nil
}

func (kr *k8sobjectsreceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize storage client for resource version persistence if storage is configured

// Validate objects against K8s API

// Register callbacks with the leader elector extension. These callbacks remain active
// for the lifetime of the receiver, allowing it to restart when leadership is regained.

// Shutdown on leader loss. The receiver will restart if leadership is regained
// since the callbacks remain registered with the leader elector extension.

func (kr *k8sobjectsreceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Close storage client if it exists

func (kr *k8sobjectsreceiver) stopWatches() { _ = "STUB: not implemented"; return }

func (kr *k8sobjectsreceiver) start(ctx context.Context, object *K8sObjectsConfig) error {
	_ = "STUB: not implemented"
	// Handle exclude_namespaces: compile regexes and filter namespaces
	// TODO: when using informers, we should find a way to get just the metadata.name of the namespace, and then filter on that
	return nil
}

// handleError handles errors according to the configured error mode
func (kr *k8sobjectsreceiver) handleError(err error, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

// This shouldn't happen as we validate ErrorMode during config validation

func getStorageClient(ctx context.Context, host component.Host, storageID *component.ID, componentID component.ID) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

// Make storage immune to component renames that add underscores to the component type.
// This is a workaround for https://github.com/open-telemetry/opentelemetry-collector/issues/14988.
