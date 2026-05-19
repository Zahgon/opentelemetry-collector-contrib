// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package drainprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/drainprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/xextension/storage"
)

const storageKey = "drain_tree"

// getStorageClient resolves a storage.Client for the processor.
func getStorageClient(ctx context.Context, host component.Host, storageID *component.ID, componentID component.ID) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

// loadSnapshot attempts to restore tree state from storage. Returns true if a
// valid snapshot was loaded, false otherwise (caller should seed the tree).
func (p *drainProcessor) loadSnapshot(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// startPeriodicSave launches a background goroutine that saves the tree
// snapshot at the configured interval.
func (p *drainProcessor) startPeriodicSave(ctx context.Context) { _ = "STUB: not implemented"; return }

// saveSnapshot serializes the tree and writes it to storage. The write is
// skipped when the snapshot hash matches the last saved hash (no changes).
func (p *drainProcessor) saveSnapshot(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
