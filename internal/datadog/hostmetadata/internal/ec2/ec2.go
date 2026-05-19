// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package ec2 contains the AWS EC2 hostname provider
package ec2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/ec2"

import (
	"context"
	"sync"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/provider"
	ec2provider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/aws/ec2"
)

var defaultPrefixes = [3]string{"ip-", "domu", "ec2amaz-"}

type HostInfo struct {
	InstanceID  string
	EC2Hostname string
	EC2Tags     []string
}

// isDefaultHostname checks if a hostname is an EC2 default
func isDefaultHostname(hostname string) bool { _ = "STUB: not implemented"; return false }

// GetHostInfo gets the hostname info from EC2 metadata
func GetHostInfo(ctx context.Context, logger *zap.Logger) (hostInfo *HostInfo) {
	_ = "STUB: not implemented"
	return nil
}

// Check if metadata service is available by trying to retrieve instance ID

func (hi *HostInfo) GetHostname(_ *zap.Logger) string { _ = "STUB: not implemented"; return "" }

var (
	_ source.Provider              = (*Provider)(nil)
	_ provider.ClusterNameProvider = (*Provider)(nil)
)

type Provider struct {
	once     sync.Once
	hostInfo HostInfo

	detector ec2provider.Provider
	logger   *zap.Logger
}

func NewProvider(logger *zap.Logger) (*Provider, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Provider) fillHostInfo(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *Provider) Source(ctx context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// instanceTags gets the EC2 tags for the current instance.
func (p *Provider) instanceTags(ctx context.Context) (*ec2.DescribeTagsOutput, error) {
	_ = "STUB: not implemented"
	// Get EC2 metadata to find the region and instance ID
	return nil, nil
}

// Get the EC2 tags for the instance id.
// Similar to:
// - https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/39dbc1ac8/processor/resourcedetectionprocessor/internal/aws/ec2/ec2.go#L118-L151
// - https://github.com/DataDog/datadog-agent/blob/1b4afdd6a03e8fabcc169b924931b2bb8935dab9/pkg/util/ec2/ec2_tags.go#L104-L134

// clusterNameFromTags gets the AWS EC2 Cluster name from the tags on an EC2 instance.
func clusterNameFromTags(ec2Tags *ec2.DescribeTagsOutput) (string, error) {
	_ = "STUB: not implemented"
	// Similar to:
	// - https://github.com/DataDog/datadog-agent/blob/1b4afdd6a03/pkg/util/ec2/ec2.go#L256-L271
	return "", nil
}

// ClusterName gets the cluster name from an AWS EC2 machine.
func (p *Provider) ClusterName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Provider) HostInfo() *HostInfo { _ = "STUB: not implemented"; return nil }
