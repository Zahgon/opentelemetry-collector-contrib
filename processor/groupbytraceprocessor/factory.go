// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package groupbytraceprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

const (
	defaultWaitDuration   = time.Second
	defaultNumTraces      = 1_000_000
	defaultNumWorkers     = 1
	defaultDiscardOrphans = false
	defaultStoreOnDisk    = false
)

var (
	errDiskStorageNotSupported    = errors.New("option 'disk storage' not supported in this release")
	errDiscardOrphansNotSupported = errors.New("option 'discard orphans' not supported in this release")
)

// NewFactory returns a new factory for the Filter processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// createDefaultConfig creates the default configuration for the processor.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// not supported for now

// createTracesProcessor creates a trace processor based on this config.
func createTracesProcessor(
	_ context.Context,
	params processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

// the only supported storage for now
