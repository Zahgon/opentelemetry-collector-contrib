// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package grafanacloudconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/grafanacloudconnector"

import (
	"sync"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

type hostMetrics struct {
	mutex sync.RWMutex
	hosts map[string]struct{}
}

func newHostMetrics() *hostMetrics { _ = "STUB: not implemented"; return nil }

func (h *hostMetrics) add(hostName string) { _ = "STUB: not implemented"; return }

func (h *hostMetrics) count() int { _ = "STUB: not implemented"; return 0 }

func (h *hostMetrics) metrics() (*pmetric.Metrics, int) { _ = "STUB: not implemented"; return nil, 0 }

func (h *hostMetrics) reset() { _ = "STUB: not implemented"; return }
