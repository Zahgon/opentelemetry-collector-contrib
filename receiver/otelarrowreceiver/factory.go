// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelarrowreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

const (
	defaultGRPCEndpoint = "0.0.0.0:4317"

	defaultMemoryLimitMiB  = 128
	defaultRequestLimitMiB = 128
	defaultWaitingLimitMiB = 32
)

// NewFactory creates a new OTel-Arrow receiver factory.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// We almost write 0 bytes, so no need to tune WriteBufferSize.

// createTraces creates a trace receiver based on provided config.
func createTraces(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

// createMetrics creates a metrics receiver based on provided config.
func createMetrics(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// createLog creates a log receiver based on provided config.
func createLog(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// This is the map of already created OTel-Arrow receivers for particular configurations.
// We maintain this map because the Factory is asked trace and metric receivers separately
// when it gets CreateTraces() and CreateMetrics() but they must not
// create separate objects, they must use one otelArrowReceiver object per configuration.
// When the receiver is shutdown it should be removed from this map so the same configuration
// can be recreated successfully.
var receivers = sharedcomponent.NewSharedComponents()
