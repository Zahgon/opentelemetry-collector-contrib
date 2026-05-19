// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package webhookeventreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/config/confighttp"
)

var (
	errMissingEndpointFromConfig   = errors.New("missing receiver server endpoint from config")
	errReadTimeoutExceedsMaxValue  = errors.New("the duration specified for read_timeout exceeds the maximum allowed value of 10s")
	errWriteTimeoutExceedsMaxValue = errors.New("the duration specified for write_timeout exceeds the maximum allowed value of 10s")
	errRequiredHeader              = errors.New("both key and value are required to assign a required_header")
	errHeaderAttributeRegexCompile = errors.New("regex for header_attribute_regex failed to compile")
)

// Config defines configuration for the Generic Webhook receiver.
type Config struct {
	confighttp.ServerConfig    `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
	ReadTimeout                string                   `mapstructure:"read_timeout"`                  // wait time for reading request headers in ms. Default is 500ms.
	WriteTimeout               string                   `mapstructure:"write_timeout"`                 // wait time for writing request response in ms. Default is 500ms.
	Path                       string                   `mapstructure:"path"`                          // path for data collection. Default is /events
	HealthPath                 string                   `mapstructure:"health_path"`                   // path for health check api. Default is /health_check
	RequiredHeader             RequiredHeader           `mapstructure:"required_header"`               // optional setting to set a required header for all requests to have
	SplitLogsAtNewLine         bool                     `mapstructure:"split_logs_at_newline"`         // optional setting to split logs into multiple log records
	SplitLogsAtJSONBoundary    bool                     `mapstructure:"split_logs_at_json_boundary"`   // optional setting to split logs at JSON object boundaries
	ConvertHeadersToAttributes bool                     `mapstructure:"convert_headers_to_attributes"` // optional to convert all headers to attributes
	HeaderAttributeRegex       string                   `mapstructure:"header_attribute_regex"`        // optional to convert headers matching a regex to log attributes
}

type RequiredHeader struct {
	Key   string `mapstructure:"key"`
	Value string `mapstructure:"value"`
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// If a user defines a custom read/write timeout there is a maximum value
// of 10s imposed here.

// Set default MaxRequestBodySize if not configured

// 20MiB
// to match default value http://github.com/open-telemetry/opentelemetry-collector/blob/release/v0.139.x/config/confighttp/server.go#L31

func (cfg *Config) ShouldSplitLogsAtJSONBoundary() bool { _ = "STUB: not implemented"; return false }
