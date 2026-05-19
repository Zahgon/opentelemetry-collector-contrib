// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"context"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"
)

type checkpointSeqNumber struct {
	// Offset only used for backwards compatibility
	Offset         string `json:"offset"`
	SequenceNumber int64  `json:"sequenceNumber"`
}

// UnmarshalJSON is a custom unmarshaller to allow for backward compatibility
// with the sequence number field
func (c *checkpointSeqNumber) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// Primary struct shape
	return nil
}

// fallback

type azPartitionClient interface {
	Close(ctx context.Context) error
	ReceiveEvents(ctx context.Context, maxBatchSize int, options *azeventhubs.ReceiveEventsOptions) ([]*azeventhubs.ReceivedEventData, error)
}

func getPollConfig(config *Config) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func getConsumerGroup(config *Config) string { _ = "STUB: not implemented"; return "" }

// createConsumerClient creates a new Azure Event Hub consumer client.
// If auth is configured, it uses the auth extension to create the client.
// Otherwise, it uses the connection string.
func createConsumerClient(
	config *Config,
	host component.Host,
	consumerGroup string,
	logger *zap.Logger,
) (*azeventhubs.ConsumerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAzeventhubWrapper(h *eventhubHandler, host component.Host) (*hubWrapperAzeventhubImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getStorageCheckpointPersister(storageClient storage.Client) *storageCheckpointPersister[checkpointSeqNumber] {
	_ = "STUB: not implemented"
	return nil
}

type azEventHubWrapper struct {
	*azeventhubs.ConsumerClient
}

func (w azEventHubWrapper) GetEventHubProperties(ctx context.Context, options *azeventhubs.GetEventHubPropertiesOptions) (azeventhubs.EventHubProperties, error) {
	_ = "STUB: not implemented"
	return *new(azeventhubs.EventHubProperties), nil
}

func (w azEventHubWrapper) GetPartitionProperties(ctx context.Context, partitionID string, options *azeventhubs.GetPartitionPropertiesOptions) (azeventhubs.PartitionProperties, error) {
	_ = "STUB: not implemented"
	return *new(azeventhubs.PartitionProperties), nil
}

func (w azEventHubWrapper) NewPartitionClient(partitionID string, options *azeventhubs.PartitionClientOptions) (azPartitionClient, error) {
	_ = "STUB: not implemented"
	return *new(azPartitionClient), nil
}

func (w azEventHubWrapper) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type azEventHub interface {
	GetEventHubProperties(ctx context.Context, options *azeventhubs.GetEventHubPropertiesOptions) (azeventhubs.EventHubProperties, error)
	GetPartitionProperties(ctx context.Context, partitionID string, options *azeventhubs.GetPartitionPropertiesOptions) (azeventhubs.PartitionProperties, error)
	NewPartitionClient(partitionID string, options *azeventhubs.PartitionClientOptions) (azPartitionClient, error)
	Close(ctx context.Context) error
}

type hubWrapperAzeventhubImpl struct {
	hub     azEventHub
	config  *Config
	storage *storageCheckpointPersister[checkpointSeqNumber]
}

func (h *hubWrapperAzeventhubImpl) GetRuntimeInformation(ctx context.Context) (*hubRuntimeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *hubWrapperAzeventhubImpl) Receive(ctx context.Context, partitionID string, handler hubHandler, applyOffset bool, logger *zap.Logger) (listenerHandleWrapper, error) {
	_ = "STUB: not implemented"
	return *new(listenerHandleWrapper), nil
}

func (h *hubWrapperAzeventhubImpl) Close(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *hubWrapperAzeventhubImpl) getStartPos(
	applyOffset bool,
	namespace string,
	eventHubName string,
	consumerGroup string,
	partitionID string,
) azeventhubs.StartPosition {
	_ = "STUB: not implemented"
	return *new(azeventhubs.StartPosition)
}

// Only apply the checkpoint seq number offset if we have one saved

func (h *hubWrapperAzeventhubImpl) namespace() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Return the first part of the namespace
// Ex: example.servicebus.windows.net => example

type partitionListener struct {
	done chan struct{}
	mu   sync.Mutex
	err  error
}

func (p *partitionListener) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (p *partitionListener) Err() error { _ = "STUB: not implemented"; return nil }

func (p *partitionListener) setErr(err error) { _ = "STUB: not implemented"; return }

// createBlobCheckpointStore creates an Azure Blob Storage-backed checkpoint store
// for use with the distributed Processor.
func createBlobCheckpointStore(config *Config, host component.Host, logger *zap.Logger) (azeventhubs.CheckpointStore, error) {
	_ = "STUB: not implemented"
	return *new(azeventhubs.CheckpointStore), nil
}

// createProcessor creates an Azure Event Hub Processor for distributed consumption.
// It returns both the Processor and the ConsumerClient so the caller can close the
// client on shutdown (the Processor does not take ownership of closing it).
func createProcessor(config *Config, host component.Host, logger *zap.Logger) (*azeventhubs.Processor, *azeventhubs.ConsumerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// processorPartitionClient is an interface for the subset of
// azeventhubs.ProcessorPartitionClient methods used by processPartitionEvents.
// This enables testing without requiring a real Azure connection.
type processorPartitionClient interface {
	PartitionID() string
	ReceiveEvents(ctx context.Context, count int, options *azeventhubs.ReceiveEventsOptions) ([]*azeventhubs.ReceivedEventData, error)
	UpdateCheckpoint(ctx context.Context, latestEvent *azeventhubs.ReceivedEventData, options *azeventhubs.UpdateCheckpointOptions) error
	Close(ctx context.Context) error
}

// processPartitionEvents receives and processes events from a single partition
// assigned by the Processor.
func processPartitionEvents(
	ctx context.Context,
	partitionClient processorPartitionClient,
	handler hubHandler,
	config *Config,
	logger *zap.Logger,
) {
	_ = "STUB: not implemented"
	return
}
