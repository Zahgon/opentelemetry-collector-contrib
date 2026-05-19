// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package apachesparkreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver/internal/models"
)

const (
	metricsPath      = "/metrics/json"
	applicationsPath = "/api/v1/applications"
)

type client interface {
	Get(path string) ([]byte, error)
	ClusterStats() (*models.ClusterProperties, error)
	Applications() ([]models.Application, error)
	StageStats(appID string) ([]models.Stage, error)
	ExecutorStats(appID string) ([]models.Executor, error)
	JobStats(appID string) ([]models.Job, error)
}

var _ client = (*apacheSparkClient)(nil)

type apacheSparkClient struct {
	client *http.Client
	cfg    *Config
	logger *zap.Logger
}

// newApacheSparkClient creates a new client to make requests for the Apache Spark receiver.
func newApacheSparkClient(ctx context.Context, cfg *Config, host component.Host, settings component.TelemetrySettings) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

// Get issues an authorized Get requests to the specified URL.
func (c *apacheSparkClient) Get(path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *apacheSparkClient) ClusterStats() (*models.ClusterProperties, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *apacheSparkClient) Applications() ([]models.Application, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *apacheSparkClient) StageStats(appID string) ([]models.Stage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *apacheSparkClient) ExecutorStats(appID string) ([]models.Executor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *apacheSparkClient) JobStats(appID string) ([]models.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *apacheSparkClient) buildReq(path string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
