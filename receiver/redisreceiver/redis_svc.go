// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"

// Wraps a client, parses the Redis info command, returning a string-string map
// containing all of the key value pairs returned by INFO. Takes a line delimiter
// from the passed in client to support testing, because Redis uses CRLF and test
// data uses LF.
type redisSvc struct {
	client    client
	delimiter string
}

// Creates a new redisSvc. Pass in a client implementation.
func newRedisSvc(client client) *redisSvc { _ = "STUB: not implemented"; return nil }

// Calls the Redis INFO and CLUSTER INFO command on the client and returns an `info` map.
func (p *redisSvc) info() (info, error) { _ = "STUB: not implemented"; return *new(info), nil }

// defensive, should always == 2
