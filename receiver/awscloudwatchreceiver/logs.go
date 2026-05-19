// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscloudwatchreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver"

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
)

const (
	noStreamName             = "THIS IS INVALID STREAM"
	maxLogGroupsPerDiscovery = int32(50)
)

type logsReceiver struct {
	settings                      receiver.Settings
	region                        string
	profile                       string
	imdsEndpoint                  string
	pollInterval                  time.Duration
	maxEventsPerRequest           int
	initialStartTime              time.Time
	groupNextStartTimes           map[string]time.Time
	groupRequests                 []groupRequest
	autodiscover                  *AutodiscoverConfig
	client                        client
	consumer                      consumer.Logs
	wg                            *sync.WaitGroup
	doneChan                      chan bool
	storageID                     *component.ID
	cloudwatchCheckpointPersister *cloudwatchCheckpointPersister
}

type client interface {
	DescribeLogGroups(ctx context.Context, input *cloudwatchlogs.DescribeLogGroupsInput, opts ...func(options *cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogGroupsOutput, error)
	FilterLogEvents(ctx context.Context, input *cloudwatchlogs.FilterLogEventsInput, opts ...func(options *cloudwatchlogs.Options)) (*cloudwatchlogs.FilterLogEventsOutput, error)
}

type streamNames struct {
	group string
	names []*string
}

func (sn *streamNames) request(limit int, nextToken string, st, et *time.Time) *cloudwatchlogs.FilterLogEventsInput {
	_ = "STUB: not implemented"
	return nil
}

func (sn *streamNames) groupName() string { _ = "STUB: not implemented"; return "" }

type streamPrefix struct {
	group  string
	prefix *string
}

func (sp *streamPrefix) request(limit int, nextToken string, st, et *time.Time) *cloudwatchlogs.FilterLogEventsInput {
	_ = "STUB: not implemented"
	return nil
}

func (sp *streamPrefix) groupName() string { _ = "STUB: not implemented"; return "" }

type groupRequest interface {
	request(limit int, nextToken string, st, et *time.Time) *cloudwatchlogs.FilterLogEventsInput
	groupName() string
}

func newLogsReceiver(cfg *Config, settings receiver.Settings, consumer consumer.Logs) *logsReceiver {
	_ = "STUB: not implemented"
	return nil
}

// safeguard from using both

func (l *logsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *logsReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (l *logsReceiver) startPolling(ctx context.Context) { _ = "STUB: not implemented"; return }

func (l *logsReceiver) poll(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Retrieve the last persisted timestamp for this log group if exists

// Poll logs for the current log group

// Persist the new end time as the checkpoint for this log group

// Update the receiver's nextStartTime for the next poll cycle

// Clean up stale entries from groupNextStartTimes map

func (l *logsReceiver) pollForLogs(ctx context.Context, pc groupRequest, startTime, endTime time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	// In case of failure, the startTime of the request will be used as the checkpoint for the next poll
	return *new(time.Time), nil
}

// the next timestamp should be 1 more millisecond than the last log

// Skip the time range in case there are no logs

func (l *logsReceiver) processEvents(now pcommon.Timestamp, logGroupName string, output *cloudwatchlogs.FilterLogEventsOutput) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// Now we know resourceLogs is initialized and has one scopeLogs so we don't have to handle any special cases.

func (l *logsReceiver) discoverGroups(ctx context.Context, auto *AutodiscoverConfig) ([]groupRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default behavior is to collect all if not stream filtered

func (l *logsReceiver) ensureSession() error { _ = "STUB: not implemented"; return nil }
