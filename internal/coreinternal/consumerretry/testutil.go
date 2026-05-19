// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumerretry // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/consumerretry"

import (
	"context"
	"sync/atomic"

	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/plog"
)

type MockLogsRejecter struct {
	consumertest.LogsSink
	rejectCount *atomic.Int32
	acceptAfter int32
}

// NewMockLogsRejecter creates new MockLogsRejecter. acceptAfter is a number of rejects before accepting,
// 0 means always accept, -1 means always reject with permanent error
func NewMockLogsRejecter(acceptAfter int32) *MockLogsRejecter {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockLogsRejecter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// mockPartialLogsRejecter is a mock LogsConsumer that accepts only one logs object and rejects the rest.
type mockPartialLogsRejecter struct {
	consumertest.LogsSink
}

func (m *mockPartialLogsRejecter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}
