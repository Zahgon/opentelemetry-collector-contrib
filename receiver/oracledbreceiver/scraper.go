// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oracledbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver/internal/metadata"
)

const (
	statsSQL                       = "select * from v$sysstat"
	enqueueDeadlocks               = "enqueue deadlocks"
	exchangeDeadlocks              = "exchange deadlocks"
	executeCount                   = "execute count"
	parseCountTotal                = "parse count (total)"
	parseCountHard                 = "parse count (hard)"
	logons                         = "logons cumulative"
	userCommits                    = "user commits"
	userRollbacks                  = "user rollbacks"
	physicalReads                  = "physical reads"
	physicalReadsDirect            = "physical reads direct"
	physicalReadIORequests         = "physical read IO requests"
	physicalWrites                 = "physical writes"
	physicalWritesDirect           = "physical writes direct"
	physicalWriteIORequests        = "physical write IO requests"
	queriesParallelized            = "queries parallelized"
	ddlStatementsParallelized      = "DDL statements parallelized"
	dmlStatementsParallelized      = "DML statements parallelized"
	parallelOpsNotDowngraded       = "Parallel operations not downgraded"
	parallelOpsDowngradedToSerial  = "Parallel operations downgraded to serial"
	parallelOpsDowngraded1To25Pct  = "Parallel operations downgraded 1 to 25 pct"
	parallelOpsDowngraded25To50Pct = "Parallel operations downgraded 25 to 50 pct"
	parallelOpsDowngraded50To75Pct = "Parallel operations downgraded 50 to 75 pct"
	parallelOpsDowngraded75To99Pct = "Parallel operations downgraded 75 to 99 pct"
	sessionLogicalReads            = "session logical reads"
	cpuTime                        = "CPU used by this session"
	pgaMemory                      = "session pga memory"
	dbBlockGets                    = "db block gets"
	consistentGets                 = "consistent gets"
	sessionCountSQL                = "select status, type, count(*) as VALUE FROM v$session GROUP BY status, type"
	systemResourceLimitsSQL        = "select RESOURCE_NAME, CURRENT_UTILIZATION, LIMIT_VALUE, CASE WHEN TRIM(INITIAL_ALLOCATION) LIKE 'UNLIMITED' THEN '-1' ELSE TRIM(INITIAL_ALLOCATION) END as INITIAL_ALLOCATION, CASE WHEN TRIM(LIMIT_VALUE) LIKE 'UNLIMITED' THEN '-1' ELSE TRIM(LIMIT_VALUE) END as LIMIT_VALUE from v$resource_limit"
	tablespaceUsageSQL             = `
		select um.TABLESPACE_NAME, um.USED_SPACE, um.TABLESPACE_SIZE, ts.BLOCK_SIZE
		FROM DBA_TABLESPACE_USAGE_METRICS um INNER JOIN DBA_TABLESPACES ts
		ON um.TABLESPACE_NAME = ts.TABLESPACE_NAME`
	dataDictHitRatioSQL = "SELECT (1-(SUM(getmisses)/SUM(gets))) * 100 as DATA_DICTIONARY_HIT_RATIO FROM v$rowcache WHERE getmisses + gets <> 0"
	recycleBinSizeSQL   = "SELECT nvl(SUM(SPACE*(SELECT value FROM v$parameter WHERE name = 'db_block_size')),0) as RECYCLE_BIN_SIZE_BYTES FROM dba_recyclebin"
	storageUsageSQL     = "WITH total_bytes AS (SELECT SUM(bytes) AS total FROM dba_data_files) SELECT (total - (SELECT SUM(bytes) FROM dba_free_space)) AS USED_DB_SIZE, total AS ALLOCATED_DB_SIZE FROM total_bytes"

	colDataDictHitRatio    = "DATA_DICTIONARY_HIT_RATIO"
	colRecycleBinSizeBytes = "RECYCLE_BIN_SIZE_BYTES"
	colUsedDBSize          = "USED_DB_SIZE"
	colAllocatedDBSize     = "ALLOCATED_DB_SIZE"

	sqlIDAttr        = "SQL_ID"
	childAddressAttr = "CHILD_ADDRESS"
	childNumberAttr  = "CHILD_NUMBER"
	sqlTextAttr      = "SQL_FULLTEXT"
	dbSystemNameVal  = "oracle"

	queryExecutionMetric        = "EXECUTIONS"
	elapsedTimeMetric           = "ELAPSED_TIME"
	cpuTimeMetric               = "CPU_TIME"
	applicationWaitTimeMetric   = "APPLICATION_WAIT_TIME"
	concurrencyWaitTimeMetric   = "CONCURRENCY_WAIT_TIME"
	userIoWaitTimeMetric        = "USER_IO_WAIT_TIME"
	clusterWaitTimeMetric       = "CLUSTER_WAIT_TIME"
	rowsProcessedMetric         = "ROWS_PROCESSED"
	bufferGetsMetric            = "BUFFER_GETS"
	physicalReadRequestsMetric  = "PHYSICAL_READ_REQUESTS"
	physicalWriteRequestsMetric = "PHYSICAL_WRITE_REQUESTS"
	physicalReadBytesMetric     = "PHYSICAL_READ_BYTES"
	physicalWriteBytesMetric    = "PHYSICAL_WRITE_BYTES"
	queryDiskReadsMetric        = "DISK_READS"
	queryDirectReadsMetric      = "DIRECT_READS"
	queryDirectWritesMetric     = "DIRECT_WRITES"
	procedureExecutionsMetric   = "PROCEDURE_EXECUTIONS"

	// Stored procedure columns
	objectIDAttr    = "PROGRAM_ID"
	objectNameAttr  = "PROCEDURE_NAME"
	objectTypeAttr  = "PROCEDURE_TYPE"
	commandTypeAttr = "COMMAND_TYPE"
)

