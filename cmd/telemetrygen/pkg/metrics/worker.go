// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"math/rand/v2"
	"sync"
	"sync/atomic"

	"github.com/lightstep/go-expohisto/structure"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	types "github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen/pkg"
)

type worker struct {
	running                *atomic.Bool                 // pointer to shared flag that indicates it's time to stop the test
	metricName             string                       // name of metric to generate
	metricType             MetricType                   // type of metric to generate
	aggregationTemporality AggregationTemporality       // Temporality type to use
	exemplars              []metricdata.Exemplar[int64] // exemplars to attach to the metric
	numMetrics             int                          // how many metrics the worker has to generate (only when duration==0)
	enforceUnique          bool                         // if true, the worker will generate unique timeseries
	totalDuration          types.DurationWithInf        // how long to run the test for (overrides `numMetrics`)
	limitPerSecond         rate.Limit                   // how many metrics per second to generate
	wg                     *sync.WaitGroup              // notify when done
	logger                 *zap.Logger                  // logger
	index                  int                          // worker index
	clock                  Clock                        // clock
	batchSize              int                          // number of metrics to batch before flushing
	metricBuffer           []metricdata.ResourceMetrics // buffer for batching metrics
	bufferMutex            sync.Mutex                   // mutex for buffer access
	batch                  bool                         // whether to batch metrics
	loadSize               int                          // desired minimum size in MB of string data for each generated metric
	allowFailures          bool                         // whether to continue on export failures
	rand                   *rand.Rand                   // random number generator for exponential histogram generation
}

// We use a 15-element bounds slice for histograms below, so there must be 16 buckets here.
// From metrics.proto:
// The number of elements in bucket_counts array must be by one greater than
// the number of elements in explicit_bounds array.
var histogramBucketSamples = []struct {
	bucketCounts []uint64
	sum          int64
}{
	{
		[]uint64{0, 0, 1, 0, 0, 0, 3, 4, 1, 1, 0, 0, 0, 0, 0, 0},
		3940,
	},
	{
		[]uint64{0, 0, 0, 0, 0, 0, 2, 4, 4, 0, 0, 0, 0, 0, 0, 0},
		4455,
	},
	{
		[]uint64{0, 0, 0, 0, 0, 0, 1, 4, 3, 2, 0, 0, 0, 0, 0, 0},
		5337,
	},
	{
		[]uint64{0, 0, 1, 0, 1, 0, 2, 2, 1, 3, 0, 0, 0, 0, 0, 0},
		4477,
	},
	{
		[]uint64{0, 0, 0, 0, 0, 1, 3, 2, 2, 2, 0, 0, 0, 0, 0, 0},
		4670,
	},
	{
		[]uint64{0, 0, 0, 1, 1, 0, 1, 1, 1, 5, 0, 0, 0, 0, 0, 0},
		5670,
	},
	{
		[]uint64{0, 0, 0, 0, 0, 2, 1, 1, 4, 2, 0, 0, 0, 0, 0, 0},
		5091,
	},
	{
		[]uint64{0, 0, 2, 0, 0, 0, 2, 4, 1, 1, 0, 0, 0, 0, 0, 0},
		3420,
	},
	{
		[]uint64{0, 0, 0, 0, 0, 0, 1, 3, 2, 4, 0, 0, 0, 0, 0, 0},
		5917,
	},
	{
		[]uint64{0, 0, 1, 0, 1, 0, 0, 4, 4, 0, 0, 0, 0, 0, 0, 0},
		3988,
	},
}

func (w *worker) simulateMetrics(res *resource.Resource, exporter sdkmetric.Exporter, signalAttrs []attribute.KeyValue, tb *timeBox) {
	_ = "STUB: not implemented"
	return
}

// Add load size attributes if specified

// Bounds from https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/sdk.md#explicit-bucket-histogram-aggregation

// Generate realistic exponential histogram data using go-expohisto

// Add random values to the histogram
// Random count between 10-30

// Create the data point and convert using utility function

func (w *worker) addToBuffer(rm metricdata.ResourceMetrics, exporter sdkmetric.Exporter) {
	_ = "STUB: not implemented"
	return
}

func (w *worker) flushBuffer(exporter sdkmetric.Exporter) { _ = "STUB: not implemented"; return }

// expoHistToSDKExponentialDataPoint copies `lightstep/go-expohisto` structure.Histogram to
// metricdata.ExponentialHistogramDataPoint
func expoHistToSDKExponentialDataPoint(agg *structure.Histogram[float64], dp *metricdata.ExponentialHistogramDataPoint[int64]) {
	_ = "STUB: not implemented"
	return
}

// go-expohisto doesn't expose ZeroThreshold, use default

// Convert positive buckets

// Convert negative buckets
