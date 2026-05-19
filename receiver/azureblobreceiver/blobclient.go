// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	"bytes"
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"go.uber.org/zap"
)

type blobClient interface {
	readBlob(ctx context.Context, containerName, blobName string) (*bytes.Buffer, error)
	listBlobs(ctx context.Context, containerName string) ([]string, error)
	deleteBlob(ctx context.Context, containerName, blobName string) error
}

type azureBlobClient struct {
	serviceClient *azblob.Client
	logger        *zap.Logger
}

var _ blobClient = (*azureBlobClient)(nil)

func (bc *azureBlobClient) listBlobs(ctx context.Context, containerName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bc *azureBlobClient) readBlob(ctx context.Context, containerName, blobName string) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bc *azureBlobClient) deleteBlob(ctx context.Context, containerName, blobName string) error {
	_ = "STUB: not implemented"
	return nil
}

func newBlobClientFromConnectionString(connectionString string, logger *zap.Logger) (*azureBlobClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newBlobClientFromCredential(storageAccountURL string, cred azcore.TokenCredential, logger *zap.Logger) (*azureBlobClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
