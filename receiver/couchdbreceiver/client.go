// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package couchdbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

// client defines the basic HTTP client interface.
type client interface {
	Get(path string) ([]byte, error)
	GetStats(nodeName string) (map[string]any, error)
}

var _ client = (*couchDBClient)(nil)

type couchDBClient struct {
	client *http.Client
	cfg    *Config
	logger *zap.Logger
}

// newCouchDBClient creates a new client to make requests for the CouchDB receiver.
func newCouchDBClient(ctx context.Context, cfg *Config, host component.Host, settings component.TelemetrySettings) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

// Get issues an authorized Get requests to the specified url.
func (c *couchDBClient) Get(path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStats gets couchdb stats at a specific node name endpoint.
func (c *couchDBClient) GetStats(nodeName string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *couchDBClient) buildReq(path string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
