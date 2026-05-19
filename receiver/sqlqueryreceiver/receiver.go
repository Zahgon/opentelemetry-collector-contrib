// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlqueryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"

import (
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"
)

func createLogsReceiverFunc(sqlOpenerFunc sqlquery.SQLOpenerFunc, clientProviderFunc sqlquery.ClientProviderFunc) receiver.CreateLogsFunc {
	_ = "STUB: not implemented"
	return *new(receiver.CreateLogsFunc)
}

func createMetricsReceiverFunc(sqlOpenerFunc sqlquery.SQLOpenerFunc, clientProviderFunc sqlquery.ClientProviderFunc) receiver.CreateMetricsFunc {
	_ = "STUB: not implemented"
	return *new(receiver.CreateMetricsFunc)
}
