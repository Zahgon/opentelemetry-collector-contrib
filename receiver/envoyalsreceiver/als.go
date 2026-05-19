// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package envoyalsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/envoyalsreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"google.golang.org/grpc"
)

type alsReceiver struct {
	conf         *Config
	nextConsumer consumer.Logs
	settings     receiver.Settings
	serverGRPC   *grpc.Server

	obsrepGRPC *receiverhelper.ObsReport
}

func (r *alsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *alsReceiver) startGRPCServer(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *alsReceiver) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func newALSReceiver(cfg *Config, nextConsumer consumer.Logs, settings receiver.Settings) (*alsReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
