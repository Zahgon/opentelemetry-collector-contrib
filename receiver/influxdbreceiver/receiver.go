// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package influxdbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver"

import (
	"context"
	"net/http"
	"sync"

	"github.com/influxdata/influxdb-observability/common"
	"github.com/influxdata/influxdb-observability/influx2otel"
	"github.com/influxdata/line-protocol/v2/lineprotocol"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

type metricsReceiver struct {
	nextConsumer       consumer.Metrics
	httpServerSettings *confighttp.ServerConfig
	converter          *influx2otel.LineProtocolToOtelMetrics

	server *http.Server
	wg     sync.WaitGroup

	logger common.Logger

	obsrecv *receiverhelper.ObsReport

	settings component.TelemetrySettings
}

func newMetricsReceiver(config *Config, settings receiver.Settings, nextConsumer consumer.Metrics) (*metricsReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *metricsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// InfluxDB 1.x
// InfluxDB 2.x

func (r *metricsReceiver) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

const (
	defaultPrecision = lineprotocol.Nanosecond
	dataFormat       = "influxdb"
)

var precisions = map[string]lineprotocol.Precision{
	"ns": lineprotocol.Nanosecond,
	"n":  lineprotocol.Nanosecond,
	"µs": lineprotocol.Microsecond,
	"µ":  lineprotocol.Microsecond,
	"us": lineprotocol.Microsecond,
	"u":  lineprotocol.Microsecond,
	"ms": lineprotocol.Millisecond,
	"s":  lineprotocol.Second,
}

func (r *metricsReceiver) handleWrite(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (*metricsReceiver) handlePing(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}
