// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sshcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver"

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver/internal/configssh"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver/internal/metadata"
)

var errClientNotInit = errors.New("client not initialized")

type sshcheckScraper struct {
	*configssh.Client
	*Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

// start starts the scraper by creating a new SSH Client on the scraper
func (s *sshcheckScraper) start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sshcheckScraper) scrapeSSH(now pcommon.Timestamp) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sshcheckScraper) scrapeSFTP(now pcommon.Timestamp) error {
	_ = "STUB: not implemented"
	return nil
}

// upgrade to SFTP and read fs

// timeout chooses the shorter duration between a given deadline and timeout
func timeout(deadline time.Time, timeout time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// scrape connects to the endpoint and produces metrics based on the response. TBH the flow-of-control
// is a bit awkward here, because the SFTP checks are not enabled by default and they would panic on nil
// ref to the underlying Conn when SSH checks failed.
func (s *sshcheckScraper) scrape(ctx context.Context) (_ pmetric.Metrics, err error) {
	_ = "STUB: not implemented"
	return *

	// check cancellation
	new(pmetric.Metrics), nil
}

// if the context carries a shorter deadline then timeout that quickly

func newScraper(conf *Config, settings receiver.Settings) *sshcheckScraper {
	_ = "STUB: not implemented"
	return nil
}
