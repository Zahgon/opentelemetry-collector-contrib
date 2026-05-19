// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tinybirdexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tinybirdexporter"

// option represents a configuration option that can be passed to the tinybird exporter.
type option func(*tinybirdExporter) error

// withMaxRequestBodySize sets the maximum size of the request body in bytes.
func withMaxRequestBodySize(size int) option { _ = "STUB: not implemented"; return *new(option) }
