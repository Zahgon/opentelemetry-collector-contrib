// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/translation"

import (
	"context"
	"net/http"
)

// Provider allows for collector extensions to be used to look up schemaURLs
type Provider interface {
	// Retrieve whill check the underlying provider to see if content exists
	// for the provided schemaURL, in the even that it doesn't an error is returned.
	Retrieve(ctx context.Context, schemaURL string) (string, error)
}

type httpProvider struct {
	client *http.Client
}

var _ Provider = (*httpProvider)(nil)

// NewHTTPProvider creates a new HTTP-based Provider.
func NewHTTPProvider(client *http.Client) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (hp *httpProvider) Retrieve(ctx context.Context, schemaURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
