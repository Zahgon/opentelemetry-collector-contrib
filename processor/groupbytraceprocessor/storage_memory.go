// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package groupbytraceprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"

import (
	"sync"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor/internal/metadata"
)

type memoryStorage struct {
	sync.RWMutex
	content                   map[pcommon.TraceID][]ptrace.ResourceSpans
	telemetry                 *metadata.TelemetryBuilder
	stopped                   bool
	stoppedLock               sync.RWMutex
	metricsCollectionInterval time.Duration
}

var _ storage = (*memoryStorage)(nil)

func newMemoryStorage(telemetry *metadata.TelemetryBuilder) *memoryStorage {
	_ = "STUB: not implemented"
	return nil
}

func (st *memoryStorage) createOrAppend(traceID pcommon.TraceID, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// getting zero value is fine

func (st *memoryStorage) get(traceID pcommon.TraceID) ([]ptrace.ResourceSpans, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// delete will return a reference to a ResourceSpans. Changes to the returned object may not be applied
// to the version in the storage.
func (st *memoryStorage) delete(traceID pcommon.TraceID) ([]ptrace.ResourceSpans, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (st *memoryStorage) start() error { _ = "STUB: not implemented"; return nil }

func (st *memoryStorage) shutdown() error { _ = "STUB: not implemented"; return nil }

func (st *memoryStorage) periodicMetrics() { _ = "STUB: not implemented"; return }

func (st *memoryStorage) count() int { _ = "STUB: not implemented"; return 0 }
