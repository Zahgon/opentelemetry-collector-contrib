// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetry // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

const (
	envAWSHostname     = "AWS_HOSTNAME"
	envAWSInstanceID   = "AWS_INSTANCE_ID"
	metadataHostname   = "hostname"
	metadataInstanceID = "instance-id"

	defaultQueueSize = 30
	defaultBatchSize = 10
	defaultInterval  = time.Minute
)

// Sender wraps a Recorder and periodically sends the records.
type Sender interface {
	Recorder
	// Start send loop.
	Start(ctx context.Context)
	// Stop send loop.
	Stop()
}

type telemetrySender struct {
	// Recorder is the recorder wrapped by the sender.
	Recorder

	// logger is used to log dropped records.
	logger *zap.Logger
	// client is used to send the records.
	client awsxray.XRayClient

	resourceARN string
	instanceID  string
	hostname    string
	// interval is the amount of time between record rotation and sending attempts.
	interval time.Duration
	// queueSize is the capacity of the queue.
	queueSize int
	// batchSize is the max number of records sent in one request.
	batchSize int

	// queue is used to keep records that failed to send for retry during
	// the next period.
	queue []types.TelemetryRecord

	startOnce sync.Once
	stopWait  sync.WaitGroup
	stopOnce  sync.Once
	// stopCh is the channel used to stop the loop.
	stopCh chan struct{}
}

type Option interface {
	apply(ts *telemetrySender)
}

type optionFunc func(ts *telemetrySender)

func (o optionFunc) apply(ts *telemetrySender) { _ = "STUB: not implemented"; return }

func WithResourceARN(resourceARN string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithInstanceID(instanceID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHostname(hostname string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogger(logger *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithInterval(interval time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithQueueSize(queueSize int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBatchSize(batchSize int) Option { _ = "STUB: not implemented"; return *new(Option) }

type metadataProvider interface {
	get(ctx context.Context) string
}

func getMetadata(ctx context.Context, providers ...metadataProvider) string {
	_ = "STUB: not implemented"
	return ""
}

type simpleMetadataProvider struct {
	metadata string
}

func (p simpleMetadataProvider) get(_ context.Context) string { _ = "STUB: not implemented"; return "" }

type envMetadataProvider struct {
	envKey string
}

func (p envMetadataProvider) get(_ context.Context) string { _ = "STUB: not implemented"; return "" }

type ec2MetadataProvider struct {
	client      *imds.Client
	metadataKey string
}

func (p ec2MetadataProvider) get(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// ToOptions returns the metadata options if enabled by the config.
func ToOptions(ctx context.Context, cfg Config, awsConfig aws.Config, settings *awsutil.AWSSessionSettings) []Option {
	_ = "STUB: not implemented"
	return nil
}

// NewSender creates a new Sender with a default interval and queue size.
func NewSender(client awsxray.XRayClient, opts ...Option) Sender {
	_ = "STUB: not implemented"
	return *new(Sender)
}

func newSender(client awsxray.XRayClient, opts ...Option) *telemetrySender {
	_ = "STUB: not implemented"
	return nil
}

// Start starts the loop to send the records.
func (ts *telemetrySender) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

// Stop closes the stopCh channel to stop the loop.
func (ts *telemetrySender) Stop() { _ = "STUB: not implemented"; return }

// run sends the queued records once a minute if telemetry data was updated.
func (ts *telemetrySender) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// enqueue the record. If queue is full, drop the head of the queue and add.
func (ts *telemetrySender) enqueue(record types.TelemetryRecord) { _ = "STUB: not implemented"; return }

// send the records in the queue in batches. Updates the queue.
func (ts *telemetrySender) send(ctx context.Context) { _ = "STUB: not implemented"; return }
