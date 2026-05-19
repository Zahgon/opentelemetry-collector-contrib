// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package postgresqlreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"

import (
	"context"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver/internal/metadata"
)

const (
	readmeURL                 = "https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/v0.88.0/receiver/postgresqlreceiver/README.md"
	defaultPostgreSQLDatabase = "postgres"
)

type postgreSQLScraper struct {
	logger        *zap.Logger
	config        *Config
	clientFactory postgreSQLClientFactory
	mb            *metadata.MetricsBuilder
	lb            *metadata.LogsBuilder
	excludes      map[string]struct{}
	cache         *lru.Cache[string, float64]
	// if enabled, uses a separated attribute for the schema
	separateSchemaAttr     bool
	queryPlanCache         *expirable.LRU[string, string]
	newestQueryTimestamp   float64
	serviceInstanceID      string
	lastExecutionTimestamp time.Time
}

type errsMux struct {
	sync.RWMutex
	errs scrapererror.ScrapeErrors
}

func (e *errsMux) add(err error) { _ = "STUB: not implemented"; return }

func (e *errsMux) addPartial(err error) { _ = "STUB: not implemented"; return }

func (e *errsMux) combine() error { _ = "STUB: not implemented"; return nil }

func newPostgreSQLScraper(
	settings receiver.Settings,
	config *Config,
	clientFactory postgreSQLClientFactory,
	cache *lru.Cache[string, float64],
	queryPlanCache *expirable.LRU[string, string],
) *postgreSQLScraper {
	_ = "STUB: not implemented"
	return nil
}

type dbRetrieval struct {
	sync.RWMutex
	activityMap map[databaseName]int64
	dbSizeMap   map[databaseName]int64
	dbStats     map[databaseName]databaseStats
}

// scrape scrapes the metric stats, transforms them and attributes them into a metric slices.
func (p *postgreSQLScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (p *postgreSQLScraper) scrapeQuerySamples(ctx context.Context, maxRowsPerQuery int64) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (p *postgreSQLScraper) scrapeTopQuery(ctx context.Context, maxRowsPerQuery, topNQuery, maxExplainEachInterval int64, collectionInterval time.Duration) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (p *postgreSQLScraper) isCollectionDue(collectionTime time.Time, interval time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// This is the first collection

func (p *postgreSQLScraper) collectQuerySamples(ctx context.Context, dbClient client, limit int64, mux *errsMux, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Use a background context so query-sample logs are not automatically linked to the scrape context.

func (p *postgreSQLScraper) collectTopQuery(ctx context.Context, clientFactory postgreSQLClientFactory, limit, topNQuery, maxExplainEachInterval int64, mux *errsMux, logger *zap.Logger, collectionTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// this should not happen, but in case

// Use raw query (with $1, $2 placeholders) for EXPLAIN, not the obfuscated one (with ?)

// to avoid flood the error message. there are some internal queries meant to not be
// explained. we wait for the cache to expire and report the error again.

func (p *postgreSQLScraper) shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *postgreSQLScraper) retrieveDBMetrics(
	ctx context.Context,
	listClient client,
	databases []string,
	r *dbRetrieval,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) recordDatabase(now pcommon.Timestamp, db string, r *dbRetrieval, numTables int64) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) collectTables(ctx context.Context, now pcommon.Timestamp, dbClient client, db string, errs *errsMux) (numTables int64) {
	_ = "STUB: not implemented"
	return 0
}

func (p *postgreSQLScraper) collectIndexes(
	ctx context.Context,
	now pcommon.Timestamp,
	client client,
	database string,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) collectFunctions(
	ctx context.Context,
	now pcommon.Timestamp,
	client client,
	database string,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) collectBGWriterStats(
	ctx context.Context,
	now pcommon.Timestamp,
	client client,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) collectDatabaseLocks(
	ctx context.Context,
	now pcommon.Timestamp,
	client client,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) collectMaxConnections(
	ctx context.Context,
	now pcommon.Timestamp,
	client client,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) collectReplicationStats(
	ctx context.Context,
	now pcommon.Timestamp,
	client client,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) collectWalAge(
	ctx context.Context,
	now pcommon.Timestamp,
	client client,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

// return no error as there is no last archive to derive the value from

func (p *postgreSQLScraper) retrieveDatabaseStats(
	ctx context.Context,
	wg *sync.WaitGroup,
	client client,
	databases []string,
	r *dbRetrieval,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) retrieveDatabaseSize(
	ctx context.Context,
	wg *sync.WaitGroup,
	client client,
	databases []string,
	r *dbRetrieval,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (*postgreSQLScraper) retrieveBackends(
	ctx context.Context,
	wg *sync.WaitGroup,
	client client,
	databases []string,
	r *dbRetrieval,
	errs *errsMux,
) {
	_ = "STUB: not implemented"
	return
}

func (p *postgreSQLScraper) setupResourceBuilder(rb *metadata.ResourceBuilder, database, schema, table, index string) *metadata.ResourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func getInstanceID(instanceString string, logger *zap.Logger) string {
	_ = "STUB: not implemented"
	return ""
}
