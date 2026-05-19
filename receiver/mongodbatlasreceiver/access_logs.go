// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	rcvr "go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
)

const (
	accessLogStorageKey           = "last_endtime_access_logs_%s"
	defaultAccessLogsPollInterval = 5 * time.Minute
	defaultAccessLogsPageSize     = 20000
	defaultAccessLogsMaxPages     = 10
)

type accessLogStorageRecord struct {
	ClusterName       string    `json:"cluster_name"`
	NextPollStartTime time.Time `json:"next_poll_start_time"`
}

type accessLogClient interface {
	GetProject(ctx context.Context, groupID string) (*mongodbatlas.Project, error)
	GetClusters(ctx context.Context, groupID string) ([]mongodbatlas.Cluster, error)
	GetAccessLogs(ctx context.Context, groupID, clusterName string, opts *internal.GetAccessLogsOptions) (ret []*mongodbatlas.AccessLogs, err error)
}

type accessLogsReceiver struct {
	client        accessLogClient
	logger        *zap.Logger
	storageClient storage.Client
	cfg           *Config
	consumer      consumer.Logs

	record     map[string][]*accessLogStorageRecord
	authResult *bool
	wg         *sync.WaitGroup
	cancel     context.CancelFunc
}

func newAccessLogsReceiver(settings rcvr.Settings, cfg *Config, consumer consumer.Logs) (*accessLogsReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (alr *accessLogsReceiver) Start(ctx context.Context, _ component.Host, storageClient storage.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (alr *accessLogsReceiver) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (alr *accessLogsReceiver) startPolling(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (alr *accessLogsReceiver) pollAccessLogs(ctx context.Context, pc *LogsProjectConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (alr *accessLogsReceiver) pollCluster(ctx context.Context, pc *LogsProjectConfig, project *mongodbatlas.Project, cluster *mongodbatlas.Cluster, startTime, now time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Assume failure, in which case we poll starting with the same startTime
// unless we successfully make request(s) for access logs and they are successfully sent to the consumer

// No logs retrieved, try again on next interval with the same start time as the API may not have
// all logs for the given time available to be queried yet (undocumented behavior)

// The first page of results will have the latest data, so we want to update the nextPollStartTime
// There is risk of data loss at this point if we are unable to then process the remaining pages
// of data, but that is a limitation of the API that we can't work around.

// This slice access is safe as we have previously confirmed that the slice is not empty

// If we are not able to get the latest log timestamp, we have to assume that we are collecting all
// data and don't want to risk duplicated data by re-polling the same data again.

// If we get back less than the maximum number of logs, we can assume that we've retrieved all of the logs
// that are currently available for this time period, though some logs may not be available in the API yet.

// If we get back the maximum number of logs, we need to re-query with a new end time. While undocumented, the API
// returns the most recent logs first. If we get the maximum number of logs back, we can assume that
// there are more logs to be retrieved. We'll re-query with the same start time, but the end
// time set to just before the timestamp of the oldest log entry returned.

// If the new max date is before the min date, we've retrieved all of the logs for this time period
// and receiving the maximum number of logs back is a coincidence.

func getTimestamp(log *mongodbatlas.AccessLogs) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// If body couldn't be parsed, we'll still use the outer Timestamp field to determine the new max date.

func getTimestampPreparsedBody(log *mongodbatlas.AccessLogs, body map[string]any) (time.Time, error) {
	_ = "STUB: not implemented"
	// If the log message has a timestamp, use that. When present, it has more precision than the timestamp from the access log entry.
	return *new(time.Time), nil
}

// If the log message doesn't have a timestamp, use the timestamp from the outer access log entry.

// The documentation claims ISO8601/RFC3339, but the API has been observed returning timestamps in UnixDate format
// UnixDate looks like Wed Apr 26 02:38:56 GMT 2023

// Return the original error as the documentation claims ISO8601

func parseLogMessage(log *mongodbatlas.AccessLogs) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func transformAccessLogs(now pcommon.Timestamp, accessLogs []*mongodbatlas.AccessLogs, p *mongodbatlas.Project, c *mongodbatlas.Cluster, logger *zap.Logger) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// Expected format documented https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Access-Tracking/operation/listAccessLogsByClusterName

func accessLogsCheckpointKey(groupID string) string { _ = "STUB: not implemented"; return "" }

func (alr *accessLogsReceiver) checkpoint(ctx context.Context, groupID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (alr *accessLogsReceiver) loadCheckpoint(ctx context.Context, groupID string) {
	_ = "STUB: not implemented"
	return
}

func (alr *accessLogsReceiver) getClusterCheckpoint(groupID, clusterName string) *accessLogStorageRecord {
	_ = "STUB: not implemented"
	return nil
}

func (alr *accessLogsReceiver) setClusterCheckpoint(groupID string, clusterCheckpoint *accessLogStorageRecord) {
	_ = "STUB: not implemented"
	return
}
