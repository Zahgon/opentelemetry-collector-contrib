// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package riakreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver/internal/model"
)

// statsPath is the path to stats endpoint
const statsPath = "/stats"

type client interface {
	// GetStats calls "/stats" endpoint to get list of stats for the target node
	GetStats(ctx context.Context) (*model.Stats, error)
}

var _ client = (*riakClient)(nil)

type riakClient struct {
	client       *http.Client
	hostEndpoint string
	creds        riakCredentials
	logger       *zap.Logger
}

type riakCredentials struct {
	username string
	password string
}

func newClient(ctx context.Context, cfg *Config, host component.Host, settings component.TelemetrySettings, logger *zap.Logger) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

func (c *riakClient) GetStats(ctx context.Context) (*model.Stats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *riakClient) get(ctx context.Context, path string, respObj any) error {
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
