// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mysqlreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"

import (
	"database/sql"
	_ "embed"
	"regexp"
	"time"

	// registers the mysql driver

	"github.com/hashicorp/go-version"
	"go.uber.org/zap"
)

// dbProduct identifies the database product (MySQL or MariaDB).
type dbProduct int

const (
	dbProductMySQL dbProduct = iota
	dbProductMariaDB
)

// minMySQLReplicaStatusVersion is the MySQL version at which SHOW REPLICA STATUS
// replaced SHOW SLAVE STATUS. Initialized at package load; panics on bad literal.
var minMySQLReplicaStatusVersion = version.Must(version.NewVersion("8.0.22"))

// dbVersion holds the parsed database version and product identity.
// Capability predicates keep version-specific branching out of callers.
type dbVersion struct {
	product dbProduct
	version *version.Version
}

// isValid reports whether a version was successfully detected.
func (v dbVersion) isValid() bool { _ = "STUB: not implemented"; return false }

// productString returns a human-readable product name for logging.
func (v dbVersion) productString() string { _ = "STUB: not implemented"; return "" }

// supportsQuerySampleText reports whether the server's
// performance_schema.events_statements_summary_by_digest table includes the
// query_sample_text column, introduced in MySQL 8.0.3. Absent on MySQL <8.0.3
// and all MariaDB versions.
var minMySQLQuerySampleTextVersion = version.Must(version.NewVersion("8.0.3"))

func (v dbVersion) supportsQuerySampleText() bool { _ = "STUB: not implemented"; return false }

// supportsReplicaStatus reports whether the server uses SHOW REPLICA STATUS
// (MySQL 8.0.22+). Older MySQL versions and all MariaDB versions use the
// deprecated SHOW SLAVE STATUS syntax.
func (v dbVersion) supportsReplicaStatus() bool { _ = "STUB: not implemented"; return false }

// supportsProcesslist reports whether performance_schema.processlist is available
// (MySQL 8.0.22+). This table exposes HOST as "host:port" for TCP/IP connections,
// enabling population of client.port and network.peer.port on query sample events.
//
// The only alternative source that exposes host:port is information_schema.PROCESSLIST,
// but it is not used: its implementation iterates active threads while holding a global
// mutex (same as SHOW PROCESSLIST), which has negative performance consequences on busy
// systems. It was also deprecated in MySQL 8.0, removed in MySQL 9.0, and was already
// removed from this receiver in a prior change. performance_schema.processlist is
// lock-free and reads directly from Performance Schema data structures.
//
// As a result, client.port and network.peer.port remain 0 on MariaDB (all versions)
// and MySQL < 8.0.22, where this table is not available.
func (v dbVersion) supportsProcesslist() bool { _ = "STUB: not implemented"; return false }

type client interface {
	Connect() error
	getDBVersion() dbVersion
	getGlobalStats() (map[string]string, error)
	getInnodbStats() (map[string]string, error)
	getTableStats() ([]tableStats, error)
	getTableIoWaitsStats() ([]tableIoWaitsStats, error)
	getIndexIoWaitsStats() ([]indexIoWaitsStats, error)
	getStatementEventsStats() ([]statementEventStats, error)
	getTableLockWaitEventStats() ([]tableLockWaitEventStats, error)
	// supportsReplicaStatus controls whether SHOW REPLICA STATUS or the
	// deprecated SHOW SLAVE STATUS is used.
	getReplicaStatusStats(supportsReplicaStatus bool) ([]replicaStatusStats, error)
	// supportsSampleText controls top-query template selection (MySQL 8+ vs fallback).
	getTopQueries(topN, lookback uint64, supportsSampleText bool) ([]topQuery, error)
	// supportsProcesslist controls whether the performance_schema.processlist JOIN
	// is included to populate client.port and network.peer.port (MySQL 8.0.22+).
	getQuerySamples(limit uint64, supportsProcesslist bool) ([]querySample, error)
	explainQuery(digestText, sampleStatement, schema, digest string, logger *zap.Logger) string
	Close() error
}

type mySQLClient struct {
	connStr                        string
	client                         *sql.DB
	statementEventsDigestTextLimit int
	statementEventsLimit           int
	statementEventsTimeLimit       time.Duration
	dbVersion                      dbVersion
}

type ioWaitsStats struct {
	schema      string
	name        string
	countDelete int64
	countFetch  int64
	countInsert int64
	countUpdate int64
	timeDelete  int64
	timeFetch   int64
	timeInsert  int64
	timeUpdate  int64
}

type tableIoWaitsStats struct {
	ioWaitsStats
}

type indexIoWaitsStats struct {
	ioWaitsStats
	index string
}

type tableStats struct {
	schema           string
	name             string
	rows             int64
	averageRowLength int64
	dataLength       int64
	indexLength      int64
}

