// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package state // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector/internal/state"

import (
	"context"
	"sync"
	"time"
)

type PSConstants struct {
	RetryInterval time.Duration
	RetryGap      time.Duration
	MaxRetries    int
}

type TryLock struct {
	lock sync.Mutex
}

func (t *TryLock) TryExecute(fn func(int), arg int) { _ = "STUB: not implemented"; return }

func NewTryLock() *TryLock { _ = "STUB: not implemented"; return nil }

type CancelManager struct {
	cancelFunc context.CancelFunc
}

func (c *CancelManager) Cancel() { _ = "STUB: not implemented"; return }

func (c *CancelManager) UpdateFn(cancelFunc context.CancelFunc) { _ = "STUB: not implemented"; return }
