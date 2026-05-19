// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver/internal/metadata"
)

// Runs intermittently, fetching info from Redis, creating metrics/datapoints,
// and feeding them to a metricsConsumer.
type redisScraper struct {
	client     client
	redisSvc   *redisSvc
	settings   component.TelemetrySettings
	mb         *metadata.MetricsBuilder
	uptime     time.Duration
	configInfo configInfo
}

const redisMaxDbs = 16 // Maximum possible number of redis databases

func newRedisScraper(cfg *Config, settings receiver.Settings) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}

// Avoid background goroutines that trigger goleak in tests and on shutdown races.

func newRedisScraperWithClient(client client, settings receiver.Settings, cfg *Config) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}

func (rs *redisScraper) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// Scrape is called periodically, querying Redis and building Metrics to send to
// the next consumer. First builds 'fixed' metrics (non-keyspace metrics)
// defined at startup time. Then builds 'keyspace' metrics if there are any
// keyspace lines returned by Redis. There should be one keyspace line per
// active Redis database, of which there can be 16.
func (rs *redisScraper) Scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// recordCommonMetrics records metrics from Redis info key-value pairs.
func (rs *redisScraper) recordCommonMetrics(ts pcommon.Timestamp, inf info, recorders map[string]any) {
	_ = "STUB: not implemented"
	return
}

// Skip unregistered metric.

// recordKeyspaceMetrics records metrics from 'keyspace' Redis info key-value pairs,
// e.g. "db0: keys=1,expires=2,avg_ttl=3".
func (rs *redisScraper) recordKeyspaceMetrics(ts pcommon.Timestamp, inf info) {
	_ = "STUB: not implemented"
	return
}

// getRedisVersion retrieves version string from 'redis_version' Redis info key-value pairs
// e.g. "redis_version:5.0.7"
func (*redisScraper) getRedisVersion(inf info) string { _ = "STUB: not implemented"; return "" }

// getRedisMode retrieves mode string from 'redis_mode' Redis info key-value pairs
// e.g. "redis_mode:standalone"
func (*redisScraper) getRedisMode(inf info) string { _ = "STUB: not implemented"; return "" }

// recordModeMetrics records metrics from 'redis_mode' Redis info key-value pairs
// e.g. "redis_mode:standalone"
func (rs *redisScraper) recordModeMetrics(ts pcommon.Timestamp, mode string) {
	_ = "STUB: not implemented"
	return
}

// recordRoleMetrics records metrics from 'role' Redis info key-value pairs
// e.g. "role:master"
func (rs *redisScraper) recordRoleMetrics(ts pcommon.Timestamp, inf info) {
	_ = "STUB: not implemented"
	return
}

// recordCmdMetrics records per-command metrics from Redis info.
// These include command stats and command latency percentiles.
// Examples:
//
//	"cmdstat_mget:calls=1685,usec=6032,usec_per_call=3.58,rejected_calls=0,failed_calls=0"
//	"latency_percentiles_usec_lastsave:p50=1.003,p99=1.003,p99.9=1.003"
func (rs *redisScraper) recordCmdMetrics(ts pcommon.Timestamp, inf info) {
	_ = "STUB: not implemented"
	return
}

// recordCmdStatsMetrics records metrics for a particular Redis command.
// Only 'calls' and 'usec' are recorded at the moment.
// 'cmd' is the Redis command, 'val' is the values string (e.g. "calls=1685,usec=6032,usec_per_call=3.58,rejected_calls=0,failed_calls=0").
func (rs *redisScraper) recordCmdStatsMetrics(ts pcommon.Timestamp, cmd, val string) {
	_ = "STUB: not implemented"
	return
}

// skip bad items

// recordCmdLatencyMetrics record latency metrics of a particular Redis command.
// 'cmd' is the Redis command, 'val' is the values string (e.g. "p50=1.003,p99=1.003,p99.9=1.003).
// Latency values in the values string are expressed in microseconds.
func (rs *redisScraper) recordCmdLatencyMetrics(ts pcommon.Timestamp, cmd, val string) {
	_ = "STUB: not implemented"
	return
}

// metric is in seconds

// sentinelDataPointRecorders returns the map of supported Sentinel metrics.
func (rs *redisScraper) sentinelDataPointRecorders() map[string]any {
	_ = "STUB: not implemented"
	return nil
}
