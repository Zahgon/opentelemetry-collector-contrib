// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/testutil"

import (
	"testing"

	"go.opentelemetry.io/collector/featuregate"
)

type portpair struct {
	first string
	last  string
}

// GetAvailableLocalAddress finds an available local port on tcp network and returns an endpoint
// describing it. The port is available for opening when this function returns
// provided that there is no race by some other code to grab the same port
// immediately.
func GetAvailableLocalAddress(tb testing.TB) string { _ = "STUB: not implemented"; return "" }

// GetAvailableLocalNetworkAddress finds an available local port on specified network and returns an endpoint
// describing it. The port is available for opening when this function returns
// provided that there is no race by some other code to grab the same port
// immediately.
func GetAvailableLocalNetworkAddress(tb testing.TB, network string) string {
	_ = "STUB: not implemented"
	// Retry has been added for windows as net.Listen can return a port that is not actually available. Details can be
	// found in https://github.com/docker/for-win/issues/3171 but to summarize Hyper-V will reserve ranges of ports
	// which do not show up under the "netstat -ano" but can only be found by
	// "netsh interface ipv4 show excludedportrange protocol=tcp".  We'll use []exclusions to hold those ranges and
	// retry if the port returned by GetAvailableLocalAddress falls in one of those them.
	return ""
}

func findAvailableAddress(tb testing.TB, network string) string {
	_ = "STUB: not implemented"

	// net.Listen supported network strings
	return ""
}

// There is a possible race if something else takes this same port before
// the test uses it, however, that is unlikely in practice.

// net.ListenPacket supported network strings

// There is a possible race if something else takes this same port before
// the test uses it, however, that is unlikely in practice.

// Get excluded ports on Windows from the command: netsh interface ipv4 show excludedportrange protocol=tcp
func getExclusionsList(tb testing.TB) []portpair { _ = "STUB: not implemented"; return nil }

func createExclusionsList(tb testing.TB, exclusionsText string) []portpair {
	_ = "STUB: not implemented"
	return nil
}

// original text may have a suffix like " - Administered port exclusions."

// Force the state of feature gate for a test
// usage: defer SetFeatureGateForTest("gateName", true)()
func SetFeatureGateForTest(tb testing.TB, gate *featuregate.Gate, enabled bool) func() {
	_ = "STUB: not implemented"
	return nil
}

func GetAvailablePort(tb testing.TB) int { _ = "STUB: not implemented"; return 0 }

// EndpointForPort gets the endpoint for a given port using localhost.
func EndpointForPort(port int) string { _ = "STUB: not implemented"; return "" }

// TempDir creates a temporary directory in a safe way. On Linux, it just calls tb.TempDir.
//
// On Windows, it uses os.MkdirTemp to create directory since t.TempDir results in an error during cleanup in scoped-tests,
// possibly due to an interaction with the -count argument in `go test`. It does not do any cleanup (i.e. the directory is left on the filesystem).
// Prefer using tb.TempDir directly if possible.
//
// See https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/42639
func TempDir(tb testing.TB) string { _ = "STUB: not implemented"; return "" }

//nolint:usetesting

// sanitizePattern so that it works correctly with os.MkdirTemp.
//
// The following code is taken from the t.TempDir implementation.
// Copyright 2009 The Go Authors. All rights reserved.
// Taken from https://cs.opensource.google/go/go/+/refs/tags/go1.25.1:src/testing/testing.go;l=1340-1364;drc=49cdf0c42e320dfed044baa551610f081eafb781
func sanitizePattern(pattern string) string {
	_ = "STUB: not implemented"
	// Limit length of file names on disk.
	// Invalid runes from slicing are dropped by strings.Map below.
	return ""
}

// Drop unusual characters (such as path separators or
// characters interacting with globs) from the directory name to
// avoid surprising os.MkdirTemp behavior.
