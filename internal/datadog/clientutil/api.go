// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clientutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/clientutil"

import (
	"context"
	"errors"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.uber.org/zap"
)

// GZipSubmitMetricsOptionalParameters is used to enable gzip compression for metric payloads submitted by native datadog client
var GZipSubmitMetricsOptionalParameters = datadogV2.NewSubmitMetricsOptionalParameters().WithContentEncoding(datadogV2.METRICCONTENTENCODING_GZIP)

// CreateAPIClient creates a new Datadog API client
func CreateAPIClient(buildInfo component.BuildInfo, endpoint string, hcs confighttp.ClientConfig) *datadog.APIClient {
	_ = "STUB: not implemented"
	return nil
}

// ValidateAPIKey checks if the API key (not the APP key) is valid
func ValidateAPIKey(ctx context.Context, apiKey string, logger *zap.Logger, apiClient *datadog.APIClient) error {
	_ = "STUB: not implemented"
	return nil
}

// GetRequestContext creates a new context with API key for DatadogV2 requests
func GetRequestContext(ctx context.Context, apiKey string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

var ErrInvalidAPI = errors.New("API Key validation failed")
