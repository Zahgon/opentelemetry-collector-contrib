// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package status // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"

import (
	"container/list"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// Note: this interface had to be introduced because we need to be able to rewrite the
// timestamps of some events during aggregation. The implementation in core doesn't currently
// allow this, but this interface provides a workaround.
type Event interface {
	Status() componentstatus.Status
	Err() error
	Timestamp() time.Time
	Attributes() pcommon.Map
}

// Scope refers to a part of an AggregateStatus. The zero-value, aka ScopeAll,
// refers to the entire AggregateStatus. ScopeExtensions refers to the extensions
// subtree, and any other value refers to a pipeline subtree.
type Scope string

const (
	ScopeAll        Scope  = ""
	ScopeExtensions Scope  = "extensions"
	pipelinePrefix  string = "pipeline:"
)

func (s Scope) toKey() string { _ = "STUB: not implemented"; return "" }

type Verbosity bool

const (
	Verbose Verbosity = true
	Concise           = false
)

// AggregateStatus contains a map of child AggregateStatuses and an embedded Event.
// It can be used to represent a single, top-level status when the ComponentStatusMap
// is empty, or a nested structure when map is non-empty.
type AggregateStatus struct {
	Event

	ComponentStatusMap map[string]*AggregateStatus
}

func (a *AggregateStatus) clone(verbosity Verbosity) *AggregateStatus {
	_ = "STUB: not implemented"
	return nil
}

type subscription struct {
	statusCh  chan *AggregateStatus
	verbosity Verbosity
}

// UnsubscribeFunc is a function used to unsubscribe from a stream.
type UnsubscribeFunc func()

// Aggregator records individual status events for components and aggregates statuses for the
// pipelines they belong to and the collector overall.
type Aggregator struct {
	// mu protects aggregateStatus and subscriptions from concurrent modification
	mu              sync.RWMutex
	aggregateStatus *AggregateStatus
	subscriptions   map[string]*list.List
	aggregationFunc aggregationFunc
}

// NewAggregator returns a *status.Aggregator.
func NewAggregator(errPriority ErrorPriority) *Aggregator { _ = "STUB: not implemented"; return nil }

// AggregateStatus returns an *AggregateStatus for the given scope. The scope can be the collector
// overall (ScopeAll), extensions (ScopeExtensions), or a pipeline by name. Detail specifies whether
// or not subtrees should be returned with the *AggregateStatus. The boolean return value indicates
// whether or not the scope was found.
func (a *Aggregator) AggregateStatus(scope Scope, verbosity Verbosity) (*AggregateStatus, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// RecordStatus stores and aggregates a StatusEvent for the given component instance.
func (a *Aggregator) RecordStatus(source *componentstatus.InstanceID, event *componentstatus.Event) {
	_ = "STUB: not implemented"
	return
}

// extensions are treated as a pseudo-pipeline

func (a *Aggregator) updateStatus(pipelineScope Scope, source *componentstatus.InstanceID, event *componentstatus.Event) {
	_ = "STUB: not implemented"
	return
}

// Subscribe allows you to subscribe to a stream of events for the given scope. The scope can be
// the collector overall (ScopeAll), extensions (ScopeExtensions), or a pipeline name.
// It is possible to subscribe to a pipeline that has not yet reported. An initial nil
// will be sent on the channel and events will start streaming if and when it starts reporting.
// A `Verbose` verbosity specifies that subtrees should be returned with the *AggregateStatus.
// To unsubscribe, call the returned UnsubscribeFunc.
func (a *Aggregator) Subscribe(scope Scope, verbosity Verbosity) (<-chan *AggregateStatus, UnsubscribeFunc) {
	_ = "STUB: not implemented"
	return nil, *new(UnsubscribeFunc)
}

// Close terminates all existing subscriptions.
func (a *Aggregator) Close() { _ = "STUB: not implemented"; return }

func (a *Aggregator) notifySubscribers(scope Scope, status *AggregateStatus) {
	_ = "STUB: not implemented"
	return
}

// clear unread events
