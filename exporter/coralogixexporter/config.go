// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	cxAppNameAttrName       = "cx.application.name"
	cxSubsystemNameAttrName = "cx.subsystem.name"
	httpProtocol            = "http"
	grpcProtocol            = "grpc"
)

// TransportConfig extends configgrpc.ClientConfig with additional HTTP-specific settings
type TransportConfig struct {
	// Embed the gRPC configuration to ensure backward compatibility
	configgrpc.ClientConfig `mapstructure:",squash"`

	// The following fields are only used when protocol is "http"
	ProxyURL string        `mapstructure:"proxy_url,omitempty"` // Used only if protocol is http
	Timeout  time.Duration `mapstructure:"timeout,omitempty"`   // Used only if protocol is http

	// AcceptEncoding specifies the compression encoding to accept for gRPC responses.
	// Defaults to "gzip" if not set. Only used when protocol is "grpc".
	AcceptEncoding string `mapstructure:"accept_encoding,omitempty"`
}

func (c *TransportConfig) ToHTTPClient(ctx context.Context, host component.Host, settings component.TelemetrySettings) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAcceptEncoding returns the accept encoding to use for gRPC responses.
func (c *TransportConfig) GetAcceptEncoding() string { _ = "STUB: not implemented"; return "" }

// Config defines configuration for Coralogix exporter.
type Config struct {
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	TimeoutSettings           exporterhelper.TimeoutConfig `mapstructure:",squash"`

	// Protocol to use for communication. Options: "grpc" (default), "http"
	Protocol string `mapstructure:"protocol"`

	// Coralogix domain
	Domain string `mapstructure:"domain"`

	// Transport settings used with Domain (supports both gRPC and HTTP)
	DomainSettings TransportConfig `mapstructure:"domain_settings"`

	// Use AWS PrivateLink for the domain
	PrivateLink bool `mapstructure:"private_link"`

	// Coralogix traces ingress endpoint (supports both gRPC and HTTP)
	Traces TransportConfig `mapstructure:"traces"`

	// The Coralogix metrics ingress endpoint (supports both gRPC and HTTP)
	Metrics TransportConfig `mapstructure:"metrics"`

	// The Coralogix logs ingress endpoint (supports both gRPC and HTTP)
	Logs TransportConfig `mapstructure:"logs"`

	// The Coralogix profiles ingress endpoint (gRPC only)
	Profiles configgrpc.ClientConfig `mapstructure:"profiles"`

	// Your Coralogix private key (sensitive) for authentication
	PrivateKey configopaque.String `mapstructure:"private_key"`

	// Ordered list of Resource attributes that are used for Coralogix
	// AppName and SubSystem values. The first non-empty Resource attribute is used.
	// Example: AppNameAttributes: ["k8s.namespace.name", "service.namespace"]
	// Example: SubSystemAttributes: ["k8s.deployment.name", "k8s.daemonset.name", "service.name"]
	AppNameAttributes   []string `mapstructure:"application_name_attributes"`
	SubSystemAttributes []string `mapstructure:"subsystem_name_attributes"`
	// Default Coralogix application and subsystem name values.
	AppName   string `mapstructure:"application_name"`
	SubSystem string `mapstructure:"subsystem_name"`

	RateLimiter RateLimiterConfig `mapstructure:"rate_limiter"`
}

var _ confmap.Unmarshaler = (*Config)(nil)

// ensureHTTPScheme ensures the endpoint has an https:// scheme
func ensureHTTPScheme(endpoint string) string { _ = "STUB: not implemented"; return "" }

func (c *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// Only auto-populate profiles endpoint if protocol is not HTTP (profiles only support gRPC)

// Only set profiles endpoint fallback if protocol is not HTTP and we have something to fallback to
// This avoid validation issues

type RateLimiterConfig struct {
	Enabled   bool          `mapstructure:"enabled"`
	Threshold int           `mapstructure:"threshold"`
	Duration  time.Duration `mapstructure:"duration"`
}

func isEmpty(endpoint string) bool { _ = "STUB: not implemented"; return false }

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// validate that at least one endpoint is set up correctly

// Validate that HTTP protocol is not used with profiles

// Validate accept_encoding for gRPC protocol

// validateAcceptEncoding checks if the given encoding is supported by gRPC.
// Empty string is allowed (defaults to gzip).
func validateAcceptEncoding(encodingName string) error {
	_ = "STUB: not implemented"
	// Empty string is valid (will default to gzip)
	return nil
}

// Check if the compressor is registered in gRPC

func (c *Config) getMetadataFromResource(res pcommon.Resource) (appName, subsystem string) {
	_ = "STUB: not implemented"
	// Example application name attributes: service.namespace, k8s.namespace.name
	return "", ""
}

// Example subsystem name attributes: service.name, k8s.deployment.name, k8s.statefulset.name

func setDomainGrpcSettings(c *Config) string { _ = "STUB: not implemented"; return "" }

// If PrivateLink is enabled, use the private link endpoint.
// However, if the domain already contains "private", don't add it again.

// setMergedTransportConfigWithConf returns a TransportConfig that merges signal-specific settings with domain settings.
// Signal-specific settings take precedence over domain settings.
// This is used when unmarshaling from confmap to get domain settings.
// We pass the conf to get a fresh copy of the domain settings from the confmap.
//
// NOTE: This function modifies *Config and *TransportConfig passed as arguments.
// DO NOT ADD FURTHER USES OUTSIDE OF Unmarshal, see github.com/open-telemetry/opentelemetry-collector-contrib/issues/44731
func setMergedTransportConfigWithConf(conf *confmap.Conf, c *Config, signalConfig *TransportConfig) (TransportConfig, error) {
	_ = "STUB: not implemented"
	return *new(TransportConfig), nil
}

// setMergedTransportConfig returns a TransportConfig that merges signal-specific settings with domain settings.
// Signal-specific settings take precedence over domain settings.
//
// NOTE: This function modifies *Config and *TransportConfig passed as arguments.
// DO NOT ADD FURTHER USES OUTSIDE OF Unmarshal, see github.com/open-telemetry/opentelemetry-collector-contrib/issues/44731
func setMergedTransportConfig(c *Config, merged, signalConfig *TransportConfig) *TransportConfig {
	_ = "STUB: not implemented"
	return nil
}

// MapList.Set copies the backing array, so this does not mutate c.DomainSettings
