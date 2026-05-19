// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/internal"

// GOOSToOSType maps a runtime.GOOS-like value to os.type style.
func GOOSToOSType(goos string) string { _ = "STUB: not implemented"; return "" }

func GOARCHtoHostArch(goarch string) string {
	_ = "STUB: not implemented"
	// These cases differ from the spec well-known values
	return ""
}

// Other cases either match the spec or are not well-known (so we use a custom value)
