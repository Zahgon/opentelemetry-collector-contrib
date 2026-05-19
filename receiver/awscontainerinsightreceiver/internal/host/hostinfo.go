// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/host"

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
)

// Info contains information about a host
type Info struct {
	cancel                context.CancelFunc
	logger                *zap.Logger
	awsConfig             aws.Config
	refreshInterval       time.Duration
	containerOrchestrator string
	instanceIDReadyC      chan bool // close of this channel indicates instance ID is ready
	instanceIPReadyC      chan bool // close of this channel indicates instance Ip is ready

	ebsVolumeReadyC chan bool // close of this channel indicates ebsVolume is initialized. It is used only in test
	ec2TagsReadyC   chan bool // close of this channel indicates ec2Tags is initialized. It is used only in test

	nodeCapacity nodeCapacityProvider
	ec2Metadata  ec2MetadataProvider
	ebsVolume    ebsVolumeProvider
	ec2Tags      ec2TagsProvider

	awsConfigCreator    func(context.Context, *zap.Logger, *awsutil.AWSSessionSettings) (aws.Config, error)
	nodeCapacityCreator func(*zap.Logger, ...nodeCapacityOption) (nodeCapacityProvider, error)
	ec2MetadataCreator  func(context.Context, aws.Config, time.Duration, chan bool, chan bool, *zap.Logger, ...ec2MetadataOption) ec2MetadataProvider
	ebsVolumeCreator    func(context.Context, aws.Config, string, string, time.Duration, *zap.Logger, ...ebsVolumeOption) ebsVolumeProvider
	ec2TagsCreator      func(context.Context, aws.Config, string, string, string, time.Duration, *zap.Logger, ...ec2TagsOption) ec2TagsProvider
}

type machineInfoOption func(*Info)

// NewInfo creates a new Info struct
func NewInfo(containerOrchestrator string, refreshInterval time.Duration, logger *zap.Logger, options ...machineInfoOption) (*Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// used in test only

func (m *Info) lazyInitEBSVolume(ctx context.Context) {
	_ = "STUB: not implemented"
	// wait until the instance id is ready
	return
}

// Because ebs volumes only change occasionally, we refresh every 5 collection intervals to reduce ec2 api calls

func (m *Info) lazyInitEC2Tags(ctx context.Context) {
	_ = "STUB: not implemented"
	// wait until the instance id is ready
	return
}

// GetInstanceID returns the ec2 instance id for the host
func (m *Info) GetInstanceID() string { _ = "STUB: not implemented"; return "" }

// GetInstanceType returns the ec2 instance type for the host
func (m *Info) GetInstanceType() string { _ = "STUB: not implemented"; return "" }

// GetRegion returns the region for the host
func (m *Info) GetRegion() string { _ = "STUB: not implemented"; return "" }

// GetInstanceIP returns the IP address of the host
func (m *Info) GetInstanceIP() string { _ = "STUB: not implemented"; return "" }

// GetNumCores returns the number of cpu cores on the host
func (m *Info) GetNumCores() int64 { _ = "STUB: not implemented"; return 0 }

// GetMemoryCapacity returns the total memory (in bytes) on the host
func (m *Info) GetMemoryCapacity() int64 { _ = "STUB: not implemented"; return 0 }

// GetEBSVolumeID returns the ebs volume id corresponding to the given device name
func (m *Info) GetEBSVolumeID(devName string) string { _ = "STUB: not implemented"; return "" }

// GetClusterName returns the cluster name associated with the host
func (m *Info) GetClusterName() string { _ = "STUB: not implemented"; return "" }

// GetInstanceIPReadyC returns the channel to show the status of host IP
func (m *Info) GetInstanceIPReadyC() chan bool { _ = "STUB: not implemented"; return nil }

// GetAutoScalingGroupName returns the auto scaling group associated with the host
func (m *Info) GetAutoScalingGroupName() string { _ = "STUB: not implemented"; return "" }

// ExtractEbsIDsUsedByKubernetes extracts the ebs volume id used by kubernetes cluster from host mount file
func (m *Info) ExtractEbsIDsUsedByKubernetes() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown stops the host Info
func (m *Info) Shutdown() { _ = "STUB: not implemented"; return }
