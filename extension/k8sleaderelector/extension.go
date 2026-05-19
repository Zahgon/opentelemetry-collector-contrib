// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sleaderelector // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/k8sleaderelector"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

type (
	StartCallback = func(context.Context)
	StopCallback  = func()
)

// LeaderElection Interface allows the invoker to set the callback functions
// that would be invoked when the leader wins or loss the election.
type LeaderElection interface {
	extension.Extension
	SetCallBackFuncs(StartCallback, StopCallback)
}

type callBackFuncs struct {
	onStartLeading StartCallback
	onStopLeading  StopCallback
}

// leaderElectionExtension is the main struct implementing the extension's behavior.
type leaderElectionExtension struct {
	config        *Config
	client        kubernetes.Interface
	logger        *zap.Logger
	leaseHolderID string
	cancel        context.CancelFunc
	waitGroup     sync.WaitGroup

	callBackFuncs []callBackFuncs

	isLeader bool

	mu sync.Mutex
}

// SetCallBackFuncs set the functions that can be invoked when the leader wins or loss the election
func (lee *leaderElectionExtension) SetCallBackFuncs(onStartLeading StartCallback, onStopLeading StopCallback) {
	_ = "STUB: not implemented"
	// Have a write lock while setting the callbacks.
	return
}

// Immediately invoke the callback since we are already leader

// If the receiver sets a callback function then it would be invoked when the leader wins the election
func (lee *leaderElectionExtension) startedLeading(ctx context.Context) {
	_ = "STUB: not implemented"
	// Have read lock so that we no new callbacks can be added while we are invoking the callbacks.
	return
}

// If the receiver sets a callback function then it would be invoked when the leader loss the election
func (lee *leaderElectionExtension) stoppedLeading() {
	_ = "STUB: not implemented"
	// Have a read lock while stopping the receivers. This would make sure that if we have executed any onStartLeading callbacks
	// after becoming leader, we would execute the onStopLeading callbacks for them as well.
	return
}

// Start begins the extension's processing.
func (lee *leaderElectionExtension) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Create the K8s leader elector

// Leader election loop stops if context is canceled or the leader elector loses the lease.
// The loop allows continued participation in leader election, even if the lease is lost.

// Shutdown ends the extension's processing.
func (lee *leaderElectionExtension) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
