// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"

import (
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/scraper/scrapererror"
)

type datacenterStats struct {
	ClusterStatusCounts map[types.ManagedEntityStatus]int64
	HostStats           map[string]map[types.ManagedEntityStatus]int64
	VMStats             map[string]map[types.ManagedEntityStatus]int64
	DatastoreCount      int64
	DiskCapacity        int64
	DiskFree            int64
	CPULimit            int64
	MemoryLimit         int64
}

// processDatacenterData creates all of the vCenter metrics from the stored scraped data under a single Datacenter
func (v *vcenterMetricScraper) processDatacenterData(dc *mo.Datacenter, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	// Init for current collection
	return
}

// buildDatacenterMetrics builds a resource and metrics for a given scraped vCenter Datacenter
func (v *vcenterMetricScraper) buildDatacenterMetrics(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	dcStats *datacenterStats,
) {
	_ = "STUB: not implemented"
	// Create Datacenter resource builder
	return
}

// Record & emit Datacenter metric data points

// processDatastores creates the vCenter Datastore metrics and resources from the stored scraped data under a single Datacenter
func (v *vcenterMetricScraper) processDatastores(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	dcStats *datacenterStats,
) {
	_ = "STUB: not implemented"
	return
}

// buildDatastoreMetrics builds a resource and metrics for a given scraped vCenter Datastore
func (v *vcenterMetricScraper) buildDatastoreMetrics(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	ds *mo.Datastore,
) {
	_ = "STUB: not implemented"
	// Create Datastore resource builder
	return
}

// Record & emit Datastore metric data points

// processResourcePools creates the vCenter Resource Pool metrics and resources from the stored scraped data under a single Datacenter
func (v *vcenterMetricScraper) processResourcePools(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	errs *scrapererror.ScrapeErrors,
) {
	_ = "STUB: not implemented"
	return
}

// Don't make metrics for vApps

// buildResourcePoolMetrics builds a resource and metrics for a given scraped vCenter Resource Pool
func (v *vcenterMetricScraper) buildResourcePoolMetrics(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	rp *mo.ResourcePool,
) error {
	_ = "STUB: not implemented"
	// Get related ResourcePool Compute info
	return nil
}

// Create ResourcePool resource builder

// Record & emit Resource Pool metric data points

// processHosts creates the vCenter HostS metrics and resources from the stored scraped data under a single Datacenter
//
// returns a map containing the ComputeResource info for each VM
func (v *vcenterMetricScraper) processHosts(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	dcStats *datacenterStats,
	errs *scrapererror.ScrapeErrors,
) map[string]*types.ManagedObjectReference {
	_ = "STUB: not implemented"
	return nil
}

// Populate master VM to CR relationship map from
// single Host based version of it

// buildHostMetrics builds a resource and metrics for a given scraped Host
//
// returns a map containing the ComputeResource info for each VM running on the Host
func (v *vcenterMetricScraper) buildHostMetrics(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	hs *mo.HostSystem,
) (vmRefToComputeRef map[string]*types.ManagedObjectReference, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get related Host ComputeResource info

// Store VM to ComputeResource relationship for all child VMs

// Create Host resource builder

// Record & emit Host metric data points

// processVMs creates the vCenter VM metrics and resources from the stored scraped data under a single Datacenter
//
// returns a map of all VM State counts for each ComputeResource
func (v *vcenterMetricScraper) processVMs(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	vmRefToComputeRef map[string]*types.ManagedObjectReference,
	dcStats *datacenterStats,
	errs *scrapererror.ScrapeErrors,
) map[string]*vmGroupInfo {
	_ = "STUB: not implemented"
	return nil
}

// Update master ComputeResource VM power state counts with VM power info

// ensureInnerMapInitialized is a helper function that ensures maps are initialized in order to freely aggregate VM and Host stats
func ensureInnerMapInitialized(stats map[string]map[types.ManagedEntityStatus]int64, key string) {
	_ = "STUB: not implemented"
	return
}

// buildVMMetrics builds a resource and metrics for a given scraped VM
//
// returns the ComputeResource and power state info associated with this VM
func (v *vcenterMetricScraper) buildVMMetrics(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	vm *mo.VirtualMachine,
	vmRefToComputeRef map[string]*types.ManagedObjectReference,
) (crRef *types.ManagedObjectReference, groupInfo *vmGroupInfo, err error) {
	_ = "STUB: not implemented"
	// Get related VM compute info
	return nil, nil, nil
}

// Powered-off/suspended/template VMs can have incomplete data from vSphere.

// Get related VM host info

// VMs may not have a ResourcePool reported (templates)
// But grab it if available

// Create VM resource builder

// Record VM metric data points

// processClusters creates the vCenter Cluster metrics and resources from the stored scraped data under a single Datacenter
func (v *vcenterMetricScraper) processClusters(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	vmStatesByComputeRef map[string]*vmGroupInfo,
	dcStats *datacenterStats,
	errs *scrapererror.ScrapeErrors,
) {
	_ = "STUB: not implemented"
	return
}

// Don't make metrics for anything that's not a Cluster (otherwise it should be the same as a HostSystem)

// buildClusterMetrics builds a resource and metrics for a given scraped Cluster
func (v *vcenterMetricScraper) buildClusterMetrics(
	ts pcommon.Timestamp,
	dc *mo.Datacenter,
	cr *mo.ComputeResource,
	vmGroupInfo *vmGroupInfo,
) (err error) {
	_ = "STUB: not implemented"
	// Create Cluster resource builder
	return nil
}

// Record and emit Cluster metric data points