type statementEventStats struct {
	schema                    string
	digest                    string
	digestText                string
	sumTimerWait              int64
	countErrors               int64
	countWarnings             int64
	countRowsAffected         int64
	countRowsSent             int64
	countRowsExamined         int64
	countCreatedTmpDiskTables int64
	countCreatedTmpTables     int64
	countSortMergePasses      int64
	countSortRows             int64
	countNoIndexUsed          int64
}

type tableLockWaitEventStats struct {
	schema                        string
	name                          string
	countReadNormal               int64
	countReadWithSharedLocks      int64
	countReadHighPriority         int64
	countReadNoInsert             int64
	countReadExternal             int64
	countWriteAllowWrite          int64
	countWriteConcurrentInsert    int64
	countWriteLowPriority         int64
	countWriteNormal              int64
	countWriteExternal            int64
	sumTimerReadNormal            int64
	sumTimerReadWithSharedLocks   int64
	sumTimerReadHighPriority      int64
	sumTimerReadNoInsert          int64
	sumTimerReadExternal          int64
	sumTimerWriteAllowWrite       int64
	sumTimerWriteConcurrentInsert int64
	sumTimerWriteLowPriority      int64
	sumTimerWriteNormal           int64
	sumTimerWriteExternal         int64
}

type replicaStatusStats struct {
	replicaIOState              string
	sourceHost                  string
	sourceUser                  string
	sourcePort                  int64
	connectRetry                int64
	sourceLogFile               string
	readSourceLogPos            int64
	relayLogFile                string
	relayLogPos                 int64
	relaySourceLogFile          string
	replicaIORunning            string
	replicaSQLRunning           string
	replicateDoDB               string
	replicateIgnoreDB           string
	replicateDoTable            string
	replicateIgnoreTable        string
	replicateWildDoTable        string
	replicateWildIgnoreTable    string
	lastErrno                   int64
	lastError                   string
	skipCounter                 int64
	execSourceLogPos            int64
	relayLogSpace               int64
	untilCondition              string
	untilLogFile                string
	untilLogPos                 string
	sourceSSLAllowed            string
	sourceSSLCAFile             string
	sourceSSLCAPath             string
	sourceSSLCert               string
	sourceSSLCipher             string
	sourceSSLKey                string
	secondsBehindSource         sql.NullInt64
	sourceSSLVerifyServerCert   string
	lastIOErrno                 int64
	lastIOError                 string
	lastSQLErrno                int64
	lastSQLError                string
	replicateIgnoreServerIDs    string
	sourceServerID              int64
	sourceUUID                  string
	sourceInfoFile              string
	sqlDelay                    int64
	sqlRemainingDelay           sql.NullInt64
	replicaSQLRunningState      string
	sourceRetryCount            int64
	sourceBind                  string
	lastIOErrorTimestamp        string
	lastSQLErrorTimestamp       string
	sourceSSLCrl                string
	sourceSSLCrlpath            string
	retrievedGtidSet            string
	executedGtidSet             string
	autoPosition                string
	replicateRewriteDB          string
	channelName                 string
	sourceTLSVersion            string
	sourcePublicKeyPath         string
	getSourcePublicKey          int64
	networkNamespace            string
	usingGtid                   string
	gtidIoPos                   string
	slaveDdlGroups              int64
	slaveNonTransactionalGroups int64
	slaveTransactionalGroups    int64
	retriedTransactions         int64
	maxRelayLogSize             int64
	executedLogEntries          int64
	slaveReceivedHeartbeats     int64
	slaveHeartbeatPeriod        int64
	gtidSlavePos                string
	masterLastEventTime         string
	slaveLastEventTime          string
	masterSlaveTimeDiff         string
	parallelMode                string
	replicateDoDomainIDs        string
	replicateIgnoreDomainIDs    string
}

type querySample struct {
	sessionID          int64
	threadID           int64
	processlistUser    string
	processlistHost    string
	clientPort         uint64
	processlistDB      string
	processlistCommand string
	processlistState   string
	sqlText            string
	digest             string
	eventID            int64
	sessionStatus      string
	waitEvent          string
	waitTime           float64
	statementTimerWait float64
	traceparent        string
}

type topQuery struct {
	schemaName                string
	digest                    string
	digestText                string
	countStar                 int64
	sumTimerWaitInPicoSeconds int64
	querySampleText           string
}

var _ client = (*mySQLClient)(nil)

func newMySQLClient(conf *Config) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

func (c *mySQLClient) Connect() error { _ = "STUB: not implemented"; return nil }

// Version detection runs exactly once during Connect and is non-fatal.
// If the query fails (e.g. no database reachable at startup) or the
// version string cannot be parsed, dbVersion stays at its zero value,
// which selects the safe fallback template (MySQL <8 / MariaDB behavior)
// for the entire lifetime of this receiver instance. Any connection error
// encountered here will cause the receiver to operate with incorrect
// version information; it will not be retried.
//
// This is intentional: sql.Open is lazy and the component lifecycle test
// calls start() against 127.0.0.1:3306 with no live database. A hard
// failure here would break that test with no way to inject a mock client
// before Connect runs. Real connection errors surface on the first scrape.

