// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package collectd // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/collectd"

// LabelsFromName tries to pull out dimensions out of name in the format
// "name[k=v,f=x]-more_name".
// For the example above it would return "name-more_name" and extract dimensions
// (k,v) and (f,x).
// If something unexpected is encountered it returns the original metric name.
//
// The code tries to avoid allocation by using local slices and avoiding calls
// to functions like strings.Slice.
func LabelsFromName(val *string) (metricName string, labels map[string]string) {
	_ = "STUB: not implemented"
	return "", nil
}
