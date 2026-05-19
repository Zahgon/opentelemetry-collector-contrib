// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

type connectionVars struct {
	InstrumentationKey string
	IngestionURL       string
}

const (
	ApplicationInsightsConnectionString = "APPLICATIONINSIGHTS_CONNECTION_STRING"
	DefaultIngestionEndpoint            = "https://dc.services.visualstudio.com/"
	IngestionEndpointKey                = "IngestionEndpoint"
	InstrumentationKey                  = "InstrumentationKey"
	ConnectionStringMaxLength           = 4096
)

func parseConnectionString(exporterConfig *Config) (*connectionVars, error) {
	_ = "STUB: not implemented"
	// First, try to get the connection string from the environment variable
	return nil, nil
}

// If not found in the environment, use the one from the configuration

func getIngestionURL(ingestionEndpoint string) string { _ = "STUB: not implemented"; return "" }
