// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package upload // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter/internal/upload"

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.uber.org/zap"
)

type Manager interface {
	Upload(ctx context.Context, data []byte, opts *UploadOptions) error
}

type ManagerOpt func(Manager)

type UploadOptions struct {
	OverrideBucket string
	OverridePrefix string
}

type s3manager struct {
	logger       *zap.Logger
	bucket       string
	builder      *PartitionKeyBuilder
	uploader     *transfermanager.Client
	storageClass s3types.StorageClass
	acl          s3types.ObjectCannedACL
}

var _ Manager = (*s3manager)(nil)

func NewS3Manager(logger *zap.Logger, bucket string, builder *PartitionKeyBuilder, service *s3.Client, storageClass s3types.StorageClass, opts ...ManagerOpt) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

func (sw *s3manager) Upload(ctx context.Context, data []byte, opts *UploadOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Only use ContentEncoding for non-archive formats
// Archive formats store files compressed permanently (like .tar.gz)
// while ContentEncoding is for HTTP transfer compression

// Only set ContentEncoding if we have a non-empty encoding value

func (sw *s3manager) contentBuffer(raw []byte) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithACL(acl s3types.ObjectCannedACL) func(Manager) { _ = "STUB: not implemented"; return nil }
