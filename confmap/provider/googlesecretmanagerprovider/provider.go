// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package googlesecretmanagerprovider // import "github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/googlesecretmanagerprovider"

import (
	"context"
	"errors"

	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	gax "github.com/googleapis/gax-go/v2"
	"go.opentelemetry.io/collector/confmap"
)

type secretsManagerClient interface {
	AccessSecretVersion(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error)
	Close() error
}

const (
	schemeName = "googlesecretmanager"
)

var (
	ErrURINotSupported     = errors.New("uri is not supported by Google Secret Manager Provider")
	ErrAccessSecretVersion = errors.New("failed to access secret version")
)

type provider struct {
	client secretsManagerClient
}

func NewFactory() confmap.ProviderFactory {
	_ = "STUB: not implemented"
	return *new(confmap.ProviderFactory)
}

func newProvider(confmap.ProviderSettings) confmap.Provider {
	_ = "STUB: not implemented"
	return *new(confmap.Provider)
}

func (p *provider) Retrieve(ctx context.Context, uri string, _ confmap.WatcherFunc) (*confmap.Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*provider) Scheme() string { _ = "STUB: not implemented"; return "" }

func (p *provider) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
