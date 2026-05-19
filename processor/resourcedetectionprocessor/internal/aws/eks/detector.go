// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package eks // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/eks"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	imdsprovider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/aws/ec2"
	apiprovider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/aws/eks"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/eks/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "eks"
	// Environment variable that is set when running on Kubernetes.
	kubernetesServiceHostEnvVar = "KUBERNETES_SERVICE_HOST"
	// EKS cluster version string identifier.
	eksClusterStringIdentifier = "-eks-"
	// EKS OIDC issuer identifier.
	eksOIDCIssuerIdentifier = "oidc.eks."
	// EKS IRSA token file path identifier.
	eksIRSATokenPathIdentifier = "eks.amazonaws.com" //nolint:gosec // not a credential
	// EKS Pod Identity token file path identifier.
	eksPodIdentityPathIdentifier = "eks-pod-identity"
	// Environment variable for IRSA web identity token file.
	awsWebIdentityTokenFileEnvVar = "AWS_WEB_IDENTITY_TOKEN_FILE" //nolint:gosec // env var name, not a credential
	// Environment variable for EKS Pod Identity authorization token file.
	awsContainerAuthTokenFileEnvVar = "AWS_CONTAINER_AUTHORIZATION_TOKEN_FILE" //nolint:gosec // env var name, not a credential

	imdsCheckMaxRetry = 1
)

// detector for EKS
type detector struct {
	cfg          Config
	logger       *zap.Logger
	imdsProvider imdsprovider.Provider
	apiProvider  apiprovider.Provider
	ra           metadata.ResourceAttributesConfig
	rb           *metadata.ResourceBuilder
	utils        detectorUtils
}

type eksDetectorUtils struct {
	cfg    Config
	logger *zap.Logger
}

type detectorUtils interface {
	isIMDSAccessible(ctx context.Context) bool
	getAWSConfig(ctx context.Context, testAccess bool) (aws.Config, error)
}

var _ internal.Detector = (*detector)(nil)

var _ detectorUtils = (*eksDetectorUtils)(nil)

// NewDetector returns a resource detector that will detect AWS EKS resources.
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect returns a Resource describing the Amazon EKS environment being run in.
func (d *detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	// Check if running on EKS.
	return *new(pcommon.Resource), "", nil
}

func (d *detector) detectFromIMDS(ctx context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

func (d *detector) detectFromAPI(ctx context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

func (d *detector) isEKS(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check for EKS-specific IRSA token path

// Check for EKS Pod Identity token path

// Check OIDC issuer for EKS-specific identifier

// Fallback: check cluster version for -eks- identifier

func (e eksDetectorUtils) isIMDSAccessible(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (eksDetectorUtils) getAWSConfig(ctx context.Context, checkAccess bool) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}
