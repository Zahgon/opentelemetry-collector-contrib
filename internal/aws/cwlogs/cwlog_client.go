// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwlogs // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs"

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/smithy-go/middleware"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

const (
	// this is the retry count, the total attempts will be at most retry count + 1.
	defaultRetryCount = 1
)

var containerInsightsRegexPattern = regexp.MustCompile(`^/aws/.*containerinsights/.*/(performance|prometheus)$`)

type cloudWatchClient interface {
	CreateLogGroup(ctx context.Context, params *cloudwatchlogs.CreateLogGroupInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.CreateLogGroupOutput, error)
	CreateLogStream(ctx context.Context, params *cloudwatchlogs.CreateLogStreamInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.CreateLogStreamOutput, error)
	PutLogEvents(ctx context.Context, params *cloudwatchlogs.PutLogEventsInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.PutLogEventsOutput, error)
	PutRetentionPolicy(ctx context.Context, params *cloudwatchlogs.PutRetentionPolicyInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.PutRetentionPolicyOutput, error)
	TagResource(ctx context.Context, params *cloudwatchlogs.TagResourceInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.TagResourceOutput, error)
}

// Possible exceptions are combination of common errors (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/CommonErrors.html)
// and API specific erros (e.g. https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html#API_PutLogEvents_Errors)
type Client struct {
	svc          cloudWatchClient
	logRetention int32
	tags         map[string]string
	logger       *zap.Logger
}

type ClientOption func(*cwLogClientConfig)

type cwLogClientConfig struct {
	userAgentExtras []string
}

func WithUserAgentExtras(userAgentExtras ...string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// Create a log client based on the actual cloudwatch logs client.
func newCloudWatchLogClient(svc cloudWatchClient, logRetention int32, tags map[string]string, logger *zap.Logger) *Client {
	_ = "STUB: not implemented"
	return nil
}

func newCollectorUserAgent(buildInfo component.BuildInfo, logGroupName, componentName string, opts ...ClientOption) string {
	_ = "STUB: not implemented"
	// Loop through each option
	return ""
}

// NewClient create Client
func NewClient(logger *zap.Logger, awsConfig aws.Config, buildInfo component.BuildInfo, logGroupName string, logRetention int32, tags map[string]string, componentName string, opts ...ClientOption) *Client {
	_ = "STUB: not implemented"
	return nil
}

// PutLogEvents mainly handles different possible error could be returned from server side, and retries them
// if necessary.
func (client *Client) PutLogEvents(ctx context.Context, input *cloudwatchlogs.PutLogEventsInput, retryCnt int) error {
	_ = "STUB: not implemented"
	return nil
}

// CloudWatch Logs API was changed to ignore the sequenceToken
// PutLogEvents actions are now accepted and never return
// InvalidSequenceTokenException or DataAlreadyAcceptedException even
// if the sequence token is not valid.
// Finally, InvalidSequenceTokenException and DataAlreadyAcceptedException are
// never returned by the PutLogEvents action.

// Should never happen

// Retry request if OperationAbortedException happens

// Retry request if ServiceUnavailableException happens

// ThrottlingException is handled here because the type cloudwatch.ThrottlingException is not yet available in public SDK
// Drop request if ThrottlingException happens

// TODO: Should have metrics to provide visibility of these failures

// Prepare the readiness for the log group and log stream.
func (client *Client) CreateStream(ctx context.Context, logGroup, streamName *string) error {
	_ = "STUB: not implemented"
	// CreateLogStream / CreateLogGroup
	return nil
}

// Create Log Group with tags if they exist and were specified in the config

// For newly created log groups, set the log retention policy if specified or non-zero. Otherwise, set to Never Expire

// After a log stream is created the token is always empty.

type addToUserAgentHeader struct {
	id, val string
}

var _ middleware.SerializeMiddleware = (*addToUserAgentHeader)(nil)

func (a *addToUserAgentHeader) ID() string { _ = "STUB: not implemented"; return "" }

func (a *addToUserAgentHeader) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (out middleware.SerializeOutput, metadata middleware.Metadata, err error) {
	_ = "STUB: not implemented"
	return *new(middleware.SerializeOutput), *new(middleware.Metadata), nil
}

func AddToUserAgentHeader(id, val string, pos middleware.RelativePosition) func(options *cloudwatchlogs.Options) {
	_ = "STUB: not implemented"
	return nil
}
