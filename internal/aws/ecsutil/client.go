// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"

import (
	"context"
	"net/http"
	"net/url"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

// Client defines the basic HTTP client interface with GET response validation and content parsing
type Client interface {
	Get(path string) ([]byte, error)
}

// NewClientProvider creates the default rest client provider
func NewClientProvider(baseURL url.URL, clientSettings confighttp.ClientConfig, host component.Host, settings component.TelemetrySettings) ClientProvider {
	_ = "STUB: not implemented"
	return *new(ClientProvider)
}

// ClientProvider defines
type ClientProvider interface {
	BuildClient() (Client, error)
}

type defaultClientProvider struct {
	baseURL        url.URL
	clientSettings confighttp.ClientConfig
	host           component.Host
	settings       component.TelemetrySettings
}

func (dcp *defaultClientProvider) BuildClient() (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func defaultClient(
	ctx context.Context,
	baseURL url.URL,
	clientSettings confighttp.ClientConfig,
	host component.Host,
	settings component.TelemetrySettings,
) (*clientImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ Client = (*clientImpl)(nil)

type clientImpl struct {
	baseURL    url.URL
	httpClient http.Client
	settings   component.TelemetrySettings
}

func (c *clientImpl) Get(path string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *clientImpl) buildReq(path string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
