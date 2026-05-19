// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"
import (
	"context"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal/metadata"
)

var _ receiver.Metrics = (*vcenterMetricScraper)(nil)

type vmGroupInfo struct {
	poweredOn  int64
	poweredOff int64
	suspended  int64
	templates  int64
}

type vcenterScrapeData struct {
	datacenters              []*mo.Datacenter
	datastores               []*mo.Datastore
	clusterRefs              []*types.ManagedObjectReference
	rPoolIPathsByRef         map[string]*string
	vAppIPathsByRef          map[string]*string
	rPoolsByRef              map[string]*mo.ResourcePool
	computesByRef            map[string]*mo.ComputeResource
	hostsByRef               map[string]*mo.HostSystem
	hostPerfMetricsByRef     map[string]*performance.EntityMetric
	vmsByRef                 map[string]*mo.VirtualMachine
	vmPerfMetricsByRef       map[string]*performance.EntityMetric
	vmVSANMetricsByUUID      map[string]*vSANMetricResults
	hostVSANMetricsByUUID    map[string]*vSANMetricResults
	clusterVSANMetricsByUUID map[string]*vSANMetricResults
}

type vcenterMetricScraper struct {
	client     *vcenterClient
	config     *Config
	mb         *metadata.MetricsBuilder
	logger     *zap.Logger
	scrapeData *vcenterScrapeData
}

func newVmwareVcenterScraper(
	logger *zap.Logger,
	config *Config,
	settings receiver.Settings,
) *vcenterMetricScraper {
	_ = "STUB: not implemented"
	return nil
}

func newVcenterScrapeData() *vcenterScrapeData { _ = "STUB: not implemented"; return nil }

func (v *vcenterMetricScraper) hasEnabledVSANMetrics() bool {
	_ = "STUB: not implemented"
	return false

	// Check cluster vSAN metrics
}

// Check host vSAN metrics

// Check VM vSAN metrics

func (v *vcenterMetricScraper) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// don't fail to start if we cannot establish connection, just log an error

func (v *vcenterMetricScraper) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *vcenterMetricScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// ensure connection before scraping

// scrapeAndProcessAllMetrics collects & converts all relevant resources managed by vCenter to OTEL resources & metrics
func (v *vcenterMetricScraper) scrapeAndProcessAllMetrics(ctx context.Context, errs *scrapererror.ScrapeErrors) error {
	_ = "STUB: not implemented"
	return nil
}

// Build metrics now that all vCenter data has been scraped for a single datacenter

// Clear scrape data

// scrapeDatacenterInventoryListObjects scrapes and stores all Datacenter objects with their InventoryLists
func (v *vcenterMetricScraper) scrapeDatacenterInventoryListObjects(
	ctx context.Context,
	errs *scrapererror.ScrapeErrors,
) []*object.Datacenter {
	_ = "STUB: not implemented"
	// Get Datacenters with InventoryLists and store for later retrieval
	return nil
}

// scrapeResourcePoolInventoryListObjects scrapes and stores all ResourcePool objects with their InventoryLists
func (v *vcenterMetricScraper) scrapeResourcePoolInventoryListObjects(
	ctx context.Context,
	dcs []*object.Datacenter,
	errs *scrapererror.ScrapeErrors,
) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// Get ResourcePools with InventoryLists and store for later retrieval

// scrapeVAppInventoryListObjects scrapes and stores all vApp objects with their InventoryLists
func (v *vcenterMetricScraper) scrapeVAppInventoryListObjects(
	ctx context.Context,
	dcs []*object.Datacenter,
	errs *scrapererror.ScrapeErrors,
) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// Get vApps with InventoryLists and store for later retrieval

// scrapeDatacenters scrapes and stores all relevant property data for all Datacenters
func (v *vcenterMetricScraper) scrapeDatacenters(ctx context.Context, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// Get Datacenters w/properties and store for later retrieval

// scrapeDatastores scrapes and stores all relevant property data for a Datacenter's Datastores
func (v *vcenterMetricScraper) scrapeDatastores(ctx context.Context, dc *mo.Datacenter, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// Get Datastores w/properties and store for later retrieval

// scrapeComputes scrapes and stores all relevant property data for a Datacenter's ComputeResources/ClusterComputeResources
func (v *vcenterMetricScraper) scrapeComputes(ctx context.Context, dc *mo.Datacenter, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// Get ComputeResources/ClusterComputeResources w/properties and store for later retrieval

// Get all Cluster vSAN metrics and store for later retrieval (only if vSAN metrics are enabled)

// scrapeHosts scrapes and stores all relevant metric/property data for a Datacenter's HostSystems
func (v *vcenterMetricScraper) scrapeHosts(ctx context.Context, dc *mo.Datacenter, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// Get HostSystems w/properties and store for later retrieval

// Just grabbing real time performance metrics of the current
// supported metrics by this receiver. If more are added we may need
// a system of making this user customizable or adapt to use a 5 minute interval per metric

// Get all HostSystem performance metrics and store for later retrieval

// Get all Host vSAN metrics and store for later retrieval

// scrapeResourcePools scrapes and stores all relevant property data for a Datacenter's ResourcePools/vApps
func (v *vcenterMetricScraper) scrapeResourcePools(ctx context.Context, dc *mo.Datacenter, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// Get ResourcePools/vApps w/properties and store for later retrieval

// scrapeVirtualMachines scrapes and stores all relevant metric/property data for a Datacenter's VirtualMachines
func (v *vcenterMetricScraper) scrapeVirtualMachines(ctx context.Context, dc *mo.Datacenter, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// Get VirtualMachines w/properties and store for later retrieval

// Just grabbing real time performance metrics of the current
// supported metrics by this receiver. If more are added we may need
// a system of making this user customizable or adapt to use a 5 minute interval per metric

// Get all VirtualMachine performance metrics and store for later retrieval

// Get all VirtualMachine vSAN metrics and store for later retrieval
