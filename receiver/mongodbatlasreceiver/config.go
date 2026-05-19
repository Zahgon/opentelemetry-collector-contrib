// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/metadata"
)

var _ component.Config = (*Config)(nil)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	BaseURL                        string                                `mapstructure:"base_url"`
	PublicKey                      string                                `mapstructure:"public_key"`
	PrivateKey                     configopaque.String                   `mapstructure:"private_key"`
	Granularity                    string                                `mapstructure:"granularity"`
	MetricsBuilderConfig           metadata.MetricsBuilderConfig         `mapstructure:",squash"`
	Projects                       []ProjectConfig                       `mapstructure:"projects"`
	Alerts                         AlertConfig                           `mapstructure:"alerts"`
	Events                         configoptional.Optional[EventsConfig] `mapstructure:"events"`
	Logs                           LogConfig                             `mapstructure:"logs"`
	BackOffConfig                  configretry.BackOffConfig             `mapstructure:"retry_on_failure"`
	StorageID                      *component.ID                         `mapstructure:"storage"`
}

type AlertConfig struct {
	Enabled  bool                    `mapstructure:"enabled"`
	Endpoint string                  `mapstructure:"endpoint"`
	Secret   configopaque.String     `mapstructure:"secret"`
	TLS      *configtls.ServerConfig `mapstructure:"tls"`
	Mode     string                  `mapstructure:"mode"`

	// these parameters are only relevant in retrieval mode
	Projects     []*ProjectConfig `mapstructure:"projects"`
	PollInterval time.Duration    `mapstructure:"poll_interval"`
	PageSize     int64            `mapstructure:"page_size"`
	MaxPages     int64            `mapstructure:"max_pages"`
}

type LogConfig struct {
	Enabled  bool                 `mapstructure:"enabled"`
	Projects []*LogsProjectConfig `mapstructure:"projects"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// EventsConfig is the configuration options for events collection
type EventsConfig struct {
	Projects      []*ProjectConfig `mapstructure:"projects"`
	Organizations []*OrgConfig     `mapstructure:"organizations"`
	PollInterval  time.Duration    `mapstructure:"poll_interval"`
	Types         []string         `mapstructure:"types"`
	PageSize      int64            `mapstructure:"page_size"`
	MaxPages      int64            `mapstructure:"max_pages"`
}

type LogsProjectConfig struct {
	ProjectConfig `mapstructure:",squash"`

	EnableAuditLogs bool              `mapstructure:"collect_audit_logs"`
	EnableHostLogs  *bool             `mapstructure:"collect_host_logs"`
	AccessLogs      *AccessLogsConfig `mapstructure:"access_logs"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type AccessLogsConfig struct {
	Enabled      *bool         `mapstructure:"enabled"`
	PollInterval time.Duration `mapstructure:"poll_interval"`
	PageSize     int64         `mapstructure:"page_size"`
	MaxPages     int64         `mapstructure:"max_pages"`
	AuthResult   *bool         `mapstructure:"auth_result"`
}

func (alc *AccessLogsConfig) IsEnabled() bool { _ = "STUB: not implemented"; return false }

type ProjectConfig struct {
	Name            string   `mapstructure:"name"`
	ExcludeClusters []string `mapstructure:"exclude_clusters"`
	IncludeClusters []string `mapstructure:"include_clusters"`

	includesByClusterName map[string]struct{}
	excludesByClusterName map[string]struct{}
}

type OrgConfig struct {
	ID string `mapstructure:"id"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func (pc *ProjectConfig) populateIncludesAndExcludes() { _ = "STUB: not implemented"; return }

var (
	// Alerts Receiver Errors
	errNoEndpoint        = errors.New("an endpoint must be specified")
	errNoSecret          = errors.New("a webhook secret must be specified")
	errNoCert            = errors.New("tls was configured, but no cert file was specified")
	errNoKey             = errors.New("tls was configured, but no key file was specified")
	errNoModeRecognized  = fmt.Errorf("alert mode not recognized for mode. Known alert modes are: %s,%s", alertModeListen, alertModePoll)
	errPageSizeIncorrect = errors.New("page size must be a value between 1 and 500")

	// Logs Receiver Errors
	errNoProjects    = errors.New("at least one 'project' must be specified")
	errNoEvents      = errors.New("at least one 'project' or 'organizations' event type must be specified")
	errClusterConfig = errors.New("only one of 'include_clusters' or 'exclude_clusters' may be specified")

	// Access Logs Errors
	errMaxPageSize = errors.New("the maximum value for 'page_size' is 20000")

	// Config Errors
	errConfigEmptyEndpoint = errors.New("baseurl must not be empty")
)

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (l *LogConfig) validate() error { _ = "STUB: not implemented"; return nil }

func (a *AlertConfig) validate() error {
	_ = "STUB: not implemented"

	// No need to further validate, receiving alerts is disabled.
	return nil
}

func (a AlertConfig) validatePollConfig() error { _ = "STUB: not implemented"; return nil }

// based off API limits https://www.mongodb.com/docs/atlas/reference/api/alerts-get-all-alerts/

func (a AlertConfig) validateListenConfig() error { _ = "STUB: not implemented"; return nil }

func (e EventsConfig) validate() error { _ = "STUB: not implemented"; return nil }

func validateEndpoint(endpoint string) error { _ = "STUB: not implemented"; return nil }
