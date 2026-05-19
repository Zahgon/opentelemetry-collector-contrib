// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"

import (
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal/metadata"
)

// recordDatacenterStats records stat metrics for a vSphere Datacenter
func (v *vcenterMetricScraper) recordDatacenterStats(
	ts pcommon.Timestamp,
	dcStat *datacenterStats,
) {
	_ = "STUB: not implemented"
	// Cluster metrics
	return
}

// VM metrics

// Host metrics

// Datacenter stats

func getEntityStatusAttribute(status types.ManagedEntityStatus) (metadata.AttributeEntityStatus, bool) {
	_ = "STUB: not implemented"
	return *new(metadata.AttributeEntityStatus), false
}

func getVMPowerStateAttribute(state string) (metadata.AttributeVMCountPowerState, bool) {
	_ = "STUB: not implemented"
	return *new(metadata.AttributeVMCountPowerState), false
}

func getHostPowerStateAttribute(state string) (metadata.AttributeHostPowerState, bool) {
	_ = "STUB: not implemented"
	return *new(metadata.AttributeHostPowerState), false
}

// recordDatastoreStats records stat metrics for a vSphere Datastore
func (v *vcenterMetricScraper) recordDatastoreStats(
	ts pcommon.Timestamp,
	ds *mo.Datastore,
) {
	_ = "STUB: not implemented"
	return
}

// recordClusterStats records stat metrics for a vSphere Cluster
func (v *vcenterMetricScraper) recordClusterStats(
	ts pcommon.Timestamp,
	cr *mo.ComputeResource,
	vmGroupInfo *vmGroupInfo,
) {
	_ = "STUB: not implemented"
	return
}

// recordClusterVSANMetrics records vSAN metrics for a vSphere Cluster
func (v *vcenterMetricScraper) recordClusterVSANMetrics(vSANMetrics *vSANMetricResults) {
	_ = "STUB: not implemented"
	return
}

// recordResourcePoolStats records stat metrics for a vSphere Resource Pool
func (v *vcenterMetricScraper) recordResourcePoolStats(
	ts pcommon.Timestamp,
	rp *mo.ResourcePool,
) {
	_ = "STUB: not implemented"
	return
}

// recordClusterStats records stat metrics for a vSphere Host
func (v *vcenterMetricScraper) recordHostSystemStats(
	ts pcommon.Timestamp,
	hs *mo.HostSystem,
) {
	_ = "STUB: not implemented"
	return
}

// recordHostVSANMetrics records vSAN metrics for a vSphere host
func (v *vcenterMetricScraper) recordHostVSANMetrics(vSANMetrics *vSANMetricResults) {
	_ = "STUB: not implemented"
	return
}

// recordVMStats records stat metrics for a vSphere Virtual Machine
func (v *vcenterMetricScraper) recordVMStats(
	ts pcommon.Timestamp,
	vm *mo.VirtualMachine,
	hs *mo.HostSystem,
) {
	_ = "STUB: not implemented"
	return
}

// Most likely the VM is unavailable or is unreachable.

// https://communities.vmware.com/t5/VMware-code-Documents/Resource-Management/ta-p/2783456
// VirtualMachine.runtime.maxCpuUsage is a property of the virtual machine, indicating the limit value.
// This value is always equal to the limit value set for that virtual machine.
// If no limit, it has full host mhz * vm.Config.Hardware.NumCPU.

// This shouldn't happen, but protect against division by zero.

var hostPerfMetricList = []string{
	// network metrics
	"net.bytesTx.average",
	"net.bytesRx.average",
	"net.packetsTx.summation",
	"net.packetsRx.summation",
	"net.usage.average",
	"net.errorsRx.summation",
	"net.errorsTx.summation",
	"net.droppedTx.summation",
	"net.droppedRx.summation",
	// disk metrics
	"disk.totalReadLatency.average",
	"disk.totalWriteLatency.average",
	"disk.maxTotalLatency.latest",
	"disk.read.average",
	"disk.write.average",
	// cpu metrics
	"cpu.reservedCapacity.average",
	"cpu.totalCapacity.average",
}

// recordHostPerformanceMetrics records performance metrics for a vSphere Host
func (v *vcenterMetricScraper) recordHostPerformanceMetrics(entityMetric *performance.EntityMetric) {
	_ = "STUB: not implemented"
	return
}

/******************************************/
// Performance Monitoring Level 1 Metrics //
/******************************************/
// (per device requires level 3)

// (per device requires level 4)

/******************************************/
// Following Requires Performance Level 2 //
/******************************************/
// (per device requires level 3)

// vmPerfMetricList may be customizable in the future but here is the full list of Virtual Machine Performance Counters
// https://docs.vmware.com/en/vRealize-Operations/8.6/com.vmware.vcom.metrics.doc/GUID-1322F5A4-DA1D-481F-BBEA-99B228E96AF2.html
var vmPerfMetricList = []string{
	// network metrics
	"net.packetsTx.summation",
	"net.packetsRx.summation",
	"net.droppedTx.summation",
	"net.droppedRx.summation",
	"net.bytesRx.average",
	"net.bytesTx.average",
	"net.usage.average",
	"net.broadcastRx.summation",
	"net.broadcastTx.summation",
	"net.multicastRx.summation",
	"net.multicastTx.summation",

	// disk metrics
	"disk.totalWriteLatency.average",
	"disk.totalReadLatency.average",
	"disk.maxTotalLatency.latest",
	"virtualDisk.totalWriteLatency.average",
	"virtualDisk.totalReadLatency.average",
	"virtualDisk.read.average",
	"virtualDisk.write.average",

	// cpu metrics
	"cpu.idle.summation",
	"cpu.wait.summation",
	"cpu.ready.summation",
}

// recordVMPerformanceMetrics records performance metrics for a vSphere Virtual Machine
func (v *vcenterMetricScraper) recordVMPerformanceMetrics(entityMetric *performance.EntityMetric) {
	_ = "STUB: not implemented"
	return
}

/******************************************/
// Performance Monitoring Level 1 Metrics //
/******************************************/
// (per device requires level 3)

// (per device requires level 4)

/******************************************/
// Following Requires Performance Level 2 //
/******************************************/
// (per device requires level 2)

// (per device requires level 3)

// recordVMVSANMetrics records vSAN metrics for a vSphere Virtual Machine
func (v *vcenterMetricScraper) recordVMVSANMetrics(vSANMetrics *vSANMetricResults) {
	_ = "STUB: not implemented"
	return
}
