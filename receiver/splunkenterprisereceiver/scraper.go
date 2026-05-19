// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkenterprisereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver"

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver/internal/metadata"
)

var errMaxSearchWaitTimeExceeded = errors.New("maximum search wait time exceeded for metric")

const receiverScope = "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver"

type splunkScraper struct {
	splunkClient  *splunkEntClient
	settings      component.TelemetrySettings
	conf          *Config
	mb            *metadata.MetricsBuilder
	customMetrics pmetric.Metrics
	customMu      sync.Mutex
}

func newSplunkMetricsScraper(params receiver.Settings, cfg *Config) splunkScraper {
	_ = "STUB: not implemented"
	return *new(splunkScraper)
}

// Create a client instance and add to the splunkScraper
func (s *splunkScraper) start(ctx context.Context, h component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// listens to the error channel and combines errors sent from different metric scrape functions,
// returning the combined error list should context timeout or a nil error value is sent in the
// channel signifying the end of a scrape cycle
func errorListener(ctx context.Context, eQueue <-chan error, eOut chan<- *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// The big one: Describes how all scraping tasks should be performed. Part of the scraper interface
func (s *splunkScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// if the build and version info has been configured that is pulled here

// actual function body

// Each metric has its own scrape function associated with it
func (s *splunkScraper) scrapeLicenseUsageByIndex(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeAvgExecLatencyByHost(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeIndexerAvgRate(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeIndexerPipelineQueues(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeBucketsSearchableStatus(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeIndexesBucketCountAdHoc(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeSchedulerCompletionRatioByHost(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeIndexerRawWriteSecondsByHost(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeIndexerCPUSecondsByHost(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeAvgIopsByHost(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

func (s *splunkScraper) scrapeSchedulerRunTimeByHost(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// Because we have to utilize network resources for each KPI we should check that each metrics
	// is enabled before proceeding
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// Record the results

// Helper function for unmarshaling search endpoint requests
func unmarshallSearchReq(res *http.Response, sr *searchResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// Scrape index throughput introspection endpoint
func (s *splunkScraper) scrapeIndexThroughput(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape indexes extended total size
func (s *splunkScraper) scrapeIndexesTotalSize(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape indexes extended total event count
func (s *splunkScraper) scrapeIndexesEventCount(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape indexes extended total bucket count
func (s *splunkScraper) scrapeIndexesBucketCount(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape indexes extended raw size
func (s *splunkScraper) scrapeIndexesRawSize(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape indexes extended bucket event count
func (s *splunkScraper) scrapeIndexesBucketEventCount(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape indexes extended bucket hot/warm count
func (s *splunkScraper) scrapeIndexesBucketHotWarmCount(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape introspection queues
func (s *splunkScraper) scrapeIntrospectionQueues(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape introspection queues bytes
func (s *splunkScraper) scrapeIntrospectionQueuesBytes(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape introspection kv store status
func (s *splunkScraper) scrapeKVStoreStatus(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// overall status

// a 0 gauge value means that the metric was not reported in the api call
// to the introspection endpoint.

// set to 0 to indicate no status being reported

// Scrape dispatch artifacts
func (s *splunkScraper) scrapeSearchArtifacts(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	// if NONE of the metrics set in this scrape are set we return early
	return
}

// Scrape Health Introspection Endpoint
func (s *splunkScraper) scrapeHealth(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

func (s *splunkScraper) traverseHealthDetailFeatures(details healthDetails, now pcommon.Timestamp, i infoContent) {
	_ = "STUB: not implemented"
	return
}

// somewhat unique scrape function for gathering the info attribute
func (s *splunkScraper) scrapeInfo(_ context.Context, _ pcommon.Timestamp, errs chan error) map[any]info {
	_ = "STUB: not implemented"
	// there could be an endpoint configured for each type (never more than 3)
	return nil
}

// Scrape Search Metrics
func (s *splunkScraper) scrapeSearch(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// if its a 204 the body will be empty because we are still waiting on search results

// set search TTL once we know the sid

// log the error but it doesn't need to fail

// if no errors and 200 returned scrape was successful, return. Note we must make sure that
// the 200 is coming after the first request which provides a jobId to retrieve results

// searchStates list of all Splunk Search Status.

// get the 1 search we submitted

// record for all possible search states

// wait for 2s until trying again

// At this point either the search is done or the timeout has been reached.
// We don't also check for time due to this and the fact that the last call
// could have a successful DONE state at the last second

func (s *splunkScraper) recordSplunkSearchInitiationDataPoint(now pcommon.Timestamp, value int64, i infoContent) {
	_ = "STUB: not implemented"
	return
}

func (s *splunkScraper) recordSplunkSearchStatusDataPoint(now pcommon.Timestamp, value int64, state string, i infoContent) {
	_ = "STUB: not implemented"
	return
}

func (s *splunkScraper) recordSplunkSearchDurationDataPoint(now pcommon.Timestamp, value float64, i infoContent) {
	_ = "STUB: not implemented"
	return
}

func (s *splunkScraper) recordSplunkSearchSuccessDataPoint(now pcommon.Timestamp, value int64, i infoContent) {
	_ = "STUB: not implemented"
	return
}

// get the MetaEntries from a search to use for introspection
func (s *splunkScraper) getSearchEntries(sid string) (searchMetaEntries, error) {
	_ = "STUB: not implemented"
	return *new(searchMetaEntries), nil
}

// setSearchJobTTLById sets the SearchJob's TTL on the remote Splunk server to Timeout and returns a ControlResponse.
func (s *splunkScraper) setSearchJobTTLByID(sid string) error {
	_ = "STUB: not implemented"
	return nil
}

// Scrape Indexer Cluster Manger Status Endpoint
func (s *splunkScraper) scrapeIndexerClusterManagerStatus(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// Scrape License Endpoint
func (s *splunkScraper) scrapeLicenses(_ context.Context, now pcommon.Timestamp, info infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

// expiry time - current time in seconds converted to int64

func (s *splunkScraper) formatSPLForSearch(search SearchConfig) string {
	_ = "STUB: not implemented"
	return ""
}

// Strip "search=" prefix so we don't duplicate

// `| tstats` doesn't get time window normally so we need to check for it and treat it with specialness
func isTstatsCommand(spl string) bool { _ = "STUB: not implemented"; return false }

func (s *splunkScraper) injectTstatsTimeRange(spl string, search SearchConfig) string {
	_ = "STUB: not implemented"
	return ""
}

// Inject before "by" in bare tstats if present

func formatDurationForSplunk(d time.Duration) string { _ = "STUB: not implemented"; return "" }

func (s *splunkScraper) scrapeCustomSearches(_ context.Context, now pcommon.Timestamp, _ infoDict, errs chan error) {
	_ = "STUB: not implemented"
	return
}

func (s *splunkScraper) executeCustomSearch(search SearchConfig, eptType string) ([]*field, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// appendSearchMetrics writes metrics for a single search using customMu
func appendSearchMetrics(now pcommon.Timestamp, sm pmetric.ScopeMetrics, search SearchConfig, fields []*field) {
	_ = "STUB: not implemented"
	return
}

func parseFields(fields []*field) []map[string]string { _ = "STUB: not implemented"; return nil }
