// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package xk8stest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/xk8stest"

import (
	"testing"

	"k8s.io/apimachinery/pkg/labels"
)

func HostEndpoint(t *testing.T) string { _ = "STUB: not implemented"; return "" }

// Prefer IPv4 gateways, but fallback to IPv6 if no IPv4 gateway is found.
// IPv6 addresses are wrapped in brackets so that callers can safely append
// ":port" (e.g. [fc00:f853:ccd:e793::1]:4317).

func SelectorFromMap(labelMap map[string]any) labels.Selector {
	_ = "STUB: not implemented"
	return *new(labels.Selector)
}
