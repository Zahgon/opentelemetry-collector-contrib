// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/testutil"

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type portpair struct {
	first string
	last  string
}

// GetAvailableLocalAddress finds an available local port and returns an endpoint
// describing it. The port is available for opening when this function returns
// provided that there is no race by some other code to grab the same port
// immediately.
func GetAvailableLocalAddress(tb testing.TB) string {
	_ = "STUB: not implemented"
	// Retry has been added for windows as net.Listen can return a port that is not actually available. Details can be
	// found in https://github.com/docker/for-win/issues/3171 but to summarize Hyper-V will reserve ranges of ports
	// which do not show up under the "netstat -ano" but can only be found by
	// "netsh interface ipv4 show excludedportrange protocol=tcp".  We'll use []exclusions to hold those ranges and
	// retry if the port returned by GetAvailableLocalAddress falls in one of those them.
	return ""
}

func findAvailableAddress(tb testing.TB) string { _ = "STUB: not implemented"; return "" }

// There is a possible race if something else takes this same port before
// the test uses it, however, that is unlikely in practice.

// Get excluded ports on Windows from the command: netsh interface ipv4 show excludedportrange protocol=tcp
func getExclusionsList(tb testing.TB) []portpair { _ = "STUB: not implemented"; return nil }

func createExclusionsList(tb testing.TB, exclusionsText string) []portpair {
	_ = "STUB: not implemented"
	return nil
}

// original text may have a suffix like " - Administered port exclusions."

// UInt64ToTraceID is from collector-contrib/internal/xidutils
func UInt64ToTraceID(high, low uint64) pcommon.TraceID {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID)
}

// UInt64ToSpanID is from collector-contrib/internal/xidutils
func UInt64ToSpanID(id uint64) pcommon.SpanID {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID)
}
