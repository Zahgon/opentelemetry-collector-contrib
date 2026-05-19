// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package sqlserverreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver/internal/metadata"
)

// SQL Server Perf Counter (PC) Scraper. This is used to scrape metrics from Windows Perf Counters.
type sqlServerPCScraper struct {
	logger           *zap.Logger
	config           *Config
	watcherRecorders []watcherRecorder
	mb               *metadata.MetricsBuilder
}

// watcherRecorder is a struct containing perf counter watcher along with corresponding value recorder.
type watcherRecorder struct {
	watcher  winperfcounters.PerfCounterWatcher
	recorder recordFunc
}

// curriedRecorder is a recorder function that already has value to be recorded,
// it needs metadata.MetricsBuilder and timestamp as arguments.
type curriedRecorder func(*metadata.MetricsBuilder, pcommon.Timestamp)

// newSQLServerPCScraper returns a new sqlServerPCScraper.
func newSQLServerPCScraper(params receiver.Settings, cfg *Config) *sqlServerPCScraper {
	_ = "STUB: not implemented"
	return nil
}

// start creates and sets the watchers for the scraper.
func (s *sqlServerPCScraper) start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// The instance name must be preceded by "MSSQL$" to indicate that it is a named instance

// scrape collects windows performance counter data from all watchers and then records/emits it using the metricBuilder
func (s *sqlServerPCScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// recordersPerDatabase scrapes perf counter values using provided []watcherRecorder and returns
// a map of database name to curriedRecorder that includes the recorded value in its closure.
func recordersPerDatabase(watcherRecorders []watcherRecorder) (map[string][]curriedRecorder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// it's important to initialize new values for the closure.

func (s *sqlServerPCScraper) emitMetricGroup(recorders []curriedRecorder, databaseName string) {
	_ = "STUB: not implemented"
	return
}

// shutdown stops all of the watchers for the scraper.
func (s sqlServerPCScraper) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
