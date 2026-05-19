// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"

import (
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

const (
	cleanInterval = 5 * time.Minute
)

// CalculateFunc defines how to process metric values by the calculator. It
// passes previously received MetricValue, and the current raw value and timestamp
// as parameters. Returns true if the calculation is executed successfully.
type CalculateFunc func(prev *MetricValue, val any, timestamp time.Time) (any, bool)

func NewFloat64DeltaCalculator() MetricCalculator {
	_ = "STUB: not implemented"
	return *new(MetricCalculator)
}

func calculateDelta(prev *MetricValue, val any, _ time.Time) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// MetricCalculator is a calculator used to adjust metric values based on its previous record.
// Shutdown() must be called to clean up goroutines before program exit.
type MetricCalculator struct {
	// lock on write
	lock sync.Mutex
	// cache stores data with expiry time. The expiry is not supported at the moment.
	cache *MapWithExpiry
	// calculateFunc is the delegation for data processing
	calculateFunc CalculateFunc
}

// NewMetricCalculator Creates a metric calculator that enforces a five-minute time to live on cache entries.
func NewMetricCalculator(calculateFunc CalculateFunc) MetricCalculator {
	_ = "STUB: not implemented"
	return *new(MetricCalculator)
}

// Calculate accepts a new metric value identified by metric key (consists of metric metadata and labels),
// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/eacfde3fcbd46ba60a6db0e9a41977390c4883bd/internal/aws/metrics/metric_calculator.go#L88-L91
// and delegates the calculation with value and timestamp back to CalculateFunc for the result. Returns
// true if the calculation is executed successfully.
func (rm *MetricCalculator) Calculate(mKey Key, value any, timestamp time.Time) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// need to also lock cache to avoid the cleanup from removing entries while they are being processed.
// This is only likely to happen when data points come in close to expiration date.

func (rm *MetricCalculator) Shutdown() error { _ = "STUB: not implemented"; return nil }

type Key struct {
	MetricMetadata any
	MetricLabels   attribute.Distinct
}

func NewKey(metricMetadata any, labels map[string]string) Key {
	_ = "STUB: not implemented"
	return *new(Key)
}

type MetricValue struct {
	RawValue  any
	Timestamp time.Time
}

// Ticker allows us to mock time.Ticker in unit tests to have more deterministic tests.
type Ticker interface {
	C() <-chan time.Time
	Stop()
}

type realTicker struct {
	ticker *time.Ticker
}

func newRealTicker(d time.Duration) Ticker { _ = "STUB: not implemented"; return *new(Ticker) }

func (r *realTicker) C() <-chan time.Time { _ = "STUB: not implemented"; return nil }

func (r *realTicker) Stop() {
	_ = "STUB: not implemented"

	// MapWithExpiry act like a map which provides a method to clean up expired entries.
	// MapWithExpiry is not thread safe and locks must be managed by the owner of the Map by the use of Lock() and Unlock()
	return
}

type MapWithExpiry struct {
	lock      *sync.Mutex
	ttl       time.Duration
	entries   map[any]*MetricValue
	doneChan  chan struct{}
	newTicker func(d time.Duration) Ticker
}

// NewMapWithExpiry automatically starts a sweeper to enforce the maps TTL. ShutDown() must be called to ensure that these
// go routines are properly cleaned up ShutDown() must be called.
func NewMapWithExpiry(ttl time.Duration) *MapWithExpiry { _ = "STUB: not implemented"; return nil }

func (m *MapWithExpiry) Get(key Key) (*MetricValue, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *MapWithExpiry) Set(key Key, value MetricValue) { _ = "STUB: not implemented"; return }

func (m *MapWithExpiry) sweep(removeFunc func(time2 time.Time)) { _ = "STUB: not implemented"; return }

func (m *MapWithExpiry) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (m *MapWithExpiry) CleanUp(now time.Time) { _ = "STUB: not implemented"; return }

func (m *MapWithExpiry) Size() int { _ = "STUB: not implemented"; return 0 }

func (m *MapWithExpiry) Lock() { _ = "STUB: not implemented"; return }

func (m *MapWithExpiry) Unlock() { _ = "STUB: not implemented"; return }
