// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

import (
	"errors"

	"go.uber.org/zap"
	"golang.org/x/sys/windows"
)

// Subscription is a subscription to a windows eventlog channel.
type Subscription struct {
	handle        uintptr
	signalEvent   windows.Handle
	Server        string
	startAt       string
	sessionHandle uintptr
	channel       string
	query         *string
	bookmark      Bookmark
	logger        *zap.Logger
}

// Open will open the subscription handle.
// It returns an error if the subscription handle is already open or if any step in the process fails.
// If the remote server is not reachable, it returns an error indicating the failure.
func (s *Subscription) Open(startAt string, sessionHandle uintptr, channel string, query *string, bookmark Bookmark) error {
	_ = "STUB: not implemented"
	return nil
}

// Close the signal handle on any failure path below.

// success — handle is now owned by the Subscription

// Close will close the subscription.
func (s *Subscription) Close() error { _ = "STUB: not implemented"; return nil }

var errSubscriptionHandleNotOpen = errors.New("subscription handle is not open")

// Wait blocks until the subscription has new events, the cancel event is signaled, or the timeout
// elapses. Returns true if events may be available (subscription signal fired or timeout elapsed),
// false if the cancel event was signaled (caller should stop).
func (s *Subscription) Wait(cancelEvent windows.Handle, timeoutMs uint32) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// WAIT_OBJECT_0+1 means the cancel event (index 1) fired — caller should stop.

// WAIT_OBJECT_0 (subscription signal) or WAIT_TIMEOUT (safety-net poll) — events may be available.

func (s *Subscription) Read(maxReads int) ([]Event, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// readWithRetry will read events from the subscription with dynamic batch sizing if the RPC_S_INVALID_BOUND error occurs.
func (s *Subscription) readWithRetry(maxReads int) ([]Event, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// close current subscription

// reopen subscription with the same parameters

// retry with half the batch size

// createFlags will create the necessary subscription flags from the supplied arguments.
func (*Subscription) createFlags(startAt string, bookmark Bookmark) uint32 {
	_ = "STUB: not implemented"
	return 0
}

// NewRemoteSubscription will create a new remote subscription with an empty handle.
func NewRemoteSubscription(server string, logger *zap.Logger) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// NewLocalSubscription will create a new local subscription with an empty handle.
func NewLocalSubscription(logger *zap.Logger) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}
