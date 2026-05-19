// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zookeeperscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/scraper/zookeeperscraper"

import (
	"bufio"
	"context"
	"net"
	"regexp"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/scraper/zookeeperscraper/internal/metadata"
)

var zookeeperFormatRE = regexp.MustCompile(`(^zk_\w+)\s+([\w.\-]+)`)

const (
	mntrCommand = "mntr"
	ruokCommand = "ruok"
)

type zookeeperMetricsScraper struct {
	component.StartFunc
	logger *zap.Logger
	config *Config
	cancel context.CancelFunc
	rb     *metadata.ResourceBuilder
	mb     *metadata.MetricsBuilder

	// For mocking.
	closeConnection       func(net.Conn) error
	setConnectionDeadline func(net.Conn, time.Time) error
	sendCmd               func(net.Conn, string) (*bufio.Scanner, error)
}

func newZookeeperMetricsScraper(settings scraper.Settings, config *Config) *zookeeperMetricsScraper {
	_ = "STUB: not implemented"
	return nil
}

func (z *zookeeperMetricsScraper) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (z *zookeeperMetricsScraper) ScrapeMetrics(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (z *zookeeperMetricsScraper) runCommand(ctx context.Context, command string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *zookeeperMetricsScraper) processMntr(response []string) { _ = "STUB: not implemented"; return }

// Skip metric if there is no descriptor associated with it.

// Unexported metric, just move to the next line.

// zk_avg_latency changed to float in ZK 3.7+; truncate to int64.

// Generate computed metrics

func (z *zookeeperMetricsScraper) processRuok(response []string) { _ = "STUB: not implemented"; return }

func closeConnection(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func setConnectionDeadline(conn net.Conn, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func sendCmd(conn net.Conn, cmd string) (*bufio.Scanner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
