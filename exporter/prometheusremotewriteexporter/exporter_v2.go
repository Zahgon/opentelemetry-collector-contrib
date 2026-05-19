// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"context"
	"net/http"

	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func (prwe *prwExporter) pushMetricsV2(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Call export even if a conversion error, since there may be points that were successfully converted.

// exportV2 sends a Snappy-compressed writev2.Request containing writev2.TimeSeries to a remote write endpoint.
func (prwe *prwExporter) exportV2(ctx context.Context, requests []*writev2.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// used to wait for workers to be finished

// Run concurrencyLimit of workers until there
// is no more requests to execute in the input channel.

func (prwe *prwExporter) handleRequestsV2(ctx context.Context, input chan *writev2.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Check firstly to ensure that the context wasn't cancelled.

func (prwe *prwExporter) handleExportV2(ctx context.Context, symbolsTable writev2.SymbolsTable, tsMap map[string]*writev2.TimeSeries) error {
	_ = "STUB: not implemented"
	// There are no metrics to export, so return.
	return nil
}

// TODO implement WAl support, can be done after #15277 is fixed

func (prwe *prwExporter) handleHeader(ctx context.Context, resp *http.Response, headerName, metricType string, recordFunc func(context.Context, int64)) {
	_ = "STUB: not implemented"
	return
}

func (prwe *prwExporter) handleWrittenHeaders(ctx context.Context, resp *http.Response) {
	_ = "STUB: not implemented"
	return
}
