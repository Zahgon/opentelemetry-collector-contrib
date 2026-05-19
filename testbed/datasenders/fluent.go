// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"context"
	"os"
	"testing"

	"github.com/fluent/fluent-logger-golang/fluent"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

const (
	fluentDatafileVar = "FLUENT_DATA_SENDER_DATA_FILE"
	fluentPortVar     = "FLUENT_DATA_SENDER_RECEIVER_PORT"
)

// FluentLogsForwarder forwards logs to fluent forwader
type FluentLogsForwarder struct {
	testbed.DataSenderBase
	fluentLogger *fluent.Fluent
	dataFile     *os.File
}

// Ensure FluentLogsForwarder implements LogDataSender.
var _ testbed.LogDataSender = (*FluentLogsForwarder)(nil)

func NewFluentLogsForwarder(t *testing.T, port int) *FluentLogsForwarder {
	_ = "STUB: not implemented"
	return nil
}

// When FLUENT_DATA_SENDER_DATA_FILE is set, the data sender, writes to a
// file. This enables users to optionally run the e2e test against a real
// fluentd/fluentbit agent rather than using the fluent writer the data sender
// uses by default. In case, one is looking to point a real fluentd/fluentbit agent
// to the e2e test, they can do so by configuring the fluent agent to read from the
// file FLUENT_DATA_SENDER_DATA_FILE and forward data to FLUENT_DATA_SENDER_RECEIVER_PORT
// on 127.0.0.1.

func (*FluentLogsForwarder) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (*FluentLogsForwarder) Start() error { _ = "STUB: not implemented"; return nil }

func (f *FluentLogsForwarder) Stop() error { _ = "STUB: not implemented"; return nil }

func (f *FluentLogsForwarder) ConsumeLogs(_ context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*FluentLogsForwarder) convertLogToMap(lr plog.LogRecord) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (*FluentLogsForwarder) convertLogToJSON(lr plog.LogRecord) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (f *FluentLogsForwarder) Flush() { _ = "STUB: not implemented"; return }

func (f *FluentLogsForwarder) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*FluentLogsForwarder) ProtocolName() string { _ = "STUB: not implemented"; return "" }
