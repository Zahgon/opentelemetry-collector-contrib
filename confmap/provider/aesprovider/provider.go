// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package aesprovider // import "github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/aesprovider"

import (
	"context"

	"go.opentelemetry.io/collector/confmap"
	"go.uber.org/zap"
)

const (
	schemaName = "aes"
	// This environment variable holds a base64-encoded AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	keyEnvVar = "OTEL_AES_CREDENTIAL_PROVIDER"
)

type provider struct {
	logger *zap.Logger
	key    []byte
}

// NewFactory creates a new provider factory
func NewFactory() confmap.ProviderFactory {
	_ = "STUB: not implemented"
	return *new(confmap.ProviderFactory)
}

func (*provider) Scheme() string { _ = "STUB: not implemented"; return "" }

func (*provider) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *provider) Retrieve(_ context.Context, uri string, _ confmap.WatcherFunc) (*confmap.Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// base64 decode env var

// Remove schemaName

func (p *provider) decrypt(cipherText string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
