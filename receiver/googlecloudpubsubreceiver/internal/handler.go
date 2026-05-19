// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver/internal"

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver/internal/metadata"
)

type StreamHandler struct {
	stream      pubsubpb.Subscriber_StreamingPullClient
	pushMessage func(ctx context.Context, message *pubsubpb.ReceivedMessage) error
	acks        []string
	mutex       sync.Mutex
	client      SubscriberClient

	clientID     string
	subscription string

	cancel context.CancelFunc
	// wait group for the send/receive function
	streamWaitGroup sync.WaitGroup
	// wait group for the handler
	handlerWaitGroup sync.WaitGroup
	settings         receiver.Settings
	telemetryBuilder *metadata.TelemetryBuilder

	// flow control settings, like max durations, counts and triggers
	flowControlConfig *FlowControlConfig

	isRunning    atomic.Bool
	retryAttempt int
}

// ack adds the ackID to the list of message to be acknowledged asynchronously
func (handler *StreamHandler) ack(ackID string) { _ = "STUB: not implemented"; return }

func NewHandler(
	ctx context.Context,
	settings receiver.Settings,
	telemetryBuilder *metadata.TelemetryBuilder,
	client SubscriberClient,
	clientID string,
	subscription string,
	config *FlowControlConfig,
	callback func(ctx context.Context, message *pubsubpb.ReceivedMessage) error,
) (*StreamHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initStream creates a new streaming pull stream. When the previous stream was closed, the
// pending acknowledge messages will be acknowledged at stream re-creation.
func (handler *StreamHandler) initStream(ctx context.Context) error {
	_ = "STUB: not implemented"

	// Create a stream, but with the receivers context as we don't want to cancel and ongoing operation
	return nil
}

// RecoverableStream starts the Pub/Sub stream loop and recovers it if it fails
func (handler *StreamHandler) RecoverableStream(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (handler *StreamHandler) recoverableStream(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Create a new cancelable context for the handler, so we can recover the stream

func (handler *StreamHandler) CancelNow() { _ = "STUB: not implemented"; return }

func (handler *StreamHandler) Wait() { _ = "STUB: not implemented"; return }

// acknowledgeMessages will acknowledge the messages, and only clear the outstanding messages when the
// acknowledgement is send successfully
func (handler *StreamHandler) acknowledgeMessages() error { _ = "STUB: not implemented"; return nil }

// requestStream waits for triggers to acknowledge messages that have been processed by the collector. If
// a stream got restarted, the messages that still needed to be acknowledged are acknowledged at the start
// of the new stream, so we don't need to start with an acknowledgeMessages.
func (handler *StreamHandler) requestStream(ctx context.Context, cancel context.CancelFunc) {
	_ = "STUB: not implemented"
	return
}

// whatever happens, we need to acknowledge the messages

// if the context is canceled, we break the loop

func (handler *StreamHandler) responseStream(ctx context.Context, cancel context.CancelFunc) {
	_ = "STUB: not implemented"
	return
}

// block until the next message or timeout expires

// handle all the messages in the response, could be one or more

// When sending a message though the pipeline fails, we ignore the error. We'll let Pubsub
// handle the flow control.

// Canceling the loop, collector is probably stopping

// exponentialBackoff will backoff exponentially with a maximum of 2 minutes
func exponentialBackoff(retryAttempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
