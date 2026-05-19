// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlserverreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver/internal/metadata"
)

type QuerySample struct {
	MaxRowsPerQuery uint64 `mapstructure:"max_rows_per_query"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type TopQueryCollection struct {
	// Enabled enables the collection of the top queries by the execution time.
	// It will collect the top N queries based on totalElapsedTimeDiffs during the last collection interval.
	// The query statement will also be reported, hence, it is not ideal to send it as a metric. Hence
	// we are reporting them as logs.
	// The `N` is configured via `TopQueryCount`
	LookbackTime        time.Duration `mapstructure:"lookback_time"`
	MaxQuerySampleCount uint          `mapstructure:"max_query_sample_count"`
	TopQueryCount       uint          `mapstructure:"top_query_count"`
	CollectionInterval  time.Duration `mapstructure:"collection_interval"`
}

// Config defines configuration for a sqlserver receiver.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	metadata.LogsBuilderConfig     `mapstructure:",squash"`
	// EnableTopQueryCollection enables the collection of the top queries by the execution time.
	// It will collect the top N queries based on totalElapsedTimeDiffs during the last collection interval.
	// The query statement will also be reported, hence, it is not ideal to send it as a metric. Hence
	// we are reporting them as logs.
	// The `N` is configured via `TopQueryCount`
	TopQueryCollection `mapstructure:"top_query_collection"`

	QuerySample `mapstructure:"query_sample_collection"`

	InstanceName string `mapstructure:"instance_name"`
	ComputerName string `mapstructure:"computer_name"`

	DataSource string `mapstructure:"datasource"`

	Password configopaque.String `mapstructure:"password"`
	Port     uint                `mapstructure:"port"`
	Server   string              `mapstructure:"server"`
	Username string              `mapstructure:"username"`

	// Flag to check if the connection is direct or not. It should only be
	// used after a successful call to the `Validate` method.
	isDirectDBConnectionEnabled bool
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func directDBConnectionEnabled(config *Config) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If no connection information is provided, we can't connect directly and this is a valid config.

// It is a valid direct connection configuration

func (cfg *Config) EffectiveLookbackTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
