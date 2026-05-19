// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsxrayreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const defaultEndpoint = "localhost:2000"

// NewFactory creates a factory for AWS receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	// reference the existing default configurations provided
	// in the X-Ray daemon:
	// https://github.com/aws/aws-xray-daemon/blob/master/pkg/cfg/cfg.go#L99
	return *new(component.Config)
}

func createTracesReceiver(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}
