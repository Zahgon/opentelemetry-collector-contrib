// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwlogs // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs"

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"go.uber.org/zap"
)

const (
	// http://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/cloudwatch_limits_cwl.html
	// In truncation logic, it assuming this constant value is larger than perEventHeaderBytes + len(truncatedSuffix)
	defaultMaxEventPayloadBytes = 1024 * 256 // 256KB
	// http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html
	maxRequestEventCount   = 10000
	perEventHeaderBytes    = 26
	maxRequestPayloadBytes = 1024 * 1024 * 1

	truncatedSuffix = "[Truncated...]"

	eventTimestampLimitInPast  = 14 * 24 * time.Hour // None of the log events in the batch can be older than 14 days
	evenTimestampLimitInFuture = -2 * time.Hour      // None of the log events in the batch can be more than 2 hours in the future.
)

var maxEventPayloadBytes = defaultMaxEventPayloadBytes

// Event struct to present a log event.
type Event struct {
	InputLogEvent types.InputLogEvent
	// The time which log generated.
	GeneratedTime time.Time
	// Identify what is the stream of destination of this event
	StreamKey
}

// NewEvent creates a new log event
// logType will be propagated to LogEventBatch and used by logPusher to determine which client to call PutLogEvent
func NewEvent(timestampMs int64, message string) *Event { _ = "STUB: not implemented"; return nil }

// Uniquely identify a cloudwatch logs stream
type StreamKey struct {
	LogGroupName  string
	LogStreamName string
}

func (logEvent *Event) Validate(logger *zap.Logger) error { _ = "STUB: not implemented"; return nil }

// http://docs.aws.amazon.com/goto/SdkForGoV1/logs-2014-03-28/PutLogEvents
// * None of the log events in the batch can be more than 2 hours in the
// future.
// * None of the log events in the batch can be older than 14 days or the
// retention period of the log group.

// Calculate the log event payload bytes.
func (logEvent *Event) eventPayloadBytes() int { _ = "STUB: not implemented"; return 0 }

// eventBatch struct to present a log event batch
type eventBatch struct {
	putLogEventsInput *cloudwatchlogs.PutLogEventsInput
	// the total bytes already in this log event batch
	byteTotal int
	// min timestamp recorded in this log event batch (ms)
	minTimestampMs int64
	// max timestamp recorded in this log event batch (ms)
	maxTimestampMs int64
}

// Create a new log event batch if needed.
func newEventBatch(key StreamKey) *eventBatch { _ = "STUB: not implemented"; return nil }

func (batch *eventBatch) exceedsLimit(nextByteTotal int) bool {
	_ = "STUB: not implemented"
	return false
}

// isActive checks whether the eventBatch spans more than 24 hours. Returns
// false if the condition does not match, and this batch should not be processed
// any further.
func (batch *eventBatch) isActive(targetTimestampMs *int64) bool {
	_ = "STUB: not implemented"
	// new log event batch
	return false
}

func (batch *eventBatch) append(event *Event) { _ = "STUB: not implemented"; return }

// Sort the log events based on the timestamp.
func (batch *eventBatch) sortLogEvents() { _ = "STUB: not implemented"; return }

type ByTimestamp []types.InputLogEvent

func (inputLogEvents ByTimestamp) Len() int { _ = "STUB: not implemented"; return 0 }

func (inputLogEvents ByTimestamp) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (inputLogEvents ByTimestamp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Pusher is created by log group and log stream
type Pusher interface {
	AddLogEntry(ctx context.Context, logEvent *Event) error
	ForceFlush(ctx context.Context) error
}

// Struct of logPusher implemented Pusher interface.
type logPusher struct {
	mu     sync.Mutex
	logger *zap.Logger
	// log group name of the current logPusher
	logGroupName *string
	// log stream name of the current logPusher
	logStreamName *string

	logEventBatch *eventBatch

	svcStructuredLog Client
	retryCnt         int
}

// NewPusher creates a logPusher instance
func NewPusher(streamKey StreamKey, retryCnt int,
	svcStructuredLog Client, logger *zap.Logger,
) Pusher {
	_ = "STUB: not implemented"
	return *new(Pusher)
}

// Only create a logPusher, but not start the instance.
func newLogPusher(streamKey StreamKey,
	svcStructuredLog Client, logger *zap.Logger,
) *logPusher {
	_ = "STUB: not implemented"
	return nil
}

// AddLogEntry Besides the limit specified by PutLogEvents API, there are some overall limit for the cloudwatchlogs
// listed here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/cloudwatch_limits_cwl.html
//
// Need to pay attention to the below 2 limits:
// Event size 256 KB (maximum). This limit cannot be changed.
// Batch size 1 MB (maximum). This limit cannot be changed.
func (p *logPusher) AddLogEntry(ctx context.Context, logEvent *Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *logPusher) ForceFlush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *logPusher) pushEventBatch(ctx context.Context, req any) error {
	_ = "STUB: not implemented"
	// http://docs.aws.amazon.com/goto/SdkForGoV1/logs-2014-03-28/PutLogEvents
	// The log events in the batch must be in chronological ordered by their
	// timestamp (the time the event occurred, expressed as the number of milliseconds
	// since Jan 1, 1970 00:00:00 UTC).
	return nil
}

func (p *logPusher) addLogEvent(logEvent *Event) *eventBatch { _ = "STUB: not implemented"; return nil }

func (p *logPusher) renewEventBatch() *eventBatch { _ = "STUB: not implemented"; return nil }

// A Pusher that is able to send events to multiple streams.
type multiStreamPusher struct {
	logStreamManager LogStreamManager
	client           Client
	pusherMap        map[StreamKey]Pusher
	logger           *zap.Logger
}

func newMultiStreamPusher(logStreamManager LogStreamManager, client Client, logger *zap.Logger) *multiStreamPusher {
	_ = "STUB: not implemented"
	return nil
}

func (m *multiStreamPusher) AddLogEntry(ctx context.Context, event *Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *multiStreamPusher) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Factory for a Pusher that has capability of sending events to multiple log streams
type MultiStreamPusherFactory interface {
	CreateMultiStreamPusher() Pusher
}

type multiStreamPusherFactory struct {
	logStreamManager LogStreamManager
	logger           *zap.Logger
	client           Client
}

// Creates a new MultiStreamPusherFactory
func NewMultiStreamPusherFactory(logStreamManager LogStreamManager, client Client, logger *zap.Logger) MultiStreamPusherFactory {
	_ = "STUB: not implemented"
	return *new(MultiStreamPusherFactory)
}

// Factory method to create a Pusher that has support to sending events to multiple log streams
func (msf *multiStreamPusherFactory) CreateMultiStreamPusher() Pusher {
	_ = "STUB: not implemented"
	return *new(Pusher)
}

// Manages the creation of streams
type LogStreamManager interface {
	// Initialize a stream so that it can receive logs
	// This will make sure that the stream exists and if it does not exist,
	// It will create one. Implementations of this method MUST be safe for concurrent use.
	InitStream(ctx context.Context, streamKey StreamKey) error
}

type logStreamManager struct {
	logStreamMutex sync.Mutex
	streams        map[StreamKey]bool
	client         Client
}

func NewLogStreamManager(svcStructuredLog Client) LogStreamManager {
	_ = "STUB: not implemented"
	return *new(LogStreamManager)
}

func (lsm *logStreamManager) InitStream(ctx context.Context, streamKey StreamKey) error {
	_ = "STUB: not implemented"
	return nil
}

// does not do anything if stream already exists
