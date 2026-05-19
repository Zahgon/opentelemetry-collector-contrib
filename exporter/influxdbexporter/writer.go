// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package influxdbexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/influxdbexporter"

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/influxdata/influxdb-observability/common"
	"github.com/influxdata/influxdb-observability/otel2influx"
	"github.com/influxdata/line-protocol/v2/lineprotocol"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

var _ otel2influx.InfluxWriter = (*influxHTTPWriter)(nil)

type influxHTTPWriter struct {
	encoderPool sync.Pool
	httpClient  *http.Client

	httpClientSettings confighttp.ClientConfig
	telemetrySettings  component.TelemetrySettings
	writeURL           string
	payloadMaxLines    int
	payloadMaxBytes    int

	logger common.Logger
}

func newInfluxHTTPWriter(logger common.Logger, config *Config, telemetrySettings component.TelemetrySettings) (*influxHTTPWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func composeWriteURL(config *Config) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Start implements component.StartFunc
func (w *influxHTTPWriter) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *influxHTTPWriter) NewBatch() otel2influx.InfluxWriterBatch {
	_ = "STUB: not implemented"
	return *new(otel2influx.InfluxWriterBatch)
}

var _ otel2influx.InfluxWriterBatch = (*influxHTTPWriterBatch)(nil)

type influxHTTPWriterBatch struct {
	*influxHTTPWriter
	encoder      *lineprotocol.Encoder
	payloadLines int
}

func newInfluxHTTPWriterBatch(w *influxHTTPWriter) *influxHTTPWriterBatch {
	_ = "STUB: not implemented"
	return nil
}

// EnqueuePoint emits a set of line protocol attributes (metrics, tags, fields, timestamp)
// to the internal line protocol buffer.
// If the buffer is full, it will be flushed by calling WriteBatch.
func (b *influxHTTPWriterBatch) EnqueuePoint(ctx context.Context, measurement string, tags map[string]string, fields map[string]any, ts time.Time, _ common.InfluxMetricValueType) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteBatch sends the internal line protocol buffer to InfluxDB.
func (b *influxHTTPWriterBatch) WriteBatch(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Success

// Retryable error

// Terminal error

type tag struct {
	k, v string
}

// optimizeTags sorts tags by key and removes tags with empty keys or values
func (b *influxHTTPWriterBatch) optimizeTags(m map[string]string) []tag {
	_ = "STUB: not implemented"
	return nil
}

func (b *influxHTTPWriterBatch) convertFields(m map[string]any) (fields map[string]lineprotocol.Value) {
	_ = "STUB: not implemented"
	return nil
}
