// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8seventsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/dynamic"
)

type k8seventsReceiver struct {
	config          *Config
	settings        receiver.Settings
	logsConsumer    consumer.Logs
	stopperChanList []chan struct{}
	startTime       time.Time
	ctx             context.Context
	cancel          context.CancelFunc
	obsrecv         *receiverhelper.ObsReport
	mu              sync.Mutex
	client          dynamic.Interface
	storageClient   storage.Client
	wg              sync.WaitGroup
}

// newReceiver creates the Kubernetes events receiver with the given configuration.
func newReceiver(
	set receiver.Settings,
	config *Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func (kr *k8seventsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize storage client for resource version persistence if storage is configured.

// Register callbacks with the leader elector extension. These callbacks remain active
// for the lifetime of the receiver, allowing it to restart when leadership is regained.

// Shutdown on leader loss. The receiver will restart if leadership is regained
// since the callbacks remain registered with the leader elector extension.

// No leader election: start immediately.

func (kr *k8seventsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Close storage client if it exists.

// startWatchers creates and starts the k8sinventory watch observer
func (kr *k8seventsReceiver) startWatchers() {
	_ = "STUB: not implemented"
	// Events GVR (GroupVersionResource)
	return
}

// Empty string means all namespaces

// Don't send initial state, only new events
// Skip DELETED events (matches old Informer behavior)

// The k8sinventory watch observer uses dynamic client which returns unstructured objects
// We need to convert them to corev1.Event

// Convert unstructured to corev1.Event

// handleEvent processes a Kubernetes event and sends it to the logs consumer
func (kr *k8seventsReceiver) handleEvent(ev *corev1.Event) { _ = "STUB: not implemented"; return }

// Allow events with eventTimestamp(EventTime/LastTimestamp/FirstTimestamp)
// not older than the receiver start time so that
// event flood can be avoided upon startup.
func (kr *k8seventsReceiver) allowEvent(ev *corev1.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func getStorageClient(ctx context.Context, host component.Host, storageID *component.ID, componentID component.ID) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

// Make storage immune to component renames that add underscores to the component type.
// This is a workaround for https://github.com/open-telemetry/opentelemetry-collector/issues/14988.
