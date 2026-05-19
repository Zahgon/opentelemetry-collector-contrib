// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package staleness // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/staleness"

import (
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
)

// PriorityQueue represents a way to store entries sorted by their priority.
// Pop() will return the oldest entry of the set.
type PriorityQueue interface {
	// Update will add or update an entry, and reshuffle the queue internally as needed to keep it sorted
	Update(id identity.Stream, newPrio time.Time)
	// Peek will return the entry at the HEAD of the queue *without* removing it from the queue
	Peek() (identity.Stream, time.Time)
	// Pop will remove the entry at the HEAD of the queue and return it
	Pop() (identity.Stream, time.Time)
	// Len will return the number of entries in the queue
	Len() int
}

// heapQueue implements heap.Interface.
// We use it as the inner implementation of a heap-based sorted queue
type heapQueue []*queueItem

type queueItem struct {
	key   identity.Stream
	prio  time.Time
	index int
}

func (pq heapQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq heapQueue) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// We want Pop to give us the lowest priority
	return false
}

func (pq heapQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *heapQueue) Push(x any) { _ = "STUB: not implemented"; return }

func (pq *heapQueue) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// avoid memory leak
// for safety

type heapPriorityQueue struct {
	inner      heapQueue
	itemLookup map[identity.Stream]*queueItem
}

func NewPriorityQueue() PriorityQueue { _ = "STUB: not implemented"; return *new(PriorityQueue) }

func (pq *heapPriorityQueue) Update(id identity.Stream, newPrio time.Time) {
	_ = "STUB: not implemented"
	// Check if the entry already exists in the queue
	return
}

// If so, we can update it in place

func (pq *heapPriorityQueue) Peek() (identity.Stream, time.Time) {
	_ = "STUB: not implemented"
	return *new(identity.Stream), *new(time.Time)
}

func (pq *heapPriorityQueue) Pop() (identity.Stream, time.Time) {
	_ = "STUB: not implemented"
	return *new(identity.Stream), *new(time.Time)
}

func (pq *heapPriorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }
