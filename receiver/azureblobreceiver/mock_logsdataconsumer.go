// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	consumer "go.opentelemetry.io/collector/consumer"
)

type mockLogsDataConsumer struct {
	mock.Mock
}

// ConsumeLogsJSON provides a mock function with given fields: ctx, json
func (_m *mockLogsDataConsumer) consumeLogsJSON(ctx context.Context, json []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SetNextLogsConsumer provides a mock function with given fields: nextLogsConsumer
func (_m *mockLogsDataConsumer) setNextLogsConsumer(nextLogsConsumer consumer.Logs) {
	_ = "STUB: not implemented"
	return
}

func newMockLogsDataConsumer() *mockLogsDataConsumer { _ = "STUB: not implemented"; return nil }
