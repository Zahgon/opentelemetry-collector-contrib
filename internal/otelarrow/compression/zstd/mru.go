// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zstd // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/compression/zstd"

import (
	"sync"
	"time"
)

// mru is a freelist whose two main benefits compared to sync.Pool are:
//
//   - It doesn't perform any per-CPU caching; it has only a single
//     cache. The cache is modeled as a stack, meaning that the most
//     recently used item is always the next to be used. (Hence the name
//     MRU.)
//
//   - It isn't cleared when GC runs. Instead, items that haven't been used
//     in a long time (1min) are released.
//
// An MRU freelist is most useful when the objects being freelisted are
// sufficiently valuable, or expensive to create, that they are worth keeping
// across GC passes. The drawbacks are that MRU isn't as performant under
// heavy concurrent access as sync.Pool, and that its sizing logic (1min TTL)
// is less sophisticated than sync.Pool's.
//
// A zero-initialized MRU is safe to use. Threadsafe.
type mru[T generational] struct {
	mu       sync.Mutex
	reset    Gen
	freelist []T
	putTimes []time.Time // putTimes[i] is when freelist[i] was Put()
	zero     T
}

// Gen is the reset time.
type Gen time.Time

type generational interface {
	// generation uses monotonic time
	generation() Gen
}

// TTL is modified in testing.
var TTL = time.Minute

// Get returns an object from the freelist. If the list is empty, the return
// value is the zero value of T.
func (mru *mru[T]) Get() (T, Gen) { _ = "STUB: not implemented"; return *new(T), *new(Gen) }

// Allow GC to occur.

func before(a, b Gen) bool { _ = "STUB: not implemented"; return false }

func (mru *mru[T]) Put(item T) { _ = "STUB: not implemented"; return }

// Evict any objects that haven't been touched recently.

// Shift values by one index in the slice, to preserve capacity.

// Allow GC to occur.

func (mru *mru[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (mru *mru[T]) Reset() Gen { _ = "STUB: not implemented"; return *new(Gen) }
