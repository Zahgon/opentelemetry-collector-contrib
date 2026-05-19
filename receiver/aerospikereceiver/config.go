// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aerospikereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver/internal/metadata"
)

var (
	errBadEndpoint          = errors.New("endpoint must be specified as host:port")
	errBadPort              = errors.New("invalid port in endpoint")
	errEmptyEndpoint        = errors.New("endpoint must be specified")
	errEmptyEndpointTLSName = errors.New("endpoint TLSName must be specified")
	errEmptyPassword        = errors.New("password must be set if username is set")
	errEmptyUsername        = errors.New("username must be set if password is set")
	errNegativeTimeout      = errors.New("timeout must be non-negative")
	errFailedTLSLoad        = errors.New("failed to load TLS config")
)

// Config is the receiver configuration
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	Endpoint                       string                        `mapstructure:"endpoint"`
	TLSName                        string                        `mapstructure:"tlsname"`
	Username                       string                        `mapstructure:"username"`
	Password                       configopaque.String           `mapstructure:"password"`
	CollectClusterMetrics          bool                          `mapstructure:"collect_cluster_metrics"`
	Timeout                        time.Duration                 `mapstructure:"timeout"`
	MetricsBuilderConfig           metadata.MetricsBuilderConfig `mapstructure:",squash"`
	TLS                            *configtls.ClientConfig       `mapstructure:"tls,omitempty"`
}

// Validate validates the values of the given Config, and returns an error if validation fails
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
