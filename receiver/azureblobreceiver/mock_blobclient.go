// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	bytes "bytes"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

type mockBlobClient struct {
	mock.Mock
}

// ReadBlob provides a mock function with given fields: ctx, containerName, blobName
func (_m *mockBlobClient) readBlob(ctx context.Context, containerName, blobName string) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (_m *mockBlobClient) listBlobs(ctx context.Context, containerName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteBlob provides a mock function with given fields: ctx, containerName, blobName
func (_m *mockBlobClient) deleteBlob(ctx context.Context, containerName, blobName string) error {
	_ = "STUB: not implemented"
	return nil
}

func newMockBlobClient() *mockBlobClient { _ = "STUB: not implemented"; return nil }
