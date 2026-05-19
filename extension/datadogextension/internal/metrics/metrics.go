// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/metrics"

import (
	"github.com/DataDog/datadog-agent/pkg/metrics"
	"go.opentelemetry.io/collector/component"
)

// TagsFromBuildInfo returns a list of tags derived from buildInfo to be used when creating metrics.
func TagsFromBuildInfo(buildInfo component.BuildInfo) []string {
	_ = "STUB: not implemented"
	return nil
}

// CreateLivenessSerie creates a liveness metric serie to report that the extension is running.
// The timestamp should be in Unix nanoseconds.
func CreateLivenessSerie(hostname string, timestampNs uint64, tags []string) *metrics.Serie {
	_ = "STUB: not implemented"
	// Transform UnixNano timestamp into Unix timestamp (seconds)
	return nil
}
