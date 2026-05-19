// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/text/message"
)

var printer = message.NewPrinter(message.MatchLanguage("en"))

// LoadGenerator is intended to be exercised by a TestCase to generate and send telemetry to an OtelcolRunner instance.
// The simplest ready implementation is the ProviderSender that unites a DataProvider with a DataSender.
type LoadGenerator interface {
	Start(options LoadOptions)
	Stop()
	IsReady() bool
	DataItemsSent() uint64
	IncDataItemsSent()
	PermanentErrors() uint64
	GetStats() string
}

// LoadOptions defines the options to use for generating the load.
type LoadOptions struct {
	// DataItemsPerSecond specifies how many spans, metric data points, or log
	// records to generate each second.
	DataItemsPerSecond int

	// ItemsPerBatch specifies how many spans, metric data points, or log
	// records per batch to generate. Should be greater than zero. The number
	// of batches generated per second will be DataItemsPerSecond/ItemsPerBatch.
	ItemsPerBatch int

	// Attributes to add to each generated data item. Can be empty.
	Attributes map[string]string

	// Parallel specifies how many goroutines to send from.
	Parallel int

	// MaxDelay defines the longest amount of time we can continue retrying for non-permanent errors.
	MaxDelay time.Duration
}

var _ LoadGenerator = (*ProviderSender)(nil)

// ProviderSender is a simple load generator.
type ProviderSender struct {
	Provider DataProvider
	Sender   DataSender

	// Number of data items (spans or metric data points) sent.
	dataItemsSent atomic.Uint64
	startedAt     time.Time
	startMutex    sync.Mutex

	// Number of permanent errors received
	permanentErrors    atomic.Uint64
	nonPermanentErrors atomic.Uint64

	stopOnce   sync.Once
	stopWait   sync.WaitGroup
	stopSignal chan struct{}

	options LoadOptions

	sendType     string
	generateFunc func() error
}

// NewLoadGenerator creates a ProviderSender to send DataProvider-generated telemetry via a DataSender.
func NewLoadGenerator(dataProvider DataProvider, sender DataSender) (LoadGenerator, error) {
	_ = "STUB: not implemented"
	return *new(LoadGenerator), nil
}

// Start the load.
func (ps *ProviderSender) Start(options LoadOptions) { _ = "STUB: not implemented"; return }

// 10 items per batch by default.

// retry for an additional 10 seconds by default

// Indicate that generation is in progress.

// Begin generation

// Stop the load.
func (ps *ProviderSender) Stop() { _ = "STUB: not implemented"; return }

// Signal generate() to stop.

// Wait for it to stop.

// Print stats.

func (ps *ProviderSender) IsReady() bool { _ = "STUB: not implemented"; return false }

// GetStats returns the stats as a printable string.
func (ps *ProviderSender) GetStats() string { _ = "STUB: not implemented"; return "" }

func (ps *ProviderSender) DataItemsSent() uint64 { _ = "STUB: not implemented"; return 0 }

func (ps *ProviderSender) PermanentErrors() uint64 { _ = "STUB: not implemented"; return 0 }

func (ps *ProviderSender) NonPermanentErrors() uint64 { _ = "STUB: not implemented"; return 0 }

// IncDataItemsSent is used when a test bypasses the ProviderSender and sends data
// directly via its Sender. This is necessary so that the total number of sent
// items in the end is correct, because the reports are printed from ProviderSender's
// fields. This is not the best way, a better approach would be to refactor the
// reports to use their own counter and load generator and other sending sources
// to contribute to this counter. This could be done as a future improvement.
func (ps *ProviderSender) IncDataItemsSent() { _ = "STUB: not implemented"; return }

func (ps *ProviderSender) generate() {
	_ = "STUB: not implemented"
	// Indicate that generation is done at the end
	return
}

// log the error if it is different from the previous result

// Send all pending generated data.

func (ps *ProviderSender) generateTrace() error { _ = "STUB: not implemented"; return nil }

// Generated data MUST be consumed once since the data counters
// are updated by the provider and not consuming the generated
// data will lead to accounting errors.

func (ps *ProviderSender) generateMetrics() error { _ = "STUB: not implemented"; return nil }

// Generated data MUST be consumed once since the data counters
// are updated by the provider and not consuming the generated
// data will lead to accounting errors.

func (ps *ProviderSender) generateLog() error { _ = "STUB: not implemented"; return nil }

// Generated data MUST be consumed once since the data counters
// are updated by the provider and not consuming the generated
// data will lead to accounting errors.

// perWorkerTickDuration calculates the tick interval each worker must observe in order to
// produce the desired average DataItemsPerSecond given the constraints of ItemsPerBatch and numWorkers.
//
// Of particular note are cases when the batchesPerSecond required of each worker is less than one due to a high
// number of workers relative to the desired DataItemsPerSecond. If the total batchesPerSecond is less than the
// number of workers then we are dealing with fractional batches per second per worker, so we need float arithmetic.
func (ps *ProviderSender) perWorkerTickDuration(numWorkers int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
