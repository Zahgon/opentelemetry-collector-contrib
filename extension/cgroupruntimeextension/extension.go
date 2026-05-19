// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cgroupruntimeextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/cgroupruntimeextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

type (
	undoFunc            func()
	maxProcsFn          func() (undoFunc, error)
	memLimitWithRatioFn func(float64) (undoFunc, error)
)

type cgroupRuntimeExtension struct {
	config *Config
	logger *zap.Logger

	// runtime modifiers
	maxProcsFn
	undoMaxProcsFn undoFunc

	memLimitWithRatioFn
	undoMemLimitFn        undoFunc
	memLimitRefreshCancel context.CancelFunc
	memLimitRefreshDone   chan struct{}
}

func newCgroupRuntime(cfg *Config, logger *zap.Logger, maxProcsFn maxProcsFn, memLimitFn memLimitWithRatioFn) *cgroupRuntimeExtension {
	_ = "STUB: not implemented"
	return nil
}

func (c *cgroupRuntimeExtension) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// automemlimit.WithRefreshInterval currently starts an unmanaged goroutine.
// Keep refresh scheduling in this extension so it can be cleanly stopped in Shutdown.
// See: https://github.com/KimMachineGun/automemlimit/issues/29

func (c *cgroupRuntimeExtension) refreshGoMemLimit(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (c *cgroupRuntimeExtension) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
