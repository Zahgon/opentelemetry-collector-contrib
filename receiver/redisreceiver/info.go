// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redisreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"

import (
	"time"
)

// A map of the INFO data returned from Redis.
type info map[string]string

func (i info) getUptimeInSeconds() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
