// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver"

import (
	"errors"
	"io"
	"net/url"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	translator "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/splunk"
)

const (
	// splunk metadata
	index      = "index"
	source     = "source"
	sourcetype = "sourcetype"
	host       = "host"
	queryTime  = "time"
)

var errCannotConvertValue = errors.New("cannot convert field value to attribute")

// splunkHecToLogData transforms splunk events into logs
func splunkHecToLogData(logger *zap.Logger, events []*translator.Event, resourceCustomizer func(pcommon.Resource), config *Config) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// The SourceType field is the most logical "name" of the event.

// Set event fields first, so the specialized attributes overwrite them if needed.

// splunkHecRawToLogData transforms raw splunk event into log
func splunkHecRawToLogData(bodyReader io.Reader, query url.Values, resourceCustomizer func(pcommon.Resource), config *Config, timestamp pcommon.Timestamp) (plog.Logs, int, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), 0, nil
}

func appendSplunkMetadata(rl plog.ResourceLogs, attrs translator.HecToOtelAttrs, host, source, sourceType, index string) {
	_ = "STUB: not implemented"
	return
}

func convertToValue(logger *zap.Logger, src any, dest pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func convertToSliceVal(logger *zap.Logger, value []any, dest pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func convertToAttributeMap(logger *zap.Logger, value map[string]any, dest pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}
