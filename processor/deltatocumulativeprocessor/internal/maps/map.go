// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package maps // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/maps"

import (
	"sync/atomic"

	"github.com/puzpuzpuz/xsync/v4"
)

func Limit(limit int64) Context { _ = "STUB: not implemented"; return *new(Context) }

func New[K comparable, V any](ctx Context) *Parallel[K, V] { _ = "STUB: not implemented"; return nil }

// lowercase alias for unexported embedding
type ctx = Context

// Parallel is a lock-free map-like structure. It can be safely used by multiple
// routines concurrently.
//
// Due to the lock-free nature, typical get, put, delete operations are not
// available. Instead, [Parallel.LoadOrStore] returns an existing value or
// inserts a new one if missing. As such, values themselves should be mutable by
// being reference types (pointers or pmetric.* types).
//
// Parallel enforces the [Context] size limit.
type Parallel[K comparable, V any] struct {
	ctx
	elems *xsync.Map[K, V]
}

// Context holds size information about one or more maps.
// Can be shared across maps for a common limit.
type Context struct {
	limit int64
	guard *atomic.Int64
	total *atomic.Int64
}

func (ctx Context) String() string { _ = "STUB: not implemented"; return "" }

// LoadOrStore loads existing values from the map or creates missing ones initialized to <def>.
//
// Return Value:
//   - <value>, true: m[k] already existed and was loaded
//   - <def>, false: m[k] was created and initialized to <def>
//   - <zero>, false: m[k] did not exist but was not created due to size limit
func (m *Parallel[K, V]) LoadOrStore(k K, def V) (_ V, loaded bool) {
	_ = "STUB: not implemented"
	// multiple routines may attempt to LoadOrStore the same value at once. as
	// such, we cannot use data-dependent instructions such as if(not exist)
	// {...}, because the <not exist> may have changed right after we checked
	// it.
	return *new(V), false
}

// as long as there appears to be actual space, try to store

// multiple routines may do this. to enforce the limit, try to claim a
// "slot" below the limit

// slot we got is above the limit. either the map is now full (loop
// will exit) or routines that won't actually store hold slots, in
// which case we will try again.

// we got a valid slot. others may too. as such, we try to store, but
// may end up loading instead if another routine stored just before us.

// another routine stored, but we got a value. give up slot

// we stored. increase the total size

// we didn't store, because we hit the limit. attempt another load, just in
// case another routine stored by now.

// LoadAndDelete deletes m[k], returning the value it had if it existed
func (m *Parallel[K, V]) LoadAndDelete(k K) (_ V, loaded bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// m[k] did exist. decrease size and open up a slot

func (ctx Context) Size() int64 { _ = "STUB: not implemented"; return 0 }

// Exceeded reports whether a [Limited.LoadOrStore] failed due to the limit being exceeded.
func Exceeded[T comparable](v T, loaded bool) bool { _ = "STUB: not implemented"; return false }
