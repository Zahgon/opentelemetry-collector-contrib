// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package priorityqueue // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/priorityqueue"

import (
	"cmp"
)

type PriorityValueType cmp.Ordered

type QueueItem[V any, P PriorityValueType] struct {
	Value    V
	Priority P
	Index    int
}

type PriorityQueue[V any, P PriorityValueType] []*QueueItem[V, P]

func (pq PriorityQueue[V, P]) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq PriorityQueue[V, P]) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pq PriorityQueue[V, P]) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *PriorityQueue[V, P]) Push(x any) { _ = "STUB: not implemented"; return }

func (pq *PriorityQueue[V, P]) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// avoid memory leak
// for safety
