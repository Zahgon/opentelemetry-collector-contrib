// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package icmpcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/icmpcheckreceiver"

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

type pingStats struct {
	minRtt    time.Duration
	avgRtt    time.Duration
	maxRtt    time.Duration
	stdDevRtt time.Duration
	lossRatio float64
}

type pinger = interface {
	Run() error
	Stats() *pingStats
	IPString() string
	HostName() string
}

type defaultPinger struct {
	*probing.Pinger
}

func (p *defaultPinger) IPString() string { _ = "STUB: not implemented"; return "" }

func (p *defaultPinger) HostName() string { _ = "STUB: not implemented"; return "" }

func (p *defaultPinger) Stats() *pingStats { _ = "STUB: not implemented"; return nil }

func defaultPingerFactory(target PingTarget) (pinger, error) {
	_ = "STUB: not implemented"
	return *new(pinger), nil
}
