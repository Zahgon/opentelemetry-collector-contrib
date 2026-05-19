// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bmchelixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter"

import (
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configretry"
)

// Config struct is used to store the configuration of the exporter
type Config struct {
	confighttp.ClientConfig `mapstructure:",squash"`
	// APIKey authenticates requests to BMC Helix Operations Management.
	// Connect to BMC Helix Operations Management, go to the Administration > Repository page, and
	// click on the Copy API Key button to get your API Key.
	// Alternatively, it is recommended to create and use a dedicated authentication key for external integration:
	// https://docs.bmc.com/xwiki/bin/view/Helix-Common-Services/BMC-Helix-Portal/BMC-Helix-Portal/helixportal261/Administering/Using-API-keys-for-external-integrations/
	APIKey      configopaque.String       `mapstructure:"api_key"`
	RetryConfig configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	// EnrichMetricWithAttributes enables enriched metric creation by appending datapoint
	// attribute values to the metric name. This increases metric cardinality but provides
	// more detailed identification in BMC Helix Operations Management. Default is true.
	EnrichMetricWithAttributes bool `mapstructure:"enrich_metric_with_attributes"`
}

// validate the configuration
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
