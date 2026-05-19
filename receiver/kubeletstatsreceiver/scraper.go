// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubeletstatsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver"

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

type scraperOptions struct {
	collectionInterval    time.Duration
	extraMetadataLabels   []kubelet.MetadataLabel
	metricGroupsToCollect map[kubelet.MetricGroup]bool
	allNetworkInterfaces  map[kubelet.MetricGroup]bool
	k8sAPIClient          kubernetes.Interface
}

type kubeletScraper struct {
	statsProvider         *kubelet.StatsProvider
	metadataProvider      *kubelet.MetadataProvider
	logger                *zap.Logger
	extraMetadataLabels   []kubelet.MetadataLabel
	metricGroupsToCollect map[kubelet.MetricGroup]bool
	allNetworkInterfaces  map[kubelet.MetricGroup]bool
	k8sAPIClient          kubernetes.Interface
	cachedVolumeSource    map[string]v1.PersistentVolumeSource
	mbs                   *metadata.MetricsBuilders
	needsResources        bool
	nodeInformer          cache.SharedInformer
	stopCh                chan struct{}
	m                     sync.RWMutex

	// A struct that keeps Node's information
	nodeInfo *kubelet.NodeInfo
}

func newKubeletScraper(
	restClient kubelet.RestClient,
	set receiver.Settings,
	rOptions *scraperOptions,
	metricsConfig metadata.MetricsBuilderConfig,
	nodeName string,
) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}

func (r *kubeletScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// fetch metadata only when extra metadata labels are needed

func (r *kubeletScraper) detailedPVCLabelsSetter() func(rb *metadata.ResourceBuilder, volCacheID, volumeClaim, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// Cache collected source.

func (r *kubeletScraper) node() kubelet.NodeInfo {
	_ = "STUB: not implemented"
	return *new(kubelet.NodeInfo)
}

func (r *kubeletScraper) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *kubeletScraper) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *kubeletScraper) handleNodeAdd(obj any) { _ = "STUB: not implemented"; return }

func (r *kubeletScraper) handleNodeUpdate(_, newNode any) { _ = "STUB: not implemented"; return }

func (r *kubeletScraper) addOrUpdateNode(node *v1.Node) { _ = "STUB: not implemented"; return }

// ie: 32564740Ki
