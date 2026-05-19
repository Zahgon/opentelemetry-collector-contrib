// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package state // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector/internal/state"

import (
	"sync"
)

type PipelineSelector struct {
	currentPipeline   int
	constants         PSConstants
	lock              sync.RWMutex
	retryEnabledToken chan struct{}
	retryChan         chan<- struct{}

	retryCancel CancelManager
	done        chan struct{}
}

// HandleError is called when an error is returned on a healthy pipeline
func (p *PipelineSelector) HandleError(idx int) { _ = "STUB: not implemented"; return }

// NextStableLevel increments the level to the next in the priority list
func (p *PipelineSelector) NextStableLevel() { _ = "STUB: not implemented"; return }

// TryEnableRetry checks if a retry is already in effect and if not starts the retry goroutine
func (p *PipelineSelector) TryEnableRetry() { _ = "STUB: not implemented"; return }

// LaunchRetry invokes the goroutine responsible for notifying the failover component to retry
func (p *PipelineSelector) LaunchRetry() { _ = "STUB: not implemented"; return }

// returnRetryToken returns the token back to the buffered channel allowing the next retry function to consume the token
func (p *PipelineSelector) returnRetryToken() { _ = "STUB: not implemented"; return }

// CurrentLevel returns the current healthy pipeline level
func (p *PipelineSelector) CurrentPipeline() int { _ = "STUB: not implemented"; return 0 }

// ResetHealthyPipeline resets a pipeline level that was successfully retries back to healthy/active
func (p *PipelineSelector) ResetHealthyPipeline(pipelineIndex int) {
	_ = "STUB: not implemented"
	return
}

func NewPipelineSelector(retryChan chan<- struct{}, done chan struct{}, consts PSConstants) *PipelineSelector {
	_ = "STUB: not implemented"
	return nil
}

// For Testing
func (p *PipelineSelector) TestSetCurrentPipeline(idx int) { _ = "STUB: not implemented"; return }
