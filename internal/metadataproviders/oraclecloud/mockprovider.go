// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oraclecloud // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/oraclecloud"

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockProvider implements the Provider interface for test purposes.
type MockProvider struct {
	mock.Mock
}

// Metadata mocks the retrieval of OracleCloud instance metadata.
func (m *MockProvider) Metadata(_ context.Context) (*ComputeMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract the mocked ComputeMetadata (if set) from method arguments.
