// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operationsmanagement // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter/internal/operationsmanagement"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.uber.org/zap"
)

// MetricsClient is responsible for sending the metrics payload to BMC Helix Operations Management
type MetricsClient struct {
	url        string
	httpClient *http.Client
	apiKey     configopaque.String
	logger     *zap.Logger
}

// NewMetricsClient creates a new MetricsClient
func NewMetricsClient(ctx context.Context, clientConfig confighttp.ClientConfig, apiKey configopaque.String, host component.Host, settings component.TelemetrySettings, logger *zap.Logger) (*MetricsClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SendHelixPayload sends the metrics payload to BMC Helix Operations Management
func (mc *MetricsClient) SendHelixPayload(ctx context.Context, payload []BMCHelixOMMetric) error {
	_ = "STUB: not implemented"
	return nil
}

// Log the payload being sent

// Get the JSON encoded payload

// Create a new HTTP request to send the payload

// Send the request

// Check the response status code

// createNewHTTPRequest creates a new HTTP request with the payload
func (mc *MetricsClient) createNewHTTPRequest(ctx context.Context, payloadBytes []byte) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set required headers
