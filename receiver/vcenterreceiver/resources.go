// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"

import (
	"github.com/vmware/govmomi/vim25/mo"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal/metadata"
)

// createDatastoreResourceBuilder returns a ResourceBuilder with
// attributes set for a vSphere Datastore
func (v *vcenterMetricScraper) createDatastoreResourceBuilder(
	dc *mo.Datacenter,
	ds *mo.Datastore,
) *metadata.ResourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// createDatacenterResourceBuilder returns a ResourceBuilder with
// attributes set for a vSphere datacenter
func (v *vcenterMetricScraper) createDatacenterResourceBuilder(
	dc *mo.Datacenter,
) *metadata.ResourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// createClusterResourceBuilder returns a ResourceBuilder with
// attributes set for a vSphere Cluster
func (v *vcenterMetricScraper) createClusterResourceBuilder(
	dc *mo.Datacenter,
	cr *mo.ComputeResource,
) *metadata.ResourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// createResourcePoolResourceBuilder returns a ResourceBuilder with
// attributes set for a vSphere Resource Pool
func (v *vcenterMetricScraper) createResourcePoolResourceBuilder(
	dc *mo.Datacenter,
	cr *mo.ComputeResource,
	rp *mo.ResourcePool,
) (*metadata.ResourceBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createHostResourceBuilder returns a ResourceBuilder with
// attributes set for a vSphere Host
func (v *vcenterMetricScraper) createHostResourceBuilder(
	dc *mo.Datacenter,
	cr *mo.ComputeResource,
	hs *mo.HostSystem,
) *metadata.ResourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// createVMResourceBuilder returns a ResourceBuilder with
// attributes set for a vSphere Virtual Machine
func (v *vcenterMetricScraper) createVMResourceBuilder(
	dc *mo.Datacenter,
	cr *mo.ComputeResource,
	hs *mo.HostSystem,
	rp *mo.ResourcePool,
	vm *mo.VirtualMachine,
) (*metadata.ResourceBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