var (
	//go:embed templates/oracleQuerySampleSql.tmpl
	samplesQuery string
	//go:embed templates/oracleQueryMetricsAndTextSql.tmpl
	oracleQueryMetricsSQL string
	//go:embed templates/oracleQueryPlanSql.tmpl
	oracleQueryPlanDataSQL string
)

type dbProviderFunc func() (*sql.DB, error)

type clientProviderFunc func(*sql.DB, string, *zap.Logger) dbClient

type oracleScraper struct {
	statsClient                dbClient
	tablespaceUsageClient      dbClient
	systemResourceLimitsClient dbClient
	sessionCountClient         dbClient
	oracleQueryMetricsClient   dbClient
	oraclePlanDataClient       dbClient
	samplesQueryClient         dbClient
	dataDictHitRatioClient     dbClient
	recycleBinSizeClient       dbClient
	storageUsageClient         dbClient
	db                         *sql.DB
	clientProviderFunc         clientProviderFunc
	mb                         *metadata.MetricsBuilder
	lb                         *metadata.LogsBuilder
	dbProviderFunc             dbProviderFunc
	logger                     *zap.Logger
	id                         component.ID
	instanceName               string
	hostName                   string
	scrapeCfg                  scraperhelper.ControllerConfig
	startTime                  pcommon.Timestamp
	metricsBuilderConfig       metadata.MetricsBuilderConfig
	logsBuilderConfig          metadata.LogsBuilderConfig
	metricCache                *lru.Cache[string, map[string]int64]
	topQueryCollectCfg         TopQueryCollection
	obfuscator                 *obfuscator
	querySampleCfg             QuerySample
	serviceInstanceID          string
	lastExecutionTimestamp     time.Time
}

func newScraper(metricsBuilder *metadata.MetricsBuilder, metricsBuilderConfig metadata.MetricsBuilderConfig, scrapeCfg scraperhelper.ControllerConfig, logger *zap.Logger, providerFunc dbProviderFunc, clientProviderFunc clientProviderFunc, instanceName, hostName string) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}

