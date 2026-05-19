// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"errors"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal/metrics"
)

// Config defines configuration for clickhouse exporter.
type Config struct {
	// collectorVersion is the build version of the collector. This is overridden when an exporter is initialized.
	collectorVersion string

	TimeoutSettings           exporterhelper.TimeoutConfig `mapstructure:",squash"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`

	// Endpoint is the clickhouse endpoint.
	Endpoint string `mapstructure:"endpoint"`
	// Username is the authentication username.
	Username string `mapstructure:"username"`
	// Password is the authentication password.
	Password configopaque.String `mapstructure:"password"`
	// Database is the database name to export.
	Database string `mapstructure:"database"`
	// TLS is the TLS config for connecting to ClickHouse.
	TLS configtls.ClientConfig `mapstructure:"tls"`
	// ConnectionParams is the extra connection parameters with map format. for example compression/dial_timeout
	ConnectionParams map[string]string `mapstructure:"connection_params"`
	// LogsTableName is the table name for logs. default is `otel_logs`.
	LogsTableName string `mapstructure:"logs_table_name"`
	// TracesTableName is the table name for traces. default is `otel_traces`.
	TracesTableName string `mapstructure:"traces_table_name"`
	// MetricsTableName is the table name for metrics. default is `otel_metrics`.
	//
	// Deprecated: MetricsTableName exists for historical compatibility
	// and should not be used. To set the metrics tables name,
	// use the MetricsTables parameter instead.
	MetricsTableName string `mapstructure:"metrics_table_name"`
	// TTL is The data time-to-live example 30m, 48h. 0 means no ttl.
	TTL time.Duration `mapstructure:"ttl"`
	// TableEngine is the table engine to use. default is `MergeTree()`.
	TableEngine TableEngine `mapstructure:"table_engine"`
	// ClusterName if set will append `ON CLUSTER` with the provided name when creating tables.
	ClusterName string `mapstructure:"cluster_name"`
	// CreateSchema if set to true will run the DDL for creating the database and tables. default is true.
	CreateSchema bool `mapstructure:"create_schema"`
	// Compress controls the compression algorithm. Valid options: `none` (disabled), `zstd`, `lz4` (default), `gzip`, `deflate`, `br`, `true` (lz4).
	Compress string `mapstructure:"compress"`
	// AsyncInsert if true will enable async inserts. Default is `true`.
	// Ignored if async inserts are configured in the `endpoint` or `connection_params`.
	// Async inserts may still be overridden server-side.
	AsyncInsert bool `mapstructure:"async_insert"`
	// JSON enables the JSON column type for attributes in logs and traces tables.
	// When false (default), Map columns are used. When true, JSON columns are used.
	// ClickHouse v25+ is recommended for reliable JSON support.
	// You may also need to add `enable_json_type=1` to your endpoint or connection_params.
	JSON bool `mapstructure:"json"`
	// MetricsTables defines the table names for metric types.
	MetricsTables MetricTablesConfig `mapstructure:"metrics_tables"`
}

type MetricTablesConfig struct {
	// Gauge is the table name for gauge metric type. default is `otel_metrics_gauge`.
	Gauge metrics.MetricTypeConfig `mapstructure:"gauge"`
	// Sum is the table name for sum metric type. default is `otel_metrics_sum`.
	Sum metrics.MetricTypeConfig `mapstructure:"sum"`
	// Summary is the table name for summary metric type. default is `otel_metrics_summary`.
	Summary metrics.MetricTypeConfig `mapstructure:"summary"`
	// Histogram is the table name for histogram metric type. default is `otel_metrics_histogram`.
	Histogram metrics.MetricTypeConfig `mapstructure:"histogram"`
	// ExponentialHistogram is the table name for exponential histogram metric type. default is `otel_metrics_exponential_histogram`.
	ExponentialHistogram metrics.MetricTypeConfig `mapstructure:"exponential_histogram"`
}

// TableEngine defines the ENGINE string value when creating the table.
type TableEngine struct {
	Name   string `mapstructure:"name"`
	Params string `mapstructure:"params"`
}

const (
	defaultDatabase           = "default"
	defaultTableEngineName    = "MergeTree"
	defaultMetricTableName    = "otel_metrics"
	defaultGaugeSuffix        = "_gauge"
	defaultSumSuffix          = "_sum"
	defaultSummarySuffix      = "_summary"
	defaultHistogramSuffix    = "_histogram"
	defaultExpHistogramSuffix = "_exponential_histogram"
)

var (
	errConfigNoEndpoint      = errors.New("endpoint must be specified")
	errConfigInvalidEndpoint = errors.New("endpoint must be url format")
)

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Validate the ClickHouse server configuration.
func (cfg *Config) Validate() (err error) { _ = "STUB: not implemented"; return nil }

// Validate DSN with clickhouse driver.
// Last chance to catch invalid config.

func (cfg *Config) buildDSN() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Add connection params to query params.

// Enable TLS if scheme is https. This flag is necessary to support https connections.

// Use async_insert from config if not specified in DSN.

// Override username and password if specified in config.

func (cfg *Config) buildClickHouseOptions() (*clickhouse.Options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load TLS config if any TLS-related field is set (not just cert/key).

// shouldCreateSchema returns true if the exporter should run the DDL for creating database/tables.
func (cfg *Config) shouldCreateSchema() bool { _ = "STUB: not implemented"; return false }

func (cfg *Config) buildMetricTableNames() { _ = "STUB: not implemented"; return }

func (cfg *Config) areMetricTableNamesSet() bool { _ = "STUB: not implemented"; return false }

// tableEngineString generates the ENGINE string.
func (cfg *Config) tableEngineString() string { _ = "STUB: not implemented"; return "" }

// database returns the preferred database for creating tables and inserting data.
// The config option takes precedence over the DSN's settings.
// Falls back to default if neither are set.
// Assumes config has passed Validate.
func (cfg *Config) database() string { _ = "STUB: not implemented"; return "" }

// clusterString generates the ON CLUSTER string. Returns empty string if not set.
func (cfg *Config) clusterString() string { _ = "STUB: not implemented"; return "" }
