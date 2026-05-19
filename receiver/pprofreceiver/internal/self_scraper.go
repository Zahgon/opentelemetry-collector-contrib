// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pprofreceiver/internal"

import (
	"bufio"
	"bytes"
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/scraper/xscraper"
)

var _ xscraper.Profiles = &SelfScraper{}

type SelfScraper struct {
	BlockProfileFraction int
	MutexProfileFraction int
	buf                  *bytes.Buffer
	writer               *bufio.Writer
}

func (hcs *SelfScraper) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*SelfScraper) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (hcs *SelfScraper) ScrapeProfiles(_ context.Context) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}
