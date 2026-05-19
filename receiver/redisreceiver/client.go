// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"

import (
	"github.com/redis/go-redis/v9"
)

// Interface for a Redis client. Implementation can be faked for testing.
type client interface {
	// retrieves a string of key/value pairs of redis metadata
	retrieveInfo() (string, error)
	// retrieves a string of key/value pairs of redis cluster metadata
	retrieveClusterInfo() (string, error)
	// line delimiter
	// redis lines are delimited by \r\n, files (for testing) by \n
	delimiter() string
	// close release redis client connection pool
	close() error
}

// Wraps a real Redis client, implements `client` interface.
type redisClient struct {
	client *redis.Client
}

var _ client = (*redisClient)(nil)

// Creates a new real Redis client from the passed-in redis.Options.
func newRedisClient(options *redis.Options) client { _ = "STUB: not implemented"; return *new(client) }

// Redis strings are CRLF delimited.
func (*redisClient) delimiter() string {
	_ = "STUB: not implemented"

	// Retrieve Redis INFO. We retrieve all of the 'sections'.
	return ""
}

func (c *redisClient) retrieveInfo() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Retrieve Redis CLUSTER INFO.
func (c *redisClient) retrieveClusterInfo() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// close client to release connection pool.
func (c *redisClient) close() error { _ = "STUB: not implemented"; return nil }
