// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ttlmap // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/ttlmap"

import (
	"sync"
)

// TTLMap is a map that evicts entries after the configured ttl has elapsed.
type TTLMap struct {
	md            *ttlMapData
	done          chan struct{}
	sweepInterval int64
}

// New creates a TTLMap. The sweepIntervalSeconds arg indicates how often
// entries are checked for expiration. The maxAgeSeconds arg indicates how long
// entries can persist before getting evicted. Call Start() on the returned
// TTLMap to begin periodic sweeps which check for expiration and evict entries
// as needed.
// done is the channel that will be used to signal to the timer to stop its work.
func New(sweepIntervalSeconds, maxAgeSeconds int64, done chan struct{}) *TTLMap {
	_ = "STUB: not implemented"
	return nil
}

// Start starts periodic sweeps for expired entries in the underlying map.
func (m *TTLMap) Start() { _ = "STUB: not implemented"; return }

// Put adds the passed-in key and value to the underlying map. The current time
// is attached to the entry for periodic expiration checking and eviction when
// necessary.
func (m *TTLMap) Put(k string, v any) { _ = "STUB: not implemented"; return }

// Get returns the object in the underlying map at the given key. If there is no
// value at that key, Get returns nil.
func (m *TTLMap) Get(k string) any { _ = "STUB: not implemented"; return *new(any) }

func (m *TTLMap) Shutdown() { _ = "STUB: not implemented"; return }

type entry struct {
	v          any
	createTime int64
}

type ttlMapData struct {
	m      map[string]entry
	maxAge int64
	mux    sync.Mutex
}

func newTTLMapData(maxAgeSeconds int64) *ttlMapData { _ = "STUB: not implemented"; return nil }

func (d *ttlMapData) put(k string, v any, currTime int64) { _ = "STUB: not implemented"; return }

func (d *ttlMapData) get(k string) any { _ = "STUB: not implemented"; return *new(any) }

func (d *ttlMapData) sweep(currTime int64) { _ = "STUB: not implemented"; return }