// fetchDBVersion queries the database for its version string and parses it
// into a dbVersion. Called once during Connect. A short context timeout
// prevents a blackholed or slow endpoint from stalling collector startup.
func (c *mySQLClient) fetchDBVersion() (dbVersion, error) {
	_ = "STUB: not implemented"
	return *new(dbVersion), nil
}

// mariaDBVersionRe extracts the leading semver triplet from a MariaDB VERSION()
// string after the optional MySQL-compat "5.5.5-" prefix has been stripped.
// Examples handled:
//   - "10.11.6-MariaDB"                          → "10.11.6"
//   - "5.5.5-10.11.6-MariaDB"                    → "10.11.6"
//   - "10.6.14-MariaDB-1:10.6.14+maria~ubu2204"  → "10.6.14"
//   - "10.11.6-MariaDB-log"                      → "10.11.6"
var mariaDBVersionRe = regexp.MustCompile(`^(\d+\.\d+\.\d+)`)

// mysqlCompatPrefix is prepended by older MariaDB 10.x builds to fool
// MySQL clients that require a version >= 5.5.5. Strip it before parsing.
const mysqlCompatPrefix = "5.5.5-"

// parseDBVersion parses a raw VERSION() string into a dbVersion.
//
// MySQL strings (no "MariaDB" substring): the semver is everything before the
// first "-" (handles suffixes like "-log").
//
// MariaDB strings: strip the optional MySQL-compat "5.5.5-" prefix, then
// extract the leading dotted-decimal triplet via regex.
func parseDBVersion(versionStr string) (dbVersion, error) {
	_ = "STUB: not implemented"
	return *new(dbVersion), nil
}

// MySQL: strip any suffix after the first "-"

// getDBVersion returns the cached database version established during Connect.
func (c *mySQLClient) getDBVersion() dbVersion {
	_ = "STUB: not implemented"

	// getGlobalStats queries the db for global status metrics.
	return *new(dbVersion)
}

func (c *mySQLClient) getGlobalStats() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getInnodbStats queries the db for innodb metrics.
func (c *mySQLClient) getInnodbStats() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getTableStats queries the db for information_schema table size metrics.
func (c *mySQLClient) getTableStats() ([]tableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getTableIoWaitsStats queries the db for table_io_waits metrics.
func (c *mySQLClient) getTableIoWaitsStats() ([]tableIoWaitsStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getIndexIoWaitsStats queries the db for index_io_waits metrics.
func (c *mySQLClient) getIndexIoWaitsStats() ([]indexIoWaitsStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *mySQLClient) getStatementEventsStats() ([]statementEventStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *mySQLClient) getTableLockWaitEventStats() ([]tableLockWaitEventStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *mySQLClient) getReplicaStatusStats(supportsReplicaStatus bool) ([]replicaStatusStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//go:embed templates/topQuery.tmpl
var topQueryTemplate string

//go:embed templates/topQueryNoSampleText.tmpl
var topQueryNoSampleTextTemplate string

func (c *mySQLClient) getTopQueries(topNValue, lookbackTime uint64, supportsSampleText bool) ([]topQuery, error) {
	_ = "STUB: not implemented"
	// Select the appropriate template based on version support.
	// MySQL <8 and all MariaDB versions lack query_sample_text in
	// events_statements_summary_by_digest, so we use the 5-column fallback.
	return nil, nil
}

// scanRow is defined once outside the loop to avoid re-evaluating
// supportsSampleText on every iteration.

// querySampleText stays "" — sentinel for "no sample available"

//go:embed templates/querySample.tmpl
var querySampleTemplate string

func (c *mySQLClient) getQuerySamples(limit uint64, supportsProcesslist bool) ([]querySample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pl.HOST from performance_schema.processlist uses "host:port" for IPv4 and
// "[::1]:port" for IPv6, both parseable by net.SplitHostPort. When processlist
// is not joined (supportsProcesslist=false), processlistHost is a bare hostname
// from thread.processlist_host and SplitHostPort will return an error — in that
// case we leave processlistHost as-is and clientPort as 0.

func (c *mySQLClient) explainQuery(digestText, sampleStatement, schema, digest string, logger *zap.Logger) string {
	_ = "STUB: not implemented"
	return ""
}

// This function filters out queries that are unsupported by 'EXPLAIN'
// ref: https://dev.mysql.com/doc/refman/8.4/en/using-explain.html
func isQueryExplainable(query string) bool { _ = "STUB: not implemented"; return false }

// stripLeadingSQLComments removes leading block (/* ... */)
// and line comments (-- ... and # ...) from a query.
func stripLeadingSQLComments(query string) string { _ = "STUB: not implemented"; return "" }

// remaining text is a comment

func query(c mySQLClient, query string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *mySQLClient) Close() error { _ = "STUB: not implemented"; return nil }
