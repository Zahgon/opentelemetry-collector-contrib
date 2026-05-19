// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gitlabreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/gitlabreceiver"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/confmap"
)

const (
	defaultReadTimeout  = 500 * time.Millisecond
	defaultWriteTimeout = 500 * time.Millisecond

	defaultEndpoint = "localhost:8080"

	defaultPath       = "/events"
	defaultHealthPath = "/health"

	// GitLab default headers: https://docs.gitlab.com/ee/user/project/integrations/webhooks.html#delivery-headers
	defaultUserAgentHeader         = "User-Agent"
	defaultGitLabInstanceHeader    = "X-Gitlab-Instance"
	defaultGitLabWebhookUUIDHeader = "X-Gitlab-Webhook-UUID"
	defaultGitLabEventHeader       = "X-Gitlab-Event"
	defaultGitLabEventUUIDHeader   = "X-Gitlab-Event-UUID"
	defaultIdempotencyKeyHeader    = "Idempotency-Key"
	// #nosec G101 - Not an actual secret, just the name of a header: https://docs.gitlab.com/user/project/integrations/webhooks/#create-a-webhook
	defaultGitLabSecretTokenHeader = "X-Gitlab-Token"
)

var (
	errReadTimeoutExceedsMaxValue  = errors.New("the duration specified for read_timeout exceeds the maximum allowed value of 10s")
	errWriteTimeoutExceedsMaxValue = errors.New("the duration specified for write_timeout exceeds the maximum allowed value of 10s")
	errRequiredHeader              = errors.New("both key and value are required to assign a required_header")
	errGitlabHeader                = errors.New("gitlab default headers [X-Gitlab-Webhook-UUID, X-Gitlab-Event, X-Gitlab-Event-UUID, Idempotency-Key] cannot be configured")
	errConfigNotValid              = errors.New("configuration is not valid for the gitlab receiver")
)

// Config that is exposed to this gitlab receiver through the OTEL config.yaml
type Config struct {
	WebHook WebHook `mapstructure:"webhook"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type WebHook struct {
	confighttp.ServerConfig `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct

	Path       string `mapstructure:"path"`        // path for data collection. default is /events
	HealthPath string `mapstructure:"health_path"` // path for health check api. default is /health_check

	RequiredHeaders map[string]configopaque.String `mapstructure:"required_headers"` // optional setting to set one or more required headers for all requests to have (except the health check)
	GitlabHeaders   GitlabHeaders                  `mapstructure:",squash"`          // GitLab headers set by default

	Secret string `mapstructure:"secret"` // secret for webhook

	// IncludeUserAttributes controls whether user information (commit author, pipeline actor) is included
	// Default: false (user information is excluded by default for privacy)
	IncludeUserAttributes bool `mapstructure:"include_user_attributes"`
}

type GitlabHeaders struct {
	Customizable map[string]string `mapstructure:","` // can be overwritten via required_headers
	Fixed        map[string]string `mapstructure:","` // are not allowed to be overwritten

	// prevent unkeyed literal initialization
	_ struct{}
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// load the non-dynamic config normally

// overwrite customizable GitLab default headers if configured within the required_headers
