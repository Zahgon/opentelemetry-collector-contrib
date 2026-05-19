// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ackextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension"

import (
	"context"
	"sync/atomic"

	lru "github.com/hashicorp/golang-lru/v2"
	"go.opentelemetry.io/collector/component"
)

// inMemoryAckExtension is the in-memory implementation of the AckExtension
// When MaxNumPartition is reached, the acks associated with the least recently used partition are evicted.
// When MaxNumPendingAcksPerPartition is reached, the least recently used ack is evicted
type inMemoryAckExtension struct {
	partitionMap                  *lru.Cache[string, *ackPartition]
	maxNumPendingAcksPerPartition uint64
}

func newInMemoryAckExtension(conf *Config) *inMemoryAckExtension {
	_ = "STUB: not implemented"
	return nil
}

type ackPartition struct {
	id     atomic.Uint64
	ackMap *lru.Cache[uint64, bool]
}

func newAckPartition(maxPendingAcks uint64) *ackPartition { _ = "STUB: not implemented"; return nil }

func (as *ackPartition) nextAck() uint64 { _ = "STUB: not implemented"; return 0 }

func (as *ackPartition) ack(key uint64) { _ = "STUB: not implemented"; return }

func (as *ackPartition) computeAcks(ackIDs []uint64) map[uint64]bool {
	_ = "STUB: not implemented"
	return nil
}

// Start of inMemoryAckExtension does nothing and returns nil
func (*inMemoryAckExtension) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"

	// Shutdown of inMemoryAckExtension does nothing and returns nil
	return nil
}

func (*inMemoryAckExtension) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// ProcessEvent marks the beginning of processing an event. It generates an ack ID for the associated partition ID.
	return nil
}

func (i *inMemoryAckExtension) ProcessEvent(partitionID string) (ackID uint64) {
	_ = "STUB: not implemented"
	return 0
}

// Ack acknowledges an event has been processed.
func (i *inMemoryAckExtension) Ack(partitionID string, ackID uint64) {
	_ = "STUB: not implemented"
	return
}

// QueryAcks checks the statuses of given ackIDs for a partition.
// ackIDs that are not generated from ProcessEvent or have been removed as a result of previous calls to QueryAcks will return false.
func (i *inMemoryAckExtension) QueryAcks(partitionID string, ackIDs []uint64) map[uint64]bool {
	_ = "STUB: not implemented"
	return nil
}
