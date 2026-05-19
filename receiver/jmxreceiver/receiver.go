// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jmxreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver/internal/subprocess"
)

var _ receiver.Metrics = (*jmxMetricReceiver)(nil)

type jmxMetricReceiver struct {
	logger       *zap.Logger
	config       *Config
	subprocess   *subprocess.Subprocess
	params       receiver.Settings
	otlpReceiver receiver.Metrics
	nextConsumer consumer.Metrics
	configFile   string
	cancel       context.CancelFunc
}

func newJMXMetricReceiver(
	params receiver.Settings,
	config *Config,
	nextConsumer consumer.Metrics,
) *jmxMetricReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (jmx *jmxMetricReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Close the file

// Overwrite these environment variables to reduce attack surface

// ensure stdout/stderr buffer is read from.
// these messages are already debug logged when captured.

func (jmx *jmxMetricReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// insertDefault is a helper function to insert a default value for a configoptional.Optional type.
func insertDefault[T any](opt *configoptional.Optional[T]) error {
	_ = "STUB: not implemented"
	return nil
}

func (jmx *jmxMetricReceiver) buildOTLPReceiver() (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// We need to know the port OTLP receiver will use to specify w/ java properties and not
// rely on gRPC server's connection.
