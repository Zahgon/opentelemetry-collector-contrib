// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package flinkmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver"

import (
	"context"
	"net/http"
	"os"
	"strings"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver/internal/models"
)

// The API endpoints required to collect metrics.
const (
	// jobmanagerMetricEndpoint gets jobmanager metrics.
	jobmanagerMetricEndpoint = "/jobmanager/metrics"
	// taskmanagersEndpoint gets taskmanager IDs.
	taskmanagersEndpoint = "/taskmanagers"
	// taskmanagersMetricEndpoint gets taskmanager using a taskmanager ID.
	taskmanagersMetricEndpoint = "/taskmanagers/%s/metrics"
	// jobsEndpoint gets job IDs.
	jobsEndpoint = "/jobs"
	// jobsOverviewEndpoint gets job IDs with associated Job names.
	jobsOverviewEndpoint = "/jobs/overview"
	// jobsWithIDEndpoint gets vertex IDs using a job ID.
	jobsWithIDEndpoint = "/jobs/%s"
	// jobsMetricEndpoint gets job metrics using a job ID.
	jobsMetricEndpoint = "/jobs/%s/metrics"
	// verticesEndpoint gets subtask index's using a job and vertex ID.
	verticesEndpoint = "/jobs/%s/vertices/%s"
	// subtaskMetricEndpoint gets subtask metrics using a job ID, vertex ID and subtask index.
	subtaskMetricEndpoint = "/jobs/%s/vertices/%s/subtasks/%v/metrics"
)

type client interface {
	GetJobmanagerMetrics(ctx context.Context) (*models.JobmanagerMetrics, error)
	GetTaskmanagersMetrics(ctx context.Context) ([]*models.TaskmanagerMetrics, error)
	GetJobsMetrics(ctx context.Context) ([]*models.JobMetrics, error)
	GetSubtasksMetrics(ctx context.Context) ([]*models.SubtaskMetrics, error)
}

type flinkClient struct {
	client       *http.Client
	hostEndpoint string
	hostName     string
	logger       *zap.Logger
}

func newClient(ctx context.Context, cfg *Config, host component.Host, settings component.TelemetrySettings, logger *zap.Logger) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

func (c *flinkClient) get(ctx context.Context, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	// Construct endpoint and create request
	return nil, nil
}

// Make request

// Defer body close

// Check for OK status code

// Attempt to extract the error payload

// getMetrics makes a request to a metric endpoint to get the metric names, the another request building a query to get the metric values.
func (c *flinkClient) getMetrics(ctx context.Context, path string) (*models.MetricsResponse, error) {
	_ = "STUB: not implemented"
	// Get the metric names
	return nil, nil
}

// Populates the metric names

// Construct a get query parameter using comma-separated list of string values to select specific metrics

// Get the metric values using the query

// Populates metric values

// GetJobManagerMetrics gets the jobmanager metrics.
func (c *flinkClient) GetJobmanagerMetrics(ctx context.Context) (*models.JobmanagerMetrics, error) {
	_ = "STUB: not implemented"
	// Get the metric names and values for jobmanager
	return nil, nil
}

// Add a hostname used to identify between multiple jobmanager instances

// GetTaskmanagersMetrics gets the Taskmanager metrics for each taskmanager.
func (c *flinkClient) GetTaskmanagersMetrics(ctx context.Context) ([]*models.TaskmanagerMetrics, error) {
	_ = "STUB: not implemented"
	// Get the taskmanager id list
	return nil, nil
}

// Populates taskmanager id names

// Get taskmanager metrics for each taskmanager id

// getTaskmanagersMetricsByIDs gets taskmanager metrics for each task manager id.
func (c *flinkClient) getTaskmanagersMetricsByIDs(ctx context.Context, taskmanagerIDs *models.TaskmanagerIDsResponse) ([]*models.TaskmanagerMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetJobsMetrics gets the job metrics for each job.
func (c *flinkClient) GetJobsMetrics(ctx context.Context) ([]*models.JobMetrics, error) {
	_ = "STUB: not implemented"
	// Get the job id and name list
	return nil, nil
}

// Populates job id and names

// Get job metrics for each job id

// getJobsMetricsByIDs gets jobs metrics for each job id.
func (c *flinkClient) getJobsMetricsByIDs(ctx context.Context, jobIDs *models.JobOverviewResponse) ([]*models.JobMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSubtasksMetrics gets subtask metrics for each job id, vertex id and subtask index.
func (c *flinkClient) GetSubtasksMetrics(ctx context.Context) ([]*models.SubtaskMetrics, error) {
	_ = "STUB: not implemented"
	// Get the job id's
	return nil, nil
}

// Populates the job id

// getSubtasksMetricsByIDs gets subtask metrics for each job id, vertex id and subtask index.
func (c *flinkClient) getSubtasksMetricsByIDs(ctx context.Context, jobsResponse *models.JobsResponse) ([]*models.SubtaskMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get vertices for each job

// Populates the job response with vertices info

// Gets subtask info for each vertex id

// Populates the vertex response with subtask info

// Gets subtask metrics for each vertex id

// Stores subtask info with additional attribute values to uniquely identify metrics

// Override for testing
var osHostname = os.Hostname

func getHostname() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Override for testing
var taskmanagerHost = strings.Split

func getTaskmanagerHost(id string) string { _ = "STUB: not implemented"; return "" }

func reflect(s string) string {
	_ = "STUB: not implemented"

	// Override for testing
	return ""
}

var taskmanagerID = reflect

func getTaskmanagerID(id string) string { _ = "STUB: not implemented"; return "" }
