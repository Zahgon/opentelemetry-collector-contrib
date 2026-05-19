// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"sync"
	"time"
)

// cacheEntry keeps information about host and expiration time
type cacheEntry struct {
	hostname   string
	expireTime time.Time
}

const (
	defaultInvalidationInterval time.Duration = 5 * time.Minute
)

type IPResolver struct {
	cache                map[string]cacheEntry
	mutex                sync.RWMutex
	done                 chan bool
	stopped              bool
	invalidationInterval time.Duration
}

// Create new resolver
func NewIPResolver() *IPResolver { _ = "STUB: not implemented"; return nil }

// Stop cache invalidation
func (r *IPResolver) Stop() { _ = "STUB: not implemented"; return }

// start runs cache invalidation every 5 minutes
func (r *IPResolver) start() { _ = "STUB: not implemented"; return }

// invalidateCache removes not longer valid entries from cache
func (r *IPResolver) invalidateCache() { _ = "STUB: not implemented"; return }

// GetHostFromIp returns hostname for given ip
// It is taken from cache if exists,
// otherwise lookup is performed and result is put into cache
func (r *IPResolver) GetHostFromIP(ip string) (host string) { _ = "STUB: not implemented"; return "" }

// lookupIPAddr resturns hostname based on ip address
func (*IPResolver) lookupIPAddr(ip string) (host string) { _ = "STUB: not implemented"; return "" }

// Trim one trailing '.'.
