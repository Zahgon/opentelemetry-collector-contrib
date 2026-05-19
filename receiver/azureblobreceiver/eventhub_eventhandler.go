// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	"context"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/v2"
	"go.uber.org/zap"
)

type eventHubEventHandler struct {
	blobClient               blobClient
	logsDataConsumer         logsDataConsumer
	tracesDataConsumer       tracesDataConsumer
	logsContainerName        string
	tracesContainerName      string
	eventHubConnectionString string
	hub                      *azeventhubs.ConsumerClient
	logger                   *zap.Logger
	wg                       sync.WaitGroup
	cancelFunc               context.CancelFunc

	pollRate      int
	maxPollEvents int
	consumerGroup string
}

var _ eventHandler = (*eventHubEventHandler)(nil)

const (
	blobCreatedEventType = "Microsoft.Storage.BlobCreated"
)

func (p *eventHubEventHandler) run(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// The event hub name is empty because it's extracted from the connection string's EntityPath.
// The consumer group "$Default" is the default consumer group for Event Hubs.

func (p *eventHubEventHandler) receiveEvents(
	ctx context.Context,
	pc *azeventhubs.PartitionClient,
	handler func(ctx context.Context, event *azeventhubs.ReceivedEventData) error,
) {
	_ = "STUB: not implemented"
	return
}

func (p *eventHubEventHandler) newMessageHandler(ctx context.Context, event *azeventhubs.ReceivedEventData) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *eventHubEventHandler) processBlobCreatedEventType(ctx context.Context, containerName, blobName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *eventHubEventHandler) close(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for all partition receiver goroutines to finish

func (p *eventHubEventHandler) setLogsDataConsumer(logsDataConsumer logsDataConsumer) {
	_ = "STUB: not implemented"
	return
}

func (p *eventHubEventHandler) setTracesDataConsumer(tracesDataConsumer tracesDataConsumer) {
	_ = "STUB: not implemented"
	return
}

func newEventHubEventHandler(eventHubConnectionString, logsContainerName, tracesContainerName string, blobClient blobClient, logger *zap.Logger) *eventHubEventHandler {
	_ = "STUB: not implemented"
	return nil
}
