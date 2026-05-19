// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package url // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor/internal/url"

import (
	"github.com/grafana/clusterurl/pkg/clusterurl"
)

type URLSanitizer struct {
	classifier *clusterurl.ClusterURLClassifier
	attributes map[string]bool
}

func NewURLSanitizer(config URLSanitizationConfig) (*URLSanitizer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *URLSanitizer) SanitizeAttributeURL(url, attributeKey string) string {
	_ = "STUB: not implemented"
	return ""
}

// SanitizeURL sanitizes the given URL by removing any gibberish words.
// https://github.com/open-telemetry/opentelemetry-ebpf-instrumentation/blob/38ca7938595409b8ffe6b897c14a0e3280dd2941/pkg/components/transform/route/cluster.go#L48
func (s *URLSanitizer) SanitizeURL(url string) string { _ = "STUB: not implemented"; return "" }
