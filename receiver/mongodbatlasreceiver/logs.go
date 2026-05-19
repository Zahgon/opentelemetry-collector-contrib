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
	rcvr "go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/model"
)

const mongoDBMajorVersion4_2 = "4.2"

type logsReceiver struct {
	log         *zap.Logger
	cfg         *Config
	client      *internal.MongoDBAtlasClient
	consumer    consumer.Logs
	stopperChan chan struct{}
	wg          sync.WaitGroup
	start       time.Time
	end         time.Time
}

type projectContext struct {
	Project mongodbatlas.Project
	orgName string
}

// MongoDB Atlas Documentation recommends a polling interval of 5 minutes: https://www.mongodb.com/docs/atlas/reference/api/logs/#logs
const collectionInterval = time.Minute * 5

func newMongoDBAtlasLogsReceiver(settings rcvr.Settings, cfg *Config, consumer consumer.Logs) (*logsReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Log receiver logic
func (s *logsReceiver) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// collection interval loop,

func (s *logsReceiver) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// parseHostNames parses out the hostname from the specified cluster host
func parseHostNames(s string, logger *zap.Logger) []string { _ = "STUB: not implemented"; return nil }

// separate hostname from scheme and port

// collect spins off functionality of the receiver from the Start function
func (s *logsReceiver) collect(ctx context.Context) { _ = "STUB: not implemented"; return }

// get clusters for each of the projects

func (s *logsReceiver) processClusters(ctx context.Context, projectCfg LogsProjectConfig, projectID string) ([]mongodbatlas.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type clusterInfo struct {
	ClusterName         string
	RegionName          string
	ProviderName        string
	MongoDBMajorVersion string
}

func (s *logsReceiver) collectClusterLogs(clusters []mongodbatlas.Cluster, projectCfg LogsProjectConfig, pc projectContext) {
	_ = "STUB: not implemented"
	return
}

// Defaults to true if not specified

// Defaults to false if not specified

func filterClusters(clusters []mongodbatlas.Cluster, projectCfg ProjectConfig) ([]mongodbatlas.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check to include or exclude clusters

// keep all clusters if include and exclude are not specified

// include is initialized

// exclude is initialized

// both are initialized

func (s *logsReceiver) getHostLogs(groupID, hostname, logName, clusterMajorVersion string) ([]model.LogEntry, error) {
	_ = "STUB: not implemented"
	// Get gzip bytes buffer from API
	return nil, nil
}

func (s *logsReceiver) getHostAuditLogs(groupID, hostname, logName string) ([]model.AuditLog, error) {
	_ = "STUB: not implemented"
	// Get gzip bytes buffer from API
	return nil, nil
}

func (s *logsReceiver) collectLogs(pc projectContext, hostname, logName string, clusterInfo clusterInfo) {
	_ = "STUB: not implemented"
	return
}

func (s *logsReceiver) collectAuditLogs(pc projectContext, hostname, logName string, clusterInfo clusterInfo) {
	_ = "STUB: not implemented"
	return
}
