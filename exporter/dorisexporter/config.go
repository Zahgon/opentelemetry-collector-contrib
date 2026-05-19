// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"time"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

type Config struct {
	// confighttp.ClientConfig.Headers is the headers of doris stream load.
	confighttp.ClientConfig   `mapstructure:",squash"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`

	// TableNames is the table name for logs, traces and metrics.
	Table `mapstructure:"table"`

	// Database is the database name.
	Database string `mapstructure:"database"`
	// Username is the authentication username.
	Username string `mapstructure:"username"`
	// Password is the authentication password.
	Password configopaque.String `mapstructure:"password"`
	// CreateSchema is whether databases and tables are created automatically.
	CreateSchema bool `mapstructure:"create_schema"`
	// MySQLEndpoint is the mysql protocol address to create the schema; ignored if create_schema is false.
	MySQLEndpoint string `mapstructure:"mysql_endpoint"`
	// Data older than these days will be deleted; ignored if create_schema is false. If set to 0, historical data will not be deleted.
	HistoryDays int32 `mapstructure:"history_days"`
	// The number of days in the history partition that was created when the table was created; ignored if create_schema is false.
	// If history_days is not 0, create_history_days needs to be less than or equal to history_days.
	CreateHistoryDays int32 `mapstructure:"create_history_days"`
	// ReplicationNum is the number of replicas of the table; ignored if create_schema is false.
	ReplicationNum int32 `mapstructure:"replication_num"`
	// Timezone is the timezone of the doris.
	TimeZone string `mapstructure:"timezone"`
	// LogResponse is whether to log the response of doris stream load.
	LogResponse bool `mapstructure:"log_response"`
	// LabelPrefix is the prefix of the label in doris stream load.
	LabelPrefix string `mapstructure:"label_prefix"`
	// ProgressInterval is the interval of the progress reporter.
	LogProgressInterval int `mapstructure:"log_progress_interval"`

	// not in config file, will be set in Validate
	timeLocation *time.Location `mapstructure:"-"`
}

type Table struct {
	// Logs is the table name for logs.
	Logs string `mapstructure:"logs"`
	// Traces is the table name for traces.
	Traces string `mapstructure:"traces"`
	// Metrics is the table name for metrics.
	Metrics string `mapstructure:"metrics"`
}

func (cfg *Config) Validate() (err error) { _ = "STUB: not implemented"; return nil }

// Preventing SQL Injection Attacks

const (
	defaultStart = -2147483648 // IntMin
)

func (cfg *Config) startHistoryDays() int32 { _ = "STUB: not implemented"; return 0 }

const (
	properties = `
PROPERTIES (
"replication_num" = "%d",
"compaction_policy" = "%s",
"dynamic_partition.enable" = "true",
"dynamic_partition.create_history_partition" = "true",
"dynamic_partition.time_unit" = "DAY",
"dynamic_partition.start" = "%d",
"dynamic_partition.history_partition_num" = "%d",
"dynamic_partition.end" = "1",
"dynamic_partition.prefix" = "p",
"compression" = "zstd",
"inverted_index_storage_format" = "V2"
)
`
)

const (
	compactionPolicySizeBased  = "size_based"
	compactionPolicyTimeSeries = "time_series"
)

// // propertiesStr returns the properties string for non-unique key tables.
func (cfg *Config) propertiesStr() string { _ = "STUB: not implemented"; return "" }

// // propertiesStrForUniqueKey returns the properties string for unique key tables.
func (cfg *Config) propertiesStrForUniqueKey() string { _ = "STUB: not implemented"; return "" }
