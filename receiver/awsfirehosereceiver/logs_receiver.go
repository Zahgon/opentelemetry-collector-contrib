// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog"
)

const defaultLogsEncoding = cwlog.TypeStr

// logsConsumer implements the firehoseConsumer
// to use a logs consumer and unmarshaler.
type logsConsumer struct {
	config   *Config
	settings receiver.Settings

	// consumer passes the translated logs on to the
	// next consumer.
	consumer consumer.Logs
	// unmarshaler is the configured plog.Unmarshaler
	// to use when processing the records.
	unmarshaler plog.Unmarshaler
}

var _ firehoseConsumer = (*logsConsumer)(nil)

// newLogsReceiver creates a new instance of the receiver
// with a logsConsumer.
func newLogsReceiver(
	config *Config,
	set receiver.Settings,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// Start sets the consumer's log unmarshaler to either a built-in
// unmarshaler or one loaded from an encoding extension.
func (c *logsConsumer) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Consume uses the configured unmarshaler to deserialize each record,
// with each resulting plog.Logs being sent to the next consumer as
// they are unmarshalled.
func (c *logsConsumer) Consume(ctx context.Context, nextRecord nextRecordFunc, commonAttributes map[string]string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
