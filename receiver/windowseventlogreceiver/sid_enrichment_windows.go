// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windowseventlogreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver/internal/sidcache"
)

// sidEnrichingConsumer wraps a logs consumer to enrich Windows events with SID resolution
type sidEnrichingConsumer struct {
	next     consumer.Logs
	sidCache sidcache.Cache
	logger   *zap.Logger
}

// newSIDEnrichingConsumer creates a new SID enriching consumer wrapper
func newSIDEnrichingConsumer(next consumer.Logs, cache sidcache.Cache, logger *zap.Logger) *sidEnrichingConsumer {
	_ = "STUB: not implemented"
	return nil
}

// Capabilities returns the consumer capabilities
func (s *sidEnrichingConsumer) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeLogs enriches logs with SID resolution before passing to the next consumer
func (s *sidEnrichingConsumer) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil

	// SID resolution disabled, pass through
}

// Iterate through all log records and enrich them

// enrichLogRecord enriches a single log record with SID resolution
func (s *sidEnrichingConsumer) enrichLogRecord(record plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

// Enrich security.user_id field

// Enrich SID fields in event_data

// enrichSecurityField enriches the security.user_id field
func (s *sidEnrichingConsumer) enrichSecurityField(bodyMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// Resolve SID

// Add resolved fields

// enrichEventDataFields enriches SID fields within event_data
func (s *sidEnrichingConsumer) enrichEventDataFields(bodyMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// Check if event_data has a "data" array (old format)

// Otherwise, treat event_data as a flat map and enrich fields directly

// tryResolveSID checks whether a key/value pair represents a SID field and resolves it.
// Returns the resolved SID or nil if the field is not a SID, is empty, or resolution fails.
func (s *sidEnrichingConsumer) tryResolveSID(key string, value pcommon.Value) *sidcache.ResolvedSID {
	_ = "STUB: not implemented"
	return nil
}

// enrichEventDataArray enriches SID fields in the event_data.data array format
func (s *sidEnrichingConsumer) enrichEventDataArray(dataSlice pcommon.Slice) {
	_ = "STUB: not implemented"
	// Track which SIDs we've seen to add companion fields after the original
	return
}

// First pass: identify SID fields and resolve them

// Second pass: add companion fields for each resolved SID

// Add {field}_Resolved

// Add {field}_Domain

// Add {field}_Account

// Add {field}_Type

// enrichEventDataMap enriches SID fields in a flat event_data map
func (s *sidEnrichingConsumer) enrichEventDataMap(eventDataMap pcommon.Map) {
	_ = "STUB: not implemented"
	// Track SIDs to enrich (we'll add fields after iteration to avoid modifying during range)
	return
}

// Add companion fields for each resolved SID
