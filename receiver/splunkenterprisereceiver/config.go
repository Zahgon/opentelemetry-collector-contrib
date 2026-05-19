// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkenterprisereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver/internal/metadata"
)

var (
	errBadOrMissingEndpoint = errors.New("missing a valid endpoint")
	errBadScheme            = errors.New("endpoint scheme must be either http or https")
	errMissingAuthExtension = errors.New("auth extension missing from config")

	errEmptySPL         = errors.New("search spl cannot be empty")
	errInvalidTarget    = errors.New("search target must be one of: indexer, search_head, cluster_master")
	errNoMetrics        = errors.New("search must have at least one metric defined")
	errEmptyMetricName  = errors.New("metric_name cannot be empty")
	errEmptyValueColumn = errors.New("value_column cannot be empty")
	errInvalidValueType = errors.New("value_type must be one of: int, double")
)

const (
	TargetIndexer       = "indexer"
	TargetSearchHead    = "search_head"
	TargetClusterMaster = "cluster_master"

	MetricValueTypeInt    = "int"
	MetricValueTypeDouble = "double"
)

type MetricConfig struct {
	MetricName       string            `mapstructure:"metric_name"`
	ValueColumn      string            `mapstructure:"value_column"`
	AttributeColumns []string          `mapstructure:"attribute_columns"`
	ValueType        string            `mapstructure:"value_type"`
	Unit             string            `mapstructure:"unit"`
	Description      string            `mapstructure:"description"`
	StaticAttributes map[string]string `mapstructure:"static_attributes"`
}

func (m MetricConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type SearchConfig struct {
	SPL      string         `mapstructure:"spl"`
	Target   string         `mapstructure:"target"`
	Earliest string         `mapstructure:"earliest"`
	Latest   string         `mapstructure:"latest"`
	Metrics  []MetricConfig `mapstructure:"metrics"`
}

func (s SearchConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (s SearchConfig) TargetType() string { _ = "STUB: not implemented"; return "" }

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	IdxEndpoint                    confighttp.ClientConfig `mapstructure:"indexer"`
	SHEndpoint                     confighttp.ClientConfig `mapstructure:"search_head"`
	CMEndpoint                     confighttp.ClientConfig `mapstructure:"cluster_master"`
	VersionInfo                    bool                    `mapstructure:"build_version_info"`
	Searches                       []SearchConfig          `mapstructure:"searches"`
}

func (cfg *Config) Validate() (errors error) { _ = "STUB: not implemented"; return nil }

// if no endpoint is set we do not start the receiver. For each set endpoint we go through and Validate
// that it contains an auth setting and a valid endpoint, if its missing either of these the receiver will
// fail to start.

// note passes for both http and https
