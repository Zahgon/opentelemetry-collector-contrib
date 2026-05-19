// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"

import (
	"context"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type hecWorker interface {
	send(context.Context, buffer, map[string]string) error
}

type defaultHecWorker struct {
	url     *url.URL
	client  *http.Client
	headers map[string]string
	logger  *zap.Logger
}

func (hec *defaultHecWorker) send(ctx context.Context, buf buffer, headers map[string]string) error {
	_ = "STUB: not implemented"
	// We copy the bytes to a new buffer to avoid corruption. This is a workaround to avoid hitting https://github.com/golang/go/issues/51907.
	return nil
}

// Set the headers configured for the client

// Set extra headers passed by the caller

// Drain the response body to avoid leaking connections.

var _ hecWorker = &defaultHecWorker{}
