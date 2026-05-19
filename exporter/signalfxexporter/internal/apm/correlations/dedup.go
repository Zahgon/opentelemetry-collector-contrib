// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Originally copied from https://github.com/signalfx/signalfx-agent/blob/fbc24b0fdd3884bd0bbfbd69fe3c83f49d4c0b77/pkg/apm/correlations/dedup.go

package correlations // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/correlations"

import (
	"container/list"
)

// deduplicator deduplicates requests and cancels pending conflicting requests and deduplicates
// this is not threadsafe
type deduplicator struct {
	// maps for deduplicating requests
	maxSize           int
	pendingCreates    *list.List
	pendingCreateKeys map[Correlation]*list.Element
	pendingDeletes    *list.List
	pendingDeleteKeys map[Correlation]*list.Element
}

func (d *deduplicator) purgeCreates() { _ = "STUB: not implemented"; return }

func (d *deduplicator) purgeDeletes() { _ = "STUB: not implemented"; return }

func (d *deduplicator) purge() { _ = "STUB: not implemented"; return }

func (d *deduplicator) evictPendingDelete() { _ = "STUB: not implemented"; return }

func (d *deduplicator) evictPendingCreate() { _ = "STUB: not implemented"; return }

func (d *deduplicator) dedupCorrelate(r *request) bool {
	_ = "STUB: not implemented"
	// look for duplicate pending creates
	return false
}

// return true if there is a context for the key and the context has not expired

// make room if necessary

// insert the request into the pendingCreates

// cancel any pending delete operations

func (d *deduplicator) dedupDelete(r *request) bool {
	_ = "STUB: not implemented"
	// look for duplicate pending creates
	return false
}

// return true if there is a context for the key and the context has not expired

// make room if necessary

// insert the request into the pendingDeletes

// cancel any pending create operations

// isDup returns true if the request is a duplicate
func (d *deduplicator) isDup(r *request) (isDup bool) { _ = "STUB: not implemented"; return false }

// newDeduplicator returns a new instance
func newDeduplicator(size int) *deduplicator { _ = "STUB: not implemented"; return nil }
