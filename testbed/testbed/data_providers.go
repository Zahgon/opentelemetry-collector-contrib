// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"sync/atomic"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pipeline"
)

// DataProvider defines the interface for generators of test data used to drive various end-to-end tests.
type DataProvider interface {
	// SetLoadGeneratorCounters supplies pointers to LoadGenerator counters.
	// The data provider implementation should increment these as it generates data.
	SetLoadGeneratorCounters(dataItemsGenerated *atomic.Uint64)
	// GenerateTraces returns an internal Traces instance with an OTLP ResourceSpans slice populated with test data.
	GenerateTraces() (ptrace.Traces, bool)
	// GenerateMetrics returns an internal MetricData instance with an OTLP ResourceMetrics slice of test data.
	GenerateMetrics() (pmetric.Metrics, bool)
	// GenerateLogs returns the internal plog.Logs format
	GenerateLogs() (plog.Logs, bool)
}

// perfTestDataProvider in an implementation of the DataProvider for use in performance tests.
// Tracing IDs are based on the incremented batch and data items counters.
type perfTestDataProvider struct {
	options            LoadOptions
	traceIDSequence    atomic.Uint64
	dataItemsGenerated *atomic.Uint64
}

// NewPerfTestDataProvider creates an instance of perfTestDataProvider which generates test data based on the sizes
// specified in the supplied LoadOptions.
func NewPerfTestDataProvider(options LoadOptions) DataProvider {
	_ = "STUB: not implemented"
	return *new(DataProvider)
}

func (dp *perfTestDataProvider) SetLoadGeneratorCounters(dataItemsGenerated *atomic.Uint64) {
	_ = "STUB: not implemented"
	return
}

func (dp *perfTestDataProvider) GenerateTraces() (ptrace.Traces, bool) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), false
}

// Create a span.

// Additional attributes.

func (dp *perfTestDataProvider) GenerateMetrics() (pmetric.Metrics, bool) {
	_ = "STUB: not implemented"
	// Generate 7 data points per metric.
	return *new(pmetric.Metrics), false
}

// Generate data points for the metric.

func (dp *perfTestDataProvider) GenerateLogs() (plog.Logs, bool) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), false
}

// goldenDataProvider is an implementation of DataProvider for use in correctness tests.
// Provided data from the "Golden" dataset generated using pairwise combinatorial testing techniques.
type goldenDataProvider struct {
	tracePairsFile     string
	spanPairsFile      string
	dataItemsGenerated *atomic.Uint64

	tracesGenerated []ptrace.Traces
	tracesIndex     int

	metricPairsFile  string
	metricsGenerated []pmetric.Metrics
	metricsIndex     int
}

// NewGoldenDataProvider creates a new instance of goldenDataProvider which generates test data based
// on the pairwise combinations specified in the tracePairsFile and spanPairsFile input variables.
func NewGoldenDataProvider(tracePairsFile, spanPairsFile, metricPairsFile string) DataProvider {
	_ = "STUB: not implemented"
	return *new(DataProvider)
}

func (dp *goldenDataProvider) SetLoadGeneratorCounters(dataItemsGenerated *atomic.Uint64) {
	_ = "STUB: not implemented"
	return
}

func (dp *goldenDataProvider) GenerateTraces() (ptrace.Traces, bool) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), false
}

func (dp *goldenDataProvider) GenerateMetrics() (pmetric.Metrics, bool) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), false
}

func (*goldenDataProvider) GenerateLogs() (plog.Logs, bool) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), false
}

// FileDataProvider in an implementation of the DataProvider for use in performance tests.
// The data to send is loaded from a file. The file should contain one JSON-encoded
// Export*ServiceRequest Protobuf message. The file can be recorded using the "file"
// exporter (note: "file" exporter writes one JSON message per line, FileDataProvider
// expects just a single JSON message in the entire file).
type FileDataProvider struct {
	dataItemsGenerated *atomic.Uint64
	logs               plog.Logs
	metrics            pmetric.Metrics
	traces             ptrace.Traces
	ItemsPerBatch      int
}

// NewFileDataProvider creates an instance of FileDataProvider which generates test data
// loaded from a file.
func NewFileDataProvider(filePath string, dataType pipeline.Signal) (*FileDataProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load the message from the file and count the data points.

func (dp *FileDataProvider) SetLoadGeneratorCounters(dataItemsGenerated *atomic.Uint64) {
	_ = "STUB: not implemented"
	return
}

func (dp *FileDataProvider) GenerateTraces() (ptrace.Traces, bool) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), false
}

func (dp *FileDataProvider) GenerateMetrics() (pmetric.Metrics, bool) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), false
}

func (dp *FileDataProvider) GenerateLogs() (plog.Logs, bool) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), false
}
