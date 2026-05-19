// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package rabbitmqreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver/internal/models"
)

const (
	// queuePath is the endpoint for RabbitMQ queues.
	queuePath = "/api/queues"

	// nodePath is the endpoint for RabbitMQ nodes.
	nodePath = "/api/nodes"
)

type client interface {
	// GetQueues calls "/api/queues" endpoint to get list of queues for the target node
	GetQueues(ctx context.Context) ([]*models.Queue, error)
	// GetNodes calls "/api/nodes" endpoint to get list of nodes for the target node
	GetNodes(ctx context.Context) ([]*models.Node, error)
}

var _ client = (*rabbitmqClient)(nil)

type rabbitmqClient struct {
	client       *http.Client
	hostEndpoint string
	creds        rabbitmqCredentials
	logger       *zap.Logger
}

type rabbitmqCredentials struct {
	username string
	password string
}

func newClient(ctx context.Context, cfg *Config, host component.Host, settings component.TelemetrySettings, logger *zap.Logger) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

func (c *rabbitmqClient) GetQueues(ctx context.Context) ([]*models.Queue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *rabbitmqClient) GetNodes(ctx context.Context) ([]*models.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *rabbitmqClient) get(ctx context.Context, path string, respObj any) error {
	_ = "STUB: not implemented"
	// Construct endpoint and create request
	return nil
}

// Set user/pass authentication

// Make request

// Defer body close

// Check for OK status code

// Attempt to extract the error payload

// Decode the payload into the passed in response object
