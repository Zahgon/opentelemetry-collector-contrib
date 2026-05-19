// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/testutils"

import (
	"k8s.io/client-go/tools/cache"
)

type MockStore struct {
	cache.Store
	WantErr bool
	Cache   map[string]any
}

func (ms *MockStore) GetByKey(id string) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func (ms *MockStore) List() []any { _ = "STUB: not implemented"; return nil }
