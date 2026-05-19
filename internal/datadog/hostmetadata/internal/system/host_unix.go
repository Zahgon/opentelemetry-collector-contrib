// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
//go:build !windows

package system // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/system"

// keep as var for testing
var hostnamePath = "/bin/hostname"

func getSystemFQDN() (string, error) {
	_ = "STUB: not implemented"
	// Go does not provide a way to get the full hostname
	// so we make a best-effort by running the hostname binary
	// if available
	return "", nil
}

// if stat failed for any reason, fail silently
