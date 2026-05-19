// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/host"

import (
	"context"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"go.uber.org/zap"
)

var ebsMountPointRegex = regexp.MustCompile(`kubernetes\.io/aws-ebs/mounts/aws/(.+)/(vol-\w+)$`)

type ebsVolumeClient interface {
	DescribeVolumes(ctx context.Context, params *ec2.DescribeVolumesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error)
}

type ebsVolumeProvider interface {
	getEBSVolumeID(devName string) string
	extractEbsIDsUsedByKubernetes() map[string]string
}

type ebsVolume struct {
	refreshInterval time.Duration
	maxJitterTime   time.Duration
	instanceID      string
	client          ebsVolumeClient
	logger          *zap.Logger
	shutdownC       chan bool

	mu sync.RWMutex
	// device name to volumeID mapping
	dev2Vol map[string]string

	// for testing only
	hostMounts   string
	osLstat      func(name string) (os.FileInfo, error)
	evalSymLinks func(path string) (string, error)
}

type ebsVolumeOption func(*ebsVolume)

func newEBSVolume(ctx context.Context, cfg aws.Config, instanceID, region string,
	refreshInterval time.Duration, logger *zap.Logger, options ...ebsVolumeOption,
) ebsVolumeProvider {
	_ = "STUB: not implemented"
	return *new(ebsVolumeProvider)
}

// keep refreshing to get updated ebs volumes

func (e *ebsVolume) refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *ebsVolume) addEBSVolumeMapping(zone *string, attachment ec2types.VolumeAttachment) string {
	_ = "STUB: not implemented"
	// *attachment.Device is sth like: /dev/xvda
	return ""
}

// find nvme block name by symlink, if symlink doesn't exist, return ""
func (e *ebsVolume) findNvmeBlockNameIfPresent(devName string) string {
	_ = "STUB: not implemented"
	// for nvme(ssd), there is a symlink from devName to nvme block name, i.e. /dev/xvda -> /dev/nvme0n1
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nvme-ebs-volumes.html
	return ""
}

func (e *ebsVolume) getEBSVolumeID(devName string) string { _ = "STUB: not implemented"; return "" }

// The key of dev2Vol is device name like nvme0n1, while the input devName could be a partition name like nvme0n1p1

// extract the ebs volume id used by kubernetes cluster
func (e *ebsVolume) extractEbsIDsUsedByKubernetes() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// example line: /dev/nvme1n1 /var/lib/kubelet/plugins/kubernetes.io/aws-ebs/mounts/aws/us-west-2b/vol-0d9f0816149eb2050 ext4 rw,relatime,data=ordered 0 0

// Set {"/dev/nvme1n1": "aws://us-west-2b/vol-0d9f0816149eb2050"}
