// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor"

import (
	"context"
	"errors"
	"net/netip"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider"
)

var (
	errIPNotFound        = errors.New("no IP address found in the resource attributes")
	errUnspecifiedIP     = errors.New("unspecified address")
	errUnspecifiedSource = errors.New("no source attributes defined")
)

// newGeoIPProcessor creates a new instance of geoIPProcessor with the specified fields.
type geoIPProcessor struct {
	providers []provider.GeoIPProvider
	logger    *zap.Logger

	cfg *Config
}

func newGeoIPProcessor(processorConfig *Config, providers []provider.GeoIPProvider, params processor.Settings) *geoIPProcessor {
	_ = "STUB: not implemented"
	return nil
}

// parseIP parses a string to a net.IP type and returns an error if the IP is invalid or unspecified.
func parseIP(strIP string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// ipFromAttributes extracts an IP address from the given attributes based on the specified fields.
// It returns the first IP address if found, or an error if no valid IP address is found.
func ipFromAttributes(attributes []attribute.Key, resource pcommon.Map) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// The attribute might contain a domain name. Skip any net.ParseIP error until we have a fine-grained error propagation strategy.
// TODO: propagate an error once error_mode configuration option is available (e.g. transformprocessor)

// geoLocation fetches geolocation information for the given IP address using the configured providers.
// It returns a set of attributes containing the geolocation data, or an error if the location could not be determined.
func (g *geoIPProcessor) geoLocation(ctx context.Context, ip netip.Addr) (attribute.Set, error) {
	_ = "STUB: not implemented"
	return *new(attribute.Set), nil
}

// continue if no metadata is found

// processAttributes processes a pcommon.Map by adding geolocation attributes based on the found IP address.
func (g *geoIPProcessor) processAttributes(ctx context.Context, metadata pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: log IP error not found

func (g *geoIPProcessor) shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
