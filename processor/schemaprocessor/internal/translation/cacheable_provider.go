// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/translation"

import (
	"context"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/metadata"
)

// CacheableProvider is a provider that caches the result of another provider.
// If the provider returns an error, the CacheableProvider will retry the call till limit
// If the provider returns an error multiple times in a row, the CacheableProvider will rate limit the calls.
type CacheableProvider struct {
	provider Provider
	cache    *cache.Cache
	mu       sync.Mutex
	// cooldown is the time to wait before retrying a failed call.
	cooldown time.Duration
	// callcount tracks the number of failed calls in a row.
	callcount int
	// limit is the number of failed calls to allow before setting the cooldown period.
	limit int
	// lastErr is the last error returned by the provider
	lastErr error
	// resetTime is the time when the rate limit will be reset
	resetTime time.Time
	// telemetryBuilder is used to record cache hit/miss metrics
	telemetryBuilder *metadata.TelemetryBuilder
}

// NewCacheableProvider creates a new CacheableProvider.
// The cooldown parameter is the time to wait before retrying a failed call.
// The limit parameter is the number of failed calls to allow before setting the cooldown period.
func NewCacheableProvider(provider Provider, cooldown time.Duration, limit int, telemetryBuilder *metadata.TelemetryBuilder) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (p *CacheableProvider) Retrieve(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	// Check if the key is in the cache.
	return "", nil
}

// Check if the key is in the cache again in case it was added while waiting for the lock.

// Check if the function is currently rate-limited

// After the cooldown expires, allow one retry but keep the count so the
// next failure triggers a new cooldown immediately.

// Release the lock before the HTTP call so other goroutines are not blocked
// for the duration of the network request.

// If the call limit is reached, set the cooldown period
