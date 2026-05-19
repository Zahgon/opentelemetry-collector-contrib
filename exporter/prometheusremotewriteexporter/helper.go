// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"github.com/prometheus/prometheus/prompb"
)

type batchTimeSeriesState struct {
	// Track batch sizes sent to avoid over allocating huge buffers.
	// This helps in the case where large batches are sent to avoid allocating too much unused memory
	nextTimeSeriesBufferSize     int
	nextMetricMetadataBufferSize int
	nextRequestBufferSize        int
}

func newBatchTimeServicesState() *batchTimeSeriesState { _ = "STUB: not implemented"; return nil }

// batchTimeSeries splits series into multiple batch write requests.
func batchTimeSeries(tsMap map[string]*prompb.TimeSeries, maxBatchByteSize int, m []*prompb.MetricMetadata, state *batchTimeSeriesState) ([]*prompb.WriteRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allocate a buffer size of at least 10, or twice the last # of requests we sent

// Allocate a time series buffer 2x the last time series batch size or the length of the input if smaller

// Allocate a metric metadata buffer 2x the last metric metadata batch size or the length of the input if smaller

func convertTimeseriesToRequest(tsArray []prompb.TimeSeries) *prompb.WriteRequest {
	_ = "STUB: not implemented"
	// the remote_write endpoint only requires the timeseries.
	// otlp defines its own way to handle metric metadata
	return nil
}

// Prometheus requires time series to be sorted by Timestamp to avoid out of order problems.
// See:
// * https://github.com/open-telemetry/wg-prometheus/issues/10
// * https://github.com/open-telemetry/opentelemetry-collector/issues/2315

func convertMetadataToRequest(m []prompb.MetricMetadata) *prompb.WriteRequest {
	_ = "STUB: not implemented"
	return nil
}

func orderBySampleTimestamp(tsArray []prompb.TimeSeries) []prompb.TimeSeries {
	_ = "STUB: not implemented"
	return nil
}
