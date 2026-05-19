// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package store // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector/internal/store"

import (
	"container/list"
	"errors"
	"sync"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

var ErrTooManyItems = errors.New("too many items")

type Callback func(e *Edge)

type Key struct {
	tid pcommon.TraceID
	sid pcommon.SpanID
}

func (k *Key) SpanIDIsEmpty() bool { _ = "STUB: not implemented"; return false }

func NewKey(tid pcommon.TraceID, sid pcommon.SpanID) Key {
	_ = "STUB: not implemented"
	return *new(Key)
}

type Store struct {
	l   *list.List
	mtx sync.Mutex
	m   map[Key]*list.Element

	onComplete Callback
	onExpire   Callback

	ttl      time.Duration
	maxItems int
}

// NewStore creates a Store to build service graphs. The store caches edges, each representing a
// request between two services. Once an edge is complete its metrics can be collected. Edges that
// have not found their pair are deleted after ttl time.
func NewStore(ttl time.Duration, maxItems int, onComplete, onExpire Callback) *Store {
	_ = "STUB: not implemented"
	return nil
}

// Len is only used for testing.
func (s *Store) Len() int {
	_ = "STUB: not implemented"

	// UpsertEdge fetches an Edge from the store and updates it using the given callback. If the Edge
	// doesn't exist yet, it creates a new one with the default TTL.
	// If the Edge is complete after applying the callback, it's completed and removed.
	return 0
}

func (s *Store) UpsertEdge(key Key, update Callback) (isNew bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check we can add new edges

// TODO: try to evict expired items

// Expire evicts all expired items in the store.
func (s *Store) Expire() { _ = "STUB: not implemented"; return }

// Iterates until no more items can be evicted

// tryEvictHead checks if the oldest item (head of list) can be evicted and will delete it if so.
// Returns true if the head was evicted.
//
// Must be called holding lock.
func (s *Store) tryEvictHead() bool { _ = "STUB: not implemented"; return false }

// list is empty
