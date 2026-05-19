// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
//go:build windows

package docker // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/docker"

import (
	"net"
	"testing"
)

func testListener(t *testing.T) (net.Listener, string) {
	_ = "STUB: not implemented"
	return *new(net.Listener), ""
}
