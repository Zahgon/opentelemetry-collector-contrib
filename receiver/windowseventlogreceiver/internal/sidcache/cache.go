// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package sidcache provides a high-performance LRU cache for Windows SID-to-name resolution.
// It uses the Windows Local Security Authority (LSA) API to resolve Security Identifiers (SIDs)
// to human-readable user and group names, with support for well-known SIDs and TTL-based expiration.
package sidcache // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver/internal/sidcache"

import (
	"regexp"
	"time"
)

// Default configuration values
const (
	DefaultCacheSize = 10000
	DefaultCacheTTL  = 15 * time.Minute
)

// sidPattern validates SID format: S-1-<revision>-<authority>-<sub-authorities>
var sidPattern = regexp.MustCompile(`^S-1-\d+(-\d+)+$`)

// isSIDFormat checks if a string matches the SID format
func isSIDFormat(sid string) bool { _ = "STUB: not implemented"; return false }

// IsSIDField checks if a field name likely contains a SID
// Used by the receiver to identify which fields need resolution
func IsSIDField(fieldName string) bool {
	_ = "STUB: not implemented"
	// Common SID field patterns in Windows events
	return false
}

// ends with "Sid"
// exact match for security.user_id
