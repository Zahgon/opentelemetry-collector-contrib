// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"

// Holds fields returned by the Keyspace section of the INFO command: e.g.
// "db0:keys=1,expires=2,avg_ttl=3"
type keyspace struct {
	db      string
	keys    int
	expires int
	avgTTL  int
}

// Turns a keyspace value (the part after the colon
// e.g. "keys=1,expires=2,avg_ttl=3") into a keyspace struct
func parseKeyspaceString(db int, str string) (*keyspace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
