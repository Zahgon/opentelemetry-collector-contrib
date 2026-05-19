// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package experimentalmetricmetadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// See entity event design document:
// https://docs.google.com/document/d/1Tg18sIck3Nakxtd3TFFcIjrmRO_0GLMdHXylVqBQmJA/edit#heading=h.pokdp8i2dmxy

const (
	semconvOtelEntityEventName    = "otel.entity.event.type"
	semconvEventEntityEventState  = "entity_state"
	semconvEventEntityEventDelete = "entity_delete"

	semconvOtelEntityID         = "otel.entity.id"
	semconvOtelEntityType       = "otel.entity.type"
	semconvOtelEntityInterval   = "otel.entity.interval"
	semconvOtelEntityAttributes = "otel.entity.attributes"

	SemconvOtelEntityEventAsScope = "otel.entity.event_as_log"
)

// EntityEventsSlice is a slice of EntityEvent.
type EntityEventsSlice struct {
	orig plog.LogRecordSlice
}

// NewEntityEventsSlice creates an empty EntityEventsSlice.
func NewEntityEventsSlice() EntityEventsSlice {
	_ = "STUB: not implemented"
	return *new(EntityEventsSlice)
}

// NewEntityEventsSliceFromLogs creates an EntityEventsSlice from a plog.LogRecordSlice.
func NewEntityEventsSliceFromLogs(logs plog.LogRecordSlice) EntityEventsSlice {
	_ = "STUB: not implemented"
	return *new(EntityEventsSlice)
}

// AppendEmpty will append to the end of the slice an empty EntityEvent.
// It returns the newly added EntityEvent.
func (s EntityEventsSlice) AppendEmpty() EntityEvent {
	_ = "STUB: not implemented"
	return *new(EntityEvent)
}

// Len returns the number of elements in the slice.
func (s EntityEventsSlice) Len() int { _ = "STUB: not implemented"; return 0 }

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
func (s EntityEventsSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

// At returns the element at the given index.
func (s EntityEventsSlice) At(i int) EntityEvent {
	_ = "STUB: not implemented"
	return *new(EntityEvent)
}

// ConvertAndMoveToLogs converts entity events to log representation and moves them
// from this EntityEventsSlice into plog.Logs. This slice becomes empty after this call.
func (s EntityEventsSlice) ConvertAndMoveToLogs() plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// Set the scope marker.

// Move all events. Note that this remove all

// EntityEvent is an entity event.
type EntityEvent struct {
	orig plog.LogRecord
}

// Timestamp of the event.
func (e EntityEvent) Timestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

// SetTimestamp sets the event timestamp.
func (e EntityEvent) SetTimestamp(timestamp pcommon.Timestamp) { _ = "STUB: not implemented"; return }

// ID of the entity.
func (e EntityEvent) ID() pcommon.Map { _ = "STUB: not implemented"; return *new(pcommon.Map) }

// SetEntityState makes this an EntityStateDetails event.
func (e EntityEvent) SetEntityState() EntityStateDetails {
	_ = "STUB: not implemented"
	return *new(EntityStateDetails)
}

// EntityStateDetails returns the entity state details of this event.
func (e EntityEvent) EntityStateDetails() EntityStateDetails {
	_ = "STUB: not implemented"
	return *new(EntityStateDetails)
}

// SetEntityDelete makes this an EntityDeleteDetails event.
func (e EntityEvent) SetEntityDelete() EntityDeleteDetails {
	_ = "STUB: not implemented"
	return *new(EntityDeleteDetails)
}

// EntityDeleteDetails return the entity delete details of this event.
func (e EntityEvent) EntityDeleteDetails() EntityDeleteDetails {
	_ = "STUB: not implemented"
	return *new(EntityDeleteDetails)
}

// EventType is the type of the entity event.
type EventType int

const (
	// EventTypeNone indicates an invalid or unknown event type.
	EventTypeNone EventType = iota
	// EventTypeState is the "entity state" event.
	EventTypeState
	// EventTypeDelete is the "entity delete" event.
	EventTypeDelete
)

// EventType returns the type of the event.
func (e EntityEvent) EventType() EventType { _ = "STUB: not implemented"; return *new(EventType) }

// EntityStateDetails represents the details of an EntityState event.
type EntityStateDetails struct {
	orig plog.LogRecord
}

// Attributes returns the attributes of the entity.
func (s EntityStateDetails) Attributes() pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// EntityType returns the type of the entity.
func (s EntityStateDetails) EntityType() string { _ = "STUB: not implemented"; return "" }

// SetEntityType sets the type of the entity.
func (s EntityStateDetails) SetEntityType(t string) { _ = "STUB: not implemented"; return }

// SetInterval sets the reporting period
// i.e. how frequently the information about this entity is reported via EntityState events even if the entity does not change.
func (s EntityStateDetails) SetInterval(t time.Duration) { _ = "STUB: not implemented"; return }

// Interval returns the reporting period
func (s EntityStateDetails) Interval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// EntityDeleteDetails represents the details of an EntityDelete event.
type EntityDeleteDetails struct {
	orig plog.LogRecord
}

// EntityType returns the type of the entity.
// TODO: Move the entity type methods to EntityEvent as they are needed for both EntityState and EntityDelete events.
func (d EntityDeleteDetails) EntityType() string { _ = "STUB: not implemented"; return "" }

// SetEntityType sets the type of the entity.
func (d EntityDeleteDetails) SetEntityType(t string) { _ = "STUB: not implemented"; return }
