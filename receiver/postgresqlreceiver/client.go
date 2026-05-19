// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package postgresqlreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"

	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/config/configtls"
	"go.uber.org/zap"
)

const querySampleTraceContextKey = "_otel_trace_context"

// databaseName is a name that refers to a database so that it can be uniquely referred to later
// i.e. database1
type databaseName string

// tableIdentifier is an identifier that contains both the database and table separated by a "|"
// i.e. database1|table2
type tableIdentifier string

// indexIdentifier is a unique string that identifies a particular index and is separated by the "|" character
type indexIdentifer string

// functionIdentifier is a unique string that identifies a particular function and is separated by the "|" character
type functionIdentifer string

// errNoLastArchive is an error that occurs when there is no previous wal archive, so there is no way to compute the
// last archived point
var errNoLastArchive = errors.New("no last archive found, not able to calculate oldest WAL age")

type client interface {
	Close() error
	getDatabaseStats(ctx context.Context, databases []string) (map[databaseName]databaseStats, error)
	getDatabaseLocks(ctx context.Context) ([]databaseLocks, error)
	getBGWriterStats(ctx context.Context) (*bgStat, error)
	getBackends(ctx context.Context, databases []string) (map[databaseName]int64, error)
	getDatabaseSize(ctx context.Context, databases []string) (map[databaseName]int64, error)
	getDatabaseTableMetrics(ctx context.Context, db string) (map[tableIdentifier]tableStats, error)
	getBlocksReadByTable(ctx context.Context, db string) (map[tableIdentifier]tableIOStats, error)
	getReplicationStats(ctx context.Context) ([]replicationStats, error)
	getLatestWalAgeSeconds(ctx context.Context) (int64, error)
	getMaxConnections(ctx context.Context) (int64, error)
	getIndexStats(ctx context.Context, database string) (map[indexIdentifer]indexStat, error)
	getFunctionStats(ctx context.Context, database string) (map[functionIdentifer]functionStat, error)
	listDatabases(ctx context.Context) ([]string, error)
	getVersion(ctx context.Context) (string, error)
	getQuerySamples(ctx context.Context, limit int64, newestQueryTimestamp float64, logger *zap.Logger) ([]map[string]any, float64, error)
	getTopQuery(ctx context.Context, limit int64, logger *zap.Logger) ([]map[string]any, error)
	explainQuery(query, queryID string, logger *zap.Logger) (string, error)
}

type postgreSQLClient struct {
	client  *sql.DB
	closeFn func() error
}

// explainableStatements is a whitelist of SQL statements that PostgreSQL can EXPLAIN.
var explainableStatements = map[string]struct{}{
	"SELECT": {},
	"TABLE":  {}, // TABLE is shorthand for SELECT * FROM
	"DELETE": {},
	"INSERT": {},
	"UPDATE": {},
	"WITH":   {}, // CTEs
	"MERGE":  {}, // PostgreSQL 15+
	"VALUES": {},
}

// isExplainableQuery checks if a query can be explained by PostgreSQL.
// Uses a whitelist approach, only allows known DML statements.
func isExplainableQuery(query string) bool { _ = "STUB: not implemented"; return false }

// Remove leading comments (both -- and /* */ style)

// Extract and uppercase only the first word to check against the whitelist

// explainQuery implements client.
func (c *postgreSQLClient) explainQuery(query, queryID string, logger *zap.Logger) (string, error) {
	_ = "STUB: not implemented"
	// Check if the query is explainable before attempting EXPLAIN
	return "", nil
}

// PostgreSQL's pg_stat_statements returns queries with $1, $2 placeholders

// Build nulls array for placeholders

// if there is no parameter needed, we can not put an empty bracket

var _ client = (*postgreSQLClient)(nil)

type postgreSQLConfig struct {
	username string
	password string
	database string
	address  confignet.AddrConfig
	tls      configtls.ClientConfig
}

func sslConnectionString(tls configtls.ClientConfig) string { _ = "STUB: not implemented"; return "" }

func (c postgreSQLConfig) ConnectionString() (string, error) {
	_ = "STUB: not implemented"
	// postgres will assume the supplied user as the database name if none is provided,
	// so we must specify a database name even when we are just collecting the list of databases.
	return "", nil
}

// lib/pg expects a unix socket host to start with a "/" and appends the appropriate .s.PGSQL.port internally

func (c *postgreSQLClient) Close() error { _ = "STUB: not implemented"; return nil }

type databaseStats struct {
	transactionCommitted int64
	transactionRollback  int64
	deadlocks            int64
	tempIo               int64
	tempFiles            int64
	tupUpdated           int64
	tupReturned          int64
	tupFetched           int64
	tupInserted          int64
	tupDeleted           int64
	blksHit              int64
	blksRead             int64
}

