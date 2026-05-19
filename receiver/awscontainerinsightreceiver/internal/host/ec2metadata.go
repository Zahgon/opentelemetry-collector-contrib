// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/host"

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"go.uber.org/zap"
)

type metadataClient interface {
	GetInstanceIdentityDocument(ctx context.Context, params *imds.GetInstanceIdentityDocumentInput, optFns ...func(*imds.Options)) (*imds.GetInstanceIdentityDocumentOutput, error)
}

type ec2MetadataProvider interface {
	getInstanceID() string
	getInstanceType() string
	getRegion() string
	getInstanceIP() string
}

type ec2Metadata struct {
	logger           *zap.Logger
	client           metadataClient
	refreshInterval  time.Duration
	instanceID       string
	instanceType     string
	instanceIP       string
	region           string
	instanceIDReadyC chan bool
	instanceIPReadyC chan bool
}

type ec2MetadataOption func(*ec2Metadata)

func newEC2Metadata(ctx context.Context, cfg aws.Config, refreshInterval time.Duration,
	instanceIDReadyC, instanceIPReadyC chan bool, logger *zap.Logger, options ...ec2MetadataOption,
) ec2MetadataProvider {
	_ = "STUB: not implemented"
	return *new(ec2MetadataProvider)
}

// stop the refresh once we get instance ID and type successfully

func (emd *ec2Metadata) refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

// notify ec2tags and ebsvolume that the instance id is ready

// notify ecsinfo that the instance id is ready

func (emd *ec2Metadata) getInstanceID() string { _ = "STUB: not implemented"; return "" }

func (emd *ec2Metadata) getInstanceType() string { _ = "STUB: not implemented"; return "" }

func (emd *ec2Metadata) getRegion() string { _ = "STUB: not implemented"; return "" }

func (emd *ec2Metadata) getInstanceIP() string { _ = "STUB: not implemented"; return "" }
