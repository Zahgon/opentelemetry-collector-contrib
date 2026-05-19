// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

// The code for sanitize is mostly copied from:
//  https://github.com/open-telemetry/opentelemetry-collector/blob/2e84285efc665798d76773b9901727e8836e9d8f/exporter/prometheusexporter/sanitize.go

// sanitize replaces non-alphanumeric characters with underscores in s.
func sanitize(s string) string { _ = "STUB: not implemented"; return "" }

// Note: No length limit for label keys because Prometheus doesn't
// define a length limit, thus we should NOT be truncating label keys.
// See https://github.com/orijtech/prometheus-go-metrics-exporter/issues/4.

// converts anything that is not a letter or digit to an underscore
func sanitizeRune(r rune) rune { _ = "STUB: not implemented"; return 0 }

// Everything else turns into an underscore
