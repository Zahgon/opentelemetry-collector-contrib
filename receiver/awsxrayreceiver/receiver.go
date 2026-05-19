// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsxrayreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/udppoller"
)

const (
	// number of goroutines polling the UDP socket.
	// https://github.com/aws/aws-xray-daemon/blob/master/pkg/cfg/cfg.go#L184
	maxPollerCount = 2
)

// xrayReceiver implements the receiver.Traces interface for converting
// AWS X-Ray segment document into the OT internal trace format.
type xrayReceiver struct {
	poller   udppoller.Poller
	server   proxy.Server
	proxyCfg *proxy.Config
	settings receiver.Settings
	consumer consumer.Traces
	obsrecv  *receiverhelper.ObsReport
	registry telemetry.Registry
}

func newReceiver(config *Config,
	consumer consumer.Traces,
	set receiver.Settings,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func (x *xrayReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Might want to pass `host` into read() below to report a fatal error

func (x *xrayReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (x *xrayReceiver) start() { _ = "STUB: not implemented"; return }
