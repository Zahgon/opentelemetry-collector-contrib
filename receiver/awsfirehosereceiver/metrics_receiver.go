// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"
)

const defaultMetricsEncoding = cwmetricstream.TypeStr

// The metricsConsumer implements the firehoseConsumer
// to use a metrics consumer and unmarshaler.
type metricsConsumer struct {
	config   *Config
	settings receiver.Settings
	// consumer passes the translated metrics on to the
	// next consumer.
	consumer consumer.Metrics
	// unmarshaler is the configured pmetric.Unmarshaler
	// to use when processing the records.
	unmarshaler pmetric.Unmarshaler
}

var _ firehoseConsumer = (*metricsConsumer)(nil)

// newMetricsReceiver creates a new instance of the receiver
// with a metricsConsumer.
func newMetricsReceiver(
	config *Config,
	set receiver.Settings,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func (c *metricsConsumer) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// newUnmarshalerFromEncoding creates a new unmarshaler from
// aws cloudwatch metric streams encoding extension.
func (c *metricsConsumer) newUnmarshalerFromEncoding(
	ctx context.Context,
	encoding string,
	format string,
) (pmetric.Unmarshaler, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Unmarshaler), nil
}

// Consume uses the configured unmarshaler to deserialize each record,
// with each resulting pmetric.Metrics being sent to the next consumer
// as they are unmarshalled.
func (c *metricsConsumer) Consume(ctx context.Context, nextRecord nextRecordFunc, commonAttributes map[string]string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
