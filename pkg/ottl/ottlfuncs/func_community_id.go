// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context" //gosec:disable G505 -- SHA1 is intentionally used for generating unique identifiers
	"net"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

var communityIDProtocols = map[string]uint8{
	"ICMP":  1,
	"TCP":   6,
	"UDP":   17,
	"RSVP":  46,
	"ICMP6": 58,
	"SCTP":  132,
}

type CommunityIDArguments[K any] struct {
	SourceIP        ottl.StringGetter[K]
	SourcePort      ottl.IntGetter[K]
	DestinationIP   ottl.StringGetter[K]
	DestinationPort ottl.IntGetter[K]
	Protocol        ottl.Optional[ottl.StringGetter[K]]
	Seed            ottl.Optional[ottl.IntGetter[K]]
}

func NewCommunityIDFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createCommunityIDFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type communityIDHash struct {
	srcIPBytes []byte
	dstIPBytes []byte
	srcPort    uint16
	dstPort    uint16
	protocol   uint8
	seed       uint16
}

func (h *communityIDHash) normalize() { _ = "STUB: not implemented"; return }

func (h *communityIDHash) compute() string { _ = "STUB: not implemented"; return "" }

// Add seed (2 bytes, network order)

// Add source, destination IPs and 1-byte protocol

// Add source and destination ports (2 bytes each, network order)

// Generate the SHA1 hash
//gosec:disable G401 -- we are not using SHA1 for security, but for generating unique identifier, conflicts will be solved with the seed

// Add version prefix (1) and return

func communityID[K any](
	sourceIP ottl.StringGetter[K],
	sourcePort ottl.IntGetter[K],
	destinationIP ottl.StringGetter[K],
	destinationPort ottl.IntGetter[K],
	protocol ottl.Optional[ottl.StringGetter[K]],
	seed ottl.Optional[ottl.IntGetter[K]],
) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

func extractIPAndPort[K any](
	ctx context.Context,
	tCtx K,
	ipGetter ottl.StringGetter[K],
	portGetter ottl.IntGetter[K],
	endpointName string,
) (net.IP, int64, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), 0, nil
}

func makeCommunityIDHash[K any](
	ctx context.Context,
	tCtx K,
	sourceIP ottl.StringGetter[K],
	sourcePort ottl.IntGetter[K],
	destinationIP ottl.StringGetter[K],
	destinationPort ottl.IntGetter[K],
	protocol ottl.Optional[ottl.StringGetter[K]],
	seed ottl.Optional[ottl.IntGetter[K]],
) (*communityIDHash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaults to TCP

// Get seed value (default: 0) if applied
