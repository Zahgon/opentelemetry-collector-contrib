// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheus // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus/internal/metadata"
)

var DropSanitizationGate = metadata.PkgTranslatorPrometheusPermissiveLabelSanitizationFeatureGate

// Normalizes the specified label to follow Prometheus label names standard
//
// See rules at https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels
//
// Labels that start with non-letter rune will be prefixed with "key_"
//
// Exception is made for double-underscores which are allowed
func NormalizeLabel(label string) string {
	_ = "STUB: not implemented"
	// Trivial case
	return ""
}

// Replace all non-alphanumeric runes with underscores

// If label starts with a number, prepend with "key_"

// Return '_' for anything non-alphanumeric
func sanitizeRune(r rune) rune { _ = "STUB: not implemented"; return 0 }
