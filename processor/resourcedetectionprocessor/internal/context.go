// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"

import (
	"context"
	"net/http"
)

type contextKey int

const clientContextKey contextKey = iota

// ContextWithClient returns a new context.Context with the provided *http.Client stored as a value.
func ContextWithClient(ctx context.Context, client *http.Client) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ClientFromContext attempts to extract an *http.Client from the provided context.Context.
func ClientFromContext(ctx context.Context) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
