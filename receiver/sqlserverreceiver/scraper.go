// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlserverreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"

import (
	"context"
	"database/sql"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver/internal/metadata"
)

const (
	computerNameKey = "computer_name"
	databaseNameKey = "database_name"
	instanceNameKey = "sql_instance"
)

type sqlServerScraperHelper struct {
	id                     component.ID
	config                 *Config
	sqlQuery               string
	instanceName           string
	clientProviderFunc     sqlquery.ClientProviderFunc
	dbProviderFunc         sqlquery.DbProviderFunc
	logger                 *zap.Logger
	telemetry              sqlquery.TelemetryConfig
	client                 sqlquery.DbClient
	db                     *sql.DB
	mb                     *metadata.MetricsBuilder
	lb                     *metadata.LogsBuilder
	cache                  *lru.Cache[string, int64]
	lastExecutionTimestamp time.Time
	obfuscator             *obfuscator
	serviceInstanceID      string
}

var (
	_ scraper.Metrics = (*sqlServerScraperHelper)(nil)
	_ scraper.Logs    = (*sqlServerScraperHelper)(nil)
)

func newSQLServerScraper(id component.ID,
	query string,
	telemetry sqlquery.TelemetryConfig,
	dbProviderFunc sqlquery.DbProviderFunc,
	clientProviderFunc sqlquery.ClientProviderFunc,
	params receiver.Settings,
	cfg *Config,
	cache *lru.Cache[string, int64],
) *sqlServerScraperHelper {
	_ = "STUB: not implemented"
	// Compute service instance ID
	return nil
}

func (s *sqlServerScraperHelper) ID() component.ID {
	_ = "STUB: not implemented"
	return *new(component.ID)
}

func (s *sqlServerScraperHelper) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlServerScraperHelper) ScrapeMetrics(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *sqlServerScraperHelper) ScrapeLogs(ctx context.Context) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func sanitizeQuerySampleOptionalAttributes(logs plog.Logs) { _ = "STUB: not implemented"; return }

func parseWaitResource(waitResource string) (resourceType, resourceID string) {
	_ = "STUB: not implemented"
	return "", ""
}

func splitTwoSegments(s string) (first, second string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func splitAfterFirstSegment(s string) (tail string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func isDigits(s string) bool { _ = "STUB: not implemented"; return false }

func isTwoNumericSegments(s string) bool { _ = "STUB: not implemented"; return false }

func isThreeNumericSegments(s string) bool { _ = "STUB: not implemented"; return false }

func (s *sqlServerScraperHelper) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// setupResourceBuilder configures common resource attributes for metrics and logs.
func (s *sqlServerScraperHelper) setupResourceBuilder(rb *metadata.ResourceBuilder, row sqlquery.StringMap) *metadata.ResourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlServerScraperHelper) recordDatabaseIOMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlServerScraperHelper) recordDatabasePerfCounterMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Constants are the columns for metrics from query

func (s *sqlServerScraperHelper) recordDatabaseStatusMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Constants are the column names of the database status
	return nil
}

func (s *sqlServerScraperHelper) recordDatabaseWaitMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Constants are the columns for metrics from query
	return nil
}

// The value is divided here because it's stored in SQL Server in ms, need to convert to s

func (s *sqlServerScraperHelper) recordDatabaseQueryTextAndPlan(ctx context.Context) (pcommon.Resource, error) {
	_ = "STUB: not implemented"
	// Constants are the column names of the database status
	return *new(pcommon.Resource), nil
}

// the time returned from mssql is in microsecond

// the time returned from mssql is in microsecond

// stored procedure columns

// defaulted to '0' if not present

// we're trying to get the queries that used the most time.
// caching the total elapsed time (in microsecond) and compare in the next scrape.

// sort the rows based on the totalElapsedTimeDiffs in descending order,
// only report first T(T=topQueryCount) rows.

// sort the totalElapsedTimeDiffs in descending order as well

// reporting human-readable query hash and query hash plan

func (s *sqlServerScraperHelper) retrieveValue(
	row sqlquery.StringMap,
	column string,
	errs *[]error,
	valueRetriever func(sqlquery.StringMap, string) (any, error),
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// cacheAndDiff store row(in int) with query hash and query plan hash variables
// (1) returns true if the key is cached before
// (2) returns positive value if the value is larger than the cached value
func (s *sqlServerScraperHelper) cacheAndDiff(queryHash, queryPlanHash, procedureID, column string, val int64) (bool, int64) {
	_ = "STUB: not implemented"
	return false, 0
}

// procedureID is '0' when not a stored procedure

// sortRows sorts the rows based on the `values` slice in descending order and return the first M(M=maximum) rows
// Input: (row: [row1, row2, row3], values: [100, 10, 1000], maximum: 2
// Expected Output: (row: [row3, row1]
func sortRows(rows []sqlquery.StringMap, values []int64, maximum uint) []sqlquery.StringMap {
	_ = "STUB: not implemented"
	return nil
}

func retrieveInt(row sqlquery.StringMap, columnName string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// SQL Server stores large integers in scientific e notation
// (eg 123456 is stored as 1.23456e+5)
// This value cannot be parsed by strconv.ParseInt, but is successfully
// parsed by strconv.ParseFloat. The goal is here to convert to int
// even if the stored value is in scientific e notation.

func retrieveIntAndConvert(convert func(int64) any) func(row sqlquery.StringMap, columnName string) (any, error) {
	_ = "STUB: not implemented"
	return nil
}

// need to convert even if it failed

func retrieveFloat(row sqlquery.StringMap, columnName string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *sqlServerScraperHelper) recordDatabaseSampleQuery(ctx context.Context) (pcommon.Resource, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), nil
}

// stored procedure columns

// in case the sql returned rows contains null value, we just log a warning and continue

// client.address: use host_name if it has value, if not, use client_net_address.
// this value may not be accurate if
// - there is proxy in the middle of sql client and sql server. Or
// - host_name value is empty or not accurate.
