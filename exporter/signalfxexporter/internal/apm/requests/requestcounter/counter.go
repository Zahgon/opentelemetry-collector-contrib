// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Originally copied from https://github.com/signalfx/signalfx-agent/blob/fbc24b0fdd3884bd0bbfbd69fe3c83f49d4c0b77/pkg/apm/requests/requestcounter/counter.go

package requestcounter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/requests/requestcounter"

import (
	"context"
)

type key int

const (
	getRequestCountKey       key = 1
	incrementRequestCountKey key = 2
	resetRequestCountKey     key = 3
)

type (
	getRequestCount       func() uint32
	incrementRequestCount func()
	resetRequestCount     func()
)

// checks if a counter already exists on the context
func counterExists(ctx context.Context) (exists bool) { _ = "STUB: not implemented"; return false }

// ContextWithRequestCounter adds a counter to the context if one does not already exist
func ContextWithRequestCounter(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *

	// don't create a new context with counter if a counter already exists on the counter
	new(context.Context)
}

// ResetRequestCount resets the request counter on the provided context if the context has one
func ResetRequestCount(ctx context.Context) { _ = "STUB: not implemented"; return }

// IncrementRequestCount increments the request counter on the provided context if the context has one
func IncrementRequestCount(ctx context.Context) { _ = "STUB: not implemented"; return }

// GetRequestCount retrieves the current request count on the provided context.  It returns 0 if the context does not
// have a request counter
func GetRequestCount(ctx context.Context) (count uint32) { _ = "STUB: not implemented"; return 0 }
