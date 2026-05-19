// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package failoverconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector"

import (
	"errors"

	"go.opentelemetry.io/collector/pipeline"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector/internal/state"
)

var (
	errNoValidPipeline = errors.New("All provided pipelines return errors")
	errConsumer        = errors.New("Error registering consumer")
)

type consumerProvider[C any] func(...pipeline.ID) (C, error)

// baseFailoverRouter provides the common infrastructure for failover routing
type baseFailoverRouter[C any] struct {
	cfg       *Config
	pS        *state.PipelineSelector
	consumers []C

	errTryLock  *state.TryLock
	notifyRetry chan struct{}
	done        chan struct{}
}

// getCurrentConsumer returns the consumer for the current healthy level
func (f *baseFailoverRouter[C]) getCurrentConsumer() (C, int) {
	_ = "STUB: not implemented"
	return *new(C), 0
}

// getConsumerAtIndex returns the consumer at a specific index
func (f *baseFailoverRouter[C]) getConsumerAtIndex(idx int) C {
	_ = "STUB: not implemented"
	return *

	// reportConsumerError ensures only one consumer is reporting an error at a time to avoid multiple failovers
	new(C)
}

func (f *baseFailoverRouter[C]) reportConsumerError(idx int) { _ = "STUB: not implemented"; return }

func (f *baseFailoverRouter[C]) Shutdown() { _ = "STUB: not implemented"; return }

func newBaseFailoverRouter[C any](provider consumerProvider[C], cfg *Config) (*baseFailoverRouter[C], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For Testing
func (f *baseFailoverRouter[C]) ModifyConsumerAtIndex(idx int, c C) {
	_ = "STUB: not implemented"
	return
}

func (f *baseFailoverRouter[C]) TestGetCurrentConsumerIndex() int {
	_ = "STUB: not implemented"
	return 0
}

func (f *baseFailoverRouter[C]) TestSetStableConsumerIndex(idx int) {
	_ = "STUB: not implemented"
	return
}

func (f *baseFailoverRouter[C]) TestGetConsumerAtIndex(idx int) C {
	_ = "STUB: not implemented"
	return *new(C)
}
