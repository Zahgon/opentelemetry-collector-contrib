// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package proxy // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy"

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
)

const (
	idleConnTimeout                = 30 * time.Second
	remoteProxyMaxIdleConnsPerHost = 2

	awsRegionEnvVar                   = "AWS_REGION"
	awsDefaultRegionEnvVar            = "AWS_DEFAULT_REGION"
	ecsContainerMetadataEnabledEnvVar = "ECS_ENABLE_CONTAINER_METADATA"
	ecsMetadataFileEnvVar             = "ECS_CONTAINER_METADATA_FILE"
)

var getEC2Region = func(ctx context.Context, cfg aws.Config) (string, error) {
	client := imds.NewFromConfig(cfg)
	output, err := client.GetRegion(ctx, &imds.GetRegionInput{})
	if err != nil {
		return "", err
	}
	return output.Region, nil
}

// newAWSConfig creates an AWS config using awsutil.GetAWSConfig, which handles
// region setting, retry configuration, and STS AssumeRole credentials if a role
// ARN is provided. The SDK's built-in EndpointResolverV2 handles partition-aware
// endpoint resolution for all AWS partitions.
var newAWSConfig = func(ctx context.Context, roleArn, region string, log *zap.Logger) (aws.Config, error) {
	settings := &awsutil.AWSSessionSettings{
		Region:     region,
		RoleARN:    roleArn,
		MaxRetries: 2,
	}
	return awsutil.GetAWSConfig(ctx, log, settings)
}

func getAWSConfigSession(ctx context.Context, c *Config, logger *zap.Logger) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

// resolveRegion determines the AWS region using the following priority:
//  1. Config file (c.Region)
//  2. Environment variables (AWS_DEFAULT_REGION, AWS_REGION)
//  3. ECS container metadata file (proxy-specific, not in awsutil)
//  4. EC2 instance metadata service (IMDS)
func resolveRegion(ctx context.Context, c *Config, logger *zap.Logger) (string, error) {
	_ = "STUB: not implemented"
	// 1. Config file takes highest priority
	return "", nil
}

// 2. Environment variables

// 3 & 4. Metadata services (only if not in local mode)

// getRegionFromEnv returns the region from environment variables.
// AWS_DEFAULT_REGION takes precedence over AWS_REGION.
func getRegionFromEnv() string { _ = "STUB: not implemented"; return "" }

// getRegionFromMetadata attempts to get region from ECS metadata first,
// then falls back to EC2 IMDS.
func getRegionFromMetadata(ctx context.Context, logger *zap.Logger) (string, error) {
	_ = "STUB: not implemented"
	// Try ECS metadata first (proxy-specific feature not in awsutil)
	return "", nil
}

// Fall back to EC2 IMDS

// Return combined error

// getRegionFromEC2Metadata fetches region from EC2 Instance Metadata Service.
var getRegionFromEC2Metadata = func(ctx context.Context, logger *zap.Logger) (string, error) {
	tempSettings := &awsutil.AWSSessionSettings{}
	tempCfg, err := awsutil.GetAWSConfig(ctx, logger, tempSettings)
	if err != nil {
		return "", err
	}
	return getEC2Region(ctx, tempCfg)
}

func getRegionFromECSMetadata() (string, error) { _ = "STUB: not implemented"; return "", nil }

// proxyServerTransport configures HTTP transport for TCP Proxy Server.
func proxyServerTransport(cfg *Config) (*http.Transport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If not disabled the transport will add a gzip encoding header
// to requests with no `accept-encoding` header value. The header
// is added after we sign the request which invalidates the
// signature.

// isValidRegion checks that the region is a non-empty string containing only
// lowercase ASCII letters, digits, and hyphens (valid DNS label characters).
func isValidRegion(region string) bool { _ = "STUB: not implemented"; return false }

// getServiceEndpoint returns the AWS service endpoint URL for a given service and region.
// It leverages the STS EndpointResolverV2 (which internally uses awsrulesfn.GetPartition
// covering all 8 AWS partitions) to resolve the correct DNS suffix, then replaces the
// service name in the resolved URL.
func getServiceEndpoint(region, serviceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// The STS endpoint URL has the format https://sts.<region>.<dnsSuffix>.
// Replace the "sts" service prefix with the target service name.
