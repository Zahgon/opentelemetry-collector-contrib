// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlserverreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"

import (
	"errors"

	lru "github.com/hashicorp/golang-lru/v2"
	_ "github.com/microsoft/go-mssqldb"                     // register Db driver
	_ "github.com/microsoft/go-mssqldb/integratedauth/krb5" // register Db driver
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver/internal/metadata"
)

var errConfigNotSQLServer = errors.New("config was not a sqlserver receiver config")

// newCache creates a new cache with the given size.
// If the size is less or equal to 0, it will be set to 1.
// It will never return an error.
func newCache(size int) *lru.Cache[string, int64] { _ = "STUB: not implemented"; return nil }

// lru will only returns error when the size is less than 0

// NewFactory creates a factory for SQL Server receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func setupQueries(cfg *Config) []string { _ = "STUB: not implemented"; return nil }

func setupLogQueries(cfg *Config) []string { _ = "STUB: not implemented"; return nil }

// Assumes config has all information necessary to directly connect to the database
func getDBConnectionString(config *Config) string { _ = "STUB: not implemented"; return "" }

// SQL Server scraper creation is split out into a separate method for the sake of testing.
func setupSQLServerScrapers(params receiver.Settings, cfg *Config) []*sqlServerScraperHelper {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Test if this needs to be re-defined for each scraper
// This should be tested when there is more than one query being made.

// lru only returns error when the size is less than 0

// SQL Server scraper creation is split out into a separate method for the sake of testing.
func setupSQLServerLogsScrapers(params receiver.Settings, cfg *Config) []*sqlServerScraperHelper {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Test if this needs to be re-defined for each scraper
// This should be tested when there is more than one query being made.

// we have 8 metrics in this query and multiple 2 to allow to cache more queries.

// Note: This method will fail silently if there is no work to do. This is an acceptable use case
// as this receiver can still get information on Windows from performance counters without a direct
// connection. Messages will be logged at the INFO level in such cases.
func setupScrapers(params receiver.Settings, cfg *Config) ([]scraperhelper.ControllerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: This method will fail silently if there is no work to do. This is an acceptable use case
// as this receiver can still get information on Windows from performance counters without a direct
// connection. Messages will be logged at the INFO level in such cases.
func setupLogsScrapers(params receiver.Settings, cfg *Config) ([]scraperhelper.ControllerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isDatabaseIOQueryEnabled(metrics *metadata.MetricsConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func isPerfCounterQueryEnabled(metrics *metadata.MetricsConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func isWaitStatsQueryEnabled(metrics *metadata.MetricsConfig) bool {
	_ = "STUB: not implemented"
	return false
}
