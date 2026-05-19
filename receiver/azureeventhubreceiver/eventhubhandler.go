// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
)

type hubWrapper interface {
	GetRuntimeInformation(ctx context.Context) (*hubRuntimeInfo, error)
	Receive(ctx context.Context, partitionID string, handler hubHandler, applyOffset bool, logger *zap.Logger) (listenerHandleWrapper, error)
	Close(ctx context.Context) error
}

type listenerHandleWrapper interface {
	Done() <-chan struct{}
	Err() error
}

type hubHandler func(ctx context.Context, event *azureEvent) error

type hubRuntimeInfo struct {
	Path           string
	CreatedAt      time.Time
	PartitionCount int
	PartitionIDs   []string
}

var errNoConfig = errors.New("Configuration error, hub not accessible")

type eventhubHandler struct {
	hub            hubWrapper
	dataConsumer   dataConsumer
	config         *Config
	settings       receiver.Settings
	cancel         context.CancelFunc
	storageClient  storage.Client
	consumerClient *azeventhubs.ConsumerClient // non-nil when in distributed mode
	wg             sync.WaitGroup              // tracks goroutines spawned by runDistributed
}

func shouldInitializeStorageClient(storageClient storage.Client, storageID *component.ID) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *eventhubHandler) run(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *eventhubHandler) runSingle(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// set manually for testing.

// set manually for testing.

// listen to each partition of the Event Hub

func (h *eventhubHandler) runDistributed(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Dispatch partition clients as the Processor assigns them.

// Processor has stopped

// Run the processor's load balancer in a background goroutine.
// It returns when the context is cancelled.

func (h *eventhubHandler) setUpOnePartition(ctx context.Context, partitionID string, applyOffset bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *eventhubHandler) newMessageHandler(ctx context.Context, event *azureEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *eventhubHandler) close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Wait for goroutines spawned by runDistributed to finish before closing
// the consumer client they depend on.

func (h *eventhubHandler) setDataConsumer(dataConsumer dataConsumer) {
	_ = "STUB: not implemented"
	return
}

func newEventhubHandler(config *Config, settings receiver.Settings) *eventhubHandler {
	_ = "STUB: not implemented"
	return nil
}
