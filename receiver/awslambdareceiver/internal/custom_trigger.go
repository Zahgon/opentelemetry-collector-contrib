// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awslambdareceiver/internal"

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.uber.org/zap"
)

// Types

// CustomTriggerHandler defines the generic contract for custom event handling
// Custom events are expected to be consumed by configured event handlers.
// Contract includes,
//   - Implements Iterator to iterate through available custom content
//   - IsDryRun to indicate if trigger is in dry-run mode
//   - PostProcess to perform any post-processing after content is consumed
type CustomTriggerHandler interface {
	Iterator[[]byte]
	IsDryRun() bool
	PostProcess(context.Context)
}

// Iterator expose a generic contract to iterate and obtain next available content
type Iterator[T any] interface {
	HasNext(context.Context) bool
	GetNext(context.Context) (T, error)
	Error() error
}

// errorReplayTrigger defines error replay custom trigger structure
type errorReplayTrigger struct {
	Config ReplayTriggerConfig `json:"replayFailedEvents"`
}

// newDefaultErrorTrigger generate a trigger event with default settings
func newDefaultErrorTrigger() errorReplayTrigger {
	_ = "STUB: not implemented"
	return *new(errorReplayTrigger)
}

// ReplayTriggerConfig defines the configuration for error replay trigger
type ReplayTriggerConfig struct {
	Dryrun          bool `json:"dryrun,omitempty"`
	RemoveOnSuccess bool `json:"removeOnSuccess,omitempty"`
}

// errorEvent defines the structure of the error event stored in S3. Here, only required section is defined
// See https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations
type errorEvent struct {
	RequestPayload json.RawMessage `json:"requestPayload"`
}

// Implementations

// ErrorReplayTriggerHandler implements CustomTriggerHandler for error replaying.
type ErrorReplayTriggerHandler struct {
	trigger    errorReplayTrigger
	bucket     string
	s3Service  S3Service
	s3Iterator Iterator[*types.Object]

	logger *zap.Logger

	currentKey string
}

func NewErrorReplayTriggerHandler(log *zap.Logger, event []byte, bucketName string, s3Service S3Service) (CustomTriggerHandler, error) {
	_ = "STUB: not implemented"
	return *new(CustomTriggerHandler), nil
}

func (m *ErrorReplayTriggerHandler) IsDryRun() bool { _ = "STUB: not implemented"; return false }

func (m *ErrorReplayTriggerHandler) HasNext(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// GetNext implementation retrieves the S3 content from the error S3 bucket.
func (m *ErrorReplayTriggerHandler) GetNext(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ErrorReplayTriggerHandler) Error() error { _ = "STUB: not implemented"; return nil }

func (m *ErrorReplayTriggerHandler) PostProcess(ctx context.Context) {
	_ = "STUB: not implemented"
	return

	// nothing to do
}

// skip if dry-run mode

// s3ListIterator implements Iterator and isolates S3 object listing for consumers
type s3ListIterator struct {
	s3Service S3Service
	bucket    string
	prefix    string

	latest     *s3.ListObjectsV2Output
	currentObj *types.Object
	index      int
	done       bool

	err error
}

func newS3ListIterator(s3Service S3Service, bucket, prefix string) Iterator[*types.Object] {
	_ = "STUB: not implemented"
	return nil
}

func (i *s3ListIterator) HasNext(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// check for initial fetch

// check for limits and fetch next if available

// set current, advance and return

func (i *s3ListIterator) GetNext(_ context.Context) (*types.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *s3ListIterator) Error() error { _ = "STUB: not implemented"; return nil }
