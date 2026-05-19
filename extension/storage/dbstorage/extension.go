// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dbstorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"

import (
	"context"
	"database/sql"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"
)

type databaseStorage struct {
	driverName     string
	datasourceName string
	logger         *zap.Logger
	db             *sql.DB
}

// Ensure this storage extension implements the appropriate interface
var _ storage.Extension = (*databaseStorage)(nil)

func newDBStorage(logger *zap.Logger, config *Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

// Start opens a connection to the database
func (ds *databaseStorage) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Log warning about legacy driver usage

// Change legacy driver to a new one

// Try to convert legacy driver options and log errors if any

// Shutdown closes the connection to the database
func (ds *databaseStorage) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// GetClient returns a storage client for an individual component
func (ds *databaseStorage) GetClient(ctx context.Context, kind component.Kind, ent component.ID, name string) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

func kindString(k component.Kind) string { _ = "STUB: not implemented"; return "" }

// not expected
