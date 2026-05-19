// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"go.opentelemetry.io/collector/extension/xextension/storage"
)

const (
	storageKeyFormat = "%s/%s/%s/%s"
)

type storageCheckpointPersister[T any] struct {
	storageClient storage.Client
	defaultValue  T
}

func (s *storageCheckpointPersister[T]) Write(namespace, name, consumerGroup, partitionID string, checkpoint T) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *storageCheckpointPersister[T]) Read(namespace, name, consumerGroup, partitionID string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
