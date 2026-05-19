// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package dns provides a DNS-based lookup source.
package dns // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/internal/source/dns"

import (
	"context"
	"net"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/lookupsource"
)

const sourceType = "dns"

type RecordType string

const (
	// Only PTR for now.
	RecordTypePTR RecordType = "PTR"
)

type Config struct {
	// RecordType specifies the DNS record type to look up.
	RecordType RecordType `mapstructure:"record_type"`

	// Timeout is the maximum time to wait for a DNS query.
	// Default: 1 second
	Timeout time.Duration `mapstructure:"timeout"`

	// Server is the DNS server to use (e.g., "8.8.8.8:53").
	// If empty, uses the system resolver.
	Server string `mapstructure:"server"`

	// Cache configures caching for DNS lookups.
	// Enabled by default.
	// Disabling is not recommended due to potential performance impact.
	Cache lookupsource.CacheConfig `mapstructure:"cache"`
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Valid

func NewFactory() lookupsource.SourceFactory {
	_ = "STUB: not implemented"
	return *new(lookupsource.SourceFactory)
}

func createDefaultConfig() lookupsource.SourceConfig {
	_ = "STUB: not implemented"
	return *new(lookupsource.SourceConfig)
}

func createSource(
	_ context.Context,
	_ lookupsource.CreateSettings,
	cfg lookupsource.SourceConfig,
) (lookupsource.Source, error) {
	_ = "STUB: not implemented"
	return *new(lookupsource.Source), nil
}

// Create the lookup function, optionally wrapped with cache

// no start needed
// no shutdown needed

type dnsSource struct {
	recordType RecordType
	timeout    time.Duration
	resolver   *net.Resolver
}

func (s *dnsSource) lookup(ctx context.Context, key string) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

// Currently only PTR is supported

// lookupPTR performs reverse DNS lookup (IP -> hostname).
func (s *dnsSource) lookupPTR(ctx context.Context, ip string) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

// DNS errors for non-existent records should return not found, not error

// Return the first hostname, trimming trailing dot
