// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package storagetest // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/storagetest"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

type StorageHost struct {
	component.Host
	extensions map[component.ID]component.Component
}

func NewStorageHost() *StorageHost { _ = "STUB: not implemented"; return nil }

func (h *StorageHost) WithExtension(id component.ID, ext extension.Extension) *StorageHost {
	_ = "STUB: not implemented"
	return nil
}

func (h *StorageHost) WithInMemoryStorageExtension(name string) *StorageHost {
	_ = "STUB: not implemented"
	return nil
}

func (h *StorageHost) WithFileBackedStorageExtension(name, storageDir string) *StorageHost {
	_ = "STUB: not implemented"
	return nil
}

func (h *StorageHost) WithNonStorageExtension(name string) *StorageHost {
	_ = "STUB: not implemented"
	return nil
}

func (h *StorageHost) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}
