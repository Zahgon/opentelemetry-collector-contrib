// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	consumer "go.opentelemetry.io/collector/consumer"
)

type mockTracesDataConsumer struct {
	mock.Mock
}

// ConsumeTracesJSON provides a mock function with given fields: ctx, json
func (_m *mockTracesDataConsumer) consumeTracesJSON(ctx context.Context, json []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SetNextTracesConsumer provides a mock function with given fields: nextracesConsumer
func (_m *mockTracesDataConsumer) setNextTracesConsumer(nextracesConsumer consumer.Traces) {
	_ = "STUB: not implemented"
	return
}

func newMockTracesDataConsumer() *mockTracesDataConsumer { _ = "STUB: not implemented"; return nil }