func (c *postgreSQLClient) getDatabaseStats(ctx context.Context, databases []string) (map[databaseName]databaseStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type databaseLocks struct {
	relation string
	mode     string
	lockType string
	locks    int64
}

func (c *postgreSQLClient) getDatabaseLocks(ctx context.Context) ([]databaseLocks, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getBackends returns a map of database names to the number of active connections
func (c *postgreSQLClient) getBackends(ctx context.Context, databases []string) (map[databaseName]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postgreSQLClient) getDatabaseSize(ctx context.Context, databases []string) (map[databaseName]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tableStats contains a result for a row of the getDatabaseTableMetrics result
type tableStats struct {
	database    string
	schema      string
	table       string
	live        int64
	dead        int64
	inserts     int64
	upd         int64
	del         int64
	hotUpd      int64
	seqScans    int64
	size        int64
	vacuumCount int64
}

func (c *postgreSQLClient) getDatabaseTableMetrics(ctx context.Context, db string) (map[tableIdentifier]tableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type tableIOStats struct {
	database  string
	schema    string
	table     string
	heapRead  int64
	heapHit   int64
	idxRead   int64
	idxHit    int64
	toastRead int64
	toastHit  int64
	tidxRead  int64
	tidxHit   int64
}

func (c *postgreSQLClient) getBlocksReadByTable(ctx context.Context, db string) (map[tableIdentifier]tableIOStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type indexStat struct {
	index    string
	table    string
	schema   string
	database string
	size     int64
	scans    int64
}

func (c *postgreSQLClient) getIndexStats(ctx context.Context, database string) (map[indexIdentifer]indexStat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type functionStat struct {
	function string
	schema   string
	database string
	calls    int64
}

func (c *postgreSQLClient) getFunctionStats(ctx context.Context, database string) (map[functionIdentifer]functionStat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type bgStat struct {
	checkpointsReq       int64
	checkpointsScheduled int64
	checkpointWriteTime  float64
	checkpointSyncTime   float64
	bgWrites             int64
	bufferBackendWrites  int64
	bufferFsyncWrites    int64
	bufferCheckpoints    int64
	buffersAllocated     int64
	maxWritten           int64
}

func (c *postgreSQLClient) getBGWriterStats(ctx context.Context) (*bgStat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Not found in pg17+ tables
// Not found in pg17+ tables

func (c *postgreSQLClient) getMaxConnections(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type replicationStats struct {
	clientAddr   string
	pendingBytes int64
	flushLagInt  int64 // Deprecated
	replayLagInt int64 // Deprecated
	writeLagInt  int64 // Deprecated
	flushLag     float64
	replayLag    float64
	writeLag     float64
}

func (c *postgreSQLClient) getDeprecatedReplicationStats(ctx context.Context) ([]replicationStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postgreSQLClient) getReplicationStats(ctx context.Context) ([]replicationStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postgreSQLClient) getLatestWalAgeSeconds(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *postgreSQLClient) listDatabases(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postgreSQLClient) getVersion(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseMajorVersion(ver string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func filterQueryByDatabases(baseQuery string, databases []string, groupBy bool) string {
	_ = "STUB: not implemented"
	return ""
}

func tableKey(database, schema, table string) tableIdentifier {
	_ = "STUB: not implemented"
	return *new(tableIdentifier)
}

func indexKey(database, schema, table, index string) indexIdentifer {
	_ = "STUB: not implemented"
	return *new(indexIdentifer)
}

func functionKey(database, schema, function string) functionIdentifer {
	_ = "STUB: not implemented"
	return *new(functionIdentifer)
}

//go:embed templates/querySampleTemplate.tmpl
var querySampleTemplate string

func (c *postgreSQLClient) getQuerySamples(ctx context.Context, limit int64, newestQueryTimestamp float64, logger *zap.Logger) ([]map[string]any, float64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// in case the sql returned rows contains null value, we just log a warning and continue

// Use a background context so we don't accidentally inherit cancellation or span context
// from the scrape context; the only trace linkage should come from the extracted traceparent.

// TODO: check if the query is truncated.

func convertMillisecondToSecond(column, value string, logger *zap.Logger) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func convertToInt(column, value string, logger *zap.Logger) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

//go:embed templates/topQueryTemplate.tmpl
var topQueryTemplate string

// getTopQuery implements client.
func (c *postgreSQLClient) getTopQuery(ctx context.Context, limit int64, logger *zap.Logger) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Only get query after the oldest query we got from the previous sample query colelction.
// For instance, if from the last sample query we got queries executed between 8:00 ~ 8:15,
// in this query, we should only gather query after 8:15

// in case the sql returned rows contains null value, we just log a warning and continue

// Store raw query before obfuscation (needed for EXPLAIN with $N placeholders)

// Obfuscate query for display/logging (converts $1,$2 to ?)
// Raw query is already stored separately for EXPLAIN
