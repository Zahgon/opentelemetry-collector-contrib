// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver/internal/metadata"
)

const (
	scrapersKey = "scrapers"

	// GitHub Delivery Headers: https://docs.github.com/en/webhooks/webhook-events-and-payloads#delivery-headers
	defaultGitHubHookIDHeader       = "X-GitHub-Hook-ID"    // Unique identifier of the webhook.
	defaultGitHubEventHeader        = "X-GitHub-Event"      // The name of the event that triggered the delivery.
	defaultGitHubDeliveryHeader     = "X-GitHub-Delivery"   // A globally unique identifier (GUID) to identify the event.
	defaultGitHubSignature256Header = "X-Hub-Signature-256" // The HMAC hex digest of the request body; generated using the SHA-256 hash function and the secret as the HMAC key.
	defaultUserAgentHeader          = "User-Agent"          // Value always prefixed with "GitHub-Hookshot/"
)

// Config that is exposed to this github receiver through the OTEL config.yaml
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	Scrapers                       map[string]internal.Config `mapstructure:"scrapers"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	WebHook                        WebHook `mapstructure:"webhook"`
}

type WebHook struct {
	confighttp.ServerConfig `mapstructure:",squash"`       // squash ensures fields are correctly decoded in embedded struct
	Path                    string                         `mapstructure:"path"`             // path for data collection. Default is /events
	HealthPath              string                         `mapstructure:"health_path"`      // path for health check api. Default is /health_check
	RequiredHeaders         map[string]configopaque.String `mapstructure:"required_headers"` // optional setting to set one or more required headers for all requests to have (except the health check)
	GitHubHeaders           GitHubHeaders                  `mapstructure:",squash"`          // GitLab headers set by default
	Secret                  string                         `mapstructure:"secret"`           // secret for webhook
	ServiceName             string                         `mapstructure:"service_name"`
	IncludeSpanEvents       bool                           `mapstructure:"include_span_events"` // attach raw webhook event JSON as span events
}

type GitHubHeaders struct {
	Customizable map[string]string `mapstructure:","` // can be overwritten via required_headers
	Fixed        map[string]string `mapstructure:","` // are not allowed to be overwritten
}

var (
	_ component.Config    = (*Config)(nil)
	_ confmap.Unmarshaler = (*Config)(nil)

	errMissingEndpointFromConfig   = errors.New("missing receiver server endpoint from config")
	errReadTimeoutExceedsMaxValue  = errors.New("the duration specified for read_timeout exceeds the maximum allowed value of 10s")
	errWriteTimeoutExceedsMaxValue = errors.New("the duration specified for write_timeout exceeds the maximum allowed value of 10s")
	errRequiredHeader              = errors.New("both key and value are required to assign a required_header")
	errRequireOneScraper           = errors.New("must specify at least one scraper")
	errGitHubHeader                = errors.New("github default headers [X-GitHub-Event, X-GitHub-Delivery, X-GitHub-Hook-ID, X-Hub-Signature-256] cannot be configured")
)

// Validate the configuration passed through the OTEL config.yaml
func (cfg *Config) Validate() error {
	_ = "STUB: not implemented"

	// For now, scrapers are required to be defined in the config. As tracing
	// and other signals are added, this requirement will change.
	return nil
}

// Unmarshal a config.Parser into the config struct.
func (cfg *Config) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// load the non-dynamic config normally

// dynamically load the individual collector configs based on the key name
