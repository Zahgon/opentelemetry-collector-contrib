// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"

import (
	"sync"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/types"
)

// ObjStore implements the cache.Store interface:
// https://github.com/kubernetes/client-go/blob/release-1.20/tools/cache/store.go#L26-L71
// It is used by cache.Reflector to keep the updated information about resources
// https://github.com/kubernetes/client-go/blob/release-1.20/tools/cache/reflector.go#L48
type ObjStore struct {
	mu sync.RWMutex

	refreshed bool
	objs      map[types.UID]any

	transformFunc func(any) (any, error)
	logger        *zap.Logger
}

func NewObjStore(transformFunc func(any) (any, error), logger *zap.Logger) *ObjStore {
	_ = "STUB: not implemented"
	return nil
}

// GetResetRefreshStatus tracks whether the underlying data store is refreshed or not.
// Calling this func itself will reset the state to false.
func (s *ObjStore) GetResetRefreshStatus() bool { _ = "STUB: not implemented"; return false }

// Add implements the Add method of the store interface.
// Add adds an entry to the ObjStore.
func (s *ObjStore) Add(obj any) error { _ = "STUB: not implemented"; return nil }

// Update implements the Update method of the store interface.
// Update updates the existing entry in the ObjStore.
func (s *ObjStore) Update(obj any) error {
	_ = "STUB: not implemented"

	// Delete implements the Delete method of the store interface.
	// Delete deletes an existing entry in the ObjStore.
	return nil
}

func (s *ObjStore) Delete(obj any) error { _ = "STUB: not implemented"; return nil }

// List implements the List method of the store interface.
// List lists all the objects in the ObjStore
func (s *ObjStore) List() []any { _ = "STUB: not implemented"; return nil }

// ListKeys implements the ListKeys method of the store interface.
// ListKeys lists the keys for all objects in the ObjStore
func (s *ObjStore) ListKeys() []string { _ = "STUB: not implemented"; return nil }

// Get implements the Get method of the store interface.
func (*ObjStore) Get(any) (item any, exists bool, err error) {
	_ = "STUB: not implemented"
	return *

	// GetByKey implements the GetByKey method of the store interface.
	new(any), false, nil
}

func (*ObjStore) GetByKey(string) (item any, exists bool, err error) {
	_ = "STUB: not implemented"
	return *

	// Replace implements the Replace method of the store interface.
	// Replace will delete the contents of the store, using instead the given list.
	new(any), false, nil
}

func (s *ObjStore) Replace(list []any, _ string) error { _ = "STUB: not implemented"; return nil }

// Resync implements the Resync method of the store interface.
func (*ObjStore) Resync() error { _ = "STUB: not implemented"; return nil }
