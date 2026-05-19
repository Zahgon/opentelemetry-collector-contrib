// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
)

// combinedLogsReceiver wraps alerts and log receivers in a single log receiver to be consumed by the factory
type combinedLogsReceiver struct {
	alerts     *alertsReceiver
	logs       *logsReceiver
	events     *eventsReceiver
	accessLogs *accessLogsReceiver
	storageID  *component.ID
	id         component.ID
}

// Starts up the combined MongoDB Atlas Logs and Alert Receiver
func (c *combinedLogsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shuts down the combined MongoDB Atlas Logs and Alert Receiver
func (c *combinedLogsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
