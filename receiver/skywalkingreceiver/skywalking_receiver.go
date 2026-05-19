// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalkingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver/internal/metrics"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver/internal/trace"
)

// configuration defines the behavior and the ports that
// the Skywalking receiver will use.
type configuration struct {
	CollectorHTTPPort           int
	CollectorHTTPSettings       confighttp.ServerConfig
	CollectorGRPCPort           int
	CollectorGRPCServerSettings configgrpc.ServerConfig
}

// Receiver type is used to receive spans that were originally intended to be sent to Skywalking.
// This receiver is basically a Skywalking collector.
type swReceiver struct {
	config *configuration

	grpc            *grpc.Server
	collectorServer *http.Server

	goroutines sync.WaitGroup

	settings receiver.Settings

	traceReceiver *trace.Receiver

	metricsReceiver *metrics.Receiver

	dummyReportService *dummyReportService
}

// newSkywalkingReceiver creates a TracesReceiver that receives traffic as a Skywalking collector
func newSkywalkingReceiver(
	config *configuration,
	set receiver.Settings,
) *swReceiver {
	_ = "STUB: not implemented"
	return nil
}

// registerTraceConsumer register a TracesReceiver that receives trace
func (sr *swReceiver) registerTraceConsumer(tc consumer.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// registerTraceConsumer register a TracesReceiver that receives trace
func (sr *swReceiver) registerMetricsConsumer(mc consumer.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (sr *swReceiver) collectorGRPCAddr() string { _ = "STUB: not implemented"; return "" }

func (sr *swReceiver) collectorGRPCEnabled() bool { _ = "STUB: not implemented"; return false }

func (sr *swReceiver) collectorHTTPEnabled() bool { _ = "STUB: not implemented"; return false }

func (sr *swReceiver) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (sr *swReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (sr *swReceiver) startCollector(host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}
