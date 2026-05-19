// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package classic // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/ibmcloud/classic"

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockProvider struct {
	mock.Mock
}

var _ Provider = (*MockProvider)(nil)

func (m *MockProvider) InstanceMetadata(_ context.Context) (*InstanceMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
