// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package netflowreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/netflowreceiver"

import (
	"context"

	"github.com/netsampler/goflow2/v2/utils"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
)

var _ utils.ReceiverCallback = (*dropHandler)(nil)

type dropHandler struct {
	logger *zap.Logger
}

func (d dropHandler) Dropped(msg utils.Message) { _ = "STUB: not implemented"; return }

type netflowReceiver struct {
	config      Config
	logger      *zap.Logger
	udpReceiver *utils.UDPReceiver
	logConsumer consumer.Logs
}

func newNetflowLogsReceiver(params receiver.Settings, cfg Config, consumer consumer.Logs) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	// UDP receiver configuration
	return *new(receiver.Logs), nil
}

func (nr *netflowReceiver) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	// The function that will decode packets
	return nil
}

// This runs until the receiver is stoppped, consuming from an error channel

func (nr *netflowReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// buildDecodeFunc creates a decode function based on the scheme
// This is the fuction that will be invoked for every netflow packet received
// The function depends on the type of schema (netflow, sflow, flow)
func (nr *netflowReceiver) buildDecodeFunc() (utils.DecoderFunc, error) {
	_ = "STUB: not implemented"
	// Eventually this can be used to configure mappings
	return *new(utils.DecoderFunc), nil
}

// converts configuration into a format that can be used by a protobuf producer

// We use a goflow2 proto producer to produce messages using protobuf format

// the otel log producer converts those messages into OpenTelemetry logs
// it is a wrapper around the protobuf producer

// handleErrors handles errors from the listener
// We don't want the receiver to stop if there is an error processing a packet
func (nr *netflowReceiver) handleErrors() { _ = "STUB: not implemented"; return }
