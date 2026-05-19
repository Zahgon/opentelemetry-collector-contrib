// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"

// Holds percentile latencies, e.g. "p99" -> 1.5.
type latencies map[string]float64

// parseLatencyStats parses the values part of one entry in Redis latencystats section,
// e.g. "p50=181.247,p99=309.247,p99.9=1023.999".
func parseLatencyStats(str string) (latencies, error) {
	_ = "STUB: not implemented"
	return *new(latencies), nil
}
