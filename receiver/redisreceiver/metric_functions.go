// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// dataPointRecorders is called once at startup. Returns recorders for all metrics (except keyspace)
// we want to extract from Redis INFO and CLUSTER INFO.
func (rs *redisScraper) dataPointRecorders() map[string]any { _ = "STUB: not implemented"; return nil }

func (rs *redisScraper) recordUsedCPUSys(now pcommon.Timestamp, val float64) {
	_ = "STUB: not implemented"
	return
}

func (rs *redisScraper) recordUsedCPUSysChildren(now pcommon.Timestamp, val float64) {
	_ = "STUB: not implemented"
	return
}

func (rs *redisScraper) recordUsedCPUSysMainThread(now pcommon.Timestamp, val float64) {
	_ = "STUB: not implemented"
	return
}

func (rs *redisScraper) recordUsedCPUUser(now pcommon.Timestamp, val float64) {
	_ = "STUB: not implemented"
	return
}

func (rs *redisScraper) recordUsedCPUUserChildren(now pcommon.Timestamp, val float64) {
	_ = "STUB: not implemented"
	return
}

func (rs *redisScraper) recordUsedCPUUserMainThread(now pcommon.Timestamp, val float64) {
	_ = "STUB: not implemented"
	return
}
