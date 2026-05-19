// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlquery // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	Driver                         string              `mapstructure:"driver"`
	DataSource                     string              `mapstructure:"datasource"`
	Host                           string              `mapstructure:"host"`
	Port                           int                 `mapstructure:"port"`
	Database                       string              `mapstructure:"database"`
	Username                       string              `mapstructure:"username"`
	Password                       configopaque.String `mapstructure:"password"`
	AdditionalParams               map[string]any      `mapstructure:"additional_params"`
	Queries                        []Query             `mapstructure:"queries"`
	StorageID                      *component.ID       `mapstructure:"storage"`
	Telemetry                      TelemetryConfig     `mapstructure:"telemetry"`
}

func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }

// check if driver is supported

// If datasource is set, none of the individual connection parameters should be set

// If datasource is not set, host, port, and database are required

// For sqlserver, port is optional

type Query struct {
	SQL                string      `mapstructure:"sql"`
	Metrics            []MetricCfg `mapstructure:"metrics"`
	Logs               []LogsCfg   `mapstructure:"logs"`
	TrackingColumn     string      `mapstructure:"tracking_column"`
	TrackingStartValue string      `mapstructure:"tracking_start_value"`
	IgnoreNullValues   bool        `mapstructure:"ignore_null_values"`
}

func (q Query) Validate() error { _ = "STUB: not implemented"; return nil }

type LogsCfg struct {
	BodyColumn       string   `mapstructure:"body_column"`
	AttributeColumns []string `mapstructure:"attribute_columns"`
}

func (config LogsCfg) Validate() error { _ = "STUB: not implemented"; return nil }

// RowCondition filters query result rows for a metric. Only rows where the
// specified column equals the specified value are used to produce the metric.
// This is useful when a single query returns a pivot-style result set (e.g.
// pgbouncer's SHOW LISTS) where each row represents a different metric and
// must be selected individually.
type RowCondition struct {
	Column string `mapstructure:"column"`
	Value  string `mapstructure:"value"`
}

type MetricCfg struct {
	MetricName       string            `mapstructure:"metric_name"`
	ValueColumn      string            `mapstructure:"value_column"`
	AttributeColumns []string          `mapstructure:"attribute_columns"`
	Monotonic        bool              `mapstructure:"monotonic"`
	ValueType        MetricValueType   `mapstructure:"value_type"`
	DataType         MetricType        `mapstructure:"data_type"`
	Aggregation      MetricAggregation `mapstructure:"aggregation"`
	Unit             string            `mapstructure:"unit"`
	Description      string            `mapstructure:"description"`
	StaticAttributes map[string]string `mapstructure:"static_attributes"`
	StartTsColumn    string            `mapstructure:"start_ts_column"`
	TsColumn         string            `mapstructure:"ts_column"`
	RowCondition     *RowCondition     `mapstructure:"row_condition"`
}

func (c MetricCfg) Validate() error { _ = "STUB: not implemented"; return nil }

type MetricType string

const (
	MetricTypeUnspecified MetricType = ""
	MetricTypeGauge       MetricType = "gauge"
	MetricTypeSum         MetricType = "sum"
)

func (t MetricType) Validate() error { _ = "STUB: not implemented"; return nil }

type MetricValueType string

const (
	MetricValueTypeUnspecified MetricValueType = ""
	MetricValueTypeInt         MetricValueType = "int"
	MetricValueTypeDouble      MetricValueType = "double"
)

func (t MetricValueType) Validate() error { _ = "STUB: not implemented"; return nil }

type MetricAggregation string

const (
	MetricAggregationUnspecified MetricAggregation = ""
	MetricAggregationCumulative  MetricAggregation = "cumulative"
	MetricAggregationDelta       MetricAggregation = "delta"
)

func (a MetricAggregation) Validate() error { _ = "STUB: not implemented"; return nil }

type TelemetryConfig struct {
	Logs TelemetryLogsConfig `mapstructure:"logs"`
}

type TelemetryLogsConfig struct {
	Query bool `mapstructure:"query"`
}
