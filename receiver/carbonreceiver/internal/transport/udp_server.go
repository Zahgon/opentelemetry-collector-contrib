// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transport // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/internal/transport"

import (
	"net"
	"sync"

	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"
)

type udpServer struct {
	wg         sync.WaitGroup
	packetConn net.PacketConn
	reporter   Reporter
}

var _ Server = (*udpServer)(nil)

// NewUDPServer creates a transport.Server using UDP as its transport.
func NewUDPServer(addr string) (Server, error) { _ = "STUB: not implemented"; return *new(Server), nil }

func (u *udpServer) ListenAndServe(
	parser protocol.Parser,
	nextConsumer consumer.Metrics,
	reporter Reporter,
) error {
	_ = "STUB: not implemented"
	return nil
}

// max size for udp packet body (assuming ipv6)

func (u *udpServer) Close() error { _ = "STUB: not implemented"; return nil }

func (u *udpServer) handlePacket(
	p protocol.Parser,
	nextConsumer consumer.Metrics,
	data []byte,
) {
	_ = "STUB: not implemented"
	return
}

// Completed without errors.
