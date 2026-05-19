// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/host"

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"go.uber.org/zap"
)

const (
	clusterNameKey          = "container-insight-eks-cluster-name"
	clusterNameTagKeyPrefix = "kubernetes.io/cluster/"
	autoScalingGroupNameTag = "aws:autoscaling:groupName"
)

type ec2TagsClient interface {
	DescribeTags(ctx context.Context, input *ec2.DescribeTagsInput, optFns ...func(options *ec2.Options)) (*ec2.DescribeTagsOutput, error)
}

type ec2TagsProvider interface {
	getClusterName() string
	getAutoScalingGroupName() string
}

type ec2Tags struct {
	refreshInterval       time.Duration
	maxJitterTime         time.Duration
	containerOrchestrator string
	instanceID            string
	client                ec2TagsClient
	clusterName           string
	autoScalingGroupName  string
	isSuccess             chan bool // only used in testing
	logger                *zap.Logger
}

type ec2TagsOption func(*ec2Tags)

func newEC2Tags(ctx context.Context, cfg aws.Config, instanceID, region, containerOrchestrator string,
	refreshInterval time.Duration, logger *zap.Logger, options ...ec2TagsOption,
) ec2TagsProvider {
	_ = "STUB: not implemented"
	return *new(ec2TagsProvider)
}

// stop once we get the cluster name

func (et *ec2Tags) fetchEC2Tags(ctx context.Context) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (et *ec2Tags) getClusterName() string { _ = "STUB: not implemented"; return "" }

func (et *ec2Tags) getAutoScalingGroupName() string { _ = "STUB: not implemented"; return "" }

func (et *ec2Tags) refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

// this will be executed only in testing
