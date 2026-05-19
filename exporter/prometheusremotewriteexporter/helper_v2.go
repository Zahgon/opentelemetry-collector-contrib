// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
)

func batchTimeSeriesV2(tsMap map[string]*writev2.TimeSeries, symbolsTable writev2.SymbolsTable, maxBatchByteSize int, state *batchTimeSeriesState) ([]*writev2.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate symbols table size once since it's shared across batches

// Initialize with symbols table size

// Reset to symbols table size for new batch

// TODO only sent necessary part of the symbolsTable

func convertTimeseriesToRequestV2(tsArray []writev2.TimeSeries, symbolsTable writev2.SymbolsTable) *writev2.Request {
	_ = "STUB: not implemented"
	return nil

	// Prometheus requires time series to be sorted by Timestamp to avoid out of order problems.
	// See:
	// * https://github.com/open-telemetry/wg-prometheus/issues/10
	// * https://github.com/open-telemetry/opentelemetry-collector/issues/2315
	// TODO: try to sort while batching?
}

func orderBySampleTimestampV2(tsArray []writev2.TimeSeries) []writev2.TimeSeries {
	_ = "STUB: not implemented"
	return nil
}
