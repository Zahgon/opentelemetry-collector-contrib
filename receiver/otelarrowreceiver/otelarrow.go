// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelarrowreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/admission2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/netstats"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/arrow"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/logs"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/metrics"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/trace"
)

// otelArrowReceiver is the type that exposes Trace and Metrics reception.
type otelArrowReceiver struct {
	cfg        *Config
	serverGRPC *grpc.Server

	tracesReceiver  *trace.Receiver
	metricsReceiver *metrics.Receiver
	logsReceiver    *logs.Receiver
	arrowReceiver   *arrow.Receiver
	shutdownWG      sync.WaitGroup

	obsrepGRPC   *receiverhelper.ObsReport
	netReporter  *netstats.NetworkReporter
	boundedQueue admission2.Queue

	settings receiver.Settings
}

// newOTelArrowReceiver just creates the OpenTelemetry receiver services. It is the caller's
// responsibility to invoke the respective Start*Reception methods as well
// as the various Stop*Reception methods to end it.
func newOTelArrowReceiver(cfg *Config, set receiver.Settings) (*otelArrowReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *otelArrowReceiver) startGRPCServer(cfg configgrpc.ServerConfig, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *otelArrowReceiver) startProtocolServers(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// in which case the default is selected in the arrowRecord package.

// Start runs the trace receiver on the gRPC server. Currently
// it also enables the metrics receiver too.
func (r *otelArrowReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown is a method to turn off receiving.
func (r *otelArrowReceiver) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *otelArrowReceiver) registerTraceConsumer(tc consumer.Traces) {
	_ = "STUB: not implemented"
	return
}

func (r *otelArrowReceiver) registerMetricsConsumer(mc consumer.Metrics) {
	_ = "STUB: not implemented"
	return
}

func (r *otelArrowReceiver) registerLogsConsumer(lc consumer.Logs) {
	_ = "STUB: not implemented"
	return
}

var _ arrow.Consumers = &otelArrowReceiver{}

func (r *otelArrowReceiver) Traces() consumer.Traces {
	_ = "STUB: not implemented"
	return *new(consumer.Traces)
}

func (r *otelArrowReceiver) Metrics() consumer.Metrics {
	_ = "STUB: not implemented"
	return *new(consumer.Metrics)
}

func (r *otelArrowReceiver) Logs() consumer.Logs {
	_ = "STUB: not implemented"
	return *new(consumer.Logs)
}
