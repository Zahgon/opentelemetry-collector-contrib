// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transport // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/transport"

type udpServer struct {
	packetServer
}

// Ensure that Server is implemented on UDP Server.
var _ Server = (*udpServer)(nil)

// NewUDPServer creates a transport.Server using UDP as its transport.
func NewUDPServer(transport Transport, address string) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

// Close closes the server.
func (u *udpServer) Close() error { _ = "STUB: not implemented"; return nil }
