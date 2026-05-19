// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfxreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver"

import (
	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"go.opentelemetry.io/collector/pdata/plog"
)

// signalFxV2ToMetricsData converts SignalFx event proto data points to
// plog.LogRecordSlice. Returning the converted data and the number of dropped log
// records.
func signalFxV2EventsToLogRecords(events []*sfxpb.Event, lrs plog.LogRecordSlice) {
	_ = "STUB: not implemented"
	return
}

// The EventType field is stored as an attribute.

// SignalFx timestamps are in millis so convert to nanos by multiplying
// by 1 million.

// This gives us an unambiguous way of determining that a log record
// represents a SignalFx event, even if category is missing from the
// event.

// No way to tell what value type is without testing each
// individually.

// If there is no property value, just insert a null to
// record that the key was present.
