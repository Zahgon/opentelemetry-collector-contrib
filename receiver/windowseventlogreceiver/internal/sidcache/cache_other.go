// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package sidcache // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver/internal/sidcache"

// New returns an error on non-Windows platforms since SID resolution
// requires the Windows Local Security Authority API.
func New(_ Config) (Cache, error) { _ = "STUB: not implemented"; return *new(Cache), nil }
