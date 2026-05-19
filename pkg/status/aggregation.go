// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package status // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"

import (
	"time"

	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// statusEvent contains a status and timestamp, and can contain an error. Note:
// this is duplicated from core because we need to be able to "rewrite" the
// timestamps of some events during aggregation.
type statusEvent struct {
	status     componentstatus.Status
	err        error
	timestamp  time.Time
	attributes pcommon.Map
}

var _ Event = (*statusEvent)(nil)

// Status returns the Status (enum) associated with the StatusEvent
func (ev *statusEvent) Status() componentstatus.Status {
	_ = "STUB: not implemented"

	// Err returns the error associated with the StatusEvent.
	return *new(componentstatus.Status)
}

func (ev *statusEvent) Err() error {
	_ = "STUB: not implemented"

	// Timestamp returns the timestamp associated with the StatusEvent
	return nil
}

func (ev *statusEvent) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (ev *statusEvent) Attributes() pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

type ErrorPriority int

const (
	PriorityPermanent ErrorPriority = iota
	PriorityRecoverable
)

type aggregationFunc func(*AggregateStatus) Event

// The purpose of aggregation is to ensure that the most relevant status bubbles
// upwards in the aggregate status. This aggregation func prioritizes lifecycle
// events (including FatalError) over PermanentError and RecoverableError
// events. The priority argument determines the priority of PermanentError
// events vs RecoverableError events. Lifecycle events will have the timestamp
// of the most recent event and error events will have the timestamp of the
// first occurrence. We use the first occurrence of an error event as this marks
// the beginning of a possible failure. This is important for two reasons:
// recovery duration and causality. We expect a RecoverableError to recover
// before the RecoveryDuration elapses. We need to use the earliest timestamp so
// that a later RecoverableError does not shadow an earlier event in the
// aggregate status. Additionally, this makes sense in the case where a
// RecoverableError in one component cascades to other components; the earliest
// error event is likely to be correlated with the cause. For non-error stauses
// we use the latest event as it represents the last time a successful status was
// reported.
func newAggregationFunc(priority ErrorPriority) aggregationFunc {
	_ = "STUB: not implemented"
	return *new(aggregationFunc)
}

// All statuses are the same. Note, this will handle StatusOK and StatusStopped as these two
// cases require all components be in the same state.

// Handle mixed status cases

// Use earliest to mark beginning of a failure

// Use most recent for last successful status

// the error status will be the first matching event

// the aggregate status matches an existing event

// the aggregate status requires a synthetic event
