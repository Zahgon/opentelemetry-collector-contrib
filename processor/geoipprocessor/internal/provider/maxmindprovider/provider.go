// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package maxmind // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider/maxmindprovider"

import (
	"context"
	"errors"
	"net/netip"

	"github.com/oschwald/geoip2-golang/v2"
	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider"
)

var (
	// defaultLanguageCode specifies English as the default Geolocation language code, see https://dev.maxmind.com/geoip/docs/web-services/responses#languages
	defaultLanguageCode = "en"
	geoIP2CityDBType    = "GeoIP2-City"
	geoLite2CityDBType  = "GeoLite2-City"

	errUnsupportedDB = errors.New("unsupported geo IP database type")
)

type maxMindProvider struct {
	geoReader *geoip2.Reader
	// language code to be used in name retrieval, e.g. "en" or "pt-BR"
	langCode string
}

var _ provider.GeoIPProvider = (*maxMindProvider)(nil)

func newMaxMindProvider(cfg *Config) (*maxMindProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Location implements provider.GeoIPProvider for MaxMind. If a non City database type is used or no metadata is found in the database, an error will be returned.
func (g *maxMindProvider) Location(_ context.Context, ipAddress netip.Addr) (attribute.Set, error) {
	_ = "STUB: not implemented"
	return *new(attribute.Set), nil
}

// Close unmaps the geo database file from virtual memory and returns the
// resources to the system.
func (g *maxMindProvider) Close(context.Context) error { _ = "STUB: not implemented"; return nil }

// cityAttributes returns a list of key-values containing geographical metadata associated to the provided IP. The key names are populated using the internal geo IP conventions package. If an invalid or nil IP is provided, an error is returned.
func (g *maxMindProvider) cityAttributes(ipAddress netip.Addr) (*[]attribute.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The exact set of top-level keys varies based on the particular GeoIP2 web service you are using. If a key maps to an undefined or empty value, it is not included in the JSON object. The following anonymous function appends the given key-value only if the value is not empty.

// city

// country

// continent

// postal code

// region

// The most specific subdivision is located at the last array position, see https://github.com/maxmind/GeoIP2-java/blob/2fe4c65424fed2c3c2449e5530381b6452b0560f/src/main/java/com/maxmind/geoip2/model/AbstractCityResponse.java#L112

// location
