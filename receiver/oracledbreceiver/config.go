// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oracledbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver/internal/metadata"
)

var (
	errBadDataSource       = errors.New("datasource is invalid")
	errBadEndpoint         = errors.New("endpoint must be specified as host:port")
	errBadPort             = errors.New("invalid port in endpoint")
	errEmptyEndpoint       = errors.New("endpoint must be specified")
	errEmptyPassword       = errors.New("password must be set")
	errEmptyService        = errors.New("service must be specified")
	errEmptyUsername       = errors.New("username must be set")
	errMaxQuerySampleCount = errors.New("`max_query_sample_count` must be between 1 and 10000")
	errTopQueryCount       = errors.New("`top_query_count` must be between 1 and 200 and less than or equal to `max_query_sample_count`")
)

type TopQueryCollection struct {
	MaxQuerySampleCount uint          `mapstructure:"max_query_sample_count"`
	TopQueryCount       uint          `mapstructure:"top_query_count"`
	CollectionInterval  time.Duration `mapstructure:"collection_interval"`
}

type QuerySample struct {
	MaxRowsPerQuery uint64 `mapstructure:"max_rows_per_query"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type Config struct {
	DataSource                     string `mapstructure:"datasource"`
	Endpoint                       string `mapstructure:"endpoint"`
	Password                       string `mapstructure:"password"`
	Service                        string `mapstructure:"service"`
	Username                       string `mapstructure:"username"`
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	metadata.LogsBuilderConfig     `mapstructure:",squash"`

	TopQueryCollection `mapstructure:"top_query_collection"`
	QuerySample        `mapstructure:"query_sample_collection"`
}

func (c Config) Validate() error {
	_ = "STUB: not implemented"

	// If DataSource is defined it takes precedence over the rest of the connection options.
	return nil
}
