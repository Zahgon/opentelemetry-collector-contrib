// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

type blobEventHandler struct {
	blobClient          blobClient
	logsDataConsumer    logsDataConsumer
	tracesDataConsumer  tracesDataConsumer
	logsContainerName   string
	tracesContainerName string
	logger              *zap.Logger
	wg                  sync.WaitGroup
	cancelFunc          context.CancelFunc
	pollInterval        time.Duration
}

var _ eventHandler = (*blobEventHandler)(nil)

func (p *blobEventHandler) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *blobEventHandler) pollBlobs(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *blobEventHandler) processContainers(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (p *blobEventHandler) processContainer(ctx context.Context, containerName string, consume func(context.Context, []byte) error) {
	_ = "STUB: not implemented"
	return
}

func (p *blobEventHandler) processBlob(ctx context.Context, containerName, blobName string, consume func(context.Context, []byte) error) {
	_ = "STUB: not implemented"
	return
}

func (p *blobEventHandler) close(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *blobEventHandler) setLogsDataConsumer(logsDataConsumer logsDataConsumer) {
	_ = "STUB: not implemented"
	return
}

func (p *blobEventHandler) setTracesDataConsumer(tracesDataConsumer tracesDataConsumer) {
	_ = "STUB: not implemented"
	return
}

func newBlobEventHandler(logsContainerName, tracesContainerName string, blobClient blobClient, logger *zap.Logger) *blobEventHandler {
	_ = "STUB: not implemented"
	return nil
}
