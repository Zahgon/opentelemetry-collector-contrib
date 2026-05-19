// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Originally copied from https://github.com/signalfx/signalfx-agent/blob/fbc24b0fdd3884bd0bbfbd69fe3c83f49d4c0b77/pkg/apm/tracetracker/cache.go

package tracetracker // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/tracetracker"

import (
	"container/list"
	"sync"
	"time"
)

type CacheKey struct {
	dimName  string
	dimValue string
	value    string
}

type cacheElem struct {
	LastSeen time.Time
	Obj      *CacheKey
}

type TimeoutCache struct {
	sync.Mutex

	// How long to keep sending metrics for a particular service name after it
	// is last seen
	timeout time.Duration
	// A linked list of keys sorted by time last seen
	keysByTime *list.List
	// Which keys are active currently.  The value is an entry in the
	// keysByTime linked list so that it can be quickly accessed and
	// moved to the back of the list.
	keysActive map[CacheKey]*list.Element

	// Internal metrics
	ActiveCount int64
	PurgedCount int64

	maxSize         int64
	maxSizeExpiryTS time.Time
}

// returns whether the cache is full
func (t *TimeoutCache) IsFull() bool { _ = "STUB: not implemented"; return false }

func (t *TimeoutCache) SetMaxSize(maxSize int64, now time.Time) { _ = "STUB: not implemented"; return }

// RunIfKeyDoesNotExist locks and runs the supplied function if the key does not exist.
// Be careful not to perform cache operations inside of this function because they will deadlock
func (t *TimeoutCache) RunIfKeyDoesNotExist(o *CacheKey, fn func()) {
	_ = "STUB: not implemented"
	return
}

// UpdateOrCreate
func (t *TimeoutCache) UpdateOrCreate(o *CacheKey, now time.Time) (isNew bool) {
	_ = "STUB: not implemented"
	return false
}

// UpdateIfExists
func (t *TimeoutCache) UpdateIfExists(o *CacheKey, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *TimeoutCache) GetPurgeable(now time.Time) []*CacheKey {
	_ = "STUB: not implemented"
	return nil
}

// If this one isn't timed out, nothing else in the list is either.

func (t *TimeoutCache) Delete(key *CacheKey) { _ = "STUB: not implemented"; return }

// PurgeOld
func (t *TimeoutCache) PurgeOld(now time.Time, onPurge func(*CacheKey)) {
	_ = "STUB: not implemented"
	return
}

// If this one isn't timed out, nothing else in the list is either.

func (t *TimeoutCache) GetActiveCount() int64 { _ = "STUB: not implemented"; return 0 }

func (t *TimeoutCache) GetPurgedCount() int64 { _ = "STUB: not implemented"; return 0 }

func NewTimeoutCache(timeout time.Duration) *TimeoutCache { _ = "STUB: not implemented"; return nil }
