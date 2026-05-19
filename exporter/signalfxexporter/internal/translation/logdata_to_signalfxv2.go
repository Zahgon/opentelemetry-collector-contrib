// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation"

import (
	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

func LogRecordSliceToSignalFxV2(
	logger *zap.Logger,
	logs plog.LogRecordSlice,
	resourceAttrs pcommon.Map,
) ([]*sfxpb.Event, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func convertLogRecord(lr plog.LogRecord, resourceAttrs pcommon.Map, logger *zap.Logger) (*sfxpb.Event, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// keep a record of Resource attributes to add as dimensions
// so as not to modify LogRecord attributes

// LogRecord attribute takes priority

// Skip internal attributes

// Convert nanoseconds to nearest milliseconds, which is the unit of
// SignalFx event timestamps.

// EventType is a required field, if not set sfx event ingest will drop it

func attributeValToPropertyVal(v pcommon.Value) (*sfxpb.PropertyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
