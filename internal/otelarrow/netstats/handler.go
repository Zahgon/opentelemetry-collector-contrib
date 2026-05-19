// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package netstats // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/netstats"

import (
	"context"

	"google.golang.org/grpc/stats"
)

type netstatsContext struct{} // value: string

type statsHandler struct {
	rep *NetworkReporter
}

var _ stats.Handler = statsHandler{}

func (rep *NetworkReporter) Handler() stats.Handler {
	_ = "STUB: not implemented"
	return *new(stats.Handler)
}

// TagRPC implements grpc/stats.Handler
func (statsHandler) TagRPC(ctx context.Context, s *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// trustUncompressed is a super hacky way of knowing when the
// uncompressed size is realistic.  nothing else would work -- the
// same handler is used by both arrow and non-arrow, and the
// `*stats.Begin` which indicates streaming vs. not streaming does not
// appear in TagRPC() where we could store it in context.  this
// approach is considered a better alternative than others, however
// ugly.  when non-arrow RPCs are sent, the instrumentation works
// correctly, this avoids special instrumentation outside of the Arrow
// components.
func trustUncompressed(method string) bool { _ = "STUB: not implemented"; return false }

func (h statsHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

// Note we have some info about header WireLength,
// but intentionally not counting.

// TagConn implements grpc/stats.Handler
func (statsHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"

	// HandleConn implements grpc/stats.Handler
	return *new(context.Context)
}

func (statsHandler) HandleConn(_ context.Context, _ stats.ConnStats) {
	_ = "STUB: not implemented"
	// Note: ConnBegin and ConnEnd
	return
}
