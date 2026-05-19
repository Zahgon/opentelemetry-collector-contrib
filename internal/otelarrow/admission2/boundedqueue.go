// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package admission2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/admission2"

import (
	"container/list"
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/trace"
	grpccodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	internalmetadata "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/internal/metadata"
)

var (
	ErrTooMuchWaiting  = status.Error(grpccodes.ResourceExhausted, "rejecting request, too much pending data")
	ErrRequestTooLarge = status.Errorf(grpccodes.InvalidArgument, "rejecting request, request is too large")
)

// BoundedQueue is a LIFO-oriented admission-controlled Queue.
type BoundedQueue struct {
	maxLimitAdmit    uint64
	maxLimitWait     uint64
	tracer           trace.Tracer
	telemetryBuilder *internalmetadata.TelemetryBuilder

	// lock protects currentAdmitted, currentWaiting, and waiters

	lock            sync.Mutex
	currentAdmitted uint64
	currentWaiting  uint64
	waiters         *list.List // of *waiter
}

var _ Queue = &BoundedQueue{}

// waiter is an item in the BoundedQueue waiters list.
type waiter struct {
	notify  N
	pending uint64
}

// NewBoundedQueue returns a LIFO-oriented Queue implementation which
// admits `maxLimitAdmit` bytes concurrently and allows up to
// `maxLimitWait` bytes to wait for admission.
func NewBoundedQueue(id component.ID, ts component.TelemetrySettings, maxLimitAdmit, maxLimitWait uint64) (Queue, error) {
	_ = "STUB: not implemented"
	return *new(Queue), nil
}

func (bq *BoundedQueue) inFlightCB() int64 {
	_ = "STUB: not implemented"
	// Note, see https://github.com/open-telemetry/otel-arrow/issues/270
	return 0
}

func (bq *BoundedQueue) waitingCB() int64 {
	_ = "STUB: not implemented"
	// Note, see https://github.com/open-telemetry/otel-arrow/issues/270
	return 0
}

// acquireOrGetWaiter returns with three distinct conditions depending
// on whether it was accepted, rejected, or asked to wait.
//
// - element=nil, error=nil: the fast success path
// - element=nil, error=non-nil: the fast failure path
// - element=non-nil, error=non-nil: the slow success path
func (bq *BoundedQueue) acquireOrGetWaiter(pending uint64) (*list.Element, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// when the request will never succeed because it is
	// individually over the total limit, fail fast.
}

// the fast success path.

// since we were unable to admit, check if we can wait.

// otherwise we need to wait

// Acquire implements Queue.
func (bq *BoundedQueue) Acquire(ctx context.Context, pending uint64) (ReleaseFunc, error) {
	_ = "STUB: not implemented"
	return *new(ReleaseFunc), nil
}

// We were also admitted, which can happen
// concurrently with cancellation. Make sure
// to release since no one else will do it.

// Remove ourselves from the list of waiters
// so that we can't be admitted in the future.

func (bq *BoundedQueue) admitWaitersLocked() { _ = "STUB: not implemented"; return }

// Ensure there is enough room to admit the next waiter.

// Returning means continuing to wait for the
// most recent arrival to get service by another release.

// Release the next waiter and tell it that it has been admitted.

func (bq *BoundedQueue) addWaiterLocked(pending uint64) *list.Element {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BoundedQueue) removeWaiterLocked(pending uint64, element *list.Element) {
	_ = "STUB: not implemented"
	return
}

func (bq *BoundedQueue) releaseLocked(pending uint64) { _ = "STUB: not implemented"; return }

func (bq *BoundedQueue) releaseFunc(pending uint64) ReleaseFunc {
	_ = "STUB: not implemented"
	return *new(ReleaseFunc)
}
