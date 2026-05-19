// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transport // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/internal/transport"

import (
	"net"
	"sync"
	"time"

	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"
)

type tcpServer struct {
	ln          net.Listener
	wg          sync.WaitGroup
	idleTimeout time.Duration
	reporter    Reporter
}

var _ Server = (*tcpServer)(nil)

// NewTCPServer creates a transport.Server using TCP as its transport.
func NewTCPServer(
	addr string,
	idleTimeout time.Duration,
) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

func (t *tcpServer) ListenAndServe(
	parser protocol.Parser,
	nextConsumer consumer.Metrics,
	reporter Reporter,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Close any lingering connection

func (t *tcpServer) Close() error { _ = "STUB: not implemented"; return nil }

func (t *tcpServer) handleConnection(
	p protocol.Parser,
	nextConsumer consumer.Metrics,
	conn net.Conn,
) {
	_ = "STUB: not implemented"
	return
}

// reader.ReadBytes call below will block until either:
//
// * a '\n' char is read
// * the connection is closed (either by client or server)
// * an idle timeout happens (see call to conn.SetDeadline above)
//
// Notice that it is possible for the function to return with error at
// the same time that it returns data (typically the error is io.EOF in
// this case).

// The protocol doesn't account for returning errors.
// Since this is a TCP connection it seems reasonable to close the
// connection as a way to report "error" back to client and minimize
// the effect of a client constantly submitting bad data.

// We want to end on timeout so idle connections are purged.
