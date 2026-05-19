// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlquery // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"

import (
	"context"
	"database/sql"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.uber.org/zap"
)

type SQLOpenerFunc func(driverName, dataSourceName string) (*sql.DB, error)

type DbProviderFunc func() (*sql.DB, error)

type ClientProviderFunc func(Db, string, *zap.Logger, TelemetryConfig) DbClient

type Scraper struct {
	id                   component.ID
	Query                Query
	ScrapeCfg            scraperhelper.ControllerConfig
	StartTime            pcommon.Timestamp
	ClientProviderFunc   ClientProviderFunc
	DbProviderFunc       DbProviderFunc
	Logger               *zap.Logger
	Telemetry            TelemetryConfig
	Client               DbClient
	Db                   *sql.DB
	InstrumentationScope pcommon.InstrumentationScope
}

var _ scraper.Metrics = (*Scraper)(nil)

func NewScraper(id component.ID, query Query, scrapeCfg scraperhelper.ControllerConfig, logger *zap.Logger, telemetry TelemetryConfig, dbProviderFunc DbProviderFunc, clientProviderFunc ClientProviderFunc, instrumentationScope pcommon.InstrumentationScope) *Scraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scraper) ID() component.ID { _ = "STUB: not implemented"; return *new(component.ID) }

func (s *Scraper) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scraper) ScrapeMetrics(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *Scraper) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func BuildDataSourceString(config Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MySQL doesn't need URL escaping

// ClickHouse connection string format: clickhouse://user:pass@host:port/db?param1=value1&param2=value2

// HDB connection string format: hdb://user:pass@host:port?param1=value1

// MySQL connection string format: user:pass@tcp(host:port)/db?param1=value1&param2=value2

// Oracle connection string format: oracle://user:pass@host:port/service_name?param1=value1&param2=value2

// PostgreSQL connection string format: postgresql://user:pass@host:port/db?param1=value1&param2=value2

// Snowflake connection string format: user:pass@host:port/database?param1=value1&param2=value2

// SQL Server connection string format: sqlserver://username:password@host:port/instance

// replace all backslashes with forward slashes

// if host contains a "/", split it into hostname and instance

// TDS connection string format: tds://user:pass@host:port/database

// Append query parameters if any exist
