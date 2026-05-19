// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/metadata"
)

type mongodbatlasreceiver struct {
	log         *zap.Logger
	cfg         *Config
	client      *internal.MongoDBAtlasClient
	lastRun     time.Time
	mb          *metadata.MetricsBuilder
	stopperChan chan struct{}
}

type timeconstraints struct {
	start      string
	end        string
	resolution string
}

func newMongoDBAtlasReceiver(settings receiver.Settings, cfg *Config) (*mongodbatlasreceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use index-based iteration: cfg.Projects is []ProjectConfig (value slice),
// so a range-copy would make populateIncludesAndExcludes a no-op on the
// original elements.

func newMongoDBAtlasScraper(recv *mongodbatlasreceiver) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}

func (s *mongodbatlasreceiver) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *mongodbatlasreceiver) timeConstraints(now time.Time) timeconstraints {
	_ = "STUB: not implemented"
	return *new(timeconstraints)
}

func (s *mongodbatlasreceiver) shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil

	// poll decides whether to poll all projects or a specific project based on the configuration.
}

func (s *mongodbatlasreceiver) poll(ctx context.Context, time timeconstraints) error {
	_ = "STUB: not implemented"
	return nil
}

// pollAllProjects handles polling across all projects within the organizations.
func (s *mongodbatlasreceiver) pollAllProjects(ctx context.Context, time timeconstraints) error {
	_ = "STUB: not implemented"
	return nil
}

// Since there is no specific ProjectConfig for these projects, pass nil.

// pollProject handles polling for specific projects as configured.
func (s *mongodbatlasreceiver) pollProjects(ctx context.Context, time timeconstraints) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *mongodbatlasreceiver) processProject(ctx context.Context, time timeconstraints, orgName string, project *mongodbatlas.Project, projectCfg *ProjectConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip processing for this cluster

// shouldProcessCluster checks whether a given cluster should be processed based on the project configuration.
func shouldProcessCluster(projectCfg *ProjectConfig, clusterName string) bool {
	_ = "STUB: not implemented"
	return false

	// If there is no project config, process all clusters.
}

// Return false immediately if the cluster is excluded.

// If IncludeClusters is empty, or the cluster is explicitly included, return true.

type providerValues struct {
	RegionName   string
	ProviderName string
}

func (s *mongodbatlasreceiver) getNodeClusterNameMap(
	ctx context.Context,
	projectID string,
) (map[string]string, map[string]providerValues, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// URI in the form mongodb://host1.mongodb.net:27017,host2.mongodb.net:27017,host3.mongodb.net:27017

// Remove the port from the node

func (s *mongodbatlasreceiver) extractProcessMetrics(
	ctx context.Context,
	time timeconstraints,
	orgName string,
	project *mongodbatlas.Project,
	process *mongodbatlas.Process,
	clusterName string,
	providerValues providerValues,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *mongodbatlasreceiver) extractProcessDatabaseMetrics(
	ctx context.Context,
	time timeconstraints,
	orgName string,
	project *mongodbatlas.Project,
	process *mongodbatlas.Process,
	clusterName string,
	providerValues providerValues,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *mongodbatlasreceiver) extractProcessDiskMetrics(
	ctx context.Context,
	time timeconstraints,
	orgName string,
	project *mongodbatlas.Project,
	process *mongodbatlas.Process,
	clusterName string,
	providerValues providerValues,
) error {
	_ = "STUB: not implemented"
	return nil
}
