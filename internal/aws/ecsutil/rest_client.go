// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"

import (
	"net/url"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

func NewRestClient(baseEndpoint url.URL, clientSettings confighttp.ClientConfig, settings component.TelemetrySettings) (RestClient, error) {
	_ = "STUB: not implemented"
	return *new(RestClient), nil
}

// TODO: Instead of using this, expose it as a argument to NewRestClient.
type nopHost struct {
	component.Host
}

func (*nopHost) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}

// RestClient is swappable for testing.
type RestClient interface {
	GetResponse(path string) ([]byte, error)
}

// TaskMetadataRestClient is a thin wrapper around an ecs task metadata client, encapsulating endpoints
// and their corresponding http methods.
type TaskMetadataRestClient struct {
	client Client
}

// NewRestClientFromClient creates a new copy of the Client
func NewRestClientFromClient(client Client) *TaskMetadataRestClient {
	_ = "STUB: not implemented"
	return nil
}

// GetResponse gets the desired path from the configured metadata endpoint
func (c *TaskMetadataRestClient) GetResponse(path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
