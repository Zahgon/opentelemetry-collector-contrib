// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package serializeprofiles // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer/serializeprofiles"

import (
	"encoding/json"
)

// unixTime64 represents nanoseconds since epoch.
type unixTime64 uint64

// newUnixTime64 creates a unixTime64 from either seconds or nanoseconds since the epoch.
func newUnixTime64(t uint64) unixTime64 { _ = "STUB: not implemented"; return *new(unixTime64) }

func (t unixTime64) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Nanoseconds, ES does not support 'epoch_nanoseconds' so
	// we have to pass it a value formatted as 'strict_date_optional_time_nanos'.
	return nil, nil
}

// Compile-time interface checks
var _ json.Marshaler = (*unixTime64)(nil)
