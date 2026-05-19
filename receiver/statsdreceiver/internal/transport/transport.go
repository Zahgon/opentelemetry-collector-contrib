// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transport // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/transport"

import "errors"

// Transport is a set of constants of the transport supported by this receiver.
type Transport string

var (
	ErrUnsupportedTransport       = errors.New("unsupported transport")
	ErrUnsupportedPacketTransport = errors.New("unsupported Packet transport")
	ErrUnsupportedStreamTransport = errors.New("unsupported Stream transport")
)

const (
	UDP  Transport = "udp"
	UDP4 Transport = "udp4"
	UDP6 Transport = "udp6"
	TCP  Transport = "tcp"
	TCP4 Transport = "tcp4"
	TCP6 Transport = "tcp6"
	UDS  Transport = "unixgram"
)

// NewTransport creates a Transport based on the transport string or returns an empty Transport.
func NewTransport(ts string) Transport { _ = "STUB: not implemented"; return *new(Transport) }

// String casts the transport to a String if the Transport is supported. Return an empty Transport overwise.
func (trans Transport) String() string { _ = "STUB: not implemented"; return "" }

// IsPacketTransport returns true if the transport is packet based.
func (trans Transport) IsPacketTransport() bool { _ = "STUB: not implemented"; return false }

// IsStreamTransport returns true if the transport is stream based.
func (trans Transport) IsStreamTransport() bool { _ = "STUB: not implemented"; return false }
