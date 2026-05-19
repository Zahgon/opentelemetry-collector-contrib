// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ec2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/ec2"

import (
	"context"
	"net/http"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	ec2provider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/aws/ec2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/ec2/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr   = "ec2"
	tagPrefix = "ec2.tag."
)

var _ internal.Detector = (*Detector)(nil)

type ec2ifaceBuilder interface {
	buildClient(ctx context.Context, region string, client *http.Client) (ec2.DescribeTagsAPIClient, error)
}

type ec2ClientBuilder struct{}

func (*ec2ClientBuilder) buildClient(ctx context.Context, region string, client *http.Client) (ec2.DescribeTagsAPIClient, error) {
	_ = "STUB: not implemented"
	return *new(ec2.DescribeTagsAPIClient), nil
}

type Detector struct {
	metadataProvider      ec2provider.Provider
	tagKeyRegexes         []*regexp.Regexp
	logger                *zap.Logger
	rb                    *metadata.ResourceBuilder
	ec2ClientBuilder      ec2ifaceBuilder
	failOnMissingMetadata bool
	tagsFromIMDS          bool
}

func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// Use IMDS: no IAM permissions needed, requires InstanceMetadataTags=enabled on the instance

// Use EC2 DescribeTags API (default): requires ec2:DescribeTags IAM permission

func getClientConfig(ctx context.Context, logger *zap.Logger) *http.Client {
	_ = "STUB: not implemented"
	return nil
}

func fetchEC2Tags(ctx context.Context, svc ec2.DescribeTagsAPIClient, instanceID string, tagKeyRegexes []*regexp.Regexp) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchIMDSTags(ctx context.Context, provider ec2provider.Provider, tagKeyRegexes []*regexp.Regexp) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compileRegexes(cfg Config) ([]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func regexArrayMatch(arr []*regexp.Regexp, val string) bool {
	_ = "STUB: not implemented"
	return false
}
