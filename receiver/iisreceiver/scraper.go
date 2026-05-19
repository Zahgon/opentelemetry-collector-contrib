// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package iisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver"

import (
	"context"
	"regexp"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver/internal/metadata"
)

type iisReceiver struct {
	params                  component.TelemetrySettings
	config                  *Config
	consumer                consumer.Metrics
	totalWatcherRecorders   []watcherRecorder
	siteWatcherRecorders    []watcherRecorder
	appPoolWatcherRecorders []watcherRecorder
	queueMaxAgeWatchers     []instanceWatcher
	rb                      *metadata.ResourceBuilder
	mb                      *metadata.MetricsBuilder

	// for mocking
	newWatcher         func(string, string, string) (winperfcounters.PerfCounterWatcher, error)
	newWatcherFromPath func(string) (winperfcounters.PerfCounterWatcher, error)
	expandWildcardPath func(string) ([]string, error)
}

// watcherRecorder is a struct containing perf counter watcher along with corresponding value recorder.
type watcherRecorder struct {
	watcher  winperfcounters.PerfCounterWatcher
	recorder recordFunc
}

// instanceWatcher is a struct containing a perf counter watcher, along with the single instance the watcher records.
type instanceWatcher struct {
	watcher  winperfcounters.PerfCounterWatcher
	instance string
}

// newIisReceiver returns an iisReceiver
func newIisReceiver(settings receiver.Settings, cfg *Config, consumer consumer.Metrics) *iisReceiver {
	_ = "STUB: not implemented"
	return nil
}

// start builds the paths to the watchers
func (rcvr *iisReceiver) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// scrape pulls counter values from the watchers
func (rcvr *iisReceiver) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Maintain maps of site -> {val, recordFunc} and app -> {val, recordFunc}
// so that we can emit all metrics for a particular instance (site or app_pool) at once,
// keeping them in a single resource metric.

func (rcvr *iisReceiver) scrapeTotalMetrics(now pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// resource for total metrics is empty
// this makes it so that the order that the scrape functions are called doesn't matter

type valRecorder struct {
	val    float64
	record recordFunc
}

func (rcvr *iisReceiver) scrapeInstanceMetrics(wrs []watcherRecorder, instanceToRecorders map[string][]valRecorder) {
	_ = "STUB: not implemented"
	return
}

// This avoids recording the _Total instance.
// The _Total instance may be the only instance, because some instances require elevated permissions
// to list and scrape. In these cases, the per-instance metric is not available, and should not be recorded.

var negativeDenominatorError = "A counter with a negative denominator value was detected.\r\n"

func (rcvr *iisReceiver) scrapeMaxQueueAgeMetrics(appToRecorders map[string][]valRecorder) {
	_ = "STUB: not implemented"
	return
}

// This error occurs when there are no items in the queue;
// in this case, we would like to emit a 0 instead of logging an error (this is an expected scenario).

// No counters scraped

// emitInstanceMap records all metrics for each instance, then emits them all as a single resource metric
func (rcvr *iisReceiver) emitInstanceMap(now pcommon.Timestamp, instanceToRecorders map[string][]valRecorder, resourceSetter func(string)) {
	_ = "STUB: not implemented"
	return
}

// shutdown closes the watchers
func (rcvr iisReceiver) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (rcvr *iisReceiver) buildWatcherRecorders(confs []perfCounterRecorderConf, scrapeErrors *scrapererror.ScrapeErrors) []watcherRecorder {
	_ = "STUB: not implemented"
	return nil
}

var maxQueueItemAgeInstanceRegex = regexp.MustCompile(`\\HTTP Service Request Queues\((?P<instance>[^)]+)\)\\MaxQueueItemAge$`)

// buildMaxQueueItemAgeWatchers builds a watcher for each individual instance of the MaxQueueItemAge counter.
// This is done in order to capture the error when scraping each individual instance, because we want to ignore
// negative denominator errors.
func (rcvr *iisReceiver) buildMaxQueueItemAgeWatchers(scrapeErrors *scrapererror.ScrapeErrors) []instanceWatcher {
	_ = "STUB: not implemented"
	return nil
}

// skip total instance

func closeWatcherRecorders(wrs []watcherRecorder) error { _ = "STUB: not implemented"; return nil }

func closeInstanceWatchers(wrs []instanceWatcher) error { _ = "STUB: not implemented"; return nil }
