// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transport // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/transport"

import (
	"errors"
	"net"
	"sync"

	"go.opentelemetry.io/collector/consumer"
)

var errTCPServerDone = errors.New("server stopped")

type tcpServer struct {
	listener  net.Listener
	wg        sync.WaitGroup
	transport Transport
}

// Ensure that Server is implemented on TCP Server.
var _ Server = (*tcpServer)(nil)

// NewTCPServer creates a transport.Server using TCP as its transport.
func NewTCPServer(transport Transport, address string) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

// ListenAndServe starts the server ready to receive metrics.
func (t *tcpServer) ListenAndServe(nextConsumer consumer.Metrics, reporter Reporter, transferChan chan<- Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// handleTCPConn is helper that parses the buffer and split it line by line to be parsed upstream.
func handleTCPConn(c net.Conn, reporter Reporter, transferChan chan<- Metric) {
	_ = "STUB: not implemented"
	return
}

// Close closes the server.
func (t *tcpServer) Close() error { _ = "STUB: not implemented"; return nil }
