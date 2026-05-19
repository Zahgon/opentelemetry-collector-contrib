// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fluentforwardreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"

import (
	"context"
	"net"
	"sync"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver/internal/metadata"
)

// The initial size of the read buffer. Messages can come in that are bigger
// than this, but this serves as a starting point.
const readBufferSize = 10 * 1024

type server struct {
	outCh            chan<- event
	logger           *zap.Logger
	telemetryBuilder *metadata.TelemetryBuilder
	conns            map[net.Conn]struct{}
	mu               sync.Mutex
}

func newServer(outCh chan<- event, logger *zap.Logger, telemetryBuilder *metadata.TelemetryBuilder) *server {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) Start(ctx context.Context, listener net.Listener) {
	_ = "STUB: not implemented"
	return
}

func (s *server) handleConnections(ctx context.Context, listener net.Listener) {
	_ = "STUB: not implemented"
	return
}

// If there is an error and the receiver isn't shutdown, we need to
// keep trying to accept connections if at all possible. Put in a sleep
// to prevent hot loops in case the error persists.

func (s *server) handleConn(ctx context.Context, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

// We must acknowledge the 'chunk' option if given. We could do this in
// another goroutine if it is too much of a bottleneck to reading
// messages -- this is the only thing that sends data back to the
// client.

// determineNextEventMode inspects the next bit of data from the given peeker
// reader to determine which type of event mode it is.  According to the
// forward protocol spec: "Server MUST detect the carrier mode by inspecting
// the second element of the array."  It is assumed that peeker is aligned at
// the start of a new event, otherwise the result is undefined and will
// probably error.
func determineNextEventMode(peeker peeker) (eventMode, error) {
	_ = "STUB: not implemented"
	return *new(eventMode), nil
}

// The first byte is the array header, which will always be 1 byte since no
// message modes have more than 4 entries. So skip to the second byte which
// is the tag string header.

// We already read the first type for the type

// Skip past the first byte (array header) and the entire tag and then get
// one byte into the second field -- that is enough to know its type.

func (s *server) addConn(c net.Conn) { _ = "STUB: not implemented"; return }

func (s *server) removeConn(c net.Conn) { _ = "STUB: not implemented"; return }

func (s *server) closeAllConns() { _ = "STUB: not implemented"; return }

// Ignore errors
