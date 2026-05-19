// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package apachesparkreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver/internal/models"
)

var (
	errFailedAppIDCollection = errors.New("failed to retrieve app ids")
	errNoMatchingAllowedApps = errors.New("no apps matched allowed names")
)

type sparkScraper struct {
	client   client
	logger   *zap.Logger
	config   *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

func newSparkScraper(logger *zap.Logger, cfg *Config, settings receiver.Settings) *sparkScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *sparkScraper) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *sparkScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Call applications endpoint to get ids and names for all apps in the cluster

// Check apps against allowed app names from config

// If no app names specified, allow all apps

// Some allowed app names specified, compare to app names from applications endpoint

// Get stats from the 'metrics' endpoint

// For each application id, get stats from stages & executors endpoints

func (s *sparkScraper) recordCluster(clusterStats *models.ClusterProperties, now pcommon.Timestamp, appID, appName string) {
	_ = "STUB: not implemented"
	return
}

func (s *sparkScraper) recordStages(stageStats []models.Stage, now pcommon.Timestamp, appID, appName string) {
	_ = "STUB: not implemented"
	return
}

func (s *sparkScraper) recordExecutors(executorStats []models.Executor, now pcommon.Timestamp, appID, appName string) {
	_ = "STUB: not implemented"
	return
}

func (s *sparkScraper) recordJobs(jobStats []models.Job, now pcommon.Timestamp, appID, appName string) {
	_ = "STUB: not implemented"
	return
}
