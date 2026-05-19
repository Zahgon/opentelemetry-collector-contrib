// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"

import (
	"bytes"
	"context"
	"net/http"
	"sync"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/config/configretry"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/metadata"
)

type clientRoundTripper struct {
	originalTransport http.RoundTripper
	log               *zap.Logger
	backoffConfig     configretry.BackOffConfig
	stopped           bool
	mutex             sync.Mutex
	shutdownChan      chan struct{}
}

func newClientRoundTripper(
	originalTransport http.RoundTripper,
	log *zap.Logger,
	backoffConfig configretry.BackOffConfig,
) *clientRoundTripper {
	_ = "STUB: not implemented"
	return nil
}

func (rt *clientRoundTripper) isStopped() bool { _ = "STUB: not implemented"; return false }

func (rt *clientRoundTripper) stop() { _ = "STUB: not implemented"; return }

func (rt *clientRoundTripper) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (rt *clientRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Can't do anything

// MongoDBAtlasClient wraps the official MongoDB Atlas client to manage pagination
// and mapping to OpenTelemetry metric and log structures.
type MongoDBAtlasClient struct {
	log          *zap.Logger
	client       *mongodbatlas.Client
	transport    *http.Transport
	roundTripper *clientRoundTripper
}

// NewMongoDBAtlasClient creates a new MongoDB Atlas client wrapper
func NewMongoDBAtlasClient(
	baseURL string,
	publicKey string,
	privateKey string,
	backoffConfig configretry.BackOffConfig,
	log *zap.Logger,
) (*MongoDBAtlasClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MongoDBAtlasClient) Shutdown() error { _ = "STUB: not implemented"; return nil }

// Check both the returned error and the status of the HTTP response
func checkMongoDBClientErr(err error, response *mongodbatlas.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func hasNext(links []*mongodbatlas.Link) bool { _ = "STUB: not implemented"; return false }

// Organizations returns a list of all organizations available with the supplied credentials
func (s *MongoDBAtlasClient) Organizations(ctx context.Context) ([]*mongodbatlas.Organization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Add error to a metric
// Stop, returning what we have (probably empty slice)

func (s *MongoDBAtlasClient) getOrganizationsPage(
	ctx context.Context,
	pageNum int,
) ([]*mongodbatlas.Organization, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// GetOrganization retrieves a single organization specified by orgID
func (s *MongoDBAtlasClient) GetOrganization(ctx context.Context, orgID string) (*mongodbatlas.Organization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Projects returns a list of projects accessible within the provided organization
func (s *MongoDBAtlasClient) Projects(
	ctx context.Context,
	orgID string,
) ([]*mongodbatlas.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetProject returns a single project specified by projectName
func (s *MongoDBAtlasClient) GetProject(ctx context.Context, projectName string) (*mongodbatlas.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MongoDBAtlasClient) getProjectsPage(
	ctx context.Context,
	orgID string,
	pageNum int,
) ([]*mongodbatlas.Project, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Processes returns the list of processes running for a given project.
func (s *MongoDBAtlasClient) Processes(
	ctx context.Context,
	projectID string,
) ([]*mongodbatlas.Process, error) {
	_ = "STUB: not implemented"
	// A paginated API, but the MongoDB client just returns the values from the first page
	return nil, nil
}

// Note: MongoDB Atlas also has the idea of a Cluster- we can retrieve a list of clusters from
// the Project, but a Cluster does not have a link to its Process list and a Process does not
// have a link to its Cluster (save through the hostname, which is not a documented relationship).

func (s *MongoDBAtlasClient) getProcessDatabasesPage(
	ctx context.Context,
	projectID string,
	host string,
	port int,
	pageNum int,
) ([]*mongodbatlas.ProcessDatabase, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// ProcessDatabases lists databases that are running in a given MongoDB Atlas process
func (s *MongoDBAtlasClient) ProcessDatabases(
	ctx context.Context,
	projectID string,
	host string,
	port int,
) ([]*mongodbatlas.ProcessDatabase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProcessMetrics returns a set of metrics associated with the specified running process.
func (s *MongoDBAtlasClient) ProcessMetrics(
	ctx context.Context,
	mb *metadata.MetricsBuilder,
	projectID string,
	host string,
	port int,
	start string,
	end string,
	resolution string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Return partial results

func (s *MongoDBAtlasClient) getProcessMeasurementsPage(
	ctx context.Context,
	projectID string,
	host string,
	port int,
	pageNum int,
	start string,
	end string,
	resolution string,
) ([]*mongodbatlas.Measurements, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// ProcessDatabaseMetrics returns metrics about a particular database running within a MongoDB Atlas process
func (s *MongoDBAtlasClient) ProcessDatabaseMetrics(
	ctx context.Context,
	mb *metadata.MetricsBuilder,
	projectID string,
	host string,
	port int,
	dbname string,
	start string,
	end string,
	resolution string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MongoDBAtlasClient) getProcessDatabaseMeasurementsPage(
	ctx context.Context,
	projectID string,
	host string,
	port int,
	dbname string,
	pageNum int,
	start string,
	end string,
	resolution string,
) ([]*mongodbatlas.Measurements, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// ProcessDisks enumerates the disks accessible to a specified MongoDB Atlas process
func (s *MongoDBAtlasClient) ProcessDisks(
	ctx context.Context,
	projectID string,
	host string,
	port int,
) []*mongodbatlas.ProcessDisk {
	_ = "STUB: not implemented"
	return nil
}

// Return partial results

func (s *MongoDBAtlasClient) getProcessDisksPage(
	ctx context.Context,
	projectID string,
	host string,
	port int,
	pageNum int,
) ([]*mongodbatlas.ProcessDisk, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// ProcessDiskMetrics returns metrics supplied for a particular disk partition used by a MongoDB Atlas process
func (s *MongoDBAtlasClient) ProcessDiskMetrics(
	ctx context.Context,
	mb *metadata.MetricsBuilder,
	projectID string,
	host string,
	port int,
	partitionName string,
	start string,
	end string,
	resolution string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MongoDBAtlasClient) processDiskMeasurementsPage(
	ctx context.Context,
	projectID string,
	host string,
	port int,
	partitionName string,
	pageNum int,
	start string,
	end string,
	resolution string,
) ([]*mongodbatlas.Measurements, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// GetLogs retrieves the logs from the mongo API using API call: https://www.mongodb.com/docs/atlas/reference/api/logs/#syntax
func (s *MongoDBAtlasClient) GetLogs(ctx context.Context, groupID, hostname, logName string, start, end time.Time) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusters retrieves the clusters from the mongo API using API call: https://www.mongodb.com/docs/atlas/reference/api/clusters-get-all/#request
func (s *MongoDBAtlasClient) GetClusters(ctx context.Context, groupID string) ([]mongodbatlas.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AlertPollOptions struct {
	PageNum  int
	PageSize int
}

// GetAlerts returns the alerts specified for the set projects
func (s *MongoDBAtlasClient) GetAlerts(ctx context.Context, groupID string, opts *AlertPollOptions) (ret []mongodbatlas.Alert, nextPage bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// GetEventsOptions are the options to use for making a request to get Project Events
type GetEventsOptions struct {
	// Which page of the paginated events
	PageNum int
	// How large the Pages will be
	PageSize int
	// The list of Event Types https://www.mongodb.com/docs/atlas/reference/api/events-projects-get-all/#event-type-values
	// to grab from the API
	EventTypes []string
	// The oldest date to look back for the events
	MinDate time.Time
	// the newest time to accept events
	MaxDate time.Time
}

// GetProjectEvents returns the events specified for the set projects
func (s *MongoDBAtlasClient) GetProjectEvents(ctx context.Context, groupID string, opts *GetEventsOptions) (ret []*mongodbatlas.Event, nextPage bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Earliest Timestamp in ISO 8601 date and time format in UTC from when Atlas should return events.

// GetOrgEvents returns the events specified for the set organizations
func (s *MongoDBAtlasClient) GetOrganizationEvents(ctx context.Context, orgID string, opts *GetEventsOptions) (ret []*mongodbatlas.Event, nextPage bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Earliest Timestamp in ISO 8601 date and time format in UTC from when Atlas should return events.

// GetAccessLogsOptions are the options to use for making a request to get Access Logs
type GetAccessLogsOptions struct {
	// The oldest date to look back for the events
	MinDate time.Time
	// the newest time to accept events
	MaxDate time.Time
	// If true, only return successful access attempts; if false, only return failed access attempts
	// If nil, return both successful and failed access attempts
	AuthResult *bool
	// Maximum number of entries to return
	NLogs int
}

// GetAccessLogs returns the access logs specified for the cluster requested
func (s *MongoDBAtlasClient) GetAccessLogs(ctx context.Context, groupID, clusterName string, opts *GetAccessLogsOptions) (ret []*mongodbatlas.AccessLogs, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Earliest Timestamp in epoch milliseconds from when Atlas should access log results

// Latest Timestamp in epoch milliseconds from when Atlas should access log results

// If true, only return successful access attempts; if false, only return failed access attempts
// If nil, return both successful and failed access attempts

// Maximum number of entries to return (0-20000)

func toUnixString(t time.Time) string { _ = "STUB: not implemented"; return "" }
