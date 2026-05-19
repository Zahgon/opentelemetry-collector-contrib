// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package storagetest // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/storagetest"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/xextension/storage"
)

var testStorageType component.Type = component.MustNewType("test_storage")

// TestStorage is an in memory storage extension designed for testing
type TestStorage struct {
	component.StartFunc
	component.ShutdownFunc
	ID         component.ID
	storageDir string
}

// Ensure this storage extension implements the appropriate interface
var _ storage.Extension = (*TestStorage)(nil)

func NewStorageID(name string) component.ID { _ = "STUB: not implemented"; return *new(component.ID) }

// NewInMemoryStorageExtension creates a TestStorage extension
func NewInMemoryStorageExtension(name string) *TestStorage { _ = "STUB: not implemented"; return nil }

// NewFileBackedStorageExtension creates a TestStorage extension
func NewFileBackedStorageExtension(name, storageDir string) *TestStorage {
	_ = "STUB: not implemented"
	return nil
}

// GetClient returns a storage client for an individual component
func (s *TestStorage) GetClient(ctx context.Context, kind component.Kind, ent component.ID, name string) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

var nonStorageType component.Type = component.MustNewType("non_storage")

// NonStorage is useful for testing expected behaviors that involve
// non-storage extensions
type NonStorage struct {
	component.StartFunc
	component.ShutdownFunc
	ID component.ID
}

// Ensure this extension implements the appropriate interface
var _ extension.Extension = (*NonStorage)(nil)

func NewNonStorageID(name string) component.ID {
	_ = "STUB: not implemented"
	return *new(component.ID)
}

// NewNonStorageExtension creates a NonStorage extension
func NewNonStorageExtension(name string) *NonStorage { _ = "STUB: not implemented"; return nil }
