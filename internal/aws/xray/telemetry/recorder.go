// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetry // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"

import (
	"sync/atomic"

	"github.com/aws/aws-sdk-go-v2/service/xray/types"
)

type Recorder interface {
	// Rotate the current record by swapping it out with a new one. Returns
	// the rotated record.
	Rotate() types.TelemetryRecord
	// HasRecording indicates whether any of the record functions were called
	// with the current record.
	HasRecording() bool
	// RecordSegmentsReceived adds the count to the current record.
	RecordSegmentsReceived(count int)
	// RecordSegmentsSent adds the count to the current record.
	RecordSegmentsSent(count int)
	// RecordSegmentsSpillover adds the count to the current record.
	RecordSegmentsSpillover(count int)
	// RecordSegmentsRejected adds the count to the current record.
	RecordSegmentsRejected(count int)
	// RecordConnectionError categorizes the error and increments the count by one
	// for the current record.
	RecordConnectionError(err error)
}

type telemetryRecorder struct {
	// record is the pointer to the count metrics for the current period.
	record types.TelemetryRecord
	// hasRecording is set to true when any count is updated. Indicates
	// that telemetry data is available.
	hasRecording *atomic.Bool
}

// NewRecorder creates a new Recorder with a default interval and queue size.
func NewRecorder() Recorder { _ = "STUB: not implemented"; return *new(Recorder) }

// NewRecord creates a new xray.TelemetryRecord with all of its fields initialized
// and set to 0.
func NewRecord() types.TelemetryRecord {
	_ = "STUB: not implemented"
	return *new(types.TelemetryRecord)
}

func (tr *telemetryRecorder) HasRecording() bool { _ = "STUB: not implemented"; return false }

// Rotate the current record and swaps it out with a new record.
// Sets the timestamp and returns the old record.
func (tr *telemetryRecorder) Rotate() types.TelemetryRecord {
	_ = "STUB: not implemented"
	return *new(types.TelemetryRecord)
}

func (tr *telemetryRecorder) RecordSegmentsReceived(count int) { _ = "STUB: not implemented"; return }

func (tr *telemetryRecorder) RecordSegmentsSent(count int) { _ = "STUB: not implemented"; return }

func (tr *telemetryRecorder) RecordSegmentsSpillover(count int) { _ = "STUB: not implemented"; return }

func (tr *telemetryRecorder) RecordSegmentsRejected(count int) { _ = "STUB: not implemented"; return }

func (tr *telemetryRecorder) RecordConnectionError(err error) { _ = "STUB: not implemented"; return }