func newLogsScraper(logsBuilder *metadata.LogsBuilder, logsBuilderConfig metadata.LogsBuilderConfig, scrapeCfg scraperhelper.ControllerConfig,
	logger *zap.Logger, providerFunc dbProviderFunc, clientProviderFunc clientProviderFunc, instanceName string, metricCache *lru.Cache[string, map[string]int64],
	topQueryCollectCfg TopQueryCollection, querySampleCfg QuerySample, hostName string,
) (scraper.Logs, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Logs), nil
}

func (s *oracleScraper) start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *oracleScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// divide by 100 as the value is expressed in tens of milliseconds

// Tablespace size should never be empty using the DBA_TABLESPACE_USAGE_METRICS query. This logic is done
// to preserve backward compatibility for with the original metric gathered from querying DBA_TABLESPACES

func (s *oracleScraper) collectDataDictHitRatio(ctx context.Context, scrapeErrors *[]error) {
	_ = "STUB: not implemented"
	return
}

func (s *oracleScraper) collectRecycleBinSize(ctx context.Context, scrapeErrors *[]error) {
	_ = "STUB: not implemented"
	return
}

func (s *oracleScraper) collectStorageUsage(ctx context.Context, scrapeErrors *[]error) {
	_ = "STUB: not implemented"
	return
}

type queryMetricCacheHit struct {
	sqlID        string
	childNumber  string
	childAddress string
	queryText    string
	metrics      map[string]int64
	objectID     int64
	objectName   string
	objectType   string
	commandType  int64
}

func (s *oracleScraper) scrapeLogs(ctx context.Context) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (s *oracleScraper) collectTopNMetricData(ctx context.Context, logs plog.Logs, collectionTime time.Time, lookbackTimeSeconds int) error {
	_ = "STUB: not implemented"

	// get metrics and query texts from DB
	return nil
}

// if we have a cache hit and the query doesn't belong to top N, cache is updated anyway
// as a result, once it finally makes its way to the top N queries, only the latest delta will be sent downstream

// Parse stored procedure PROGRAM_ID

// if any of the deltas is less than zero, cursor was likely purged from the shared pool

// skip if possible purge or no new executions since last scrape

// if cache updates is not equal to rows returned, that indicates there is problem somewhere

// order by elapsed time delta, descending

// keep at most maxHitSize

func (s *oracleScraper) collectQuerySamples(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse stored procedure PROCEDURE_ID

// Parse wait time in seconds

func asFloatInSeconds(value int64) float64 { _ = "STUB: not implemented"; return 0 }

func (s *oracleScraper) obfuscateCacheHits(hits []queryMetricCacheHit) []queryMetricCacheHit {
	_ = "STUB: not implemented"
	return nil
}

// obfuscate and normalize the query text

func (s *oracleScraper) getChildAddressToPlanMap(ctx context.Context, hits []queryMetricCacheHit) map[string][]metricRow {
	_ = "STUB: not implemented"
	return nil
}

// child address was for internal use only, it's not going to be used beyond this point

func (*oracleScraper) getTopNMetricNames() []string { _ = "STUB: not implemented"; return nil }

func (s *oracleScraper) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *oracleScraper) setupResourceBuilder(rb *metadata.ResourceBuilder) *metadata.ResourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func getInstanceID(instanceString string, logger *zap.Logger) string {
	_ = "STUB: not implemented"
	return ""
}

// Replace the host value with machine name if connecting to localhost target

func constructInstanceID(host, port, service string) string { _ = "STUB: not implemented"; return "" }

func (s *oracleScraper) calculateLookbackSeconds() int { _ = "STUB: not implemented"; return 0 }

// vsqlRefreshLag is the buffer to account for v$sql maximum refresh latency (5 seconds) + 5 seconds to offset any collection delays.
// PS: https://docs.oracle.com/en/database/oracle/oracle-database/21/refrn/V-SQL.html
