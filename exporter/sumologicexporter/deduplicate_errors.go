// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"

// deduplicateErrors replaces duplicate instances of the same error in a slice
// with a single error containing the number of times it occurred added as a suffix.
// For example, three occurrences of "error: 502 Bad Gateway"
// are replaced with a single instance of "error: 502 Bad Gateway (x3)".
func deduplicateErrors(errs []error) []error { _ = "STUB: not implemented"; return nil }

type errorWithCount struct {
	err   error
	count int
}
