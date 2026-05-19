// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3receiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	ingestedTag    = "otel-collector:status"
	ingestedStatus = "ingested"
)

type ListObjectsV2Pager interface {
	HasMorePages() bool
	NextPage(context.Context, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

type ListObjectsAPI interface {
	NewListObjectsV2Paginator(params *s3.ListObjectsV2Input) ListObjectsV2Pager
}

type SingleObjectAPI interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	GetObjectTagging(ctx context.Context, params *s3.GetObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error)
	PutObjectTagging(ctx context.Context, params *s3.PutObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error)
}

type s3ListObjectsAPIImpl struct {
	client *s3.Client
}

func newS3Client(ctx context.Context, cfg S3DownloaderConfig) (ListObjectsAPI, SingleObjectAPI, error) {
	_ = "STUB: not implemented"
	return *new(ListObjectsAPI), *new(SingleObjectAPI), nil
}

func (api *s3ListObjectsAPIImpl) NewListObjectsV2Paginator(params *s3.ListObjectsV2Input) ListObjectsV2Pager {
	_ = "STUB: not implemented"
	return *new(ListObjectsV2Pager)
}

// retrieveS3Object retrieves S3 object content for a given bucket and key
func retrieveS3Object(ctx context.Context, client SingleObjectAPI, bucket, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tagS3Object tags an S3 object for a given bucket and key
func tagS3Object(ctx context.Context, client SingleObjectAPI, bucket, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// hasIngestedTag checks if an S3 object has the ingested tag set
func hasIngestedTag(ctx context.Context, client SingleObjectAPI, bucket, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
