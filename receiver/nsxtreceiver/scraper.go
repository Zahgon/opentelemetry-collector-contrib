// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nsxtreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver/internal/metadata"
	dm "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver/internal/model"
)

type nsxtScraper struct {
	config   *Config
	settings component.TelemetrySettings
	host     component.Host
	client   Client
	mb       *metadata.MetricsBuilder
}

func newScraper(cfg *Config, settings receiver.Settings) *nsxtScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *nsxtScraper) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

type nodeClass int

const (
	transportClass nodeClass = iota
	managerClass
)

func (s *nsxtScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

type nodeInfo struct {
	nodeProps  dm.NodeProperties
	nodeType   string
	interfaces []interfaceInformation
	stats      *dm.NodeStatus
}

type interfaceInformation struct {
	iFace dm.NetworkInterface
	stats *dm.NetworkInterfaceStats
}

func (s *nsxtScraper) retrieve(ctx context.Context) ([]*nodeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no useful stats are recorded for controller nodes

func (s *nsxtScraper) retrieveInterfaces(
	ctx context.Context,
	nodeProps dm.NodeProperties,
	nodeInfo *nodeInfo,
	nodeClass nodeClass,
	wg *sync.WaitGroup,
	errs *scrapererror.ScrapeErrors,
) {
	_ = "STUB: not implemented"
	return
}

func (s *nsxtScraper) retrieveNodeStats(
	ctx context.Context,
	nodeProps dm.NodeProperties,
	nodeInfo *nodeInfo,
	nodeClass nodeClass,
	wg *sync.WaitGroup,
	errs *scrapererror.ScrapeErrors,
) {
	_ = "STUB: not implemented"
	return
}

func (s *nsxtScraper) process(
	nodes []*nodeInfo,
	colTime pcommon.Timestamp,
) {
	_ = "STUB: not implemented"
	return
}

func (s *nsxtScraper) recordNodeInterface(colTime pcommon.Timestamp, nodeProps dm.NodeProperties, i interfaceInformation) {
	_ = "STUB: not implemented"
	return
}

func (s *nsxtScraper) recordNode(
	colTime pcommon.Timestamp,
	info *nodeInfo,
) {
	_ = "STUB: not implemented"
	return
}

// ensure division by zero is safeguarded

func clusterNodeType(node dm.ClusterNode) string { _ = "STUB: not implemented"; return "" }
