// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package eventhub // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver/internal/eventhub"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver/internal/trigger"
)

// LogsConsumer is a trigger.Consumer that unmarshals each Content message as logs
type LogsConsumer struct {
	unmarshaler plog.Unmarshaler
	nextLogs    consumer.Logs
}

// NewLogsConsumer returns a trigger.Consumer for Event Hub log bindings.
func NewLogsConsumer(unmarshaler plog.Unmarshaler, nextLogs consumer.Logs) *LogsConsumer {
	_ = "STUB: not implemented"
	return nil
}

// ConsumeEvents implements trigger.Consumer.
func (c *LogsConsumer) ConsumeEvents(ctx context.Context, req trigger.ParsedRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Decision: Log events that result in zero records are treated
// as anomalies and rejected as permanent errors.
