// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transport // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/transport"

import (
	"net"

	"go.opentelemetry.io/collector/consumer"
)

type packetServer struct {
	packetConn net.PacketConn
	transport  Transport
}

// ListenAndServe starts the server ready to receive metrics.
func (u *packetServer) ListenAndServe(
	nextConsumer consumer.Metrics,
	reporter Reporter,
	transferChan chan<- Metric,
) error {
	_ = "STUB: not implemented"
	return nil
}

// max size for udp packet body (assuming ipv6)

// handlePacket is helper that parses the buffer and split it line by line to be parsed upstream.
func (*packetServer) handlePacket(
	numBytes int,
	data []byte,
	addr net.Addr,
	transferChan chan<- Metric,
) {
	_ = "STUB: not implemented"
	return
}

type udsAddr struct {
	network string
	address string
}

func (u *udsAddr) Network() string { _ = "STUB: not implemented"; return "" }

func (u *udsAddr) String() string { _ = "STUB: not implemented"; return "" }
