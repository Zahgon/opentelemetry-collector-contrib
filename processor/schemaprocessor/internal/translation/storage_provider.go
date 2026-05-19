// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/translation"

import (
	"context"

	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"
)

// StorageProvider wraps another Provider and adds persistent storage as a
// read-through cache. On Retrieve, it checks storage first. On a miss, it
// delegates to the wrapped provider and persists the result for future use.
type StorageProvider struct {
	wrapped Provider
	client  storage.Client
	log     *zap.Logger
}

// NewStorageProvider creates a provider that checks persistent storage before
// delegating to the wrapped provider. Fetched schemas are persisted to storage
// for use across collector restarts.
func NewStorageProvider(wrapped Provider, client storage.Client, log *zap.Logger) *StorageProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *StorageProvider) Retrieve(ctx context.Context, schemaURL string) (string, error) {
	_ = "STUB: not implemented"
	// Check persistent storage first.
	return "", nil
}

// Fall through to the wrapped provider (typically HTTP).

// Persist to storage for next time. Best-effort — don't fail on storage errors.
