// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package arrow // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/otelarrowexporter/internal/arrow"

import (
	"context"
	"math/rand/v2"
	"time"
)

// bestOfNPrioritizer is a prioritizer that selects a less-loaded stream to write.
// https://smallrye.io/smallrye-stork/1.1.1/load-balancer/power-of-two-choices/
type bestOfNPrioritizer struct {
	doneCancel

	// input from the pipeline, as processed data with headers and
	// a return channel for the result.  This channel is never
	// closed and is buffered.  At shutdown, items of telemetry can
	// be left in this channel, but users are expected to complete
	// their requests before calling shutdown (and the collector's
	// graph package ensures this).
	input chan writeItem

	// state tracks the work being handled by all streams.
	state []*streamWorkState

	// numChoices is the number of streams to consider in each decision.
	numChoices int

	// loadFunc is the load function.
	loadFunc loadFunc
}

type loadFunc func(*streamWorkState) float64

type streamSorter struct {
	work *streamWorkState
	load float64
}

var _ streamPrioritizer = &bestOfNPrioritizer{}

func newBestOfNPrioritizer(dc doneCancel, numChoices, numStreams int, lf loadFunc, maxLifetime time.Duration) (*bestOfNPrioritizer, []*streamWorkState) {
	_ = "STUB: not implemented"
	return nil,

		// Limit numChoices to the number of streams.
		nil
}

// TODO It's not clear if/when the prioritizer can
// become a bottleneck.

func (lp *bestOfNPrioritizer) downgrade(ctx context.Context) { _ = "STUB: not implemented"; return }

func (lp *bestOfNPrioritizer) sendOne(item writeItem, rnd *rand.Rand, tmp []streamSorter) {
	_ = "STUB: not implemented"
	return
}

// All other cases: signal restart.

func (lp *bestOfNPrioritizer) run() { _ = "STUB: not implemented"; return }

// sendAndWait implements streamWriter
func (lp *bestOfNPrioritizer) sendAndWait(ctx context.Context, errCh <-chan error, wri writeItem) error {
	_ = "STUB: not implemented"
	return nil
}

func (lp *bestOfNPrioritizer) nextWriter() streamWriter {
	_ = "STUB: not implemented"
	return *

	// In case of downgrade, return nil to return into a
	// non-Arrow code path.
	new(streamWriter)
}

// Fall through to sendAndWait().

func (lp *bestOfNPrioritizer) streamFor(_ writeItem, rnd *rand.Rand, tmp []streamSorter) *streamWorkState {
	_ = "STUB: not implemented"
	// Place all streams into the temporary slice.
	return nil
}

// Select numChoices at random by shifting the selection into the start
// of the temporary slice.

// TODO: skip channels w/ a pending item (maybe)
