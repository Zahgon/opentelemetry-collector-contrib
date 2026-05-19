// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vpc // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/ibmcloud/vpc"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	vpcprovider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/ibmcloud/vpc"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/ibmcloud/vpc/internal/metadata"
)

const (
	// TypeStr is the detector type string.
	TypeStr = "ibmcloud_vpc"
)

var _ internal.Detector = (*Detector)(nil)

// Detector queries the IBM Cloud VPC Instance Metadata Service and emits resource attributes.
type Detector struct {
	provider vpcprovider.Provider
	logger   *zap.Logger
	rb       *metadata.ResourceBuilder
}

// NewDetector creates an IBM Cloud VPC detector.
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects IBM Cloud VPC instance metadata and returns a resource with the available attributes.
func (d *Detector) Detect(ctx context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// TODO: Use semconv constant once CloudPlatformIBMCloudVPC is added.

// regionFromZone extracts the region from a zone name.
// Example: "us-south-1" -> "us-south", "eu-de-2" -> "eu-de"
func regionFromZone(zone string) string { _ = "STUB: not implemented"; return "" }

// accountIDFromCRN extracts the account ID from a CRN.
// CRN format: crn:v1:cname:ctype:service-name:region:account-id:service-instance:resource-type:resource-id
// Example: "crn:v1:bluemix:public:is:us-south-1:a/123456::instance:0717_xxx"
// The account segment is at index 6 and may be prefixed with "a/"
func accountIDFromCRN(crn string) string { _ = "STUB: not implemented"; return "" }
