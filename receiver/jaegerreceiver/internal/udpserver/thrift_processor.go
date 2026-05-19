// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package udpserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/udpserver"

import (
	"context"
	"sync"

	"github.com/apache/thrift/lib/go/thrift"
	"go.uber.org/zap"
)

// ThriftProcessor is a server that processes spans using a TBuffered Server
type ThriftProcessor struct {
	server        *UDPServer
	handler       AgentProcessor
	protocolPool  *sync.Pool
	numProcessors int
	processing    sync.WaitGroup
	logger        *zap.Logger
}

// AgentProcessor handler used by the processor to process thrift and call the reporter
// with the deserialized struct. This interface is implemented directly by Thrift generated
// code, e.g. jaegerThrift.NewAgentProcessor(handler), where handler implements the Agent
// Thrift service interface, which is invoked with the deserialized struct.
type AgentProcessor interface {
	Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException)
}

// NewThriftProcessor creates a TBufferedServer backed ThriftProcessor
func NewThriftProcessor(
	server *UDPServer,
	numProcessors int,
	factory thrift.TProtocolFactory,
	handler AgentProcessor,
	logger *zap.Logger,
) (*ThriftProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Serve starts serving traffic
func (s *ThriftProcessor) Serve() {
	_ = "STUB: not implemented"

	// IsServing indicates whether the server is currently serving traffic
	return
}

func (s *ThriftProcessor) IsServing() bool { _ = "STUB: not implemented"; return false }

// Stop stops the serving of traffic and waits until the queue is
// emptied by the readers
func (s *ThriftProcessor) Stop() { _ = "STUB: not implemented"; return }

// processBuffer reads data off the channel and puts it into a custom transport for
// the processor to process
func (s *ThriftProcessor) processBuffer() { _ = "STUB: not implemented"; return }

// writes to memory transport don't fail

// NB: oddly, thrift-gen/agent/agent.go:L156 does this: `return true, thrift.WrapTException(err2)`
// So we check for both OK and error.

// acknowledge receipt and release the buffer
