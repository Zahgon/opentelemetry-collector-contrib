// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package adapter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/xextension/storage"
)

func GetStorageClient(ctx context.Context, host component.Host, storageID *component.ID, componentID component.ID) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

// Make storage immune to component renames that add underscores to the component type.
// This is a workaround for https://github.com/open-telemetry/opentelemetry-collector/issues/14988.

func (r *receiver) setStorageClient(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}
