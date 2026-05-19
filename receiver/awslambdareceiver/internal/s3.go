// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awslambdareceiver/internal"

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// s3API abstracts out the S3 APIs allowing mocking for tests
type s3API interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

// S3Service define services exposed for consumers
type S3Service interface {
	GetReader(ctx context.Context, bucketName, objectKey string) (io.ReadCloser, error)
	ReadObject(ctx context.Context, bucketName, objectKey string) ([]byte, error)
	ListObjects(ctx context.Context, bucketName, continuationToken, prefix string) (*s3.ListObjectsV2Output, error)
	DeleteObject(ctx context.Context, bucketName, objectKey string) error
}

// S3Provider expose contract to get S3Service
type S3Provider interface {
	GetService(ctx context.Context) (S3Service, error)
}

// S3ServiceProvider provides S3Service instances.
type S3ServiceProvider struct{}

func (*S3ServiceProvider) GetService(ctx context.Context) (S3Service, error) {
	_ = "STUB: not implemented"
	return *new(S3Service), nil
}

// s3ServiceClient implements the S3Service
type s3ServiceClient struct {
	api s3API
}

func (s *s3ServiceClient) ReadObject(ctx context.Context, bucketName, objectKey string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// s3 body read error is marked for retrying

func (s *s3ServiceClient) GetReader(ctx context.Context, bucketName, objectKey string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (s *s3ServiceClient) ListObjects(ctx context.Context, bucketName, continuationToken, prefix string) (*s3.ListObjectsV2Output, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3ServiceClient) DeleteObject(ctx context.Context, bucketName, objectKey string) error {
	_ = "STUB: not implemented"
	return nil
}
