// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"

import (
	"context"
	"time"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	vt "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan"
	"github.com/vmware/govmomi/vsan/types"
	"go.uber.org/zap"
)

// vcenterClient is a client that collects data from a vCenter endpoint.
type vcenterClient struct {
	logger         *zap.Logger
	sessionManager *session.Manager
	vimDriver      *vim25.Client
	vsanDriver     *vsan.Client
	finder         *find.Finder
	pm             *performance.Manager
	vm             *view.Manager
	cfg            *Config
}

var newVcenterClient = defaultNewVcenterClient

func defaultNewVcenterClient(l *zap.Logger, c *Config) *vcenterClient {
	_ = "STUB: not implemented"
	return nil
}

// EnsureConnection will establish a connection to the vSphere SDK if not already established
func (vc *vcenterClient) EnsureConnection(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Disconnect will logout of the authenticated session
func (vc *vcenterClient) Disconnect(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Datacenters returns the Datacenters of the vSphere SDK
func (vc *vcenterClient) Datacenters(ctx context.Context) ([]mo.Datacenter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Datastores returns the Datastores of the vSphere SDK
func (vc *vcenterClient) Datastores(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.Datastore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ComputeResources returns the ComputeResources (& ClusterComputeResources) of the vSphere SDK
func (vc *vcenterClient) ComputeResources(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.ComputeResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HostSystems returns the HostSystems of the vSphere SDK
func (vc *vcenterClient) HostSystems(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.HostSystem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResourcePools returns the ResourcePools (&VirtualApps) of the vSphere SDK
func (vc *vcenterClient) ResourcePools(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.ResourcePool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VMS returns the VirtualMachines of the vSphere SDK
func (vc *vcenterClient) VMs(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DatacenterInventoryListObjects returns the Datacenters (with populated InventoryLists) of the vSphere SDK
func (vc *vcenterClient) DatacenterInventoryListObjects(ctx context.Context) ([]*object.Datacenter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResourcePoolInventoryListObjects returns the ResourcePools (with populated InventoryLists) of the vSphere SDK
func (vc *vcenterClient) ResourcePoolInventoryListObjects(
	ctx context.Context,
	dcs []*object.Datacenter,
) ([]*object.ResourcePool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VAppInventoryListObjects returns the vApps (with populated InventoryLists) of the vSphere SDK
func (vc *vcenterClient) VAppInventoryListObjects(
	ctx context.Context,
	dcs []*object.Datacenter,
) ([]*object.VirtualApp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// perfMetricsQueryResult contains performance metric related data
type perfMetricsQueryResult struct {
	// Contains performance metrics keyed by MoRef string
	resultsByRef map[string]*performance.EntityMetric
}

// PerfMetricsQuery returns the requested performance metrics for the requested resources
// over a given sample interval and sample count
func (vc *vcenterClient) PerfMetricsQuery(
	ctx context.Context,
	spec vt.PerfQuerySpec,
	names []string,
	objs []vt.ManagedObjectReference,
) (*perfMetricsQueryResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vc *vcenterClient) batchSizeForMetrics(numMetrics int) int {
	_ = "STUB: not implemented"
	return 0
}

// vSANQueryResults contains all returned vSAN metric related data
type vSANQueryResults struct {
	// Contains vSAN metric data keyed by UUID string
	MetricResultsByUUID map[string]*vSANMetricResults
}

// vSANMetricResults contains vSAN metric related data for a single resource
type vSANMetricResults struct {
	// Contains UUID info for related resource
	UUID string
	// Contains returned metric value info for all metrics
	MetricDetails []*vSANMetricDetails
}

// vSANMetricDetails contains vSAN metric data for a single metric
type vSANMetricDetails struct {
	// Contains the metric label
	MetricLabel string
	// Contains the metric interval in seconds
	Interval int32
	// Contains timestamps for all metric values
	Timestamps []*time.Time
	// Contains all values for vSAN metric label
	Values []int64
}

// vSANQueryType represents the type of VSAN query
type vSANQueryType string

const (
	vSANQueryTypeClusters        vSANQueryType = "cluster-domclient:*"
	vSANQueryTypeHosts           vSANQueryType = "host-domclient:*"
	vSANQueryTypeVirtualMachines vSANQueryType = "virtual-machine:*"
)

// getLabelsForQueryType returns the appropriate labels for each query type
func (*vcenterClient) getLabelsForQueryType(queryType vSANQueryType) []string {
	_ = "STUB: not implemented"
	return nil
}

// VSANClusters returns back cluster vSAN performance metrics
func (vc *vcenterClient) VSANClusters(
	ctx context.Context,
	clusterRefs []*vt.ManagedObjectReference,
) (*vSANQueryResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VSANHosts returns host VSAN performance metrics for a group of clusters
func (vc *vcenterClient) VSANHosts(
	ctx context.Context,
	clusterRefs []*vt.ManagedObjectReference,
) (*vSANQueryResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VSANVirtualMachines returns virtual machine vSAN performance metrics for a group of clusters
func (vc *vcenterClient) VSANVirtualMachines(
	ctx context.Context,
	clusterRefs []*vt.ManagedObjectReference,
) (*vSANQueryResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// vSANQuery performs a vSAN query for the specified type across all clusters
func (vc *vcenterClient) vSANQuery(
	ctx context.Context,
	queryType vSANQueryType,
	clusterRefs []*vt.ManagedObjectReference,
) (*vSANQueryResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// vSANQueryByCluster performs a vSAN query for the specified type for one cluster
func (vc *vcenterClient) vSANQueryByCluster(
	ctx context.Context,
	queryType vSANQueryType,
	clusterRef *vt.ManagedObjectReference,
) (*vSANQueryResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Not all vCenters support vSAN so just return an empty result

func (vc *vcenterClient) handleVSANError(
	err error,
	queryType vSANQueryType,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (vc *vcenterClient) convertVSANResultToMetricResults(vSANResult types.VsanPerfEntityMetricCSV) (*vSANMetricResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse all timestamps

// Assuming the collector is making the request in the same time zone as the localized response
// from the vSAN API. Not a great assumption, but otherwise it will almost definitely be wrong
// if we assume that it is UTC. There is precedent for this method at least.

// Parse all metrics

func (vc *vcenterClient) convertVSANValueToMetricDetails(
	vSANValue types.VsanPerfMetricSeriesCSV,
	timestamps []time.Time,
) (*vSANMetricDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If not found assume the interval is 5m

// Match up timestamps with metric values

// uuidFromEntityRefID returns the UUID portion of the EntityRefId
func (*vcenterClient) uuidFromEntityRefID(id string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
