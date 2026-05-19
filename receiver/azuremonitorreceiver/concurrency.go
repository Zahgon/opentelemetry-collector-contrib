// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver"

import (
	"sync"

	cmap "github.com/orcaman/concurrent-map/v2"
)

type concurrentMetricsBuilderMap[V any] interface {
	Get(string) (V, bool)
	Set(string, V)
	Clear()
	Range(func(string, V))
}

// Implementation with concurrent-map (generic API)
type concurrentMapImpl[V any] struct {
	m cmap.ConcurrentMap[string, V]
}

func newConcurrentMapImpl[V any]() concurrentMetricsBuilderMap[V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *concurrentMapImpl[V]) Get(key string) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (c *concurrentMapImpl[V]) Set(key string, value V) { _ = "STUB: not implemented"; return }

func (c *concurrentMapImpl[V]) Clear() { _ = "STUB: not implemented"; return }

func (c *concurrentMapImpl[V]) Range(f func(string, V)) {
	_ = "STUB: not implemented"

	// Implementation with sync.Map
	return
}

type syncMapImpl[V any] struct {
	m sync.Map
}

func newSyncMapImpl[V any]() concurrentMetricsBuilderMap[V] { _ = "STUB: not implemented"; return nil }

func (s *syncMapImpl[V]) Get(key string) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (s *syncMapImpl[V]) Set(key string, value V) { _ = "STUB: not implemented"; return }

func (s *syncMapImpl[V]) Clear() { _ = "STUB: not implemented"; return }

func (s *syncMapImpl[V]) Range(f func(string, V)) { _ = "STUB: not implemented"; return }

// Implementation with classic map and mutex

type mutexMapImpl[V any] struct {
	m     map[string]V
	mutex sync.RWMutex
}

func newMutexMapImpl[V any]() concurrentMetricsBuilderMap[V] { _ = "STUB: not implemented"; return nil }

func (mm *mutexMapImpl[V]) Get(key string) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (mm *mutexMapImpl[V]) Set(key string, value V) { _ = "STUB: not implemented"; return }

func (mm *mutexMapImpl[V]) Clear() { _ = "STUB: not implemented"; return }

func (mm *mutexMapImpl[V]) Range(f func(string, V)) { _ = "STUB: not implemented"; return }
