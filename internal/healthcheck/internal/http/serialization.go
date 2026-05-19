// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package http // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/http"

import (
	"time"

	"go.opentelemetry.io/collector/component/componentstatus"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"
)

type healthyFunc func(status.Event) bool

func (f healthyFunc) isHealthy(ev status.Event) bool { _ = "STUB: not implemented"; return false }

type serializationOptions struct {
	includeStartTime  bool
	startTimestamp    *time.Time
	healthyFunc       healthyFunc
	includeAttributes bool
}

type serializableStatus struct {
	StartTimestamp *time.Time `json:"start_time,omitempty"`
	*SerializableEvent
	ComponentStatuses map[string]*serializableStatus `json:"components,omitempty"`
}

// SerializableEvent is exported for json.Unmarshal
type SerializableEvent struct {
	Healthy      bool           `json:"healthy"`
	StatusString string         `json:"status"`
	Error        string         `json:"error,omitempty"`
	Timestamp    time.Time      `json:"status_time"`
	Attributes   map[string]any `json:"attributes"`
}

var stringToStatusMap = map[string]componentstatus.Status{
	"StatusNone":             componentstatus.StatusNone,
	"StatusStarting":         componentstatus.StatusStarting,
	"StatusOK":               componentstatus.StatusOK,
	"StatusRecoverableError": componentstatus.StatusRecoverableError,
	"StatusPermanentError":   componentstatus.StatusPermanentError,
	"StatusFatalError":       componentstatus.StatusFatalError,
	"StatusStopping":         componentstatus.StatusStopping,
	"StatusStopped":          componentstatus.StatusStopped,
}

func (ev *SerializableEvent) Status() componentstatus.Status {
	_ = "STUB: not implemented"
	return *new(componentstatus.Status)
}

func toSerializableEvent(ev status.Event, isHealthy, includeAttributes bool) *SerializableEvent {
	_ = "STUB: not implemented"
	return nil
}

func toSerializableStatus(
	st *status.AggregateStatus,
	opts *serializationOptions,
) *serializableStatus {
	_ = "STUB: not implemented"
	return nil
}
